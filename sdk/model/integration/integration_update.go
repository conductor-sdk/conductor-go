package integration

type IntegrationUpdate struct {
	Category      string                    `json:"category,omitempty"`
	Configuration map[ConfigKey]interface{} `json:"configuration,omitempty"`
	Description   string                    `json:"description,omitempty"`
	Enabled       bool                      `json:"enabled,omitempty"`
	Type_         string                    `json:"type,omitempty"`
}
