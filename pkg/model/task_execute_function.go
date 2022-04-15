package model

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type TaskExecuteFunction func(t *http_model.Task) (*http_model.TaskResult, error)

func GetTaskResultFromTask(task *http_model.Task) *http_model.TaskResult {
	return &http_model.TaskResult{
		TaskId:             task.TaskId,
		WorkflowInstanceId: task.WorkflowInstanceId,
		WorkerId:           "TESTING...",
	}
}
