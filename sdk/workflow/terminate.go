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

func NewTerminateTask(taskRefName string, status model.WorkflowStatus, terminationReason string) *Task {
	return &Task{
		model.WorkflowTask{
			Name:              taskRefName,
			TaskReferenceName: taskRefName,
			Description:       "",
			WorkflowTaskType:  string(TERMINATE),
			Optional:          false,
			InputParameters: map[string]interface{}{
				"terminationStatus": status,
				"terminationReason": terminationReason,
			},
		},
	}
}
