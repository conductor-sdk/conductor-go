package integration

import "github.com/conductor-sdk/conductor-go/sdk/model"

type PromptTemplate struct {
	CreatedBy    string            `json:"createdBy,omitempty"`
	CreatedOn    int64             `json:"createdOn,omitempty"`
	Description  string            `json:"description,omitempty"`
	Integrations []string          `json:"integrations,omitempty"`
	Name         string            `json:"name,omitempty"`
	Tags         []model.TagObject `json:"tags,omitempty"`
	Template     string            `json:"template,omitempty"`
	UpdatedBy    string            `json:"updatedBy,omitempty"`
	UpdatedOn    int64             `json:"updatedOn,omitempty"`
	Variables    []string          `json:"variables,omitempty"`
}
