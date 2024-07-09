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
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
)

type ApplicationClient interface {
	AddRoleToApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error)
	CreateAccessKey(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error)
	CreateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest) (*rbac.ConductorApplication, *http.Response, error)
	DeleteAccessKey(ctx context.Context, applicationId string, keyId string) (*http.Response, error)
	DeleteApplication(ctx context.Context, id string) (interface{}, *http.Response, error)
	DeleteTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error)
	GetAccessKeys(ctx context.Context, id string) ([]rbac.AccessKeyResponse, *http.Response, error)
	GetAppByAccessKeyId(ctx context.Context, accessKeyId string) (interface{}, *http.Response, error)
	GetApplication(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error)
	GetTagsForApplication(ctx context.Context, id string) ([]model.Tag, *http.Response, error)
	ListApplications(ctx context.Context) ([]rbac.ConductorApplication, *http.Response, error)
	PutTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error)
	RemoveRoleFromApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error)
	ToggleAccessKeyStatus(ctx context.Context, applicationId string, keyId string) (interface{}, *http.Response, error)
	UpdateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest, id string) (*rbac.ConductorApplication, *http.Response, error)
}

func NewApplicationClient(apiClient *APIClient) *ApplicationResourceApiService {
	return &ApplicationResourceApiService{apiClient}
}
