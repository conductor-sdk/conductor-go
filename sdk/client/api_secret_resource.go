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

type SecretResourceApiService struct {
	*APIClient
}

/*
SecretResourceApiService Clear local cache
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return map[string]string
*/
func (a *SecretResourceApiService) ClearLocalCache(ctx context.Context) (map[string]string, *http.Response, error) {
	var result map[string]string
	path := "/secrets/clearLocalCache"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
SecretResourceApiService Clear redis cache
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return map[string]string
*/
func (a *SecretResourceApiService) ClearRedisCache(ctx context.Context) (map[string]string, *http.Response, error) {
	var result map[string]string

	path := "/secrets/clearRedisCache"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
SecretResourceApiService Delete a secret value by key
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param key
    @return interface{}
*/
func (a *SecretResourceApiService) DeleteSecret(ctx context.Context, key string) (interface{}, *http.Response, error) {
	var result interface{}

	path := fmt.Sprintf("/secrets/%s", key)
	resp, err := a.Delete(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
SecretResourceApiService Delete tags of the secret
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param key
*/
func (a *SecretResourceApiService) DeleteTagForSecret(ctx context.Context, body []model.Tag, key string) (*http.Response, error) {
	path := fmt.Sprintf("/secrets/%s/tags", key)

	resp, err := a.DeleteWithBody(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
SecretResourceApiService Get secret value by key
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param key
    @return string
*/
func (a *SecretResourceApiService) GetSecret(ctx context.Context, key string) (string, *http.Response, error) {
	var result string

	path := fmt.Sprintf("/secrets/%s", key)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

/*
SecretResourceApiService Get tags by secret
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param key
    @return []model.Tag
*/
func (a *SecretResourceApiService) GetTags(ctx context.Context, key string) ([]model.Tag, *http.Response, error) {
	var result []model.Tag

	path := fmt.Sprintf("/secrets/%s/tags", key)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
SecretResourceApiService List all secret names
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []string
*/
func (a *SecretResourceApiService) ListAllSecretNames(ctx context.Context) ([]string, *http.Response, error) {
	var result []string

	path := "/secrets"

	resp, err := a.Post(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
SecretResourceApiService List all secret names user can grant access to
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []string
*/
func (a *SecretResourceApiService) ListSecretsThatUserCanGrantAccessTo(ctx context.Context) ([]string, *http.Response, error) {
	var result []string

	path := "/secrets"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
SecretResourceApiService List all secret names along with tags user can grant access to
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []model.Secret
*/
func (a *SecretResourceApiService) ListSecretsWithTagsThatUserCanGrantAccessTo(ctx context.Context) ([]model.Secret, *http.Response, error) {
	var result []model.Secret

	path := "/secrets-v2"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
SecretResourceApiService Put a secret value by key
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param key
    @return interface{}
*/
func (a *SecretResourceApiService) PutSecret(ctx context.Context, body string, key string) (interface{}, *http.Response, error) {
	var result interface{}

	path := fmt.Sprintf("/secrets/%s", key)
	resp, err := a.Put(ctx, path, body, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
SecretResourceApiService Tag a secret
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param key
*/
func (a *SecretResourceApiService) PutTagForSecret(ctx context.Context, body []model.Tag, key string) (*http.Response, error) {
	path := fmt.Sprintf("/secrets/%s/tags", key)

	resp, err := a.Put(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
SecretResourceApiService Check if secret exists
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param key
    @return interface{}
*/
func (a *SecretResourceApiService) SecretExists(ctx context.Context, key string) (interface{}, *http.Response, error) {
	var result interface{}

	path := fmt.Sprintf("/secrets/%s/exists", key)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
