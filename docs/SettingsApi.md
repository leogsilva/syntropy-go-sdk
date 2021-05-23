# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingCreate**](SettingsApi.md#SettingCreate) | **Post** /api/settings | 
[**SettingIndex**](SettingsApi.md#SettingIndex) | **Get** /api/settings | 

# **SettingCreate**
> ModelMap SettingCreate(ctx, body)


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

# **SettingIndex**
> []ModelMap SettingIndex(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiSettingIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiSettingIndexOpts struct
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

