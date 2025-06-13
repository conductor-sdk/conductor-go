package integration

import "github.com/conductor-sdk/conductor-go/sdk/model"

type IntegrationApi struct {
	Api             string                 `json:"api,omitempty"`
	Configuration   map[string]interface{} `json:"configuration,omitempty"`
	CreatedBy       string                 `json:"createdBy,omitempty"`
	CreatedOn       int64                  `json:"createdOn,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Enabled         bool                   `json:"enabled,omitempty"`
	IntegrationName string                 `json:"integrationName,omitempty"`
	Tags            []model.TagObject      `json:"tags,omitempty"`
	UpdatedBy       string                 `json:"updatedBy,omitempty"`
	UpdatedOn       int64                  `json:"updatedOn,omitempty"`
}
