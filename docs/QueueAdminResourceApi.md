# {{classname}}

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Names**](QueueAdminResourceApi.md#Names) | **Get** /api/queue/ | Get Queue Names
[**Size1**](QueueAdminResourceApi.md#Size1) | **Get** /api/queue/size | Get the queue length
[**Update1**](QueueAdminResourceApi.md#Update1) | **Post** /api/queue/update/{workflowId}/{taskRefName}/{status} | Publish a message in queue to mark a wait task as completed.
[**UpdateByTaskId**](QueueAdminResourceApi.md#UpdateByTaskId) | **Post** /api/queue/update/{workflowId}/task/{taskId}/{status} | Publish a message in queue to mark a wait task (by taskId) as completed.

# **Names**
> map[string]string Names(ctx, )
Get Queue Names

### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Size1**
> map[string]int64 Size1(ctx, )
Get the queue length

### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update1**
> Update1(ctx, body, workflowId, taskRefName, status)
Publish a message in queue to mark a wait task as completed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**map[string]interface{}**](map.md)|  | 
  **workflowId** | **string**|  | 
  **taskRefName** | **string**|  | 
  **status** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateByTaskId**
> UpdateByTaskId(ctx, body, workflowId, taskId, status)
Publish a message in queue to mark a wait task (by taskId) as completed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**map[string]interface{}**](map.md)|  | 
  **workflowId** | **string**|  | 
  **taskId** | **string**|  | 
  **status** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

