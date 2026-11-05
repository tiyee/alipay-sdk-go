package project

// InvoiceExpenseRulesProjectEmployeeModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesProjectEmployeeModifyRequest struct {
	// AccountId 是共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// AddEmployeeList 是切换open_id前使用此字段：	需要添加的员工UID列表	约束：一次最多50个
	AddEmployeeList *[]string `json:"add_employee_list,omitempty"`
	// AddEmployeeOpenIdList 是切换open_id后使用此字段：	需要添加的员工open_id列表	约束：一次最多50个
	AddEmployeeOpenIdList *[]string `json:"add_employee_open_id_list,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// ProjectId 是项目ID
	ProjectId *string `json:"project_id,omitempty"`
	// RemoveEmployeeList 是切换open_id前使用此字段：	需要移除的员工UID列表	约束：一次最多50个
	RemoveEmployeeList *[]string `json:"remove_employee_list,omitempty"`
	// RemoveEmployeeOpenIdList 是切换open_id后前使用此字段： 需要移除的员工open_id列表 约束：一次最多50个
	RemoveEmployeeOpenIdList *[]string `json:"remove_employee_open_id_list,omitempty"`
}

// InvoiceExpenseRulesProjectInfoModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesProjectInfoModifyRequest struct {
	// AccountId 是企业ID
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EffectiveEndDate 是有效期截止（yyyy-MM-dd HH:mm:ss）	特殊说明：	1）与起始时间必须同时传，且大于起始时间	2）不传则默认不修改
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是有效期起始（yyyy-MM-dd HH:mm:ss）	特殊说明：	1）与截止时间必须同时传，且小于截止时间	2）不传入则默认不修改
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// ProjectId 是项目ID
	ProjectId *string `json:"project_id,omitempty"`
	// ProjectName 是项目名	特殊说明：不传入则默认不修改
	ProjectName *string `json:"project_name,omitempty"`
}

// InvoiceExpenseRulesProjectRuleCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesProjectRuleCreateRequest struct {
	// AccountId 是共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EffectiveEndDate 是有效期截止（yyyy-MM-dd HH:mm:ss，必须大于起始时间）
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是有效期起始（yyyy-MM-dd HH:mm:ss，必须小于截止时间）
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// EmployeeList 是切换open_id前使用此字段：员工支付宝ID列表	特殊说明：单次传入的最大员工数为50，后续通过修改项目的员工列表接口进行员工调整
	EmployeeList *[]string `json:"employee_list,omitempty"`
	// EmployeeOpenIdList 是切换open_id后使用此字段：员工open_id列表 特殊说明：单次传入的最大员工数为50，后续通过修改项目的员工列表接口进行员工调整
	EmployeeOpenIdList *[]string `json:"employee_open_id_list,omitempty"`
	// ExpenseCtrlRuleInfoGroupList 是费控规则列表
	ExpenseCtrlRuleInfoGroupList *[]ExpenseCtrRuleGroupInfo `json:"expense_ctrl_rule_info_group_list,omitempty"`
	// ProjectName 是项目名称	约束：	1）敏感词校验	2）不能重复
	ProjectName *string `json:"project_name,omitempty"`
}

// InvoiceExpenseRulesProjectRulesModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesProjectRulesModifyRequest struct {
	// AccountId 是企业ID
	AccountId *string `json:"account_id,omitempty"`
	// Action 是修改操作	枚举值：MODIFY_RULE（修改费控条件），仅支持MODIFY_RULE
	Action *string `json:"action,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// ExpenseCtrlRuleInfoGroupList 是费控规则列表
	ExpenseCtrlRuleInfoGroupList *[]ExpenseCtrRuleGroupInfo `json:"expense_ctrl_rule_info_group_list,omitempty"`
	// ProjectId 是项目ID
	ProjectId *string `json:"project_id,omitempty"`
}

// InvoiceExpenseRulesProjectRuleQueryRequest is the query request model for alipay.ebpp.invoice.expenserules.projectrule.query.
type InvoiceExpenseRulesProjectRuleQueryRequest struct {
	// AccountId 是共同账号ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// ProjectId 是项目id（可空）
	ProjectId *string `json:"project_id,omitempty" query:"project_id,omitempty"`
	// EmployeeId 是员工UID（可空）
	EmployeeId *string `json:"employee_id,omitempty" query:"employee_id,omitempty"`
	// EmployeeOpenId 是切换open_id后请使用：员工open_id（可空）
	EmployeeOpenId *string `json:"employee_open_id,omitempty" query:"employee_open_id,omitempty"`
	// PageNum 是页数
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,required"`
	// PageSize 是页大小
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,required"`
}
