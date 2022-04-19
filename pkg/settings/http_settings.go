package settings

type HttpSettings struct {
	BaseUrl                 string
	Headers                 map[string]string
	ExternalStorageSettings *ExternalStorageSettings
}

func NewHttpDefaultSettings() *HttpSettings {
	return NewHttpSettings(
		"http://localhost:8080/api",
		nil,
	)
}

func NewHttpSettings(baseUrl string, externalStorageSettings *ExternalStorageSettings) *HttpSettings {
	return &HttpSettings{
		BaseUrl: baseUrl,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
		ExternalStorageSettings: externalStorageSettings,
	}
}
