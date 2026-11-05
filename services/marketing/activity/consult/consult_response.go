// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package consult

// ActivityConsultResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityConsultResponse struct {
	// ConsultResultInfoList 是咨询后的活动结果信息
	ConsultResultInfoList *[]ConsultActivityResultInfo `json:"consult_result_info_list,omitempty"`
	// OpenId 是领券的用户openId
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是领券的用户uid
	UserId *string `json:"user_id,omitempty"`
}
