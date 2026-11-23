package app

// AppMessageTemplateSubscribeQueryRequest 是支付宝官方 V3 OpenAPI 模型。
type AppMessageTemplateSubscribeQueryRequest struct {
	// OpenId 是支付宝openId，用于支付宝用户在当前应用下的用户标识。。
	OpenId *string `json:"open_id,omitempty"`
	// TemplateIdList 是消息模板id，可以填写多个，最多不超过3个。模板id需要保持同一个应用主体，并且展示在同一个订阅组件中的模板id。 模板id获取详情参见<a href="https://opendocs.alipay.com/mini/01rnqx">模板消息</a>。。
	TemplateIdList *[]string `json:"template_id_list,omitempty"`
	// UserId 是订阅消息模板用户的支付宝唯一标识，2088开头。。
	UserId *string `json:"user_id,omitempty"`
}

// AppMiniTemplateMessageSendRequest 是支付宝官方 V3 OpenAPI 模型。
type AppMiniTemplateMessageSendRequest struct {
	// Data 是模板消息内容。商家/开发者将模板占位符替换为自定义内容。
	Data *string `json:"data,omitempty"`
	// FormId 是支付消息模板：需传入用户发生的交易行为的支付宝交易号 trade_no；</br>表单提交模板：需传入用户在小程序触发表单提交事件获得的表单号；</br>刷脸消息模板：需传入在IOT刷脸后得到的ftoken等，用于信息发送的校验。</br>。
	FormId *string `json:"form_id,omitempty"`
	// Page 是小程序的跳转页面。用于用户点击模板消息 进入小程序查看 按钮后，跳转至商家小程序对应页面。。
	Page *string `json:"page,omitempty"`
	// ToOpenId 是支付宝openId，用于支付宝用户在当前应用下的用户标识。。
	ToOpenId *string `json:"to_open_id,omitempty"`
	// ToUserId 是接收模板消息用户支付宝 user_id，可通过 <a href="https://opendocs.alipay.com/mini/api/openapi-authorize">用户授权</a> 获取。
	ToUserId *string `json:"to_user_id,omitempty"`
	// UserTemplateId 是商家在商家自运营中心选用的消息模板ID，详情参见 <a href="https://opendocs.alipay.com/mini/03l9bb?pathHash=19d2e0aa&ref=api#%E9%A2%86%E7%94%A8%E8%AE%A2%E9%98%85%E6%B6%88%E6%81%AF%E6%A8%A1%E6%9D%BF">选用消息模板</a> 。。
	UserTemplateId *string `json:"user_template_id,omitempty"`
}

// AppQRCodeCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type AppQRCodeCreateRequest struct {
	// Color 是圆形二维码颜色（十六进制颜色色值），仅圆形二维码支持颜色设置，方形二维码默认为黑色。。
	Color *string `json:"color,omitempty"`
	// Describe 是对应的二维码描述。
	Describe *string `json:"describe,omitempty"`
	// QueryParam 是小程序启动参数，内容按照格式为参数名=参数值&参数名=参数值。打开小程序的query ，在小程序 onLaunch的方法中获取。可查看 <a href="https://opendocs.alipay.com/support/01rb2a">如何获取各种场景的启动参数</a>。。
	QueryParam *string `json:"query_param,omitempty"`
	// Size 是合成后图片的大小规格，有s、m、l三档可选。。
	Size *string `json:"size,omitempty"`
	// UrlParam 是page/component/component-pages/view/view为小程序中能访问到的页面路径。
	UrlParam *string `json:"url_param,omitempty"`
}
