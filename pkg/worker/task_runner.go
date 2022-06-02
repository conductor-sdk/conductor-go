package worker

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/concurrency"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metrics_counter"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metrics_gauge"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	log "github.com/sirupsen/logrus"
)

type TaskRunner struct {
	conductorTaskResourceClient *conductor_http_client.TaskResourceApiService
	hostName                    string

	maxAllowedWorkersByTaskType map[string]int
	runningWorkersByTaskType    map[string]int
	mutex                       sync.Mutex
	workerWaitGroup             sync.WaitGroup
}

func NewTaskRunner(authenticationSettings *settings.AuthenticationSettings, httpSettings *settings.HttpSettings) *TaskRunner {
	apiClient := conductor_http_client.NewAPIClient(
		authenticationSettings,
		httpSettings,
	)
	return NewTaskRunnerWithApiClient(apiClient)
}

func NewTaskRunnerWithApiClient(
	apiClient *conductor_http_client.APIClient,
) *TaskRunner {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	return &TaskRunner{
		conductorTaskResourceClient: &conductor_http_client.TaskResourceApiService{
			APIClient: apiClient,
		},
		hostName:                    hostname,
		maxAllowedWorkersByTaskType: make(map[string]int),
		runningWorkersByTaskType:    make(map[string]int),
	}
}

// StartWorkerWithDomain
//  - taskType Task Type to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize No. of tasks to poll for.  Each polled task is executed in a goroutine.  Batching improves the throughput
//  - pollInterval Time to wait for between polls if there are no tasks available.  Reduces excessive polling on the server when there is no work
//  - domain Task domain. Optional for polling
func (c *TaskRunner) StartWorkerWithDomain(taskType string, executeFunction model.TaskExecuteFunction, threadCount int, pollInterval time.Duration, domain string) error {
	return c.startWorker(taskType, executeFunction, threadCount, pollInterval, domain)
}

// StartWorker
//  - taskType Task Type to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize No. of tasks to poll for.  Each polled task is executed in a goroutine.  Batching improves the throughput
//  - pollInterval Time to wait for between polls if there are no tasks available.  Reduces excessive polling on the server when there is no work
func (c *TaskRunner) StartWorker(taskType string, executeFunction model.TaskExecuteFunction, batchSize int, pollInterval time.Duration) error {
	return c.startWorker(taskType, executeFunction, batchSize, pollInterval, "")
}

func (c *TaskRunner) RemoveWorker(taskType string, threadCount int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if threadCount >= c.maxAllowedWorkersByTaskType[taskType] {
		c.maxAllowedWorkersByTaskType[taskType] = 0
	} else {
		c.maxAllowedWorkersByTaskType[taskType] -= threadCount
	}
	return nil
}

func (c *TaskRunner) WaitWorkers() {
	c.workerWaitGroup.Wait()
}

func (c *TaskRunner) startWorker(taskType string, executeFunction model.TaskExecuteFunction, threadCount int, pollInterval time.Duration, taskDomain string) error {
	var domain optional.String
	if taskDomain != "" {
		domain = optional.NewString(taskDomain)
	}
	previousMaxAllowedWorkers, err := c.getMaxAllowedWorkers(taskType)
	if err != nil {
		return err
	}
	err = c.increaseMaxAllowedWorkers(taskType, threadCount)
	if err != nil {
		return err
	}
	if previousMaxAllowedWorkers == 0 {
		c.workerWaitGroup.Add(1)
		go c.pollAndExecute(taskType, executeFunction, pollInterval, domain)
	}
	log.Debug(
		"Started worker for task: ", taskType,
		", threadCount / batchSize: ", threadCount,
		", polling interval: ", pollInterval.Milliseconds(), "ms",
	)
	return nil
}

func (c *TaskRunner) pollAndExecute(taskType string, executeFunction model.TaskExecuteFunction, pollInterval time.Duration, domain optional.String) error {
	defer concurrency.HandlePanicError("poll_and_execute")
	for c.isWorkerAlive(taskType) {
		isTaskQueueEmpty, err := c.runBatch(taskType, executeFunction, pollInterval, domain)
		if err != nil {
			log.Warning(
				"Failed to poll and execute",
				", reason: ", err.Error(),
				", taskType: ", taskType,
				", pollInterval: ", pollInterval.Milliseconds(), "ms",
				", domain: ", domain,
			)
		} else if isTaskQueueEmpty {
			log.Debug("No tasks available for: ", taskType)
			time.Sleep(pollInterval)
		}
	}
	c.workerWaitGroup.Done()
	return nil
}

func (c *TaskRunner) runBatch(taskType string, executeFunction model.TaskExecuteFunction, pollInterval time.Duration, domain optional.String) (bool, error) {
	batchSize, err := c.getAvailableWorkerAmount(taskType)
	if err != nil {
		return false, err
	}
	if batchSize < 1 {
		// TODO wait until there is available workers
		time.Sleep(1 * time.Millisecond)
		return false, nil
	}
	tasks, err := c.batchPoll(taskType, batchSize, pollInterval, domain)
	if err != nil {
		return false, err
	}
	if len(tasks) < 1 {
		return true, nil
	}
	c.increaseRunningWorkers(taskType, len(tasks))
	for _, task := range tasks {
		go c.executeAndUpdateTask(taskType, task, executeFunction)
	}
	return false, nil
}

func (c *TaskRunner) executeAndUpdateTask(taskType string, task http_model.Task, executeFunction model.TaskExecuteFunction) error {
	defer concurrency.HandlePanicError("execute_and_update_task")
	taskResult, err := c.executeTask(&task, executeFunction)
	if err != nil {
		return err
	}
	err = c.updateTask(taskType, taskResult)
	if err != nil {
		return err
	}
	return c.runningWorkerDone(taskType)
}

func (c *TaskRunner) batchPoll(taskType string, count int, timeout time.Duration, domain optional.String) ([]http_model.Task, error) {
	log.Debug(
		"Polling for task: ", taskType,
		", in batches of size: ", count,
	)
	metrics_counter.IncrementTaskPoll(taskType)
	startTime := time.Now()
	tasks, response, err := c.conductorTaskResourceClient.BatchPoll(
		context.Background(),
		taskType,
		&conductor_http_client.TaskResourceApiBatchPollOpts{
			Domain:   domain,
			Workerid: optional.NewString(c.hostName),
			Count:    optional.NewInt32(int32(count)),
			Timeout:  optional.NewInt32(int32(timeout.Milliseconds())),
		},
	)
	spentTime := time.Since(startTime)
	metrics_gauge.RecordTaskPollTime(
		taskType,
		spentTime.Seconds(),
	)
	if err != nil {
		metrics_counter.IncrementTaskPollError(
			taskType, err,
		)
		return nil, err
	}
	if response.StatusCode == 204 {
		return nil, nil
	}
	log.Debug("Polled tasks: ", len(tasks), " for taskType ", taskType)
	return tasks, nil
}

func (c *TaskRunner) executeTask(t *http_model.Task, executeFunction model.TaskExecuteFunction) (*http_model.TaskResult, error) {
	startTime := time.Now()
	taskResult, err := executeFunction(t)
	spentTime := time.Since(startTime)
	metrics_gauge.RecordTaskExecuteTime(
		t.TaskDefName, float64(spentTime.Milliseconds()),
	)
	if err != nil {
		taskResult.Status = task_result_status.FAILED
		taskResult.ReasonForIncompletion = err.Error()
		metrics_counter.IncrementTaskExecuteError(
			t.TaskDefName, err,
		)
		return nil, err
	}
	if taskResult == nil {
		return nil, fmt.Errorf("task result cannot be nil")
	}
	log.Trace(fmt.Sprintf("Polled task: %+v", *t))
	return taskResult, nil
}

func (c *TaskRunner) updateTask(taskType string, taskResult *http_model.TaskResult) error {
	retryCount := 3
	for i := 0; i < retryCount; i++ {
		err := c._updateTask(taskType, taskResult)
		if err == nil {
			return nil
		}
		amount := (1 << i)
		time.Sleep(time.Duration(amount) * time.Second)
	}
	return fmt.Errorf("failed to update task %s after %d attempts", taskType, retryCount)
}

func (c *TaskRunner) _updateTask(taskType string, taskResult *http_model.TaskResult) error {
	startTime := time.Now()
	_, response, err := c.conductorTaskResourceClient.UpdateTask(context.Background(), taskResult)
	spentTime := time.Since(startTime)
	metrics_gauge.RecordTaskUpdateTime(
		taskType, float64(spentTime.Milliseconds()),
	)
	if err != nil {
		metrics_counter.IncrementTaskUpdateError(taskType, err)
		log.Debug(
			"Failed to update task",
			", reason: ", err.Error(),
			", task type: ", taskType,
			", task result: ", *taskResult,
			", response: ", response,
		)
		return err
	}
	log.Debug("Updated task: ", taskResult.TaskId, ", status: ", taskResult.Status)
	return nil
}

func (c *TaskRunner) getAvailableWorkerAmount(taskType string) (int, error) {
	allowed, err := c.getMaxAllowedWorkers(taskType)
	if err != nil {
		return -1, err
	}
	running, err := c.getRunningWorkers(taskType)
	if err != nil {
		return -1, err
	}
	return allowed - running, nil
}

func (c *TaskRunner) getMaxAllowedWorkers(taskType string) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	amount, ok := c.maxAllowedWorkersByTaskType[taskType]
	if !ok {
		return 0, nil
	}
	return amount, nil
}

func (c *TaskRunner) getRunningWorkers(taskType string) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	amount, ok := c.runningWorkersByTaskType[taskType]
	if !ok {
		return 0, nil
	}
	return amount, nil
}

func (c *TaskRunner) isWorkerAlive(taskType string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	allowed, ok := c.maxAllowedWorkersByTaskType[taskType]
	return ok && allowed > 0
}

func (c *TaskRunner) increaseRunningWorkers(taskType string, amount int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.runningWorkersByTaskType[taskType] += amount
	log.Trace("Increased running workers for task: ", taskType, ", by: ", amount)
	return nil
}

func (c *TaskRunner) runningWorkerDone(taskType string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.runningWorkersByTaskType[taskType] -= 1
	log.Trace("Running worker done for task: ", taskType)
	return nil
}

func (c *TaskRunner) increaseMaxAllowedWorkers(taskType string, threadCount int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.maxAllowedWorkersByTaskType[taskType] += threadCount
	log.Debug("Increased max allowed workers of task: ", taskType, ", by: ", threadCount)
	return nil
}
