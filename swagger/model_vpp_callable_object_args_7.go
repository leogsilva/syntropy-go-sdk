/*
 * syntropy-controller
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VppCallableObjectArgs7 struct {
	Endpoints []string `json:"endpoints"`
	Port float64 `json:"port"`
	Addr string `json:"addr,omitempty"`
	Ifname string `json:"ifname"`
}
