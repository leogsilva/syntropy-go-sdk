# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentConfig**](AgentsApi.md#AgentConfig) | **Get** /api/agents/{key}/config | 
[**AgentCreate**](AgentsApi.md#AgentCreate) | **Post** /api/agents | 
[**AgentDestroy**](AgentsApi.md#AgentDestroy) | **Delete** /api/agents/{id} | 
[**AgentIndex**](AgentsApi.md#AgentIndex) | **Get** /api/agents | 
[**AgentServerInfo**](AgentsApi.md#AgentServerInfo) | **Get** /api/agents/{key}/server-info | 
[**AgentShow**](AgentsApi.md#AgentShow) | **Get** /api/agents/{id} | 
[**AgentUpdate**](AgentsApi.md#AgentUpdate) | **Patch** /api/agents/{id} | 
[**FindByApiKey**](AgentsApi.md#FindByApiKey) | **Get** /api/agents/api-key/{api_key_id} | 

# **AgentConfig**
> Object AgentConfig(ctx, key, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 
 **optional** | ***AgentsApiAgentConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ip** | **optional.String**| a Agent IP is determined from request, but can be overwritten. Affects only agents with global key. | 
 **linkTag** | [**optional.Interface of LinkTag**](.md)|  | 

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentCreate**
> ModelMap AgentCreate(ctx, body)


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

# **AgentDestroy**
> AgentDestroy(ctx, id)


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

# **AgentIndex**
> []ModelMap AgentIndex(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentsApiAgentIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentIndexOpts struct
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

# **AgentServerInfo**
> ModelMap AgentServerInfo(ctx, key, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 
 **optional** | ***AgentsApiAgentServerInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentServerInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ip** | **optional.String**| Agent IP is determined from request, but can be overwritten. Affects only agents with global key. | 

### Return type

[**ModelMap**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentShow**
> ModelMap AgentShow(ctx, id)


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

# **AgentUpdate**
> ModelMap AgentUpdate(ctx, body, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TsoaPartialAgentObject_**](TsoaPartialAgentObject_.md)|  | 
  **id** | **float64**|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindByApiKey**
> []ModelMap FindByApiKey(ctx, apiKeyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyId** | [**int32**](.md)|  | 

### Return type

[**[]ModelMap**](map.md)

### Authorization

[api_key](../README.md#api_key), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

