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
		string(model.COMPLETED),
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
