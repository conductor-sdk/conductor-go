package http_client_e2e

import (
	"context"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
)

func TestUpdateTaskRefByName(t *testing.T) {
	workflowId, response, err := e2e_properties.WorkflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		e2e_properties.WORKFLOW_NAME,
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
		e2e_properties.TASK_NAME,
		string(task_result_status.COMPLETED),
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
