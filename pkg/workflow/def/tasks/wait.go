package tasks

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

func Wait(taskRefName string) *decision {
	return &decision{
		task: task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          WAIT,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
}

type wait struct {
	task
}

func (task *wait) Description(description string) *wait {
	task.task.Description(description)
	return task
}

func (task *wait) Optional(optional bool) *wait {
	task.task.Optional(optional)
	return task
}

func (task *wait) ToWorkflowTask() *http_model.WorkflowTask {
	return task.task.ToWorkflowTask()
}
