package metrics

import (
	"github.com/netflix/conductor/client/go/metrics/metric_model/metric_name"
	"github.com/netflix/conductor/client/go/metrics/metrics_counter"
	"github.com/netflix/conductor/client/go/metrics/metrics_gauge"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsCollector struct {
	counterByName map[metric_name.MetricName]*prometheus.CounterVec
	gaugeByName   map[metric_name.MetricName]*prometheus.GaugeVec
}

func NewMetricsCollector() *MetricsCollector {
	metricsCollector := new(MetricsCollector)
	metricsCollector.counterByName = metrics_counter.NewCounterByName()
	metricsCollector.gaugeByName = metrics_gauge.NewGaugeByName()
	return metricsCollector
}

func (c *MetricsCollector) IncrementTaskPoll(taskType string) {
	c.incrementCounter(
		metric_name.TASK_POLL,
		[]string{
			taskType,
		},
	)
}

func (c *MetricsCollector) IncrementTaskExecutionQueueFull(taskType string) {
	c.incrementCounter(
		metric_name.TASK_EXECUTION_QUEUE_FULL,
		[]string{
			taskType,
		},
	)
}

func (c *MetricsCollector) IncrementUncaughtException() {
	c.incrementCounter(
		metric_name.THREAD_UNCAUGHT_EXCEPTION,
		[]string{},
	)
}

func (c *MetricsCollector) IncrementTaskPollError(taskType string, err error) {
	c.incrementCounter(
		metric_name.TASK_POLL_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func (c *MetricsCollector) IncrementTaskPaused(taskType string) {
	c.incrementCounter(
		metric_name.TASK_PAUSED,
		[]string{
			taskType,
		},
	)
}

func (c *MetricsCollector) IncrementTaskExecuteError(taskType string, err error) {
	c.incrementCounter(
		metric_name.TASK_EXECUTE_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func (c *MetricsCollector) IncrementTaskAckFailed(taskType string) {
	c.incrementCounter(
		metric_name.TASK_ACK_FAILED,
		[]string{
			taskType,
		},
	)
}

func (c *MetricsCollector) IncrementTaskAckError(taskType string, err error) {
	c.incrementCounter(
		metric_name.TASK_ACK_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func (c *MetricsCollector) IncrementTaskUpdateError(taskType string, err error) {
	c.incrementCounter(
		metric_name.TASK_UPDATE_ERROR,
		[]string{
			taskType,
			err.Error(),
		},
	)
}

func (c *MetricsCollector) IncrementExternalPayloadUsed(entityName string, operation string, payloadType string) {
	c.incrementCounter(
		metric_name.EXTERNAL_PAYLOAD_USED,
		[]string{
			entityName,
			operation,
			payloadType,
		},
	)
}

func (c *MetricsCollector) IncrementWorkflowStartError(workflowType string, err error) {
	c.incrementCounter(
		metric_name.WORKFLOW_START_ERROR,
		[]string{
			workflowType,
			err.Error(),
		},
	)
}

func (c *MetricsCollector) RecordWorkflowInputPayloadSize(workflowType string, version string, payloadSize float64) {
	c.setGauge(
		metric_name.WORKFLOW_INPUT_SIZE,
		[]string{
			workflowType,
			version,
		},
		payloadSize,
	)
}

func (c *MetricsCollector) RecordTaskResultPayloadSize(taskType string, payloadSize float64) {
	c.setGauge(
		metric_name.TASK_RESULT_SIZE,
		[]string{
			taskType,
		},
		payloadSize,
	)
}

func (c *MetricsCollector) RecordTaskPollTime(taskType string, timeSpent float64) {
	c.setGauge(
		metric_name.TASK_POLL_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
}

func (c *MetricsCollector) RecordTaskExecuteTime(taskType string, timeSpent float64) {
	c.setGauge(
		metric_name.TASK_EXECUTE_TIME,
		[]string{
			taskType,
		},
		timeSpent,
	)
}

func (c *MetricsCollector) incrementCounter(metricName metric_name.MetricName, labelValues []string) {
	counter := c.getCounter(metricName, labelValues)
	if counter != nil {
		(*counter).Inc()
	}
}

func (c *MetricsCollector) setGauge(metricName metric_name.MetricName, labelValues []string, value float64) {
	gauge := c.getGauge(metricName, labelValues)
	if gauge != nil {
		(*gauge).Set(value)
	}
}

func (c *MetricsCollector) getCounter(metricName metric_name.MetricName, labelValues []string) *prometheus.Counter {
	if counterVec, ok := c.counterByName[metricName]; ok {
		counter, _ := counterVec.GetMetricWithLabelValues(
			labelValues...,
		)
		return &counter
	}
	return nil
}

func (c *MetricsCollector) getGauge(metricName metric_name.MetricName, labelValues []string) *prometheus.Gauge {
	if gaugeVec, ok := c.gaugeByName[metricName]; ok {
		gauge, _ := gaugeVec.GetMetricWithLabelValues(
			labelValues...,
		)
		return &gauge
	}
	return nil
}
