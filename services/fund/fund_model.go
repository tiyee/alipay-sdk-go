package fund

// ExtCardInfo 是支付宝官方 V3 OpenAPI 模型。
type ExtCardInfo struct {
	// BankAccName 是记账的外卡户名。
	BankAccName *string `json:"bank_acc_name,omitempty"`
	// CardBank 是记账的外卡开户行。
	CardBank *string `json:"card_bank,omitempty"`
	// CardBranch 是记账的外卡支行。
	CardBranch *string `json:"card_branch,omitempty"`
	// CardDeposit 是记账的外卡联行号。
	CardDeposit *string `json:"card_deposit,omitempty"`
	// CardLocation 是记账的外卡开户地址。
	CardLocation *string `json:"card_location,omitempty"`
	// CardNo 是记账的外卡卡号。
	CardNo *string `json:"card_no,omitempty"`
	// Status 是记账的外卡状态, A：正常状态; 其他：异常。
	Status *string `json:"status,omitempty"`
}

// FundExtInfo 是支付宝官方 V3 OpenAPI 模型。
type FundExtInfo struct {
	// FundIdentity 是出资主体账号，fund_identity_type是ALIPAY_USER_ID填支付宝会员ID（2088开头）； 是ALIPAY_LOGON_ID 填支付宝登录号。
	FundIdentity *string `json:"fund_identity,omitempty"`
	// FundIdentityType 是出资主体账号类型，目前支持如下类型： 1、ALIPAY_USER_ID 支付宝的会员ID 2、ALIPAY_LOGON_ID：支付宝登录号，支持邮箱和手机号格式。
	FundIdentityType *string `json:"fund_identity_type,omitempty"`
}
