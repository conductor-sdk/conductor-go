package model

type RequestParam struct {
	Name     string  `json:"name,omitempty"`
	Required bool    `json:"required,omitempty"`
	Schema   *Schema `json:"schema,omitempty"`
	Type_    string  `json:"type,omitempty"`
}
