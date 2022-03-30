package settings

type HttpSettings struct {
	BaseUrl     string
	BearerToken *string
	Headers     map[string]string
}

func NewHttpDefaultSettings() *HttpSettings {
	return NewHttpSettings("http://localhost:8080/api")
}

func NewHttpSettings(baseUrl string) *HttpSettings {
	settings := new(HttpSettings)
	settings.BaseUrl = baseUrl
	settings.BearerToken = nil
	settings.Headers = map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	return settings
}
