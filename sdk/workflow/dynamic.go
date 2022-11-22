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

type DynamicTask struct {
	Task
}

const dynamicTaskNameParameter = "taskToExecute"

// NewDynamicTask
//   - taskRefName Reference name for the task.  MUST be unique within the workflow
//   - taskNameParameter Parameter that contains the expression for the dynamic task name.  e.g. ${workflow.input.dynamicTask}
func NewDynamicTask(taskRefName string, taskNameParameter string) *DynamicTask {
	return &DynamicTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          DYNAMIC,
			optional:          false,
			inputParameters: map[string]interface{}{
				dynamicTaskNameParameter: taskNameParameter,
			},
		},
	}
}

func (task *DynamicTask) toWorkflowTask() []model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].DynamicTaskNameParam = dynamicTaskNameParameter
	return workflowTasks
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *DynamicTask) Input(key string, value interface{}) *DynamicTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *DynamicTask) InputMap(inputMap map[string]interface{}) *DynamicTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *DynamicTask) Optional(optional bool) *DynamicTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *DynamicTask) Description(description string) *DynamicTask {
	task.Task.Description(description)
	return task
}
