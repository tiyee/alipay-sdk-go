// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package file

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// FileUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type FileUploadRequest struct {
	Data        *FileUpload         `json:"data,omitempty" multipart:"data,json"`
	FileContent *core.MultipartFile `json:"file_content,omitempty" multipart:"file_content,file"`
}
