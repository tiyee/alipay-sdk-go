package account

// LogisticsAccountStatusDTO 是支付宝官方 V3 OpenAPI 模型。
type LogisticsAccountStatusDTO struct {
	// AuditDesc 是商家在配送公司的账户的审核说明。
	AuditDesc *string `json:"audit_desc,omitempty"`
	// LogisticsCode 是配送公司物流编码。。
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// LogisticsName 是配送公司名称。。
	LogisticsName *string `json:"logistics_name,omitempty"`
	// Status 是账户状态，AUDITING:审核中，AUDIT_REJECT:审核驳回，VALID:已生效。
	Status *string `json:"status,omitempty"`
}
