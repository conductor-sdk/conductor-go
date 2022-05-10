package conductor_http_client

import (
	"context"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/tests/e2e/conductor_client"
)

func StartWorkflow(workflowName string) (string, error) {
	apiClient := conductor_http_client.NewAPIClient(
		conductor_client.GetAuthenticationSettings(),
		conductor_client.GetHttpSettingsWithAuth(),
	)
	workflowClient := *&conductor_http_client.WorkflowResourceApiService{
		APIClient: apiClient,
	}
	workflowId, _, err := workflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		workflowName,
		nil,
	)
	if err != nil {
		return "", err
	}
	return workflowId, nil
}
