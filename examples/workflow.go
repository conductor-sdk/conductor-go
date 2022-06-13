package examples

import (
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

func IsWorkflowCompleted(workflow *model.Workflow) bool {
	return workflow.Status == model.COMPLETED
}

func NewHttpTaskConductorWorkflow(workflowExecutor *executor.WorkflowExecutor) *def.ConductorWorkflow {
	return newConductorWorkflow(
		workflowExecutor,
		"go_workflow_with_http_task",
		HttpTask,
	)
}

func NewSimpleTaskConductorWorkflow(workflowExecutor *executor.WorkflowExecutor) *def.ConductorWorkflow {
	return newConductorWorkflow(
		workflowExecutor,
		"go_workflow_with_simple_task",
		SimpleTask,
	)
}

func newConductorWorkflow(workflowExecutor *executor.WorkflowExecutor, workflowName string, task def.TaskInterface) *def.ConductorWorkflow {
	return def.NewConductorWorkflow(workflowExecutor).
		Name(workflowName).
		Version(1).
		Add(task)
}
