# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProxyGetWgInternalIpv4**](ProxyApi.md#ProxyGetWgInternalIpv4) | **Get** /api/proxy/platform-agents/get-wg-internal-ipv4 | 
[**ProxySendNftablesConf**](ProxyApi.md#ProxySendNftablesConf) | **Post** /api/proxy/sdn-agents/send-nftables-conf | 
[**ProxySendVppConf**](ProxyApi.md#ProxySendVppConf) | **Post** /api/proxy/sdn-agents/send-vpp-conf | 

# **ProxyGetWgInternalIpv4**
> string ProxyGetWgInternalIpv4(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProxySendNftablesConf**
> ModelMap ProxySendNftablesConf(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body2**](Body2.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProxySendVppConf**
> ModelMap ProxySendVppConf(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body3**](Body3.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

