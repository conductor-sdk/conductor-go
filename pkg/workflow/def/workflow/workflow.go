package workflow

import (
	"net/http"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

type ConductorWorkflow struct {
	executor *executor.WorkflowExecutor
	name     string
	version  int32

	tasks []TaskInterface
}

func NewConductorWorkflow(executor *executor.WorkflowExecutor) *ConductorWorkflow {
	return &ConductorWorkflow{
		executor: executor,
	}
}

func (workflow *ConductorWorkflow) Name(name string) *ConductorWorkflow {
	workflow.name = name
	return workflow
}

func (workflow *ConductorWorkflow) GetName() string {
	return workflow.name
}

func (workflow *ConductorWorkflow) GetVersion() int32 {
	return workflow.version
}

func (workflow *ConductorWorkflow) Version(version int32) *ConductorWorkflow {
	workflow.version = version
	return workflow
}

func (workflow *ConductorWorkflow) Add(task TaskInterface) *ConductorWorkflow {
	workflow.tasks = append(workflow.tasks, task)
	return workflow
}

func (workflow *ConductorWorkflow) Register() (*http.Response, error) {
	return workflow.executor.RegisterWorkflow(
		workflow.toWorkflowDef(),
	)
}

func (workflow *ConductorWorkflow) Start(input interface{}) (executor.WorkflowExecutionChannel, error) {
	return workflow.executor.ExecuteWorkflow(
		workflow.name,
		workflow.version,
		input,
	)
}

func (workflow *ConductorWorkflow) toWorkflowDef() *http_model.WorkflowDef {
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

func getWorkflowTasksFromConductorWorkflow(workflow *ConductorWorkflow) []http_model.WorkflowTask {
	workflowTasks := make([]http_model.WorkflowTask, 0)
	for _, task := range workflow.tasks {
		workflowTasks = append(
			workflowTasks,
			task.toWorkflowTask()...,
		)
	}
	return workflowTasks
}
