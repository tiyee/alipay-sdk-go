// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expensecontrol

// InvoiceExpenseControlAggregationCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlAggregationCreateRequest struct {
	// AccountId 是企业ID
	AccountId *string `json:"account_id,omitempty"`
	// AggregationName 是聚合关系名称	特殊说明：敏感词校验
	AggregationName *string `json:"aggregation_name,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// StandardIdList 是费控规则ID列表
	StandardIdList *[]string `json:"standard_id_list,omitempty"`
}

// InvoiceExpenseControlIssueBatchCancelRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlIssueBatchCancelRequest struct {
	// AccountId 是共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// IssueBatchId 是发放批次id
	IssueBatchId *string `json:"issue_batch_id,omitempty"`
}

// InvoiceExpenseControlIssueBatchCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlIssueBatchCreateRequest struct {
	// AccountId 是共同账户id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// BatchNo 是发放批次号，用于幂等校验本次发放
	BatchNo *string `json:"batch_no,omitempty"`
	// EffectiveEndDate 是额度有效结束时间
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是额度有效起始时间
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// InstitutionId 是制度id
	InstitutionId *string `json:"institution_id,omitempty"`
	// IssueDesc 是发放说明
	IssueDesc *string `json:"issue_desc,omitempty"`
	// IssueName 是发放名称
	IssueName *string `json:"issue_name,omitempty"`
	// IssueTargetInfoList 是员工发放信息列表
	IssueTargetInfoList *[]IssueTargetInfoContent `json:"issue_target_info_list,omitempty"`
	// QuotaType 是额度类型
	QuotaType *string `json:"quota_type,omitempty"`
	// ShareMode 是是否可转赠
	ShareMode *string `json:"share_mode,omitempty"`
}

// InvoiceExpenseControlQuotaCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlQuotaCreateRequest struct {
	// AccountId 是共同账号id（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EffectiveEndDate 是额度失效时间（格式：yyyy-MM-dd HH:mm:ss）
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是额度生效时间（格式：yyyy-MM-dd HH:mm:ss）
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// IssueDesc 是批量发放时，可填写该字段
	IssueDesc *string `json:"issue_desc,omitempty"`
	// IssueName 是批量发放时，可填写该字段作为发放名称，如果未填写取outer_source_id为默认值
	IssueName *string `json:"issue_name,omitempty"`
	// IssueQuotaTargetList 是创建额度发放明细列表
	IssueQuotaTargetList *[]IssueTargetInfoContent `json:"issue_quota_target_list,omitempty"`
	// OuterSourceId 是外部操作幂等ID，标识创建额度的唯一性，防止重复创建
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// OwnerId 是额度所属者ID（未切换open_id请使用此字段）：(字段升级，请使用issue_target_info_list中owner_id字段)
	OwnerId *string `json:"owner_id,omitempty"`
	// OwnerOpenId 是额度所属者ID（切换open_id后请使用此字段）：(字段升级，请使用issue_quota_target_list中owner_open_id)
	OwnerOpenId *string `json:"owner_open_id,omitempty"`
	// OwnerType 是额度所属者类型(字段升级，请使用issue_quota_target_list中owner_type字段)
	OwnerType *string `json:"owner_type,omitempty"`
	// Platform 是外部平台编码（通常为接入方大写英文缩写）(历史版本字段，不推荐使用)
	Platform *string `json:"platform,omitempty"`
	// QuotaType 是创建额度类型
	QuotaType *string `json:"quota_type,omitempty"`
	// QuotaValue 是额度值，以（分）为单位(字段升级，请使用issue_quota_target_list中issue_quota字段)
	QuotaValue *string `json:"quota_value,omitempty"`
	// ShareMode 是0:不可转赠	1:可以转增
	ShareMode *string `json:"share_mode,omitempty"`
	// TargetId 是额度维度ID
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是额度维度	枚举值：	INSTITUTION（制度维度），	EXPENSE_TYPE（费用类型维度）
	TargetType *string `json:"target_type,omitempty"`
}

// InvoiceExpenseControlQuotaModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlQuotaModifyRequest struct {
	// AccountId 是企业共同账户ID（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AccountId *string `json:"account_id,omitempty"`
	// Action 是变更模式
	Action *string `json:"action,omitempty"`
	// AgreementNo 是授权签约协议号（该字段将废弃，不建议使用，可用enterprise_id字段替换）(该字段将废弃，不建议使用，可用enterprise_id字段替换)
	AgreementNo *string `json:"agreement_no,omitempty"`
	// AllowNegative 是修改余额对应的额度时是否允许修改为负数，true表示允许，false表示不允许。默认允许
	AllowNegative *bool `json:"allow_negative,omitempty"`
	// Amount 是变更的金额/余额，日元或者韩元以（元）为单位，其他币种以（分）为单位，当变更的资产类型为次卡时，单位为次
	Amount *string `json:"amount,omitempty"`
	// EnterpriseId 是企业ID
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// OuterSourceId 是外部操作幂等ID（接入方接口调用幂等控制ID）
	OuterSourceId *string `json:"outer_source_id,omitempty"`
	// Platform 是外部平台编码(历史字段，不推荐使用)
	Platform *string `json:"platform,omitempty"`
	// QuotaId 是额度ID
	QuotaId *string `json:"quota_id,omitempty"`
	// ShareMode 是是否可转赠，1表示可转赠，0表示不可转赠
	ShareMode *string `json:"share_mode,omitempty"`
}

// InvoiceExpenseControlQuotaQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseControlQuotaQueryRequest struct {
	// AccountId 是企业共同账户ID
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EnterpriseId 是企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// OwnerId 是额度所属者ID
	OwnerId *string `json:"owner_id,omitempty"`
	// OwnerOpenId 是额度所属者开放ID
	OwnerOpenId *string `json:"owner_open_id,omitempty"`
	// OwnerType 是额度所属者类型
	OwnerType *string `json:"owner_type,omitempty"`
	// PageNum 是页码
	PageNum *int64 `json:"page_num,omitempty"`
	// PageSize 是每页条数
	PageSize *int64 `json:"page_size,omitempty"`
	// QuotaIdList 是额度ID列表
	QuotaIdList *[]string `json:"quota_id_list,omitempty"`
	// QuotaType 是额度类型
	QuotaType *string `json:"quota_type,omitempty"`
	// TargetId 是额度维度ID
	TargetId *string `json:"target_id,omitempty"`
	// TargetType 是额度维度
	TargetType *string `json:"target_type,omitempty"`
}
