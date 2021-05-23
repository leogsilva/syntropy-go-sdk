# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicHealth**](PublicApi.md#PublicHealth) | **Get** /api/public/health | 
[**PublicInfo**](PublicApi.md#PublicInfo) | **Get** /api/public/info | 
[**PublicLanguageIndex**](PublicApi.md#PublicLanguageIndex) | **Get** /api/public/languages | 
[**PublicLinkIndex**](PublicApi.md#PublicLinkIndex) | **Get** /api/public/links | 
[**PublicTranslationIndex**](PublicApi.md#PublicTranslationIndex) | **Get** /api/public/translations | 
[**PublicVersion**](PublicApi.md#PublicVersion) | **Get** /api/public/versions | 

# **PublicHealth**
> PublicHealth(ctx, )


Health check used for smoke tests.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicInfo**
> ModelMap PublicInfo(ctx, )


Social providers application ids and versions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ModelMap**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicLanguageIndex**
> ModelMap PublicLanguageIndex(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ModelMap**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicLinkIndex**
> ModelMap PublicLinkIndex(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicApiPublicLinkIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiPublicLinkIndexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of ContextType**](.md)|  | 
 **lang** | **optional.String**|  | [default to en]

### Return type

[**ModelMap**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicTranslationIndex**
> ModelMap PublicTranslationIndex(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicApiPublicTranslationIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicApiPublicTranslationIndexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of ContextType**](.md)|  | 
 **lang** | **optional.String**|  | [default to en]

### Return type

[**ModelMap**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublicVersion**
> ModelMap PublicVersion(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ModelMap**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

