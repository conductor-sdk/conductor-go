package orkestrator

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/metrics"
	"github.com/conductor-sdk/conductor-go/pkg/orkestrator"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
)

func TestWorkerOrkestratorWithoutAuthenticationSettings(t *testing.T) {
	metricsCollector := metrics.NewMetricsCollector()
	apiClient := conductor_http_client.NewAPIClient(
		nil,
		settings.NewHttpDefaultSettings(),
		metricsCollector,
	)
	workerOrkestrator := orkestrator.NewWorkerOrkestrator(
		metricsCollector,
		apiClient,
	)
	if workerOrkestrator == nil {
		t.Fail()
	}
}

func TestWorkerOrkestratorWithAuthenticationSettings(t *testing.T) {
	metricsCollector := metrics.NewMetricsCollector()
	authenticationSettings := settings.NewAuthenticationSettings(
		"keyId",
		"keySecret",
	)
	apiClient := conductor_http_client.NewAPIClient(
		authenticationSettings,
		settings.NewHttpDefaultSettings(),
		metricsCollector,
	)
	workerOrkestrator := orkestrator.NewWorkerOrkestrator(
		metricsCollector,
		apiClient,
	)
	if workerOrkestrator == nil {
		t.Fail()
	}
}
