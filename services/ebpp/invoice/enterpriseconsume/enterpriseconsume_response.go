// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterpriseconsume

// InvoiceEnterpriseConsumeConsumeBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeConsumeBatchQueryResponse struct {
	// ExpenseConsumeInfoList 是企业查询到的员工账单列表
	ExpenseConsumeInfoList *[]ExpenseConsumeInfo `json:"expense_consume_info_list,omitempty"`
}

// InvoiceEnterpriseConsumeDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeDetailQueryResponse struct {
	ExpenseConsumeInfo *ExpenseConsumeInfo `json:"expense_consume_info,omitempty"`
	ExpenseInvoiceInfo *ExpenseInvoiceInfo `json:"expense_invoice_info,omitempty"`
	VoucherInfo        *ExpenseVoucherInfo `json:"voucher_info,omitempty"`
}

// InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateResponse struct {
	// InvoiceRuleId 是开票规则ID
	InvoiceRuleId *string `json:"invoice_rule_id,omitempty"`
}

// InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyResponse struct {
	// Success 是success
	Success *bool `json:"success,omitempty"`
}

// InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryResponse struct {
	EnterpriseOpenRuleInfo *EnterpriseOpenRuleInfo `json:"enterprise_open_rule_info,omitempty"`
}

// InvoiceEnterpriseConsumeOpenRuleRelationCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeOpenRuleRelationCreateResponse struct {
	// Success 是是否成功
	Success *bool `json:"success,omitempty"`
}

// InvoiceEnterpriseConsumeRelateDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeRelateDetailQueryResponse struct {
	ExpenseCtrlConsumeInfo *ExpenseCtrlConsumeInfo `json:"expense_ctrl_consume_info,omitempty"`
}

// InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryResponse struct {
	// AccountId 是企业汇总场景下，即企业ID
	AccountId *string `json:"account_id,omitempty"`
	// ApplyDate 是申请时间
	ApplyDate *string `json:"apply_date,omitempty"`
	// BatchId 是已废弃，开票批次 ID（建议使用summary_id）
	BatchId *string `json:"batch_id,omitempty"`
	// BatchStatus 是批次状态	    INIT：汇总批次初始化状态，未进行汇总	    SUMMARY：汇总批次化汇总完成，汇总金额计算完成	    PROCESS：汇总批次初始化状态，汇总开票申请提交成功	    NOT_SUPPORT：汇总开票商户业务能力不支持	    FAIL：汇总批次申请失败状态	    SUCCESS：汇总批次申请成功状态，此时发票已回传	    MANUAL_CONFIRM：汇总批次人工确认收票
	BatchStatus *string `json:"batch_status,omitempty"`
	// CompleteVoucherList 是发票内容
	CompleteVoucherList *[]CompleteVoucherInfo `json:"complete_voucher_list,omitempty"`
	// GmtBizEnd 是账期结束时间
	GmtBizEnd *string `json:"gmt_biz_end,omitempty"`
	// GmtBizStart 是账期起始时间
	GmtBizStart *string `json:"gmt_biz_start,omitempty"`
	// StatusShowContent 是状态对应的展示明细内容	    初始化	    汇总完成	    开票中...	    不支持开票	    开票失败	    已开票	    确认已收票
	StatusShowContent *string `json:"status_show_content,omitempty"`
	// SummaryId 是开票批次 ID
	SummaryId *string `json:"summary_id,omitempty"`
	// UserId 是企业汇总场景下，即企业ID（请使用account_id）
	UserId *string `json:"user_id,omitempty"`
}
