package client

import (
	"context"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"net/http"
)

type GetIntegrationProvidersOpts struct {
	Category   optional.String
	ActiveOnly optional.Bool
}

type IntegrationClient interface {

	// Integration Provider Management
	//

	GetIntegrationProviders(ctx context.Context, opts *GetIntegrationProvidersOpts) ([]integration.Integration, *http.Response, error)
	GetIntegrationProvider(ctx context.Context, name string) (integration.Integration, *http.Response, error)
	SaveIntegrationProvider(ctx context.Context, update integration.IntegrationUpdate, name string) (*http.Response, error)
	DeleteIntegrationProvider(ctx context.Context, name string) (*http.Response, error)
	GetIntegrationProviderDefs(ctx context.Context) ([]model.IntegrationDef, *http.Response, error)
	// Tag management
	//

	GetTagsForIntegrationProvider(ctx context.Context, name string) ([]model.TagObject, *http.Response, error)
	UpdateTagForIntegrationProvider(ctx context.Context, tags []model.TagObject, name string) (*http.Response, error)
	DeleteTagForIntegrationProvider(ctx context.Context, tags []model.TagObject, name string) (*http.Response, error)
	GetTagsForIntegration(ctx context.Context, name string, model string) ([]model.TagObject, *http.Response, error)
	UpdateTagForIntegration(ctx context.Context, tags []model.TagObject, name string, model string) (*http.Response, error)
	DeleteTagForIntegration(ctx context.Context, tags []model.TagObject, name string, model string) (*http.Response, error)

	// Integration Management
	//

	GetIntegrationApis(ctx context.Context, name string, activeOnly optional.Bool) ([]integration.IntegrationApi, *http.Response, error)
	GetIntegrationApi(ctx context.Context, name string, model string) (integration.IntegrationApi, *http.Response, error)
	SaveIntegrationApi(ctx context.Context, update integration.IntegrationApiUpdate, name string, model string) (*http.Response, error)
	DeleteIntegrationApi(ctx context.Context, name string, model string) (*http.Response, error)

	// LLM specific
	//

	GetPromptsWithIntegration(ctx context.Context, integrationName string, model string) ([]integration.PromptTemplate, *http.Response, error)
	AssociatePromptWithIntegration(ctx context.Context, integrationName string, model string, promptName string) (*http.Response, error)
	GetTokenUsageForIntegration(ctx context.Context, integrationName string, model string) (int32, *http.Response, error)
	GetTokenUsageForIntegrationProvider(ctx context.Context, name string) (map[string]string, *http.Response, error)

	GetAllIntegrations(ctx context.Context, optionals *IntegrationResourceApiGetAllIntegrationsOpts) ([]model.Integration, *http.Response, error)
	RecordEventStats(ctx context.Context, body []model.EventLog, type_ string) (*http.Response, error)
}

func NewIntegrationClient(apiClient *APIClient) IntegrationClient {
	return &IntegrationResourceApiService{apiClient}
}
