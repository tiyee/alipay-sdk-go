// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

// ActivityMerchantInfo 是支付宝官方 V3 OpenAPI 模型。
type ActivityMerchantInfo struct {
	// MerchantId 是商户PID,默认为当前接口调用商户。
	MerchantId *string `json:"merchant_id,omitempty"`
	// MerchantIdType 是商户支付宝id类型。
	MerchantIdType *string `json:"merchant_id_type,omitempty"`
}
