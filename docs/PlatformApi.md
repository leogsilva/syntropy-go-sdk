# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformAdminAgentConfig**](PlatformApi.md#PlatformAdminAgentConfig) | **Get** /api/platform/admin/agent/{agent_id}/config | 
[**PlatformAgentConfig**](PlatformApi.md#PlatformAgentConfig) | **Get** /api/platform/agent/{agent_id}/config | 
[**PlatformAgentCoordinates**](PlatformApi.md#PlatformAgentCoordinates) | **Post** /api/platform/agents/coordinates | 
[**PlatformAgentCreate**](PlatformApi.md#PlatformAgentCreate) | **Post** /api/platform/agents | 
[**PlatformAgentDestroy**](PlatformApi.md#PlatformAgentDestroy) | **Delete** /api/platform/agents/{agent_id} | 
[**PlatformAgentGroupDestroy**](PlatformApi.md#PlatformAgentGroupDestroy) | **Delete** /api/platform/network/agent-groups/{group_id} | 
[**PlatformAgentGroupUpdate**](PlatformApi.md#PlatformAgentGroupUpdate) | **Put** /api/platform/network/agent-groups/{group_id} | 
[**PlatformAgentIdNamePairs**](PlatformApi.md#PlatformAgentIdNamePairs) | **Get** /api/platform/agents/filters | 
[**PlatformAgentIndex**](PlatformApi.md#PlatformAgentIndex) | **Get** /api/platform/agents | 
[**PlatformAgentProviderIndex**](PlatformApi.md#PlatformAgentProviderIndex) | **Get** /api/platform/agent-providers | 
[**PlatformAgentProviderShow**](PlatformApi.md#PlatformAgentProviderShow) | **Get** /api/platform/agent-providers/{id} | 
[**PlatformAgentServiceDestroy**](PlatformApi.md#PlatformAgentServiceDestroy) | **Post** /api/platform/agent-services-delete | 
[**PlatformAgentServiceIndex**](PlatformApi.md#PlatformAgentServiceIndex) | **Get** /api/platform/agent-services | 
[**PlatformAgentServiceSubnetUpdate**](PlatformApi.md#PlatformAgentServiceSubnetUpdate) | **Post** /api/platform/agent-services-subnets | 
[**PlatformAgentTagIndex**](PlatformApi.md#PlatformAgentTagIndex) | **Get** /api/platform/agent-tags | 
[**PlatformAgentUpdate**](PlatformApi.md#PlatformAgentUpdate) | **Patch** /api/platform/agents/{agent_id} | 
[**PlatformAgentsDestroy**](PlatformApi.md#PlatformAgentsDestroy) | **Post** /api/platform/agents/remove | 
[**PlatformApiKeyCreate**](PlatformApi.md#PlatformApiKeyCreate) | **Post** /api/platform/api-keys | 
[**PlatformApiKeyDestroy**](PlatformApi.md#PlatformApiKeyDestroy) | **Delete** /api/platform/api-keys/{api_key_id} | 
[**PlatformApiKeyIndex**](PlatformApi.md#PlatformApiKeyIndex) | **Get** /api/platform/api-keys | 
[**PlatformConnectionAgentDestroy**](PlatformApi.md#PlatformConnectionAgentDestroy) | **Delete** /api/platform/connections/agents/{agent_id} | 
[**PlatformConnectionAgentsDestroy**](PlatformApi.md#PlatformConnectionAgentsDestroy) | **Post** /api/platform/connections/agents/remove | 
[**PlatformConnectionCreate**](PlatformApi.md#PlatformConnectionCreate) | **Post** /api/platform/connections | 
[**PlatformConnectionCreateMesh**](PlatformApi.md#PlatformConnectionCreateMesh) | **Post** /api/platform/connections/mesh | 
[**PlatformConnectionCreateP2p**](PlatformApi.md#PlatformConnectionCreateP2p) | **Post** /api/platform/connections/point-to-point | 
[**PlatformConnectionDestroy**](PlatformApi.md#PlatformConnectionDestroy) | **Post** /api/platform/connections/remove | 
[**PlatformConnectionDestroyDeprecated**](PlatformApi.md#PlatformConnectionDestroyDeprecated) | **Delete** /api/platform/connections/{connection_id} | 
[**PlatformConnectionIndex**](PlatformApi.md#PlatformConnectionIndex) | **Get** /api/platform/connections | 
[**PlatformConnectionServiceShow**](PlatformApi.md#PlatformConnectionServiceShow) | **Get** /api/platform/connection-services | 
[**PlatformConnectionServiceUpdate**](PlatformApi.md#PlatformConnectionServiceUpdate) | **Post** /api/platform/connection-services | 
[**PlatformConnectionSubnetDestroy**](PlatformApi.md#PlatformConnectionSubnetDestroy) | **Post** /api/platform/connection-services-delete | 
[**PlatformLogsReadTimestamp**](PlatformApi.md#PlatformLogsReadTimestamp) | **Post** /api/platform/logs-reads-timestamp | 
[**PlatformNetworkAgentCreate**](PlatformApi.md#PlatformNetworkAgentCreate) | **Post** /api/platform/network/{network_id}/agents/add | 
[**PlatformNetworkAgentCreateDeprecated**](PlatformApi.md#PlatformNetworkAgentCreateDeprecated) | **Post** /api/platform/network/{network_id}/agents | 
[**PlatformNetworkAgentDestroy**](PlatformApi.md#PlatformNetworkAgentDestroy) | **Delete** /api/platform/networks/{network_id}/agents/{agent_id} | 
[**PlatformNetworkAgentGroupCreate**](PlatformApi.md#PlatformNetworkAgentGroupCreate) | **Post** /api/platform/network/{network_id}/agent-groups/{group_name} | 
[**PlatformNetworkAgentRemove**](PlatformApi.md#PlatformNetworkAgentRemove) | **Post** /api/platform/networks/{network_id}/agents/remove | 
[**PlatformNetworkAgentRemoveDeprecated**](PlatformApi.md#PlatformNetworkAgentRemoveDeprecated) | **Delete** /api/platform/networks/{network_id}/agents | 
[**PlatformNetworkCreate**](PlatformApi.md#PlatformNetworkCreate) | **Post** /api/platform/networks | 
[**PlatformNetworkDestroy**](PlatformApi.md#PlatformNetworkDestroy) | **Delete** /api/platform/networks/{network_id} | 
[**PlatformNetworkIndex**](PlatformApi.md#PlatformNetworkIndex) | **Get** /api/platform/networks | 
[**PlatformNetworkInfo**](PlatformApi.md#PlatformNetworkInfo) | **Get** /api/platform/network/{network_id}/info | 
[**PlatformNetworkNetworkAgentDestroyDeprecated**](PlatformApi.md#PlatformNetworkNetworkAgentDestroyDeprecated) | **Post** /api/platform/network/{network_id}/agents/delete | 
[**PlatformNetworkTopology**](PlatformApi.md#PlatformNetworkTopology) | **Get** /api/platform/networks/topology | 

# **PlatformAdminAgentConfig**
> ModelMap PlatformAdminAgentConfig(ctx, agentId)


Returns agent configuration details (admin).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentId** | [**int32**](.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentConfig**
> ModelMap PlatformAgentConfig(ctx, agentId)


Returns agent configuration details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentId** | [**int32**](.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentCoordinates**
> ModelMap PlatformAgentCoordinates(ctx, body)


Retrieves Platform agent location coordinates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body1**](Body1.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentCreate**
> ModelMap PlatformAgentCreate(ctx, body)


Creates new `platform agent`.

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

# **PlatformAgentDestroy**
> InlineResponse204 PlatformAgentDestroy(ctx, agentId)


Deletes `platform agent` ands its `connections`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentGroupDestroy**
> InlineResponse204 PlatformAgentGroupDestroy(ctx, groupId)


Remove agent group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentGroupUpdate**
> InlineResponse204 PlatformAgentGroupUpdate(ctx, body, groupId)


Update network agents group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]float64**](float64.md)|  | 
  **groupId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentIdNamePairs**
> ModelMap PlatformAgentIdNamePairs(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlatformApiPlatformAgentIdNamePairsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformAgentIdNamePairsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**optional.Interface of string**](.md)| ids[]: array of agent ids, example: \&quot;1;2;3\&quot;, id|name: agent id or agent name, example: \&quot;name1\&quot; or \&quot;132\&quot;, name: exact agent name, example: \&quot;name1\&quot;, statuses[]: one of: connected, connected_with_errors, disconnected, example: \&quot;connected;connected_with_errors\&quot;, networks[]: array of network ids, example: \&quot;1;2;3\&quot;, providers[]: array of providers ids, example: \&quot;1;2;3\&quot;, tags[]: array of tags ids, example: \&quot;1;2;3\&quot;, tags_names[]: array of tags name, example: \&quot;name1;name2;name3\&quot;, networks_names[]: array of networks names, example: \&quot;name1;name2;name3\&quot;, connected[]: boolean to check if agent belongs to connection, example: \&quot;true\&quot;, | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentIndex**
> ModelMap PlatformAgentIndex(ctx, optional)


Retrieves agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlatformApiPlatformAgentIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformAgentIndexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | [**optional.Interface of int32**](.md)|  | 
 **take** | [**optional.Interface of int32**](.md)|  | 
 **order** | [**optional.Interface of string**](.md)| string: \&quot;ASC\&quot; | \&quot;DESC\&quot; | 
 **filter** | [**optional.Interface of string**](.md)| ids[]: array of agent ids, example: \&quot;1;2;3\&quot;, id|name: agent id or agent name or ip, example: \&quot;name1\&quot; or \&quot;132\&quot; or \&quot;192.168.0.1\&quot;, name: exact agent name, example: \&quot;name1\&quot;, types[]: array of agent types, example: Windows;macOS;Linux;Virtual, statuses[]: one of: connected, connected_with_errors, disconnected, example: \&quot;connected;connected_with_errors\&quot;, networks[]: array of network ids, example: \&quot;0;1;2;3\&quot;, providers[]: array of providers ids, example: \&quot;0;1;2;3\&quot;, tags[]: array of tags ids, example: \&quot;0;1;2;3\&quot;, tags_names[]: array of tags name, example: \&quot;name1;name2;name3\&quot;, networks_names[]: array of networks names, example: \&quot;name1;name2;name3\&quot;, agent_modified_at_from: date from which agent was last modified agent_modified_at_to: date to which agent was last modified agent_versions[]: array of agent versions, example: \&quot;0.0.75;0.0.74\&quot; locations[]: array of locations, example: \&quot;ES;US\&quot; | 
 **loadRelations** | **optional.Bool**|  | [default to true]
 **showLogsState** | **optional.Bool**|  | [default to false]

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentProviderIndex**
> ModelMap PlatformAgentProviderIndex(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlatformApiPlatformAgentProviderIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformAgentProviderIndexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | [**optional.Interface of int32**](.md)|  | 
 **take** | [**optional.Interface of int32**](.md)|  | 
 **order** | [**optional.Interface of string**](.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentProviderShow**
> ModelMap PlatformAgentProviderShow(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**int32**](.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentServiceDestroy**
> PlatformAgentServiceDestroy(ctx, body)


Deletes agent Services

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

# **PlatformAgentServiceIndex**
> ModelMap PlatformAgentServiceIndex(ctx, agentIds)


Retrieves agent services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentIds** | [**[]float64**](float64.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentServiceSubnetUpdate**
> PlatformAgentServiceSubnetUpdate(ctx, body)


Updates agent services status

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

# **PlatformAgentTagIndex**
> ModelMap PlatformAgentTagIndex(ctx, )


Get user agent tags.

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

# **PlatformAgentUpdate**
> ModelMap PlatformAgentUpdate(ctx, body, agentId)


Patches agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)|  | 
  **agentId** | [**int32**](.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformAgentsDestroy**
> InlineResponse204 PlatformAgentsDestroy(ctx, body)


Deletes `platform agent` ands its `connections`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ModelMap**](map.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformApiKeyCreate**
> ModelMap PlatformApiKeyCreate(ctx, body)


Creates API key.

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

# **PlatformApiKeyDestroy**
> InlineResponse204 PlatformApiKeyDestroy(ctx, apiKeyId)


Deletes API key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKeyId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformApiKeyIndex**
> ModelMap PlatformApiKeyIndex(ctx, optional)


Get API keys.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlatformApiPlatformApiKeyIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformApiKeyIndexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | [**optional.Interface of int32**](.md)|  | 
 **take** | [**optional.Interface of int32**](.md)|  | 
 **order** | **optional.String**| string: \&quot;ASC\&quot; | \&quot;DESC\&quot; | 
 **filter** | **optional.String**| api_key_id: string, api_key_name: string | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformConnectionAgentDestroy**
> InlineResponse204 PlatformConnectionAgentDestroy(ctx, agentId)


Deletes agent `connections`. Does not remove `platform agent` from `networks`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformConnectionAgentsDestroy**
> InlineResponse204 PlatformConnectionAgentsDestroy(ctx, body)


Deletes agent `connections`. Does not remove `platform agent` from `networks`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ModelMap**](map.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformConnectionCreate**
> ModelMap PlatformConnectionCreate(ctx, body, optional)


Creates agents connections. If connection type is POINT_TO_POINT, then agent_ids should contain pairs of ids, i.e.: [[1,2], [2,3], ...]. In other types of networks a flat array of ids should be passed, i.e.: [1, 2, 3, ...].

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)|  | 
 **optional** | ***PlatformApiPlatformConnectionCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformConnectionCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paths** | [**optional.Interface of []string**](string.md)| Comma separated servers ids list for SDN path. | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformConnectionCreateMesh**
> ModelMap PlatformConnectionCreateMesh(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)|  | 
 **optional** | ***PlatformApiPlatformConnectionCreateMeshOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformConnectionCreateMeshOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paths** | [**optional.Interface of []string**](string.md)| Comma separated servers ids list for SDN path. | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformConnectionCreateP2p**
> ModelMap PlatformConnectionCreateP2p(ctx, body, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)|  | 
 **optional** | ***PlatformApiPlatformConnectionCreateP2pOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformConnectionCreateP2pOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paths** | [**optional.Interface of []string**](string.md)| Comma separated servers ids list for SDN path. | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformConnectionDestroy**
> InlineResponse204 PlatformConnectionDestroy(ctx, body, optional)


Deletes `connections` by supplied pairs of `platform agents`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ModelMap**](map.md)|  | 
 **optional** | ***PlatformApiPlatformConnectionDestroyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformConnectionDestroyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkUpdatedBy** | [**optional.Interface of NetworkGenesisType**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformConnectionDestroyDeprecated**
> InlineResponse204 PlatformConnectionDestroyDeprecated(ctx, connectionId, optional)


Removes agent pair (agent1, agent2) connections (PUBLIC, SDN{1,2,3}).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**int32**](.md)|  | 
 **optional** | ***PlatformApiPlatformConnectionDestroyDeprecatedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformConnectionDestroyDeprecatedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkUpdatedBy** | [**optional.Interface of NetworkGenesisType**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformConnectionIndex**
> ModelMap PlatformConnectionIndex(ctx, optional)


Retrieves connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlatformApiPlatformConnectionIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformConnectionIndexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | [**optional.Interface of int32**](.md)|  | 
 **take** | [**optional.Interface of int32**](.md)|  | 
 **order** | [**optional.Interface of string**](.md)| string: \&quot;ASC\&quot; | \&quot;DESC\&quot; | 
 **filter** | [**optional.Interface of string**](.md)| id|name: string, example: \&quot;1\&quot; or \&quot;name\&quot;, connectionId: string, example: \&quot;1\&quot;, name: exact agent name, example: \&quot;name1\&quot;, agent_ids[]: array of agent ids, example: \&quot;1;2;3\&quot;, statuses[]: array of statuses, one of PENDING, WARNING, ERROR, CONNECTED, OFFLINE, example: \&quot;OFFLINE;ERROR;WARNING\&quot;, networks[]: array of networks ids, example: \&quot;1;2;3\&quot;, providers[]: array of providers ids, example: \&quot;1;2;3\&quot;, agent_connection_updated_at_from: date from which connection was last modified, agent_connection_updated_at_to: date to which connection was last modified | 
 **showSdnConnections** | [**optional.Interface of ShowSdnConnections**](.md)|  | 
 **loadRelations** | **optional.Bool**|  | [default to true]

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformConnectionServiceShow**
> ModelMap PlatformConnectionServiceShow(ctx, connectionIds)


Retrieves connection services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionIds** | [**[]float64**](float64.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformConnectionServiceUpdate**
> PlatformConnectionServiceUpdate(ctx, body)


Updates agent connection services ips

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

# **PlatformConnectionSubnetDestroy**
> PlatformConnectionSubnetDestroy(ctx, body)


Deletes agent connection services/subnets.

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

# **PlatformLogsReadTimestamp**
> InlineResponse204 PlatformLogsReadTimestamp(ctx, body)


Save last logs read timestamp.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ModelMap**](map.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkAgentCreate**
> InlineResponse204 PlatformNetworkAgentCreate(ctx, body, networkId)


Adds agents to network without modifying connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ModelMap**](map.md)|  | 
  **networkId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkAgentCreateDeprecated**
> InlineResponse204 PlatformNetworkAgentCreateDeprecated(ctx, body, networkId)


Adds `platform agents` to `network` view. Does not modify `connections`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]ModelMap**](map.md)|  | 
  **networkId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkAgentDestroy**
> InlineResponse204 PlatformNetworkAgentDestroy(ctx, networkId, agentId)


Removes agent from network view and connections associated with it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**int32**](.md)|  | 
  **agentId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkAgentGroupCreate**
> InlineResponse204 PlatformNetworkAgentGroupCreate(ctx, networkId, groupName, optional)


Creates agent group relation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**int32**](.md)|  | 
  **groupName** | [**string**](.md)|  | 
 **optional** | ***PlatformApiPlatformNetworkAgentGroupCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformNetworkAgentGroupCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []float64**](float64.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkAgentRemove**
> InlineResponse204 PlatformNetworkAgentRemove(ctx, body, networkId)


Remove agents from network view without unconfiguring connections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]float64**](float64.md)|  | 
  **networkId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkAgentRemoveDeprecated**
> InlineResponse204 PlatformNetworkAgentRemoveDeprecated(ctx, networkId)


Remove `platform agents` from `network` view. Does not modify `connections`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkCreate**
> ModelMap PlatformNetworkCreate(ctx, body)


Creates network.

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

# **PlatformNetworkDestroy**
> InlineResponse204 PlatformNetworkDestroy(ctx, networkId)


Deletes `network`. Does not modify `connections`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkIndex**
> ModelMap PlatformNetworkIndex(ctx, optional)


Retrieves networks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlatformApiPlatformNetworkIndexOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformNetworkIndexOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | [**optional.Interface of int32**](.md)|  | 
 **take** | [**optional.Interface of int32**](.md)|  | 
 **order** | [**optional.Interface of string**](.md)| string: \&quot;ASC\&quot; | \&quot;DESC\&quot; | 
 **filter** | [**optional.Interface of string**](.md)| id|name: string, name: string, | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkInfo**
> ModelMap PlatformNetworkInfo(ctx, networkId)


Get network info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | [**int32**](.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkNetworkAgentDestroyDeprecated**
> InlineResponse204 PlatformNetworkNetworkAgentDestroyDeprecated(ctx, body, networkId)


Removes network agents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]float64**](float64.md)|  | 
  **networkId** | [**int32**](.md)|  | 

### Return type

[**InlineResponse204**](inline_response_204.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlatformNetworkTopology**
> ModelMap PlatformNetworkTopology(ctx, optional)


Retrieves networks topology.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlatformApiPlatformNetworkTopologyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlatformApiPlatformNetworkTopologyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkId** | [**optional.Interface of int32**](.md)|  | 

### Return type

[**ModelMap**](map.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

