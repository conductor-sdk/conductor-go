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
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type TagsApiService struct {
	*APIClient
}

/*
TagsApiService Adds the tag to the task
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param taskName

@return interface{}
*/
func (a *TagsApiService) AddTaskTag(ctx context.Context, body model.TagObject, taskName string) (interface{}, *http.Response, error) {
	var result interface{}

	path := fmt.Sprintf("/metadata/task/%s/tags", taskName)

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TagsApiService Adds the tag to the workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name

@return interface{}
*/
func (a *TagsApiService) AddWorkflowTag(ctx context.Context, body model.TagObject, name string) (interface{}, *http.Response, error) {
	var result interface{}

	path := fmt.Sprintf("/metadata/workflow/%s/tags", name)
	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TagsApiService Removes the tag of the task
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param taskName

@return interface{}
*/
func (a *TagsApiService) DeleteTaskTag(ctx context.Context, body model.TagString, taskName string) (interface{}, *http.Response, error) {
	var result interface{}

	path := fmt.Sprintf("/metadata/task/%s/tags", taskName)

	resp, err := a.DeleteWithBody(ctx, path, body, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TagsApiService Removes the tag of the workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name

@return interface{}
*/
func (a *TagsApiService) DeleteWorkflowTag(ctx context.Context, body model.TagObject, name string) (interface{}, *http.Response, error) {
	var result interface{}

	localVarPath := fmt.Sprintf("/metadata/workflow/%s/tags", name)
	resp, err := a.DeleteWithBody(ctx, localVarPath, body, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TagsApiService List all tags
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []model.TagObject
*/
func (a *TagsApiService) GetTags1(ctx context.Context) ([]model.TagObject, *http.Response, error) {
	var result []model.TagObject

	path := "/metadata/tags"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TagsApiService Returns all the tags of the task
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskName

@return []model.TagObject
*/
func (a *TagsApiService) GetTaskTags(ctx context.Context, taskName string) ([]model.TagObject, *http.Response, error) {
	var result []model.TagObject

	localVarPath := fmt.Sprintf("/metadata/task/%s/tags", taskName)
	resp, err := a.Get(ctx, localVarPath, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TagsApiService Returns all the tags of the workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name

@return []model.TagObject
*/
func (a *TagsApiService) GetWorkflowTags(ctx context.Context, name string) ([]model.TagObject, *http.Response, error) {
	var result []model.TagObject

	path := fmt.Sprintf("/metadata/workflow/%s/tags", name)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TagsApiService Adds the tag to the task
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param taskName

@return interface{}
*/
func (a *TagsApiService) SetTaskTags(ctx context.Context, body []model.TagObject, taskName string) (interface{}, *http.Response, error) {
	var result interface{}

	localVarPath := fmt.Sprintf("/metadata/task/%s/tags", taskName)
	resp, err := a.Put(ctx, localVarPath, body, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TagsApiService Set the tags of the workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name

@return interface{}
*/
func (a *TagsApiService) SetWorkflowTags(ctx context.Context, body []model.TagObject, name string) (interface{}, *http.Response, error) {
	var result interface{}

	localVarPath := fmt.Sprintf("/metadata/workflow/%s/tags", name)
	resp, err := a.Put(ctx, localVarPath, body, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
