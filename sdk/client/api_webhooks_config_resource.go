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

type WebhooksConfigResourceApiService struct {
	*APIClient
}

/*
WebhooksConfigResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
    @return WebhookConfig
*/
func (a *WebhooksConfigResourceApiService) CreateWebhook(ctx context.Context, body model.WebhookConfig) (model.WebhookConfig, *http.Response, error) {
	var result model.WebhookConfig

	path := "/metadata/webhook"

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return model.WebhookConfig{}, resp, err
	}
	return result, resp, nil
}

/*
WebhooksConfigResourceApiService Delete a tag for webhook id
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *WebhooksConfigResourceApiService) DeleteTagForWebhook(ctx context.Context, id string, body []model.Tag) (*http.Response, error) {
	path := fmt.Sprintf("/metadata/webhook/%s/tags", id)

	resp, err := a.DeleteWithBody(ctx, path, body, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
WebhooksConfigResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
*/
func (a *WebhooksConfigResourceApiService) DeleteWebhook(ctx context.Context, id string) (*http.Response, error) {
	path := fmt.Sprintf("/metadata/webhook/%s", id)

	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
WebhooksConfigResourceApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []WebhookConfig
*/
func (a *WebhooksConfigResourceApiService) GetAllWebhook(ctx context.Context) ([]model.WebhookConfig, *http.Response, error) {
	var result []model.WebhookConfig

	path := "/metadata/webhook"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return []model.WebhookConfig{}, resp, err
	}
	return result, resp, nil
}

/*
WebhooksConfigResourceApiService Get tags by webhook id
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return []Tag
*/
func (a *WebhooksConfigResourceApiService) GetTagsForWebhook(ctx context.Context, id string) ([]model.Tag, *http.Response, error) {
	var result []model.Tag

	path := fmt.Sprintf("/metadata/webhook/%s/tags", id)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return []model.Tag{}, resp, err
	}
	return result, resp, nil
}

/*
WebhooksConfigResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return WebhookConfig
*/
func (a *WebhooksConfigResourceApiService) GetWebhook(ctx context.Context, id string) (model.WebhookConfig, *http.Response, error) {
	var result model.WebhookConfig

	path := fmt.Sprintf("/metadata/webhook/%s", id)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return model.WebhookConfig{}, resp, err
	}
	return result, resp, nil
}

/*
WebhooksConfigResourceApiService Put a tag to webhook id
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param id
*/
func (a *WebhooksConfigResourceApiService) PutTagForWebhook(ctx context.Context, body []model.Tag, id string) (*http.Response, error) {
	path := fmt.Sprintf("/metadata/webhook/%s/tags", id)
	resp, err := a.Put(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
WebhooksConfigResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param id
    @return WebhookConfig
*/
func (a *WebhooksConfigResourceApiService) UpdateWebhook(ctx context.Context, body model.WebhookConfig, id string) (model.WebhookConfig, *http.Response, error) {
	var result model.WebhookConfig

	path := fmt.Sprintf("/metadata/webhook/%s", id)
	resp, err := a.Put(ctx, path, body, &result)
	if err != nil {
		return model.WebhookConfig{}, resp, err
	}
	return result, resp, nil
}
