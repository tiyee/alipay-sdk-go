package trans

import (
	"strings"

	"github.com/tiyee/alipay-sdk-go/core"
)

// validAmount 校验金额为精确的十进制字符串：非负、至多两位小数、
// 至少含一个非零数字。金额不得使用浮点数建模或计算。
func validAmount(amount, field string) error {
	if amount == "" {
		return &core.ValidationError{Field: field, Reason: "amount is required"}
	}
	intPart, fracPart, found := strings.Cut(amount, ".")
	if intPart == "" && fracPart == "" {
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
