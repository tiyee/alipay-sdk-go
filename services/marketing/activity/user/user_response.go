// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package user

// ActivityUserBatchQueryVoucherResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityUserBatchQueryVoucherResponse struct {
	// PageNum 是分页查询页码
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是分页查询单页数据条数
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalSize 是总数量
	TotalSize *int64 `json:"total_size,omitempty"`
	// UserVoucherInfos 是用户券实例
	UserVoucherInfos *[]UserVoucherInfo `json:"user_voucher_infos,omitempty"`
}

// ActivityUserQueryVoucherResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityUserQueryVoucherResponse struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// ActivityId 是活动 id
	ActivityId *string `json:"activity_id,omitempty"`
	// AssociateTradeNo 是若商家券操作过关联商户订单信息，则该字段返回商家券已关联的商户订单号。
	AssociateTradeNo *string `json:"associate_trade_no,omitempty"`
	// AvailableBeginTime 是券可用开始时间。格式为：yyyy-MM-dd HH:mm:ss
	AvailableBeginTime *string `json:"available_begin_time,omitempty"`
	// AvailableEndTime 是券可用结束时间。格式为：yyyy-MM-dd HH:mm:ss
	AvailableEndTime *string `json:"available_end_time,omitempty"`
	// BelongMerchantId 是券归属 pid
	BelongMerchantId *string `json:"belong_merchant_id,omitempty"`
	// CreateTime 是领券时间
	CreateTime                *string                    `json:"create_time,omitempty"`
	UserVoucherBaseInfo       *UserVoucherBaseInfo       `json:"user_voucher_base_info,omitempty"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayInfo        *CommonVoucherDisplayInfo  `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	// VoucherName 是对消费者展示的券(商品)名称。
	VoucherName         *string                `json:"voucher_name,omitempty"`
	VoucherSendModeInfo *VoucherSendModeInfo   `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule     *CommonVoucherSendRule `json:"voucher_send_rule,omitempty"`
	// VoucherStatus 是券状态。
	VoucherStatus *string `json:"voucher_status,omitempty"`
	// VoucherType 是券类型。
	VoucherType        *string               `json:"voucher_type,omitempty"`
	VoucherUseRule     *CommonVoucherUseRule `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo *VoucherUseRuleInfo   `json:"voucher_use_rule_info,omitempty"`
}
