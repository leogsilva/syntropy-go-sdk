# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentProviderCreate**](AgentProvidersApi.md#AgentProviderCreate) | **Post** /api/agent-providers | 
[**AgentProviderDestroy**](AgentProvidersApi.md#AgentProviderDestroy) | **Delete** /api/agent-providers/{id} | 
[**AgentProviderIndex**](AgentProvidersApi.md#AgentProviderIndex) | **Get** /api/agent-providers | 
[**AgentProviderShow**](AgentProvidersApi.md#AgentProviderShow) | **Get** /api/agent-providers/{id} | 
[**AgentProviderUpdate**](AgentProvidersApi.md#AgentProviderUpdate) | **Patch** /api/agent-providers/{id} | 

# **AgentProviderCreate**
> ModelMap AgentProviderCreate(ctx, body)


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

# **AgentProviderDestroy**
> AgentProviderDestroy(ctx, id)


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

# **AgentProviderIndex**
> []ModelMap AgentProviderIndex(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentProvidersApiAgentProviderIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentProvidersApiAgentProviderIndexOpts struct
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

# **AgentProviderShow**
> ModelMap AgentProviderShow(ctx, id)


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

# **AgentProviderUpdate**
> ModelMap AgentProviderUpdate(ctx, body, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TsoaPartialAgentProviderObject_**](TsoaPartialAgentProviderObject_.md)|  | 
  **id** | **float64**|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

