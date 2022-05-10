package settings

type HttpSettings struct {
	BaseUrl string
	Headers map[string]string
}

func NewHttpDefaultSettings() *HttpSettings {
	return NewHttpSettings(
		"http://localhost:8080/api",
	)
}

func NewHttpSettings(baseUrl string) *HttpSettings {
	return &HttpSettings{
		BaseUrl: baseUrl,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
	}
}
