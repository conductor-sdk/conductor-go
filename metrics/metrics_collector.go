package metrics

import "github.com/prometheus/client_golang/prometheus"

type MetricsCollector struct {
	counterByName map[MetricName]*prometheus.CounterVec
}

func newCounterTemplates() map[MetricName]*MetricDetails {
	return map[MetricName]*MetricDetails{
		TASK_POLL: NewMetricDetails(
			"task_poll",
			"Counter for TaskPoll",
			[]string{"task_type"},
		),
	}
}

func newCounterByName(counterTemplates map[MetricName]*MetricDetails) map[MetricName]*prometheus.CounterVec {
	counterByName := map[MetricName]*prometheus.CounterVec{}
	for metricName, metricDetails := range counterTemplates {
		counterByName[metricName] = newCounter(metricDetails)
	}
	return counterByName
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

func NewMetricsCollector() *MetricsCollector {
	metricsCollector := new(MetricsCollector)
	counterTemplates := newCounterTemplates()
	metricsCollector.counterByName = newCounterByName(counterTemplates)
	for _, counter := range metricsCollector.counterByName {
		prometheus.MustRegister(counter)
	}
	return metricsCollector
}

func (c *MetricsCollector) IncrementCounter(metricName MetricName, labelValues []string) {
	counter := c.getCounter(metricName, labelValues)
	(*counter).Inc()
}

func (c *MetricsCollector) getCounter(metricName MetricName, labelValues []string) *prometheus.Counter {
	if counterVec, ok := c.counterByName[metricName]; ok {
		counter, _ := counterVec.GetMetricWithLabelValues(
			labelValues...,
		)
		return &counter
	}
	return nil
}
