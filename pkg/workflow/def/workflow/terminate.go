package workflow

import "github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"

type TerminateTask struct {
	Task
}

func NewTerminateTask(taskRefName string, status workflow_status.WorkflowStatus, terminationReason string) *TerminateTask {
	return &TerminateTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          TERMINATE,
			optional:          false,
			inputParameters: map[string]interface{}{
				"terminationStatus": status,
				"terminationReason": terminationReason,
			},
		},
	}
}
func (task *TerminateTask) Description(description string) *TerminateTask {
	task.Task.Description(description)
	return task
}
func (task *TerminateTask) Optional(optional bool) *TerminateTask {
	task.Task.Optional(optional)
	return task
}
func (task *TerminateTask) Input(key string, value interface{}) *TerminateTask {
	task.Task.Input(key, value)
	return task
}
func (task *TerminateTask) InputMap(inputMap map[string]interface{}) *TerminateTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
