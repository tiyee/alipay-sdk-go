package employee

// InvoiceExpenseRulesGroupEmployeeModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesGroupEmployeeModifyRequest struct {
	// AccountId 是企业共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// AddEmployeeList 是未切换 open_id 时请使用此字段：	需要添加的员工UID列表	特殊说明：一次最多50个
	AddEmployeeList *[]string `json:"add_employee_list,omitempty"`
	// AddEmployeeOpenIdList 是切换 open_id 后请使用此字段：	需要添加的open_id/企业码员工ID列表 特殊说明：一次最多50个
	AddEmployeeOpenIdList *[]string `json:"add_employee_open_id_list,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EnterpriseId 是企业码企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// GroupIdList 是费控规则ID列表
	GroupIdList *[]string `json:"group_id_list,omitempty"`
	// RemoveEmployeeList 是未切换 open_id 时请使用此字段：	需要移除的员工UID列表	特殊说明：一次最多50个
	RemoveEmployeeList *[]string `json:"remove_employee_list,omitempty"`
	// RemoveEmployeeOpenIdList 是切换 open_id 后请使用此字段：	需要移除的open_id/企业码员工ID列表 特殊说明：一次最多50个
	RemoveEmployeeOpenIdList *[]string `json:"remove_employee_open_id_list,omitempty"`
}

// InvoiceExpenseRulesEmployeeQueryRequest is the query request model for alipay.ebpp.invoice.expenserules.employee.query.
type InvoiceExpenseRulesEmployeeQueryRequest struct {
	// AccountId 是共同账号ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// EnterpriseId 是企业码企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// StandardId 是费控规则ID
	StandardId *string `json:"standard_id,omitempty" query:"standard_id,omitempty"`
	// PageNum 是页数
	PageNum *string `json:"page_num,omitempty" query:"page_num,omitempty"`
	// PageSize 是页大小
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
}

// InvoiceExpenseRulesEmployeeRulesQueryRequest is the query request model for alipay.ebpp.invoice.expenserules.employeerules.query.
type InvoiceExpenseRulesEmployeeRulesQueryRequest struct {
	// AccountId 是企业ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// EmployeeId 是未切换open_id时使用此字段：员工ID
	EmployeeId *string `json:"employee_id,omitempty" query:"employee_id,omitempty"`
	// EmployeeOpenId 是切换open_id后使用此字段：员工的open_id
	EmployeeOpenId *string `json:"employee_open_id,omitempty" query:"employee_open_id,omitempty"`
	// PageNum 是页码，默认值为1
	PageNum *string `json:"page_num,omitempty" query:"page_num,omitempty"`
	// PageSize 是每页行数，默认值为20
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
}
