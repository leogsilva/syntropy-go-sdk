/*
 * syntropy-controller
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ChangePathObjectDataCosts struct {
	Price float64 `json:"price,omitempty"`
	Latency float64 `json:"latency,omitempty"`
	Jitter float64 `json:"jitter,omitempty"`
	Bandwidth float64 `json:"bandwidth,omitempty"`
}