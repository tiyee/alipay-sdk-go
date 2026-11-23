package account

import (
	"github.com/tiyee/alipay-sdk-go/services/fund"
)

// FundAccountQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FundAccountQueryResponse struct {
	AmountDetail *BalanceAccountDetail `json:"amount_detail,omitempty"`
	// AvailableAmount 是账户可用余额，单位元，精确到小数点后两位。。
	AvailableAmount *string           `json:"available_amount,omitempty"`
	ExtCardInfo     *fund.ExtCardInfo `json:"ext_card_info,omitempty"`
	// FreezeAmount 是当前支付宝账户的实时冻结余额。
	FreezeAmount *string `json:"freeze_amount,omitempty"`
	// TotalAmount 是支付宝账户的余额总数，仅当查询日终余额明细时返回。单位：元。
	TotalAmount *string `json:"total_amount,omitempty"`
}
