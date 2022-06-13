package quickstart

import (
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/definition"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"os"
	"time"
)

type Address struct {
	Name    string
	Address []string
	Country string
}

type ShippingCost struct {
	Amount float32
}

func NewSimpleWorkflow() *definition.ConductorWorkflow {
	apiClient := client.NewAPIClient(
		settings.NewAuthenticationSettings(
			os.Getenv("KEY"),
			os.Getenv("SECRET"),
		),
		settings.NewHttpSettings(
			"https://play.orkes.io/api",
		))
	executor := executor.NewWorkflowExecutor(apiClient)

	workflow := definition.NewConductorWorkflow(executor).
		Name("my_first_workflow").
		Version(1).
		Description("My First Workflow").
		TimeoutPolicy(definition.TimeOutWorkflow, 60)

	//Create a task that calculates the shipping cost
	calculateShipmentCost := definition.NewSimpleTask("shipping_cost_cal", "shipping_cost_calc").
		Input("address", "${workflow.input.address}").
		Description("Calculates the cost of shipping based on the address")

	//Add two simple tasks
	workflow.
		Add(calculateShipmentCost).
		OutputParameters(calculateShipmentCost.OutputRef(""))

	return workflow
}

func CalculateShippingCost(task *model.Task) (interface{}, error) {
	return &ShippingCost{Amount: 100}, nil
}

func StartWorkers() {
	apiClient := client.NewAPIClient(
		settings.NewAuthenticationSettings(
			os.Getenv("KEY"),
			os.Getenv("SECRET"),
		),
		settings.NewHttpSettings(
			"https://play.orkes.io/api",
		))
	taskRunner := worker.NewTaskRunnerWithApiClient(apiClient)
	taskRunner.StartWorker("shipping_cost_cal", CalculateShippingCost, 1, time.Second*1)

	//Block
	taskRunner.WaitWorkers()
}
