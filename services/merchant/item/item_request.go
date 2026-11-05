package item

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// MerchantItemFileUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantItemFileUploadRequest struct {
	Data        *MerchantItemFileUpload `json:"data,omitempty" multipart:"data,json"`
	FileContent *core.MultipartFile     `json:"file_content,omitempty" multipart:"file_content,file"`
}
