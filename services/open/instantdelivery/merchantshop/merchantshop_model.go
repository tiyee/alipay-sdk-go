package merchantshop

// LogisticsShopStatusDTO 是支付宝官方 V3 OpenAPI 模型。
type LogisticsShopStatusDTO struct {
	// AuditDesc 是商家在配送公司的账户的审核说明.。
	AuditDesc *string `json:"audit_desc,omitempty"`
	// LogisticsCode 是配送公司物流编码。。
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// LogisticsName 是配送公司名称。。
	LogisticsName *string `json:"logistics_name,omitempty"`
	// Status 是门店状态，AUDITING:审核中，AUDIT_REJECT:审核驳回，VALID:已生效，UPDATING:更新中，WAIT_CREATE:待开通(存量门店若需要支持后面新增的配送公司，需要调用门店更新接口发起申请)。
	Status *string `json:"status,omitempty"`
}

// MerchantShopDTO 是支付宝官方 V3 OpenAPI 模型。
type MerchantShopDTO struct {
	// ShopName 是门店名称，最长不超过256个字符。。
	ShopName *string `json:"shop_name,omitempty"`
	// ShopNo 是门店编号。。
	ShopNo *string `json:"shop_no,omitempty"`
}
