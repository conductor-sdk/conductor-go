package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type join struct {
	task
	joinOn []string
}

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

func (task *join) toWorkflowTask() []http_model.WorkflowTask {
	workflowTasks := task.task.toWorkflowTask()
	workflowTasks[0].JoinOn = task.joinOn
	return workflowTasks
}
