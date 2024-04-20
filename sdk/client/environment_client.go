package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type EnvironmentClient interface {
	CreateOrUpdateEnvVariable(ctx context.Context, body string, key string) (*http.Response, error)
	DeleteEnvVariable(ctx context.Context, key string) (string, *http.Response, error)
	DeleteTagForEnvVar(ctx context.Context, body []model.Tag, name string) (*http.Response, error)
	Get(ctx context.Context, key string) (string, *http.Response, error)
	GetAll(ctx context.Context) ([]model.EnvironmentVariable, *http.Response, error)
	GetTagsForEnvVar(ctx context.Context, name string) ([]model.Tag, *http.Response, error)
	PutTagForEnvVar(ctx context.Context, body []model.Tag, name string) (*http.Response, error)
}

func NewEnvironmentClient(apiClient *APIClient) EnvironmentClient {
	return &EnvironmentResourceApiService{apiClient}
}
