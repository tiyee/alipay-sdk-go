// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package consult

// ConsultActivityInfo 是支付宝官方 V3 OpenAPI 模型。
type ConsultActivityInfo struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
}

// ConsultActivityResultInfo 是支付宝官方 V3 OpenAPI 模型。
type ConsultActivityResultInfo struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// ConsultResultCode 是咨询结果码
	ConsultResultCode *string `json:"consult_result_code,omitempty"`
}
