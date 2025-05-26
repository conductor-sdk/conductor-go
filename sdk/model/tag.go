package model

type Tag struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	// Deprecated: since 11/21/23
	Type_ string `json:"type,omitempty"`
}