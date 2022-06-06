package workflow

import (
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
)

type StartWorkflowTask struct {
	Task
}

func NewStartWorkflowTask(taskRefName string, workflowName string, version optional.Int32, startWorkflowRequest *http_model.StartWorkflowRequest) *StartWorkflowTask {
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

func (task *StartWorkflowTask) toWorkflowTask() []http_model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	return workflowTasks
}

func (task *StartWorkflowTask) Description(description string) *StartWorkflowTask {
	task.Task.Description(description)
	return task
}

func (task *StartWorkflowTask) Optional(optional bool) *StartWorkflowTask {
	task.Task.Optional(optional)
	return task
}
func (task *StartWorkflowTask) Input(key string, value interface{}) *StartWorkflowTask {
	task.Task.Input(key, value)
	return task
}
func (task *StartWorkflowTask) InputMap(inputMap map[string]interface{}) *StartWorkflowTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
