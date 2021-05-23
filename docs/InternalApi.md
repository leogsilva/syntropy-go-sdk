# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InternalAgentRestart**](InternalApi.md#InternalAgentRestart) | **Post** /api/internal/agent-restart | 
[**InternalBiStats**](InternalApi.md#InternalBiStats) | **Get** /api/internal/bi-stats | 
[**InternalMaintenance**](InternalApi.md#InternalMaintenance) | **Get** /api/internal/maintenance | 
[**InternalServerRrestart**](InternalApi.md#InternalServerRrestart) | **Post** /api/internal/server-restart | 
[**InternalServers**](InternalApi.md#InternalServers) | **Delete** /api/internal/servers | 
[**InternalTextToQrcode**](InternalApi.md#InternalTextToQrcode) | **Post** /api/internal/text-to-qrcode | 
[**InternalUsersData**](InternalApi.md#InternalUsersData) | **Delete** /api/internal/users-data | 
[**InternalWgKeygen**](InternalApi.md#InternalWgKeygen) | **Get** /api/internal/wg-keygen | 

# **InternalAgentRestart**
> InternalAgentRestart(ctx, body)


Restart agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalBiStats**
> ModelMap InternalBiStats(ctx, )


BI related statistics.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalMaintenance**
> InternalMaintenance(ctx, )


Always returns maintenance status. Used for development only.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalServerRrestart**
> InternalServerRrestart(ctx, body)


Restart single server. If server id is not supplied, restarts all servers. Deletes users data if all servers are restared.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)|  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalServers**
> InternalServers(ctx, )


Destroy servers data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalTextToQrcode**
> Object InternalTextToQrcode(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body**](Body.md)|  | 

### Return type

**Object**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalUsersData**
> InternalUsersData(ctx, )


Destroy users' data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalWgKeygen**
> ModelMap InternalWgKeygen(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

