// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package resource

// MiniResourcePromotionSourceNotifyRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniResourcePromotionSourceNotifyRequest struct {
	// AuthorId 是媒体唤起时传入的支付宝id。
	AuthorId *string `json:"author_id,omitempty"`
	// Params 是媒体唤起时提供的额外参数值列表，这里可能有多个值，打平以后拼入。即url_decode(authcbparams)。
	Params *string `json:"params,omitempty"`
	// PromotionId 是推广位id。
	PromotionId *string `json:"promotion_id,omitempty"`
	// PromotionName 是推广位名称。
	PromotionName *string `json:"promotion_name,omitempty"`
	// Source 是媒体来源，标识调用方。
	Source *string `json:"source,omitempty"`
}

// MiniResourceRecordNotifyRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniResourceRecordNotifyRequest struct {
	// AuthorId 是媒体唤起时传入的支付宝id。
	AuthorId *string `json:"author_id,omitempty"`
	// MiniAppId 是支付宝appid，代表一个小程序的唯一标识。
	MiniAppId *string `json:"mini_app_id,omitempty"`
	// Params 是媒体唤起时提供的额外参数值列表，这里可能有多个值，打平以后拼入。即url_decode(authcbparams)。
	Params *string `json:"params,omitempty"`
	// SiteId 是回调的siteid 淘宝账号唯一。
	SiteId *string `json:"site_id,omitempty"`
	// Source 是媒体来源，标识调用方。
	Source *string `json:"source,omitempty"`
	// TaobaoId 是淘宝id。
	TaobaoId *string `json:"taobao_id,omitempty"`
	// TaobaoNick 是淘宝昵称。
	TaobaoNick *string `json:"taobao_nick,omitempty"`
}
