// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchantlist

// InvoiceMerchantListEnterApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceMerchantListEnterApplyRequest struct {
	MerchantBase          *MerchantBaseEnterOpen      `json:"merchant_base,omitempty"`
	SubMerchantCommonInfo *SubMerchantCommonEnterOpen `json:"sub_merchant_common_info,omitempty"`
	// SubMerchantList 是商户门店列表信息，最多传入30个门店信息
	SubMerchantList *[]SubMerchantEnterOpen `json:"sub_merchant_list,omitempty"`
}
