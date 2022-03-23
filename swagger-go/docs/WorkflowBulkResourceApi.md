# {{classname}}

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PauseWorkflow1**](WorkflowBulkResourceApi.md#PauseWorkflow1) | **Put** /api/workflow/bulk/pause | Pause the list of workflows
[**Restart1**](WorkflowBulkResourceApi.md#Restart1) | **Post** /api/workflow/bulk/restart | Restart the list of completed workflow
[**ResumeWorkflow1**](WorkflowBulkResourceApi.md#ResumeWorkflow1) | **Put** /api/workflow/bulk/resume | Resume the list of workflows
[**Retry1**](WorkflowBulkResourceApi.md#Retry1) | **Post** /api/workflow/bulk/retry | Retry the last failed task for each workflow from the list
[**Terminate**](WorkflowBulkResourceApi.md#Terminate) | **Post** /api/workflow/bulk/terminate | Terminate workflows execution

# **PauseWorkflow1**
> BulkResponse PauseWorkflow1(ctx, body)
Pause the list of workflows

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)|  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Restart1**
> BulkResponse Restart1(ctx, body, optional)
Restart the list of completed workflow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)|  | 
 **optional** | ***WorkflowBulkResourceApiRestart1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowBulkResourceApiRestart1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useLatestDefinitions** | **optional.**|  | [default to false]

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumeWorkflow1**
> BulkResponse ResumeWorkflow1(ctx, body)
Resume the list of workflows

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)|  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Retry1**
> BulkResponse Retry1(ctx, body)
Retry the last failed task for each workflow from the list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)|  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Terminate**
> BulkResponse Terminate(ctx, body, optional)
Terminate workflows execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)|  | 
 **optional** | ***WorkflowBulkResourceApiTerminateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowBulkResourceApiTerminateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **optional.**|  | 

### Return type

[**BulkResponse**](BulkResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

