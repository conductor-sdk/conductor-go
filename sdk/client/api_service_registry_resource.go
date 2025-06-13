// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package client

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
	"net/url"
)

type ServiceRegistryResourceApiService struct {
	*APIClient
}

/*
ServiceRegistryResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param registryName
*/
func (a *ServiceRegistryResourceApiService) AddOrUpdateMethod(ctx context.Context, body model.ServiceMethod, registryName string) (*http.Response, error) {
	var fileBytes []byte

	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s/methods", registryName)

	resp, err := a.Post(ctx, path, body, &fileBytes)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
ServiceRegistryResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
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
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return CircuitBreakerTransitionResponse
*/
func (a *ServiceRegistryResourceApiService) CloseCircuitBreaker(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error) {
	var transitionResp model.CircuitBreakerTransitionResponse

	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s/circuit-breaker/close", name)

	resp, err := a.Post(ctx, path, nil, &transitionResp)
	if err != nil {
		return model.CircuitBreakerTransitionResponse{}, resp, err
	}
	return transitionResp, resp, nil
}

/*
ServiceRegistryResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
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

func (a *ServiceRegistryResourceApiService) Discover(ctx context.Context, name string, optionals *ServiceRegistryResourceApiDiscoverOpts) ([]model.ServiceMethod, *http.Response, error) {
	var returnValue []model.ServiceMethod

	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s/discover", name)

	queryParams := url.Values{}

	if optionals != nil && optionals.Create.IsSet() {
		queryParams.Add("create", parameterToString(optionals.Create.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &returnValue)
	if err != nil {
		return nil, resp, err
	}
	return returnValue, resp, nil
}

/*
ServiceRegistryResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param registryName
    @return []ProtoRegistryEntry
*/
func (a *ServiceRegistryResourceApiService) GetAllProtos(ctx context.Context, registryName string) ([]model.ProtoRegistryEntry, *http.Response, error) {
	var returnValue []model.ProtoRegistryEntry

	// create path and map variables
	path := fmt.Sprintf("/registry/service/protos/%s", registryName)

	resp, err := a.Get(ctx, path, nil, &returnValue)
	if err != nil {
		return nil, resp, err
	}
	return returnValue, resp, nil
}

/*
ServiceRegistryResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return CircuitBreakerTransitionResponse
*/
func (a *ServiceRegistryResourceApiService) GetCircuitBreakerStatus(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error) {
	var returnValue model.CircuitBreakerTransitionResponse

	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s/circuit-breaker/status", name)

	resp, err := a.Get(ctx, path, nil, &returnValue)
	if err != nil {
		return model.CircuitBreakerTransitionResponse{}, resp, err
	}
	return returnValue, resp, nil
}

/*
ServiceRegistryResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param registryName
  - @param filename
    @return string
*/
func (a *ServiceRegistryResourceApiService) GetProtoData(ctx context.Context, registryName string, filename string) (string, *http.Response, error) {
	var returnValue string

	// create path and map variables
	path := fmt.Sprintf("/registry/service/protos/%s/%s", registryName, filename)

	resp, err := a.Get(ctx, path, nil, &returnValue)
	if err != nil {
		return "", resp, err
	}
	return returnValue, resp, nil
}

/*
ServiceRegistryResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []ServiceRegistry
*/
func (a *ServiceRegistryResourceApiService) GetRegisteredServices(ctx context.Context) ([]model.ServiceRegistry, *http.Response, error) {
	var returnValue []model.ServiceRegistry

	// create path and map variables
	path := "/registry/service"

	resp, err := a.Get(ctx, path, nil, &returnValue)
	if err != nil {
		return nil, resp, err
	}
	return returnValue, resp, nil
}

/*
ServiceRegistryResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return ServiceRegistry
*/
func (a *ServiceRegistryResourceApiService) GetService(ctx context.Context, name string) (model.ServiceRegistry, *http.Response, error) {
	var returnValue model.ServiceRegistry

	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s", name)
	resp, err := a.Get(ctx, path, nil, &returnValue)
	if err != nil {
		return model.ServiceRegistry{}, resp, err
	}
	return returnValue, resp, nil
}

/*
ServiceRegistryResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return CircuitBreakerTransitionResponse
*/
func (a *ServiceRegistryResourceApiService) OpenCircuitBreaker(ctx context.Context, name string) (model.CircuitBreakerTransitionResponse, *http.Response, error) {
	var returnValue model.CircuitBreakerTransitionResponse

	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s/circuit-breaker/open", name)

	resp, err := a.Post(ctx, path, nil, &returnValue)
	if err != nil {
		return model.CircuitBreakerTransitionResponse{}, resp, err
	}
	return returnValue, resp, nil
}

/*
ServiceRegistryResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
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
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
*/
func (a *ServiceRegistryResourceApiService) RemoveService(ctx context.Context, name string) (*http.Response, error) {
	// create path and map variables
	path := fmt.Sprintf("/registry/service/%s", name)

	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
ServiceRegistryResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param registryName
  - @param filename
*/
func (a *ServiceRegistryResourceApiService) SetProtoData(ctx context.Context, body string, registryName string, filename string) (*http.Response, error) {

	// create path and map variables
	path := fmt.Sprintf("/registry/service/protos/%s/%s", registryName, filename)

	resp, err := a.PostWithContentType(ctx, path, body, "application/octet-stream", nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
