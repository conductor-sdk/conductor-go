//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

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

// Linger please
var (
	_ context.Context
)

type WorkflowResourceApiService struct {
	*APIClient
}

/*
WorkflowResourceApiService Starts the decision task for a workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workflowId
*/
func (a *WorkflowResourceApiService) Decide(ctx context.Context, workflowId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/workflow/decide/{workflowId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
WorkflowResourceApiService Removes the workflow from the system
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowId
 * @param optional nil or *WorkflowResourceApiDeleteOpts - Optional Parameters:
     * @param "ArchiveWorkflow" (optional.Bool) -

*/

type WorkflowResourceApiDeleteOpts struct {
	ArchiveWorkflow optional.Bool
}

func (a *WorkflowResourceApiService) Delete(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiDeleteOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/workflow/{workflowId}/remove"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ArchiveWorkflow.IsSet() {
		localVarQueryParams.Add("archiveWorkflow", parameterToString(localVarOptionals.ArchiveWorkflow.Value(), ""))
	}

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
WorkflowResourceApiService Gets the workflow by workflow id
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowId
 * @param optional nil or *WorkflowResourceApiGetExecutionStatusOpts - Optional Parameters:
     * @param "IncludeTasks" (optional.Bool) -
@return http_model.Workflow
*/

type WorkflowResourceApiGetExecutionStatusOpts struct {
	IncludeTasks optional.Bool
}

func (a *WorkflowResourceApiService) GetExecutionStatus(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiGetExecutionStatusOpts) (model.Workflow, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.Workflow
	)

	localVarPath := "/workflow/{workflowId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IncludeTasks.IsSet() {
		localVarQueryParams.Add("includeTasks", parameterToString(localVarOptionals.IncludeTasks.Value(), ""))
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

func (a *WorkflowResourceApiService) GetWorkflowState(ctx context.Context, workflowId string, includeOutput bool, includeVariables bool) (model.WorkflowState, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.WorkflowState
	)

	localVarPath := "/workflow/{workflowId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("includeOutput", parameterToString(includeOutput, ""))
	localVarQueryParams.Add("includeVariables", parameterToString(includeVariables, ""))

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Get the uri and path of the external storage where the workflow payload is to be stored
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param path
  - @param operation
  - @param payloadType

@return http_model.ExternalStorageLocation
*/
func (a *WorkflowResourceApiService) GetExternalStorageLocation(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.ExternalStorageLocation
	)

	localVarPath := "/workflow/externalstoragelocation"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("path", parameterToString(path, ""))
	localVarQueryParams.Add("operation", parameterToString(operation, ""))
	localVarQueryParams.Add("payloadType", parameterToString(payloadType, ""))

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Retrieve all the running workflows
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name
 * @param optional nil or *WorkflowResourceApiGetRunningWorkflowOpts - Optional Parameters:
     * @param "Version" (optional.Int32) -
     * @param "StartTime" (optional.Int64) -
     * @param "EndTime" (optional.Int64) -
@return []string
*/

type WorkflowResourceApiGetRunningWorkflowOpts struct {
	Version   optional.Int32
	StartTime optional.Int64
	EndTime   optional.Int64
}

func (a *WorkflowResourceApiService) GetRunningWorkflow(ctx context.Context, name string, localVarOptionals *WorkflowResourceApiGetRunningWorkflowOpts) ([]string, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []string
	)

	localVarPath := "/workflow/running/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Version.IsSet() {
		localVarQueryParams.Add("version", parameterToString(localVarOptionals.Version.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartTime.IsSet() {
		localVarQueryParams.Add("startTime", parameterToString(localVarOptionals.StartTime.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EndTime.IsSet() {
		localVarQueryParams.Add("endTime", parameterToString(localVarOptionals.EndTime.Value(), ""))
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Lists workflows for the given correlation id list
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param name
 * @param optional nil or *WorkflowResourceApiGetWorkflowsOpts - Optional Parameters:
     * @param "IncludeClosed" (optional.Bool) -
     * @param "IncludeTasks" (optional.Bool) -
@return map[string][]http_model.Workflow
*/

func (a *WorkflowResourceApiService) GetWorkflows(ctx context.Context, body []string, name string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue map[string][]model.Workflow
	)

	localVarPath := "/workflow/{name}/correlated"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IncludeClosed.IsSet() {
		localVarQueryParams.Add("includeClosed", parameterToString(localVarOptionals.IncludeClosed.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeTasks.IsSet() {
		localVarQueryParams.Add("includeTasks", parameterToString(localVarOptionals.IncludeTasks.Value(), ""))
	}

	localVarPostBody = &body
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

func (a *WorkflowResourceApiService) GetWorkflowsBatch(ctx context.Context, body map[string][]string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue map[string][]model.Workflow
	)

	localVarPath := "/workflow/correlated/batch"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IncludeClosed.IsSet() {
		localVarQueryParams.Add("includeClosed", parameterToString(localVarOptionals.IncludeClosed.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeTasks.IsSet() {
		localVarQueryParams.Add("includeTasks", parameterToString(localVarOptionals.IncludeTasks.Value(), ""))
	}

	localVarPostBody = &body
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Lists workflows for the given correlation id
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name
 * @param correlationId
 * @param optional nil or *WorkflowResourceApiGetWorkflowsOpts - Optional Parameters:
     * @param "IncludeClosed" (optional.Bool) -
     * @param "IncludeTasks" (optional.Bool) -
@return []http_model.Workflow
*/

type WorkflowResourceApiGetWorkflowsOpts struct {
	IncludeClosed optional.Bool
	IncludeTasks  optional.Bool
}

func (a *WorkflowResourceApiService) GetWorkflowsByCorrelationId(ctx context.Context, name string, correlationId string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) ([]model.Workflow, *http.Response, error) {
	return a.GetWorkflows1(ctx, name, correlationId, localVarOptionals)
}
func (a *WorkflowResourceApiService) GetWorkflows1(ctx context.Context, name string, correlationId string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) ([]model.Workflow, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []model.Workflow
	)

	localVarPath := "/workflow/{name}/correlated/{correlationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"correlationId"+"}", fmt.Sprintf("%v", correlationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IncludeClosed.IsSet() {
		localVarQueryParams.Add("includeClosed", parameterToString(localVarOptionals.IncludeClosed.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeTasks.IsSet() {
		localVarQueryParams.Add("includeTasks", parameterToString(localVarOptionals.IncludeTasks.Value(), ""))
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Pauses the workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workflowId
*/
func (a *WorkflowResourceApiService) PauseWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/workflow/{workflowId}/pause"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
WorkflowResourceApiService Reruns the workflow from a specific task
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param workflowId

@return string
*/
func (a *WorkflowResourceApiService) Rerun(ctx context.Context, body model.RerunWorkflowRequest, workflowId string) (string, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue string
	)

	localVarPath := "/workflow/{workflowId}/rerun"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "text/plain"
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarPostBody = &body
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Resets callback times of all non-terminal SIMPLE tasks to 0
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workflowId
*/
func (a *WorkflowResourceApiService) ResetWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/workflow/{workflowId}/resetcallbacks"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
WorkflowResourceApiService Restarts a completed workflow
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowId
 * @param optional nil or *WorkflowResourceApiRestartOpts - Optional Parameters:
     * @param "UseLatestDefinitions" (optional.Bool) -

*/

type WorkflowResourceApiRestartOpts struct {
	UseLatestDefinitions optional.Bool
}

func (a *WorkflowResourceApiService) Restart(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiRestartOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/workflow/{workflowId}/restart"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.UseLatestDefinitions.IsSet() {
		localVarQueryParams.Add("useLatestDefinitions", parameterToString(localVarOptionals.UseLatestDefinitions.Value(), ""))
	}

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
WorkflowResourceApiService Resumes the workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workflowId
*/
func (a *WorkflowResourceApiService) ResumeWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/workflow/{workflowId}/resume"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
WorkflowResourceApiService Retries the last failed task
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowId
 * @param optional nil or *WorkflowResourceApiRetryOpts - Optional Parameters:
     * @param "ResumeSubworkflowTasks" (optional.Bool) -

*/

type WorkflowResourceApiRetryOpts struct {
	ResumeSubworkflowTasks optional.Bool
}

func (a *WorkflowResourceApiService) Retry(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiRetryOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/workflow/{workflowId}/retry"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ResumeSubworkflowTasks.IsSet() {
		localVarQueryParams.Add("resumeSubworkflowTasks", parameterToString(localVarOptionals.ResumeSubworkflowTasks.Value(), ""))
	}

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
WorkflowResourceApiService Search for workflows based on payload and other parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *WorkflowResourceApiSearchOpts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return http_model.SearchResultWorkflowSummary
*/

type WorkflowResourceApiSearchOpts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *WorkflowResourceApiService) Search(ctx context.Context, localVarOptionals *WorkflowResourceApiSearchOpts) (model.SearchResultWorkflowSummary, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.SearchResultWorkflowSummary
	)

	localVarPath := "/workflow/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FreeText.IsSet() {
		localVarQueryParams.Add("freeText", parameterToString(localVarOptionals.FreeText.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Query.IsSet() {
		localVarQueryParams.Add("query", parameterToString(localVarOptionals.Query.Value(), ""))
	}
	localVarHeaderParams["Accept"] = "*/*"
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Search for workflows based on payload and other parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *WorkflowResourceApiSearchV2Opts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return http_model.SearchResultWorkflow
*/

type WorkflowResourceApiSearchV2Opts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *WorkflowResourceApiService) SearchV2(ctx context.Context, localVarOptionals *WorkflowResourceApiSearchV2Opts) (model.SearchResultWorkflow, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.SearchResultWorkflow
	)

	localVarPath := "/workflow/search-v2"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FreeText.IsSet() {
		localVarQueryParams.Add("freeText", parameterToString(localVarOptionals.FreeText.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Query.IsSet() {
		localVarQueryParams.Add("query", parameterToString(localVarOptionals.Query.Value(), ""))
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Search for workflows based on task parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *WorkflowResourceApiSearchWorkflowsByTasksOpts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return http_model.SearchResultWorkflowSummary
*/

type WorkflowResourceApiSearchWorkflowsByTasksOpts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *WorkflowResourceApiService) SearchWorkflowsByTasks(ctx context.Context, localVarOptionals *WorkflowResourceApiSearchWorkflowsByTasksOpts) (model.SearchResultWorkflowSummary, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.SearchResultWorkflowSummary
	)

	localVarPath := "/workflow/search-by-tasks"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FreeText.IsSet() {
		localVarQueryParams.Add("freeText", parameterToString(localVarOptionals.FreeText.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Query.IsSet() {
		localVarQueryParams.Add("query", parameterToString(localVarOptionals.Query.Value(), ""))
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Search for workflows based on task parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *WorkflowResourceApiSearchWorkflowsByTasksV2Opts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return http_model.SearchResultWorkflow
*/

type WorkflowResourceApiSearchWorkflowsByTasksV2Opts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *WorkflowResourceApiService) SearchWorkflowsByTasksV2(ctx context.Context, localVarOptionals *WorkflowResourceApiSearchWorkflowsByTasksV2Opts) (model.SearchResultWorkflow, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.SearchResultWorkflow
	)

	localVarPath := "/workflow/search-by-tasks-v2"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		localVarQueryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FreeText.IsSet() {
		localVarQueryParams.Add("freeText", parameterToString(localVarOptionals.FreeText.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Query.IsSet() {
		localVarQueryParams.Add("query", parameterToString(localVarOptionals.Query.Value(), ""))
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Skips a given task from a current running workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workflowId
  - @param taskReferenceName
  - @param skipTaskRequest
*/
func (a *WorkflowResourceApiService) SkipTaskFromWorkflow(ctx context.Context, workflowId string, taskReferenceName string, skipTaskRequest model.SkipTaskRequest) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/workflow/{workflowId}/skiptask/{taskReferenceName}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"taskReferenceName"+"}", fmt.Sprintf("%v", taskReferenceName), -1)

	localVarHeaderParams := make(map[string]string)

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("skipTaskRequest", parameterToString(skipTaskRequest, ""))

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
WorkflowResourceApiService Start a new workflow. Returns the ID of the workflow instance that can be later used for tracking
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param name
 * @param optional nil or *WorkflowResourceApiStartWorkflowOpts - Optional Parameters:
     * @param "Version" (optional.Int32) -
     * @param "CorrelationId" (optional.String) -
     * @param "Priority" (optional.Int32) -
@return string
*/

type WorkflowResourceApiStartWorkflowOpts struct {
	Version       optional.Int32
	CorrelationId optional.String
	Priority      optional.Int32
}

func (a *WorkflowResourceApiService) StartWorkflow(ctx context.Context, body map[string]interface{}, name string, localVarOptionals *WorkflowResourceApiStartWorkflowOpts) (string, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue string
	)

	localVarPath := "/workflow/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "text/plain"
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Version.IsSet() {
		localVarQueryParams.Add("version", parameterToString(localVarOptionals.Version.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CorrelationId.IsSet() {
		localVarQueryParams.Add("correlationId", parameterToString(localVarOptionals.CorrelationId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Priority.IsSet() {
		localVarQueryParams.Add("priority", parameterToString(localVarOptionals.Priority.Value(), ""))
	}

	localVarPostBody = &body
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

func (a *WorkflowResourceApiService) ExecuteWorkflow(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask string) (model.WorkflowRun, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.WorkflowRun
	)

	localVarPath := "/workflow/execute/{name}/{version}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "application/json"
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("requestId", parameterToString(requestId, ""))
	if len(waitUntilTask) > 0 {
		localVarQueryParams.Add("waitUntilTaskRef", parameterToString(waitUntilTask, ""))
	}

	localVarPostBody = &body
	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)

	localVarHttpResponse.Body.Close()
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
WorkflowResourceApiService Start a new workflow with http_model.StartWorkflowRequest, which allows task to be executed in a domain
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body

@return string
*/
func (a *WorkflowResourceApiService) StartWorkflowWithRequest(ctx context.Context, body model.StartWorkflowRequest) (string, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue string
	)

	localVarPath := "/workflow"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "text/plain"
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarPostBody = &body
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
WorkflowResourceApiService Terminate workflow execution
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowId
 * @param optional nil or *WorkflowResourceApiTerminateOpts - Optional Parameters:
     * @param "Reason" (optional.String) -

*/

type WorkflowResourceApiTerminateOpts struct {
	Reason                 optional.String
	TriggerFailureWorkflow optional.Bool
}

func (a *WorkflowResourceApiService) Terminate(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiTerminateOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/workflow/{workflowId}"
	localVarPath = strings.Replace(localVarPath, "{"+"workflowId"+"}", fmt.Sprintf("%v", workflowId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Reason.IsSet() {
		localVarQueryParams.Add("reason", parameterToString(localVarOptionals.Reason.Value(), ""))
	}

	if localVarOptionals != nil && localVarOptionals.TriggerFailureWorkflow.IsSet() {
		localVarQueryParams.Add("triggerFailureWorkflow", parameterToString(localVarOptionals.TriggerFailureWorkflow.Value(), ""))
	}

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

// AddWorkflowTags adds tags to a workflow
func (a *WorkflowResourceApiService) AddWorkflowTags(ctx context.Context, name string, tags []model.TagObject) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// Create path and map variables
	localVarPath := "/workflow/{name}/tags"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "application/json"
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// Set body parameter
	localVarPostBody = &tags

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
