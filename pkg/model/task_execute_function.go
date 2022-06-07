package model

import (
	"encoding/json"
	"os"

	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	log "github.com/sirupsen/logrus"
)

var hostname, _ = os.Hostname()

type ExecuteTaskFunction func(t *Task) (interface{}, error)

type ValidateWorkflowFunction func(w *Workflow) (bool, error)

func GetTaskResultFromTask(task *Task) *TaskResult {
	return &TaskResult{
		TaskId:             task.TaskId,
		WorkflowInstanceId: task.WorkflowInstanceId,
		WorkerId:           hostname,
	}

}

func GetTaskResultFromTaskWithError(t *Task, err error) *TaskResult {
	taskResult := GetTaskResultFromTask(t)
	taskResult.Status = task_result_status.FAILED
	taskResult.ReasonForIncompletion = err.Error()
	return taskResult
}

func GetTaskResultFromTaskExecutionOutput(t *Task, taskExecutionOutput interface{}) (*TaskResult, error) {
	taskResult, ok := taskExecutionOutput.(*TaskResult)
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
