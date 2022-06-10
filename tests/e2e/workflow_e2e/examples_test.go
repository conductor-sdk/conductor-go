package workflow_e2e

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/examples/shipment_example"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
)

func TestOrderWorkflow(t *testing.T) {
	err := e2e_properties.ValidateTaskRegistration(shipment_example.TaskCalculateTaxAndTotal)
	if err != nil {
		t.Fatal(err)
	}
	err = e2e_properties.ValidateTaskRegistration(shipment_example.TaskChargePayment)
	if err != nil {
		t.Fatal(err)
	}
	workflow := shipment_example.NewOrderWorkflow(e2e_properties.WorkflowExecutor)
	err = e2e_properties.ValidateWorkflowRegistration(workflow)
	if err != nil {
		t.Fatal(err)
	}
}
