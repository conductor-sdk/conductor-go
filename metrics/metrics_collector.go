package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsCollector struct {
	counter prometheus.Counter
}

func NewMetricsCollector() *MetricsCollector {
	metricsCollector := new(MetricsCollector)
	metricsCollector.counter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "counter",
			Help: "Counter example...",
		},
	)
	prometheus.MustRegister(metricsCollector.counter)
	return metricsCollector
}

func (c *MetricsCollector) IncrementCounter() {
	c.counter.Inc()
}
