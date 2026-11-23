package trade

import (
	"net/url"
	"strings"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/protocol"
	"github.com/tiyee/alipay-sdk-go/internal/tradevalidate"
)

func (r PreCreateRequest) validate() error {
	if err := validatePaymentOrderNo(r.OutTradeNo); err != nil {
		return err
	}
	for _, ch := range r.OutTradeNo {
		if !(ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' || ch == '_') {
			return &core.ValidationError{Field: "out_trade_no",
				Reason: "out_trade_no only supports letters, digits and underscore"}
		}
	}
	if err := tradevalidate.OrderAmount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if err := validateNotifyEcho(r.Subject, r.Body, r.PassbackParams); err != nil {
		return err
	}
	if r.PaymentType != nil && *r.PaymentType != PaymentTypeAgentPay {
		return &core.ValidationError{Field: "payment_type", Reason: "payment_type must be agent_pay"}
	}
	if strings.ContainsAny(r.Subject, "/=&") {
		return &core.ValidationError{Field: "subject",
			Reason: `subject must not contain "/", "=" or "&"`}
	}
	if r.DiscountableAmount != nil {
		if err := tradevalidate.OrderAmount(*r.DiscountableAmount, "discountable_amount"); err != nil {
			return err
		}
	}
	if r.UndiscountableAmount != nil {
		if err := tradevalidate.OrderAmount(*r.UndiscountableAmount, "undiscountable_amount"); err != nil {
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

func validatePaymentOrderNo(value string) error {
	if value == "" {
		return &core.ValidationError{Field: "out_trade_no", Reason: "out_trade_no is required"}
	}
	if protocol.HasNotifyBoundary(value) {
		return &core.ValidationError{Field: "out_trade_no", Reason: "ambiguous field boundary in notification order number"}
	}
	return nil
}

// validateNotifyEcho 拦截可能在普通支付通知中被重新解释为顶层字段的回传文本。
func validateNotifyEcho(subject string, body, passback *string) error {
	if subject == "" {
		return &core.ValidationError{Field: "subject", Reason: "subject is required"}
	}
	if protocol.HasNotifyBoundary(subject) {
		return &core.ValidationError{Field: "subject", Reason: "ambiguous field boundary in notification text"}
	}
	if body != nil && protocol.HasNotifyBoundary(*body) {
		return &core.ValidationError{Field: "body", Reason: "ambiguous field boundary in notification text"}
	}
	if passback != nil {
		// passback_params 在支付产物中由调用者预先 URL 编码。仅检查，保留传输原文。
		decoded, err := url.QueryUnescape(*passback)
		if err != nil || protocol.HasNotifyBoundary(decoded) || protocol.HasNotifyBoundary(*passback) {
			return &core.ValidationError{Field: "passback_params", Reason: "invalid encoding or ambiguous field boundary in notification text"}
		}
	}
	return nil
}
