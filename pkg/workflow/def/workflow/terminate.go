package workflow

import (
	"github.com/conductor-sdk/conductor-go/pkg/model"
)

type TerminateTask struct {
	Task
}

func NewTerminateTask(taskRefName string, status model.WorkflowStatus, terminationReason string) *TerminateTask {
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

// Description of the task
func (task *TerminateTask) Description(description string) *TerminateTask {
	task.Task.Description(description)
	return task
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *TerminateTask) Input(key string, value interface{}) *TerminateTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *TerminateTask) InputMap(inputMap map[string]interface{}) *TerminateTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
