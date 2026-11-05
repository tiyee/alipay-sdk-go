package app

// SubscribeRelation 是支付宝官方 V3 OpenAPI 模型。
type SubscribeRelation struct {
	// KeepState 是订阅保持状态，即用户勾选“总是保持以上选择，不再询问”选项时勾选的订阅状态。说明：若用户“总是保持以上选择，不再询问”选项，且选择订阅消息。下次触发消息订阅组件时，支付宝将自动发起一次静默订阅（不再拉起订阅组件，无需用户手动订阅）。。
	KeepState *string `json:"keep_state,omitempty"`
	// Show 是模板是否展示在订阅组件中。
	Show *bool `json:"show,omitempty"`
	// SubscribeState 是订阅状态。
	SubscribeState *string `json:"subscribe_state,omitempty"`
	// SubscribeType 是消息模板的订阅类型。
	SubscribeType *string `json:"subscribe_type,omitempty"`
	// TemplateId 是消息模板id。
	TemplateId *string `json:"template_id,omitempty"`
}
