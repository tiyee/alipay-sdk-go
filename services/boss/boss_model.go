package boss

// MultiCurrencyMoneyOpenAPI 是支付宝官方 V3 OpenAPI 模型。
type MultiCurrencyMoneyOpenAPI struct {
	// Cent 是表示金额中的分，为该货币的最小单位, 如人民币中的分、美元中的美分, 韩元日元中的元等, 具体换算比例在多币种类MultiCurrencyMoney中有字段digits
	Cent *int64 `json:"cent,omitempty"`
	// CurrencyValue 是币种类型值
	CurrencyValue *string `json:"currency_value,omitempty"`
}
