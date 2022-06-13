//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package http_client_e2e

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
)

const (
	taskName     = "TEST_GO_TASK_SIMPLE"
	workflowName = "TEST_GO_WORKFLOW_SIMPLE"
)

func TestUpdateTaskRefByName(t *testing.T) {
	workflowId, response, err := e2e_properties.WorkflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		workflowName,
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
	returnValue, response, err := e2e_properties.TaskClient.UpdateTaskByRefName(
		context.Background(),
		outputData,
		workflowId,
		taskName,
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
	go e2e_properties.ValidateWorkflowDaemon(
		5*time.Second,
		errorChannel,
		workflowId,
		outputData,
	)
	err = <-errorChannel
	if err != nil {
		t.Fatal(
			"Failed to validate workflow. Reason: ", err.Error(),
			", workflowId: ", workflowId,
		)
	}
}
