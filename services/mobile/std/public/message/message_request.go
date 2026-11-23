package message

// MobileStdPublicMessageCustomSendRequest 是支付宝官方 V3 OpenAPI 模型。
type MobileStdPublicMessageCustomSendRequest struct {
	// BizContent 是业务内容，其中包括消息类型msgType、消息体和消息接受人toUserId三大块，具体参见“表1-2 服务窗单发客服消息的biz_content参数说明”。
	BizContent *string `json:"biz_content,omitempty"`
}
