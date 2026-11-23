package tradevalidate

import (
	"strconv"
	"strings"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Identifier 校验 trade_no / out_trade_no 二选一约束：
// 不能同时为空；文档注明同时存在时网关优先使用 trade_no。
func Identifier(tradeNo, outTradeNo, field string) error {
	if tradeNo == "" && outTradeNo == "" {
		return &core.ValidationError{Field: field,
			Reason: "trade_no and out_trade_no cannot both be empty"}
	}
	return nil
}

// Amount 校验金额为精确的十进制字符串：非负、至多两位小数、
// 至少含一个非零数字。金额不得使用浮点数建模或计算。
func Amount(amount, field string) error {
	if amount == "" {
		return &core.ValidationError{Field: field, Reason: "amount is required"}
	}
	intPart, fracPart, found := strings.Cut(amount, ".")
	if intPart == "" || (found && fracPart == "") {
		return &core.ValidationError{Field: field, Reason: "amount is not a decimal string"}
	}
	if len(intPart) > 1 && intPart[0] == '0' {
		return &core.ValidationError{Field: field, Reason: "amount has leading zeros"}
	}
	for _, part := range [2]string{intPart, fracPart} {
		for _, ch := range part {
			if ch < '0' || ch > '9' {
				return &core.ValidationError{Field: field, Reason: "amount is not a decimal string"}
			}
		}
	}
	if found && len(fracPart) > 2 {
		return &core.ValidationError{Field: field, Reason: "amount supports at most two decimal places"}
	}
	nonZero := strings.ContainsFunc(intPart, func(ch rune) bool { return ch != '0' }) ||
		strings.ContainsFunc(fracPart, func(ch rune) bool { return ch != '0' })
	if !nonZero {
		return &core.ValidationError{Field: field, Reason: "amount must be greater than zero"}
	}
	return nil
}

// OrderAmount 校验官方限定在 [0.01, 100000000] 元的订单金额字段。
func OrderAmount(amount, field string) error {
	if err := Amount(amount, field); err != nil {
		return err
	}
	const maxYuan = 100000000
	whole, fraction, _ := strings.Cut(amount, ".")
	yuan, err := strconv.ParseUint(whole, 10, 64)
	if err != nil || yuan > maxYuan || (yuan == maxYuan && strings.Trim(fraction, "0") != "") {
		return &core.ValidationError{Field: field, Reason: "amount exceeds 100000000 yuan"}
	}
	return nil
}
