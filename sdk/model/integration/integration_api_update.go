package integration

type IntegrationApiUpdate struct {
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Enabled       bool                   `json:"enabled,omitempty"`
}
