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
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// NewWaitTask creates WAIT task used to wait until an external event or a timeout occurs
func NewWaitTask(taskRefName string) *Task {
	return &Task{
		model.WorkflowTask{
			Name:              taskRefName,
			TaskReferenceName: taskRefName,
			Description:       "",
			WorkflowTaskType:  string(WAIT),
			Optional:          false,
			InputParameters:   map[string]interface{}{},
		},
	}
}

func NewWaitForDurationTask(taskRefName string, duration time.Duration) *Task {
	return &Task{
		model.WorkflowTask{
			Name:              taskRefName,
			TaskReferenceName: taskRefName,
			Description:       "",
			WorkflowTaskType:  string(WAIT),
			Optional:          false,
			InputParameters: map[string]interface{}{
				"duration": duration.String(),
			},
		},
	}
}

func NewWaitUntilTask(taskRefName string, dateTime string) *Task {
	return &Task{
		model.WorkflowTask{
			Name:              taskRefName,
			TaskReferenceName: taskRefName,
			Description:       "",
			WorkflowTaskType:  string(WAIT),
			Optional:          false,
			InputParameters: map[string]interface{}{
				"until": dateTime,
			},
		},
	}
}
