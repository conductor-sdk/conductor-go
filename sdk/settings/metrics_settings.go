package settings

//MetricsSettings configures the prometheus metrics for worker SDK
type MetricsSettings struct {
	ApiEndpoint string
	Port        int
}

//NewDefaultMetricsSettings creates an endpoint at /metrics on port 2112
func NewDefaultMetricsSettings() *MetricsSettings {
	return NewMetricsSettings(
		"/metrics",
		2112,
	)
}

//NewMetricsSettings new metrics settings with endpoint and port
func NewMetricsSettings(apiEndpoint string, port int) *MetricsSettings {
	return &MetricsSettings{
		ApiEndpoint: apiEndpoint,
		Port:        port,
	}
}
