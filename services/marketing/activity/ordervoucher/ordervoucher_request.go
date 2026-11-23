// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ordervoucher

// ActivityOrderVoucherAssociateRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherAssociateRequest struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// TradeNo 是支付宝支付下单时的商户订单号，欲与该商家券关联的支付宝支付
	TradeNo *string `json:"trade_no,omitempty"`
	// VoucherCode 是商家券券码。	限制:	券码必须已发放给用户。
	VoucherCode *string `json:"voucher_code,omitempty"`
}

// ActivityOrderVoucherCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherCreateRequest struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// ActivityName 是活动名称。不对用户进行展示，仅供商家在后台管理活动使用。
	ActivityName       *string             `json:"activity_name,omitempty"`
	BelongMerchantInfo *BelongMerchantInfo `json:"belong_merchant_info,omitempty"`
	// BizTag 是该字段废弃！后续不要使用该字段！商家券业务标签，影响商家券对C端用户的展示形式。
	BizTag *string `json:"biz_tag,omitempty"`
	// CodeMode 是码模式。MERCHANT_UPLOAD：商户上传自定义code，发券时系统随机选取上传的券code发放。MERCHANT_API：发奖时指定券码发奖，此模式无须提前上传券码。
	CodeMode      *string        `json:"code_mode,omitempty"`
	CustomerGuide *CustomerGuide `json:"customer_guide,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。	幂等作用：参数不变的情况下，再次请求返回与上一次相同的结果。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PublishEndTime 是券发放结束时间。格式为：yyyy-MM-dd HH:mm:ss
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// PublishStartTime 是券发放开始时间。格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime          *string                    `json:"publish_start_time,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info,omitempty"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayInfo        *VoucherDisplayInfo        `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule           *VoucherSendRuleDetail     `json:"voucher_send_rule,omitempty"`
	// VoucherType 是券类型。	兑换券仅支持一类券类型：EXCHANGE_VOUCHER: 兑换券；
	VoucherType        *string             `json:"voucher_type,omitempty"`
	VoucherUseRule     *VoucherUseRule     `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo *VoucherUseRuleInfo `json:"voucher_use_rule_info,omitempty"`
}

// ActivityOrderVoucherDisassociateRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherDisassociateRequest struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。 外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// TradeNo 是该交易在支付宝系统中的交易流水号。
	TradeNo *string `json:"trade_no,omitempty"`
	// VoucherCode 是券码。 限制: 券码必须已发放给用户。
	VoucherCode *string `json:"voucher_code,omitempty"`
}

// ActivityOrderVoucherAppendRequest is the request model for alipay.marketing.activity.ordervoucher.append.
type ActivityOrderVoucherAppendRequest struct {
	MarketingActivityOrderVoucherAppend `json:",inline"`
	// ActivityId 是活动id
	ActivityId string `json:"-" path:"activity_id"`
}

// ActivityOrderVoucherCodeCountRequest is the query request model for alipay.marketing.activity.ordervoucher.codecount.
type ActivityOrderVoucherCodeCountRequest struct {
	// ActivityId 是活动id
	ActivityId string `json:"-" path:"activity_id"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty" query:"merchant_access_mode,omitempty"`
}

// ActivityOrderVoucherCodeDepositRequest is the request model for alipay.marketing.activity.ordervoucher.codedeposit.
type ActivityOrderVoucherCodeDepositRequest struct {
	MarketingActivityOrderVoucherCodeDeposit `json:",inline"`
	// ActivityId 是活动id
	ActivityId string `json:"-" path:"activity_id"`
}

// ActivityOrderVoucherModifyRequest is the request model for alipay.marketing.activity.ordervoucher.modify.
type ActivityOrderVoucherModifyRequest struct {
	MarketingActivityOrderVoucherModify `json:",inline"`
	// ActivityId 是活动id1
	ActivityId string `json:"-" path:"activity_id"`
}

// ActivityOrderVoucherQueryRequest is the query request model for alipay.marketing.activity.ordervoucher.query.
type ActivityOrderVoucherQueryRequest struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty" query:"activity_id,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty" query:"merchant_access_mode,omitempty"`
}

// ActivityOrderVoucherRefundRequest is the request model for alipay.marketing.activity.ordervoucher.refund.
type ActivityOrderVoucherRefundRequest struct {
	MarketingActivityOrderVoucherRefund `json:",inline"`
	// ActivityId 是活动id
	ActivityId string `json:"-" path:"activity_id"`
	// VoucherCode 是已核销待退券的券码
	VoucherCode string `json:"-" path:"voucher_code"`
}

// ActivityOrderVoucherStopRequest is the request model for alipay.marketing.activity.ordervoucher.stop.
type ActivityOrderVoucherStopRequest struct {
	MarketingActivityOrderVoucherStop `json:",inline"`
	// ActivityId 是活动id
	ActivityId string `json:"-" path:"activity_id"`
}

// ActivityOrderVoucherUseRequest is the request model for alipay.marketing.activity.ordervoucher.use.
type ActivityOrderVoucherUseRequest struct {
	MarketingActivityOrderVoucherUse `json:",inline"`
	// ActivityId 是活动id
	ActivityId string `json:"-" path:"activity_id"`
	// VoucherCode 是待核销的券码。
	VoucherCode string `json:"-" path:"voucher_code"`
}
