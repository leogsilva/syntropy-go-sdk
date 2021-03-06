/*
 * syntropy-controller
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TsoaPartialRouteObject_ struct {
	ServerVinId float64 `json:"server_vin_id,omitempty"`
	ServerVenId float64 `json:"server_ven_id,omitempty"`
	RouteBandwidth float64 `json:"route_bandwidth,omitempty"`
	RoutePrice float64 `json:"route_price,omitempty"`
	RouteJitter float64 `json:"route_jitter,omitempty"`
	RouteLatency float64 `json:"route_latency,omitempty"`
	RouteStatus *Status `json:"route_status,omitempty"`
	RouteStatusReason string `json:"route_status_reason,omitempty"`
}
