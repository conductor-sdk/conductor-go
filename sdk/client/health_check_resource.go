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
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

// Linger please
var (
	_ context.Context
)

type HealthCheckResourceApiService struct {
	*APIClient
}

/*
HealthCheckResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return http_model.HealthCheckStatus
*/
func (a *HealthCheckResourceApiService) DoCheck(ctx context.Context) (model.HealthCheckStatus, *http.Response, error) {
	var result model.HealthCheckStatus

	path := "/health"
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return model.HealthCheckStatus{}, resp, err
	}
	return result, resp, nil
}
