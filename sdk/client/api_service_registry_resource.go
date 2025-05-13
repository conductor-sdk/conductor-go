package client

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type ServiceRegistryResourceApiService struct {
	*APIClient
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param registryName
*/
func (a *ServiceRegistryResourceApiService) AddOrUpdateMethod(ctx context.Context, body model.ServiceMethod, registryName string) (*http.Response, error) {
	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s/methods", registryName)

	resp, err := a.Post(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *ServiceRegistryResourceApiService) AddOrUpdateService(ctx context.Context, body model.ServiceRegistry) (*http.Response, error) {
	// create path and map variables
	path := "/registry/service"

	resp, err := a.Post(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name

@return CircuitBreakerTransitionResponse
*/
func (a *ServiceRegistryResourceApiService) CloseCircuitBreaker(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error) {
	var result model.CircuitBreakerTransitionResponse

	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s/circuit-breaker/close", name)

	resp, err := a.Post(ctx, path, nil, &result)
	if err != nil {
		return model.CircuitBreakerTransitionResponse{}, resp, err
	}
	return result, resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param registryName
  - @param filename
*/
func (a *ServiceRegistryResourceApiService) DeleteProto(ctx context.Context, registryName string, filename string) (*http.Response, error) {
	// create path and map variables
	path := fmt.Sprintf("/registry/service/protos/%s/%s", registryName, filename)
	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
ServiceRegistryResourceApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name
 * @param optional nil or *ServiceRegistryResourceApiDiscoverOpts - Optional Parameters:
     * @param "Create" (optional.Bool) -
@return []ServiceMethod
*/

type ServiceRegistryResourceApiDiscoverOpts struct {
	Create optional.Bool
}

func (a *ServiceRegistryResourceApiService) Discover(ctx context.Context, name string, opts *ServiceRegistryResourceApiDiscoverOpts) ([]model.ServiceMethod, *http.Response, error) {
	var result []model.ServiceMethod

	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s/discover", name)

	queryParams := url.Values{}
	if opts != nil && opts.Create.IsSet() {
		queryParams.Add("create", parameterToString(opts.Create.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param registryName

@return []ProtoRegistryEntry
*/
func (a *ServiceRegistryResourceApiService) GetAllProtos(ctx context.Context, registryName string) ([]model.ProtoRegistryEntry, *http.Response, error) {
	var result []model.ProtoRegistryEntry

	// create path and map variables
	path := fmt.Sprintf("/registry/service/protos/%s", registryName)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name

@return CircuitBreakerTransitionResponse
*/
func (a *ServiceRegistryResourceApiService) GetCircuitBreakerStatus(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error) {
	var result model.CircuitBreakerTransitionResponse

	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s/circuit-breaker/status", name)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param registryName
  - @param filename

@return string
*/
func (a *ServiceRegistryResourceApiService) GetProtoData(ctx context.Context, registryName string, filename string) ([]byte, *http.Response, error) {
	var result []byte

	// create path and map variables
	path := fmt.Sprintf("/registry/service/protos/%s/%s", registryName, filename)

	acceptType := "application/octet-stream"
	resp, err := a.GetWithAcceptType(ctx, path, nil, acceptType, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []ServiceRegistry
*/
func (a *ServiceRegistryResourceApiService) GetRegisteredServices(ctx context.Context) ([]model.ServiceRegistry, *http.Response, error) {
	var result []model.ServiceRegistry

	// create path and map variables
	path := "/registry/service"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name

@return ServiceRegistry
*/
func (a *ServiceRegistryResourceApiService) GetService(ctx context.Context, name string) (model.ServiceRegistry, *http.Response, error) {
	var result model.ServiceRegistry

	// create path and map variables
	localVarPath := fmt.Sprintf("/registry/service/%s", name)
	resp, err := a.Get(ctx, localVarPath, nil, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name

@return CircuitBreakerTransitionResponse
*/
func (a *ServiceRegistryResourceApiService) OpenCircuitBreaker(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error) {
	var result model.CircuitBreakerTransitionResponse

	// create path and map variables
	localVarPath := fmt.Sprintf("/registry/service/%s/circuit-breaker/open", name)
	resp, err := a.Post(ctx, localVarPath, nil, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param registryName
  - @param serviceName
  - @param method
  - @param methodType
*/
func (a *ServiceRegistryResourceApiService) RemoveMethod(ctx context.Context, registryName string, serviceName string, method string, methodType string) (*http.Response, error) {
	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s/methods", registryName)

	queryParams := url.Values{}
	queryParams.Add("serviceName", parameterToString(serviceName, ""))
	queryParams.Add("method", parameterToString(method, ""))
	queryParams.Add("methodType", parameterToString(methodType, ""))

	resp, err := a.Delete(ctx, path, queryParams, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
*/
func (a *ServiceRegistryResourceApiService) RemoveService(ctx context.Context, name string) (*http.Response, error) {
	// create path and map variables
	localVarPath := fmt.Sprintf("/registry/service/%s", name)

	resp, err := a.Delete(ctx, localVarPath, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param registryName
  - @param filename
*/
func (a *ServiceRegistryResourceApiService) SetProtoData(ctx context.Context, body []byte, registryName string, filename string) (*http.Response, error) {
	// create path and map variables
	localVarPath := fmt.Sprintf("/registry/service/protos/%s/%s", registryName, filename)
	contentType := "application/octet-stream"

	resp, err := a.PostWithContentType(ctx, localVarPath, contentType, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
