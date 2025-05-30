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
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
	"net/url"
	"strconv"
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
	path := "/metadata/workflow"

	queryParams := url.Values{
		"overwrite": []string{strconv.FormatBool(overwrite)},
	}
	resp, err := a.PostWithParams(ctx, path, queryParams, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
MetadataResourceApiService Create a new workflow definition with tags
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) RegisterWorkflowDefWithTags(ctx context.Context, overwrite bool, body model.WorkflowDef, tags []model.MetadataTag) (*http.Response, error) {
	path := "/metadata/workflow"

	params := url.Values{
		"overwrite": []string{strconv.FormatBool(overwrite)},
	}
	tagObjects := []model.TagObject{}
	for i := 0; i < len(tags); i++ {
		tagObjects = append(tagObjects, model.NewTagObject(tags[i]))
	}
	workflowDefWithTags := body
	workflowDefWithTags.Tags = tagObjects
	workflowDefWithTags.OverwriteTags = true

	resp, err := a.PostWithParams(ctx, path, params, workflowDefWithTags, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
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
	var result model.WorkflowDef

	path := fmt.Sprintf("/metadata/workflow/%s", name)

	queryParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Version.IsSet() {
		queryParams.Add("version", parameterToString(localVarOptionals.Version.Value(), ""))
	}

	resp, err := a.APIClient.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.WorkflowDef{}, resp, err
	}
	return result, resp, nil
}

/*
MetadataResourceApiService Retrieves all workflow definition along with blueprint
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []http_model.WorkflowDef
*/
func (a *MetadataResourceApiService) GetAll(ctx context.Context) ([]model.WorkflowDef, *http.Response, error) {
	var result []model.WorkflowDef

	path := "/metadata/workflow"

	queryParams := url.Values{}

	resp, err := a.APIClient.Get(ctx, path, queryParams, &result)
	if err != nil {
		return []model.WorkflowDef{}, resp, err
	}
	return result, resp, nil
}

/*
MetadataResourceApiService Gets the task definition
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param tasktype

@return http_model.TaskDef
*/
func (a *MetadataResourceApiService) GetTaskDef(ctx context.Context, tasktype string) (model.TaskDef, *http.Response, error) {
	var result model.TaskDef
	path := fmt.Sprintf("/metadata/taskdefs/%s", tasktype)

	resp, err := a.APIClient.Get(ctx, path, nil, &result)
	if err != nil {
		return model.TaskDef{}, resp, err
	}

	return result, resp, nil
}

/*
MetadataResourceApiService Gets all task definition
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []http_model.TaskDef
*/
func (a *MetadataResourceApiService) GetTaskDefs(ctx context.Context) ([]model.TaskDef, *http.Response, error) {
	var result []model.TaskDef

	path := "/metadata/taskdefs"

	resp, err := a.APIClient.Get(ctx, path, nil, &result)
	if err != nil {
		return []model.TaskDef{}, resp, err
	}
	return result, resp, nil
}

/*
MetadataResourceApiService Update an existing task
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) UpdateTaskDef(ctx context.Context, body model.TaskDef) (*http.Response, error) {
	path := "/metadata/taskdefs"

	resp, err := a.APIClient.Put(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
MetadataResourceApiService Update an existing task along with tags
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) UpdateTaskDefWithTags(ctx context.Context, body model.TaskDef, tags []model.MetadataTag, overwriteTags bool) (*http.Response, error) {
	path := "/metadata/taskdefs"

	tagObjects := []model.TagObject{}
	for i := 0; i < len(tags); i++ {
		tagObjects = append(tagObjects, model.NewTagObject(tags[i]))
	}

	taskDefWithTags := body
	taskDefWithTags.Tags = tagObjects
	taskDefWithTags.OverwriteTags = overwriteTags

	resp, err := a.APIClient.Put(ctx, path, taskDefWithTags, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
MetadataResourceApiService Create new task definition(s)
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) RegisterTaskDef(ctx context.Context, body []model.TaskDef) (*http.Response, error) {
	path := "/metadata/taskdefs"

	resp, err := a.APIClient.Post(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
MetadataResourceApiService Create new task definition with tags
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body model.TaskDef
  - @param tags []model.MetadataTag
*/
func (a *MetadataResourceApiService) RegisterTaskDefWithTags(ctx context.Context, body model.TaskDef, tags []model.MetadataTag) (*http.Response, error) {
	path := "/metadata/taskdefs"

	tagObjects := []model.TagObject{}
	for i := 0; i < len(tags); i++ {
		tagObjects = append(tagObjects, model.NewTagObject(tags[i]))
	}

	taskDefWithTags := body
	taskDefWithTags.Tags = tagObjects
	taskDefWithTags.OverwriteTags = true
	taskDefs := []model.TaskDef{taskDefWithTags}

	resp, err := a.APIClient.Post(ctx, path, taskDefs, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
MetadataResourceApiService Remove a task definition
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param tasktype
*/
func (a *MetadataResourceApiService) UnregisterTaskDef(ctx context.Context, taskType string) (*http.Response, error) {
	path := fmt.Sprintf("/metadata/taskdefs/%s", taskType)

	resp, err := a.APIClient.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
MetadataResourceApiService Removes workflow definition. It does not remove workflows associated with the definition.
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
  - @param version
*/
func (a *MetadataResourceApiService) UnregisterWorkflowDef(ctx context.Context, name string, version int32) (*http.Response, error) {
	path := fmt.Sprintf("/metadata/workflow/%s/%d", name, version)

	resp, err := a.APIClient.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
MetadataResourceApiService Create or update workflow definition
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) Update(ctx context.Context, body []model.WorkflowDef) (*http.Response, error) {
	path := "/metadata/workflow"

	resp, err := a.APIClient.Put(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
MetadataResourceApiService Create or update workflow definition along with tags
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *MetadataResourceApiService) UpdateWorkflowDefWithTags(ctx context.Context, body model.WorkflowDef, tags []model.MetadataTag, overwriteTags bool) (*http.Response, error) {
	path := "/metadata/workflow"

	tagObjects := []model.TagObject{}
	for i := 0; i < len(tags); i++ {
		tagObjects = append(tagObjects, model.NewTagObject(tags[i]))
	}
	workflowDefWithTags := body
	workflowDefWithTags.Tags = tagObjects
	workflowDefWithTags.OverwriteTags = overwriteTags
	workflowDefs := []model.WorkflowDef{workflowDefWithTags}

	resp, err := a.APIClient.Put(ctx, path, workflowDefs, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (a *MetadataResourceApiService) GetTagsForWorkflowDef(ctx context.Context, name string) ([]model.MetadataTag, error) {
	path := fmt.Sprintf("/metadata/workflow/%s?metadata=true", name)

	var workflowDef model.WorkflowDef
	_, err := a.APIClient.Get(ctx, path, nil, &workflowDef)
	if err != nil {
		return nil, err
	}

	// Extract and convert tags as in your original implementation
	var result []model.MetadataTag
	for _, tag := range workflowDef.Tags {
		value := fmt.Sprintf("%v", tag.Value)
		result = append(result, model.MetadataTag{
			Key:   tag.Key,
			Value: value,
		})
	}

	return result, nil
}

func (a *MetadataResourceApiService) GetTagsForTaskDef(ctx context.Context, tasktype string) ([]model.MetadataTag, error) {
	path := fmt.Sprintf("/metadata/taskdefs/%s?metadata=true", tasktype)

	var taskDef model.WorkflowDef
	_, err := a.APIClient.Get(ctx, path, nil, &taskDef)
	if err != nil {
		return nil, err
	}

	// Extract and convert tags as in your original implementation
	var result []model.MetadataTag
	for _, tag := range taskDef.Tags {
		value := fmt.Sprintf("%v", tag.Value)
		result = append(result, model.MetadataTag{
			Key:   tag.Key,
			Value: value,
		})
	}

	return result, nil
}
