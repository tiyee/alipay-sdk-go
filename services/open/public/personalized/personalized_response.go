// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package personalized

// PublicPersonalizedExtensionBatchQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicPersonalizedExtensionBatchQueryResponse struct {
	// Count 是扩展区套数。
	Count *int64 `json:"count,omitempty"`
	// Extensions 是扩展区信息。
	Extensions *[]QueryExtension `json:"extensions,omitempty"`
}

// PublicPersonalizedExtensionCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicPersonalizedExtensionCreateResponse struct {
	// ExtensionKey 是扩展区套id。创建个性化扩展区成功后，支付宝会将该字段返回，后续扩展区上下线或者扩展区删除都会用到这个值。。
	ExtensionKey *string `json:"extension_key,omitempty"`
}

// PublicPersonalizedExtensionDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicPersonalizedExtensionDeleteResponse struct {
}

// PublicPersonalizedMenuCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicPersonalizedMenuCreateResponse struct {
	// MenuKey 是该套个性化菜单key。
	MenuKey *string `json:"menu_key,omitempty"`
}

// PublicPersonalizedMenuDeleteResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicPersonalizedMenuDeleteResponse struct {
}
