package metrics

import (
	"net/http"
	"strconv"

	"github.com/netflix/conductor/client/go/settings"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ProvideDefaultMetrics() {
	metricsSettings := settings.NewDefaultMetricsSettings()
	ProvideMetrics(metricsSettings)
}

func ProvideMetrics(metricsSettings *settings.MetricsSettings) {
	http.Handle(
		metricsSettings.ApiEndpoint,
		promhttp.Handler(),
	)
	portString := strconv.Itoa(metricsSettings.Port)
	http.ListenAndServe(":"+portString, nil)
}
