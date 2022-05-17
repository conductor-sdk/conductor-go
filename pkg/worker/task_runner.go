package worker

import (
	"context"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metrics_counter"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metrics_gauge"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

type TaskRunner struct {
	conductorTaskResourceClient *conductor_http_client.TaskResourceApiService
	waitGroup                   sync.WaitGroup
	hostName                    string
}

func NewTaskRunnerWithApiClient(
	apiClient *conductor_http_client.APIClient,
) *TaskRunner {
	return &TaskRunner{
		conductorTaskResourceClient: &conductor_http_client.TaskResourceApiService{
			APIClient: apiClient,
		},
	}
}

func NewTaskRunner(authenticationSettings *settings.AuthenticationSettings, httpSettings *settings.HttpSettings) *TaskRunner {
	apiClient := conductor_http_client.NewAPIClient(
		authenticationSettings,
		httpSettings,
	)
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}
	return &TaskRunner{
		conductorTaskResourceClient: &conductor_http_client.TaskResourceApiService{
			APIClient: apiClient,
		},
		hostName: hostname,
	}
}

func (c *TaskRunner) StartWorkerWithDomain(taskType string, executeFunction model.TaskExecuteFunction,
	threadCount int, pollIntervalInMillis int, domain string) {
	c.startWorker(taskType, executeFunction, threadCount, pollIntervalInMillis, domain)
}
func (c *TaskRunner) StartWorker(taskType string, executeFunction model.TaskExecuteFunction, threadCount int, pollIntervalInMillis int) {
	c.startWorker(taskType, executeFunction, threadCount, pollIntervalInMillis, "")
}

func (c *TaskRunner) startWorker(taskType string, executeFunction model.TaskExecuteFunction, threadCount int, pollIntervalInMillis int, taskDomain string) {
	c.waitGroup.Add(1)
	var domain optional.String
	if taskDomain != "" {
		domain = optional.NewString(taskDomain)
	}
	go c.pollAndExecute(taskType, executeFunction, pollIntervalInMillis, threadCount, domain)
	log.Info(
		"Started worker for task: ", taskType,
		", threadCount / batchSize: ", threadCount,
		", polling interval: ", pollIntervalInMillis, "ms",
	)
}

func (c *TaskRunner) WaitWorkers() {
	c.waitGroup.Wait()
}

func (c *TaskRunner) pollAndExecute(taskType string, executeFunction model.TaskExecuteFunction,
	pollingInterval int, batchSize int, domain optional.String,
) {
	for {
		c.runBatch(taskType, executeFunction, pollingInterval, batchSize, domain)
	}
}

func (c *TaskRunner) runBatch(
	taskType string, executeFunction model.TaskExecuteFunction,
	pollingInterval int, batchSize int,
	domain optional.String,
) {

	tasks := c.batchPoll(taskType, batchSize, pollingInterval, domain)
	if tasks == nil || len(tasks) == 0 {
		sleep(pollingInterval)
		return
	}
	for _, task := range tasks {
		c.waitGroup.Add(1)
		go c.executeAndUpdateTask(taskType, task, executeFunction)
	}
}

func (c *TaskRunner) executeAndUpdateTask(taskType string, task http_model.Task, executeFunction model.TaskExecuteFunction) {
	taskResult := c.executeTask(&task, executeFunction)
	c.updateTask(taskType, taskResult)
}

func (c *TaskRunner) pollTask(taskType string) *http_model.Task {
	log.Debug("Polling for ", taskType)
	metrics_counter.IncrementTaskPoll(taskType)
	startTime := time.Now()

	task, response, err := c.conductorTaskResourceClient.Poll(
		context.Background(),
		taskType,
		nil,
	)
	spentTime := time.Since(startTime)
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
	log.Debug("Polled task: ", task)
	return &task
}

func (c *TaskRunner) batchPoll(taskType string, count int, pollingInterval int, domain optional.String) []http_model.Task {
	log.Debug("Polling for ", taskType, ", batchSize, ", count)
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
	log.Debug("Polled task: ", len(tasks))
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
		amount := 2 ^ i
		time.Sleep(time.Duration(amount) * time.Second)
	}
}

func (c *TaskRunner) _updateTask(taskType string, taskResult *http_model.TaskResult) error {
	_, response, err := c.conductorTaskResourceClient.UpdateTask(context.Background(), taskResult)
	if err != nil {
		log.Error(
			"Error on task update. taskResult: ", *taskResult,
			", error: ", err.Error(),
			", response: ", response,
		)
		metrics_counter.IncrementTaskUpdateError(taskType, err)
		return err
	}
	log.Debug("Updated task: ", *taskResult)
	return nil
}

func sleep(pollingInterval int) {
	time.Sleep(
		time.Duration(pollingInterval) * time.Millisecond,
	)
}
