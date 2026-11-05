// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package issuerule

// InvoiceIssueRuleCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceIssueRuleCreateRequest struct {
	// AccountId 是共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// Effective 是是否启用，"0"为停用，"1"为启用；	如果target_type指定INSTITUTION，则与制度一致,无需填写
	Effective *string `json:"effective,omitempty"`
	// EffectivePeriod 是生效周期，时间不限格式："{\"all\":true}"
	EffectivePeriod *string `json:"effective_period,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InvalidMode 是是否可累计：0（不可累计）、1（可累计），2（累计天数）、3（累计到指定日期）余额只允许可累计，按城市发放只允许不可累计，设置了生效周期只允许不可累计
	InvalidMode *int64 `json:"invalid_mode,omitempty"`
	// InvalidModeValue 是累计值（预留字段）
	InvalidModeValue *string `json:"invalid_mode_value,omitempty"`
	// IssueAmountValue 是发放金额，单位（元）；	按城市不同发放标准示例值："[{\"cities\":[\"321000\",\"321100\"],\"amount\":\"50\"},{\"cities\":[\"321200\",\"321300\"],\"amount\":\"150\"}]"	统一标准发放示例值："200"
	IssueAmountValue *string `json:"issue_amount_value,omitempty"`
	// IssueEndDate 是发放规则的截止时间	如果target_type指定INSTITUTION，则使用制度有效期结束时间，无需填写
	IssueEndDate *string `json:"issue_end_date,omitempty"`
	// IssueRuleName 是发放规则名称
	IssueRuleName *string `json:"issue_rule_name,omitempty"`
	// IssueStartDate 是发放规则的开始时间	如果target_type指定INSTITUTION，则使用制度有效期起始时间，无需填写
	IssueStartDate *string `json:"issue_start_date,omitempty"`
	// IssueType 是发放类型
	IssueType *string `json:"issue_type,omitempty"`
	// OuterSourceId 是填写该字段可用于创建幂等，防止重复创建
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// QuotaType 是额度类型，枚举：COUPON（点券）、CAP（余额）
	QuotaType *string `json:"quota_type,omitempty"`
	// ShareMode 是是否可转赠：0（不可转赠）、1（可转赠）
	ShareMode *int64 `json:"share_mode,omitempty"`
	// TargetId 是发放规则关联的目标值	target_type为制度时：制度id
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是发放规则关联的目标类型
	TargetType *string `json:"target_type,omitempty"`
}

// InvoiceIssueRuleDeleteRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceIssueRuleDeleteRequest struct {
	// AccountId 是共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// IssueRuleIdList 是需要删除的发放规则id列表
	IssueRuleIdList *[]string `json:"issue_rule_id_list,omitempty"`
	// TargetId 是目标值id
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是发放规则关联的目标类型
	TargetType *string `json:"target_type,omitempty"`
}

// InvoiceIssueRuleModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceIssueRuleModifyRequest struct {
	// AccountId 是共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// Action 是修改类型，枚举：	MODIFY_BASIC_INFO（修改发放规则基本信息）
	Action *string `json:"action,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// Effective 是是否启用，"0"为停用，"1"为启用；
	Effective *string `json:"effective,omitempty"`
	// EffectivePeriod 是生效周期（不传则不修改），不限："{\"all\":true}"
	EffectivePeriod *string `json:"effective_period,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InvalidMode 是是否可累计（不传则不修改）可选值：0（不可累计）、1（可累计），2（累计天数），3（累计到固定时间）
	InvalidMode *int64 `json:"invalid_mode,omitempty"`
	// InvalidModeValue 是累计值
	InvalidModeValue *string `json:"invalid_mode_value,omitempty"`
	// IssueAmountValue 是发放金额（不传则不修改），单位（元）
	IssueAmountValue *string `json:"issue_amount_value,omitempty"`
	// IssueEndDate 是发放规则的截止时间
	IssueEndDate *string `json:"issue_end_date,omitempty"`
	// IssueRuleId 是发放规则id
	IssueRuleId *string `json:"issue_rule_id,omitempty"`
	// IssueRuleName 是发放规则名称（不传则不修改）
	IssueRuleName *string `json:"issue_rule_name,omitempty"`
	// IssueStartDate 是发放规则的开始时间
	IssueStartDate *string `json:"issue_start_date,omitempty"`
	// IssueType 是发放类型
	IssueType *string `json:"issue_type,omitempty"`
	// QuotaType 是额度类型
	QuotaType *string `json:"quota_type,omitempty"`
	// ShareMode 是是否可转赠（不传则不修改），可选值：0（不可转赠）、1（可转赠），默认为0，按城市发放不支持可转赠
	ShareMode *int64 `json:"share_mode,omitempty"`
	// TargetId 是发放规则关联的目标值 target_type为制度时：制度id
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是发放规则关联的目标类型
	TargetType *string `json:"target_type,omitempty"`
}
