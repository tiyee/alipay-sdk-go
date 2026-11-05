// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package consult

// ActivityConsultRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityConsultRequest struct {
	// ConsultActivityInfoList 是待咨询的活动信息列表
	ConsultActivityInfoList *[]ConsultActivityInfo `json:"consult_activity_info_list,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// MerchantId 是商户PID,默认为当前接口调用商户。
	MerchantId *string `json:"merchant_id,omitempty"`
	// OpenId 是领券的支付宝用户openId
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是领券的用户uid
	UserId *string `json:"user_id,omitempty"`
}
