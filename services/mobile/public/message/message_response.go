package message

// MobilePublicMessageCustomSendResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMessageCustomSendResponse struct {
	// Code 是结果码
	Code *string `json:"code,omitempty"`
	// Msg 是结果描述
	Msg *string `json:"msg,omitempty"`
}

// MobilePublicMessageSingleSendResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMessageSingleSendResponse struct {
	// Code 是结果码
	Code *string `json:"code,omitempty"`
	// Msg 是结果描述
	Msg *string `json:"msg,omitempty"`
}

// MobilePublicMessageTotalSendResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMessageTotalSendResponse struct {
	// Code 是结果码
	Code *string `json:"code,omitempty"`
	// Data 是消息ID
	Data *string `json:"data,omitempty"`
	// Msg 是结果描述
	Msg *string `json:"msg,omitempty"`
}
