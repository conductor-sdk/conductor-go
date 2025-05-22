package model

type Schema struct {
	DefaultValue *interface{} `json:"defaultValue,omitempty"`
	Format       string       `json:"format,omitempty"`
	Type_        string       `json:"type,omitempty"`
}
