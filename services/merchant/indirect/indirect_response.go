package indirect

import (
	"encoding/json"
)

// MerchantIndirectAuthOrderCloseResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantIndirectAuthOrderCloseResponse = json.RawMessage

// MerchantIndirectAuthOrderCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantIndirectAuthOrderCreateResponse struct {
	// OrderNo 是申请单号
	OrderNo *string `json:"order_no,omitempty"`
	// OrderStatus 是意愿申请单状态，枚举值：审核中(AUDITING)、待联系人确认（CONTACT_CONFIRM），待法人确认（LEGAL_CONFIRM）、审核通过(AUDIT_PASS)、审核驳回(AUDIT_REJECT)、已冻结(AUDIT_FREEZE)、已撤回(CANCELED)、联系人处理中(CONTACT_PROCESSING)
	OrderStatus *string `json:"order_status,omitempty"`
}

// MerchantIndirectAuthOrderQueryStatusResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantIndirectAuthOrderQueryStatusResponse struct {
	// OrderNo 是商家认证申请单号
	OrderNo *string `json:"order_no,omitempty"`
	// OrderStatus 是商家认证申请单状态，枚举值：审核中(AUDITING)、待联系人确认（CONTACT_CONFIRM），待法人确认（LEGAL_CONFIRM）、审核通过(AUDIT_PASS)、审核驳回(AUDIT_REJECT)、已冻结(AUDIT_FREEZE)、已撤回(CANCELED)、联系人处理中（CONTACT_PROCESSING）
	OrderStatus *string `json:"order_status,omitempty"`
	// QrCode 是商家认证小程序二维码图片链接，申请单状态处于审核通过、待联系人确认、待法人确认、已冻结、审核失败时返回，商户使用支付宝APP扫码后完成后续确认、修改信息等动作。	此字段是平台生成并返回的二维码码值，通过浏览器访问即可获得一个二维码的图片。服务商在推广过程中，可引导商家使用支付宝APP扫描此二维码，进入认证申请单页面，完成认证申请的后续确认操作。
	QrCode *string `json:"qr_code,omitempty"`
	// VerifyList 是申请单审核失败时的原因描述列表
	VerifyList *[]IndirectAuthOrderFailedReason `json:"verify_list,omitempty"`
}

// MerchantIndirectSMIDBindQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MerchantIndirectSMIDBindQueryResponse struct {
	// CheckResult 是确认状态，枚举：已完成确认（AUTHORIZED）/未完成确认（UNAUTHORIZED）/已销户（CLOSED）/sub_merchant_id不存在（SMID_NOT_EXIST）
	CheckResult *string `json:"check_result,omitempty"`
}
