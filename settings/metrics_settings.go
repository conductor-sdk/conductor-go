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
	metricsSettings := new(MetricsSettings)
	metricsSettings.ApiEndpoint = apiEndpoint
	metricsSettings.Port = port
	return metricsSettings
}
