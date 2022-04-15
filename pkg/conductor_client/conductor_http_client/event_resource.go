package conductor_http_client

// import (
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"strings"

// 	"github.com/antihax/optional"
// )

// // Linger please
// var (
// 	_ context.Context
// )

// type EventResourceApiService service

// /*
// EventResourceApiService Add a new event handler.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body

// */
// func (a *EventResourceApiService) AddEventHandler(ctx context.Context, body EventHandler) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Post")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/event"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json"}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	// body params
// 	localVarPostBody = &body
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }

// /*
// EventResourceApiService Get all the event handlers
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
// @return []EventHandler
// */
// func (a *EventResourceApiService) GetEventHandlers(ctx context.Context) ([]EventHandler, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Get")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue []EventHandler
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/event"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{"*/*"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return localVarReturnValue, nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode < 300 {
// 		// If we succeed, return the data, otherwise pass on to decode error.
// 		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 		if err == nil {
// 			return localVarReturnValue, localVarHttpResponse, err
// 		}
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		if localVarHttpResponse.StatusCode == 200 {
// 			var v []EventHandler
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 			if err != nil {
// 				newErr.error = err.Error()
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 			}
// 			newErr.model = v
// 			return localVarReturnValue, localVarHttpResponse, newErr
// 		}
// 		return localVarReturnValue, localVarHttpResponse, newErr
// 	}

// 	return localVarReturnValue, localVarHttpResponse, nil
// }

// /*
// EventResourceApiService Get event handlers for a given event
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param event
//  * @param optional nil or *EventResourceApiGetEventHandlersForEventOpts - Optional Parameters:
//      * @param "ActiveOnly" (optional.Bool) -
// @return []EventHandler
// */

// type EventResourceApiGetEventHandlersForEventOpts struct {
// 	ActiveOnly optional.Bool
// }

// func (a *EventResourceApiService) GetEventHandlersForEvent(ctx context.Context, event string, localVarOptionals *EventResourceApiGetEventHandlersForEventOpts) ([]EventHandler, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Get")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue []EventHandler
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/event/{event}"
// 	localVarPath = strings.Replace(localVarPath, "{"+"event"+"}", fmt.Sprintf("%v", event), -1)

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	if localVarOptionals != nil && localVarOptionals.ActiveOnly.IsSet() {
// 		localVarQueryParams.Add("activeOnly", parameterToString(localVarOptionals.ActiveOnly.Value(), ""))
// 	}
// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{"*/*"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return localVarReturnValue, nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode < 300 {
// 		// If we succeed, return the data, otherwise pass on to decode error.
// 		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 		if err == nil {
// 			return localVarReturnValue, localVarHttpResponse, err
// 		}
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		if localVarHttpResponse.StatusCode == 200 {
// 			var v []EventHandler
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
// 			if err != nil {
// 				newErr.error = err.Error()
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 			}
// 			newErr.model = v
// 			return localVarReturnValue, localVarHttpResponse, newErr
// 		}
// 		return localVarReturnValue, localVarHttpResponse, newErr
// 	}

// 	return localVarReturnValue, localVarHttpResponse, nil
// }

// /*
// EventResourceApiService Remove an event handler
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param name

// */
// func (a *EventResourceApiService) RemoveEventHandlerStatus(ctx context.Context, name string) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Delete")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/event/{name}"
// 	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }

// /*
// EventResourceApiService Update an existing event handler.
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body

// */
// func (a *EventResourceApiService) UpdateEventHandler(ctx context.Context, body EventHandler) (*http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Put")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/event"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json"}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	// body params
// 	localVarPostBody = &body
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body:  localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}
// 		return localVarHttpResponse, newErr
// 	}

// 	return localVarHttpResponse, nil
// }
