# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentInterfaceBwCreate**](AgentInterfaceBwApi.md#AgentInterfaceBwCreate) | **Post** /api/agent-interface-bw | 
[**AgentInterfaceBwDestroy**](AgentInterfaceBwApi.md#AgentInterfaceBwDestroy) | **Delete** /api/agent-interface-bw/{id} | 
[**AgentInterfaceBwIndex**](AgentInterfaceBwApi.md#AgentInterfaceBwIndex) | **Get** /api/agent-interface-bw | 
[**AgentInterfaceBwShow**](AgentInterfaceBwApi.md#AgentInterfaceBwShow) | **Get** /api/agent-interface-bw/{id} | 
[**AgentInterfaceBwUpdate**](AgentInterfaceBwApi.md#AgentInterfaceBwUpdate) | **Patch** /api/agent-interface-bw/{id} | 

# **AgentInterfaceBwCreate**
> ModelMap AgentInterfaceBwCreate(ctx, body)


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

# **AgentInterfaceBwDestroy**
> AgentInterfaceBwDestroy(ctx, id)


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

# **AgentInterfaceBwIndex**
> []ModelMap AgentInterfaceBwIndex(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentInterfaceBwApiAgentInterfaceBwIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentInterfaceBwApiAgentInterfaceBwIndexOpts struct
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

# **AgentInterfaceBwShow**
> ModelMap AgentInterfaceBwShow(ctx, id)


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

# **AgentInterfaceBwUpdate**
> ModelMap AgentInterfaceBwUpdate(ctx, body, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TsoaPartialAgentInterfaceBwObject_**](TsoaPartialAgentInterfaceBwObject_.md)|  | 
  **id** | **float64**|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

