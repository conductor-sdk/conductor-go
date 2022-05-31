package worker

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metrics_counter"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metrics_gauge"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	log "github.com/sirupsen/logrus"
)

const updateTaskRetryAttempts = 3

var hostname, _ = os.Hostname()

func batchPoll(taskType string, count int, timeout time.Duration, domain optional.String, conductorClient *conductor_http_client.TaskResourceApiService) ([]http_model.Task, error) {
	log.Debug("Polling for task: ", taskType, ", in batches of size: ", count)
	metrics_counter.IncrementTaskPoll(taskType)
	startTime := time.Now()
	tasks, response, err := conductorClient.BatchPoll(
		context.Background(),
		taskType,
		&conductor_http_client.TaskResourceApiBatchPollOpts{
			Domain:   domain,
			Workerid: optional.NewString(hostname),
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
	log.Debug("Polled ", len(tasks), " tasks of taskType ", taskType)
	return tasks, nil
}

func executeTask(t *http_model.Task, executeFunction model.TaskExecuteFunction) (*http_model.TaskResult, error) {
	log.Debug(
		"Executing task",
		", taskType: ", t.TaskType,
		", workflowId: ", t.WorkflowInstanceId,
		", taskId: ", t.TaskId,
	)
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
	log.Debug(
		"Executed task",
		", taskType: ", t.TaskType,
		", workflowId: ", t.WorkflowInstanceId,
		", taskId: ", t.TaskId,
	)
	return taskResult, nil
}

func updateTaskWithRetry(taskType string, taskResult *http_model.TaskResult, conductorClient *conductor_http_client.TaskResourceApiService) error {
	log.Debug(
		"Updating task with retry",
		", taskType: ", taskType,
		", workflowId: ", taskResult.WorkflowInstanceId,
		", taskId: ", taskResult.TaskId,
	)
	for attempts := 0; attempts < updateTaskRetryAttempts; attempts += 1 {
		response, err := updateTask(taskType, taskResult, conductorClient)
		if err == nil {
			log.Debug(
				"Updated task",
				", taskType: ", taskType,
				", workflowId: ", taskResult.WorkflowInstanceId,
				", taskId: ", taskResult.TaskId,
			)
			return nil
		}
		log.Debug(
			"Failed to update task. Reason: ", err.Error(),
			", attempt: ", attempts+1, " of ", updateTaskRetryAttempts,
			", taskType: ", taskType,
			", taskResult: ", fmt.Sprintf("%+v", taskResult),
			", response: ", response,
		)
		amount := (1 << attempts)
		time.Sleep(time.Duration(amount) * time.Second)
	}
	return fmt.Errorf("failed to update task after %d attempts", updateTaskRetryAttempts)
}

func updateTask(taskType string, taskResult *http_model.TaskResult, conductorClient *conductor_http_client.TaskResourceApiService) (*http.Response, error) {
	startTime := time.Now()
	_, response, err := conductorClient.UpdateTask(context.Background(), taskResult)
	spentTime := time.Since(startTime)
	metrics_gauge.RecordTaskUpdateTime(
		taskType, float64(spentTime.Milliseconds()),
	)
	if err != nil {
		metrics_counter.IncrementTaskUpdateError(taskType, err)
		return response, err
	}
	return response, nil
}
