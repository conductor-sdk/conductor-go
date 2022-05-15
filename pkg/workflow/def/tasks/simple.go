package tasks

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

func Simple(name string, taskRefName string) *simpleTask {
	return &simpleTask{task{
		name:              name,
		taskReferenceName: taskRefName,
		description:       "",
		taskType:          SIMPLE,
		optional:          false,
		inputParameters:   map[string]interface{}{},
	}}
}

type simpleTask struct {
	task
}

func (task *simpleTask) Description(description string) *simpleTask {
	task.task.Description(description)
	return task
}

func (task *simpleTask) Optional(optional bool) *simpleTask {
	task.task.Optional(optional)
	return task
}

func (task *simpleTask) ToWorkflowTask() *http_model.WorkflowTask {
	return task.task.ToWorkflowTask()
}
