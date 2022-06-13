package def

import (
	"github.com/conductor-sdk/conductor-go/pkg/model"
)

type StartWorkflowTask struct {
	Task
}

func NewStartWorkflowTask(taskRefName string, workflowName string, version *int32, startWorkflowRequest *model.StartWorkflowRequest) *StartWorkflowTask {
	return &StartWorkflowTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          START_WORKFLOW,
			optional:          false,
			inputParameters: map[string]interface{}{
				"startWorkflow": map[string]interface{}{
					"name":          workflowName,
					"version":       version,
					"input":         startWorkflowRequest.Input,
					"correlationId": startWorkflowRequest.CorrelationId,
				},
			},
		},
	}
}

func (task *StartWorkflowTask) toWorkflowTask() []model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	return workflowTasks
}

// Description of the task
func (task *StartWorkflowTask) Description(description string) *StartWorkflowTask {
	task.Task.Description(description)
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *StartWorkflowTask) Optional(optional bool) *StartWorkflowTask {
	task.Task.Optional(optional)
	return task
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *StartWorkflowTask) Input(key string, value interface{}) *StartWorkflowTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *StartWorkflowTask) InputMap(inputMap map[string]interface{}) *StartWorkflowTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
