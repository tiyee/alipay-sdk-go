// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package message

// PublicMessageContentCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageContentCreateResponse struct {
	// ContentId 是内容id。
	ContentId *string `json:"content_id,omitempty"`
	// ContentUrl 是内容详情内链url。
	ContentUrl *string `json:"content_url,omitempty"`
}

// PublicMessageContentModifyResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageContentModifyResponse struct {
	// ContentId 是内容id。
	ContentId *string `json:"content_id,omitempty"`
	// ContentUrl 是内容详情内链url。
	ContentUrl *string `json:"content_url,omitempty"`
}

// PublicMessageCustomSendResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageCustomSendResponse struct {
}

// PublicMessageGroupSendResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageGroupSendResponse struct {
	// MessageId 是2013121100055554f1000b64-3f05-4581-a5e1-16e29d242950。
	MessageId *string `json:"message_id,omitempty"`
}

// PublicMessageLabelSendResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageLabelSendResponse struct {
}

// PublicMessageSingleSendResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageSingleSendResponse struct {
}

// PublicMessageTotalSendResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicMessageTotalSendResponse struct {
	// MessageId 是消息ID。
	MessageId *string `json:"message_id,omitempty"`
}
