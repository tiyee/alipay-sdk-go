// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package defaultapi

// ExtensionArea 是支付宝官方 V3 OpenAPI 模型。
type ExtensionArea struct {
	// GotoUrl 是跳转链接，当type为"image"时必传，必须是https或alipays开头的url链接。已废弃。
	GotoUrl *string `json:"goto_url,omitempty"`
	// Height 是扩展区高度，当type值为"h5"时必填，取值范围为200-500的整数,单位是px。
	Height *int64 `json:"height,omitempty"`
	// Name 是扩展区名字。
	Name *string `json:"name,omitempty"`
	// Type 是扩展区类型，当前支持的值为h5，表示扩展区内容类型为h5。。
	Type *string `json:"type,omitempty"`
	// Url 是扩展区url，传入h5页面url，必须是https开头的链接。
	Url *string `json:"url,omitempty"`
}
