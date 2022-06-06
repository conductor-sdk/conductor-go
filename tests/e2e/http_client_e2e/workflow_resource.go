package http_client_e2e

import (
	"context"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"net/http"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
)

var workflowClient = conductor_http_client.WorkflowResourceApiService{
	APIClient: e2e_properties.API_CLIENT,
}

func StartWorkflows(workflowQty int, workflowName string) ([]string, error) {
	workflowIdList := make([]string, workflowQty)
	for i := 0; i < workflowQty; i += 1 {
		workflowId, _, err := StartWorkflow(workflowName)
		if err != nil {
			return nil, err
		}
		workflowIdList[i] = workflowId
	}
	return workflowIdList, nil
}

func StartWorkflow(workflowName string) (string, *http.Response, error) {
	return workflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		workflowName,
		nil,
	)
}

func GetWorkflowExecutionStatus(workflowId string) (model.Workflow, *http.Response, error) {
	return workflowClient.GetExecutionStatus(
		context.Background(),
		workflowId,
		nil,
	)
}
