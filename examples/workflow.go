package examples

import (
	model2 "github.com/conductor-sdk/conductor-go/model"
	def2 "github.com/conductor-sdk/conductor-go/workflow/def"
	"github.com/conductor-sdk/conductor-go/workflow/executor"
)

func IsWorkflowCompleted(workflow *model2.Workflow) bool {
	return workflow.Status == model2.COMPLETED
}

func NewHttpTaskConductorWorkflow(workflowExecutor *executor.WorkflowExecutor) *def2.ConductorWorkflow {
	return newConductorWorkflow(
		workflowExecutor,
		"go_workflow_with_http_task",
		HttpTask,
	)
}

func NewSimpleTaskConductorWorkflow(workflowExecutor *executor.WorkflowExecutor) *def2.ConductorWorkflow {
	return newConductorWorkflow(
		workflowExecutor,
		"go_workflow_with_simple_task",
		SimpleTask,
	)
}

func newConductorWorkflow(workflowExecutor *executor.WorkflowExecutor, workflowName string, task def2.TaskInterface) *def2.ConductorWorkflow {
	return def2.NewConductorWorkflow(workflowExecutor).
		Name(workflowName).
		Version(1).
		Add(task)
}
