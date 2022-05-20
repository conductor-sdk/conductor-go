package workflow_task_e2e

import (
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/tests/e2e/workflow_e2e/workflow_e2e_properties"
)

var (
	SIMPLE_TASK_WORKFLOW_NAME = "GO_WORKFLOW_WITH_SIMPLE_TASK"
	SIMPLE_TASK_NAME          = "GO_TASK_OF_SIMPLE_TYPE"

	SIMPLE_TASK = workflow.NewSimpleTask(
		SIMPLE_TASK_NAME,
		SIMPLE_TASK_NAME,
	)

	SIMPLE_WORKFLOW = workflow.NewConductorWorkflow(workflow_e2e_properties.WorkflowExecutor).
			Name(SIMPLE_TASK_WORKFLOW_NAME).
			Version(1).
			Add(SIMPLE_TASK)
)
