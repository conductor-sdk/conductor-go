package conductor_client_tests

import (
	"context"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
)

func StartWorkflow(workflowName string) (string, error) {
	apiClient := GetApiClientWithAuthentication()
	workflowClient := conductor_http_client.WorkflowResourceApiService{
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

func GetApiClientWithAuthentication() *conductor_http_client.APIClient {
	return conductor_http_client.NewAPIClient(
		getAuthenticationSettings(),
		getHttpSettingsWithAuth(),
	)
}

func getAuthenticationSettings() *settings.AuthenticationSettings {
	return settings.NewAuthenticationSettings(
		"",
		"",
	)
}

func getHttpSettingsWithAuth() *settings.HttpSettings {
	return settings.NewHttpSettings(
		"https://play.orkes.io",
		nil,
	)
}
