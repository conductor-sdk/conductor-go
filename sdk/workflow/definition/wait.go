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
	"time"
)

type WaitTask struct {
	Task
}

//NewWaitTask creates WAIT task used to wait until an external event or a timeout occurs
func NewWaitTask(taskRefName string) *WaitTask {
	return &WaitTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          WAIT,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
	}
}
func NewWaitForDurationTask(taskRefName string, duration time.Duration) *WaitTask {
	return &WaitTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          WAIT,
			optional:          false,
			inputParameters: map[string]interface{}{
				"duration": duration.String(),
			},
		},
	}
}

func NewWaitUntilTask(taskRefName string, dateTime string) *WaitTask {
	return &WaitTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          WAIT,
			optional:          false,
			inputParameters: map[string]interface{}{
				"until": dateTime,
			},
		},
	}
}

// Description of the task
func (task *WaitTask) Description(description string) *WaitTask {
	task.Task.Description(description)
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *WaitTask) Optional(optional bool) *WaitTask {
	task.Task.Optional(optional)
	return task
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *WaitTask) Input(key string, value interface{}) *WaitTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *WaitTask) InputMap(inputMap map[string]interface{}) *WaitTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
