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
	"context"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/internal/testdata"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

func TestUpdateTaskRefByName(t *testing.T) {
	workflowId, response, err := testdata.WorkflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		testdata.WorkflowName,
		nil,
	)
	if err != nil {
		t.Fatal(
			"Failed to start workflow. Reason: ", err.Error(),
			", workflowId: ", workflowId,
			", response:, ", *response,
		)
	}
	outputData := map[string]interface{}{
		"key": "value",
	}
	returnValue, response, err := testdata.TaskClient.UpdateTaskByRefName(
		context.Background(),
		outputData,
		workflowId,
		testdata.TaskName,
		string(model.CompletedTask),
	)
	if err != nil {
		t.Fatal(
			"Failed to updated task by ref name. Reason: ", err.Error(),
			", workflowId: ", workflowId,
			", return_value: ", returnValue,
			", response:, ", *response,
		)
	}
	errorChannel := make(chan error)
	go testdata.ValidateWorkflowDaemon(
		5*time.Second,
		errorChannel,
		workflowId,
		outputData,
		model.CompletedWorkflow,
	)
	err = <-errorChannel
	if err != nil {
		t.Fatal(
			"Failed to validate workflow. Reason: ", err.Error(),
			", workflowId: ", workflowId,
		)
	}
}
