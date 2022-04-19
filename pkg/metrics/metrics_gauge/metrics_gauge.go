package metrics_gauge

import (
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model/metric_documentation"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model/metric_label"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model/metric_name"
	"github.com/prometheus/client_golang/prometheus"
)

func NewGaugeByName() map[metric_name.MetricName]*prometheus.GaugeVec {
	gaugeByName := map[metric_name.MetricName]*prometheus.GaugeVec{}
	for metricName, metricDetails := range gaugeTemplates {
		gaugeByName[metricName] = newGauge(metricDetails)
		prometheus.MustRegister(gaugeByName[metricName])
	}
	return gaugeByName
}

var gaugeTemplates = map[metric_name.MetricName]*metric_model.MetricDetails{
	metric_name.WORKFLOW_INPUT_SIZE: metric_model.NewMetricDetails(
		metric_name.WORKFLOW_INPUT_SIZE,
		metric_documentation.WORKFLOW_INPUT_SIZE,
		[]metric_label.MetricLabel{
			metric_label.WORKFLOW_TYPE,
			metric_label.WORKFLOW_VERSION,
		},
	),
	metric_name.TASK_RESULT_SIZE: metric_model.NewMetricDetails(
		metric_name.TASK_RESULT_SIZE,
		metric_documentation.TASK_RESULT_SIZE,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
		},
	),
	metric_name.TASK_POLL_TIME: metric_model.NewMetricDetails(
		metric_name.TASK_POLL_TIME,
		metric_documentation.TASK_POLL_TIME,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
		},
	),
	metric_name.TASK_EXECUTE_TIME: metric_model.NewMetricDetails(
		metric_name.TASK_EXECUTE_TIME,
		metric_documentation.TASK_EXECUTE_TIME,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
		},
	),
}

func newGauge(metricDetails *metric_model.MetricDetails) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricDetails.Name,
			Help: metricDetails.Description,
		},
		metricDetails.Labels,
	)
}
