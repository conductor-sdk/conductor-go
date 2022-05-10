package worker

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
)

func TestSimpleWorkerOrkerstrator(t *testing.T) {
	workerOrkestrator := worker.NewTaskRunner(nil, nil)
	if workerOrkestrator == nil {
		t.Fail()
	}
}

func TestWorkerOrkestratorWithoutAuthenticationSettings(t *testing.T) {
	apiClient := conductor_http_client.NewAPIClient(
		nil,
		settings.NewHttpDefaultSettings(),
	)
	workerOrkestrator := worker.NewWorkerOrkestratorWithApiClient(
		apiClient,
	)
	if workerOrkestrator == nil {
		t.Fail()
	}
}

func TestWorkerOrkestratorWithAuthenticationSettings(t *testing.T) {
	authenticationSettings := settings.NewAuthenticationSettings(
		"keyId",
		"keySecret",
	)
	apiClient := conductor_http_client.NewAPIClient(
		authenticationSettings,
		settings.NewHttpDefaultSettings(),
	)
	workerOrkestrator := worker.NewWorkerOrkestratorWithApiClient(
		apiClient,
	)
	if workerOrkestrator == nil {
		t.Fail()
	}
}
