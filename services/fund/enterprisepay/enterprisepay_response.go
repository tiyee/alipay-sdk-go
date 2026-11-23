package enterprisepay

// FundEnterprisePayGroupAddResponse 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePayGroupAddResponse struct {
	// AccountId 是因公付账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// FundGroupId 是资金群组号。
	FundGroupId *string `json:"fund_group_id,omitempty"`
	// OutGroupId 是外部群组号，业务号。
	OutGroupId *string `json:"out_group_id,omitempty"`
}

// FundEnterprisePayGroupDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePayGroupDeleteResponse struct {
}

// FundEnterprisePayGroupModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePayGroupModifyResponse struct {
}

// FundEnterprisePayGroupQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePayGroupQueryResponse struct {
	// AccountId 是企业签约账户ID。
	AccountId *string `json:"account_id,omitempty"`
	// FundGroupId 是内部群组号。
	FundGroupId *string `json:"fund_group_id,omitempty"`
	// FundIdentity 是该群所对应的出资主体账号。
	FundIdentity *string `json:"fund_identity,omitempty"`
	// GroupName 是群组名称。
	GroupName *string `json:"group_name,omitempty"`
	// OutGroupId 是外部群组号。
	OutGroupId *string `json:"out_group_id,omitempty"`
}

// FundEnterprisePayMemberModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePayMemberModifyResponse struct {
}

// FundEnterprisePaySignResponse 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePaySignResponse struct {
	// SignToken 是签约凭证。
	SignToken *string `json:"sign_token,omitempty"`
	// SignUrl 是签约链接，用于跳转到开户页面。
	SignUrl *string `json:"sign_url,omitempty"`
}

// FundEnterprisePayUnsignResponse 是支付宝官方 V3 OpenAPI 模型。
type FundEnterprisePayUnsignResponse struct {
	// UnsignToken 是解约凭证。
	UnsignToken *string `json:"unsign_token,omitempty"`
	// UnsignUrl 是解约链接。
	UnsignUrl *string `json:"unsign_url,omitempty"`
}
