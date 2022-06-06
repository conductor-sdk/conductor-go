package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

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

func (task *JoinTask) toWorkflowTask() []http_model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].JoinOn = task.joinOn
	return workflowTasks
}

// Input to the task
func (task *JoinTask) Input(key string, value interface{}) *JoinTask {
	task.Task.Input(key, value)
	return task
}
func (task *JoinTask) InputMap(inputMap map[string]interface{}) *JoinTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
func (task *JoinTask) Optional(optional bool) *JoinTask {
	task.Task.Optional(optional)
	return task
}
func (task *JoinTask) Description(description string) *JoinTask {
	task.Task.Description(description)
	return task
}
