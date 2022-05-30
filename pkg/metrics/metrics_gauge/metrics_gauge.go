package metrics_gauge

import (
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model/metric_documentation"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model/metric_label"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model/metric_name"
	"github.com/prometheus/client_golang/prometheus"
)

var gaugeByName = map[metric_name.MetricName]*prometheus.GaugeVec{}

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
	metric_name.TASK_UPDATE_TIME: metric_model.NewMetricDetails(
		metric_name.TASK_UPDATE_TIME,
		metric_documentation.TASK_UPDATE_TIME,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
		},
	),
}

func init() {
	for metricName, metricDetails := range gaugeTemplates {
		gaugeByName[metricName] = newGauge(metricDetails)
		prometheus.MustRegister(gaugeByName[metricName])
	}
}

func RecordWorkflowInputPayloadSize(workflowType string, version string, payloadSize float64) {
	setGauge(
		metric_name.WORKFLOW_INPUT_SIZE,
		[]string{
			workflowType,
			version,
		},
		payloadSize,
	)
}

func RecordTaskResultPayloadSize(taskType string, payloadSize float64) {
	setGauge(
		metric_name.TASK_RESULT_SIZE,
		[]string{
			taskType,
		},
		payloadSize,
	)
}

func RecordTaskPollTime(taskType string, timeSpent float64) {
	setGauge(
		metric_name.TASK_POLL_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
}

func RecordTaskUpdateTime(taskType string, timeSpent float64) {
	setGauge(
		metric_name.TASK_UPDATE_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
}

func RecordTaskExecuteTime(taskType string, timeSpent float64) {
	setGauge(
		metric_name.TASK_EXECUTE_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
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

func setGauge(metricName metric_name.MetricName, labelValues []string, value float64) {
	gauge := getGauge(metricName, labelValues)
	if gauge != nil {
		(*gauge).Set(value)
	}
}

func getGauge(metricName metric_name.MetricName, labelValues []string) *prometheus.Gauge {
	gaugeVec, ok := gaugeByName[metricName]
	if !ok {
		return nil
	}
	gauge, _ := gaugeVec.GetMetricWithLabelValues(
		labelValues...,
	)
	return &gauge
}
