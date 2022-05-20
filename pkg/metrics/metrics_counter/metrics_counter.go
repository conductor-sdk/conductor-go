package metrics_counter

import (
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model/metric_documentation"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model/metric_label"
	"github.com/conductor-sdk/conductor-go/pkg/metrics/metric_model/metric_name"
	"github.com/prometheus/client_golang/prometheus"
)

var counterByName = map[metric_name.MetricName]*prometheus.CounterVec{}

var counterTemplates = map[metric_name.MetricName]*metric_model.MetricDetails{
	metric_name.TASK_POLL: metric_model.NewMetricDetails(
		metric_name.TASK_POLL,
		metric_documentation.TASK_POLL,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
		},
	),
	metric_name.TASK_EXECUTION_QUEUE_FULL: metric_model.NewMetricDetails(
		metric_name.TASK_EXECUTION_QUEUE_FULL,
		metric_documentation.TASK_EXECUTION_QUEUE_FULL,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
		},
	),
	metric_name.THREAD_UNCAUGHT_EXCEPTION: metric_model.NewMetricDetails(
		metric_name.THREAD_UNCAUGHT_EXCEPTION,
		metric_documentation.THREAD_UNCAUGHT_EXCEPTION,
		[]metric_label.MetricLabel{
			metric_label.EXCEPTION,
		},
	),
	metric_name.TASK_POLL_ERROR: metric_model.NewMetricDetails(
		metric_name.TASK_POLL_ERROR,
		metric_documentation.TASK_POLL_ERROR,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
			metric_label.EXCEPTION,
		},
	),
	metric_name.TASK_PAUSED: metric_model.NewMetricDetails(
		metric_name.TASK_PAUSED,
		metric_documentation.TASK_PAUSED,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
		},
	),
	metric_name.TASK_EXECUTE_ERROR: metric_model.NewMetricDetails(
		metric_name.TASK_EXECUTE_ERROR,
		metric_documentation.TASK_EXECUTE_ERROR,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
			metric_label.EXCEPTION,
		},
	),
	metric_name.TASK_ACK_FAILED: metric_model.NewMetricDetails(
		metric_name.TASK_ACK_FAILED,
		metric_documentation.TASK_ACK_FAILED,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
		},
	),
	metric_name.TASK_ACK_ERROR: metric_model.NewMetricDetails(
		metric_name.TASK_ACK_ERROR,
		metric_documentation.TASK_ACK_ERROR,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
			metric_label.EXCEPTION,
		},
	),
	metric_name.TASK_UPDATE_ERROR: metric_model.NewMetricDetails(
		metric_name.TASK_UPDATE_ERROR,
		metric_documentation.TASK_UPDATE_ERROR,
		[]metric_label.MetricLabel{
			metric_label.TASK_TYPE,
			metric_label.EXCEPTION,
		},
	),
	metric_name.EXTERNAL_PAYLOAD_USED: metric_model.NewMetricDetails(
		metric_name.EXTERNAL_PAYLOAD_USED,
		metric_documentation.EXTERNAL_PAYLOAD_USED,
		[]metric_label.MetricLabel{
			metric_label.ENTITY_NAME,
			metric_label.OPERATION,
			metric_label.PAYLOAD_TYPE,
		},
	),
	metric_name.WORKFLOW_START_ERROR: metric_model.NewMetricDetails(
		metric_name.WORKFLOW_START_ERROR,
		metric_documentation.WORKFLOW_START_ERROR,
		[]metric_label.MetricLabel{
			metric_label.WORKFLOW_TYPE,
			metric_label.EXCEPTION,
		},
	),
}

func init() {
	for metricName, metricDetails := range counterTemplates {
		counterByName[metricName] = newCounter(metricDetails)
		prometheus.MustRegister(counterByName[metricName])
	}
}

func IncrementTaskPoll(taskType string) {
	incrementCounter(
		metric_name.TASK_POLL,
		[]string{
			taskType,
		},
	)
}

func IncrementTaskExecutionQueueFull(taskType string) {
	incrementCounter(
		metric_name.TASK_EXECUTION_QUEUE_FULL,
		[]string{
			taskType,
		},
	)
}

func IncrementUncaughtException(message string) {
	incrementCounter(
		metric_name.THREAD_UNCAUGHT_EXCEPTION,
		[]string{
			message,
		},
	)
}

func IncrementTaskPollError(taskType string, err error) {
	incrementCounter(
		metric_name.TASK_POLL_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func IncrementTaskPaused(taskType string) {
	incrementCounter(
		metric_name.TASK_PAUSED,
		[]string{
			taskType,
		},
	)
}

func IncrementTaskExecuteError(taskType string, err error) {
	incrementCounter(
		metric_name.TASK_EXECUTE_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func IncrementTaskAckFailed(taskType string) {
	incrementCounter(
		metric_name.TASK_ACK_FAILED,
		[]string{
			taskType,
		},
	)
}

func IncrementTaskAckError(taskType string, err error) {
	incrementCounter(
		metric_name.TASK_ACK_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func IncrementTaskUpdateError(taskType string, err error) {
	incrementCounter(
		metric_name.TASK_UPDATE_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func IncrementExternalPayloadUsed(entityName string, operation string, payloadType string) {
	incrementCounter(
		metric_name.EXTERNAL_PAYLOAD_USED,
		[]string{
			entityName,
			operation,
			payloadType,
		},
	)
}

func IncrementWorkflowStartError(workflowType string, err error) {
	incrementCounter(
		metric_name.WORKFLOW_START_ERROR,
		[]string{
			workflowType,
			err.Error(),
		},
	)
}

func incrementCounter(metricName metric_name.MetricName, labelValues []string) {
	counter := getCounter(metricName, labelValues)
	if counter != nil {
		(*counter).Inc()
	}
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

func getCounter(metricName metric_name.MetricName, labelValues []string) *prometheus.Counter {
	counterVec, ok := counterByName[metricName]
	if !ok {
		return nil
	}
	counter, _ := counterVec.GetMetricWithLabelValues(
		labelValues...,
	)
	return &counter
}
