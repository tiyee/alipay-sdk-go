// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterprise

// InvoiceEnterpriseMerchantRelationCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseMerchantRelationCreateRequest struct {
	// EnterpriseId 在 URL query 中传递。企业ID
	EnterpriseId *string `query:"enterprise_id" json:"-"`
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// OperatorId 是传入操作人员标识
	OperatorId *string     `json:"operator_id,omitempty"`
	ShopInfo   *EcShopInfo `json:"shop_info,omitempty"`
	// TradeNo 是交易流水号
	TradeNo *string `json:"trade_no,omitempty"`
}
