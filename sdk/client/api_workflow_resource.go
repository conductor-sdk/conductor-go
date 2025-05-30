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
	"fmt"
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type WorkflowResourceApiService struct {
	*APIClient
}

/*
WorkflowResourceApiService Starts the decision task for a workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workflowId
*/
func (a *WorkflowResourceApiService) Decide(ctx context.Context, workflowId string) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/decide/%s", workflowId)

	resp, err := a.Put(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
WorkflowResourceApiService Removes the workflow from the system
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowId
 * @param optional nil or *WorkflowResourceApiDeleteOpts - Optional Parameters:
     * @param "ArchiveWorkflow" (optional.Bool) -

*/

type WorkflowResourceApiDeleteOpts struct {
	ArchiveWorkflow optional.Bool
}

func (a *WorkflowResourceApiService) Delete(ctx context.Context, workflowId string, localVarOptionals *WorkflowResourceApiDeleteOpts) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/remove", workflowId)

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.ArchiveWorkflow.IsSet() {
		queryParams.Add("archiveWorkflow", parameterToString(localVarOptionals.ArchiveWorkflow.Value(), ""))
	}

	resp, err := a.APIClient.Delete(ctx, path, queryParams, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
WorkflowResourceApiService Gets the workflow by workflow id
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowId
 * @param optional nil or *WorkflowResourceApiGetExecutionStatusOpts - Optional Parameters:
     * @param "IncludeTasks" (optional.Bool) -
@return http_model.Workflow
*/

type WorkflowResourceApiGetExecutionStatusOpts struct {
	IncludeTasks optional.Bool
}

func (a *WorkflowResourceApiService) GetExecutionStatus(ctx context.Context, workflowId string, opts *WorkflowResourceApiGetExecutionStatusOpts) (model.Workflow, *http.Response, error) {
	var result model.Workflow

	path := fmt.Sprintf("/workflow/%s", workflowId)

	queryParams := url.Values{}
	if opts != nil && opts.IncludeTasks.IsSet() {
		queryParams.Add("includeTasks", parameterToString(opts.IncludeTasks.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.Workflow{}, resp, err
	}

	return result, resp, nil
}

func (a *WorkflowResourceApiService) GetWorkflowState(ctx context.Context, workflowId string, includeOutput bool, includeVariables bool) (model.WorkflowState, *http.Response, error) {
	var result model.WorkflowState

	path := fmt.Sprintf("/workflow/%s/status", workflowId)

	queryParams := url.Values{}
	queryParams.Add("includeOutput", parameterToString(includeOutput, ""))
	queryParams.Add("includeVariables", parameterToString(includeVariables, ""))

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.WorkflowState{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Get the uri and path of the external storage where the workflow payload is to be stored
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param path
  - @param operation
  - @param payloadType

@return http_model.ExternalStorageLocation
*/
func (a *WorkflowResourceApiService) GetExternalStorageLocation(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error) {
	var result model.ExternalStorageLocation

	path = "/workflow/externalstoragelocation"

	queryParams := url.Values{}
	queryParams.Add("path", parameterToString(path, ""))
	queryParams.Add("operation", parameterToString(operation, ""))
	queryParams.Add("payloadType", parameterToString(payloadType, ""))

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.ExternalStorageLocation{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Retrieve all the running workflows
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name
 * @param optional nil or *WorkflowResourceApiGetRunningWorkflowOpts - Optional Parameters:
     * @param "Version" (optional.Int32) -
     * @param "StartTime" (optional.Int64) -
     * @param "EndTime" (optional.Int64) -
@return []string
*/

type WorkflowResourceApiGetRunningWorkflowOpts struct {
	Version   optional.Int32
	StartTime optional.Int64
	EndTime   optional.Int64
}

func (a *WorkflowResourceApiService) GetRunningWorkflow(ctx context.Context, name string, opts *WorkflowResourceApiGetRunningWorkflowOpts) ([]string, *http.Response, error) {
	var result []string

	path := fmt.Sprintf("/workflow/running/%s", name)

	queryParams := url.Values{}
	if opts != nil && opts.Version.IsSet() {
		queryParams.Add("version", parameterToString(opts.Version.Value(), ""))
	}
	if opts != nil && opts.StartTime.IsSet() {
		queryParams.Add("startTime", parameterToString(opts.StartTime.Value(), ""))
	}
	if opts != nil && opts.EndTime.IsSet() {
		queryParams.Add("endTime", parameterToString(opts.EndTime.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Lists workflows for the given correlation id list
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param name
 * @param optional nil or *WorkflowResourceApiGetWorkflowsOpts - Optional Parameters:
     * @param "IncludeClosed" (optional.Bool) -
     * @param "IncludeTasks" (optional.Bool) -
@return map[string][]http_model.Workflow
*/

func (a *WorkflowResourceApiService) GetWorkflows(ctx context.Context, body []string, name string, opts *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error) {
	var result map[string][]model.Workflow

	path := fmt.Sprintf("/workflow/%s/correlated", name)

	queryParams := url.Values{}
	if opts != nil && opts.IncludeClosed.IsSet() {
		queryParams.Add("includeClosed", parameterToString(opts.IncludeClosed.Value(), ""))
	}
	if opts != nil && opts.IncludeTasks.IsSet() {
		queryParams.Add("includeTasks", parameterToString(opts.IncludeTasks.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (a *WorkflowResourceApiService) GetWorkflowsBatch(ctx context.Context, body map[string][]string, localVarOptionals *WorkflowResourceApiGetWorkflowsOpts) (map[string][]model.Workflow, *http.Response, error) {
	var result map[string][]model.Workflow

	path := "/workflow/correlated/batch"

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.IncludeClosed.IsSet() {
		queryParams.Add("includeClosed", parameterToString(localVarOptionals.IncludeClosed.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeTasks.IsSet() {
		queryParams.Add("includeTasks", parameterToString(localVarOptionals.IncludeTasks.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Lists workflows for the given correlation id
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name
 * @param correlationId
 * @param optional nil or *WorkflowResourceApiGetWorkflowsOpts - Optional Parameters:
     * @param "IncludeClosed" (optional.Bool) -
     * @param "IncludeTasks" (optional.Bool) -
@return []http_model.Workflow
*/

type WorkflowResourceApiGetWorkflowsOpts struct {
	IncludeClosed optional.Bool
	IncludeTasks  optional.Bool
}

func (a *WorkflowResourceApiService) GetWorkflowsByCorrelationId(ctx context.Context, name string, correlationId string, opts *WorkflowResourceApiGetWorkflowsOpts) ([]model.Workflow, *http.Response, error) {
	return a.GetWorkflows1(ctx, name, correlationId, opts)
}
func (a *WorkflowResourceApiService) GetWorkflows1(ctx context.Context, name string, correlationId string, opts *WorkflowResourceApiGetWorkflowsOpts) ([]model.Workflow, *http.Response, error) {
	var result []model.Workflow

	localVarPath := fmt.Sprintf("/workflow/%s/correlated/%s", name, correlationId)

	queryParams := url.Values{}
	if opts != nil && opts.IncludeClosed.IsSet() {
		queryParams.Add("includeClosed", parameterToString(opts.IncludeClosed.Value(), ""))
	}
	if opts != nil && opts.IncludeTasks.IsSet() {
		queryParams.Add("includeTasks", parameterToString(opts.IncludeTasks.Value(), ""))
	}

	resp, err := a.Get(ctx, localVarPath, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Pauses the workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workflowId
*/
func (a *WorkflowResourceApiService) PauseWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/pause", workflowId)

	resp, err := a.Put(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
WorkflowResourceApiService Reruns the workflow from a specific task
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param workflowId

@return string
*/
func (a *WorkflowResourceApiService) Rerun(ctx context.Context, body model.RerunWorkflowRequest, workflowId string) (string, *http.Response, error) {
	var result string

	path := fmt.Sprintf("/workflow/%s/rerun", workflowId)

	resp, err := a.Post(ctx, path, body, nil)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Resets callback times of all non-terminal SIMPLE tasks to 0
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workflowId
*/
func (a *WorkflowResourceApiService) ResetWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/resetcallbacks", workflowId)

	resp, err := a.Post(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
WorkflowResourceApiService Restarts a completed workflow
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowId
 * @param optional nil or *WorkflowResourceApiRestartOpts - Optional Parameters:
     * @param "UseLatestDefinitions" (optional.Bool) -

*/

type WorkflowResourceApiRestartOpts struct {
	UseLatestDefinitions optional.Bool
}

func (a *WorkflowResourceApiService) Restart(ctx context.Context, workflowId string, opts *WorkflowResourceApiRestartOpts) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/restart", workflowId)

	queryParams := url.Values{}
	if opts != nil && opts.UseLatestDefinitions.IsSet() {
		queryParams.Add("useLatestDefinitions", parameterToString(opts.UseLatestDefinitions.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
WorkflowResourceApiService Resumes the workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workflowId
*/
func (a *WorkflowResourceApiService) ResumeWorkflow(ctx context.Context, workflowId string) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/resume", workflowId)

	resp, err := a.Put(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
WorkflowResourceApiService Retries the last failed task
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowId
 * @param optional nil or *WorkflowResourceApiRetryOpts - Optional Parameters:
     * @param "ResumeSubworkflowTasks" (optional.Bool) -

*/

type WorkflowResourceApiRetryOpts struct {
	ResumeSubworkflowTasks optional.Bool
}

func (a *WorkflowResourceApiService) Retry(ctx context.Context, workflowId string, opts *WorkflowResourceApiRetryOpts) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/retry", workflowId)

	queryParams := url.Values{}
	if opts != nil && opts.ResumeSubworkflowTasks.IsSet() {
		queryParams.Add("resumeSubworkflowTasks", parameterToString(opts.ResumeSubworkflowTasks.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
WorkflowResourceApiService Search for workflows based on payload and other parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *WorkflowResourceApiSearchOpts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return http_model.SearchResultWorkflowSummary
*/

type WorkflowResourceApiSearchOpts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *WorkflowResourceApiService) Search(ctx context.Context, opts *WorkflowResourceApiSearchOpts) (model.SearchResultWorkflowSummary, *http.Response, error) {
	var result model.SearchResultWorkflowSummary

	path := "/workflow/search"

	queryParams := url.Values{}
	if opts != nil && opts.Start.IsSet() {
		queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
	}
	if opts != nil && opts.Size.IsSet() {
		queryParams.Add("size", parameterToString(opts.Size.Value(), ""))
	}
	if opts != nil && opts.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(opts.Sort.Value(), ""))
	}
	if opts != nil && opts.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(opts.FreeText.Value(), ""))
	}
	if opts != nil && opts.Query.IsSet() {
		queryParams.Add("query", parameterToString(opts.Query.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.SearchResultWorkflowSummary{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Search for workflows based on payload and other parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *WorkflowResourceApiSearchV2Opts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return http_model.SearchResultWorkflow
*/

type WorkflowResourceApiSearchV2Opts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *WorkflowResourceApiService) SearchV2(ctx context.Context, opts *WorkflowResourceApiSearchV2Opts) (model.SearchResultWorkflow, *http.Response, error) {
	var result model.SearchResultWorkflow

	path := "/workflow/search-v2"

	queryParams := url.Values{}
	if opts != nil && opts.Start.IsSet() {
		queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
	}
	if opts != nil && opts.Size.IsSet() {
		queryParams.Add("size", parameterToString(opts.Size.Value(), ""))
	}
	if opts != nil && opts.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(opts.Sort.Value(), ""))
	}
	if opts != nil && opts.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(opts.FreeText.Value(), ""))
	}
	if opts != nil && opts.Query.IsSet() {
		queryParams.Add("query", parameterToString(opts.Query.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.SearchResultWorkflow{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Search for workflows based on task parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *WorkflowResourceApiSearchWorkflowsByTasksOpts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return http_model.SearchResultWorkflowSummary
*/

type WorkflowResourceApiSearchWorkflowsByTasksOpts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *WorkflowResourceApiService) SearchWorkflowsByTasks(ctx context.Context, opts *WorkflowResourceApiSearchWorkflowsByTasksOpts) (model.SearchResultWorkflowSummary, *http.Response, error) {
	var result model.SearchResultWorkflowSummary

	localVarPath := "/workflow/search-by-tasks"

	queryParams := url.Values{}
	if opts != nil && opts.Start.IsSet() {
		queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
	}
	if opts != nil && opts.Size.IsSet() {
		queryParams.Add("size", parameterToString(opts.Size.Value(), ""))
	}
	if opts != nil && opts.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(opts.Sort.Value(), ""))
	}
	if opts != nil && opts.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(opts.FreeText.Value(), ""))
	}
	if opts != nil && opts.Query.IsSet() {
		queryParams.Add("query", parameterToString(opts.Query.Value(), ""))
	}

	resp, err := a.Get(ctx, localVarPath, queryParams, &result)
	if err != nil {
		return model.SearchResultWorkflowSummary{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Search for workflows based on task parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *WorkflowResourceApiSearchWorkflowsByTasksV2Opts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return http_model.SearchResultWorkflow
*/

type WorkflowResourceApiSearchWorkflowsByTasksV2Opts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *WorkflowResourceApiService) SearchWorkflowsByTasksV2(ctx context.Context, opts *WorkflowResourceApiSearchWorkflowsByTasksV2Opts) (model.SearchResultWorkflow, *http.Response, error) {
	var result model.SearchResultWorkflow

	localVarPath := "/workflow/search-by-tasks-v2"

	queryParams := url.Values{}
	if opts != nil && opts.Start.IsSet() {
		queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
	}
	if opts != nil && opts.Size.IsSet() {
		queryParams.Add("size", parameterToString(opts.Size.Value(), ""))
	}
	if opts != nil && opts.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(opts.Sort.Value(), ""))
	}
	if opts != nil && opts.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(opts.FreeText.Value(), ""))
	}
	if opts != nil && opts.Query.IsSet() {
		queryParams.Add("query", parameterToString(opts.Query.Value(), ""))
	}

	resp, err := a.Get(ctx, localVarPath, queryParams, &result)
	if err != nil {
		return model.SearchResultWorkflow{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Skips a given task from a current running workflow
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param workflowId
  - @param taskReferenceName
  - @param skipTaskRequest
*/
func (a *WorkflowResourceApiService) SkipTaskFromWorkflow(ctx context.Context, workflowId string, taskReferenceName string, skipTaskRequest model.SkipTaskRequest) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s/skiptask/%s", workflowId, taskReferenceName)

	queryParams := url.Values{}
	queryParams.Add("skipTaskRequest", parameterToString(skipTaskRequest, ""))

	resp, err := a.PutWithParams(ctx, path, queryParams, nil, &model.SkipTaskRequest{})
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
WorkflowResourceApiService Start a new workflow. Returns the ID of the workflow instance that can be later used for tracking
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param name
 * @param optional nil or *WorkflowResourceApiStartWorkflowOpts - Optional Parameters:
     * @param "Version" (optional.Int32) -
     * @param "CorrelationId" (optional.String) -
     * @param "Priority" (optional.Int32) -
@return string
*/

type WorkflowResourceApiStartWorkflowOpts struct {
	Version       optional.Int32
	CorrelationId optional.String
	Priority      optional.Int32
}

func (a *WorkflowResourceApiService) StartWorkflow(ctx context.Context, body map[string]interface{}, name string, opts *WorkflowResourceApiStartWorkflowOpts) (string, *http.Response, error) {
	var result string

	path := fmt.Sprintf("/workflow/%s", name)

	queryParams := url.Values{}
	if opts != nil && opts.Version.IsSet() {
		queryParams.Add("version", parameterToString(opts.Version.Value(), ""))
	}
	if opts != nil && opts.CorrelationId.IsSet() {
		queryParams.Add("correlationId", parameterToString(opts.CorrelationId.Value(), ""))
	}
	if opts != nil && opts.Priority.IsSet() {
		queryParams.Add("priority", parameterToString(opts.Priority.Value(), ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

func (a *WorkflowResourceApiService) ExecuteWorkflow(ctx context.Context, body model.StartWorkflowRequest, requestId string, name string, version int32, waitUntilTask string) (model.WorkflowRun, *http.Response, error) {
	var result model.WorkflowRun

	path := fmt.Sprintf("/workflow/execute/%s/%d", name, version)

	queryParams := url.Values{}
	queryParams.Add("requestId", parameterToString(requestId, ""))
	if len(waitUntilTask) > 0 {
		queryParams.Add("waitUntilTaskRef", parameterToString(waitUntilTask, ""))
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	if err != nil {
		return model.WorkflowRun{}, resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Start a new workflow with http_model.StartWorkflowRequest, which allows task to be executed in a domain
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body

@return string
*/
func (a *WorkflowResourceApiService) StartWorkflowWithRequest(ctx context.Context, body model.StartWorkflowRequest) (string, *http.Response, error) {
	var result string

	path := "/workflow"

	resp, err := a.Post(ctx, path, body, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

/*
WorkflowResourceApiService Terminate workflow execution
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param workflowId
 * @param optional nil or *WorkflowResourceApiTerminateOpts - Optional Parameters:
     * @param "Reason" (optional.String) -

*/

type WorkflowResourceApiTerminateOpts struct {
	Reason                 optional.String
	TriggerFailureWorkflow optional.Bool
}

func (a *WorkflowResourceApiService) Terminate(ctx context.Context, workflowId string, opts *WorkflowResourceApiTerminateOpts) (*http.Response, error) {
	path := fmt.Sprintf("/workflow/%s", workflowId)

	queryParams := url.Values{}
	if opts != nil && opts.Reason.IsSet() {
		queryParams.Add("reason", parameterToString(opts.Reason.Value(), ""))
	}
	if opts != nil && opts.TriggerFailureWorkflow.IsSet() {
		queryParams.Add("triggerFailureWorkflow", parameterToString(opts.TriggerFailureWorkflow.Value(), ""))
	}

	resp, err := a.APIClient.Delete(ctx, path, queryParams, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
