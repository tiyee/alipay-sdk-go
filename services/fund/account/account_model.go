package account

// BalanceAccountDetail 是支付宝官方 V3 OpenAPI 模型。
type BalanceAccountDetail struct {
	// Acs 是acs余额，单位：元。
	Acs *string `json:"acs,omitempty"`
	// Bank 是bank余额，单位：元。
	Bank *string `json:"bank,omitempty"`
}
