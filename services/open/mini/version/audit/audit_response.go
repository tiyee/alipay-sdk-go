package audit

// MiniVersionAuditApplyResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionAuditApplyResponse struct {
	// SpeedUp 是是否加速审核 加速：true； 不加速：false。
	SpeedUp *string `json:"speed_up,omitempty"`
	// SpeedUpMemo 是提审加速审核说明。
	SpeedUpMemo *string `json:"speed_up_memo,omitempty"`
}

// MiniVersionAuditCancelResponse 是支付宝官方 V3 OpenAPI 模型。
type MiniVersionAuditCancelResponse struct {
}
