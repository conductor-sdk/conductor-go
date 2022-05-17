package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type JoinTask struct {
	task   Task
	joinOn []string
}

func Join(taskRefName string, joinOn []string, inputParameters map[string]interface{}) *JoinTask {
	return &JoinTask{
		task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          JOIN,
			optional:          false,
			inputParameters:   inputParameters,
		},
		joinOn: joinOn,
	}
}

func (task *JoinTask) toWorkflowTask() []http_model.WorkflowTask {
	workflowTasks := task.task.toWorkflowTask()
	workflowTasks[0].JoinOn = task.joinOn
	return workflowTasks
}
