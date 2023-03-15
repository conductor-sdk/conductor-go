//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

import (
	"encoding/json"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

var hostname string
var once sync.Once

type ExecuteTaskFunction func(t *Task) (interface{}, error)

type ValidateWorkflowFunction func(w *Workflow) (bool, error)

func NewTaskResultFromTask(task *Task) *TaskResult {
	return &TaskResult{
		TaskId:             task.TaskId,
		WorkflowInstanceId: task.WorkflowInstanceId,
		WorkerId:           getHostname(),
	}
}

func NewTaskResultFromTaskWithError(t *Task, err error) *TaskResult {
	taskResult := NewTaskResultFromTask(t)
	taskResult.ReasonForIncompletion = err.Error()
	switch err.(type) {
	case *NonRetryableError:
		taskResult.Status = FailedWithTerminalErrorTask
	default:
		taskResult.Status = FailedTask
	}
	return taskResult
}

func NewTaskResult(taskId string, workflowInstanceId string) *TaskResult {
	return &TaskResult{
		TaskId:             taskId,
		WorkflowInstanceId: workflowInstanceId,
		WorkerId:           getHostname(),
	}

}

func GetTaskResultFromTaskExecutionOutput(t *Task, taskExecutionOutput interface{}) (*TaskResult, error) {
	taskResult, ok := taskExecutionOutput.(*TaskResult)
	if !ok {
		taskResult = NewTaskResultFromTask(t)
		outputData, err := ConvertToMap(taskExecutionOutput)
		if err != nil {
			return nil, err
		}
		taskResult.OutputData = outputData
		taskResult.Status = CompletedTask
	}
	return taskResult, nil
}

func ConvertToMap(input interface{}) (map[string]interface{}, error) {
	if input == nil {
		return nil, nil
	}
	data, err := json.Marshal(input)
	if err != nil {
		log.Debug(
			"Failed to parse input",
			", reason: ", err.Error(),
		)
		return nil, err
	}
	var parsedInput map[string]interface{}
	json.Unmarshal(data, &parsedInput)
	return parsedInput, nil
}

func getHostname() string {
	once.Do(updateHostname)
	return hostname
}

func updateHostname() {
	hostname, _ = os.Hostname()
}
