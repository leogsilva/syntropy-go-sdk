# Go API client for swagger

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AgentConnectionsApi* | [**AgentConnectionCreate**](docs/AgentConnectionsApi.md#agentconnectioncreate) | **Post** /api/agent-connections | 
*AgentConnectionsApi* | [**AgentConnectionDestroy**](docs/AgentConnectionsApi.md#agentconnectiondestroy) | **Delete** /api/agent-connections/{id} | 
*AgentConnectionsApi* | [**AgentConnectionIndex**](docs/AgentConnectionsApi.md#agentconnectionindex) | **Get** /api/agent-connections | 
*AgentConnectionsApi* | [**AgentConnectionRoutingInfo**](docs/AgentConnectionsApi.md#agentconnectionroutinginfo) | **Get** /api/agent-connections/routing-info | 
*AgentConnectionsApi* | [**AgentConnectionShow**](docs/AgentConnectionsApi.md#agentconnectionshow) | **Get** /api/agent-connections/{id} | 
*AgentConnectionsApi* | [**AgentConnectionUpdate**](docs/AgentConnectionsApi.md#agentconnectionupdate) | **Patch** /api/agent-connections/{id} | 
*AgentInterfaceBwApi* | [**AgentInterfaceBwCreate**](docs/AgentInterfaceBwApi.md#agentinterfacebwcreate) | **Post** /api/agent-interface-bw | 
*AgentInterfaceBwApi* | [**AgentInterfaceBwDestroy**](docs/AgentInterfaceBwApi.md#agentinterfacebwdestroy) | **Delete** /api/agent-interface-bw/{id} | 
*AgentInterfaceBwApi* | [**AgentInterfaceBwIndex**](docs/AgentInterfaceBwApi.md#agentinterfacebwindex) | **Get** /api/agent-interface-bw | 
*AgentInterfaceBwApi* | [**AgentInterfaceBwShow**](docs/AgentInterfaceBwApi.md#agentinterfacebwshow) | **Get** /api/agent-interface-bw/{id} | 
*AgentInterfaceBwApi* | [**AgentInterfaceBwUpdate**](docs/AgentInterfaceBwApi.md#agentinterfacebwupdate) | **Patch** /api/agent-interface-bw/{id} | 
*AgentInterfacesApi* | [**AgentInterfaceCreate**](docs/AgentInterfacesApi.md#agentinterfacecreate) | **Post** /api/agent-interfaces | 
*AgentInterfacesApi* | [**AgentInterfaceDestroy**](docs/AgentInterfacesApi.md#agentinterfacedestroy) | **Delete** /api/agent-interfaces/{id} | 
*AgentInterfacesApi* | [**AgentInterfaceIndex**](docs/AgentInterfacesApi.md#agentinterfaceindex) | **Get** /api/agent-interfaces | 
*AgentInterfacesApi* | [**AgentInterfaceShow**](docs/AgentInterfacesApi.md#agentinterfaceshow) | **Get** /api/agent-interfaces/{id} | 
*AgentInterfacesApi* | [**AgentInterfaceUpdate**](docs/AgentInterfacesApi.md#agentinterfaceupdate) | **Patch** /api/agent-interfaces/{id} | 
*AgentNetworksApi* | [**AgentNetworkCreate**](docs/AgentNetworksApi.md#agentnetworkcreate) | **Post** /api/agent-networks | 
*AgentNetworksApi* | [**AgentNetworkDestroy**](docs/AgentNetworksApi.md#agentnetworkdestroy) | **Delete** /api/agent-networks/{id} | 
*AgentNetworksApi* | [**AgentNetworkIndex**](docs/AgentNetworksApi.md#agentnetworkindex) | **Get** /api/agent-networks | 
*AgentNetworksApi* | [**AgentNetworkShow**](docs/AgentNetworksApi.md#agentnetworkshow) | **Get** /api/agent-networks/{id} | 
*AgentNetworksApi* | [**AgentNetworkUpdate**](docs/AgentNetworksApi.md#agentnetworkupdate) | **Patch** /api/agent-networks/{id} | 
*AgentPathsApi* | [**AgentPathCreate**](docs/AgentPathsApi.md#agentpathcreate) | **Post** /api/agent-paths | 
*AgentPathsApi* | [**AgentPathDestroy**](docs/AgentPathsApi.md#agentpathdestroy) | **Delete** /api/agent-paths/{id} | 
*AgentPathsApi* | [**AgentPathIndex**](docs/AgentPathsApi.md#agentpathindex) | **Get** /api/agent-paths | 
*AgentPathsApi* | [**AgentPathShow**](docs/AgentPathsApi.md#agentpathshow) | **Get** /api/agent-paths/{id} | 
*AgentPathsApi* | [**AgentPathUpdate**](docs/AgentPathsApi.md#agentpathupdate) | **Patch** /api/agent-paths/{id} | 
*AgentProvidersApi* | [**AgentProviderCreate**](docs/AgentProvidersApi.md#agentprovidercreate) | **Post** /api/agent-providers | 
*AgentProvidersApi* | [**AgentProviderDestroy**](docs/AgentProvidersApi.md#agentproviderdestroy) | **Delete** /api/agent-providers/{id} | 
*AgentProvidersApi* | [**AgentProviderIndex**](docs/AgentProvidersApi.md#agentproviderindex) | **Get** /api/agent-providers | 
*AgentProvidersApi* | [**AgentProviderShow**](docs/AgentProvidersApi.md#agentprovidershow) | **Get** /api/agent-providers/{id} | 
*AgentProvidersApi* | [**AgentProviderUpdate**](docs/AgentProvidersApi.md#agentproviderupdate) | **Patch** /api/agent-providers/{id} | 
*AgentsApi* | [**AgentConfig**](docs/AgentsApi.md#agentconfig) | **Get** /api/agents/{key}/config | 
*AgentsApi* | [**AgentCreate**](docs/AgentsApi.md#agentcreate) | **Post** /api/agents | 
*AgentsApi* | [**AgentDestroy**](docs/AgentsApi.md#agentdestroy) | **Delete** /api/agents/{id} | 
*AgentsApi* | [**AgentIndex**](docs/AgentsApi.md#agentindex) | **Get** /api/agents | 
*AgentsApi* | [**AgentServerInfo**](docs/AgentsApi.md#agentserverinfo) | **Get** /api/agents/{key}/server-info | 
*AgentsApi* | [**AgentShow**](docs/AgentsApi.md#agentshow) | **Get** /api/agents/{id} | 
*AgentsApi* | [**AgentUpdate**](docs/AgentsApi.md#agentupdate) | **Patch** /api/agents/{id} | 
*AgentsApi* | [**FindByApiKey**](docs/AgentsApi.md#findbyapikey) | **Get** /api/agents/api-key/{api_key_id} | 
*ApiKeysApi* | [**ApiKeyCreate**](docs/ApiKeysApi.md#apikeycreate) | **Post** /api/api-keys | 
*ApiKeysApi* | [**ApiKeyDestroy**](docs/ApiKeysApi.md#apikeydestroy) | **Delete** /api/api-keys/{id} | 
*ApiKeysApi* | [**ApiKeyIndex**](docs/ApiKeysApi.md#apikeyindex) | **Get** /api/api-keys | 
*ApiKeysApi* | [**ApiKeyShow**](docs/ApiKeysApi.md#apikeyshow) | **Get** /api/api-keys/{id} | 
*ApiKeysApi* | [**ApiKeyUpdate**](docs/ApiKeysApi.md#apikeyupdate) | **Patch** /api/api-keys/{id} | 
*ColorsApi* | [**ColorCreate**](docs/ColorsApi.md#colorcreate) | **Post** /api/colors | 
*ColorsApi* | [**ColorDestroy**](docs/ColorsApi.md#colordestroy) | **Delete** /api/colors/{id} | 
*ColorsApi* | [**ColorIndex**](docs/ColorsApi.md#colorindex) | **Get** /api/colors | 
*ColorsApi* | [**ColorShow**](docs/ColorsApi.md#colorshow) | **Get** /api/colors/{id} | 
*ColorsApi* | [**ColorSync**](docs/ColorsApi.md#colorsync) | **Post** /api/colors/sync | 
*ColorsApi* | [**ColorUpdate**](docs/ColorsApi.md#colorupdate) | **Patch** /api/colors/{id} | 
*ContentApi* | [**ContentCreate**](docs/ContentApi.md#contentcreate) | **Post** /api/content | 
*ContentApi* | [**ContentDestroy**](docs/ContentApi.md#contentdestroy) | **Delete** /api/content/{id} | 
*ContentApi* | [**ContentIndex**](docs/ContentApi.md#contentindex) | **Get** /api/content | 
*ContentApi* | [**ContentShow**](docs/ContentApi.md#contentshow) | **Get** /api/content/{id} | 
*ContentApi* | [**ContentUpdate**](docs/ContentApi.md#contentupdate) | **Patch** /api/content/{id} | 
*InterfacesApi* | [**InterfaceCreate**](docs/InterfacesApi.md#interfacecreate) | **Post** /api/interfaces | 
*InterfacesApi* | [**InterfaceDestroy**](docs/InterfacesApi.md#interfacedestroy) | **Delete** /api/interfaces/{id} | 
*InterfacesApi* | [**InterfaceIndex**](docs/InterfacesApi.md#interfaceindex) | **Get** /api/interfaces | 
*InterfacesApi* | [**InterfaceShow**](docs/InterfacesApi.md#interfaceshow) | **Get** /api/interfaces/{id} | 
*InterfacesApi* | [**InterfaceUpdate**](docs/InterfacesApi.md#interfaceupdate) | **Patch** /api/interfaces/{id} | 
*InternalApi* | [**InternalAgentRestart**](docs/InternalApi.md#internalagentrestart) | **Post** /api/internal/agent-restart | 
*InternalApi* | [**InternalBiStats**](docs/InternalApi.md#internalbistats) | **Get** /api/internal/bi-stats | 
*InternalApi* | [**InternalMaintenance**](docs/InternalApi.md#internalmaintenance) | **Get** /api/internal/maintenance | 
*InternalApi* | [**InternalServerRrestart**](docs/InternalApi.md#internalserverrrestart) | **Post** /api/internal/server-restart | 
*InternalApi* | [**InternalServers**](docs/InternalApi.md#internalservers) | **Delete** /api/internal/servers | 
*InternalApi* | [**InternalTextToQrcode**](docs/InternalApi.md#internaltexttoqrcode) | **Post** /api/internal/text-to-qrcode | 
*InternalApi* | [**InternalUsersData**](docs/InternalApi.md#internalusersdata) | **Delete** /api/internal/users-data | 
*InternalApi* | [**InternalWgKeygen**](docs/InternalApi.md#internalwgkeygen) | **Get** /api/internal/wg-keygen | 
*LanguagesApi* | [**LanguageCreate**](docs/LanguagesApi.md#languagecreate) | **Post** /api/languages | 
*LanguagesApi* | [**LanguageDestroy**](docs/LanguagesApi.md#languagedestroy) | **Delete** /api/languages/{id} | 
*LanguagesApi* | [**LanguageIndex**](docs/LanguagesApi.md#languageindex) | **Get** /api/languages | 
*LanguagesApi* | [**LanguageShow**](docs/LanguagesApi.md#languageshow) | **Get** /api/languages/{id} | 
*LanguagesApi* | [**LanguageUpdate**](docs/LanguagesApi.md#languageupdate) | **Patch** /api/languages/{id} | 
*LinksApi* | [**LinkCreate**](docs/LinksApi.md#linkcreate) | **Post** /api/links | 
*LinksApi* | [**LinkDestroy**](docs/LinksApi.md#linkdestroy) | **Delete** /api/links/{id} | 
*LinksApi* | [**LinkIndex**](docs/LinksApi.md#linkindex) | **Get** /api/links | 
*LinksApi* | [**LinkShow**](docs/LinksApi.md#linkshow) | **Get** /api/links/{id} | 
*LinksApi* | [**LinkUpdate**](docs/LinksApi.md#linkupdate) | **Patch** /api/links/{id} | 
*NetworksApi* | [**NetworkCreate**](docs/NetworksApi.md#networkcreate) | **Post** /api/networks | 
*NetworksApi* | [**NetworkDestroy**](docs/NetworksApi.md#networkdestroy) | **Delete** /api/networks/{id} | 
*NetworksApi* | [**NetworkIndex**](docs/NetworksApi.md#networkindex) | **Get** /api/networks | 
*NetworksApi* | [**NetworkShow**](docs/NetworksApi.md#networkshow) | **Get** /api/networks/{id} | 
*NetworksApi* | [**NetworkUpdate**](docs/NetworksApi.md#networkupdate) | **Patch** /api/networks/{id} | 
*PlatformApi* | [**PlatformAdminAgentConfig**](docs/PlatformApi.md#platformadminagentconfig) | **Get** /api/platform/admin/agent/{agent_id}/config | 
*PlatformApi* | [**PlatformAgentConfig**](docs/PlatformApi.md#platformagentconfig) | **Get** /api/platform/agent/{agent_id}/config | 
*PlatformApi* | [**PlatformAgentCoordinates**](docs/PlatformApi.md#platformagentcoordinates) | **Post** /api/platform/agents/coordinates | 
*PlatformApi* | [**PlatformAgentCreate**](docs/PlatformApi.md#platformagentcreate) | **Post** /api/platform/agents | 
*PlatformApi* | [**PlatformAgentDestroy**](docs/PlatformApi.md#platformagentdestroy) | **Delete** /api/platform/agents/{agent_id} | 
*PlatformApi* | [**PlatformAgentGroupDestroy**](docs/PlatformApi.md#platformagentgroupdestroy) | **Delete** /api/platform/network/agent-groups/{group_id} | 
*PlatformApi* | [**PlatformAgentGroupUpdate**](docs/PlatformApi.md#platformagentgroupupdate) | **Put** /api/platform/network/agent-groups/{group_id} | 
*PlatformApi* | [**PlatformAgentIdNamePairs**](docs/PlatformApi.md#platformagentidnamepairs) | **Get** /api/platform/agents/filters | 
*PlatformApi* | [**PlatformAgentIndex**](docs/PlatformApi.md#platformagentindex) | **Get** /api/platform/agents | 
*PlatformApi* | [**PlatformAgentProviderIndex**](docs/PlatformApi.md#platformagentproviderindex) | **Get** /api/platform/agent-providers | 
*PlatformApi* | [**PlatformAgentProviderShow**](docs/PlatformApi.md#platformagentprovidershow) | **Get** /api/platform/agent-providers/{id} | 
*PlatformApi* | [**PlatformAgentServiceDestroy**](docs/PlatformApi.md#platformagentservicedestroy) | **Post** /api/platform/agent-services-delete | 
*PlatformApi* | [**PlatformAgentServiceIndex**](docs/PlatformApi.md#platformagentserviceindex) | **Get** /api/platform/agent-services | 
*PlatformApi* | [**PlatformAgentServiceSubnetUpdate**](docs/PlatformApi.md#platformagentservicesubnetupdate) | **Post** /api/platform/agent-services-subnets | 
*PlatformApi* | [**PlatformAgentTagIndex**](docs/PlatformApi.md#platformagenttagindex) | **Get** /api/platform/agent-tags | 
*PlatformApi* | [**PlatformAgentUpdate**](docs/PlatformApi.md#platformagentupdate) | **Patch** /api/platform/agents/{agent_id} | 
*PlatformApi* | [**PlatformAgentsDestroy**](docs/PlatformApi.md#platformagentsdestroy) | **Post** /api/platform/agents/remove | 
*PlatformApi* | [**PlatformApiKeyCreate**](docs/PlatformApi.md#platformapikeycreate) | **Post** /api/platform/api-keys | 
*PlatformApi* | [**PlatformApiKeyDestroy**](docs/PlatformApi.md#platformapikeydestroy) | **Delete** /api/platform/api-keys/{api_key_id} | 
*PlatformApi* | [**PlatformApiKeyIndex**](docs/PlatformApi.md#platformapikeyindex) | **Get** /api/platform/api-keys | 
*PlatformApi* | [**PlatformConnectionAgentDestroy**](docs/PlatformApi.md#platformconnectionagentdestroy) | **Delete** /api/platform/connections/agents/{agent_id} | 
*PlatformApi* | [**PlatformConnectionAgentsDestroy**](docs/PlatformApi.md#platformconnectionagentsdestroy) | **Post** /api/platform/connections/agents/remove | 
*PlatformApi* | [**PlatformConnectionCreate**](docs/PlatformApi.md#platformconnectioncreate) | **Post** /api/platform/connections | 
*PlatformApi* | [**PlatformConnectionCreateMesh**](docs/PlatformApi.md#platformconnectioncreatemesh) | **Post** /api/platform/connections/mesh | 
*PlatformApi* | [**PlatformConnectionCreateP2p**](docs/PlatformApi.md#platformconnectioncreatep2p) | **Post** /api/platform/connections/point-to-point | 
*PlatformApi* | [**PlatformConnectionDestroy**](docs/PlatformApi.md#platformconnectiondestroy) | **Post** /api/platform/connections/remove | 
*PlatformApi* | [**PlatformConnectionDestroyDeprecated**](docs/PlatformApi.md#platformconnectiondestroydeprecated) | **Delete** /api/platform/connections/{connection_id} | 
*PlatformApi* | [**PlatformConnectionIndex**](docs/PlatformApi.md#platformconnectionindex) | **Get** /api/platform/connections | 
*PlatformApi* | [**PlatformConnectionServiceShow**](docs/PlatformApi.md#platformconnectionserviceshow) | **Get** /api/platform/connection-services | 
*PlatformApi* | [**PlatformConnectionServiceUpdate**](docs/PlatformApi.md#platformconnectionserviceupdate) | **Post** /api/platform/connection-services | 
*PlatformApi* | [**PlatformConnectionSubnetDestroy**](docs/PlatformApi.md#platformconnectionsubnetdestroy) | **Post** /api/platform/connection-services-delete | 
*PlatformApi* | [**PlatformLogsReadTimestamp**](docs/PlatformApi.md#platformlogsreadtimestamp) | **Post** /api/platform/logs-reads-timestamp | 
*PlatformApi* | [**PlatformNetworkAgentCreate**](docs/PlatformApi.md#platformnetworkagentcreate) | **Post** /api/platform/network/{network_id}/agents/add | 
*PlatformApi* | [**PlatformNetworkAgentCreateDeprecated**](docs/PlatformApi.md#platformnetworkagentcreatedeprecated) | **Post** /api/platform/network/{network_id}/agents | 
*PlatformApi* | [**PlatformNetworkAgentDestroy**](docs/PlatformApi.md#platformnetworkagentdestroy) | **Delete** /api/platform/networks/{network_id}/agents/{agent_id} | 
*PlatformApi* | [**PlatformNetworkAgentGroupCreate**](docs/PlatformApi.md#platformnetworkagentgroupcreate) | **Post** /api/platform/network/{network_id}/agent-groups/{group_name} | 
*PlatformApi* | [**PlatformNetworkAgentRemove**](docs/PlatformApi.md#platformnetworkagentremove) | **Post** /api/platform/networks/{network_id}/agents/remove | 
*PlatformApi* | [**PlatformNetworkAgentRemoveDeprecated**](docs/PlatformApi.md#platformnetworkagentremovedeprecated) | **Delete** /api/platform/networks/{network_id}/agents | 
*PlatformApi* | [**PlatformNetworkCreate**](docs/PlatformApi.md#platformnetworkcreate) | **Post** /api/platform/networks | 
*PlatformApi* | [**PlatformNetworkDestroy**](docs/PlatformApi.md#platformnetworkdestroy) | **Delete** /api/platform/networks/{network_id} | 
*PlatformApi* | [**PlatformNetworkIndex**](docs/PlatformApi.md#platformnetworkindex) | **Get** /api/platform/networks | 
*PlatformApi* | [**PlatformNetworkInfo**](docs/PlatformApi.md#platformnetworkinfo) | **Get** /api/platform/network/{network_id}/info | 
*PlatformApi* | [**PlatformNetworkNetworkAgentDestroyDeprecated**](docs/PlatformApi.md#platformnetworknetworkagentdestroydeprecated) | **Post** /api/platform/network/{network_id}/agents/delete | 
*PlatformApi* | [**PlatformNetworkTopology**](docs/PlatformApi.md#platformnetworktopology) | **Get** /api/platform/networks/topology | 
*ProvidersApi* | [**ProviderCreate**](docs/ProvidersApi.md#providercreate) | **Post** /api/providers | 
*ProvidersApi* | [**ProviderDestroy**](docs/ProvidersApi.md#providerdestroy) | **Delete** /api/providers/{id} | 
*ProvidersApi* | [**ProviderIndex**](docs/ProvidersApi.md#providerindex) | **Get** /api/providers | 
*ProvidersApi* | [**ProviderShow**](docs/ProvidersApi.md#providershow) | **Get** /api/providers/{id} | 
*ProvidersApi* | [**ProviderUpdate**](docs/ProvidersApi.md#providerupdate) | **Patch** /api/providers/{id} | 
*ProxyApi* | [**ProxyGetWgInternalIpv4**](docs/ProxyApi.md#proxygetwginternalipv4) | **Get** /api/proxy/platform-agents/get-wg-internal-ipv4 | 
*ProxyApi* | [**ProxySendNftablesConf**](docs/ProxyApi.md#proxysendnftablesconf) | **Post** /api/proxy/sdn-agents/send-nftables-conf | 
*ProxyApi* | [**ProxySendVppConf**](docs/ProxyApi.md#proxysendvppconf) | **Post** /api/proxy/sdn-agents/send-vpp-conf | 
*PublicApi* | [**PublicHealth**](docs/PublicApi.md#publichealth) | **Get** /api/public/health | 
*PublicApi* | [**PublicInfo**](docs/PublicApi.md#publicinfo) | **Get** /api/public/info | 
*PublicApi* | [**PublicLanguageIndex**](docs/PublicApi.md#publiclanguageindex) | **Get** /api/public/languages | 
*PublicApi* | [**PublicLinkIndex**](docs/PublicApi.md#publiclinkindex) | **Get** /api/public/links | 
*PublicApi* | [**PublicTranslationIndex**](docs/PublicApi.md#publictranslationindex) | **Get** /api/public/translations | 
*PublicApi* | [**PublicVersion**](docs/PublicApi.md#publicversion) | **Get** /api/public/versions | 
*RegionsApi* | [**RegionCreate**](docs/RegionsApi.md#regioncreate) | **Post** /api/regions | 
*RegionsApi* | [**RegionDestroy**](docs/RegionsApi.md#regiondestroy) | **Delete** /api/regions/{id} | 
*RegionsApi* | [**RegionIndex**](docs/RegionsApi.md#regionindex) | **Get** /api/regions | 
*RegionsApi* | [**RegionShow**](docs/RegionsApi.md#regionshow) | **Get** /api/regions/{id} | 
*RegionsApi* | [**RegionUpdate**](docs/RegionsApi.md#regionupdate) | **Patch** /api/regions/{id} | 
*RoutesApi* | [**RouteCreate**](docs/RoutesApi.md#routecreate) | **Post** /api/routes | 
*RoutesApi* | [**RouteDestroy**](docs/RoutesApi.md#routedestroy) | **Delete** /api/routes/{id} | 
*RoutesApi* | [**RouteIndex**](docs/RoutesApi.md#routeindex) | **Get** /api/routes | 
*RoutesApi* | [**RouteShow**](docs/RoutesApi.md#routeshow) | **Get** /api/routes/{id} | 
*RoutesApi* | [**RouteUpdate**](docs/RoutesApi.md#routeupdate) | **Patch** /api/routes/{id} | 
*S3ServiceApi* | [**S3AgentProviderImageDestroy**](docs/S3ServiceApi.md#s3agentproviderimagedestroy) | **Delete** /api/s3/delete-agent-provider-img/{key} | 
*S3ServiceApi* | [**S3AgentProviderImageUpdate**](docs/S3ServiceApi.md#s3agentproviderimageupdate) | **Post** /api/s3/upload-or-update-agent-provider-img | 
*S3ServiceApi* | [**S3AgentProviderImagesIndex**](docs/S3ServiceApi.md#s3agentproviderimagesindex) | **Get** /api/s3/get-agent-provider-imgs | 
*SearchApi* | [**SearchAgents**](docs/SearchApi.md#searchagents) | **Post** /api/search/agent | 
*SearchApi* | [**SearchController**](docs/SearchApi.md#searchcontroller) | **Post** /api/search/controller | 
*SearchApi* | [**SearchPlatformAgent**](docs/SearchApi.md#searchplatformagent) | **Post** /api/search/platform-agents | 
*SearchApi* | [**SearchPlatformAgentError**](docs/SearchApi.md#searchplatformagenterror) | **Post** /api/search/platform-agents-errors | 
*ServersApi* | [**ServerCreate**](docs/ServersApi.md#servercreate) | **Post** /api/servers | 
*ServersApi* | [**ServerDestroy**](docs/ServersApi.md#serverdestroy) | **Delete** /api/servers/{id} | 
*ServersApi* | [**ServerIndex**](docs/ServersApi.md#serverindex) | **Get** /api/servers | 
*ServersApi* | [**ServerShow**](docs/ServersApi.md#servershow) | **Get** /api/servers/{id} | 
*ServersApi* | [**ServerSpeedtest**](docs/ServersApi.md#serverspeedtest) | **Post** /api/servers/speedtest | 
*ServersApi* | [**ServerUpdate**](docs/ServersApi.md#serverupdate) | **Patch** /api/servers/{id} | 
*SettingsApi* | [**SettingCreate**](docs/SettingsApi.md#settingcreate) | **Post** /api/settings | 
*SettingsApi* | [**SettingIndex**](docs/SettingsApi.md#settingindex) | **Get** /api/settings | 
*SrPathsApi* | [**SrPathIndex**](docs/SrPathsApi.md#srpathindex) | **Get** /api/sr-paths | 
*SrPathsApi* | [**SrPathShow**](docs/SrPathsApi.md#srpathshow) | **Get** /api/sr-paths/{id} | 
*SrPoliciesApi* | [**SrPolicyIndex**](docs/SrPoliciesApi.md#srpolicyindex) | **Get** /api/sr-policies | 
*SrPoliciesApi* | [**SrPolicyShow**](docs/SrPoliciesApi.md#srpolicyshow) | **Get** /api/sr-policies/{id} | 
*TopologiesApi* | [**TopologyCreate**](docs/TopologiesApi.md#topologycreate) | **Post** /api/topologies | 
*TopologiesApi* | [**TopologyDestroy**](docs/TopologiesApi.md#topologydestroy) | **Delete** /api/topologies/{id} | 
*TopologiesApi* | [**TopologyIndex**](docs/TopologiesApi.md#topologyindex) | **Get** /api/topologies | 
*TopologiesApi* | [**TopologyShow**](docs/TopologiesApi.md#topologyshow) | **Get** /api/topologies/{id} | 
*TopologiesApi* | [**TopologyUpdate**](docs/TopologiesApi.md#topologyupdate) | **Patch** /api/topologies/{id} | 
*TranslationsApi* | [**Download**](docs/TranslationsApi.md#download) | **Post** /api/translations/download-synch-file | 
*TranslationsApi* | [**TranslationCreate**](docs/TranslationsApi.md#translationcreate) | **Post** /api/translations | 
*TranslationsApi* | [**TranslationDestroy**](docs/TranslationsApi.md#translationdestroy) | **Delete** /api/translations/{id} | 
*TranslationsApi* | [**TranslationIndex**](docs/TranslationsApi.md#translationindex) | **Get** /api/translations | 
*TranslationsApi* | [**TranslationShow**](docs/TranslationsApi.md#translationshow) | **Get** /api/translations/{id} | 
*TranslationsApi* | [**TranslationUpdate**](docs/TranslationsApi.md#translationupdate) | **Patch** /api/translations/{id} | 
*TranslationsApi* | [**UploadFile**](docs/TranslationsApi.md#uploadfile) | **Post** /api/translations/upload-sync-file | 
*TunnelsApi* | [**TunnelCreate**](docs/TunnelsApi.md#tunnelcreate) | **Post** /api/tunnels | 
*TunnelsApi* | [**TunnelDestroy**](docs/TunnelsApi.md#tunneldestroy) | **Delete** /api/tunnels/{id} | 
*TunnelsApi* | [**TunnelIndex**](docs/TunnelsApi.md#tunnelindex) | **Get** /api/tunnels | 
*TunnelsApi* | [**TunnelShow**](docs/TunnelsApi.md#tunnelshow) | **Get** /api/tunnels/{id} | 
*TunnelsApi* | [**TunnelUpdate**](docs/TunnelsApi.md#tunnelupdate) | **Patch** /api/tunnels/{id} | 
*UserSrsApi* | [**UserSrCreate**](docs/UserSrsApi.md#usersrcreate) | **Post** /api/user-srs | 
*UserSrsApi* | [**UserSrDestroy**](docs/UserSrsApi.md#usersrdestroy) | **Delete** /api/user-srs/{id} | 
*UserSrsApi* | [**UserSrIndex**](docs/UserSrsApi.md#usersrindex) | **Get** /api/user-srs | 
*UserSrsApi* | [**UserSrShow**](docs/UserSrsApi.md#usersrshow) | **Get** /api/user-srs/{id} | 
*UserSrsApi* | [**UserSrUpdate**](docs/UserSrsApi.md#usersrupdate) | **Patch** /api/user-srs/{id} | 
*UsersApi* | [**UserIndex**](docs/UsersApi.md#userindex) | **Get** /api/users | 
*UsersApi* | [**UserShow**](docs/UsersApi.md#usershow) | **Get** /api/users/{id} | 
*VersionsApi* | [**VersionCreate**](docs/VersionsApi.md#versioncreate) | **Post** /api/versions | 
*VersionsApi* | [**VersionDestroy**](docs/VersionsApi.md#versiondestroy) | **Delete** /api/versions/{id} | 
*VersionsApi* | [**VersionIndex**](docs/VersionsApi.md#versionindex) | **Get** /api/versions | 
*VersionsApi* | [**VersionShow**](docs/VersionsApi.md#versionshow) | **Get** /api/versions/{id} | 
*VersionsApi* | [**VersionUpdate**](docs/VersionsApi.md#versionupdate) | **Patch** /api/versions/{id} | 
*VpnsApi* | [**VpnCreate**](docs/VpnsApi.md#vpncreate) | **Post** /api/vpns | 
*VpnsApi* | [**VpnDestroy**](docs/VpnsApi.md#vpndestroy) | **Delete** /api/vpns/{id} | 
*VpnsApi* | [**VpnIndex**](docs/VpnsApi.md#vpnindex) | **Get** /api/vpns | 
*VpnsApi* | [**VpnShow**](docs/VpnsApi.md#vpnshow) | **Get** /api/vpns/{id} | 
*VpnsApi* | [**VpnUpdate**](docs/VpnsApi.md#vpnupdate) | **Patch** /api/vpns/{id} | 
*WebhooksApi* | [**WebhookChangePath**](docs/WebhooksApi.md#webhookchangepath) | **Post** /api/webhooks/change-path/{id} | 

## Documentation For Models

 - [AgentConnectionStatus](docs/AgentConnectionStatus.md)
 - [AgentCoordinatesObject](docs/AgentCoordinatesObject.md)
 - [AgentMessagePayload](docs/AgentMessagePayload.md)
 - [AgentMessageType](docs/AgentMessageType.md)
 - [AgentServiceTypes](docs/AgentServiceTypes.md)
 - [AnyOfAgentMessagePayload](docs/AnyOfAgentMessagePayload.md)
 - [AnyOfPlatformResponseErrorItemValue](docs/AnyOfPlatformResponseErrorItemValue.md)
 - [AnyOfVppCallableObject](docs/AnyOfVppCallableObject.md)
 - [AnyOfWgCallableObject](docs/AnyOfWgCallableObject.md)
 - [AnyOfinlineResponse204](docs/AnyOfinlineResponse204.md)
 - [BehaviorType](docs/BehaviorType.md)
 - [BiStatisticsEdgesPost](docs/BiStatisticsEdgesPost.md)
 - [Body](docs/Body.md)
 - [Body1](docs/Body1.md)
 - [Body2](docs/Body2.md)
 - [Body3](docs/Body3.md)
 - [Body4](docs/Body4.md)
 - [Body5](docs/Body5.md)
 - [ChangePathObjectData](docs/ChangePathObjectData.md)
 - [ChangePathObjectDataCosts](docs/ChangePathObjectDataCosts.md)
 - [ConstraintEnum](docs/ConstraintEnum.md)
 - [ContextType](docs/ContextType.md)
 - [InlineResponse204](docs/InlineResponse204.md)
 - [InterfaceType](docs/InterfaceType.md)
 - [LinkTag](docs/LinkTag.md)
 - [LogsReadTimestampEntityTypes](docs/LogsReadTimestampEntityTypes.md)
 - [MetadataNetworkType](docs/MetadataNetworkType.md)
 - [NetworkGenesisType](docs/NetworkGenesisType.md)
 - [PlatformAgentStatus](docs/PlatformAgentStatus.md)
 - [PlatformAgentTypeLocal](docs/PlatformAgentTypeLocal.md)
 - [PlatformAgentsErrorResponseHits](docs/PlatformAgentsErrorResponseHits.md)
 - [PlatformAgentsHitObject](docs/PlatformAgentsHitObject.md)
 - [PlatformAgentsHitObjectSource](docs/PlatformAgentsHitObjectSource.md)
 - [PlatformResponseErrorItem](docs/PlatformResponseErrorItem.md)
 - [PlatformResponseSuccessbooleanData](docs/PlatformResponseSuccessbooleanData.md)
 - [S3ObjectListItemOwner](docs/S3ObjectListItemOwner.md)
 - [ServerAgentStatus](docs/ServerAgentStatus.md)
 - [ServerSrSoftware](docs/ServerSrSoftware.md)
 - [SettingsTypes](docs/SettingsTypes.md)
 - [ShowSdnConnections](docs/ShowSdnConnections.md)
 - [Status](docs/Status.md)
 - [TsoaPartialAgentConnectionObject_](docs/TsoaPartialAgentConnectionObject_.md)
 - [TsoaPartialAgentInterfaceBwObject_](docs/TsoaPartialAgentInterfaceBwObject_.md)
 - [TsoaPartialAgentInterfaceObject_](docs/TsoaPartialAgentInterfaceObject_.md)
 - [TsoaPartialAgentNetworkObject_](docs/TsoaPartialAgentNetworkObject_.md)
 - [TsoaPartialAgentObject_](docs/TsoaPartialAgentObject_.md)
 - [TsoaPartialAgentPathObject_](docs/TsoaPartialAgentPathObject_.md)
 - [TsoaPartialAgentProviderObject_](docs/TsoaPartialAgentProviderObject_.md)
 - [TsoaPartialApiKeyObject_](docs/TsoaPartialApiKeyObject_.md)
 - [TsoaPartialColorObject_](docs/TsoaPartialColorObject_.md)
 - [TsoaPartialContentObject_](docs/TsoaPartialContentObject_.md)
 - [TsoaPartialInterfaceObject_](docs/TsoaPartialInterfaceObject_.md)
 - [TsoaPartialLanguageObject_](docs/TsoaPartialLanguageObject_.md)
 - [TsoaPartialLinkObject_](docs/TsoaPartialLinkObject_.md)
 - [TsoaPartialNetworkObject_](docs/TsoaPartialNetworkObject_.md)
 - [TsoaPartialProviderObject_](docs/TsoaPartialProviderObject_.md)
 - [TsoaPartialRegionObject_](docs/TsoaPartialRegionObject_.md)
 - [TsoaPartialRouteObject_](docs/TsoaPartialRouteObject_.md)
 - [TsoaPartialServerObject_](docs/TsoaPartialServerObject_.md)
 - [TsoaPartialTopologyObject_](docs/TsoaPartialTopologyObject_.md)
 - [TsoaPartialTranslationObject_](docs/TsoaPartialTranslationObject_.md)
 - [TsoaPartialTunnelObject_](docs/TsoaPartialTunnelObject_.md)
 - [TsoaPartialUserSrObject_](docs/TsoaPartialUserSrObject_.md)
 - [TsoaPartialVersionObject_](docs/TsoaPartialVersionObject_.md)
 - [TsoaPartialVpnObject_](docs/TsoaPartialVpnObject_.md)
 - [TsoaPickAgentAgentIdOrAgentLocationLatOrAgentLocationLon_](docs/TsoaPickAgentAgentIdOrAgentLocationLatOrAgentLocationLon_.md)
 - [TsoaPickAgentAgentLocationCountry_](docs/TsoaPickAgentAgentLocationCountry_.md)
 - [TsoaPickAgentAgentNameOrAgentId_](docs/TsoaPickAgentAgentNameOrAgentId_.md)
 - [TsoaPickAgentAgentVersion_](docs/TsoaPickAgentAgentVersion_.md)
 - [UpdateStatusBodySubnetsToUpdate](docs/UpdateStatusBodySubnetsToUpdate.md)
 - [UserSrDirection](docs/UserSrDirection.md)
 - [VersionType](docs/VersionType.md)
 - [VisibilityType](docs/VisibilityType.md)
 - [VppCallableObject](docs/VppCallableObject.md)
 - [VppCallableObjectArgs](docs/VppCallableObjectArgs.md)
 - [VppCallableObjectArgs1](docs/VppCallableObjectArgs1.md)
 - [VppCallableObjectArgs10](docs/VppCallableObjectArgs10.md)
 - [VppCallableObjectArgs11](docs/VppCallableObjectArgs11.md)
 - [VppCallableObjectArgs12](docs/VppCallableObjectArgs12.md)
 - [VppCallableObjectArgs13](docs/VppCallableObjectArgs13.md)
 - [VppCallableObjectArgs14](docs/VppCallableObjectArgs14.md)
 - [VppCallableObjectArgs15](docs/VppCallableObjectArgs15.md)
 - [VppCallableObjectArgs16](docs/VppCallableObjectArgs16.md)
 - [VppCallableObjectArgs2](docs/VppCallableObjectArgs2.md)
 - [VppCallableObjectArgs3](docs/VppCallableObjectArgs3.md)
 - [VppCallableObjectArgs4](docs/VppCallableObjectArgs4.md)
 - [VppCallableObjectArgs5](docs/VppCallableObjectArgs5.md)
 - [VppCallableObjectArgs6](docs/VppCallableObjectArgs6.md)
 - [VppCallableObjectArgs7](docs/VppCallableObjectArgs7.md)
 - [VppCallableObjectArgs8](docs/VppCallableObjectArgs8.md)
 - [VppCallableObjectArgs9](docs/VppCallableObjectArgs9.md)
 - [WgAddPeerArgs](docs/WgAddPeerArgs.md)
 - [WgAddPeerMetadata](docs/WgAddPeerMetadata.md)
 - [WgAddPeerMetadataAllowedIpsInfo](docs/WgAddPeerMetadataAllowedIpsInfo.md)
 - [WgCallableObject](docs/WgCallableObject.md)
 - [WgCreateInterfaceArgs](docs/WgCreateInterfaceArgs.md)
 - [WgCreateInterfaceMetadata](docs/WgCreateInterfaceMetadata.md)
 - [WgRemoveInterfaceArgs](docs/WgRemoveInterfaceArgs.md)
 - [WgRemovePeerArgs](docs/WgRemovePeerArgs.md)

## Documentation For Authorization

## api_key
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```
## jwt

## Author


