package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type ServiceRegistryClient interface {
	AddOrUpdateMethod(ctx context.Context, body model.ServiceMethod, registryName string) (*http.Response, error)
	AddOrUpdateService(ctx context.Context, body model.ServiceRegistry) (*http.Response, error)
	CloseCircuitBreaker(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error)
	DeleteProto(ctx context.Context, registryName string, filename string) (*http.Response, error)
	Discover(ctx context.Context, name string, optionals *ServiceRegistryResourceApiDiscoverOpts) ([]model.ServiceMethod, *http.Response, error)
	GetAllProtos(ctx context.Context, registryName string) ([]model.ProtoRegistryEntry, *http.Response, error)
	GetCircuitBreakerStatus(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error)
	GetProtoData(ctx context.Context, registryName string, filename string) (string, *http.Response, error)
	GetRegisteredServices(ctx context.Context) ([]model.ServiceRegistry, *http.Response, error)
	GetService(ctx context.Context, name string) (model.ServiceRegistry, *http.Response, error)
	OpenCircuitBreaker(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error)
	RemoveMethod(ctx context.Context, registryName string, serviceName string, method string, methodType string) (*http.Response, error)
	RemoveService(ctx context.Context, name string) (*http.Response, error)
	SetProtoData(ctx context.Context, body string, registryName string, filename string) (*http.Response, error)
}

func NewServiceRegistryClient(apiClient *APIClient) ServiceRegistryClient {
	return &ServiceRegistryResourceApiService{apiClient}
}
