package shipment_example

import (
	"github.com/conductor-sdk/conductor-go/examples/shipment_example/shipment_method_example"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

var TaskCalculateTaxAndTotal = def.NewSimpleTask("calculate_tax_and_total", "calculate_tax_and_total").
	Input("orderDetail", "${workflow.input.orderDetail}")

var TaskChargePayment = def.NewSimpleTask("charge_payment", "charge_payment").
	InputMap(
		map[string]interface{}{
			"billingId":   "${workflow.input.userDetails.billingId}",
			"billingType": "${workflow.input.userDetails.billingType}",
			"amount":      "${calculate_tax_and_total.output.total_amount}",
		},
	)

var (
	shippingLabelInputMap = map[string]interface{}{
		"name":    "${workflow.input.userDetails.name}",
		"address": "${workflow.input.userDetails.addressLine}",
		"orderNo": "${workflow.input.orderDetails.orderNumber}",
	}

	TaskGroundShippingLabel  = def.NewSimpleTask("ground_shipping_label", "ground_shipping_label").InputMap(shippingLabelInputMap)
	SameDayShippingLabel     = def.NewSimpleTask("same_day_shipping_label", "same_day_shipping_label").InputMap(shippingLabelInputMap)
	AirShippingLabel         = def.NewSimpleTask("air_shipping_label", "air_shipping_label").InputMap(shippingLabelInputMap)
	UnsupportedShippingLabel = def.NewTerminateTask("unsupported_shipping_type", model.FAILED, "Unsupported Shipping Method")

	TaskShippingLabel = def.NewSwitchTask("shipping_label", "${workflow.input.orderDetail.shippingMethod}").
				SwitchCase(string(shipment_method_example.Ground), TaskGroundShippingLabel).
				SwitchCase(string(shipment_method_example.SameDay), SameDayShippingLabel).
				SwitchCase(string(shipment_method_example.NextDayAir), AirShippingLabel).
				DefaultCase(UnsupportedShippingLabel)
)

var TaskSendEmail = def.NewSimpleTask("send_email", "send_email").
	InputMap(
		map[string]interface{}{
			"name":    "${workflow.input.userDetails.name}",
			"email":   "${workflow.input.userDetails.email}",
			"orderNo": "${workflow.input.orderDetails.orderNumber}",
		},
	)

var TaskGetOrderDetails = def.NewSimpleTask("get_order_details", "get_order_details").
	Input("orderNo", "${workflow.input.orderNo}")

var TaskGetUserDetails = def.NewSimpleTask("get_user_details", "get_user_details").
	Input("userId", "${workflow.input.userId}")

var TaskGetInParallel = def.NewForkTask(
	"get_in_parallel",
	[]def.TaskInterface{
		TaskGetOrderDetails, TaskGetUserDetails,
	},
)

var (
	TaskGenerateDynamicFork = def.NewSimpleTask("generateDynamicFork", "generateDynamicFork").
				InputMap(
			map[string]interface{}{
				"orderDetails": TaskGetOrderDetails.OutputRef("result"),
				"userDetails":  TaskGetUserDetails.OutputRef(""),
			},
		)

	TaskProcessOrder = def.NewDynamicForkTask("process_order", TaskGenerateDynamicFork)
)

var TaskUpdateState = def.NewSetVariableTask("update_state").
	Input("shipped", true)

func NewOrderWorkflow(workflowExecutor *executor.WorkflowExecutor) *def.ConductorWorkflow {
	return def.NewConductorWorkflow(workflowExecutor).
		Name("example_go_order_workflow").
		Version(1).
		OwnerEmail("developers@orkes.io").
		TimeoutPolicy(def.TimeOutWorkflow, 60).
		Description("Workflow to track order").
		Add(TaskCalculateTaxAndTotal).
		Add(TaskChargePayment).
		Add(TaskShippingLabel).
		Add(TaskSendEmail)
}

func NewShipmentWorkflow(workflowExecutor *executor.WorkflowExecutor) *def.ConductorWorkflow {
	return def.NewConductorWorkflow(workflowExecutor).
		Name("example_go_shipment_workflow").
		Version(1).
		OwnerEmail("developers@orkes.io").
		Variables(NewShipmentState()).
		TimeoutPolicy(def.TimeOutWorkflow, 60).
		Description("Workflow to track shipment").
		Add(TaskGetInParallel).
		Add(TaskProcessOrder).
		Add(TaskUpdateState)
}
