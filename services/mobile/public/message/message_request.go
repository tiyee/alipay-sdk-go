package message

// MobilePublicMessageCustomSendRequest 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMessageCustomSendRequest struct {
	// BizContent 是业务内容，其中包括消息类型msgType、消息体和消息接受人toUserId三大块，在toUserId这一层级新加eventType参数，该字段取值为follow:表示关注事件，click：表示菜单点击事件，enter_ppchat:代表进入事件。具体参见“表1-2 服务窗单发客服消息的biz_content参数说明”。 	详情请见
	BizContent *string `json:"biz_content,omitempty"`
}

// MobilePublicMessageSingleSendRequest 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMessageSingleSendRequest struct {
	// BizContent 是业务内容，其中包括模板template和消息接收人toUserId两大块，具体参见“表1-2 服务窗单发模板消息的biz_content参数说明”。 	详情请见
	BizContent *string `json:"biz_content,omitempty"`
}

// MobilePublicMessageTotalSendRequest 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicMessageTotalSendRequest struct {
	// BizContent 是业务内容，其中包括消息类型msgType和消息体两部分，具体参见“表1-2 服务窗群发消息的biz_content参数说明”。
	BizContent *string `json:"biz_content,omitempty"`
}
