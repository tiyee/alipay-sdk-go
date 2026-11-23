// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package voucher

import (
	"encoding/json"
)

// ActivityVoucherAppendResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherAppendResponse struct {
	// RechargeUrl 是预充值链接		限制	有效时间3天
	RechargeUrl *string `json:"recharge_url,omitempty"`
}

// ActivityVoucherCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherCreateResponse struct {
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
}

// ActivityVoucherModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherModifyResponse = json.RawMessage

// ActivityVoucherPublishResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherPublishResponse struct {
	// RechargeUrl 是预充值链接
	RechargeUrl *string `json:"recharge_url,omitempty"`
}

// ActivityVoucherQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherQueryResponse struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// ActivityId 是活动 ID
	ActivityId *string `json:"activity_id,omitempty"`
	// ActivityName 是活动名称。 不对用户进行展示，仅供商家在后台管理活动使用。
	ActivityName *string `json:"activity_name,omitempty"`
	// ActivityStatus 是活动状态。	如果该值为空，说明活动还未创建成功。
	ActivityStatus     *string                           `json:"activity_status,omitempty"`
	AuditResultInfo    *AuditResultInfo                  `json:"audit_result_info,omitempty"`
	BelongMerchantInfo *PaymentVoucherBelongMerchantInfo `json:"belong_merchant_info,omitempty"`
	// PublishEndTime 是券发放结束时间。 格式为：yyyy-MM-dd HH:mm:ss。
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// PublishStartTime 是券发放开始时间。 格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime          *string                    `json:"publish_start_time,omitempty"`
	Summary                   *VoucherSummary            `json:"summary,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info,omitempty"`
	VoucherBudgetInfo         *PaymentVoucherBudgetInfo  `json:"voucher_budget_info,omitempty"`
	VoucherBudgetSupplyInfo   *VoucherBudgetSupplyInfo   `json:"voucher_budget_supply_info,omitempty"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayInfo        *PaymentVoucherDisplayInfo `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	VoucherInventoryInfo      *VoucherInventoryInfo      `json:"voucher_inventory_info,omitempty"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule           *PaymentVoucherSendRule    `json:"voucher_send_rule,omitempty"`
	// VoucherType 是优惠类型。
	VoucherType        *string                      `json:"voucher_type,omitempty"`
	VoucherUseRule     *PaymentVoucherUseRuleDetail `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo *VoucherUseRuleInfo          `json:"voucher_use_rule_info,omitempty"`
}

// ActivityVoucherSendResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherSendResponse struct {
	// VoucherId 是本次发放的支付券id
	VoucherId *string `json:"voucher_id,omitempty"`
}

// ActivityVoucherStopResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityVoucherStopResponse = json.RawMessage
