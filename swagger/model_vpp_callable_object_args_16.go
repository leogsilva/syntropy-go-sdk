/*
 * syntropy-controller
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VppCallableObjectArgs16 struct {
	EndIfname string `json:"end_ifname,omitempty"`
	EndAddr string `json:"end_addr,omitempty"`
	Behavior *BehaviorType `json:"behavior"`
	Addr string `json:"addr"`
}
