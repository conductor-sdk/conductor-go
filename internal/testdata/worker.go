//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package testdata

import (
	"fmt"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

const (
	TaskName = "TEST_GO_TASK_SIMPLE"

	WorkflowName              = "TEST_GO_WORKFLOW_SIMPLE"
	WorkflowCompletionTimeout = 5 * time.Second
	WorkflowExecutionQty      = 15

	WorkerQty          = 7
	WorkerPollInterval = 250 * time.Millisecond
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
		"key":  "value",
		"key1": "value1",
		"key2": 42,
	}
	taskResult.Status = model.CompletedTask
	return taskResult, nil
}

func TestWorkers(t *testing.T) {
	outputData := map[string]interface{}{
		"key": "value",
	}
	workerWithTaskResultOutput := func(t *model.Task) (interface{}, error) {
		taskResult := model.NewTaskResultFromTask(t)
		taskResult.OutputData = outputData
		taskResult.Status = model.CompletedTask
		return taskResult, nil
	}
	workerWithGenericOutput := func(t *model.Task) (interface{}, error) {
		return outputData, nil
	}
	workers := []model.ExecuteTaskFunction{
		workerWithTaskResultOutput,
		workerWithGenericOutput,
	}
	for _, worker := range workers {
		err := validateWorker(worker, outputData)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func validateWorker(worker model.ExecuteTaskFunction, expectedOutput map[string]interface{}) error {
	workflowIdList, err := StartWorkflows(
		WorkflowExecutionQty,
		WorkflowName,
	)
	if err != nil {
		return err
	}
	err = TaskRunner.StartWorker(
		TaskName,
		worker,
		WorkerQty,
		WorkerPollInterval,
	)
	if err != nil {
		return err
	}
	runningWorkflows := make([]chan error, len(workflowIdList))
	for i, workflowId := range workflowIdList {
		runningWorkflows[i] = make(chan error)
		go ValidateWorkflowDaemon(
			WorkflowCompletionTimeout,
			runningWorkflows[i],
			workflowId,
			expectedOutput,
			model.CompletedWorkflow,
		)
	}
	for _, runningWorkflowChannel := range runningWorkflows {
		err := <-runningWorkflowChannel
		if err != nil {
			return err
		}
	}
	return TaskRunner.DecreaseBatchSize(
		TaskName,
		WorkerQty,
	)
}

func FaultyWorker(task *model.Task) (interface{}, error) {
	taskResult := model.NewTaskResultFromTask(task)
	taskResult.Status = model.FailedWithTerminalErrorTask
	taskResult.OutputData = map[string]interface{}{
		"some_relevant_key": "relevant value",
	}
	return taskResult, fmt.Errorf("random error")
}

func WorkerWithNonRetryableError(task *model.Task) (interface{}, error) {
	return nil, model.NewNonRetryableError(fmt.Errorf("testing out some error stuff"))
}
