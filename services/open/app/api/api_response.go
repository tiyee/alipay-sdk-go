package api

// AppAPIFieldApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type AppAPIFieldApplyResponse struct {
}

// AppAPIFieldQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AppAPIFieldQueryResponse struct {
	AuthFieldResponse *AuthFieldResponse `json:"auth_field_response,omitempty"`
}

// AppAPIQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AppAPIQueryResponse struct {
	// Apis 是应用可申请的接口出参敏感字段列表。
	Apis *[]AuthAPIDTO `json:"apis,omitempty"`
}

// AppAPISceneQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AppAPISceneQueryResponse struct {
	// AuthFieldScene 是接口信息字段应用场景。
	AuthFieldScene *[]AuthFieldSceneDTO `json:"auth_field_scene,omitempty"`
}

// AuthFieldResponse 是支付宝官方 V3 OpenAPI 模型。
type AuthFieldResponse struct {
	// Records 是用户信息申请记录列表。
	Records *[]AuthFieldDTO `json:"records,omitempty"`
}
