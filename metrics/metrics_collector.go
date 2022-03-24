package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricsCollector struct {
	counters [1]prometheus.Counter
}

func NewMetricsCollector() *MetricsCollector {
	metricsCollector := new(MetricsCollector)
	metricsCollector.counters = newCounters()
	for _, counter := range metricsCollector.counters {
		prometheus.MustRegister(counter)
	}
	return metricsCollector
}

func (c *MetricsCollector) IncrementCounter() {
	c.counters[0].Inc()
}

func newCounters() [1]prometheus.Counter {
	return [...]prometheus.Counter{
		prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "counter",
				Help: "Counter example...",
			},
		),
	}
}
