/*
 * syntropy-controller
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TsoaPartialLanguageObject_ struct {
	LangCode string `json:"lang_code,omitempty"`
	LangTitle string `json:"lang_title,omitempty"`
	LangNativeTitle string `json:"lang_native_title,omitempty"`
	LangWeight float64 `json:"lang_weight,omitempty"`
	LangVisibility *VisibilityType `json:"lang_visibility,omitempty"`
}
