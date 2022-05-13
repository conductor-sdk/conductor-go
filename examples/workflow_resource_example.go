package examples

import (
	"context"
	"net/http"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
)

func StartWorkflow(workflowName string, version optional.Int32, correlationId optional.String) (string, *http.Response, error) {
	workflowClient := conductor_http_client.WorkflowResourceApiService{
		APIClient: conductor_http_client.NewAPIClient(
			settings.NewAuthenticationSettings(
				"key",
				"id",
			),
			settings.NewHttpDefaultSettings(),
		),
	}
	return workflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		workflowName,
		&conductor_http_client.WorkflowResourceApiStartWorkflowOpts{
			Version:       version,
			CorrelationId: correlationId,
		},
	)
}
