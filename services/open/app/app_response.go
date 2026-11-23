package app

// AppMessageTemplateSubscribeQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AppMessageTemplateSubscribeQueryResponse struct {
	// ShowComponent 是是否显示订阅组件。
	ShowComponent *bool `json:"show_component,omitempty"`
	// SubscribeRelations 是用户对消息模板的订阅关系列表，为入参中的用户id对消息模板id的订阅关系。 限制：用户未订阅消息，该参数不返回。。
	SubscribeRelations *[]SubscribeRelation `json:"subscribe_relations,omitempty"`
}

// AppMiniTemplateMessageSendResponse 是支付宝官方 V3 OpenAPI 模型。
type AppMiniTemplateMessageSendResponse struct {
}

// AppQRCodeCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type AppQRCodeCreateResponse struct {
	// QrCodeUrl 是方形二维码图片链接地址。
	QrCodeUrl *string `json:"qr_code_url,omitempty"`
	// QrCodeUrlCircleBlue 是圆形二维码地址，蓝色slogan。
	QrCodeUrlCircleBlue *string `json:"qr_code_url_circle_blue,omitempty"`
	// QrCodeUrlCircleWhite 是圆形二维码地址，白色slogan。
	QrCodeUrlCircleWhite *string `json:"qr_code_url_circle_white,omitempty"`
}
