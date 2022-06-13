package examples

import (
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/def"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

func NewKitchenSinkWorkflow(executor *executor.WorkflowExecutor) *def.ConductorWorkflow {
	task := def.NewSimpleTask("simple_task_0", "simple_task_0")
	simpleWorkflow := def.NewConductorWorkflow(executor).
		Name("inline_sub").
		Add(
			def.NewSimpleTask("simple_task_0", "simple_task_0"),
		)
	subWorkflowInline := def.NewSubWorkflowInlineTask(
		"sub_flow_inline",
		simpleWorkflow,
	)
	decide := def.NewSwitchTask("fact_length", "$.number < 15 ? 'LONG':'LONG'").
		Description("Fail if the fact is too short").
		Input("number", "${get_data.output.number}").
		UseJavascript(true).
		SwitchCase(
			"LONG",
			def.NewSimpleTask("simple_task_1", "simple_task_1"),
			def.NewSimpleTask("simple_task_1", "simple_task_1"),
		).
		SwitchCase(
			"SHORT",
			def.NewTerminateTask(
				"too_short",
				model.FAILED,
				"value too short",
			),
		)
	doWhile := def.NewLoopTask("loop_until_success", 2, decide).
		Optional(true)
	fork := def.NewForkTask(
		"fork",
		[]def.TaskInterface{
			doWhile,
			subWorkflowInline,
		},
		[]def.TaskInterface{
			def.NewSimpleTask("simple_task_5", "simple_task_5"),
		},
	)
	dynamicFork := def.NewDynamicForkTask(
		"dynamic_fork",
		def.NewSimpleTask("dynamic_fork_prep", "dynamic_fork_prep"),
	)
	setVariable := def.NewSetVariableTask("set_state").
		Input("call_made", true).
		Input("number", task.OutputRef("number"))

	subWorkflow := def.NewSubWorkflowTask("sub_flow", "PopulationMinMax", nil)

	fmt.Println(subWorkflow)
	fmt.Println(setVariable)

	jqTask := def.NewJQTask("jq", "{ key3: (.key1.value1 + .key2.value2) }")
	jqTask.Input("key1", map[string]interface{}{
		"value1": []string{"a", "b"},
	})
	jqTask.InputMap(map[string]interface{}{
		"value2": []string{"d", "e"},
	})

	workflow := def.NewConductorWorkflow(executor).
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
