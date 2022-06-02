package examples

import (
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

func IsWorkflowCompleted(workflow *http_model.Workflow) bool {
	return workflow.Status == string(workflow_status.COMPLETED)
}

func NewHttpTaskConductorWorkflow(workflowName string, workflowExecutor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	return newConductorWorkflow(
		workflowExecutor,
		workflowName,
		HttpTask,
	)
}

func NewSimpleTaskConductorWorkflow(workflowName string, workflowExecutor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	return newConductorWorkflow(
		workflowExecutor,
		workflowName,
		SimpleTask,
	)
}

func newConductorWorkflow(workflowExecutor *executor.WorkflowExecutor, workflowName string, task workflow.TaskInterface) *workflow.ConductorWorkflow {
	return workflow.NewConductorWorkflow(workflowExecutor).
		Name(workflowName).
		Version(1).
		Add(task)
}