// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterprise

// InvoiceEnterpriseMerchantRelationCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseMerchantRelationCreateResponse struct {
	// Pid 是门店收单pid
	Pid *string `json:"pid,omitempty"`
	// RoleType 是门店直间连类型
	RoleType *string `json:"role_type,omitempty"`
	// ShopId 是门店ID
	ShopId *string `json:"shop_id,omitempty"`
	// Success 是是否成功
	Success *bool `json:"success,omitempty"`
}
