package definition

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

type SubWorkflowTask struct {
	Task
	workflowName    string
	version         *int32
	taskToDomainMap map[string]string
	workflow        *ConductorWorkflow
}

func NewSubWorkflowTask(taskRefName string, workflowName string, version *int32) *SubWorkflowTask {
	return &SubWorkflowTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SUB_WORKFLOW,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		workflowName: workflowName,
		version:      version,
	}
}

func NewSubWorkflowInlineTask(taskRefName string, workflow *ConductorWorkflow) *SubWorkflowTask {
	return &SubWorkflowTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SUB_WORKFLOW,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		workflow: workflow,
	}
}

func (task *SubWorkflowTask) TaskToDomain(taskToDomainMap map[string]string) *SubWorkflowTask {
	task.taskToDomainMap = taskToDomainMap
	return task
}
func (task *SubWorkflowTask) toWorkflowTask() []model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	if task.workflow != nil {
		workflowTasks[0].SubWorkflowParam = &model.SubWorkflowParams{
			Name:               task.workflow.name,
			TaskToDomain:       task.taskToDomainMap,
			WorkflowDefinition: task.workflow.ToWorkflowDef(),
		}
	} else {
		workflowTasks[0].SubWorkflowParam = &model.SubWorkflowParams{
			Name:               task.workflowName,
			Version:            task.version,
			TaskToDomain:       task.taskToDomainMap,
			WorkflowDefinition: nil,
		}
	}
	return workflowTasks
}

// Description of the task
func (task *SubWorkflowTask) Description(description string) *SubWorkflowTask {
	task.Task.Description(description)
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *SubWorkflowTask) Optional(optional bool) *SubWorkflowTask {
	task.Task.Optional(optional)
	return task
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SubWorkflowTask) Input(key string, value interface{}) *SubWorkflowTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SubWorkflowTask) InputMap(inputMap map[string]interface{}) *SubWorkflowTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
