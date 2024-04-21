package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"net/http"
)

type PromptClient interface {
	DeleteMessageTemplate(ctx context.Context, name string) (*http.Response, error)
	DeleteTagForPromptTemplate(ctx context.Context, tags []model.Tag, name string) (*http.Response, error)
	GetMessageTemplate(ctx context.Context, name string) (*integration.PromptTemplate, *http.Response, error)
	GetMessageTemplates(ctx context.Context) ([]integration.PromptTemplate, *http.Response, error)
	GetTagsForPromptTemplate(ctx context.Context, name string) ([]model.Tag, *http.Response, error)
	PutTagForPromptTemplate(ctx context.Context, tags []model.Tag, name string) (*http.Response, error)
	SaveMessageTemplate(ctx context.Context, templateText string, description string, name string, optionals *PromptResourceApiSaveMessageTemplateOpts) (*http.Response, error)
	TestMessageTemplate(ctx context.Context, request model.PromptTemplateTestRequest) (string, *http.Response, error)
}

func NewPromptClient(apiClient *APIClient) PromptClient {
	return &PromptResourceApiService{apiClient}
}
