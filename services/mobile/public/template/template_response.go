package template

import (
	"encoding/json"
)

// MobilePublicTemplateMessageDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicTemplateMessageDeleteResponse = json.RawMessage

// MobilePublicTemplateMessageGetResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicTemplateMessageGetResponse struct {
	// MsgTemplateId 是消息模板id--商户领取母版后生成的唯一模板id
	MsgTemplateId *string `json:"msg_template_id,omitempty"`
	// Template 是模板内容
	Template *string `json:"template,omitempty"`
}

// MobilePublicTemplateMessageQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicTemplateMessageQueryResponse struct {
	// Result 是结果值
	Result *string `json:"result,omitempty"`
}
