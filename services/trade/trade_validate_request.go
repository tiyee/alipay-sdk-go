package trade

import (
	"strings"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/tradevalidate"
)

func (r PreCreateRequest) validate() error {
	if r.OutTradeNo == "" {
		return &core.ValidationError{Field: "out_trade_no", Reason: "out_trade_no is required"}
	}
	for _, ch := range r.OutTradeNo {
		if !(ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' || ch == '_') {
			return &core.ValidationError{Field: "out_trade_no",
				Reason: "out_trade_no only supports letters, digits and underscore"}
		}
	}
	if err := tradevalidate.Amount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if r.Subject == "" {
		return &core.ValidationError{Field: "subject", Reason: "subject is required"}
	}
	if strings.ContainsAny(r.Subject, "/=&") {
		return &core.ValidationError{Field: "subject",
			Reason: `subject must not contain "/", "=" or "&"`}
	}
	if r.DiscountableAmount != nil {
		if err := tradevalidate.Amount(*r.DiscountableAmount, "discountable_amount"); err != nil {
			return err
		}
	}
	if r.UndiscountableAmount != nil {
		if err := tradevalidate.Amount(*r.UndiscountableAmount, "undiscountable_amount"); err != nil {
			return err
		}
	}
	for i, goods := range r.GoodsDetail {
		if err := validateGoodsDetail(goods, i); err != nil {
			return err
		}
	}
	return nil
}
