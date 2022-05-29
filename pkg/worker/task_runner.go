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

func (c *TaskRunner) StartWorkerWithDomain(taskType string, executeFunction model.TaskExecuteFunction, threadCount int, pollIntervalInMillis int, domain string) {
	c.startWorker(taskType, executeFunction, threadCount, pollIntervalInMillis, domain)
}

// StartWorker
//  - taskType Task Type to poll and execute the work
//  - executeFunction Task execution function
//  - batchSize No. of tasks to poll for.  Each polled task is executed in a goroutine.  Batching improves the throughput
//  - pollIntervalInMillis Time to wait for between polls if there are no tasks available.  Reduces excessive polling on the server when there is no work
func (c *TaskRunner) StartWorker(taskType string, executeFunction model.TaskExecuteFunction, batchSize int, pollIntervalInMillis int) {
	c.startWorker(taskType, executeFunction, batchSize, pollIntervalInMillis, "")
}

func (c *TaskRunner) RemoveWorker(taskType string, threadCount int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if threadCount >= c.maxAllowedWorkersByTaskType[taskType] {
		c.maxAllowedWorkersByTaskType[taskType] = 0
	} else {
		c.maxAllowedWorkersByTaskType[taskType] -= threadCount
	}
}

func (c *TaskRunner) WaitWorkers() {
	c.workerWaitGroup.Wait()
}

func (c *TaskRunner) startWorker(taskType string, executeFunction model.TaskExecuteFunction, threadCount int, pollIntervalInMillis int, taskDomain string) {
	var domain optional.String
	if taskDomain != "" {
		domain = optional.NewString(taskDomain)
	}
	previousMaxAllowedWorkers := c.getMaxAllowedWorkers(taskType)
	c.increaseMaxAllowedWorkers(taskType, threadCount)
	if previousMaxAllowedWorkers == 0 {
		c.workerWaitGroup.Add(1)
		go c.pollAndExecute(taskType, executeFunction, pollIntervalInMillis, domain)
	}
	log.Info(
		"Started worker for task: ", taskType,
		", threadCount / batchSize: ", threadCount,
		", polling interval: ", pollIntervalInMillis, "ms",
	)
}

func (c *TaskRunner) pollAndExecute(taskType string, executeFunction model.TaskExecuteFunction, pollingInterval int, domain optional.String) {
	defer func() {
		c.workerWaitGroup.Done()
		concurrency.OnError("poll_and_execute")
		log.Warning(
			"Panic at pollAndExecute",
			", taskType: ", taskType,
			", pollingInterval: ", pollingInterval,
			", domain: ", domain.Value(),
		)
	}()
	for c.isWorkerAlive(taskType) {
		c.runBatch(taskType, executeFunction, pollingInterval, domain)
	}
}

func (c *TaskRunner) runBatch(taskType string, executeFunction model.TaskExecuteFunction, pollingInterval int, domain optional.String) {
	batchSize := c.getAvailableWorkerAmount(taskType)
	if batchSize < 1 {
		sleep(1)
		return
	}
	tasks := c.batchPoll(taskType, batchSize, pollingInterval, domain)
	if len(tasks) < 1 {
		log.Debug("No tasks available for: ", taskType)
		sleep(pollingInterval)
		return
	}
	c.increaseRunningWorkers(taskType, len(tasks))
	for _, task := range tasks {
		go c.executeAndUpdateTask(taskType, task, executeFunction)
	}
}

func (c *TaskRunner) executeAndUpdateTask(taskType string, task http_model.Task, executeFunction model.TaskExecuteFunction) {
	defer func() {
		c.runningWorkerDone(taskType)
		concurrency.OnError(
			fmt.Sprintf("executeAndUpdateTask, taskType: %s, task: %s",
				taskType,
				fmt.Sprint(task),
			),
		)
	}()
	taskResult := c.executeTask(&task, executeFunction)
	c.updateTask(taskType, taskResult)
}

func (c *TaskRunner) batchPoll(taskType string, count int, pollingInterval int, domain optional.String) []http_model.Task {
	log.Debug("Polling for task: ", taskType, ", in batches of size: ", count)
	metrics_counter.IncrementTaskPoll(taskType)
	startTime := time.Now()
	tasks, response, err := c.conductorTaskResourceClient.BatchPoll(
		context.Background(),
		taskType,
		&conductor_http_client.TaskResourceApiBatchPollOpts{
			Domain:   domain,
			Workerid: optional.NewString(c.hostName),
			Count:    optional.NewInt32(int32(count)),
			Timeout:  optional.NewInt32(int32(pollingInterval)),
		},
	)
	spentTime := time.Since(startTime)
	log.Debug("Task Poll Time ", spentTime.Milliseconds())
	metrics_gauge.RecordTaskPollTime(
		taskType,
		spentTime.Seconds(),
	)
	if err != nil {
		log.Error(
			"Error polling for task: ", taskType,
			", error: ", err.Error(),
		)
		metrics_counter.IncrementTaskPollError(
			taskType, err,
		)
		return nil
	}
	if response.StatusCode == 204 {
		return nil
	}
	log.Debug("Polled tasks: ", len(tasks), " for taskType ", taskType)
	return tasks
}

func (c *TaskRunner) executeTask(t *http_model.Task, executeFunction model.TaskExecuteFunction) *http_model.TaskResult {
	startTime := time.Now()
	taskResult, err := executeFunction(t)
	spentTime := time.Since(startTime)
	metrics_gauge.RecordTaskExecuteTime(
		t.TaskDefName, spentTime.Seconds(),
	)
	if taskResult == nil {
		log.Error("TaskResult cannot be nil: ", t.TaskId)
		return nil
	}
	if err != nil {
		log.Error("Error Executing task:", err.Error())
		taskResult.Status = task_result_status.FAILED
		taskResult.ReasonForIncompletion = err.Error()
		metrics_counter.IncrementTaskExecuteError(
			t.TaskDefName, err,
		)
	}
	log.Debug("Executed task: ", (*t).TaskId)
	return taskResult
}

func (c *TaskRunner) updateTask(taskType string, taskResult *http_model.TaskResult) {
	retryCount := 3
	for i := 0; i < retryCount; i++ {
		err := c._updateTask(taskType, taskResult)
		if err == nil {
			return
		}
		amount := (1 << i)
		time.Sleep(time.Duration(amount) * time.Second)
	}
}

func (c *TaskRunner) _updateTask(taskType string, taskResult *http_model.TaskResult) error {
	startTime := time.Now()
	_, response, err := c.conductorTaskResourceClient.UpdateTask(context.Background(), taskResult)
	spentTime := time.Since(startTime)
	log.Debug("Task Update Time ", spentTime.Milliseconds())
	if err != nil {
		log.Error(
			"Error on task update. taskResult: ", *taskResult,
			", error: ", err.Error(),
			", response: ", response,
		)
		metrics_counter.IncrementTaskUpdateError(taskType, err)
		return err
	}
	log.Debug("Updated task: ", (*taskResult).TaskId, ",", (*taskResult).Status)
	return nil
}

func sleep(pollingInterval int) {
	time.Sleep(
		time.Duration(pollingInterval) * time.Millisecond,
	)
}

func (c *TaskRunner) increaseMaxAllowedWorkers(taskType string, threadCount int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.maxAllowedWorkersByTaskType[taskType] += threadCount
	log.Debug("Increased max allowed workers of task: ", taskType, ", by: ", threadCount)
}

func (c *TaskRunner) getAvailableWorkerAmount(taskType string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.maxAllowedWorkersByTaskType[taskType] - c.runningWorkersByTaskType[taskType]
}

func (c *TaskRunner) getMaxAllowedWorkers(taskType string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.maxAllowedWorkersByTaskType[taskType]
}

func (c *TaskRunner) increaseRunningWorkers(taskType string, amount int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.runningWorkersByTaskType[taskType] += amount
	log.Trace("Increased running workers of task: ", taskType, ", by: ", amount)
}

func (c *TaskRunner) runningWorkerDone(taskType string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.runningWorkersByTaskType[taskType] -= 1
	log.Debug("Running worker done for task: ", taskType)
}

func (c *TaskRunner) isWorkerAlive(taskType string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.maxAllowedWorkersByTaskType[taskType] > 0
}
