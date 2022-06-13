package model

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

var hostname, _ = os.Hostname()

type ExecuteTaskFunction func(t *Task) (interface{}, error)

type ValidateWorkflowFunction func(w *Workflow) (bool, error)

func NewTaskResultFromTask(task *Task) *TaskResult {
	return &TaskResult{
		TaskId:             task.TaskId,
		WorkflowInstanceId: task.WorkflowInstanceId,
		WorkerId:           hostname,
	}

}

func NewTaskResultFromTaskWithError(t *Task, err error) *TaskResult {
	taskResult := NewTaskResultFromTask(t)
	taskResult.Status = FailedTask
	taskResult.ReasonForIncompletion = err.Error()
	return taskResult
}

func NewTaskResult(taskId string, workflowInstanceId string) *TaskResult {
	return &TaskResult{
		TaskId:             taskId,
		WorkflowInstanceId: workflowInstanceId,
		WorkerId:           hostname,
	}

}

func GetTaskResultFromTaskExecutionOutput(t *Task, taskExecutionOutput interface{}) (*TaskResult, error) {
	taskResult, ok := taskExecutionOutput.(*TaskResult)
	if !ok {
		taskResult = NewTaskResultFromTask(t)
		outputData, err := ConvertToMap(taskExecutionOutput)
		if err != nil {
			return nil, err
		}
		taskResult.OutputData = outputData
		taskResult.Status = CompletedTask
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
