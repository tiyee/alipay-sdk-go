// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package applystatus

// OrderApplyStatusBriefDTO 是支付宝官方 V3 OpenAPI 模型。
type OrderApplyStatusBriefDTO struct {
	// ApplyStatus 是SUCCESS：开票成功	FAIL：开票失败	PROCESS：开票中	NOTEXIST：申请不存在
	ApplyStatus *string `json:"apply_status,omitempty"`
	// OrderNo 是开票申请时传入订单号（支持主单号、子单号），不限是否为支付宝体内交易单号
	OrderNo *string `json:"order_no,omitempty"`
}
