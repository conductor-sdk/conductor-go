//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

const (
	EvaluatorTypeValueParam = "value-param"
	EvaluatorTypeJavaScript = "javascript"
)

type SwitchTask struct {
	Task
	DecisionCases map[string][]TaskInterface
	defaultCase   []TaskInterface
	expression    string
	useJavascript bool
	evaluatorType string
}

func NewSwitchTask(taskRefName string, caseExpression string) *SwitchTask {
	return &SwitchTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SWITCH,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		DecisionCases: make(map[string][]TaskInterface),
		defaultCase:   make([]TaskInterface, 0),
		expression:    caseExpression,
		useJavascript: false,
		evaluatorType: EvaluatorTypeValueParam,
	}
}

func (task *SwitchTask) SwitchCase(caseName string, tasks ...TaskInterface) *SwitchTask {
	task.DecisionCases[caseName] = tasks
	return task
}
func (task *SwitchTask) DefaultCase(tasks ...TaskInterface) *SwitchTask {
	task.defaultCase = tasks
	return task
}

func (task *SwitchTask) toWorkflowTask() []model.WorkflowTask {
	var DecisionCases = map[string][]model.WorkflowTask{}
	for caseValue, tasks := range task.DecisionCases {
		for _, task := range tasks {
			DecisionCases[caseValue] = append(DecisionCases[caseValue], task.toWorkflowTask()...)
		}
	}
	var defaultCase []model.WorkflowTask
	for _, task := range task.defaultCase {
		defaultCase = append(defaultCase, task.toWorkflowTask()...)
	}
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].DecisionCases = DecisionCases
	workflowTasks[0].DefaultCase = defaultCase
	workflowTasks[0].EvaluatorType = task.evaluatorType
	if task.useJavascript {
		workflowTasks[0].Expression = task.expression
	} else {
		workflowTasks[0].Expression = "switchCaseValue"
		workflowTasks[0].InputParameters["switchCaseValue"] = task.expression
	}

	return workflowTasks
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SwitchTask) Input(key string, value interface{}) *SwitchTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SwitchTask) InputMap(inputMap map[string]interface{}) *SwitchTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Description of the task
func (task *SwitchTask) Description(description string) *SwitchTask {
	task.Task.Description(description)
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *SwitchTask) Optional(optional bool) *SwitchTask {
	task.Task.Optional(optional)
	return task
}

// UseJavascript If set to to true, the caseExpression parameter is treated as a Javascript.
// If set to false, the caseExpression follows the regular task input mapping format as described in https://conductor.netflix.com/how-tos/Tasks/task-inputs.html
func (task *SwitchTask) UseJavascript(use bool) *SwitchTask {
	task.useJavascript = use
	task.evaluatorType = EvaluatorTypeJavaScript
	return task
}
