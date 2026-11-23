package operation

// RecommendAccountDTO 是支付宝官方 V3 OpenAPI 模型。
type RecommendAccountDTO struct {
	// AlipayAccount 是推荐商家支付宝登录账号，邮箱或手机号码，脱敏处理。
	AlipayAccount *string `json:"alipay_account,omitempty"`
	// UserName 是推荐商家支付宝账号名称，脱敏处理。
	UserName *string `json:"user_name,omitempty"`
}
