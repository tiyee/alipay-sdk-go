// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package institution

// InvoiceInstitutionCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionCreateRequest struct {
	// AccountId 是企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// AdapterType 是制度适配范围(请使用institution_scope_info中adapter_type)
	AdapterType *string `json:"adapter_type,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// ConsultMode 是费控咨询模式，默认为0。
	ConsultMode *string `json:"consult_mode,omitempty"`
	// Currency 是该字段用于管控制度下发放的资产币种，不填写默认为人民币，设置该字段后，制度下手工发放的资产币种均为设置值。
	Currency *string `json:"currency,omitempty"`
	// EffectiveEndDate 是制度生效结束时间
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是制度生效起始时间
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// ExpenseType 是费用类型
	ExpenseType *string `json:"expense_type,omitempty"`
	// InstitutionDesc 是制度描述
	InstitutionDesc *string `json:"institution_desc,omitempty"`
	// InstitutionName 是制度名称
	InstitutionName      *string               `json:"institution_name,omitempty"`
	InstitutionScopeInfo *InstitutionScopeInfo `json:"institution_scope_info,omitempty"`
	// IssueRuleInfoList 是发放规则列表
	IssueRuleInfoList *[]IssueRuleInfo `json:"issue_rule_info_list,omitempty"`
	// MultiEmployeeShareMode 是制度下额度是否支持多人共享。如果支持，则制度下资产所有人共享；如果不支持，则制度下的资产为个人资产，成员之间不共享。
	MultiEmployeeShareMode *string `json:"multi_employee_share_mode,omitempty"`
	// OuterSourceId 是外部唯一标识，可用于防止重复创建
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// OwnerIdList 是适配id列表(请使用institution_scope_info中owner_id_list)
	OwnerIdList *[]string `json:"owner_id_list,omitempty"`
	// OwnerOpenIdList 是适配开放id列表(仅存量使用员工支付宝id配置适配范围的场景需要该字段，新对接方案不推荐使用员工支付宝id)
	OwnerOpenIdList *[]string `json:"owner_open_id_list,omitempty"`
	// OwnerType 是适配id类型(请使用institution_scope_info中owner_type)
	OwnerType *string `json:"owner_type,omitempty"`
	// PaymentMode 是费控制度的出资类型
	PaymentMode *string `json:"payment_mode,omitempty"`
	// SceneType 是因公场景(该字段将废弃，不建议使用)
	SceneType *string `json:"scene_type,omitempty"`
	// StandardInfoList 是使用规则列表
	StandardInfoList *[]StandardInfo `json:"standard_info_list,omitempty"`
}

// InvoiceInstitutionExpenseRuleCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionExpenseRuleCreateRequest struct {
	// AccountId 是企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo          *string               `json:"agreement_no,omitempty"`
	AssetShareSourceInfo *AssetShareSourceInfo `json:"asset_share_source_info,omitempty"`
	// ConsumeMode 是消费模式
	ConsumeMode *string `json:"consume_mode,omitempty"`
	// EnterpriseId 是企业码id
	EnterpriseId            *string             `json:"enterprise_id,omitempty"`
	ExpenseCtrlRuleInfoList *ExpenseCtrRuleInfo `json:"expense_ctrl_rule_info_list,omitempty"`
	// ExpenseTypeSubCategory 是费用类型子类
	ExpenseTypeSubCategory *string `json:"expense_type_sub_category,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// OpenRuleId 是开票规则id
	OpenRuleId *string `json:"open_rule_id,omitempty"`
	// OuterSourceId 是外部唯一标识，填写该字段可用于创建幂等，防止重复创建
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// PaymentPolicy 是支付策略
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// PaymentTimeMode 是支付时间模式，用于控制支付时的时间。该字段缺省时采用“支付时间或预计送达时间”，两个时间只要有一个命中即可使用因公付
	PaymentTimeMode *string `json:"payment_time_mode,omitempty"`
	// PersonalQrcodeMode 是个人收款码转账是否支持因公付，默认为0。可选值：0（不支持）、1（支持）
	PersonalQrcodeMode *int64 `json:"personal_qrcode_mode,omitempty"`
	// StandardConditionInfoList 是使用规则因子列表
	StandardConditionInfoList *[]StandardConditionInfo `json:"standard_condition_info_list,omitempty"`
	// StandardDesc 是使用规则描述
	StandardDesc *string `json:"standard_desc,omitempty"`
	// StandardName 是费控规则名称
	StandardName *string `json:"standard_name,omitempty"`
}

// InvoiceInstitutionExpenseRuleDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionExpenseRuleDeleteRequest struct {
	// AccountId 是企业共同账户id（alipay.ebpp.invoice.institution.expenserule.delete）(alipay.ebpp.invoice.institution.expenserule.delete)
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// StandardIdList 是删除的使用规则id列表
	StandardIdList *[]string `json:"standard_id_list,omitempty"`
}

// InvoiceInstitutionExpenseRuleModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionExpenseRuleModifyRequest struct {
	// AccountId 是企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// Action 是修改使用规则
	Action *string `json:"action,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo          *string               `json:"agreement_no,omitempty"`
	AssetShareSourceInfo *AssetShareSourceInfo `json:"asset_share_source_info,omitempty"`
	// ConsumeMode 是该使用规则支持的资产消费模式，不填写则为默认模式，默认模式下有余额时使用余额，没有余额则使用规则中的限额，点券模式为只能使用点券，点券+余额模式为可以使用点券和余额。
	ConsumeMode *string `json:"consume_mode,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// ExpenseCtrlRuleInfoList 是使用规则条件列表（已废弃）
	ExpenseCtrlRuleInfoList *[]ExpenseCtrRuleInfo `json:"expense_ctrl_rule_info_list,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// OpenRuleId 是开票规则id
	OpenRuleId *string `json:"open_rule_id,omitempty"`
	// PaymentPolicy 是当笔消费金额大于规则可用余额时，用于控制支付策略
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// PaymentTimeMode 是支付时间模式，用于控制支付时的时间。该字段缺省时采用“支付时间或预计送达时间”，两个时间只要有一个命中即可使用因公付
	PaymentTimeMode *string `json:"payment_time_mode,omitempty"`
	// PersonalQrcodeMode 是个人收款码转账是否支持因公付，默认为0。可选值：0（不支持）、1（支持）
	PersonalQrcodeMode *int64 `json:"personal_qrcode_mode,omitempty"`
	// StandardConditionInfoList 是规则条件列表
	StandardConditionInfoList *[]StandardConditionInfo `json:"standard_condition_info_list,omitempty"`
	// StandardDesc 是规则描述
	StandardDesc *string `json:"standard_desc,omitempty"`
	// StandardId 是使用规则id
	StandardId *string `json:"standard_id,omitempty"`
	// StandardName 是规则名称
	StandardName *string `json:"standard_name,omitempty"`
}

// InvoiceInstitutionModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionModifyRequest struct {
	// AccountId 是企业共同账户id	（该字段将废弃，不建议使用，可用enterprise_id字段替换）(推荐使用enterprise_id)
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号	（该字段将废弃，不建议使用，可用enterprise_id字段替换）(推荐使用enterprise_id)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// ConsultMode 是费控咨询模式：	0-支付宝内部计算规则可用金额	1-咨询外部服务商规则可用金额
	ConsultMode *string `json:"consult_mode,omitempty"`
	// Effective 是制度是否启用
	Effective *string `json:"effective,omitempty"`
	// EffectiveEndDate 是制度生效结束时间（可空），最小粒度为天，最早是当日
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是制度生效起始时间
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InstitutionDesc 是制度描述
	InstitutionDesc *string `json:"institution_desc,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// InstitutionName 是制度名称
	InstitutionName           *string                    `json:"institution_name,omitempty"`
	ModifyIssueRuleDetailInfo *ModifyIssueRuleDetailInfo `json:"modify_issue_rule_detail_info,omitempty"`
	ModifyScopeInfo           *ModifyScopeInfo           `json:"modify_scope_info,omitempty"`
	ModifyStandardDetailInfo  *ModifyStandardDetailInfo  `json:"modify_standard_detail_info,omitempty"`
}

// InvoiceInstitutionScopeModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceInstitutionScopeModifyRequest struct {
	// AccountId 是企业共同账户id
	AccountId *string `json:"account_id,omitempty"`
	// AdapterType 是制度适用范围类型
	AdapterType *string `json:"adapter_type,omitempty"`
	// AddOwnerIdList 是增加适配id列表
	AddOwnerIdList *[]string `json:"add_owner_id_list,omitempty"`
	// AddOwnerOpenIdList 是增加适配开放id列表
	AddOwnerOpenIdList *[]string `json:"add_owner_open_id_list,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// OwnerType 是归属类型
	OwnerType *string `json:"owner_type,omitempty"`
	// RemoveOwnerIdList 是删除适配id列表
	RemoveOwnerIdList *[]string `json:"remove_owner_id_list,omitempty"`
	// RemoveOwnerOpenIdList 是删除适配开放id列表
	RemoveOwnerOpenIdList *[]string `json:"remove_owner_open_id_list,omitempty"`
}

// InvoiceInstitutionDeleteRequest is the query request model for alipay.ebpp.invoice.institution.delete.
type InvoiceInstitutionDeleteRequest struct {
	// AccountId 是企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号	（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty" query:"institution_id,omitempty"`
}

// InvoiceInstitutionDetailInfoQueryRequest is the query request model for alipay.ebpp.invoice.institution.detailinfo.query.
type InvoiceInstitutionDetailInfoQueryRequest struct {
	// AccountId 是企业共同账户id
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty" query:"institution_id,omitempty"`
	// OuterSourceId 是外部唯一标识
	OuterSourceId *string `json:"outer_source_id,omitempty" query:"outer_source_id,omitempty"`
	// OwnerType 是适配id类型
	OwnerType *string `json:"owner_type,omitempty" query:"owner_type,omitempty"`
}

// InvoiceInstitutionPageInfoQueryRequest is the query request model for alipay.ebpp.invoice.institution.pageinfo.query.
type InvoiceInstitutionPageInfoQueryRequest struct {
	// AccountId 是企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// PageNum 是页码
	PageNum *string `json:"page_num,omitempty" query:"page_num,omitempty"`
	// PageSize 是页大小
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
}

// InvoiceInstitutionScopePageInfoQueryRequest is the query request model for alipay.ebpp.invoice.institution.scopepageinfo.query.
type InvoiceInstitutionScopePageInfoQueryRequest struct {
	// AccountId 是企业共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty" query:"institution_id,omitempty"`
	// OwnerType 是适配id类型
	OwnerType *string `json:"owner_type,omitempty" query:"owner_type,omitempty"`
	// PageNum 是页码
	PageNum *string `json:"page_num,omitempty" query:"page_num,omitempty"`
	// PageSize 是页大小
	PageSize *string `json:"page_size,omitempty" query:"page_size,omitempty"`
}
