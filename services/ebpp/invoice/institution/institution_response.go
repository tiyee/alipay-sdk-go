// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package institution

// InvoiceInstitutionCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionCreateResponse struct {
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// IssueRuleIdInfoList 是发放规则id列表
	IssueRuleIdInfoList *[]IssueRuleIdInfo `json:"issue_rule_id_info_list,omitempty"`
	// StandardIdInfoList 是使用规则id列表
	StandardIdInfoList *[]StandardIdInfo `json:"standard_id_info_list,omitempty"`
}

// InvoiceInstitutionDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionDeleteResponse struct {
	// Result 是删除结果是否成功
	Result *bool `json:"result,omitempty"`
}

// InvoiceInstitutionDetailInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionDetailInfoQueryResponse struct {
	// AdapterType 是若适用范围为EMPLOYEE_ALL，则表示制度对企业下全员生效，owner_id_list不返回；若适配范围为EMPLOYEE_SELECT，owner_id_list返回员工对应的id信息，返回的id类型通过owner_type区分；若适配范围为EMPLOYEE_DEPARTMENT，则表示对指定的部门生效，返回owner_id_list为部门id列表
	AdapterType *string `json:"adapter_type,omitempty"`
	// ConsultMode 是费控咨询模式
	ConsultMode *string `json:"consult_mode,omitempty"`
	// Currency 是制度下可发放币种
	Currency *string `json:"currency,omitempty"`
	// Effective 是制度启用停用状态
	Effective *string `json:"effective,omitempty"`
	// EffectiveEndDate 是制度生效结束时间
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是制度生效起始时间
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// ExpenseType 是费用类型
	ExpenseType *string `json:"expense_type,omitempty"`
	// InstitutionDesc 是制度描述
	InstitutionDesc *string `json:"institution_desc,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// InstitutionName 是制度名称
	InstitutionName *string `json:"institution_name,omitempty"`
	// IssueRuleInfoList 是发放规则信息列表
	IssueRuleInfoList *[]IssueRuleInfo `json:"issue_rule_info_list,omitempty"`
	// MultiEmployeeShareMode 是制度下额度是否支持多人共享。如果支持，则制度下资产所有人共享；如果不支持，则制度下的资产为个人资产，成员之间不共享。
	MultiEmployeeShareMode *string `json:"multi_employee_share_mode,omitempty"`
	// OuterSourceId 是外部唯一标识
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// OwnerIdList 是适配id列表
	OwnerIdList *[]string `json:"owner_id_list,omitempty"`
	// OwnerOpenIdList 是适配开放id列表
	OwnerOpenIdList *[]string `json:"owner_open_id_list,omitempty"`
	// OwnerType 是适配id类型
	OwnerType *string `json:"owner_type,omitempty"`
	// PaymentMode 是费控制度出资类型
	PaymentMode *string `json:"payment_mode,omitempty"`
	// SceneType 是因公场景
	SceneType *string `json:"scene_type,omitempty"`
	// StandardInfoDetailList 是使用规则列表
	StandardInfoDetailList *[]StandardInfo `json:"standard_info_detail_list,omitempty"`
	StandardInfoList       *StandardInfo   `json:"standard_info_list,omitempty"`
}

// InvoiceInstitutionExpenseRuleCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionExpenseRuleCreateResponse struct {
	// StandardId 是使用规则id
	StandardId *string `json:"standard_id,omitempty"`
}

// InvoiceInstitutionExpenseRuleDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionExpenseRuleDeleteResponse struct {
	// Result 是删除是否成功
	Result *bool `json:"result,omitempty"`
}

// InvoiceInstitutionExpenseRuleModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionExpenseRuleModifyResponse struct {
	// Result 是编辑结果是否成功
	Result *bool `json:"result,omitempty"`
}

// InvoiceInstitutionModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionModifyResponse struct {
	// IssueRuleIdInfoList 是新增发放规则时返回外部幂等id和发放规则id的映射关系
	IssueRuleIdInfoList *[]IssueRuleIdInfo `json:"issue_rule_id_info_list,omitempty"`
	// Result 是修改是否成功
	Result *bool `json:"result,omitempty"`
	// StandardIdInfoList 是新增使用规则后会返回外部幂等id和使用规则id的映射关系
	StandardIdInfoList *[]StandardIdInfo `json:"standard_id_info_list,omitempty"`
}

// InvoiceInstitutionPageInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionPageInfoQueryResponse struct {
	// InstitutionBasicInfoList 是制度基本信息列表
	InstitutionBasicInfoList *[]InstitutionBasicInfo `json:"institution_basic_info_list,omitempty"`
	// PageNum 是页码
	PageNum *string `json:"page_num,omitempty"`
	// PageSize 是页大小
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalPageCount 是总页数
	TotalPageCount *int64 `json:"total_page_count,omitempty"`
}

// InvoiceInstitutionScopeModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionScopeModifyResponse struct {
	// Result 是操作是否成功
	Result *bool `json:"result,omitempty"`
}

// InvoiceInstitutionScopePageInfoQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionScopePageInfoQueryResponse struct {
	// AdapterType 是制度适用范围类型
	AdapterType *string `json:"adapter_type,omitempty"`
	// OnwerOpenIdList 是适配开放id列表
	OnwerOpenIdList *[]string `json:"onwer_open_id_list,omitempty"`
	// OwnerIdList 是适配id列表
	OwnerIdList *[]string `json:"owner_id_list,omitempty"`
	// PageNum 是页码
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是页大小
	PageSize *int64 `json:"page_size,omitempty"`
	// TotalPageCount 是总页数
	TotalPageCount *int64 `json:"total_page_count,omitempty"`
}
