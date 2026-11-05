// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package security

// SecurityRiskCustomerRiskSendRequest 是支付宝官方 V3 OpenAPI 模型。
type SecurityRiskCustomerRiskSendRequest struct {
	// BankCardNo 是交易对应的银行卡的卡号（支付宝直连和间连合作伙伴均可回传此参数）
	BankCardNo *string `json:"bank_card_no,omitempty"`
	// BusinessLicenseNo 是该账号对应的营业执照号码（支付宝直连和间连合作伙伴均可回传此参数）
	BusinessLicenseNo *string `json:"business_license_no,omitempty"`
	// CertNo 是该用户对应的身份证号码（支付宝直连和间连合作伙伴均可回传此参数）
	CertNo *string `json:"cert_no,omitempty"`
	// EmailAddress 是账户注册时填写的邮箱地址(适用于支付宝直连商户回传此信息)
	EmailAddress *string `json:"email_address,omitempty"`
	// ExternalId 是适用于间连业务场景下，商户在银行被分配的ID（适用于支付宝间连合作伙伴回传此信息）
	ExternalId *string `json:"external_id,omitempty"`
	// MerchName 是该笔风险交易购买的具体商品名称（适用于支付宝直连商户回传此信息）
	MerchName *string `json:"merch_name,omitempty"`
	// Mobile 是手机（支付宝直连和间连合作伙伴均可回传此参数）
	Mobile *string `json:"mobile,omitempty"`
	// MobileIp 是该账号登录合作伙伴平台时的IP地址（适用于支付宝直连商户回传此信息）
	MobileIp *string `json:"mobile_ip,omitempty"`
	// OrderIp 是该账户在商户平台下单时的MAC地址或IP地址（适用于支付宝直连商户回传此信息）
	OrderIp *string `json:"order_ip,omitempty"`
	// Pid 是适用于直连商户场景下的支付宝合作伙伴ID（适用于支付宝直连合作伙伴回传此信息）
	Pid *string `json:"pid,omitempty"`
	// PlatAccount 是指第三方在商户平台注册成功后，平台给予的账户号（适用于支付宝直连商户回传此信息）
	PlatAccount *string `json:"plat_account,omitempty"`
	// ProcessCode 是直连场景下，商户对该账户采取的措施： 	多种处理结果可调用多次，每次回传一种处理结果。 	（01：暂停发货； 	02：延迟结算； 	03：关停账户； 	04：暂停发货+关停账户； 	05：延迟结算+关停账户； 	06：其他； 	07：平台进行退款退订； 	08：平台跟用户沟通后，用户撤诉。 	间连场景下，商户回传处理结果： 	对商户的处理情况： 	01：已延迟结算，02：关停商户，99：其他 	多重处理结果可调用多次，每次回传一种处理结果
	ProcessCode *string `json:"process_code,omitempty"`
	// Smid 是指间连交易场景下商家被支付宝分配的ID（适用于支付宝间连合作伙伴回传此信息）
	Smid *string `json:"smid,omitempty"`
	// TradeNo 是支付宝推送的风险消息中的交易号（支付宝直连和间连合作伙伴均需回传此参数）
	TradeNo *string `json:"trade_no,omitempty"`
}

// SecurityRiskVerifyIdentityMiniAppConfirmRequest 是支付宝官方 V3 OpenAPI 模型。
type SecurityRiskVerifyIdentityMiniAppConfirmRequest struct {
	// Challenge 是业务的业务流水号，例如订单id
	Challenge *string `json:"challenge,omitempty"`
	// OpenId 是蚂蚁统一会员ID，验证时会与做身份认证时候会话中的openid做比对，如不相符核验结果不通过
	OpenId *string `json:"open_id,omitempty"`
	// UserId 是蚂蚁统一会员ID，验证时会与做身份认证时候会话中的userId做比对，如不相符核验结果不通过
	UserId *string `json:"user_id,omitempty"`
	// VerifyId 是核身校验id，是一次核身校验服务中唯一性的id
	VerifyId *string `json:"verify_id,omitempty"`
}

// SecurityProdSssQueryRequest is the request model for alipay.security.prod.sss.query.
type SecurityProdSssQueryRequest struct {
	// Aaa 在 URL query 中传递。1
	Aaa *string `query:"aaa" json:"-"`
	// AaaOpenId 在 URL query 中传递。1
	AaaOpenId            *string `query:"aaa_open_id" json:"-"`
	SecurityProdSssQuery `json:",inline"`
	// Bbb 是activity_id 词条描述
	Bbb string `json:"-" path:"bbb"`
}
