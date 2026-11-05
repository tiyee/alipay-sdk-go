// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expensescene

// ExpenseMerchantInfo 是支付宝官方 V3 OpenAPI 模型。
type ExpenseMerchantInfo struct {
	// MerchantName 是商户名称
	MerchantName *string `json:"merchant_name,omitempty"`
	// Pid 是商户ID
	Pid *string `json:"pid,omitempty"`
	// RoleId 是商户ID
	RoleId *string `json:"role_id,omitempty"`
	// RoleType 是角色类型	0-直连商户	1-间联商户
	RoleType *string `json:"role_type,omitempty"`
	// ShopId 是门店ID
	ShopId *string `json:"shop_id,omitempty"`
	// ShopName 是门店名称
	ShopName *string `json:"shop_name,omitempty"`
}
