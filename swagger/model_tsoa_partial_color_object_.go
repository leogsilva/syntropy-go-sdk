/*
 * syntropy-controller
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TsoaPartialColorObject_ struct {
	ColorId float64 `json:"color_id,omitempty"`
	ColorBandwidth *ConstraintEnum `json:"color_bandwidth,omitempty"`
	ColorJitter *ConstraintEnum `json:"color_jitter,omitempty"`
	ColorLatency *ConstraintEnum `json:"color_latency,omitempty"`
	ColorPrice *ConstraintEnum `json:"color_price,omitempty"`
}