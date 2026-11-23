// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expensecontrol

// ExpenseQuotaInfo 是支付宝官方 V3 OpenAPI 模型。
type ExpenseQuotaInfo struct {
	// Currency 是额度对应的币种
	Currency *string `json:"currency,omitempty"`
	// EffectiveEndDate 是额度失效时间
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是额度生效时间
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// Freeze 是额度是否冻结，冻结后因公付不可用
	Freeze *bool `json:"freeze,omitempty"`
	// OwnerId 是额度所属者ID	owner_type为EMPLOYEE时为员工支付宝ID	owner_type为ENTERPRISE_PAY_UID时为员工企业码ID	owner_type为PHONE时为员工手机号	owner_type为ENTERPRISE时为企业ID
	OwnerId *string `json:"owner_id,omitempty"`
	// OwnerOpenId 是owner_type为EMPLOYEE时为员工open_id owner_type为PHONE时为员工手机号 owner_type为ENTERPRISE_PAY_UID时为员工企业码ID
	OwnerOpenId *string `json:"owner_open_id,omitempty"`
	// OwnerType 是额度所属者id类型
	OwnerType *string `json:"owner_type,omitempty"`
	// QuotaAvailable 是可用金额（单位分）
	QuotaAvailable *string `json:"quota_available,omitempty"`
	// QuotaId 是额度ID
	QuotaId *string `json:"quota_id,omitempty"`
	// QuotaLocked 是锁定金额（单位分）
	QuotaLocked *string `json:"quota_locked,omitempty"`
	// QuotaTotal 是总金额（单位分）
	QuotaTotal *string `json:"quota_total,omitempty"`
	// QuotaType 是额度类型
	QuotaType *string `json:"quota_type,omitempty"`
	// QuotaUsed 是已用金额（单位分）
	QuotaUsed *string `json:"quota_used,omitempty"`
	// TargetId 是额度维度ID	当 target_type=EXPENSE_TYPE 时，值为	MEAL（工作餐）	当target_type=RULE_GROUP_AGGREGATION 时，值为费控规则聚合ID	当target_type=INSTITUTION 时，值为制度ID
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是额度维度
	TargetType *string `json:"target_type,omitempty"`
}

// IssueQuotaCheckInfo 是支付宝官方 V3 OpenAPI 模型。
type IssueQuotaCheckInfo struct {
	// IssueQuota 是发放点券和余额时，单位为元，发放次卡时，单位为次
	IssueQuota *string `json:"issue_quota,omitempty"`
	// Message 是校验结果的原因
	Message *string `json:"message,omitempty"`
	// OwnerId 是入参的 owner_id
	OwnerId *string `json:"owner_id,omitempty"`
	// OwnerOpenId 是入参的  owner_open_id
	OwnerOpenId *string `json:"owner_open_id,omitempty"`
	// OwnerType 是owner_type
	OwnerType *string `json:"owner_type,omitempty"`
	// Result 是校验结果
	Result *bool `json:"result,omitempty"`
	// UserName 是用户姓名
	UserName *string `json:"user_name,omitempty"`
}

// IssueTargetInfoContent 是支付宝官方 V3 OpenAPI 模型。
type IssueTargetInfoContent struct {
	// IssueQuota 是当发放余额或点券时，单位为元，币种以创建制度时填写的currency为准，不填写默认为人民币。当发放类型为次卡时，单位为次。
	IssueQuota *string `json:"issue_quota,omitempty"`
	// OwnerId 是切换open_id前请使用：owner类型为PHONE时该字段表示员工手机号，owner类型为EMPLOYEE时该字段表示为员工支付宝uid，owner类型为ENTERPRISE_PAY_UID时该字段表示为员工企业码id，owner类型为ENTERPRISE时可不填，如果填写仅支持填企业id，owner类型为INSTITUTION时可不填，默认为制度id
	OwnerId *string `json:"owner_id,omitempty"`
	// OwnerOpenId 是切换open_id后请使用：owner类型为PHONE时，填写员工手机号，EMPLOYEE时填写员工open_id
	OwnerOpenId *string `json:"owner_open_id,omitempty"`
	// OwnerType 是owner类型，用来表示owner_id的类型
	OwnerType *string `json:"owner_type,omitempty"`
	// UserName 是员工姓名
	UserName *string `json:"user_name,omitempty"`
}
