package examples

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/definition"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

func IsWorkflowCompleted(workflow *model.Workflow) bool {
	return workflow.Status == model.CompletedWorkflow
}

func NewHttpTaskConductorWorkflow(workflowExecutor *executor.WorkflowExecutor) *definition.ConductorWorkflow {
	return newConductorWorkflow(
		workflowExecutor,
		"go_workflow_with_http_task",
		HttpTask,
	)
}

func NewSimpleTaskConductorWorkflow(workflowExecutor *executor.WorkflowExecutor) *definition.ConductorWorkflow {
	return newConductorWorkflow(
		workflowExecutor,
		"go_workflow_with_simple_task",
		SimpleTask,
	)
}

func newConductorWorkflow(workflowExecutor *executor.WorkflowExecutor, workflowName string, task definition.TaskInterface) *definition.ConductorWorkflow {
	return definition.NewConductorWorkflow(workflowExecutor).
		Name(workflowName).
		Version(1).
		Add(task)
}
