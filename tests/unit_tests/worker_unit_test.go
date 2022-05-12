package unit_tests

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
)

func TestSimpleTaskRunner(t *testing.T) {
	taskRunner := worker.NewTaskRunner(nil, nil)
	if taskRunner == nil {
		t.Fail()
	}
}

func TestTaskRunnerWithoutAuthenticationSettings(t *testing.T) {
	apiClient := conductor_http_client.NewAPIClient(
		nil,
		settings.NewHttpDefaultSettings(),
	)
	taskRunner := worker.NewTaskRunnerWithApiClient(
		apiClient,
	)
	if taskRunner == nil {
		t.Fail()
	}
}

func TestTaskRunnerWithAuthenticationSettings(t *testing.T) {
	authenticationSettings := settings.NewAuthenticationSettings(
		"keyId",
		"keySecret",
	)
	apiClient := conductor_http_client.NewAPIClient(
		authenticationSettings,
		settings.NewHttpDefaultSettings(),
	)
	taskRunner := worker.NewTaskRunnerWithApiClient(
		apiClient,
	)
	if taskRunner == nil {
		t.Fail()
	}
}
