//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow

import "github.com/conductor-sdk/conductor-go/sdk/model"

type UpdateTask struct {
	Task
}

// Create a new Update with Task Id (targetTaskId)
func NewUpdateTaskWithTaskId(taskRefName string, status model.TaskResultStatus, targetTaskId string) *UpdateTask {

	update := &UpdateTask{
		Task: Task{
			name:              string(UPDATE),
			taskReferenceName: taskRefName,
			taskType:          UPDATE,
			inputParameters: map[string]interface{}{
				"taskId":     targetTaskId,
				"taskStatus": status,
			},
		},
	}

	return update
}

// Create a new Update Task with workflow Id (targetWorkflowId) and  task ref (targetTaskRefName)
func NewUpdateTask(taskRefName string, status model.TaskResultStatus, targetWorkflowId string, targetTaskRefName string) *UpdateTask {

	update := &UpdateTask{
		Task: Task{
			name:              string(UPDATE),
			taskReferenceName: taskRefName,
			taskType:          UPDATE,
			inputParameters: map[string]interface{}{
				"workflowId":  targetWorkflowId,
				"taskRefName": targetTaskRefName,
				"taskStatus":  status,
			},
		},
	}

	return update
}

func (task *UpdateTask) MergeOutput(value bool) *UpdateTask {
	task.Task.Input("mergeOutput", value)
	return task
}

func (task *UpdateTask) TaskOutput(inputMap map[string]interface{}) *UpdateTask {
	task.Task.Input("taskOutput", inputMap)
	return task
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *UpdateTask) Input(key string, value interface{}) *UpdateTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *UpdateTask) InputMap(inputMap map[string]interface{}) *UpdateTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *UpdateTask) Optional(optional bool) *UpdateTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *UpdateTask) Description(description string) *UpdateTask {
	task.Task.Description(description)
	return task
}
