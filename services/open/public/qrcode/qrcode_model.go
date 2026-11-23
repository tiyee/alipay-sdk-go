// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package qrcode

// CodeInfo 是支付宝官方 V3 OpenAPI 模型。
type CodeInfo struct {
	// GotoUrl 是跳转URL，扫码关注服务窗后会直接跳转到此URL。
	GotoUrl *string `json:"goto_url,omitempty"`
	Scene   *Scene  `json:"scene,omitempty"`
}

// Scene 是支付宝官方 V3 OpenAPI 模型。
type Scene struct {
	// SceneId 是场景Id，最长32位，英文字母、数字以及下划线，开发者自定义。
	SceneId *string `json:"scene_id,omitempty"`
}
