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
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

type EventResourceApiService struct {
	*APIClient
}

/*
EventResourceApiService Add a new event handler.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *EventResourceApiService) AddEventHandler(ctx context.Context, body model.EventHandler) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/event"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarPostBody = &body
	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)
	if err != nil {
		return localVarHttpResponse, err
	}

	if !isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		newErr := NewGenericSwaggerError(localVarBody, localVarHttpResponse.Status, nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
EventResourceApiService Delete queue config by name
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param queueType
  - @param queueName
*/
func (a *EventResourceApiService) DeleteQueueConfig(ctx context.Context, queueType string, queueName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/event/queue/config/{queueType}/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"queueType"+"}", fmt.Sprintf("%v", queueType), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)
	if err != nil {
		return localVarHttpResponse, err
	}

	if !isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		newErr := NewGenericSwaggerError(localVarBody, localVarHttpResponse.Status, nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
EventResourceApiService Get all the event handlers
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []model.EventHandler
*/
func (a *EventResourceApiService) GetEventHandlers(ctx context.Context) ([]model.EventHandler, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []model.EventHandler
	)

	localVarPath := "/event"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		err = a.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	} else {
		newErr := NewGenericSwaggerError(localVarBody, localVarHttpResponse.Status, nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
EventResourceApiService Get event handlers for a given event
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param event
 * @param optional nil or *EventResourceApiGetEventHandlersForEventOpts - Optional Parameters:
     * @param "ActiveOnly" (optional.Bool) -
@return []model.EventHandler
*/

type EventResourceApiGetEventHandlersForEventOpts struct {
	ActiveOnly optional.Bool
}

func (a *EventResourceApiService) GetEventHandlersForEvent(ctx context.Context, event string, localVarOptionals *EventResourceApiGetEventHandlersForEventOpts) ([]model.EventHandler, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []model.EventHandler
	)

	localVarPath := "/event/{event}"
	localVarPath = strings.Replace(localVarPath, "{"+"event"+"}", fmt.Sprintf("%v", event), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ActiveOnly.IsSet() {
		localVarQueryParams.Add("activeOnly", parameterToString(localVarOptionals.ActiveOnly.Value(), ""))
	}

	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		err = a.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	} else {
		newErr := NewGenericSwaggerError(localVarBody, localVarHttpResponse.Status, nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
EventResourceApiService Get queue config by name
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param queueType
  - @param queueName

@return map[string]interface{}
*/
func (a *EventResourceApiService) GetQueueConfig(ctx context.Context, queueType string, queueName string) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue map[string]interface{}
	)

	localVarPath := "/event/queue/config/{queueType}/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"queueType"+"}", fmt.Sprintf("%v", queueType), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		err = a.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	} else {
		newErr := NewGenericSwaggerError(localVarBody, localVarHttpResponse.Status, nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
EventResourceApiService Get all queue configs
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return map[string]string
*/
func (a *EventResourceApiService) GetQueueNames(ctx context.Context) (map[string]string, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue map[string]string
	)

	localVarPath := "/event/queue/config"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		err = a.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	} else {
		newErr := NewGenericSwaggerError(localVarBody, localVarHttpResponse.Status, nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
EventResourceApiService Create or update queue config by name
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param queueType
  - @param queueName
*/
func (a *EventResourceApiService) PutQueueConfig(ctx context.Context, body string, queueType string, queueName string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/event/queue/config/{queueType}/{queueName}"
	localVarPath = strings.Replace(localVarPath, "{"+"queueType"+"}", fmt.Sprintf("%v", queueType), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"queueName"+"}", fmt.Sprintf("%v", queueName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarPostBody = &body
	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)
	if err != nil {
		return localVarHttpResponse, err
	}

	if !isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		newErr := NewGenericSwaggerError(localVarBody, localVarHttpResponse.Status, nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
EventResourceApiService Remove an event handler
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
*/
func (a *EventResourceApiService) RemoveEventHandler(ctx context.Context, name string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/event/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)
	if err != nil {
		return localVarHttpResponse, err
	}

	if !isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		newErr := NewGenericSwaggerError(localVarBody, localVarHttpResponse.Status, nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
EventResourceApiService Update an existing event handler.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *EventResourceApiService) UpdateEventHandler(ctx context.Context, body model.EventHandler) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/event"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarPostBody = &body
	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)
	if err != nil {
		return localVarHttpResponse, err
	}

	if !isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		newErr := NewGenericSwaggerError(localVarBody, localVarHttpResponse.Status, nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
