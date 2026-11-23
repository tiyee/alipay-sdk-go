// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package auth

// InvoiceAuthSignRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceAuthSignRequest struct {
	// AuthorizationType 是发票授权类型，可选值：INVOICE_AUTO_SYNC（发票自动回传）	INVOICE_TWOWAY_AUTO_SYNC（发票双向自动回传）
	AuthorizationType *string `json:"authorization_type,omitempty"`
	// ExtendFields 是扩展字段，格式为：KEY1=VALUE1,KEY2=VALUE2,KEY3=VALUE3	邮箱地址KEY为EMAIL_ADDRESS
	ExtendFields *string `json:"extend_fields,omitempty"`
	// MShortName 是开票商户品牌简称，与商户入驻时的品牌简称保持一致。
	MShortName *string `json:"m_short_name,omitempty"`
	// OpenId 是支付宝用户userId
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是支付宝用户userId
	UserId *string `json:"user_id,omitempty"`
}

// InvoiceAuthUnsignRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceAuthUnsignRequest struct {
	// AuthorizationType 是发票授权类型，可选值：INVOICE_AUTO_SYNC（发票自动回传）	INVOICE_TWOWAY_AUTO_SYNC（发票双向自动回传）
	AuthorizationType *string `json:"authorization_type,omitempty"`
	// ExtendFields 是扩展字段，格式为：KEY1=VALUE1,KEY2=VALUE2,KEY3=VALUE3	邮箱地址KEY为EMAIL_ADDRESS
	ExtendFields *string `json:"extend_fields,omitempty"`
	// MShortName 是开票商户品牌简称，与商户入驻时的品牌简称保持一致。
	MShortName *string `json:"m_short_name,omitempty"`
	// OpenId 是支付宝用户userId
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是支付宝用户userId
	UserId *string `json:"user_id,omitempty"`
}
