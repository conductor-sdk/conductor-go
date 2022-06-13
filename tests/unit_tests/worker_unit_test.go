package unit_tests

import (
	"github.com/conductor-sdk/conductor-go/client"
	settings2 "github.com/conductor-sdk/conductor-go/settings"
	"github.com/conductor-sdk/conductor-go/worker"
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
		settings2.NewHttpDefaultSettings(),
	)
	taskRunner := worker.NewTaskRunnerWithApiClient(
		apiClient,
	)
	if taskRunner == nil {
		t.Fail()
	}
}

func TestTaskRunnerWithAuthenticationSettings(t *testing.T) {
	authenticationSettings := settings2.NewAuthenticationSettings(
		"keyId",
		"keySecret",
	)
	apiClient := client.NewAPIClient(
		authenticationSettings,
		settings2.NewHttpDefaultSettings(),
	)
	taskRunner := worker.NewTaskRunnerWithApiClient(
		apiClient,
	)
	if taskRunner == nil {
		t.Fail()
	}
}
