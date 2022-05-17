package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type SimpleTask struct {
	task Task
}

func Simple(name string, taskRefName string) *SimpleTask {
	return &SimpleTask{
		task: Task{
			name:              name,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SIMPLE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
}

func (task *SimpleTask) toWorkflowTask() []http_model.WorkflowTask {
	return task.task.toWorkflowTask()
}
