// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package invite

// ProductInviteStatusInfo 是支付宝官方 V3 OpenAPI 模型。
type ProductInviteStatusInfo struct {
	// ProductCode 是产品编码。
	ProductCode *string `json:"product_code,omitempty"`
	// ProductName 是产品名称。
	ProductName *string `json:"product_name,omitempty"`
	// Status 是NONE：未签约，表示还没有签约该产品 MERCHANT_AUDITING：审核中，已经有合约在审核中，请等待审核完成 MERCHANT_CONFIRM_SUCCESS：已生效，表示合约已经生效，不需要再签约了 MERCHANT_APPLY_ORDER_CANCELED：审核未通过。
	Status *string `json:"status,omitempty"`
}
