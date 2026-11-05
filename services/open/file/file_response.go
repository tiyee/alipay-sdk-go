// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package file

// FileUploadResponse 是支付宝官方 V3 OpenAPI 模型。
type FileUploadResponse struct {
	// ExternId 是视频资源id，同步外部渠道后，外部渠道返回的id。目前支持同步优酷场景，返回优酷视频id，用于支付宝客户端播放视频使用。
	ExternId *string `json:"extern_id,omitempty"`
	// FileId 是文件id，用于提供后后续业务调用使用。
	FileId *string `json:"file_id,omitempty"`
}
