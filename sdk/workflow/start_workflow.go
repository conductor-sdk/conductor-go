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

type StartWorkflowTask struct {
	Task
}

func NewStartWorkflowTask(taskRefName string, workflowName string, version *int32, startWorkflowRequest *model.StartWorkflowRequest) *StartWorkflowTask {
	return &StartWorkflowTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          START_WORKFLOW,
			optional:          false,
			inputParameters: map[string]interface{}{
				"startWorkflow": map[string]interface{}{
					"name":          workflowName,
					"version":       version,
					"input":         startWorkflowRequest.Input,
					"correlationId": startWorkflowRequest.CorrelationId,
				},
			},
		},
	}
}

func (task *StartWorkflowTask) toWorkflowTask() []model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	return workflowTasks
}

// Description of the task
func (task *StartWorkflowTask) Description(description string) *StartWorkflowTask {
	task.Task.Description(description)
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *StartWorkflowTask) Optional(optional bool) *StartWorkflowTask {
	task.Task.Optional(optional)
	return task
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *StartWorkflowTask) Input(key string, value interface{}) *StartWorkflowTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *StartWorkflowTask) InputMap(inputMap map[string]interface{}) *StartWorkflowTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}
