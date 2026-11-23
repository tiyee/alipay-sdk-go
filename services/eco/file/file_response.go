// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package file

// FilePathQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type FilePathQueryResponse struct {
	// FileId 是文件Id
	FileId *string `json:"file_id,omitempty"`
	// UploadUrl 是文件直传地址, 可以重复使用，但是只能传一样的文件，有效期一小时
	UploadUrl *string `json:"upload_url,omitempty"`
}
