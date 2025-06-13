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
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"net/http"
	"net/url"
	"strings"
)

type IntegrationResourceApiService struct {
	*APIClient
}

/*
IntegrationResourceApiService Associate a Prompt Template with an Integration
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param integrationProvider
  - @param integrationName
  - @param promptName
*/
func (a *IntegrationResourceApiService) AssociatePromptWithIntegration(ctx context.Context, integrationProvider string, integrationName string, promptName string) (*http.Response, error) {
	path := fmt.Sprintf("/integrations/provider/%s/integration/%s/prompt/%s", integrationProvider, integrationName, promptName)

	resp, err := a.Post(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

/*
IntegrationResourceApiService Delete an Integration
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
  - @param integrationName
*/
func (a *IntegrationResourceApiService) DeleteIntegrationApi(ctx context.Context, name string, integrationName string) (*http.Response, error) {
	path := fmt.Sprintf("/integrations/provider/%s/integration/%s", name, integrationName)

	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

/*
IntegrationResourceApiService Delete an Integration Provider
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
*/
func (a *IntegrationResourceApiService) DeleteIntegrationProvider(ctx context.Context, name string) (*http.Response, error) {
	path := fmt.Sprintf("/integrations/provider/%s", name)
	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

/*
IntegrationResourceApiService Delete a tag for Integration
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
  - @param integrationName
*/
func (a *IntegrationResourceApiService) DeleteTagForIntegration(ctx context.Context, tags []model.TagObject, name string, model string) (*http.Response, error) {
	path := fmt.Sprintf("/integrations/provider/%s/integration/%s/tags", name, model)
	resp, err := a.DeleteWithBody(ctx, path, tags, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
IntegrationResourceApiService Delete a tag for Integration Provider
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
*/
func (a *IntegrationResourceApiService) DeleteTagForIntegrationProvider(ctx context.Context, tags []model.TagObject, name string) (*http.Response, error) {
	path := fmt.Sprintf("/integrations/provider/%s/tags", name)

	resp, err := a.DeleteWithBody(ctx, path, tags, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

/*
IntegrationResourceApiService Get Integration details
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
  - @param integrationName

@return IntegrationApi
*/
func (a *IntegrationResourceApiService) GetIntegrationApi(ctx context.Context, name string, model string) (integration.IntegrationApi, *http.Response, error) {
	var result integration.IntegrationApi

	path := fmt.Sprintf("/integrations/provider/%s/integration/%s", name, model)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return integration.IntegrationApi{}, resp, err
	}
	return result, resp, nil
}

/*
IntegrationResourceApiService Get Integrations of an Integration Provider
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name
 * @param optional nil or *IntegrationResourceApiGetIntegrationApisOpts - Optional Parameters:
     * @param "ActiveOnly" (optional.Bool) -
@return []IntegrationApi
*/

func (a *IntegrationResourceApiService) GetIntegrationApis(ctx context.Context, name string, ActiveOnly optional.Bool) ([]integration.IntegrationApi, *http.Response, error) {
	var result []integration.IntegrationApi
	path := fmt.Sprintf("/integrations/provider/%s/integration", name)

	queryParams := url.Values{}
	if ActiveOnly.IsSet() {
		queryParams.Add("activeOnly", parameterToString(ActiveOnly.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
IntegrationResourceApiService Get Integrations Available for an Integration Provider
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name

@return []string
*/
func (a *IntegrationResourceApiService) GetIntegrationAvailableApis(ctx context.Context, name string) ([]string, *http.Response, error) {
	var result []string
	path := fmt.Sprintf("/integrations/provider/%s/integration/all", name)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
IntegrationResourceApiService Get Integration provider
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name

@return Integration
*/
func (a *IntegrationResourceApiService) GetIntegrationProvider(ctx context.Context, name string) (integration.Integration, *http.Response, error) {
	var result integration.Integration

	path := fmt.Sprintf("/integrations/provider/%s", name)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return integration.Integration{}, resp, err
	}

	return result, resp, nil
}

/*
IntegrationResourceApiService Get all Integrations Providers
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetIntegrationProvidersOpts - Optional Parameters:
     * @param "Category" (optional.String) -
     * @param "ActiveOnly" (optional.Bool) -
@return []Integration
*/

func (a *IntegrationResourceApiService) GetIntegrationProviders(ctx context.Context, localVarOptionals *GetIntegrationProvidersOpts) ([]integration.Integration, *http.Response, error) {
	var result []integration.Integration

	path := "/integrations/provider"

	localVarHeaderParams := make(map[string]string)
	localVarHeaderParams["Accept"] = "application/json"

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.Category.IsSet() {
		queryParams.Add("category", parameterToString(localVarOptionals.Category.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ActiveOnly.IsSet() {
		queryParams.Add("activeOnly", parameterToString(localVarOptionals.ActiveOnly.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
IntegrationResourceApiService Get the list of prompt templates associated with an integration
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param integrationProvider
  - @param integrationName

@return []MessageTemplate
*/
func (a *IntegrationResourceApiService) GetPromptsWithIntegration(ctx context.Context, integrationProvider string, integrationName string) ([]integration.PromptTemplate, *http.Response, error) {
	var result []integration.PromptTemplate

	path := fmt.Sprintf("/integrations/provider/%s/integration/%s/prompt", integrationProvider, integrationName)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
IntegrationResourceApiService Get Integrations Providers and Integrations combo
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *IntegrationResourceApiGetProvidersAndIntegrationsOpts - Optional Parameters:
     * @param "Type_" (optional.String) -
     * @param "ActiveOnly" (optional.Bool) -
@return []string
*/

type IntegrationResourceApiGetProvidersAndIntegrationsOpts struct {
	Type_      optional.String
	ActiveOnly optional.Bool
}

func (a *IntegrationResourceApiService) GetProvidersAndIntegrations(ctx context.Context, localVarOptionals *IntegrationResourceApiGetProvidersAndIntegrationsOpts) ([]string, *http.Response, error) {
	var result []string

	localVarPath := "/integrations/all"

	queryParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Type_.IsSet() {
		queryParams.Add("type", parameterToString(localVarOptionals.Type_.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ActiveOnly.IsSet() {
		queryParams.Add("activeOnly", parameterToString(localVarOptionals.ActiveOnly.Value(), ""))
	}

	resp, err := a.Get(ctx, localVarPath, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
IntegrationResourceApiService Get tags by Integration
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
  - @param integrationName

@return []Tag
*/
func (a *IntegrationResourceApiService) GetTagsForIntegration(ctx context.Context, name string, integrationName string) ([]model.TagObject, *http.Response, error) {
	var result []model.TagObject

	path := fmt.Sprintf("/integrations/provider/%s/integration/%s/tags", name, integrationName)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
IntegrationResourceApiService Get tags by Integration Provider
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name

@return []Tag
*/
func (a *IntegrationResourceApiService) GetTagsForIntegrationProvider(ctx context.Context, name string) ([]model.TagObject, *http.Response, error) {
	var result []model.TagObject

	path := fmt.Sprintf("/integrations/provider/%s/tags", name)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
IntegrationResourceApiService Get Token Usage by Integration
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
  - @param integrationName

@return int32
*/
func (a *IntegrationResourceApiService) GetTokenUsageForIntegration(ctx context.Context, integration string, model string) (int32, *http.Response, error) {
	var result int32

	path := fmt.Sprintf("/integrations/provider/%s/integration/%s/metrics", integration, model)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return 0, resp, err
	}
	return result, resp, nil
}

/*
IntegrationResourceApiService Get Token Usage by Integration Provider
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name

@return map[string]string
*/
func (a *IntegrationResourceApiService) GetTokenUsageForIntegrationProvider(ctx context.Context, name string) (map[string]string, *http.Response, error) {
	var result map[string]string

	path := fmt.Sprintf("/integrations/provider/%s/metrics", name)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
IntegrationResourceApiService Put a tag to Integration
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
  - @param integrationName
*/
func (a *IntegrationResourceApiService) UpdateTagForIntegration(ctx context.Context, tags []model.TagObject, name string, model string) (*http.Response, error) {
	path := fmt.Sprintf("/integrations/provider/%s/integration/%s/tags", name, model)

	resp, err := a.Put(ctx, path, tags, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
IntegrationResourceApiService Put a tag to Integration Provider
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
*/
func (a *IntegrationResourceApiService) UpdateTagForIntegrationProvider(ctx context.Context, tags []model.TagObject, name string) (*http.Response, error) {
	path := fmt.Sprintf("/integrations/provider/%s/tags", name)

	resp, err := a.Put(ctx, path, tags, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
IntegrationResourceApiService Create or Update Integration
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
  - @param integrationName
*/
func (a *IntegrationResourceApiService) SaveIntegrationApi(ctx context.Context, integrationApiUpdate integration.IntegrationApiUpdate, name string, integrationName string) (*http.Response, error) {
	path := fmt.Sprintf("/integrations/provider/%s/integration/%s", name, integrationName)

	resp, err := a.Post(ctx, path, integrationApiUpdate, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
IntegrationResourceApiService Create or Update Integration provider
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
*/
func (a *IntegrationResourceApiService) SaveIntegrationProvider(ctx context.Context, integrationUpdate integration.IntegrationUpdate, name string) (*http.Response, error) {
	path := fmt.Sprintf("/integrations/provider/%s", name)
	resp, err := a.Post(ctx, path, integrationUpdate, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
   IntegrationResourceApiService Get all Integrations
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param optional nil or *IntegrationResourceApiGetAllIntegrationsOpts - Optional Parameters:
        * @param "Category" (optional.String) -
    * @param "ActiveOnly" (optional.Bool) -
   @return []Integration
*/

type IntegrationResourceApiGetAllIntegrationsOpts struct {
	Category   optional.String
	ActiveOnly optional.Bool
}

func (a *IntegrationResourceApiService) GetAllIntegrations(ctx context.Context, optionals *IntegrationResourceApiGetAllIntegrationsOpts) ([]model.Integration, *http.Response, error) {
	var result []model.Integration

	// create path and map variables
	path := "/integrations/"

	queryParams := url.Values{}
	if optionals != nil && optionals.Category.IsSet() {
		queryParams.Add("category", parameterToString(optionals.Category.Value(), ""))
	}
	if optionals != nil && optionals.ActiveOnly.IsSet() {
		queryParams.Add("activeOnly", parameterToString(optionals.ActiveOnly.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
IntegrationResourceApiService Get Integration provider definitions
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []IntegrationDef
*/
func (a *IntegrationResourceApiService) GetIntegrationProviderDefs(ctx context.Context) ([]model.IntegrationDef, *http.Response, error) {
	var result []model.IntegrationDef

	// create path and map variables
	path := "/integrations/def"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
IntegrationResourceApiService Record Event Stats
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param type_
*/
func (a *IntegrationResourceApiService) RecordEventStats(ctx context.Context, body []model.EventLog, type_ string) (*http.Response, error) {
	// create path and map variables
	path := fmt.Sprintf("/integrations/eventStats/%s", type_)
	path = strings.Replace(path, "{"+"type"+"}", fmt.Sprintf("%v", type_), -1)

	resp, err := a.Post(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
IntegrationResourceApiService Register Token usage
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
  - @param integrationName
*/
func (a *IntegrationResourceApiService) RegisterTokenUsage(ctx context.Context, body int32, name string, integrationName string) (*http.Response, error) {
	// create path and map variables
	path := fmt.Sprintf("/integrations/provider/%s/integration/%s/metrics", name, integrationName)
	resp, err := a.Post(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
