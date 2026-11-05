package certdoc

// UserCertDocCertVerifyConsultResponse 是支付宝官方 V3 OpenAPI 模型。
type UserCertDocCertVerifyConsultResponse struct {
	// FailParams 是具体哪些字段不一致。与预授权的输入字段对应，类型为字符串列表的JSON格式。
	FailParams *string `json:"fail_params,omitempty"`
	// FailReason 是校验不一致时，描述不一致的原因
	FailReason *string `json:"fail_reason,omitempty"`
	// Passed 是校验是否一致：T-一致、F-不一致
	Passed *string `json:"passed,omitempty"`
}

// UserCertDocCertVerifyPreConsultResponse 是支付宝官方 V3 OpenAPI 模型。
type UserCertDocCertVerifyPreConsultResponse struct {
	// VerifyId 是申请验证ID。有效期为 2 小时，过期后在校验接口使用会报错 "校验信息已过期"。
	VerifyId *string `json:"verify_id,omitempty"`
}
