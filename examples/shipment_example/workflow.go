package shipment_example

import (
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

var (
	TaskCalculateTaxAndTotal = workflow.NewSimpleTask("calculate_tax_and_total").
					Input("orderDetail", "${workflow.input.orderDetail}")

	TaskChargePayment = workflow.NewSimpleTask("charge_payment").
				InputMap(
			map[string]interface{}{
				"billingId":   "${workflow.input.userDetails.billingId}",
				"billingType": "${workflow.input.userDetails.billingType}",
				"amount":      "${calculate_tax_and_total.output.total_amount}",
			},
		)
)

func NewOrderWorkflow(workflowExecutor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	return workflow.NewConductorWorkflow(workflowExecutor).
		Name("EXAMPLE_GO_ORDER_WORKFLOW").
		Version(1).
		OwnerEmail("developers@orkes.io").
		TimeoutPolicy(workflow.TimeOutWorkflow, 60).
		Description("Workflow to track shipment").
		Add(TaskCalculateTaxAndTotal).
		Add(TaskChargePayment)
}
