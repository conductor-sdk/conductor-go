package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

func SubWorkflow(taskRefName string, workflowName string, version *int32) *subWorkflow {
	return &subWorkflow{
		task: Task{
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

func SubWorkflowInline(taskRefName string, workflow *ConductorWorkflow) *subWorkflow {
	return &subWorkflow{
		task: Task{
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

type subWorkflow struct {
	task            Task
	workflowName    string
	version         *int32
	taskToDomainMap map[string]string
	workflow        *ConductorWorkflow
}

func (task *subWorkflow) Description(description string) *subWorkflow {
	task.task.Description(description)
	return task
}

func (task *subWorkflow) Optional(optional bool) *subWorkflow {
	task.task.Optional(optional)
	return task
}
func (task *subWorkflow) Input(key string, value interface{}) *subWorkflow {
	task.task.Input(key, value)
	return task
}
func (task *subWorkflow) TaskToDomain(taskToDomainMap map[string]string) *subWorkflow {
	task.taskToDomainMap = taskToDomainMap
	return task
}
func (task *subWorkflow) toWorkflowTask() []http_model.WorkflowTask {
	workflowTasks := task.task.toWorkflowTask()
	if task.workflow != nil {
		workflowTasks[0].SubWorkflowParam = &http_model.SubWorkflowParams{
			Name:               task.workflow.name,
			TaskToDomain:       task.taskToDomainMap,
			WorkflowDefinition: task.workflow.toWorkflowDef(),
		}
	} else {
		workflowTasks[0].SubWorkflowParam = &http_model.SubWorkflowParams{
			Name:               task.workflowName,
			Version:            task.version,
			TaskToDomain:       task.taskToDomainMap,
			WorkflowDefinition: nil,
		}
	}
	return workflowTasks
}
