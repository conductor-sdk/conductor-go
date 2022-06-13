package def

import (
	"github.com/conductor-sdk/conductor-go/pkg/model"
)

type JoinTask struct {
	Task
	joinOn []string
}

func NewJoinTask(taskRefName string, joinOn ...string) *JoinTask {
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

func (task *JoinTask) toWorkflowTask() []model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].JoinOn = task.joinOn
	return workflowTasks
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *JoinTask) Optional(optional bool) *JoinTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *JoinTask) Description(description string) *JoinTask {
	task.Task.Description(description)
	return task
}
