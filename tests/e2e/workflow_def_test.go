package e2e

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

func TestWorkflowDefWithSimpleTask(t *testing.T) {
	conductorWorkflow := workflow.NewConductorWorkflow(
		workflowExecutor,
	).Name(
		WORKFLOW_NAME,
	).Version(
		1,
	).Add(
		workflow.Simple(
			TASK_NAME,
			TASK_REFERENCE_NAME,
		),
	)
	conductorWorkflow.Register()
	workflowExecutionChannel, err := conductorWorkflow.Start(nil)
	if err != nil {
		t.Error(err)
	}
	waitForCompletionOfWorkflows(
		t,
		[]executor.WorkflowExecutionChannel{
			workflowExecutionChannel,
		},
	)
}
