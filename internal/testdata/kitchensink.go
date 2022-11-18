//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package testdata

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/status"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

func NewKitchenSinkWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	task := workflow.NewSimpleTask("simple_task", "simple_task_0")
	simpleWorkflow := workflow.NewConductorWorkflow(executor).
		Name("inline_sub").
		Add(
			workflow.NewSimpleTask("simple_task", "simple_task_0"),
		)
	subWorkflowInline := workflow.NewSubWorkflowInlineTask(
		"sub_flow_inline",
		simpleWorkflow,
	)
	decide := workflow.NewSwitchTask("fact_length", "$.number < 15 ? 'LONG':'LONG'").
		Description("Fail if the fact is too short").
		Input("number", "${get_data.output.number}").
		UseJavascript(true).
		SwitchCase(
			"LONG",
			workflow.NewSimpleTask("simple_task", "simple_task_1"),
			workflow.NewSimpleTask("simple_task", "simple_task_1"),
		).
		SwitchCase(
			"SHORT",
			workflow.NewTerminateTask(
				"too_short",
				status.FailedWorkflow,
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
			workflow.NewSimpleTask("simple_task", "simple_task_5"),
		},
	)
	dynamicFork := workflow.NewDynamicForkTask(
		"dynamic_fork",
		workflow.NewSimpleTask("dynamic_fork_prep", "dynamic_fork_prep"),
	)
	setVariable := workflow.NewSetVariableTask("set_state").
		Input("call_made", true).
		Input("number", task.OutputRef("number"))

	subWorkflow := workflow.NewSubWorkflowTask("sub_flow", "PopulationMinMax", nil)

	jqTask := workflow.NewJQTask("jq", "{ key3: (.key1.value1 + .key2.value2) }")
	jqTask.Input("key1", map[string]interface{}{
		"value1": []string{"a", "b"},
	})
	jqTask.InputMap(map[string]interface{}{
		"value2": []string{"d", "e"},
	})

	workflow := workflow.NewConductorWorkflow(executor).
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

type WorkflowTask struct {
	Name              string `json:"name"`
	TaskReferenceName string `json:"taskReferenceName"`
	Type              string `json:"type,omitempty"`
}

func DynamicForkWorker(t *model.Task) (output interface{}, err error) {
	taskResult := model.NewTaskResultFromTask(t)
	tasks := []WorkflowTask{
		{
			Name:              "simple_task_1",
			TaskReferenceName: "simple_task_11",
			Type:              "SIMPLE",
		},
		{
			Name:              "simple_task_3",
			TaskReferenceName: "simple_task_12",
			Type:              "SIMPLE",
		},
		{
			Name:              "simple_task_5",
			TaskReferenceName: "simple_task_13",
			Type:              "SIMPLE",
		},
	}
	inputs := map[string]interface{}{
		"simple_task_11": map[string]interface{}{
			"key1": "value1",
			"key2": 121,
		},
		"simple_task_12": map[string]interface{}{
			"key1": "value2",
			"key2": 122,
		},
		"simple_task_13": map[string]interface{}{
			"key1": "value3",
			"key2": 123,
		},
	}

	taskResult.OutputData = map[string]interface{}{
		"forkedTasks":       tasks,
		"forkedTasksInputs": inputs,
	}
	taskResult.Status = string(status.CompletedTask)
	err = nil
	return taskResult, err
}
