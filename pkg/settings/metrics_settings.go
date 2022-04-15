package settings

type MetricsSettings struct {
	ApiEndpoint string
	Port        int
}

func NewDefaultMetricsSettings() *MetricsSettings {
	return NewMetricsSettings(
		"/metrics",
		2112,
	)
}

func NewMetricsSettings(apiEndpoint string, port int) *MetricsSettings {
	return &MetricsSettings{
		ApiEndpoint: apiEndpoint,
		Port:        port,
	}
}
