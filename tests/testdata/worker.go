//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package examples

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

func ExampleWorker(t *model.Task) (interface{}, error) {
	taskResult := model.NewTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"key0": nil,
		"key1": 3,
		"key2": false,
		"foo":  "bar",
	}
	taskResult.Logs = append(
		taskResult.Logs,
		model.TaskExecLog{
			Log: "log message",
		},
	)
	taskResult.Status = model.CompletedTask
	return taskResult, nil
}

func SimpleWorker(t *model.Task) (interface{}, error) {
	taskResult := model.NewTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"key": "value",
	}
	taskResult.Status = model.CompletedTask
	return taskResult, nil
}

func OpenTreasureChest(t *model.Task) (interface{}, error) {
	taskResult := model.NewTaskResultFromTask(t)
	taskResult.OutputData = map[string]interface{}{
		"treasure": t.InputData["importantValue"],
	}
	taskResult.Status = model.CompletedTask
	return taskResult, nil
}
