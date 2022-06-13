//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package examples

import "github.com/conductor-sdk/conductor-go/sdk/model"

type WorkflowTask struct {
	Name              string `json:"name"`
	TaskReferenceName string `json:"taskReferenceName"`
	Type              string `json:"type,omitempty"`
}

//dynamic_fork_prep
func DynamicForkWorker(t *model.Task) (output interface{}, err error) {
	taskResult := model.NewTaskResultFromTask(t)
	tasks := []WorkflowTask{
		{
			Name:              "simple_task_1",
			TaskReferenceName: "simple_task_11",
			Type:              "SIMPLE",
		},
		{
			Name:              "simple_task_3",
			TaskReferenceName: "simple_task_12",
			Type:              "SIMPLE",
		},
		{
			Name:              "simple_task_5",
			TaskReferenceName: "simple_task_13",
			Type:              "SIMPLE",
		},
	}
	inputs := map[string]interface{}{
		"simple_task_11": map[string]interface{}{
			"key1": "value1",
			"key2": 121,
		},
		"simple_task_12": map[string]interface{}{
			"key1": "value2",
			"key2": 122,
		},
		"simple_task_13": map[string]interface{}{
			"key1": "value3",
			"key2": 123,
		},
	}

	taskResult.OutputData = map[string]interface{}{
		"forkedTasks":       tasks,
		"forkedTasksInputs": inputs,
	}
	taskResult.Status = model.CompletedTask
	err = nil
	return taskResult, err
}
