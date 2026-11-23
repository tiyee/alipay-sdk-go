package scene

import (
	"github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/expenserules"
)

// InvoiceExpenseRulesSceneRuleCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesSceneRuleCreateRequest struct {
	// AccountId 是共同账号ID
	AccountId *string `json:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EffectiveEndDate 是有效期截止
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是有效期起始
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// EmployeeList 是切换open_id前请使用：员工支付宝uid列表	特殊说明：单次传入的最大员工数为50，后续通过修改费控规则员工接口进行员工调整
	EmployeeList *[]string `json:"employee_list,omitempty"`
	// EmployeeOpenIdList 是切换open_id后请使用：员工open_id/企业码员工ID列表 特殊说明：单次传入的最大员工数为50，后续通过修改费控规则员工接口进行员工调整
	EmployeeOpenIdList *[]string `json:"employee_open_id_list,omitempty"`
	// EnterpriseId 是企业码企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// ExpenseCtrlRuleInfoList 是费控条件列表	特殊说明：	1）至少存在日额度（QUOTA_DAY）、月额度（QUOTA_MONTH）、有效期总额度（QUOTA_TOTAL）三者中的一个额度条件；	2）如果费用类型为MEAL，费控维度商户（MEAL_MERCHANT）和商户类型（MCC）对应的费控条件必须存在其一且不能同时存在；	3）如果费用类型为METRO，费控维度地铁卡类型（CARD_TYPE）对应的费控条件必须存在；	4）如果因公场景为OVERTIME，费控维度时间段（ALARM_CLOCK_TIME）对应的费控条件必须存在；	5）不能存在重复的费控维度对应的费控条件；	6）非MEAL费用类型，商户仅支持MERCHANT，不支持MEAL_MERCHANT
	ExpenseCtrlRuleInfoList *[]expenserules.ExpenseCtrRuleInfo `json:"expense_ctrl_rule_info_list,omitempty"`
	// ExpenseType 是费用类型
	ExpenseType *string `json:"expense_type,omitempty"`
	// PaymentPolicy 是当笔消费金额大于规则可用余额时，用于控制支付策略，该字段缺省时采取因公账户和个人账户组合支付策略， 枚举值：PERSONAL（全部个人账户支付）, COMBINATION（因公账户和个人账户组合支付）
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// SceneType 是因公场景
	SceneType *string `json:"scene_type,omitempty"`
	// StandardDesc 是费控规则说明	特殊说明：敏感词校验
	StandardDesc *string `json:"standard_desc,omitempty"`
	// StandardName 是费控规则名称	特殊说明：	1）敏感词校验	2）不能重复，若需要重复联系支持人员定向放开
	StandardName *string `json:"standard_name,omitempty"`
}

// InvoiceExpenseRulesSceneRuleModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type InvoiceExpenseRulesSceneRuleModifyRequest struct {
	// AccountId 是企业ID
	AccountId *string `json:"account_id,omitempty"`
	// Action 是修改操作	枚举值：	ADD_RULE（新增费控条件），DELETE_RULE（删除费控条件），MODIFY_RULE（修改费控条件），MODIFY_STANDARD（修改费控规则基本信息）
	Action *string `json:"action,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty"`
	// EffectiveEndDate 是有效期截止	特殊说明：修改费控规则基本信息时必传
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是有效期起始	特殊说明：修改费控规则基本信息时必传
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// EnterpriseId 是企业码企业id
	EnterpriseId *string `json:"enterprise_id,omitempty"`
	// ExpenseCtrlRuleInfoList 是费控规则条件列表 特殊说明：新增费控条件/修改费控条件/删除费控条件操作时必填
	ExpenseCtrlRuleInfoList *[]expenserules.ExpenseCtrRuleInfo `json:"expense_ctrl_rule_info_list,omitempty"`
	// PaymentPolicy 是当笔消费金额大于规则可用余额时，用于控制支付策略，该字段缺省时采取因公账户和个人账户组合支付策略， 枚举值：PERSONAL（全部个人账户支付）, COMBINATION（因公账户和个人账户组合支付）
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// StandardDesc 是费控规则说明	特殊说明：	1）敏感词校验	2）修改费控规则基本信息时必传
	StandardDesc *string `json:"standard_desc,omitempty"`
	// StandardId 是费控规则ID
	StandardId *string `json:"standard_id,omitempty"`
	// StandardName 是费控规则名称	特殊说明：	1）敏感词校验	2）修改费控规则基本信息时必传	3）不能重复，若需要重复联系支持人员定向放开
	StandardName *string `json:"standard_name,omitempty"`
}

// InvoiceExpenseRulesSceneRuleQueryRequest is the query request model for alipay.ebpp.invoice.expenserules.scenerule.query.
type InvoiceExpenseRulesSceneRuleQueryRequest struct {
	// AccountId 是企业ID
	AccountId *string `json:"account_id,omitempty" query:"account_id,omitempty"`
	// AgreementNo 是授权签约协议号
	AgreementNo *string `json:"agreement_no,omitempty" query:"agreement_no,omitempty"`
	// EnterpriseId 是企业码企业id
	EnterpriseId *string `json:"enterprise_id,omitempty" query:"enterprise_id,omitempty"`
	// StandardId 是费控规则ID
	StandardId *string `json:"standard_id,omitempty" query:"standard_id,omitempty"`
	// PageNum 是页数
	PageNum *int64 `json:"page_num,omitempty" query:"page_num,required"`
	// PageSize 是页大小
	PageSize *int64 `json:"page_size,omitempty" query:"page_size,required"`
}
