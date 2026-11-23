package image

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// MerchantImageUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type MerchantImageUploadRequest struct {
	Data         *MerchantImageUpload `json:"data,omitempty" multipart:"data,json"`
	ImageContent *core.MultipartFile  `json:"image_content,omitempty" multipart:"image_content,file"`
}
