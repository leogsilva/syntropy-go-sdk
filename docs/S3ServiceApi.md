# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3AgentProviderImageDestroy**](S3ServiceApi.md#S3AgentProviderImageDestroy) | **Delete** /api/s3/delete-agent-provider-img/{key} | 
[**S3AgentProviderImageUpdate**](S3ServiceApi.md#S3AgentProviderImageUpdate) | **Post** /api/s3/upload-or-update-agent-provider-img | 
[**S3AgentProviderImagesIndex**](S3ServiceApi.md#S3AgentProviderImagesIndex) | **Get** /api/s3/get-agent-provider-imgs | 

# **S3AgentProviderImageDestroy**
> S3AgentProviderImageDestroy(ctx, key)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **S3AgentProviderImageUpdate**
> ModelMap S3AgentProviderImageUpdate(ctx, img)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **img** | ***os.File*****os.File**|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **S3AgentProviderImagesIndex**
> []ModelMap S3AgentProviderImagesIndex(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

