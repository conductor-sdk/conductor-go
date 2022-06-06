package model

import (
	"os"
)

var hostname, _ = os.Hostname()

type TaskExecuteFunction func(t *Task) (*TaskResult, error)

type TaskExecuteFunction2 func(t *interface{}) (*interface{}, error)

func GetTaskResultFromTask(task *Task) *TaskResult {
	return &TaskResult{
		TaskId:             task.TaskId,
		WorkflowInstanceId: task.WorkflowInstanceId,
		WorkerId:           hostname,
	}

}
