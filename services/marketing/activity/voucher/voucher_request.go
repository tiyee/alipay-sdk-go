// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package voucher

// ActivityVoucherAppendRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherAppendRequest struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。		幂等作用：	参数不变的情况下，再次请求返回与上一次相同的结果。		外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// VoucherQuantity 是发行券的数量。。 限制： 券总预算<= 99999999 特别说明： 该字段的含义是追加到的数量。 不可减少，只能增加。
	VoucherQuantity *int64 `json:"voucher_quantity,omitempty"`
}

// ActivityVoucherCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherCreateRequest struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// ActivityName 是活动名称。
	ActivityName       *string                           `json:"activity_name,omitempty"`
	BelongMerchantInfo *PaymentVoucherBelongMerchantInfo `json:"belong_merchant_info,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是用作幂等控制。		幂等作用：	参数不变的情况下，再次请求返回与上一次相同的结果。		外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PublishEndTime 是券发放结束时间。 格式为：yyyy-MM-dd HH:mm:ss 限制： 券发放结束时间 publish_end_time 与 券发放开始时间  publish_start_time 间隔必须小于等于180天
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// PublishStartTime 是券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime          *string                    `json:"publish_start_time,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info,omitempty"`
	VoucherBudgetInfo         *PaymentVoucherBudgetInfo  `json:"voucher_budget_info,omitempty"`
	VoucherBudgetSupplyInfo   *VoucherBudgetSupplyInfo   `json:"voucher_budget_supply_info,omitempty"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayInfo        *PaymentVoucherDisplayInfo `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule           *PaymentVoucherSendRule    `json:"voucher_send_rule,omitempty"`
	// VoucherType 是券类型
	VoucherType        *string                `json:"voucher_type,omitempty"`
	VoucherUseRule     *PaymentVoucherUseRule `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo *VoucherUseRuleInfo    `json:"voucher_use_rule_info,omitempty"`
}

// ActivityVoucherModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherModifyRequest struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。 幂等作用： 参数不变的情况下，再次请求返回与上一次相同的结果。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// PublishEndTime 是券发放结束时间。格式为：yyyy-MM-dd HH:mm:ss
	PublishEndTime            *string                      `json:"publish_end_time,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo   `json:"voucher_available_scope_info,omitempty"`
	VoucherSendModeInfo       *VoucherSendModeInfo         `json:"voucher_send_mode_info,omitempty"`
	VoucherUseRule            *PaymentVoucherUseRuleModify `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo        *VoucherUseRuleInfo          `json:"voucher_use_rule_info,omitempty"`
}

// ActivityVoucherPublishRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherPublishRequest struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是用作幂等控制		幂等作用：	参数不变的情况下，再次请求返回与上一次相同的结果。	外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// ActivityVoucherSendRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherSendRequest struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OpenId 是领券的支付宝用户openId
	OpenId *string `json:"open_id,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制。		幂等作用：	参数不变的情况下，再次请求返回与上一次相同的结果。		外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
	// UserId 是领券的支付宝user_id账号
	UserId *string `json:"user_id,omitempty"`
}

// ActivityVoucherStopRequest 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherStopRequest struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// MerchantAccessMode 是商户接入模式
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty"`
	// OutBizNo 是64	外部业务单号，用作幂等控制。		幂等作用：	参数不变的情况下，再次请求返回与上一次相同的结果。		外部接入方需保证业务单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// ActivityVoucherQueryRequest is the query request model for alipay.marketing.activity.voucher.query.
type ActivityVoucherQueryRequest struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty" query:"activity_id,omitempty"`
	// MerchantAccessMode 是商户接入模式。
	MerchantAccessMode *string `json:"merchant_access_mode,omitempty" query:"merchant_access_mode,omitempty"`
}
