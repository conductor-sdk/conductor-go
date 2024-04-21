package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type MetadataClient interface {
	RegisterWorkflowDef(ctx context.Context, overwrite bool, body model.WorkflowDef) (*http.Response, error)
	RegisterWorkflowDefWithTags(ctx context.Context, overwrite bool, body model.WorkflowDef, tags []model.MetadataTag) (*http.Response, error)
	Get(ctx context.Context, name string, localVarOptionals *MetadataResourceApiGetOpts) (model.WorkflowDef, *http.Response, error)
	GetAll(ctx context.Context) ([]model.WorkflowDef, *http.Response, error)
	GetTaskDef(ctx context.Context, tasktype string) (model.TaskDef, *http.Response, error)
	GetTaskDefs(ctx context.Context) ([]model.TaskDef, *http.Response, error)
	UpdateTaskDef(ctx context.Context, body model.TaskDef) (*http.Response, error)
	UpdateTaskDefWithTags(ctx context.Context, body model.TaskDef, tags []model.MetadataTag, overwriteTags bool) (*http.Response, error)
	RegisterTaskDef(ctx context.Context, body []model.TaskDef) (*http.Response, error)
	RegisterTaskDefWithTags(ctx context.Context, body model.TaskDef, tags []model.MetadataTag) (*http.Response, error)
	UnregisterTaskDef(ctx context.Context, tasktype string) (*http.Response, error)
	UnregisterWorkflowDef(ctx context.Context, name string, version int32) (*http.Response, error)
	Update(ctx context.Context, body []model.WorkflowDef) (*http.Response, error)
	UpdateWorkflowDefWithTags(ctx context.Context, body model.WorkflowDef, tags []model.MetadataTag, overwriteTags bool) (*http.Response, error)
	GetTagsForWorkflowDef(ctx context.Context, name string) ([]model.MetadataTag, error)
	GetTagsForTaskDef(ctx context.Context, tasktype string) ([]model.MetadataTag, error)
}

func NewMetadataClient(apiClient *APIClient) MetadataClient {
	return &MetadataResourceApiService{apiClient}
}
