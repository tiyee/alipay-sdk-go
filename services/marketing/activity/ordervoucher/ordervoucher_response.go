// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ordervoucher

import (
	"encoding/json"
)

// ActivityOrderVoucherAppendResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherAppendResponse = json.RawMessage

// ActivityOrderVoucherAssociateResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherAssociateResponse struct {
	// AssociateTime 是支付宝系统关联订单成功的时间。 格式为：yyyy-MM-dd HH:mm:ss
	AssociateTime *string `json:"associate_time,omitempty"`
}

// ActivityOrderVoucherCodeCountResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherCodeCountResponse struct {
	// SuccessCount 是已经导入的券码数量
	SuccessCount *int64 `json:"success_count,omitempty"`
}

// ActivityOrderVoucherCodeDepositResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherCodeDepositResponse struct {
	// FailCount 是导入失败的券码数量
	FailCount *int64 `json:"fail_count,omitempty"`
	// FailVoucherCodeDetailList 是上传失败券码详情列表，包括上传失败的券码及对应的错误码和错误原因
	FailVoucherCodeDetailList *[]FailVoucherCodeDetail `json:"fail_voucher_code_detail_list,omitempty"`
	// SuccessCount 是导入成功的券码数量
	SuccessCount *int64 `json:"success_count,omitempty"`
	// SuccessVoucherCodeList 是上传成功的券码列表
	SuccessVoucherCodeList *[]string `json:"success_voucher_code_list,omitempty"`
}

// ActivityOrderVoucherCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherCreateResponse struct {
	// ActivityId 是活动id
	ActivityId                      *string                          `json:"activity_id,omitempty"`
	VoucherAvailableScopeResultInfo *VoucherAvailableScopeResultInfo `json:"voucher_available_scope_result_info,omitempty"`
	VoucherUseRuleResult            *OrderVoucherUseRuleResult       `json:"voucher_use_rule_result,omitempty"`
}

// ActivityOrderVoucherDisassociateResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherDisassociateResponse struct {
	// DisassociateTime 是支付宝系统取消关联订单成功的时间。 格式为：yyyy-MM-dd HH:mm:ss
	DisassociateTime *string `json:"disassociate_time,omitempty"`
}

// ActivityOrderVoucherModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherModifyResponse struct {
	VoucherAvailableScopeResultInfo *VoucherAvailableScopeResultInfo `json:"voucher_available_scope_result_info,omitempty"`
	VoucherUseRuleResult            *OrderVoucherUseRuleResult       `json:"voucher_use_rule_result,omitempty"`
}

// ActivityOrderVoucherQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherQueryResponse struct {
	ActivityBaseInfo *ActivityBaseInfo `json:"activity_base_info,omitempty"`
	// ActivityId 是活动id
	ActivityId *string `json:"activity_id,omitempty"`
	// ActivityName 是活动名称。不对用户进行展示，仅供商家在后台管理活动使用。
	ActivityName *string `json:"activity_name,omitempty"`
	// ActivityOperationStatus 是活动操作状态
	ActivityOperationStatus *string `json:"activity_operation_status,omitempty"`
	// ActivityStatus 是活动状态。已激活状态代表活动正常运行，已终止状态代表商家终止活动，不可再发放。
	ActivityStatus     *string             `json:"activity_status,omitempty"`
	AuditResultInfo    *AuditResultInfo    `json:"audit_result_info,omitempty"`
	BelongMerchantInfo *BelongMerchantInfo `json:"belong_merchant_info,omitempty"`
	// BizTag 是商家券业务标签，影响商家券对C端用户的展示形式。
	BizTag        *string        `json:"biz_tag,omitempty"`
	CustomerGuide *CustomerGuide `json:"customer_guide,omitempty"`
	// PublishEndTime 是券发放结束时间。格式为：yyyy-MM-dd HH:mm:ss
	PublishEndTime *string `json:"publish_end_time,omitempty"`
	// PublishStartTime 是券发放开始时间。格式为：yyyy-MM-dd HH:mm:ss
	PublishStartTime          *string                    `json:"publish_start_time,omitempty"`
	Summary                   *VoucherSummary            `json:"summary,omitempty"`
	VoucherAvailableScopeInfo *VoucherAvailableScopeInfo `json:"voucher_available_scope_info,omitempty"`
	VoucherCustomerGuideInfo  *VoucherCustomerGuideInfo  `json:"voucher_customer_guide_info,omitempty"`
	VoucherDeductInfo         *VoucherDeductInfo         `json:"voucher_deduct_info,omitempty"`
	VoucherDisplayInfo        *VoucherDisplayInfo        `json:"voucher_display_info,omitempty"`
	VoucherDisplayPatternInfo *VoucherDisplayPatternInfo `json:"voucher_display_pattern_info,omitempty"`
	VoucherInventoryInfo      *VoucherInventoryInfo      `json:"voucher_inventory_info,omitempty"`
	VoucherSendModeInfo       *VoucherSendModeInfo       `json:"voucher_send_mode_info,omitempty"`
	VoucherSendRule           *VoucherSendRuleDetail     `json:"voucher_send_rule,omitempty"`
	// VoucherType 是券类型
	VoucherType        *string             `json:"voucher_type,omitempty"`
	VoucherUseRule     *VoucherUseRule     `json:"voucher_use_rule,omitempty"`
	VoucherUseRuleInfo *VoucherUseRuleInfo `json:"voucher_use_rule_info,omitempty"`
}

// ActivityOrderVoucherRefundResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherRefundResponse struct {
	// ActivityId 是优惠券活动id
	ActivityId                 *string                     `json:"activity_id,omitempty"`
	VoucherUseDetailResultInfo *VoucherUseDetailResultInfo `json:"voucher_use_detail_result_info,omitempty"`
}

// ActivityOrderVoucherStopResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherStopResponse = json.RawMessage

// ActivityOrderVoucherUseResponse 是支付宝官方 V3 OpenAPI 模型。
type ActivityOrderVoucherUseResponse struct {
	// ActivityId 是优惠券活动id
	ActivityId                 *string                     `json:"activity_id,omitempty"`
	VoucherUseDetailResultInfo *VoucherUseDetailResultInfo `json:"voucher_use_detail_result_info,omitempty"`
}
