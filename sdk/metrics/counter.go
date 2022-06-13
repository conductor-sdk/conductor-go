package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var counterByName = map[MetricName]*prometheus.CounterVec{}

var counterTemplates = map[MetricName]*MetricDetails{
	TASK_POLL: NewMetricDetails(
		TASK_POLL,
		TASK_POLL_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_EXECUTION_QUEUE_FULL: NewMetricDetails(
		TASK_EXECUTION_QUEUE_FULL,
		TASK_EXECUTION_QUEUE_FULL_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	THREAD_UNCAUGHT_EXCEPTION: NewMetricDetails(
		THREAD_UNCAUGHT_EXCEPTION,
		THREAD_UNCAUGHT_EXCEPTION_DOC,
		[]MetricLabel{
			EXCEPTION,
		},
	),
	TASK_POLL_ERROR: NewMetricDetails(
		TASK_POLL_ERROR,
		TASK_POLL_ERROR_DOC,
		[]MetricLabel{
			TASK_TYPE,
			EXCEPTION,
		},
	),
	TASK_PAUSED: NewMetricDetails(
		TASK_PAUSED,
		TASK_PAUSED_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_EXECUTE_ERROR: NewMetricDetails(
		TASK_EXECUTE_ERROR,
		TASK_EXECUTE_ERROR_DOC,
		[]MetricLabel{
			TASK_TYPE,
			EXCEPTION,
		},
	),
	TASK_ACK_FAILED: NewMetricDetails(
		TASK_ACK_FAILED,
		TASK_ACK_FAILED_DOC,
		[]MetricLabel{
			TASK_TYPE,
		},
	),
	TASK_ACK_ERROR: NewMetricDetails(
		TASK_ACK_ERROR,
		TASK_ACK_ERROR_DOC,
		[]MetricLabel{
			TASK_TYPE,
			EXCEPTION,
		},
	),
	TASK_UPDATE_ERROR: NewMetricDetails(
		TASK_UPDATE_ERROR,
		TASK_UPDATE_ERROR_DOC,
		[]MetricLabel{
			TASK_TYPE,
			EXCEPTION,
		},
	),
	EXTERNAL_PAYLOAD_USED: NewMetricDetails(
		EXTERNAL_PAYLOAD_USED,
		EXTERNAL_PAYLOAD_USED_DOC,
		[]MetricLabel{
			ENTITY_NAME,
			OPERATION,
			PAYLOAD_TYPE,
		},
	),
	WORKFLOW_START_ERROR: NewMetricDetails(
		WORKFLOW_START_ERROR,
		WORKFLOW_START_ERROR_DOC,
		[]MetricLabel{
			WORKFLOW_TYPE,
			EXCEPTION,
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
		TASK_POLL,
		[]string{
			taskType,
		},
	)
}

func IncrementTaskExecutionQueueFull(taskType string) {
	incrementCounter(
		TASK_EXECUTION_QUEUE_FULL,
		[]string{
			taskType,
		},
	)
}

func IncrementUncaughtException(message string) {
	incrementCounter(
		THREAD_UNCAUGHT_EXCEPTION,
		[]string{
			message,
		},
	)
}

func IncrementTaskPollError(taskType string, err error) {
	incrementCounter(
		TASK_POLL_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func IncrementTaskPaused(taskType string) {
	incrementCounter(
		TASK_PAUSED,
		[]string{
			taskType,
		},
	)
}

func IncrementTaskExecuteError(taskType string, err error) {
	incrementCounter(
		TASK_EXECUTE_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func IncrementTaskAckFailed(taskType string) {
	incrementCounter(
		TASK_ACK_FAILED,
		[]string{
			taskType,
		},
	)
}

func IncrementTaskAckError(taskType string, err error) {
	incrementCounter(
		TASK_ACK_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func IncrementTaskUpdateError(taskType string, err error) {
	incrementCounter(
		TASK_UPDATE_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func IncrementExternalPayloadUsed(entityName string, operation string, payloadType string) {
	incrementCounter(
		EXTERNAL_PAYLOAD_USED,
		[]string{
			entityName,
			operation,
			payloadType,
		},
	)
}

func IncrementWorkflowStartError(workflowType string, err error) {
	incrementCounter(
		WORKFLOW_START_ERROR,
		[]string{
			workflowType,
			err.Error(),
		},
	)
}

func incrementCounter(metricName MetricName, labelValues []string) {
	counter := getCounter(metricName, labelValues)
	if counter != nil {
		(*counter).Inc()
	}
}

func newCounter(metricDetails *MetricDetails) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metricDetails.Name,
			Help: metricDetails.Description,
		},
		metricDetails.Labels,
	)
}

func getCounter(metricName MetricName, labelValues []string) *prometheus.Counter {
	counterVec, ok := counterByName[metricName]
	if !ok {
		return nil
	}
	counter, _ := counterVec.GetMetricWithLabelValues(
		labelValues...,
	)
	return &counter
}