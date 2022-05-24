package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

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

func (task *DynamicTask) toWorkflowTask() []http_model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].DynamicTaskNameParam = dynamicTaskNameParameter
	return workflowTasks
}
