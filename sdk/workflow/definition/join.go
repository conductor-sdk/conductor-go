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

type JoinTask struct {
	Task
	joinOn []string
}

func NewJoinTask(taskRefName string, joinOn ...string) *JoinTask {
	return &JoinTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          JOIN,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		joinOn: joinOn,
	}
}

func (task *JoinTask) toWorkflowTask() []model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].JoinOn = task.joinOn
	return workflowTasks
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *JoinTask) Optional(optional bool) *JoinTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *JoinTask) Description(description string) *JoinTask {
	task.Task.Description(description)
	return task
}
