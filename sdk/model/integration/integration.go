package integration

import "github.com/conductor-sdk/conductor-go/sdk/model"

type Integration struct {
	Category      string                 `json:"category,omitempty"`
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	CreatedBy     string                 `json:"createdBy,omitempty"`
	CreatedOn     int64                  `json:"createdOn,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Enabled       bool                   `json:"enabled,omitempty"`
	ModelsCount   int64                  `json:"modelsCount,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Tags          []model.TagObject      `json:"tags,omitempty"`
	Type_         string                 `json:"type,omitempty"`
	UpdatedBy     string                 `json:"updatedBy,omitempty"`
	UpdatedOn     int64                  `json:"updatedOn,omitempty"`
}
