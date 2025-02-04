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
	"strconv"
	"strings"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

// Linger please
var (
	_ context.Context
)

type MetadataResourceApiService struct {
	*APIClient
}

/*
MetadataResourceApiService Create a new workflow definition
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) RegisterWorkflowDef(ctx context.Context, overwrite bool, body model.WorkflowDef) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/metadata/workflow"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{
		"overwrite": []string{strconv.FormatBool(overwrite)},
	}
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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
MetadataResourceApiService Create a new workflow definition with tags
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) RegisterWorkflowDefWithTags(ctx context.Context, overwrite bool, body model.WorkflowDef, tags []model.MetadataTag) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/metadata/workflow"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{
		"overwrite": []string{strconv.FormatBool(overwrite)},
	}
	localVarFormParams := url.Values{}

	tagObjects := []model.TagObject{}
	for i := 0; i < len(tags); i++ {
		tagObjects = append(tagObjects, model.NewTagObject(tags[i]))
	}

	workflowDefWithTags := body
	workflowDefWithTags.Tags = tagObjects
	workflowDefWithTags.OverwriteTags = true
	localVarPostBody = &workflowDefWithTags

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
MetadataResourceApiService Retrieves workflow definition along with blueprint
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name
 * @param optional nil or *MetadataResourceApiGetOpts - Optional Parameters:
     * @param "Version" (optional.Int32) -
@return http_model.WorkflowDef
*/

type MetadataResourceApiGetOpts struct {
	Version optional.Int32
}

func (a *MetadataResourceApiService) Get(ctx context.Context, name string, localVarOptionals *MetadataResourceApiGetOpts) (model.WorkflowDef, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.WorkflowDef
	)

	localVarPath := "/metadata/workflow/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Version.IsSet() {
		localVarQueryParams.Add("version", parameterToString(localVarOptionals.Version.Value(), ""))
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
MetadataResourceApiService Retrieves all workflow definition along with blueprint
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []http_model.WorkflowDef
*/
func (a *MetadataResourceApiService) GetAll(ctx context.Context) ([]model.WorkflowDef, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []model.WorkflowDef
	)

	localVarPath := "/metadata/workflow"

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
MetadataResourceApiService Gets the task definition
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param tasktype

@return http_model.TaskDef
*/
func (a *MetadataResourceApiService) GetTaskDef(ctx context.Context, tasktype string) (model.TaskDef, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue model.TaskDef
	)

	localVarPath := "/metadata/taskdefs/{tasktype}"
	localVarPath = strings.Replace(localVarPath, "{"+"tasktype"+"}", fmt.Sprintf("%v", tasktype), -1)

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
MetadataResourceApiService Gets all task definition
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []http_model.TaskDef
*/
func (a *MetadataResourceApiService) GetTaskDefs(ctx context.Context) ([]model.TaskDef, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []model.TaskDef
	)

	localVarPath := "/metadata/taskdefs"

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, err
}

/*
MetadataResourceApiService Update an existing task
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) UpdateTaskDef(ctx context.Context, body model.TaskDef) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/metadata/taskdefs"

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
MetadataResourceApiService Update an existing task along with tags
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) UpdateTaskDefWithTags(ctx context.Context, body model.TaskDef, tags []model.MetadataTag, overwriteTags bool) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPutBody    interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/metadata/taskdefs"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	tagObjects := []model.TagObject{}
	for i := 0; i < len(tags); i++ {
		tagObjects = append(tagObjects, model.NewTagObject(tags[i]))
	}

	taskDefWithTags := body
	taskDefWithTags.Tags = tagObjects
	taskDefWithTags.OverwriteTags = overwriteTags
	localVarPutBody = &taskDefWithTags

	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPutBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
MetadataResourceApiService Create new task definition(s)
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) RegisterTaskDef(ctx context.Context, body []model.TaskDef) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/metadata/taskdefs"

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
MetadataResourceApiService Create new task definition with tags
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body model.TaskDef
  - @param tags []model.MetadataTag
*/
func (a *MetadataResourceApiService) RegisterTaskDefWithTags(ctx context.Context, body model.TaskDef, tags []model.MetadataTag) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/metadata/taskdefs"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	tagObjects := []model.TagObject{}
	for i := 0; i < len(tags); i++ {
		tagObjects = append(tagObjects, model.NewTagObject(tags[i]))
	}

	taskDefWithTags := body
	taskDefWithTags.Tags = tagObjects
	taskDefWithTags.OverwriteTags = true
	taskDefs := []model.TaskDef{taskDefWithTags}
	localVarPostBody = &taskDefs
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
MetadataResourceApiService Remove a task definition
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param tasktype
*/
func (a *MetadataResourceApiService) UnregisterTaskDef(ctx context.Context, tasktype string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/metadata/taskdefs/{tasktype}"
	localVarPath = strings.Replace(localVarPath, "{"+"tasktype"+"}", fmt.Sprintf("%v", tasktype), -1)

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
MetadataResourceApiService Removes workflow definition. It does not remove workflows associated with the definition.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
  - @param version
*/
func (a *MetadataResourceApiService) UnregisterWorkflowDef(ctx context.Context, name string, version int32) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/metadata/workflow/{name}/{version}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

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
MetadataResourceApiService Create or update workflow definition
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) Update(ctx context.Context, body []model.WorkflowDef) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/metadata/workflow"

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
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
MetadataResourceApiService Create or update workflow definition along with tags
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) UpdateWorkflowDefWithTags(ctx context.Context, body model.WorkflowDef, tags []model.MetadataTag, overwriteTags bool) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	localVarPath := "/metadata/workflow"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Content-Type"] = "application/json"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	tagObjects := []model.TagObject{}
	for i := 0; i < len(tags); i++ {
		tagObjects = append(tagObjects, model.NewTagObject(tags[i]))
	}

	workflowDefWithTags := body
	workflowDefWithTags.Tags = tagObjects
	workflowDefWithTags.OverwriteTags = overwriteTags
	workflowDefs := []model.WorkflowDef{workflowDefWithTags}
	localVarPostBody = &workflowDefs

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

func (a *MetadataResourceApiService) GetTagsForWorkflowDef(ctx context.Context, name string) ([]model.MetadataTag, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []model.MetadataTag
	)

	localVarPath := "/metadata/workflow/{name}?metadata=true"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)

	if err != nil {
		return localVarReturnValue, err
	}

	if isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		extendedWorkflowDef := model.WorkflowDef{}
		err = a.decode(&extendedWorkflowDef, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			for i := 0; i < len(extendedWorkflowDef.Tags); i++ {
				value := fmt.Sprintf("%v", extendedWorkflowDef.Tags[i].Value)
				tag := model.MetadataTag{
					Key:   extendedWorkflowDef.Tags[i].Key,
					Value: value,
				}
				localVarReturnValue = append(localVarReturnValue, tag)
			}
		}
	} else {
		newErr := NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, err
}

func (a *MetadataResourceApiService) GetTagsForTaskDef(ctx context.Context, tasktype string) ([]model.MetadataTag, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []model.MetadataTag
	)

	localVarPath := "/metadata/taskdefs/{tasktype}?metadata=true"
	localVarPath = strings.Replace(localVarPath, "{"+"tasktype"+"}", fmt.Sprintf("%v", tasktype), -1)
	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "*/*"

	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	r, err := a.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := getDecompressedBody(localVarHttpResponse)
	if err != nil {
		return localVarReturnValue, err
	}

	if isSuccessfulStatus(localVarHttpResponse.StatusCode) {
		extendedTaskDef := model.WorkflowDef{}
		err = a.decode(&extendedTaskDef, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			for i := 0; i < len(extendedTaskDef.Tags); i++ {
				value := fmt.Sprintf("%v", extendedTaskDef.Tags[i].Value)
				tag := model.MetadataTag{
					Key:   extendedTaskDef.Tags[i].Key,
					Value: value,
				}
				localVarReturnValue = append(localVarReturnValue, tag)
			}
		}
	} else {
		return localVarReturnValue, NewGenericSwaggerError(localVarBody, string(localVarBody), nil, localVarHttpResponse.StatusCode)
	}

	return localVarReturnValue, err
}
