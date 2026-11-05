// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package signflows

// AttachmentDetail 是支付宝官方 V3 OpenAPI 模型。
type AttachmentDetail struct {
	// Code 是创建流程时指定的附件code，可能为空。（身份证正反面的固定code：正面：ESIGN_IDCARD，背面：ESIGN_BACK_IDCARD）
	Code *string `json:"code,omitempty"`
	// FileName 是附件文件名
	FileName *string `json:"file_name,omitempty"`
	// FileUrl 是附件下载地址
	FileUrl *string `json:"file_url,omitempty"`
}

// DocInfo 是支付宝官方 V3 OpenAPI 模型。
type DocInfo struct {
	// FileName 是合同文件名
	FileName *string `json:"file_name,omitempty"`
	// FileUrl 是合同下载地址
	FileUrl *string `json:"file_url,omitempty"`
}
