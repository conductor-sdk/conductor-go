//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package definition

type JQTask struct {
	Task
}

func NewJQTask(name string, script string) *JQTask {
	return &JQTask{
		Task{
			name:              name,
			taskReferenceName: name,
			taskType:          JSON_JQ_TRANSFORM,
			inputParameters: map[string]interface{}{
				"queryExpression": script,
			},
		},
	}
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *JQTask) Input(key string, value interface{}) *JQTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *JQTask) InputMap(inputMap map[string]interface{}) *JQTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *JQTask) Optional(optional bool) *JQTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *JQTask) Description(description string) *JQTask {
	task.Task.Description(description)
	return task
}
