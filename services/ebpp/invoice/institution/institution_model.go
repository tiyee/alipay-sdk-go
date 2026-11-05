// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package institution

// AssetShareSourceInfo 是支付宝官方 V3 OpenAPI 模型。
type AssetShareSourceInfo struct {
	// ShareMode 是当前规则可用的资产来源类型，搭配source_id_list使用。
	ShareMode *string `json:"share_mode,omitempty"`
	// SourceIdList 是资产共享来源id列表，与share_mode配合设置，如果share_mode为INSTITUTION，该值则填写制度id列表
	SourceIdList *[]string `json:"source_id_list,omitempty"`
}

// ExpenseCtrRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type ExpenseCtrRuleInfo struct {
	// RuleFactor 是费控维度
	RuleFactor *string `json:"rule_factor,omitempty"`
	// RuleId 是费控条件ID
	RuleId *string `json:"rule_id,omitempty"`
	// RuleName 是费控条件名称
	RuleName *string `json:"rule_name,omitempty"`
	// RuleOperator 是费控条件操作符	枚举值：	LT("=","大于等于")	GT(">","大于")
	RuleOperator *string `json:"rule_operator,omitempty"`
	// RuleValue 是费控条件值
	RuleValue *string `json:"rule_value,omitempty"`
}

// InstitutionBasicInfo 是支付宝官方 V3 OpenAPI 模型。
type InstitutionBasicInfo struct {
	// ConsultMode 是费控咨询模式
	ConsultMode *string `json:"consult_mode,omitempty"`
	// Currency 是制度可发放资产币种
	Currency *string `json:"currency,omitempty"`
	// Effective 是制度是否启用
	Effective *string `json:"effective,omitempty"`
	// EffectiveEndDate 是制度生效结束时间
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是制度生效起始时间
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// InstitutionDesc 是制度描述
	InstitutionDesc *string `json:"institution_desc,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// InstitutionName 是制度名称
	InstitutionName *string `json:"institution_name,omitempty"`
	// MultiEmployeeShareMode 是制度下额度是否支持多人共享。如果支持，则制度下资产所有人共享；如果不支持，则制度下的资产为个人资产，成员之间不共享。
	MultiEmployeeShareMode *string `json:"multi_employee_share_mode,omitempty"`
}

// InstitutionScopeInfo 是支付宝官方 V3 OpenAPI 模型。
type InstitutionScopeInfo struct {
	// AdapterType 是制度适用范围，支持配置全员、指定员工和指定部门
	AdapterType *string `json:"adapter_type,omitempty"`
	// OwnerIdList 是适用范围归属id列表
	OwnerIdList *[]string `json:"owner_id_list,omitempty"`
	// OwnerType 是员工id类型
	OwnerType *string `json:"owner_type,omitempty"`
}

// IssueRuleIdInfo 是支付宝官方 V3 OpenAPI 模型。
type IssueRuleIdInfo struct {
	// IssueRuleId 是发放规则id
	IssueRuleId *string `json:"issue_rule_id,omitempty"`
	// OuterSourceId 是外部发放规则id
	OuterSourceId *string `json:"outer_source_id,omitempty"`
}

// IssueRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type IssueRuleInfo struct {
	// EffectivePeriod 是生效时间段
	EffectivePeriod *string `json:"effective_period,omitempty"`
	// InvalidMode 是累计类型，默认为0	可选值：0（不可累计）、1（可累计）、2（累计天数）、3（累计到指定日期）
	InvalidMode *int64 `json:"invalid_mode,omitempty"`
	// InvalidModeValue 是累计类型值
	InvalidModeValue *string `json:"invalid_mode_value,omitempty"`
	// IssueAmountValue 是发放金额，单位元
	IssueAmountValue *string `json:"issue_amount_value,omitempty"`
	// IssueEndDate 是发放规则有效结束时间
	IssueEndDate *string `json:"issue_end_date,omitempty"`
	// IssueRuleId 是发放规则id
	IssueRuleId *string `json:"issue_rule_id,omitempty"`
	// IssueRuleName 是发放规则名称
	IssueRuleName *string `json:"issue_rule_name,omitempty"`
	// IssueStartDate 是发放规则有效起始时间
	IssueStartDate *string `json:"issue_start_date,omitempty"`
	// IssueType 是发放类型
	IssueType *string `json:"issue_type,omitempty"`
	// OuterSourceId 是外部发放规则id
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// QuotaType 是额度类型
	QuotaType *string `json:"quota_type,omitempty"`
	// ShareMode 是是否可转赠
	ShareMode *int64 `json:"share_mode,omitempty"`
	// TargetId 是目标id
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是发放规则归属的目标类型
	TargetType *string `json:"target_type,omitempty"`
}

// ModifyIssueRuleDetailInfo 是支付宝官方 V3 OpenAPI 模型。
type ModifyIssueRuleDetailInfo struct {
	// AddIssueRuleList 是新增发放规则列表
	AddIssueRuleList *[]IssueRuleInfo `json:"add_issue_rule_list,omitempty"`
	// DeleteIssueRuleIdList 是删除的发放规则id列表
	DeleteIssueRuleIdList *[]string            `json:"delete_issue_rule_id_list,omitempty"`
	ModifyIssueRuleList   *ModifyIssueRuleInfo `json:"modify_issue_rule_list,omitempty"`
}

// ModifyIssueRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type ModifyIssueRuleInfo struct {
	// EffectivePeriod 是发放资产的有效时间段，该字段为空时，默认为不限制。当发放资产为余额时，仅支持设置为不限，当发放资产为点券时，可支持设置多种值。具体传值格式请参考文档 接入文档
	EffectivePeriod *string `json:"effective_period,omitempty"`
	// InvalidMode 是累计类型，默认为0 可选值：0（不可累计）、1（可累计）、2（累计天数）、3（累计到指定日期）
	InvalidMode *int64 `json:"invalid_mode,omitempty"`
	// InvalidModeValue 是累计类型值
	InvalidModeValue *string `json:"invalid_mode_value,omitempty"`
	// IssueAmountValue 是发放金额（不传则不修改），单位（元）
	IssueAmountValue *string `json:"issue_amount_value,omitempty"`
	// IssueRuleId 是发放规则id
	IssueRuleId *string `json:"issue_rule_id,omitempty"`
	// IssueRuleName 是发放规则名称
	IssueRuleName *string `json:"issue_rule_name,omitempty"`
	// IssueType 是额度发放周期
	IssueType *string `json:"issue_type,omitempty"`
	// QuotaType 是发放规则发放的额度类型
	QuotaType *string `json:"quota_type,omitempty"`
	// ShareMode 是是否可转赠（不传则不修改），可选值：0（不可转赠）、1（可转赠），默认为0。该字段决定发放规则发放的资产是否可转赠给同制度下的其他人员
	ShareMode *int64 `json:"share_mode,omitempty"`
}

// ModifyScopeInfo 是支付宝官方 V3 OpenAPI 模型。
type ModifyScopeInfo struct {
	// AdapterType 是制度的适用范围类型
	AdapterType *string `json:"adapter_type,omitempty"`
	// AddOwnerIdList 是增加的归属id列表
	AddOwnerIdList *[]string `json:"add_owner_id_list,omitempty"`
	// DeleteOwnerIdList 是删除归属id列表
	DeleteOwnerIdList *[]string `json:"delete_owner_id_list,omitempty"`
	// OwnerType 是归属id类型
	OwnerType *string `json:"owner_type,omitempty"`
}

// ModifyStandardDetailInfo 是支付宝官方 V3 OpenAPI 模型。
type ModifyStandardDetailInfo struct {
	// AddStandardList 是新增的使用规则列表
	AddStandardList *[]StandardInfo `json:"add_standard_list,omitempty"`
	// DeleteStandardIdList 是待删除的使用规则id列表
	DeleteStandardIdList *[]string `json:"delete_standard_id_list,omitempty"`
	// ModifyStandardList 是待修改的使用规则信息列表
	ModifyStandardList *[]ModifyStandardInfo `json:"modify_standard_list,omitempty"`
}

// ModifyStandardInfo 是支付宝官方 V3 OpenAPI 模型。
type ModifyStandardInfo struct {
	// AddConditionList 是要添加的条件列表
	AddConditionList     *[]StandardConditionInfo `json:"add_condition_list,omitempty"`
	AssetShareSourceInfo *AssetShareSourceInfo    `json:"asset_share_source_info,omitempty"`
	// ConsumeMode 是消费模式
	ConsumeMode *string `json:"consume_mode,omitempty"`
	// DeleteConditionIdList 是待删除的条件id列表
	DeleteConditionIdList *[]string `json:"delete_condition_id_list,omitempty"`
	// ModifyConditionList 是修改的条件列表
	ModifyConditionList *[]StandardConditionInfo `json:"modify_condition_list,omitempty"`
	// OpenRuleId 是使用规则绑定的开票规则id
	OpenRuleId *string `json:"open_rule_id,omitempty"`
	// PaymentPolicy 是支付策略 当笔消费金额大于规则可用余额时，用于控制支付策略。COMBINATION表示支持因公资产和个人资产组合支付，PERSONAL表示整单个人支付。
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// PaymentTimeMode 是支付时间模式，用于控制支付时的时间。该字段缺省时采用“支付时间或预计送达时间”，两个时间只要有一个命中即可使用因公付
	PaymentTimeMode *string `json:"payment_time_mode,omitempty"`
	// PersonalQrcodeMode 是个人收款码转账是否支持因公付。可选值：0（不支持）、1（支持）
	PersonalQrcodeMode *int64 `json:"personal_qrcode_mode,omitempty"`
	// StandardDesc 是使用规则描述
	StandardDesc *string `json:"standard_desc,omitempty"`
	// StandardId 是修改的使用规则id
	StandardId *string `json:"standard_id,omitempty"`
	// StandardName 是使用规则名称
	StandardName *string `json:"standard_name,omitempty"`
}

// StandardConditionInfo 是支付宝官方 V3 OpenAPI 模型。
type StandardConditionInfo struct {
	// RuleFactor 是条件类型
	RuleFactor *string `json:"rule_factor,omitempty"`
	// RuleId 是费控条件ID
	RuleId *string `json:"rule_id,omitempty"`
	// RuleName 是费控条件名称
	RuleName *string `json:"rule_name,omitempty"`
	// RuleOperator 是费控条件操作符 枚举值： LT("=","大于等于") GT(">","大于")(该字段为保留字段，目前未被使用)
	RuleOperator *string `json:"rule_operator,omitempty"`
	// RuleValue 是费控条件值
	RuleValue *string `json:"rule_value,omitempty"`
}

// StandardIdInfo 是支付宝官方 V3 OpenAPI 模型。
type StandardIdInfo struct {
	// OuterSourceId 是外部使用规则id
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// StandardId 是使用规则id
	StandardId *string `json:"standard_id,omitempty"`
}

// StandardInfo 是支付宝官方 V3 OpenAPI 模型。
type StandardInfo struct {
	AssetShareSourceInfo *AssetShareSourceInfo `json:"asset_share_source_info,omitempty"`
	// ConsumeMode 是消费模式，不填为默认模式，枚举值：COUPON_ONLY（仅支持点券）	COUPON_AND_CAP（支持点券+余额）	COUNT（仅支持次卡）	DEFAULT（默认模式）	点券：消费时找员工的点券，没有或者用完了不可付；	点券+余额：消费时找员工的点券，没有找员工的余额，没有或者用完了不可付；	次卡：消费时找员工的次卡，没有或者用完了不可付；	默认：有给员工设置员工余额以员工余额为准，用完为止。否则只受规则里的限额和企业账户资金上限管控；
	ConsumeMode *string `json:"consume_mode,omitempty"`
	// ExpenseTypeSubCategory 是费用类型子类
	ExpenseTypeSubCategory *string `json:"expense_type_sub_category,omitempty"`
	// OpenRuleId 是开票规则id，可通过接口alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.create 创建并得到开票规则ID
	OpenRuleId *string `json:"open_rule_id,omitempty"`
	// OuterSourceId 是外部使用规则id，制度内使用规则该字段不允许重复
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// PaymentPolicy 是支付策略 当笔消费金额大于规则可用余额时，用于控制支付策略，该字段缺省时采取因公账户和个人账户组合支付策略
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// PaymentTimeMode 是支付时间模式，用于控制支付时的时间。该字段缺省时采用“支付时间或预计送达时间”，两个时间只要有一个命中即可使用因公付
	PaymentTimeMode *string `json:"payment_time_mode,omitempty"`
	// PersonalQrcodeMode 是个人收款码转账是否支持因公付，默认为0。可选值：0（不支持）、1（支持）
	PersonalQrcodeMode *int64 `json:"personal_qrcode_mode,omitempty"`
	// StandardConditionInfoList 是使用规则条件列表
	StandardConditionInfoList *[]StandardConditionInfo `json:"standard_condition_info_list,omitempty"`
	// StandardDesc 是使用规则描述(敏感词校验)
	StandardDesc *string `json:"standard_desc,omitempty"`
	// StandardId 是制度ID（创建使用规则时非必填）
	StandardId *string `json:"standard_id,omitempty"`
	// StandardName 是规则名称
	StandardName *string `json:"standard_name,omitempty"`
}
