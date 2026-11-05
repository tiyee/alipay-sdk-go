package certdoc

// UserCertDocCertVerifyPreConsultRequest 是支付宝官方 V3 OpenAPI 模型。
type UserCertDocCertVerifyPreConsultRequest struct {
	// CertNo 是证件号
	CertNo *string `json:"cert_no,omitempty"`
	// CertType 是证件类型。暂仅支持 IDENTITY_CARD （身份证）。
	CertType *string `json:"cert_type,omitempty"`
	// ExtInfo 是拓展字段,JSON格式
	ExtInfo *string `json:"ext_info,omitempty"`
	// LogonId 是支付宝登录名
	LogonId *string `json:"logon_id,omitempty"`
	// Mobile 是手机号码
	Mobile *string `json:"mobile,omitempty"`
	// UserName 是真实姓名
	UserName *string `json:"user_name,omitempty"`
}

// UserCertDocCertVerifyConsultRequest is the query request model for alipay.user.certdoc.certverify.consult.
type UserCertDocCertVerifyConsultRequest struct {
	// AuthToken 是用户授权令牌
	AuthToken *string `json:"auth_token,omitempty" query:"auth_token,omitempty"`
	// VerifyId 是信息校验验证ID。通过alipay.user.certdoc.certverify.preconsult(实名证件信息比对验证预咨询)接口获取。
	VerifyId *string `json:"verify_id,omitempty" query:"verify_id,omitempty"`
}
