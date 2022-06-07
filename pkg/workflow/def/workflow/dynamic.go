package workflow

import (
	"github.com/conductor-sdk/conductor-go/pkg/model"
)

type DynamicTask struct {
	Task
	joinOn []string
}

const dynamicTaskNameParameter = "taskToExecute"

// NewDynamicTask
//  - taskRefName Reference name for the task.  MUST be unique within the workflow
//  - taskNameParameter Parameter that contains the expression for the dynamic task name.  e.g. ${workflow.input.dynamicTask}
func NewDynamicTask(taskRefName string, taskNameParameter string) *DynamicTask {
	return &DynamicTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          DYNAMIC,
			optional:          false,
			inputParameters: map[string]interface{}{
				dynamicTaskNameParameter: taskNameParameter,
			},
		},
	}
}

func (task *DynamicTask) toWorkflowTask() []model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].DynamicTaskNameParam = dynamicTaskNameParameter
	return workflowTasks
}

// Input to the task
func (task *DynamicTask) Input(key string, value interface{}) *DynamicTask {
	task.Task.Input(key, value)
	return task
}
func (task *DynamicTask) InputMap(inputMap map[string]interface{}) *DynamicTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
func (task *DynamicTask) Optional(optional bool) *DynamicTask {
	task.Task.Optional(optional)
	return task
}
func (task *DynamicTask) Description(description string) *DynamicTask {
	task.Task.Description(description)
	return task
}
