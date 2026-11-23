package accountbook

// FundAccountBookCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAccountBookCreateRequest struct {
	// ExtInfo 是JSON格式，传递业务扩展参数。
	ExtInfo *string `json:"ext_info,omitempty"`
	// MerchantUserId 是外部商户系统会员的唯一标识，自定义传入。
	MerchantUserId *string `json:"merchant_user_id,omitempty"`
	// MerchantUserType 是外部商户用户类型:BUSINESS_EMPLOYEE。
	MerchantUserType *string `json:"merchant_user_type,omitempty"`
	// SceneCode 是资金记账本的业务场景。
	SceneCode *string `json:"scene_code,omitempty"`
}

// FundAccountBookNotifySubscribeRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAccountBookNotifySubscribeRequest struct {
	// AccountBookId 是记账本ID。
	AccountBookId *string `json:"account_book_id,omitempty"`
	// AgreementNo 是协议号。 若是基于协议的记账本，则agreement_no必传； 若是自创建的记账本，则agreement_no不传；。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BizScene 是场景码。固定为DEFAULT。
	BizScene *string `json:"biz_scene,omitempty"`
	// ProductCode 是产品码。固定为SATF_FUND_BOOK。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundAccountBookNotifyUnsubscribeRequest 是支付宝官方 V3 OpenAPI 模型。
type FundAccountBookNotifyUnsubscribeRequest struct {
	// AccountBookId 是记账本ID。
	AccountBookId *string `json:"account_book_id,omitempty"`
	// AgreementNo 是协议号。 若是基于协议的记账本，则agreement_no必传； 若是自创建的记账本，则agreement_no不传；。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BizScene 是场景码。固定为DEFAULT。
	BizScene *string `json:"biz_scene,omitempty"`
	// ProductCode 是产品码。固定为SATF_FUND_BOOK。
	ProductCode *string `json:"product_code,omitempty"`
}

// FundAccountBookNotifyQueryRequest 是 alipay.fund.accountbook.notify.query 的 V3 query 请求参数。
type FundAccountBookNotifyQueryRequest struct {
	// ProductCode 是产品码。固定为SATF_FUND_BOOK。
	ProductCode *string `query:"product_code,omitempty"`
	// BizScene 是场景码。固定为DEFAULT。
	BizScene *string `query:"biz_scene,omitempty"`
	// AccountBookId 是记账本ID。
	AccountBookId *string `query:"account_book_id,omitempty"`
	// AgreementNo 是协议号。 若是基于协议的记账本，则agreement_no必传； 若是自创建的记账本，则agreement_no不传；。
	AgreementNo *string `query:"agreement_no,omitempty"`
}

// FundAccountBookQueryRequest 是 alipay.fund.accountbook.query 的 V3 query 请求参数。
type FundAccountBookQueryRequest struct {
	// AccountBookId 是记账账簿id。
	AccountBookId *string `query:"account_book_id,omitempty"`
	// MerchantUserId 是商户会员的唯一标识，如果传入account_book_id此字段必传并比对一致性。。
	MerchantUserId *string `query:"merchant_user_id,omitempty"`
	// SceneCode 是资金记账本的开通场景码。
	SceneCode *string `query:"scene_code,omitempty"`
	// ExtInfo 是传递业务扩展参数。 注意：ext_info字段本质上是String类型，所以传递 的是转义后的json字符串。。
	ExtInfo *string `query:"ext_info,omitempty"`
}
