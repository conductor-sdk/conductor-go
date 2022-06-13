package examples

import (
	"fmt"
	"github.com/conductor-sdk/conductor-go/model"
	def2 "github.com/conductor-sdk/conductor-go/workflow/def"
	"github.com/conductor-sdk/conductor-go/workflow/executor"
)

func NewKitchenSinkWorkflow(executor *executor.WorkflowExecutor) *def2.ConductorWorkflow {
	task := def2.NewSimpleTask("simple_task_0", "simple_task_0")
	simpleWorkflow := def2.NewConductorWorkflow(executor).
		Name("inline_sub").
		Add(
			def2.NewSimpleTask("simple_task_0", "simple_task_0"),
		)
	subWorkflowInline := def2.NewSubWorkflowInlineTask(
		"sub_flow_inline",
		simpleWorkflow,
	)
	decide := def2.NewSwitchTask("fact_length", "$.number < 15 ? 'LONG':'LONG'").
		Description("Fail if the fact is too short").
		Input("number", "${get_data.output.number}").
		UseJavascript(true).
		SwitchCase(
			"LONG",
			def2.NewSimpleTask("simple_task_1", "simple_task_1"),
			def2.NewSimpleTask("simple_task_1", "simple_task_1"),
		).
		SwitchCase(
			"SHORT",
			def2.NewTerminateTask(
				"too_short",
				model.FAILED,
				"value too short",
			),
		)
	doWhile := def2.NewLoopTask("loop_until_success", 2, decide).
		Optional(true)
	fork := def2.NewForkTask(
		"fork",
		[]def2.TaskInterface{
			doWhile,
			subWorkflowInline,
		},
		[]def2.TaskInterface{
			def2.NewSimpleTask("simple_task_5", "simple_task_5"),
		},
	)
	dynamicFork := def2.NewDynamicForkTask(
		"dynamic_fork",
		def2.NewSimpleTask("dynamic_fork_prep", "dynamic_fork_prep"),
	)
	setVariable := def2.NewSetVariableTask("set_state").
		Input("call_made", true).
		Input("number", task.OutputRef("number"))

	subWorkflow := def2.NewSubWorkflowTask("sub_flow", "PopulationMinMax", nil)

	fmt.Println(subWorkflow)
	fmt.Println(setVariable)

	jqTask := def2.NewJQTask("jq", "{ key3: (.key1.value1 + .key2.value2) }")
	jqTask.Input("key1", map[string]interface{}{
		"value1": []string{"a", "b"},
	})
	jqTask.InputMap(map[string]interface{}{
		"value2": []string{"d", "e"},
	})

	workflow := def2.NewConductorWorkflow(executor).
		Name("sdk_kitchen_sink2").
		Version(1).
		OwnerEmail("orkes-workers@apps.orkes.io").
		Add(task).
		Add(jqTask).
		Add(setVariable).
		Add(subWorkflow).
		Add(dynamicFork).
		Add(fork)

	return workflow
}
