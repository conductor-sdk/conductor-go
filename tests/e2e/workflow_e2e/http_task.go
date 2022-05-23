package workflow_e2e

import (
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
)

var (
	HTTP_TASK_WORKFLOW_NAME = "GO_WORKFLOW_WITH_HTTP_TASK"
	HTTP_TASK_NAME          = "GO_TASK_OF_HTTP_TYPE"

	HTTP_TASK = workflow.NewHttpTask(
		HTTP_TASK_NAME,
		&workflow.HttpInput{
			Uri: "https://catfact.ninja/fact",
		},
	)

	HTTP_WORKFLOW = workflow.NewConductorWorkflow(workflowExecutor).
			Name(HTTP_TASK_WORKFLOW_NAME).
			Version(1).
			Add(HTTP_TASK)
)
