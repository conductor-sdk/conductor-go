package examples

import (
	"fmt"

	"github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

func NewKitchenSinkWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	task := workflow.NewSimpleTask("simple_task_0")
	simpleWorkflow := workflow.NewConductorWorkflow(executor).
		Name("inline_sub").
		Add(
			workflow.NewSimpleTask("simple_task_0"),
		)
	subWorkflowInline := workflow.NewSubWorkflowInlineTask(
		"sub_flow_inline",
		simpleWorkflow,
	)
	decide := workflow.NewSwitchTask("fact_length", "$.number < 15 ? 'SHORT':'LONG'").
		Description("Fail if the fact is too short").
		Input("number", "${get_data.output.number}").
		UseJavascript(true).
		SwitchCase(
			"LONG",
			workflow.NewSimpleTask("simple_task_1"),
			workflow.NewSimpleTask("simple_task_1"),
		).
		SwitchCase(
			"SHORT",
			workflow.NewTerminateTask(
				"too_short",
				workflow_status.FAILED,
				"value too short",
			),
		)
	doWhile := workflow.NewLoopTask("loop_until_success", 2, decide).
		Optional(true)
	fork := workflow.NewForkTask(
		"fork",
		[]workflow.TaskInterface{
			doWhile,
			subWorkflowInline,
		},
		[]workflow.TaskInterface{
			workflow.NewSimpleTask("simple_task_5"),
		},
	)
	dynamicFork := workflow.NewDynamicForkTask(
		"dynamic_fork",
		workflow.NewSimpleTask("dynamic_fork_prep"),
	)
	setVariable := workflow.NewSetVariableTask("set_state").
		Input("call_made", true).
		Input("number", task.OutputRef("number"))

	subWorkflow := workflow.NewSubWorkflowTask("sub_flow", "PopulationMinMax", nil)

	fmt.Println(subWorkflow)
	fmt.Println(setVariable)

	workflow := workflow.NewConductorWorkflow(executor).
		Name("sdk_kitchen_sink2").
		Version(1).
		OwnerEmail("viren@orkes.io").
		Add(task).
		Add(setVariable).
		Add(subWorkflow).
		Add(dynamicFork).
		Add(fork)

	return workflow
}
