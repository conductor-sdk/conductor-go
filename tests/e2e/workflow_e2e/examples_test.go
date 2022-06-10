package workflow_e2e

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/examples/shipment_example"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
)

func TestOrderWorkflow(t *testing.T) {
	err := e2e_properties.ValidateTaskRegistration(
		*shipment_example.TaskCalculateTaxAndTotal.ToTaskDef(),
		*shipment_example.TaskChargePayment.ToTaskDef(),
		*shipment_example.TaskGroundShippingLabel.ToTaskDef(),
		*shipment_example.SameDayShippingLabel.ToTaskDef(),
		*shipment_example.AirShippingLabel.ToTaskDef(),
		*shipment_example.UnsupportedShippingLabel.ToTaskDef(),
		*shipment_example.TaskShippingLabel.ToTaskDef(),
		*shipment_example.TaskSendEmail.ToTaskDef(),
	)
	if err != nil {
		t.Fatal(err)
	}
	workflow := shipment_example.NewOrderWorkflow(e2e_properties.WorkflowExecutor)
	err = e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}
