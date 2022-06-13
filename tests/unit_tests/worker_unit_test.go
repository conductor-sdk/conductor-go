package unit_tests

import (
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"testing"
)

func TestSimpleTaskRunner(t *testing.T) {
	taskRunner := worker.NewTaskRunner(nil, nil)
	if taskRunner == nil {
		t.Fail()
	}
}

func TestTaskRunnerWithoutAuthenticationSettings(t *testing.T) {
	apiClient := client.NewAPIClient(
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
	apiClient := client.NewAPIClient(
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
