package account

// InstantDeliveryAccountCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InstantDeliveryAccountCreateResponse struct {
	// LogisticsAccountStatus 是配送公司账户创建结果列表。
	LogisticsAccountStatus *[]LogisticsAccountStatusDTO `json:"logistics_account_status,omitempty"`
}

// InstantDeliveryAccountQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InstantDeliveryAccountQueryResponse struct {
	// Balance 是账户余额。
	Balance *string `json:"balance,omitempty"`
	// BusinessLicense 是企业营业执照图片正反面。
	BusinessLicense *string `json:"business_license,omitempty"`
	// BusinessScope 是经营范围。
	BusinessScope *string `json:"business_scope,omitempty"`
	// CreditCode 是统一社会信用代码。
	CreditCode *string `json:"credit_code,omitempty"`
	// Email 是联系人邮箱。
	Email *string `json:"email,omitempty"`
	// EnterpriseAddress 是企业地址。
	EnterpriseAddress *string `json:"enterprise_address,omitempty"`
	// EnterpriseCity 是企业所在地，市编码。
	EnterpriseCity *string `json:"enterprise_city,omitempty"`
	// EnterpriseDistrict 是企业所在地，区编码。
	EnterpriseDistrict *string `json:"enterprise_district,omitempty"`
	// EnterpriseName 是企业全称。
	EnterpriseName *string `json:"enterprise_name,omitempty"`
	// EnterpriseProvince 是企业所在地，省编码。
	EnterpriseProvince *string `json:"enterprise_province,omitempty"`
	// EnterpriseType 是企业类型。
	EnterpriseType *string `json:"enterprise_type,omitempty"`
	// LogisticsAccountStatus 是在配送公司的账户状态信息。
	LogisticsAccountStatus *[]LogisticsAccountStatusDTO `json:"logistics_account_status,omitempty"`
	// Phone 是联系人手机号。
	Phone *string `json:"phone,omitempty"`
}
