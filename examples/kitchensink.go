package examples

import (
	"fmt"
	workflow_status "github.com/conductor-sdk/conductor-go/pkg/model/enum"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

func NewKitchenSinkWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	simpleWorkflow := workflow.NewConductorWorkflow(executor).Name("inline_sub")
	simpleWorkflow.Add(workflow.Simple("simple_task_2", "simple_task"))
	subWorkflowInline := workflow.SubWorkflowInline("sub_flow_inline", simpleWorkflow)

	task1 := workflow.Simple("simple_task_2", "get_data")

	decide := workflow.Switch("fact_length", "$.number < 15 ? 'SHORT':'LONG'").
		Description("Fail if the fact is too short")

	decide.
		Input("number", "${get_data.output.number}").
		UseJavascript(true).
		SwitchCase("LONG", workflow.Simple("simple_task_4", "simple_task_4"), workflow.Simple("simple_task_4", "simple_task_4")).
		SwitchCase("SHORT", workflow.Terminate("too_short", workflow_status.FAILED, "value too short"))

	doWhile := workflow.Loop("loop_until_success", 2, decide).Optional(true)
	fork := workflow.Fork("fork",
		[]workflow.TaskInterface{doWhile, subWorkflowInline},
		[]workflow.TaskInterface{workflow.Simple("simple_task_5", "simple_task_5")},
	)
	dynamicFork := workflow.DynamicFork("dynamic_fork", workflow.Simple("dynamic_fork_prep", "dynamic_fork_prep"))

	setVariable := workflow.SetVariable("set_state").
		Input("call_made", true).
		Input("number", task1.OutputRef("number"))

	subWorkflow := workflow.SubWorkflow("sub_flow", "PopulationMinMax", nil)

	fmt.Println(subWorkflow)
	fmt.Println(setVariable)

	workflow := workflow.NewConductorWorkflow(executor).
		Name("sdk_kitchen_sink2").
		Version(1).
		Add(task1).
		Add(task1).
		Add(setVariable).
		Add(subWorkflow).
		Add(dynamicFork).
		Add(fork)

	return workflow
}
