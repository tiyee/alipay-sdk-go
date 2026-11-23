// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package material

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// MaterialCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type MaterialCreateRequest struct {
	// MaterialFields 是素材字段列表，素材提报的具体内容，包含三种类型：图片、文本、链接。图片类型，则根据“图片资源上传接口”中返回的resource_id字段值进行素材字段的设值。文本或链接类型，则按照实际需要设值即可，需满足具体的规范要求。规范要求详情请参见产品文档。
	MaterialFields *[]MaterialField `json:"material_fields,omitempty"`
	// MaterialName 是素材名称，接入方自行生成，建议命名时遵循一定的规律以方便管理。
	MaterialName *string `json:"material_name,omitempty"`
	// MaterialSpecId 是素材规范ID，用于标识具体场景下的素材规范，根据具体的使用场景获取对应的素材规范ID。
	MaterialSpecId *string `json:"material_spec_id,omitempty"`
	// OutBizNo 是外部业务单号，用作幂等控制，相同单号会返回上一次的结果，接入方需保证单号唯一。
	OutBizNo *string `json:"out_biz_no,omitempty"`
}

// MaterialImageUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type MaterialImageUploadRequest struct {
	Data        *MarketingMaterialImageUpload `json:"data,omitempty" multipart:"data,json"`
	FileContent *core.MultipartFile           `json:"file_content,omitempty" multipart:"file_content,file"`
}
