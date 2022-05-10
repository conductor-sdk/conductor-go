package conductor_http_client

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/tests/e2e/conductor_client"
	log "github.com/sirupsen/logrus"
)

func TestStartWorkflow(t *testing.T) {
	apiClient := conductor_http_client.NewAPIClient(
		conductor_client.GetAuthenticationSettings(),
		conductor_client.GetHttpSettingsWithAuth(),
	)
	workflowClient := *&conductor_http_client.WorkflowResourceApiService{
		APIClient: apiClient,
	}
	workflowId, response, err := workflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		"workflow_with_go_task_example_from_code",
		nil,
	)
	if err != nil {
		t.Error("response: ", response)
	}
	log.Warn("workflowId:", workflowId)
}
