//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package integration_tests

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
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

const (
	taskName = "TEST_GO_TASK_SIMPLE"

	workflowName              = "TEST_GO_WORKFLOW_SIMPLE"
	workflowCompletionTimeout = 5 * time.Second
	workflowExecutionQty      = 15

	workerQty          = 7
	workerPollInterval = 250 * time.Millisecond
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
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
		workflowExecutionQty,
		workflowName,
	)
	if err != nil {
		return err
	}
	err = TaskRunner.StartWorker(
		taskName,
		worker,
		workerQty,
		workerPollInterval,
	)
	if err != nil {
		return err
	}
	runningWorkflows := make([]chan error, len(workflowIdList))
	for i, workflowId := range workflowIdList {
		runningWorkflows[i] = make(chan error)
		go ValidateWorkflowDaemon(
			workflowCompletionTimeout,
			runningWorkflows[i],
			workflowId,
			expectedOutput,
		)
	}
	for _, runningWorkflowChannel := range runningWorkflows {
		err := <-runningWorkflowChannel
		if err != nil {
			return err
		}
	}
	return TaskRunner.RemoveWorker(
		taskName,
		workerQty,
	)
}
