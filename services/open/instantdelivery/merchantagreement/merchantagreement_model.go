package merchantagreement

// AgreementDetail 是支付宝官方 V3 OpenAPI 模型。
type AgreementDetail struct {
	// AgreementNo 是协议号。
	AgreementNo *string `json:"agreement_no,omitempty"`
	// Status 是协议状态。
	Status *string `json:"status,omitempty"`
}
