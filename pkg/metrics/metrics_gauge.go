package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var gaugeByName = map[MetricName]*prometheus.GaugeVec{}

var gaugeTemplates = map[MetricName]*MetricDetails{
	WORKFLOW_INPUT_SIZE: NewMetricDetails(
		WORKFLOW_INPUT_SIZE,
		WORKFLOW_INPUT_SIZE,
		[]MetricLabel{
			WORKFLOW_TYPE,
			WORKFLOW_VERSION,
		},
	),
	TASK_RESULT_SIZE: NewMetricDetails(
		TASK_RESULT_SIZE,
		TASK_RESULT_SIZE,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_POLL_TIME: NewMetricDetails(
		TASK_POLL_TIME,
		TASK_POLL_TIME,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_EXECUTE_TIME: NewMetricDetails(
		TASK_EXECUTE_TIME,
		TASK_EXECUTE_TIME,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_UPDATE_TIME: NewMetricDetails(
		TASK_UPDATE_TIME,
		TASK_UPDATE_TIME,
		[]MetricLabel{
			TASK_TYPE,
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
		WORKFLOW_INPUT_SIZE,
		[]string{
			workflowType,
			version,
		},
		payloadSize,
	)
}

func RecordTaskResultPayloadSize(taskType string, payloadSize float64) {
	setGauge(
		TASK_RESULT_SIZE,
		[]string{
			taskType,
		},
		payloadSize,
	)
}

func RecordTaskPollTime(taskType string, timeSpent float64) {
	setGauge(
		TASK_POLL_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
}

func RecordTaskUpdateTime(taskType string, timeSpent float64) {
	setGauge(
		TASK_UPDATE_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
}

func RecordTaskExecuteTime(taskType string, timeSpent float64) {
	setGauge(
		TASK_EXECUTE_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
}

func newGauge(metricDetails *MetricDetails) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricDetails.Name,
			Help: metricDetails.Description,
		},
		metricDetails.Labels,
	)
}

func setGauge(metricName MetricName, labelValues []string, value float64) {
	gauge := getGauge(metricName, labelValues)
	if gauge != nil {
		(*gauge).Set(value)
	}
}

func getGauge(metricName MetricName, labelValues []string) *prometheus.Gauge {
	gaugeVec, ok := gaugeByName[metricName]
	if !ok {
		return nil
	}
	gauge, _ := gaugeVec.GetMetricWithLabelValues(
		labelValues...,
	)
	return &gauge
}
