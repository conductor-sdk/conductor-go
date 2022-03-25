package metrics

import (
	"github.com/netflix/conductor/client/go/metrics/enum/metric_name"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsCollector struct {
	counterByName map[metric_name.MetricName]*prometheus.CounterVec
	gaugeByName   map[metric_name.MetricName]*prometheus.GaugeVec
}

func newCounterTemplates() map[metric_name.MetricName]*MetricDetails {
	return map[metric_name.MetricName]*MetricDetails{
		metric_name.TASK_POLL: NewMetricDetails(
			"task_poll",
			"Counter for TaskPoll",
			[]string{"task_type"},
		),
	}
}

func newGaugeTemplates() map[metric_name.MetricName]*MetricDetails {
	return map[metric_name.MetricName]*MetricDetails{
		metric_name.TASK_POLL_TIME: NewMetricDetails(
			"task_poll_time",
			"Gauge for Task poll time",
			[]string{"task_type"},
		),
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

func newGauge(metricDetails *MetricDetails) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricDetails.Name,
			Help: metricDetails.Description,
		},
		metricDetails.Labels,
	)
}

func newCounterByName(counterTemplates map[metric_name.MetricName]*MetricDetails) map[metric_name.MetricName]*prometheus.CounterVec {
	counterByName := map[metric_name.MetricName]*prometheus.CounterVec{}
	for metricName, metricDetails := range counterTemplates {
		counterByName[metricName] = newCounter(metricDetails)
	}
	return counterByName
}

func newGaugeByName(gaugeTemplates map[metric_name.MetricName]*MetricDetails) map[metric_name.MetricName]*prometheus.GaugeVec {
	gaugeByName := map[metric_name.MetricName]*prometheus.GaugeVec{}
	for metricName, metricDetails := range gaugeTemplates {
		gaugeByName[metricName] = newGauge(metricDetails)
	}
	return gaugeByName
}

func NewMetricsCollector() *MetricsCollector {
	metricsCollector := new(MetricsCollector)

	counterTemplates := newCounterTemplates()
	metricsCollector.counterByName = newCounterByName(counterTemplates)

	gaugeTemplates := newGaugeTemplates()
	metricsCollector.gaugeByName = newGaugeByName(gaugeTemplates)

	metricsCollector.registerMetricsCollectors()

	return metricsCollector
}

func (c *MetricsCollector) registerMetricsCollectors() {
	for _, counter := range c.counterByName {
		prometheus.MustRegister(counter)
	}
	for _, gauge := range c.gaugeByName {
		prometheus.MustRegister(gauge)
	}
}

func (c *MetricsCollector) IncrementCounter(metricName metric_name.MetricName, labelValues []string) {
	counter := c.getCounter(metricName, labelValues)
	(*counter).Inc()
}

func (c *MetricsCollector) SetGauge(metricName metric_name.MetricName, labelValues []string, value float64) {
	gauge := c.getGauge(metricName, labelValues)
	(*gauge).Set(value)
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
