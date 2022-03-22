package settings

type HttpSettings struct {
	BaseUrl     string
	BearerToken *string
	Debug       bool
	Headers     map[string]string
}

func NewHttpSettings() *HttpSettings {
	return NewHttpSettingsWithBaseUrlAndDebug(
		"http://localhost:8080/api",
		true,
	)
}

func NewHttpSettingsWithBaseUrlAndDebug(baseUrl string, debug bool) *HttpSettings {
	settings := new(HttpSettings)
	settings.BaseUrl = baseUrl
	settings.BearerToken = nil
	settings.Debug = debug
	settings.Headers = map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	return settings
}
