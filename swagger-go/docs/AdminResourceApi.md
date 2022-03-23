# {{classname}}

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllConfig**](AdminResourceApi.md#GetAllConfig) | **Get** /api/admin/config | Get all the configuration parameters
[**GetEventQueues**](AdminResourceApi.md#GetEventQueues) | **Get** /api/admin/queues | Get registered queues
[**RequeueSweep**](AdminResourceApi.md#RequeueSweep) | **Post** /api/admin/sweep/requeue/{workflowId} | Queue up all the running workflows for sweep
[**VerifyAndRepairWorkflowConsistency**](AdminResourceApi.md#VerifyAndRepairWorkflowConsistency) | **Post** /api/admin/consistency/verifyAndRepair/{workflowId} | Verify and repair workflow consistency
[**View**](AdminResourceApi.md#View) | **Get** /api/admin/task/{tasktype} | Get the list of pending tasks for a given task type

# **GetAllConfig**
> map[string]interface{} GetAllConfig(ctx, )
Get all the configuration parameters

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventQueues**
> map[string]interface{} GetEventQueues(ctx, optional)
Get registered queues

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdminResourceApiGetEventQueuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminResourceApiGetEventQueuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verbose** | **optional.Bool**|  | [default to false]

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequeueSweep**
> string RequeueSweep(ctx, workflowId)
Queue up all the running workflows for sweep

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyAndRepairWorkflowConsistency**
> string VerifyAndRepairWorkflowConsistency(ctx, workflowId)
Verify and repair workflow consistency

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **View**
> []Task View(ctx, tasktype, optional)
Get the list of pending tasks for a given task type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tasktype** | **string**|  | 
 **optional** | ***AdminResourceApiViewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminResourceApiViewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**|  | [default to 0]
 **count** | **optional.Int32**|  | [default to 100]

### Return type

[**[]Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

