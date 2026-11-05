// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package operation

// BankActivityInfo 是支付宝官方 V3 OpenAPI 模型。
type BankActivityInfo struct {
	// BindCard 是是否绑定银行卡
	BindCard *bool `json:"bind_card,omitempty"`
	// CardName 是银行卡名称
	CardName *string `json:"card_name,omitempty"`
	// DiscountDesc 是优惠描述
	DiscountDesc *string `json:"discount_desc,omitempty"`
	// Logo 是活动logo
	Logo *string `json:"logo,omitempty"`
}
