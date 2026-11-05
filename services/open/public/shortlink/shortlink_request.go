// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package shortlink

// PublicShortLinkCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicShortLinkCreateRequest struct {
	// Remark 是对于场景 ID 的描述，由商户自定义。。
	Remark *string `json:"remark,omitempty"`
	// SceneId 是短链接对应的场景 ID。由商户自定义，仅支持数字、字母及下划线。。
	SceneId *string `json:"scene_id,omitempty"`
}
