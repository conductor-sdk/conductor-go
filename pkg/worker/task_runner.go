package worker

import (
	"context"
	"sync"
	"time"

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
	waitGroup                   sync.WaitGroup
}

func NewWorkerOrkestratorWithApiClient(
	apiClient *conductor_http_client.APIClient,
) *TaskRunner {
	return &TaskRunner{
		conductorTaskResourceClient: &conductor_http_client.TaskResourceApiService{
			APIClient: apiClient,
		},
	}
}

func NewTaskRunner(
	authenticationSettings *settings.AuthenticationSettings,
	httpSettings *settings.HttpSettings,
) *TaskRunner {
	apiClient := conductor_http_client.NewAPIClient(
		authenticationSettings,
		httpSettings,
	)
	return &TaskRunner{
		conductorTaskResourceClient: &conductor_http_client.TaskResourceApiService{
			APIClient: apiClient,
		},
	}
}

func (c *TaskRunner) StartWorker(taskType string, executeFunction model.TaskExecuteFunction, threadCount int, pollIntervalInMillis int) {
	for goRoutines := 1; goRoutines <= threadCount; goRoutines++ {
		c.waitGroup.Add(1)
		go c.run(taskType, executeFunction, pollIntervalInMillis)
	}
	log.Debug(
		"Started worker for task: ", taskType,
		", go routines amount: ", threadCount,
		", polling interval: ", pollIntervalInMillis, "ms",
	)
}

func (c *TaskRunner) WaitWorkers() {
	c.waitGroup.Wait()
}

func (c *TaskRunner) run(taskType string, executeFunction model.TaskExecuteFunction, pollingInterval int) {
	for {
		c.runOnce(taskType, executeFunction, pollingInterval)
	}
	// c.waitGroup.Done()
}

func (c *TaskRunner) runOnce(taskType string, executeFunction model.TaskExecuteFunction, pollingInterval int) {
	task := c.pollTask(taskType)
	if task == nil {
		sleep(pollingInterval)
		return
	}
	taskResult := c.executeTask(task, executeFunction)
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
	log.Debug("Executed task: ", *t)
	return taskResult
}

func (c *TaskRunner) updateTask(taskType string, taskResult *http_model.TaskResult) {
	_, response, err := c.conductorTaskResourceClient.UpdateTask(
		taskType,
		context.Background(),
		taskResult,
	)
	if err != nil {
		log.Error(
			"Error on task update. taskResult: ", *taskResult,
			", error: ", err.Error(),
			", response: ", response,
		)
		metrics_counter.IncrementTaskUpdateError(taskType, err)
		return
	}
	log.Debug("Updated task: ", *taskResult)
}

func sleep(pollingInterval int) {
	time.Sleep(
		time.Duration(pollingInterval) * time.Millisecond,
	)
}
