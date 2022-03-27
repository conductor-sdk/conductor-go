package metrics_counter

import (
	"github.com/netflix/conductor/client/go/metrics/metric_model"
	"github.com/netflix/conductor/client/go/metrics/metric_model/metric_documentation"
	"github.com/netflix/conductor/client/go/metrics/metric_model/metric_label"
	"github.com/netflix/conductor/client/go/metrics/metric_model/metric_name"
	"github.com/prometheus/client_golang/prometheus"
)

func NewCounterByName() map[metric_name.MetricName]*prometheus.CounterVec {
	counterByName := map[metric_name.MetricName]*prometheus.CounterVec{}
	for metricName, metricDetails := range counterTemplates {
		counter := newCounter(metricDetails)
		counterByName[metricName] = counter
		prometheus.MustRegister(counter)
	}
	return counterByName
}

var counterTemplates = map[metric_name.MetricName]*metric_model.MetricDetails{
	metric_name.TASK_POLL: metric_model.NewMetricDetails(
		metric_name.TASK_POLL,
		metric_documentation.TASK_POLL,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
		},
	),
}

func newCounter(metricDetails *metric_model.MetricDetails) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricDetails.Name,
			Help: metricDetails.Description,
		},
		metricDetails.Labels,
	)
}
