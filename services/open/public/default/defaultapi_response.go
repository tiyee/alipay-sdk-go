// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package defaultapi

// PublicDefaultExtensionCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicDefaultExtensionCreateResponse struct {
	// ExtensionKey 是一套扩展区的key，创建一套扩展区成功后，支付宝会将该字段返回，后续对扩展区进行删除等操作都会用到这个值。。
	ExtensionKey *string `json:"extension_key,omitempty"`
}
