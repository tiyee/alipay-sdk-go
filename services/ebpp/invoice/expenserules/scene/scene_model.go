package scene

import (
	"github.com/tiyee/alipay-sdk-go/services/ebpp/invoice/expenserules"
)

// StandardRuleInfo 是支付宝官方 V3 OpenAPI 模型。
type StandardRuleInfo struct {
	// EffectiveEndDate 是有效期截止
	EffectiveEndDate *string `json:"effective_end_date,omitempty"`
	// EffectiveStartDate 是有效期起始
	EffectiveStartDate *string `json:"effective_start_date,omitempty"`
	// ExpenseCtrlRuleInfoList 是费控条件列表
	ExpenseCtrlRuleInfoList *[]expenserules.ExpenseCtrRuleInfo `json:"expense_ctrl_rule_info_list,omitempty"`
	// PaymentPolicy 是当笔消费金额大于规则可用余额时，用于控制支付策略，该字段缺省时采取因公账户和个人账户组合支付策略， 枚举值：PERSONAL（全部个人账户支付）
	PaymentPolicy *string `json:"payment_policy,omitempty"`
	// StandardDesc 是费控规则说明
	StandardDesc *string `json:"standard_desc,omitempty"`
	// StandardId 是费控规则ID
	StandardId *string `json:"standard_id,omitempty"`
	// StandardName 是费控规则名称
	StandardName *string `json:"standard_name,omitempty"`
}
