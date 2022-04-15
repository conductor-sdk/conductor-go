package conductor_http_client

// import (
// 	"context"
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

// type WorkflowBulkResourceApiService service

// /*
// WorkflowBulkResourceApiService Pause the list of workflows
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body
// @return BulkResponse
// */
// func (a *WorkflowBulkResourceApiService) PauseWorkflow1(ctx context.Context, body []string) (BulkResponse, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Put")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue BulkResponse
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/workflow/bulk/pause"

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
// 	localVarHttpHeaderAccepts := []string{"*/*"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	// body params
// 	localVarPostBody = &body
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
// 			var v BulkResponse
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
// WorkflowBulkResourceApiService Restart the list of completed workflow
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body
//  * @param optional nil or *WorkflowBulkResourceApiRestart1Opts - Optional Parameters:
//      * @param "UseLatestDefinitions" (optional.Bool) -
// @return BulkResponse
// */

// type WorkflowBulkResourceApiRestart1Opts struct {
// 	UseLatestDefinitions optional.Bool
// }

// func (a *WorkflowBulkResourceApiService) Restart1(ctx context.Context, body []string, localVarOptionals *WorkflowBulkResourceApiRestart1Opts) (BulkResponse, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Post")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue BulkResponse
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/workflow/bulk/restart"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	if localVarOptionals != nil && localVarOptionals.UseLatestDefinitions.IsSet() {
// 		localVarQueryParams.Add("useLatestDefinitions", parameterToString(localVarOptionals.UseLatestDefinitions.Value(), ""))
// 	}
// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json"}

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
// 	// body params
// 	localVarPostBody = &body
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
// 			var v BulkResponse
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
// WorkflowBulkResourceApiService Resume the list of workflows
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body
// @return BulkResponse
// */
// func (a *WorkflowBulkResourceApiService) ResumeWorkflow1(ctx context.Context, body []string) (BulkResponse, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Put")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue BulkResponse
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/workflow/bulk/resume"

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
// 	localVarHttpHeaderAccepts := []string{"*/*"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	// body params
// 	localVarPostBody = &body
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
// 			var v BulkResponse
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
// WorkflowBulkResourceApiService Retry the last failed task for each workflow from the list
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body
// @return BulkResponse
// */
// func (a *WorkflowBulkResourceApiService) Retry1(ctx context.Context, body []string) (BulkResponse, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Post")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue BulkResponse
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/workflow/bulk/retry"

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
// 	localVarHttpHeaderAccepts := []string{"*/*"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	// body params
// 	localVarPostBody = &body
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
// 			var v BulkResponse
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
// WorkflowBulkResourceApiService Terminate workflows execution
//  * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
//  * @param body
//  * @param optional nil or *WorkflowBulkResourceApiTerminateOpts - Optional Parameters:
//      * @param "Reason" (optional.String) -
// @return BulkResponse
// */

// type WorkflowBulkResourceApiTerminateOpts struct {
// 	Reason optional.String
// }

// func (a *WorkflowBulkResourceApiService) Terminate(ctx context.Context, body []string, localVarOptionals *WorkflowBulkResourceApiTerminateOpts) (BulkResponse, *http.Response, error) {
// 	var (
// 		localVarHttpMethod  = strings.ToUpper("Post")
// 		localVarPostBody    interface{}
// 		localVarFileName    string
// 		localVarFileBytes   []byte
// 		localVarReturnValue BulkResponse
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/api/workflow/bulk/terminate"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	if localVarOptionals != nil && localVarOptionals.Reason.IsSet() {
// 		localVarQueryParams.Add("reason", parameterToString(localVarOptionals.Reason.Value(), ""))
// 	}
// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json"}

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
// 	// body params
// 	localVarPostBody = &body
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
// 			var v BulkResponse
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
