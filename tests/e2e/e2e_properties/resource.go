package e2e_properties

import (
	"os"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
)

type WorkflowValidator func(*http_model.Workflow) bool

var (
	AUTHENTICATION_KEY_ID     = "KEY"
	AUTHENTICATION_KEY_SECRET = "SECRET"
	BASE_URL                  = "https://play.orkes.io"
)

var API_CLIENT = getApiClientWithAuthentication()

func getApiClientWithAuthentication() *conductor_http_client.APIClient {
	return conductor_http_client.NewAPIClient(
		getAuthenticationSettings(),
		getHttpSettingsWithAuth(),
	)
}

func getAuthenticationSettings() *settings.AuthenticationSettings {
	return settings.NewAuthenticationSettings(
		os.Getenv(AUTHENTICATION_KEY_ID),
		os.Getenv(AUTHENTICATION_KEY_SECRET),
	)
}

func getHttpSettingsWithAuth() *settings.HttpSettings {
	return settings.NewHttpSettings(
		BASE_URL,
	)
}
