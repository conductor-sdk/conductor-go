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
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"net/http"
)

type ApplicationResourceApiService struct {
	*APIClient
}

/*
ApplicationResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param applicationId
  - @param role
    @return interface{}
*/
func (a *ApplicationResourceApiService) AddRoleToApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/applications/%s/roles/%s", applicationId, role)
	resp, err := a.Post(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
ApplicationResourceApiService Create an access key for an application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *ApplicationResourceApiService) CreateAccessKey(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error) {
	var result rbac.ConductorApplication
	path := fmt.Sprintf("/applications/%s/accessKeys", id)
	resp, err := a.Post(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

/*
ApplicationResourceApiService Create an application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
    @return interface{}
*/
func (a *ApplicationResourceApiService) CreateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest) (*rbac.ConductorApplication, *http.Response, error) {
	var result rbac.ConductorApplication
	resp, err := a.Post(ctx, "/applications", body, &result)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

/*
ApplicationResourceApiService Delete an access key
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param applicationId
  - @param keyId
    @return interface{}
*/
func (a *ApplicationResourceApiService) DeleteAccessKey(ctx context.Context, applicationId string, keyId string) (*http.Response, error) {
	path := fmt.Sprintf("/applications/%s/accessKeys/%s", applicationId, keyId)
	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

/*
ApplicationResourceApiService Delete an application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *ApplicationResourceApiService) DeleteApplication(ctx context.Context, id string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/applications/%s", id)
	resp, err := a.Delete(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
ApplicationResourceApiService Delete a tag for application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param id
*/
func (a *ApplicationResourceApiService) DeleteTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error) {
	path := fmt.Sprintf("/applications/%s/tags", id)
	resp, err := a.DeleteWithBody(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetAccessKeys gets all access keys for an application
func (a *ApplicationResourceApiService) GetAccessKeys(ctx context.Context, id string) ([]rbac.AccessKeyResponse, *http.Response, error) {
	var result []rbac.AccessKeyResponse
	path := fmt.Sprintf("/applications/%s/accessKeys", id)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetAppByAccessKeyId gets an application by access key ID
func (a *ApplicationResourceApiService) GetAppByAccessKeyId(ctx context.Context, accessKeyId string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/applications/key/%s", accessKeyId)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetApplication gets an application by ID
func (a *ApplicationResourceApiService) GetApplication(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error) {
	var result rbac.ConductorApplication
	path := fmt.Sprintf("/applications/%s", id)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, err
}

// GetTagsForApplication gets all tags for an application
func (a *ApplicationResourceApiService) GetTagsForApplication(ctx context.Context, id string) ([]model.Tag, *http.Response, error) {
	var result []model.Tag
	path := fmt.Sprintf("/applications/%s/tags", id)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// ListApplications lists all applications
func (a *ApplicationResourceApiService) ListApplications(ctx context.Context) ([]rbac.ConductorApplication, *http.Response, error) {
	var result []rbac.ConductorApplication
	resp, err := a.Get(ctx, "/applications", nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// PutTagForApplication adds tags to an application
func (a *ApplicationResourceApiService) PutTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error) {
	path := fmt.Sprintf("/applications/%s/tags", id)
	resp, err := a.Put(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// RemoveRoleFromApplicationUser removes a role from an application user
func (a *ApplicationResourceApiService) RemoveRoleFromApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/applications/%s/roles/%s", applicationId, role)
	resp, err := a.Delete(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// ToggleAccessKeyStatus toggles the status of an access key
func (a *ApplicationResourceApiService) ToggleAccessKeyStatus(ctx context.Context, applicationId string, keyId string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/applications/%s/accessKeys/%s/status", applicationId, keyId)
	resp, err := a.Post(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// UpdateApplication updates an application
func (a *ApplicationResourceApiService) UpdateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest, id string) (*rbac.ConductorApplication, *http.Response, error) {
	var result rbac.ConductorApplication
	path := fmt.Sprintf("/applications/%s", id)
	resp, err := a.Put(ctx, path, body, &result)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
