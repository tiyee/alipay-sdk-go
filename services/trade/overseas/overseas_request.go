package overseas

import (
	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/tradevalidate"
)

// OverseasSettleRequest 是 alipay.trade.overseas.settle（境内收单跨境
// 结算）的请求。写操作：out_request_no 用于幂等控制。
type OverseasSettleRequest struct {
	// TradeNo 是支付宝交易号。
	TradeNo string `json:"trade_no"`
	// OutRequestNo 是外部请求号；调用方生成并保证唯一，用于幂等控制。
	OutRequestNo string `json:"out_request_no"`
	// Amount 是跨境结算的人民币金额，单位为元。
	Amount string `json:"amount"`
	// ForeignSettleCurrency 是跨境结算的外币币种。
	ForeignSettleCurrency string `json:"foreign_settle_currency"`
	// ExtendParams 是跨境结算扩展参数。
	ExtendParams *OverseasSettleExtendParams `json:"extend_params,omitempty"`
}

func (r OverseasSettleRequest) validate() error {
	if r.TradeNo == "" {
		return &core.ValidationError{Field: "trade_no", Reason: "trade_no is required"}
	}
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	if err := tradevalidate.Amount(r.Amount, "amount"); err != nil {
		return err
	}
	if r.ForeignSettleCurrency == "" {
		return &core.ValidationError{Field: "foreign_settle_currency", Reason: "foreign_settle_currency is required"}
	}
	return nil
}
