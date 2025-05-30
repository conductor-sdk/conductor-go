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
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"

	"fmt"
)

type EnvironmentResourceApiService struct {
	*APIClient
}

/*
EnvironmentResourceApiService Create or update an environment variable (requires metadata or admin role)
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param key
*/
func (a *EnvironmentResourceApiService) CreateOrUpdateEnvVariable(ctx context.Context, body string, key string) (*http.Response, error) {
	path := fmt.Sprintf("/environment/%s", key)

	resp, err := a.PutWithContentType(ctx, path, body, "text/plain", nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
EnvironmentResourceApiService Delete an environment variable (requires metadata or admin role)
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param key
    @return string
*/
func (a *EnvironmentResourceApiService) DeleteEnvVariable(ctx context.Context, key string) (string, *http.Response, error) {
	var result string
	path := fmt.Sprintf("/environment/%s", key)

	resp, err := a.Delete(ctx, path, nil, &result)

	if err != nil {
		return "", resp, err
	}

	return result, resp, nil
}

/*
EnvironmentResourceApiService Delete a tag for environment variable name
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
*/
func (a *EnvironmentResourceApiService) DeleteTagForEnvVar(ctx context.Context, body []model.Tag, name string) (*http.Response, error) {
	path := fmt.Sprintf("/environment/%s/tags", name)
	resp, err := a.DeleteWithBody(ctx, path, body, nil)
	return resp, err
}

/*
EnvironmentResourceApiService Get the environment value by key
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param key
    @return string
*/
func (a *EnvironmentResourceApiService) Get(ctx context.Context, key string) (string, *http.Response, error) {
	var result string
	path := fmt.Sprintf("/environment/%s", key)

	resp, err := a.APIClient.Get(ctx, path, nil, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

/*
EnvironmentResourceApiService List all the environment variables
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []model.EnvironmentVariable
*/
func (a *EnvironmentResourceApiService) GetAll(ctx context.Context) ([]model.EnvironmentVariable, *http.Response, error) {
	var result []model.EnvironmentVariable
	resp, err := a.APIClient.Get(ctx, "/environment", nil, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
EnvironmentResourceApiService Get tags by environment variable name
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return []Tag
*/
func (a *EnvironmentResourceApiService) GetTagsForEnvVar(ctx context.Context, name string) ([]model.Tag, *http.Response, error) {
	var result []model.Tag
	path := fmt.Sprintf("/environment/%s/tags", name)
	resp, err := a.APIClient.Get(ctx, path, nil, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
EnvironmentResourceApiService Put a tag to environment variable name
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
*/
func (a *EnvironmentResourceApiService) PutTagForEnvVar(ctx context.Context, body []model.Tag, name string) (*http.Response, error) {
	path := fmt.Sprintf("/environment/%s/tags", name)
	resp, err := a.Put(ctx, path, body, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
