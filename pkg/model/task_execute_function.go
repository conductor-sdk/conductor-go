package model

import (
	"encoding/json"
	"os"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	log "github.com/sirupsen/logrus"
)

var hostname, _ = os.Hostname()

type TaskExecutionFunction func(t *http_model.Task) (interface{}, error)

type WorkflowValidator func(w *http_model.Workflow) (bool, error)

func GetTaskResultFromTask(task *http_model.Task) *http_model.TaskResult {
	return &http_model.TaskResult{
		TaskId:             task.TaskId,
		WorkflowInstanceId: task.WorkflowInstanceId,
		WorkerId:           hostname,
	}
}

func GetTaskResultFromTaskWithError(t *http_model.Task, err error) *http_model.TaskResult {
	taskResult := GetTaskResultFromTask(t)
	taskResult.Status = task_result_status.FAILED
	taskResult.ReasonForIncompletion = err.Error()
	return taskResult
}

func GetTaskResultFromTaskExecutionOutput(t *http_model.Task, taskExecutionOutput interface{}) (*http_model.TaskResult, error) {
	taskResult, ok := taskExecutionOutput.(*http_model.TaskResult)
	if !ok {
		taskResult := GetTaskResultFromTask(t)
		outputData, err := ConvertToMap(taskExecutionOutput)
		if err != nil {
			return nil, err
		}
		taskResult.OutputData = outputData
		return taskResult, nil
	}
	return taskResult, nil
}

func ConvertToMap(input interface{}) (map[string]interface{}, error) {
	if input == nil {
		return nil, nil
	}
	data, err := json.Marshal(input)
	if err != nil {
		log.Debug(
			"Failed to parse input",
			", reason: ", err.Error(),
		)
		return nil, err
	}
	var parsedInput map[string]interface{}
	json.Unmarshal(data, &parsedInput)
	return parsedInput, nil
}
