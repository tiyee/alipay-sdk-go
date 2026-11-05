package accountbook

// AuthInfoDTO 是支付宝官方 V3 OpenAPI 模型。
type AuthInfoDTO struct {
	// AgreementNo 是授权协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AlipayLogonId 是授权的支付宝账号，脱敏。。
	AlipayLogonId *string `json:"alipay_logon_id,omitempty"`
}
