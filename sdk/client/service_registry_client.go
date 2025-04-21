package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

// ServiceRegistryClient interface defines methods for interacting with the service registry
type ServiceRegistryClient interface {
	// Method management
	AddOrUpdateMethod(ctx context.Context, method model.ServiceMethod, registryName string) (*http.Response, error)
	RemoveMethod(ctx context.Context, registryName string, serviceName string, method string, methodType string) (*http.Response, error)

	// Service management
	AddOrUpdateService(ctx context.Context, registry model.ServiceRegistry) (*http.Response, error)
	GetService(ctx context.Context, name string) (model.ServiceRegistry, *http.Response, error)
	GetRegisteredServices(ctx context.Context) ([]model.ServiceRegistry, *http.Response, error)
	RemoveService(ctx context.Context, name string) (*http.Response, error)

	// Service discovery
	Discover(ctx context.Context, name string, optionals *ServiceRegistryResourceApiDiscoverOpts) ([]model.ServiceMethod, *http.Response, error)

	// Circuit breaker operations
	OpenCircuitBreaker(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error)
	CloseCircuitBreaker(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error)
	GetCircuitBreakerStatus(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error)

	// Proto management
	GetAllProtos(ctx context.Context, registryName string) ([]model.ProtoRegistryEntry, *http.Response, error)
	GetProtoData(ctx context.Context, registryName string, filename string) ([]byte, *http.Response, error)
	SetProtoData(ctx context.Context, body []byte, registryName string, filename string) (*http.Response, error)
	DeleteProto(ctx context.Context, registryName string, filename string) (*http.Response, error)
}

// NewServiceRegistryClient creates a new instance of the ServiceRegistryClient
func NewServiceRegistryClient(apiClient *APIClient) ServiceRegistryClient {
	return &ServiceRegistryResourceApiService{apiClient}
}
