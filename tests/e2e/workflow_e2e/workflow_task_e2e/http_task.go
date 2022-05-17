package workflow_task_e2e

import (
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/tests/e2e/workflow_e2e/workflow_e2e_properties"
)

var (
	HTTP_TASK_WORKFLOW_NAME = "GO_WORKFLOW_WITH_HTTP_TASK"
	HTTP_TASK_NAME          = "GO_TASK_OF_HTTP_TYPE"

	HTTP_TASK = workflow.Http(
		HTTP_TASK_NAME,
		&workflow.HttpInput{
			Uri: "https://catfact.ninja/fact",
		},
	)

	HTTP_WORKFLOW = workflow.NewConductorWorkflow(workflow_e2e_properties.WorkflowExecutor).
			Name(HTTP_TASK_WORKFLOW_NAME).
			Version(1).
			Add(HTTP_TASK)
)
