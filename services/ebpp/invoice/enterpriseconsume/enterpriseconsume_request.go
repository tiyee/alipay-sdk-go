// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterpriseconsume

// InvoiceEnterpriseConsumeConsumeBatchQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeConsumeBatchQueryRequest struct {
	// AccountId 是企业ID
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// ConsumptionEnd 是交易结束时间，格式：yyyy-MM-dd HH:mm:ss
	ConsumptionEnd *string `json:"consumption_end,omitempty"`
	// ConsumptionStart 是交易开始时间，格式：yyyy-MM-dd HH:mm:ss
	ConsumptionStart *string `json:"consumption_start,omitempty"`
	// EmployeeList 是员工支付宝UID列表，单次传入最大员工数量为10
	EmployeeList *[]string `json:"employee_list,omitempty"`
	// EmployeeOpenIds 是员工支付宝UID列表，单次传入最大员工数量为10
	EmployeeOpenIds *[]string `json:"employee_open_ids,omitempty"`
}

// InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateRequest struct {
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BillMonthDay 是账期结束日期
	BillMonthDay *int64 `json:"bill_month_day,omitempty"`
	// BillScope 是开票账单范围
	BillScope *string `json:"bill_scope,omitempty"`
	// CombinedPayMode 是组合支付模式
	CombinedPayMode *string `json:"combined_pay_mode,omitempty"`
	// DefaultInvoiceKind 是默认开票类型
	DefaultInvoiceKind *string `json:"default_invoice_kind,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InvoiceRemarkValueRule 是发票备注取值规则。NONE表示不使用额度发放说明；QUOTA_ISSUE_DESC表示将因公付额度发放说明作为发票备注。不传、传null或空字符串时默认NONE。
	InvoiceRemarkValueRule *string `json:"invoice_remark_value_rule,omitempty"`
	// InvoiceRuleName 是开票规则名称
	InvoiceRuleName *string `json:"invoice_rule_name,omitempty"`
	// InvoiceTitleId 是发票抬头
	InvoiceTitleId *string `json:"invoice_title_id,omitempty"`
	// OpenMode 是开票模式
	OpenMode *string `json:"open_mode,omitempty"`
	// ReceiveAddress 是收件人地址
	ReceiveAddress *string `json:"receive_address,omitempty"`
	// ReceiveName 是收件人姓名
	ReceiveName *string `json:"receive_name,omitempty"`
	// ReceivePhone 是收件人手机号
	ReceivePhone *string `json:"receive_phone,omitempty"`
	// SellerType 是销方类型
	SellerType *string `json:"seller_type,omitempty"`
	// Tag 是开票规则标记
	Tag *string `json:"tag,omitempty"`
}

// InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyRequest struct {
	// AccountId 是企业共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BillScope 是开票账单范围
	BillScope *string `json:"bill_scope,omitempty"`
	// CombinedPayMode 是组合支付模式
	CombinedPayMode *string `json:"combined_pay_mode,omitempty"`
	// DefaultInvoiceKind 是默认开票类型
	DefaultInvoiceKind *string `json:"default_invoice_kind,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InvoiceRemarkValueRule 是发票备注取值规则。NONE表示不使用额度发放说明；QUOTA_ISSUE_DESC表示将因公付额度发放说明作为发票备注。修改接口未传、传null或空字符串时保留原配置。
	InvoiceRemarkValueRule *string `json:"invoice_remark_value_rule,omitempty"`
	// InvoiceRuleId 是开票规则ID
	InvoiceRuleId *string `json:"invoice_rule_id,omitempty"`
	// InvoiceRuleName 是开票规则名称
	InvoiceRuleName *string `json:"invoice_rule_name,omitempty"`
	// InvoiceTitleId 是发票抬头
	InvoiceTitleId *string `json:"invoice_title_id,omitempty"`
	// ReceiveAddress 是收件人地址
	ReceiveAddress *string `json:"receive_address,omitempty"`
	// ReceiveName 是收件人姓名
	ReceiveName *string `json:"receive_name,omitempty"`
	// ReceivePhone 是收件人手机号
	ReceivePhone *string `json:"receive_phone,omitempty"`
	// SellerType 是销方类型
	SellerType *string `json:"seller_type,omitempty"`
}

// InvoiceEnterpriseConsumeOpenRuleRelationCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeOpenRuleRelationCreateRequest struct {
	// AccountId 是企业共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InvoiceRuleId 是开票规则ID
	InvoiceRuleId *string `json:"invoice_rule_id,omitempty"`
	// StandardIdList 是费控规则ID列表
	StandardIdList *[]string `json:"standard_id_list,omitempty"`
}

// InvoiceEnterpriseConsumeDetailQueryRequest is the query request model for alipay.ebpp.invoice.enterpriseconsume.detail.query.
type InvoiceEnterpriseConsumeDetailQueryRequest struct {
	// AccountId 是企业ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// VoucherId 是账单凭证ID
	VoucherId *string `json:"voucher_id,omitempty" query:"voucher_id,omitempty"`
}

// InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryRequest is the query request model for alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.query.
type InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryRequest struct {
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// InvoiceRuleId 是开票规则ID
	InvoiceRuleId *string `json:"invoice_rule_id,omitempty" query:"invoice_rule_id,omitempty"`
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
}

// InvoiceEnterpriseConsumeRelateDetailQueryRequest is the query request model for alipay.ebpp.invoice.enterpriseconsume.relatedetail.query.
type InvoiceEnterpriseConsumeRelateDetailQueryRequest struct {
	// AccountId 是企业ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// VoucherId 是账单凭证ID
	VoucherId *string `json:"voucher_id,omitempty" query:"voucher_id,omitempty"`
}

// InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryRequest is the query request model for alipay.ebpp.invoice.enterpriseconsume.summaryinvoicedetail.query.
type InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryRequest struct {
	// AccountId 是企业 ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// SummaryId 是开票批次 ID
	SummaryId *string `json:"summary_id,omitempty" query:"summary_id,omitempty"`
}
