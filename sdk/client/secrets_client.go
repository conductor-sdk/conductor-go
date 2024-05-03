package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type SecretsClient interface {
	ClearLocalCache(ctx context.Context) (map[string]string, *http.Response, error)
	ClearRedisCache(ctx context.Context) (map[string]string, *http.Response, error)
	DeleteSecret(ctx context.Context, key string) (interface{}, *http.Response, error)
	DeleteTagForSecret(ctx context.Context, body []model.Tag, key string) (*http.Response, error)
	GetSecret(ctx context.Context, key string) (string, *http.Response, error)
	GetTags(ctx context.Context, key string) ([]model.Tag, *http.Response, error)
	ListAllSecretNames(ctx context.Context) ([]string, *http.Response, error)
	ListSecretsThatUserCanGrantAccessTo(ctx context.Context) ([]string, *http.Response, error)
	ListSecretsWithTagsThatUserCanGrantAccessTo(ctx context.Context) ([]model.Secret, *http.Response, error)
	PutSecret(ctx context.Context, body string, key string) (interface{}, *http.Response, error)
	PutTagForSecret(ctx context.Context, body []model.Tag, key string) (*http.Response, error)
	SecretExists(ctx context.Context, key string) (interface{}, *http.Response, error)
}

func NewSecretsClient(client *APIClient) SecretsClient {
	return &SecretResourceApiService{client}
}
