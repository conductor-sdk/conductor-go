package model

type Health struct {
	Details      map[string]interface{} `json:"details,omitempty"`
	ErrorMessage string                 `json:"errorMessage,omitempty"`
	Healthy      bool                   `json:"healthy,omitempty"`
}
