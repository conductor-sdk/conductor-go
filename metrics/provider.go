package metrics

import (
	"github.com/conductor-sdk/conductor-go/settings"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ProvideMetrics(metricsSettings *settings.MetricsSettings) {
	if metricsSettings == nil {
		metricsSettings = settings.NewDefaultMetricsSettings()
	}
	http.Handle(
		metricsSettings.ApiEndpoint,
		promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		),
	)
	portString := strconv.Itoa(metricsSettings.Port)
	http.ListenAndServe(":"+portString, nil)
}
