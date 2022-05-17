package e2e

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/tests"
)

var (
	// COMPLETED_TERMINATE_TASK = workflow.Terminate(
	// 	tests.TASK_REFERENCE_NAME,
	// 	workflow_status.COMPLETED,
	// 	"",
	// )

	HTTP_TASK = workflow.Http(
		"call_something",
		&workflow.HttpInput{Uri: "https://catfact.ninja/fact"},
	)

	SIMPLE_TASK = workflow.Simple(
		tests.TASK_NAME,
		tests.TASK_REFERENCE_NAME,
	)
)

func TestWorkflowDefWithSimpleTask(t *testing.T) {
	getConductorWorkflowWithSimpleTask(t)
}

func TestWorkflowDefWithHttpTask(t *testing.T) {
	getConductorWorkflowWithHttpTask(t)
}

func getConductorWorkflowWithSimpleTask(t *testing.T) *workflow.ConductorWorkflow {
	conductorWorkflow := workflow.NewConductorWorkflow(
		workflowExecutor,
	).Name(
		tests.WORKFLOW_NAME,
	).Version(
		1,
	).Add(
		SIMPLE_TASK,
	)
	validateWorkflowRegistration(t, conductorWorkflow)
	return conductorWorkflow
}

func getConductorWorkflowWithHttpTask(t *testing.T) *workflow.ConductorWorkflow {
	conductorWorkflow := workflow.NewConductorWorkflow(
		workflowExecutor,
	).Name(
		"WORKFLOW_WITH_HTTP_TASK",
	).Version(
		1,
	).Add(
		HTTP_TASK,
	)
	validateWorkflowRegistration(t, conductorWorkflow)
	return conductorWorkflow
}

func validateWorkflowRegistration(t *testing.T, conductorWorkflow *workflow.ConductorWorkflow) {
	_, err := conductorWorkflow.Register()
	if err != nil {
		t.Error(err)
	}
}
