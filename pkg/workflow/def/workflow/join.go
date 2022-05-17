package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type JoinTask struct {
	Task
	joinOn []string
}

func Join(taskRefName string, joinOn ...string) *JoinTask {
	return &JoinTask{
		Task: Task{
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

func (task *JoinTask) toWorkflowTask() []http_model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].JoinOn = task.joinOn
	return workflowTasks
}
