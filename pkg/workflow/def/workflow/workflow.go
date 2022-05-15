package workflow

import (
	"net/http"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

func NewConductorWorkflow(executor *executor.WorkflowExecutor) *conductorWorkflow {
	return &conductorWorkflow{
		executor: executor,
	}
}

type conductorWorkflow struct {
	executor *executor.WorkflowExecutor
	name     string
	version  int32
	tasks    []Task
}

func (workflow *conductorWorkflow) Name(name string) *conductorWorkflow {
	workflow.name = name
	return workflow
}
func (workflow *conductorWorkflow) Version(version int32) *conductorWorkflow {
	workflow.version = version
	return workflow
}
func (workflow *conductorWorkflow) Add(task Task) *conductorWorkflow {
	workflow.tasks = append(workflow.tasks, task)
	return workflow
}

func (workflow *conductorWorkflow) Register() (*http.Response, error) {
	return workflow.executor.RegisterWorkflow(
		workflow.toWorkflowDef(),
	)
}

func (workflow *conductorWorkflow) Start(input interface{}) (executor.WorkflowExecutionChannel, error) {
	return workflow.executor.ExecuteWorkflow(
		workflow.name,
		workflow.version,
		input,
	)
}

func (workflow *conductorWorkflow) toWorkflowDef() *http_model.WorkflowDef {
	return &http_model.WorkflowDef{
		Name:             workflow.name,
		Description:      "",
		Version:          workflow.version,
		Tasks:            getWorkflowTasksFromConductorWorkflow(workflow),
		InputParameters:  nil,
		OutputParameters: nil,
		FailureWorkflow:  "",
		SchemaVersion:    2,
		OwnerEmail:       "",
		TimeoutPolicy:    "",
		TimeoutSeconds:   0,
		Variables:        nil,
		InputTemplate:    nil,
	}
}

func getWorkflowTasksFromConductorWorkflow(workflow *conductorWorkflow) []http_model.WorkflowTask {
	workflowTasks := make([]http_model.WorkflowTask, 0)
	for _, task := range workflow.tasks {
		workflowTasks = append(
			workflowTasks,
			task.toWorkflowTask()...,
		)
	}
	return workflowTasks
}
