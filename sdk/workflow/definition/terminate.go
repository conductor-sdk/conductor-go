//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package definition

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

type TerminateTask struct {
	Task
}

func NewTerminateTask(taskRefName string, status model.WorkflowStatus, terminationReason string) *TerminateTask {
	return &TerminateTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          TERMINATE,
			optional:          false,
			inputParameters: map[string]interface{}{
				"terminationStatus": status,
				"terminationReason": terminationReason,
			},
		},
	}
}

// Description of the task
func (task *TerminateTask) Description(description string) *TerminateTask {
	task.Task.Description(description)
	return task
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *TerminateTask) Input(key string, value interface{}) *TerminateTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *TerminateTask) InputMap(inputMap map[string]interface{}) *TerminateTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
