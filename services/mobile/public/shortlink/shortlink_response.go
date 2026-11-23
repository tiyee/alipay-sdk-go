package shortlink

// MobilePublicShortLinkCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicShortLinkCreateResponse struct {
	// Code 是结果码
	Code *string `json:"code,omitempty"`
	// Msg 是成功
	Msg *string `json:"msg,omitempty"`
	// Shortlink 是短链接url
	Shortlink *string `json:"shortlink,omitempty"`
}
