package certify

// UserCertifyOpenInitializeResponse 是支付宝官方 V3 OpenAPI 模型。
type UserCertifyOpenInitializeResponse struct {
	// CertifyId 是本次申请操作的唯一标识，商户需要记录，后续的操作都需要用到
	CertifyId *string `json:"certify_id,omitempty"`
}

// UserCertifyOpenQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type UserCertifyOpenQueryResponse struct {
	// FailReason 是认证错误原因，扩展字段，默认不返回
	FailReason *string `json:"fail_reason,omitempty"`
	// IdentityInfo 是认证的主体信息，一般的认证场景返回为空
	IdentityInfo *string `json:"identity_info,omitempty"`
	// MaterialInfo 是认证主体附件信息，主要为图片类材料，一般的认证场景都是返回空
	MaterialInfo *string `json:"material_info,omitempty"`
	// Passed 是是否通过，通过为T，不通过为F
	Passed *string `json:"passed,omitempty"`
}
