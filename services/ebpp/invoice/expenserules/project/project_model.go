package project

import (
	"github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/expenserules"
)

// ExpenseCtrRuleGroupInfo 是支付宝官方 V3 OpenAPI 模型。
type ExpenseCtrRuleGroupInfo struct {
	// ExpenseCtrlRuleInfoList 是费控条件列表
	ExpenseCtrlRuleInfoList *[]expenserules.ExpenseCtrRuleInfo `json:"expense_ctrl_rule_info_list,omitempty"`
	// ExpenseType 是费用类型	特殊说明：仅支持费用类型DEFAULT
	ExpenseType *string `json:"expense_type,omitempty"`
	// GroupId 是费控规则ID
	GroupId *string `json:"group_id,omitempty"`
	// GroupName 是费控规则名称	约束：	1）敏感词校验	2）不能重复
	GroupName *string `json:"group_name,omitempty"`
}

// ProjectRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type ProjectRuleInfo struct {
	// EffectiveEndDate 是有效期截止
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是有效期起始
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// EmployeeList 是切换open_id前请使用此字段：员工uid列表
	EmployeeList *[]string `json:"employee_list,omitempty"`
	// EmployeeOpenIdList 是切换open_id后请使用此字段：员工open_id列表
	EmployeeOpenIdList *[]string `json:"employee_open_id_list,omitempty"`
	// ExpenseCtrlRuleInfoGroupList 是规则组列表
	ExpenseCtrlRuleInfoGroupList *[]ExpenseCtrRuleGroupInfo `json:"expense_ctrl_rule_info_group_list,omitempty"`
	// ProjectId 是项目id
	ProjectId *string `json:"project_id,omitempty"`
	// ProjectName 是项目名称
	ProjectName *string `json:"project_name,omitempty"`
}
