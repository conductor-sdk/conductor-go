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
)

type EnvironmentClient interface {
	CreateOrUpdateEnvVariable(ctx context.Context, body string, key string) (*http.Response, error)
	DeleteEnvVariable(ctx context.Context, key string) (string, *http.Response, error)
	DeleteTagForEnvVar(ctx context.Context, body []model.Tag, name string) (*http.Response, error)
	Get(ctx context.Context, key string) (string, *http.Response, error)
	GetAll(ctx context.Context) ([]model.EnvironmentVariable, *http.Response, error)
	GetTagsForEnvVar(ctx context.Context, name string) ([]model.Tag, *http.Response, error)
	PutTagForEnvVar(ctx context.Context, body []model.Tag, name string) (*http.Response, error)
}

func NewEnvironmentClient(apiClient *APIClient) EnvironmentClient {
	return &EnvironmentResourceApiService{apiClient}
}
