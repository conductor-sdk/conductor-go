//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package examples

import (
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/definition"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

func NewKitchenSinkWorkflow(executor *executor.WorkflowExecutor) *definition.ConductorWorkflow {
	task := definition.NewSimpleTask("simple_task_0", "simple_task_0")
	simpleWorkflow := definition.NewConductorWorkflow(executor).
		Name("inline_sub").
		Add(
			definition.NewSimpleTask("simple_task_0", "simple_task_0"),
		)
	subWorkflowInline := definition.NewSubWorkflowInlineTask(
		"sub_flow_inline",
		simpleWorkflow,
	)
	decide := definition.NewSwitchTask("fact_length", "$.number < 15 ? 'LONG':'LONG'").
		Description("Fail if the fact is too short").
		Input("number", "${get_data.output.number}").
		UseJavascript(true).
		SwitchCase(
			"LONG",
			definition.NewSimpleTask("simple_task_1", "simple_task_1"),
			definition.NewSimpleTask("simple_task_1", "simple_task_1"),
		).
		SwitchCase(
			"SHORT",
			definition.NewTerminateTask(
				"too_short",
				model.FailedWorkflow,
				"value too short",
			),
		)
	doWhile := definition.NewLoopTask("loop_until_success", 2, decide).
		Optional(true)
	fork := definition.NewForkTask(
		"fork",
		[]definition.TaskInterface{
			doWhile,
			subWorkflowInline,
		},
		[]definition.TaskInterface{
			definition.NewSimpleTask("simple_task_5", "simple_task_5"),
		},
	)
	dynamicFork := definition.NewDynamicForkTask(
		"dynamic_fork",
		definition.NewSimpleTask("dynamic_fork_prep", "dynamic_fork_prep"),
	)
	setVariable := definition.NewSetVariableTask("set_state").
		Input("call_made", true).
		Input("number", task.OutputRef("number"))

	subWorkflow := definition.NewSubWorkflowTask("sub_flow", "PopulationMinMax", nil)

	fmt.Println(subWorkflow)
	fmt.Println(setVariable)

	jqTask := definition.NewJQTask("jq", "{ key3: (.key1.value1 + .key2.value2) }")
	jqTask.Input("key1", map[string]interface{}{
		"value1": []string{"a", "b"},
	})
	jqTask.InputMap(map[string]interface{}{
		"value2": []string{"d", "e"},
	})

	workflow := definition.NewConductorWorkflow(executor).
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
