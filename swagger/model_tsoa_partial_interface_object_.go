/*
 * syntropy-controller
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TsoaPartialInterfaceObject_ struct {
	InterfaceName string `json:"interface_name,omitempty"`
	InterfaceType *InterfaceType `json:"interface_type,omitempty"`
	ServerId float64 `json:"server_id,omitempty"`
	ServerCidrIpv4 string `json:"server_cidr_ipv4,omitempty"`
	ServerCidrIpv6 string `json:"server_cidr_ipv6,omitempty"`
	InterfaceMonitoringCidr string `json:"interface_monitoring_cidr,omitempty"`
	InterfaceKvmGatewayIpv4 string `json:"interface_kvm_gateway_ipv4,omitempty"`
	InterfaceAdditionalRoutesCidrs string `json:"interface_additional_routes_cidrs,omitempty"`
}
