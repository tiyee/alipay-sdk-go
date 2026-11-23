package account

// FundAccountQueryRequest 是 alipay.fund.account.query 的 V3 query 请求参数。
type FundAccountQueryRequest struct {
	// MerchantUserId 是商户会员的唯一标识。如果传入的user_id为虚拟账户userId，此字段必传并比对一致性。。
	MerchantUserId *string `query:"merchant_user_id,omitempty"`
	// AlipayUserId 是支付宝会员 id。。
	AlipayUserId *string `query:"alipay_user_id,omitempty"`
	// AlipayOpenId 是支付宝openId。
	AlipayOpenId *string `query:"alipay_open_id,omitempty"`
	// AccountProductCode 是开户产品码。如果查询托管子户余额，必传且必须传入与开户时传入的值一致。。
	AccountProductCode *string `query:"account_product_code,omitempty"`
	// AccountType 是查询的账号类型，查询余额账户值为ACCTRANS_ACCOUNT。必填。。
	AccountType *string `query:"account_type,omitempty"`
	// AccountSceneCode 是开户场景码，与开户产品码不可同时传递。。
	AccountSceneCode *string `query:"account_scene_code,omitempty"`
	// ExtInfo 是JSON格式，传递业务扩展参数。。
	ExtInfo *string `query:"ext_info,omitempty"`
}
