package integration

// IntegrationDefFormField represents a form field for integration definitions
type IntegrationDefFormField struct {
	FieldType    string                    `json:"fieldType,omitempty"`
	ValueOptions []Option                  `json:"valueOptions,omitempty"`
	Label        string                    `json:"label,omitempty"`
	FieldName    ConfigKey                 `json:"fieldName,omitempty"`
	Value        string                    `json:"value,omitempty"`
	DefaultValue string                    `json:"defaultValue,omitempty"`
	DependsOn    []IntegrationDefFormField `json:"dependsOn,omitempty"`
	Description  string                    `json:"description,omitempty"`
	Optional     bool                      `json:"optional,omitempty"`
}

// Option represents a label-value pair for dropdown options
type Option struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

// Constants for IntegrationDefFormFieldType
const (
	FormFieldTypeDropdown = "DROPDOWN"
	FormFieldTypeText     = "TEXT"
	FormFieldTypePassword = "PASSWORD"
	FormFieldTypeFile     = "FILE"
)