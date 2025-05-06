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
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type WorkflowBulkResourceApiService struct {
	*APIClient
}

/*
WorkflowBulkResourceApiService Pause the list of workflows
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body

@return http_model.BulkResponse
*/
func (a *WorkflowBulkResourceApiService) PauseWorkflow1(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	localVarPath := "/workflow/bulk/pause"

	resp, err := a.Put(ctx, localVarPath, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowBulkResourceApiService Restart the list of completed workflow
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param optional nil or *WorkflowBulkResourceApiRestart1Opts - Optional Parameters:
     * @param "UseLatestDefinitions" (optional.Bool) -
@return http_model.BulkResponse
*/

type WorkflowBulkResourceApiRestart1Opts struct {
	UseLatestDefinitions optional.Bool
}

func (a *WorkflowBulkResourceApiService) Restart(ctx context.Context, body []string, localVarOptionals *WorkflowBulkResourceApiRestart1Opts) (model.BulkResponse, *http.Response, error) {
	return a.Restart1(ctx, body, localVarOptionals)
}
func (a *WorkflowBulkResourceApiService) Restart1(ctx context.Context, body []string, localVarOptionals *WorkflowBulkResourceApiRestart1Opts) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	path := "/workflow/bulk/restart"

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.UseLatestDefinitions.IsSet() {
		queryParams.Add("useLatestDefinitions", parameterToString(localVarOptionals.UseLatestDefinitions.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}

func (a *WorkflowBulkResourceApiService) ResumeWorkflow(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	return a.ResumeWorkflow1(ctx, body)
}

/*
WorkflowBulkResourceApiService Resume the list of workflows
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body

@return http_model.BulkResponse
*/
func (a *WorkflowBulkResourceApiService) ResumeWorkflow1(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	path := "/workflow/bulk/resume"

	resp, err := a.Put(ctx, path, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowBulkResourceApiService Retry the last failed task for each workflow from the list
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body

@return http_model.BulkResponse
*/
func (a *WorkflowBulkResourceApiService) Retry(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	return a.Retry1(ctx, body)
}
func (a *WorkflowBulkResourceApiService) Retry1(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	path := "/workflow/bulk/retry"

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowBulkResourceApiService Terminate workflows execution
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param optional nil or *WorkflowBulkResourceApiTerminateOpts - Optional Parameters:
     * @param "Reason" (optional.String) -
@return http_model.BulkResponse
*/

type WorkflowBulkResourceApiTerminateOpts struct {
	Reason                 optional.String
	TriggerFailureWorkflow optional.Bool
}

func (a *WorkflowBulkResourceApiService) Terminate(ctx context.Context, body []string, opts *WorkflowBulkResourceApiTerminateOpts) (model.BulkResponse, *http.Response, error) {
	var result model.BulkResponse

	path := "/workflow/bulk/terminate"

	queryParams := url.Values{}
	if opts != nil && opts.Reason.IsSet() {
		queryParams.Add("reason", parameterToString(opts.Reason.Value(), ""))
	}
	if opts != nil && opts.TriggerFailureWorkflow.IsSet() {
		queryParams.Add("triggerFailureWorkflow", parameterToString(opts.TriggerFailureWorkflow.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return model.BulkResponse{}, resp, err
	}
	return result, resp, nil
}
