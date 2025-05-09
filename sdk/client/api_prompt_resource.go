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
	"github.com/conductor-sdk/conductor-go/sdk/model/integration"
	"net/http"
	"net/url"
)

type PromptResourceApiService struct {
	*APIClient
}

/*
PromptResourceApiService Delete Template
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
*/
func (a *PromptResourceApiService) DeleteMessageTemplate(ctx context.Context, name string) (*http.Response, error) {
	path := fmt.Sprintf("/prompts/%s", name)

	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
PromptResourceApiService Delete a tag for Prompt Template
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
*/
func (a *PromptResourceApiService) DeleteTagForPromptTemplate(ctx context.Context, body []model.Tag, name string) (*http.Response, error) {
	path := fmt.Sprintf("/prompts/%s/tags", name)

	resp, err := a.DeleteWithBody(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
PromptResourceApiService Get Template
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return MessageTemplate
*/
func (a *PromptResourceApiService) GetMessageTemplate(ctx context.Context, name string) (*integration.PromptTemplate, *http.Response, error) {
	var result integration.PromptTemplate

	path := fmt.Sprintf("/prompts/%s", name)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

/*
PromptResourceApiService Get Templates
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []MessageTemplate
*/
func (a *PromptResourceApiService) GetMessageTemplates(ctx context.Context) ([]integration.PromptTemplate, *http.Response, error) {
	var result []integration.PromptTemplate
	path := "/prompts"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
PromptResourceApiService Get tags by Prompt Template
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return []model.Tag
*/
func (a *PromptResourceApiService) GetTagsForPromptTemplate(ctx context.Context, name string) ([]model.Tag, *http.Response, error) {
	var result []model.Tag
	path := fmt.Sprintf("/prompts/%s/tags", name)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
PromptResourceApiService Put a tag to Prompt Template
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
*/
func (a *PromptResourceApiService) PutTagForPromptTemplate(ctx context.Context, body []model.Tag, name string) (*http.Response, error) {
	path := fmt.Sprintf("/prompts/%s/tags", name)

	resp, err := a.Put(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

/*
   PromptResourceApiService Create or Update Template
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param body
    * @param description
    * @param name
    * @param optional nil or *PromptResourceApiSaveMessageTemplateOpts - Optional Parameters:
        * @param "Models" (optional.Interface of []string) -

*/

type PromptResourceApiSaveMessageTemplateOpts struct {
	Models []string
}

func (a *PromptResourceApiService) SaveMessageTemplate(ctx context.Context, body string, description string, name string, optionals *PromptResourceApiSaveMessageTemplateOpts) (*http.Response, error) {
	path := fmt.Sprintf("/prompts/%s", name)

	queryParams := url.Values{}
	queryParams.Add("description", parameterToString(description, ""))
	if optionals != nil {
		queryParams.Add("models", parameterToString(optionals.Models, "multi"))
	}
	resp, err := a.PostWithParams(ctx, path, queryParams, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
PromptResourceApiService Test Prompt Template
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
    @return string
*/
func (a *PromptResourceApiService) TestMessageTemplate(ctx context.Context, body model.PromptTemplateTestRequest) (string, *http.Response, error) {
	var result string

	path := "/prompts/test"

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}
