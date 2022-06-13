//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow

type SetVariableTask struct {
	Task
}

func NewSetVariableTask(taskRefName string) *SetVariableTask {
	return &SetVariableTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SET_VARIABLE,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SetVariableTask) Input(key string, value interface{}) *SetVariableTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SetVariableTask) InputMap(inputMap map[string]interface{}) *SetVariableTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *SetVariableTask) Optional(optional bool) *SetVariableTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *SetVariableTask) Description(description string) *SetVariableTask {
	task.Task.Description(description)
	return task
}
