//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package metrics

import (
	"net/http"
	"strconv"

	"github.com/conductor-sdk/conductor-go/sdk/settings"
	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var collectionEnabled bool = false

// ProvideMetrics start collecting metrics for the workers
// We use prometheus to collect metrics from the workers.  When called this function starts the metrics server and publishes the worker metrics
func ProvideMetrics(metricsSettings *settings.MetricsSettings) {
	defer handlePanicError("provide_metrics")
	if metricsSettings == nil {
		metricsSettings = settings.NewDefaultMetricsSettings()
	}

	for metricName, metricDetails := range counterTemplates {
		counterByName[metricName] = newCounter(metricDetails)
		prometheus.MustRegister(counterByName[metricName])
	}

	for metricName, metricDetails := range gaugeTemplates {
		gaugeByName[metricName] = newGauge(metricDetails)
		prometheus.MustRegister(gaugeByName[metricName])
	}
	collectionEnabled = true
	
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

func handlePanicError(message string) {
	err := recover()
	if err == nil {
		return
	}
	IncrementUncaughtException(message)
	log.Warning(
		"Uncaught panic",
		", message: ", message,
		", error: ", err,
	)
}
