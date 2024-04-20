package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"net/http"
)

type ApplicationClient interface {
	AddRoleToApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error)
	CreateAccessKey(ctx context.Context, id string) (interface{}, *http.Response, error)
	CreateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest) (interface{}, *http.Response, error)
	DeleteAccessKey(ctx context.Context, applicationId string, keyId string) (interface{}, *http.Response, error)
	DeleteApplication(ctx context.Context, id string) (interface{}, *http.Response, error)
	DeleteTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error)
	GetAccessKeys(ctx context.Context, id string) (interface{}, *http.Response, error)
	GetAppByAccessKeyId(ctx context.Context, accessKeyId string) (interface{}, *http.Response, error)
	GetApplication(ctx context.Context, id string) (interface{}, *http.Response, error)
	GetTagsForApplication(ctx context.Context, id string) ([]model.Tag, *http.Response, error)
	ListApplications(ctx context.Context) ([]rbac.ConductorApplication, *http.Response, error)
	PutTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error)
	RemoveRoleFromApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error)
	ToggleAccessKeyStatus(ctx context.Context, applicationId string, keyId string) (interface{}, *http.Response, error)
	UpdateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest, id string) (interface{}, *http.Response, error)
}

func NewApplicationClient(apiClient *APIClient) *ApplicationResourceApiService {
	return &ApplicationResourceApiService{apiClient}
}
