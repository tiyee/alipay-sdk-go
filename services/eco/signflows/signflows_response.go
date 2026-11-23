// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package signflows

// SignFlowsDetailQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SignFlowsDetailQueryResponse struct {
	Attachments *AttachmentDetail `json:"attachments,omitempty"`
	Docs        *DocInfo          `json:"docs,omitempty"`
}

// SignFlowsURLQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type SignFlowsURLQueryResponse struct {
	// PreviewShortUrl 是预览短链地址（30天有效）
	PreviewShortUrl *string `json:"preview_short_url,omitempty"`
	// PreviewUrl 是预览长链地址(永久有效)
	PreviewUrl *string `json:"preview_url,omitempty"`
	// SignShortUrl 是签署短链地址（30天有效）
	SignShortUrl *string `json:"sign_short_url,omitempty"`
	// SignUrl 是签署长链地址(永久有效)
	SignUrl *string `json:"sign_url,omitempty"`
}
