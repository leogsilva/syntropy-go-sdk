# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VersionCreate**](VersionsApi.md#VersionCreate) | **Post** /api/versions | 
[**VersionDestroy**](VersionsApi.md#VersionDestroy) | **Delete** /api/versions/{id} | 
[**VersionIndex**](VersionsApi.md#VersionIndex) | **Get** /api/versions | 
[**VersionShow**](VersionsApi.md#VersionShow) | **Get** /api/versions/{id} | 
[**VersionUpdate**](VersionsApi.md#VersionUpdate) | **Patch** /api/versions/{id} | 

# **VersionCreate**
> ModelMap VersionCreate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VersionDestroy**
> VersionDestroy(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VersionIndex**
> []ModelMap VersionIndex(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VersionsApiVersionIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VersionsApiVersionIndexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **optional.Float64**|  | 
 **take** | **optional.Float64**|  | 
 **order** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**[]ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VersionShow**
> ModelMap VersionShow(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **float64**|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VersionUpdate**
> ModelMap VersionUpdate(ctx, body, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TsoaPartialVersionObject_**](TsoaPartialVersionObject_.md)|  | 
  **id** | **float64**|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

