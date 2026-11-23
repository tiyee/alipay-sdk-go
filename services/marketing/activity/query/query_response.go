// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package query

// ActivityQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityQueryResponse struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// ActivityName 是活动名称。 不对用户进行展示，仅供商家在后台管理活动使用。
	ActivityName *string `json:"activity_name,omitempty"`
	// ActivityStatus 是活动状态。活动已激活，表示活动已经生效，等到活动开始(publish_start_time)之后用户就可以参与活动。活动已暂停，表示商户临时暂停该活动，该状态下用户不能参与活动。活动已结束，表示商户主动停止活动或活动到期结束(publish_end_time)不能再进行领取或修改等操作。
	ActivityStatus *string `json:"activity_status,omitempty"`
	// BelongMerchantId 是归属商户PID
	BelongMerchantId *string `json:"belong_merchant_id,omitempty"`
	// PublishEndTime 是券发放结束时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// PublishStartTime 是券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime          *string                    `json:"publish_start_time,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info,omitempty"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayInfo        *CommonVoucherDisplayInfo  `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule           *CommonVoucherSendRule     `json:"voucher_send_rule,omitempty"`
	// VoucherType 是券类型。
	VoucherType        *string               `json:"voucher_type,omitempty"`
	VoucherUseRule     *CommonVoucherUseRule `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo *VoucherUseRuleInfo   `json:"voucher_use_rule_info,omitempty"`
}
