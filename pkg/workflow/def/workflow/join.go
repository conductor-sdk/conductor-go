package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

func Join(taskRefName string, joinOn ...string) *join {
	return &join{
		task: task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          JOIN,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		joinOn: joinOn,
	}
}

type join struct {
	task
	joinOn []string
}

func (task *join) Description(description string) *join {
	task.task.Description(description)
	return task
}

func (task *join) Optional(optional bool) *join {
	task.task.Optional(optional)
	return task
}

func (task *join) toWorkflowTask() *[]http_model.WorkflowTask {
	workflowTasks := task.task.toWorkflowTask()
	(*workflowTasks)[0].JoinOn = task.joinOn
	return workflowTasks
}
