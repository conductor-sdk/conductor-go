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

func NewSwitchTask(taskRefName string, caseExpression string) *Task {
	task := &Task{
		model.WorkflowTask{
			Name:              taskRefName,
			TaskReferenceName: taskRefName,
			Description:       "",
			WorkflowTaskType:  string(SWITCH),
			Optional:          false,
			InputParameters: map[string]interface{}{
				"switchCaseValue": caseExpression,
			},
			DecisionCases: make(map[string][]model.WorkflowTask),
			DefaultCase:   make([]model.WorkflowTask, 0),
			Expression:    caseExpression,
			EvaluatorType: "value-param",
		},
	}
	return task
}

func (t *Task) UseJavascript(flag bool) *Task {
	return t
}

func (t *Task) SwitchCase(caseName string, tasks ...*Task) *Task {
	workflowTasks := make([]model.WorkflowTask, len(tasks))
	for index, task := range tasks {
		workflowTasks[index] = task.WorkflowTask
	}
	t.WorkflowTask.DecisionCases[caseName] = workflowTasks
	return t
}

func (t *Task) DefaultCase(tasks ...*Task) *Task {
	workflowTasks := make([]model.WorkflowTask, len(tasks))
	for index, task := range tasks {
		workflowTasks[index] = task.WorkflowTask
	}
	t.WorkflowTask.DefaultCase = workflowTasks
	return t
}
