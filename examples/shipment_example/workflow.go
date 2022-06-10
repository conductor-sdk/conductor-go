package shipment_example

import (
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

func NewShipmentWorkflow(workflowExecutor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	return workflow.NewConductorWorkflow(workflowExecutor).
		Name("EXAMPLE_GO_SHIPMENT_WORKFLOW").
		Version(1).
		OwnerEmail("developers@orkes.io").
		TimeoutPolicy(workflow.TimeOutWorkflow, 60)
}
