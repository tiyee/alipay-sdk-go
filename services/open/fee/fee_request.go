// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package fee

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// FeeAdjustApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type FeeAdjustApplyRequest struct {
	Attachment   *core.MultipartFile `json:"attachment,omitempty" multipart:"attachment,file"`
	CertPic      *core.MultipartFile `json:"cert_pic,omitempty" multipart:"cert_pic,file"`
	Data         *FeeAdjustApply     `json:"data,omitempty" multipart:"data,json"`
	ShopScenePic *core.MultipartFile `json:"shop_scene_pic,omitempty" multipart:"shop_scene_pic,file"`
	ShopSignPic  *core.MultipartFile `json:"shop_sign_pic,omitempty" multipart:"shop_sign_pic,file"`
}
