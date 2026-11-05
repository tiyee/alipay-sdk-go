// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package miniapp

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// MiniMiniAppBrandSubmitRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniMiniAppBrandSubmitRequest struct {
	// ApplyType 是申请的资质类型 SELF_BRAND_APPLY 自有品牌申请(当需要一个新的品牌时选择此项) AUTHORIZE_BRAND_APPLY 授权品牌申请(当需要授权一个已有的品牌给商户时选择此项) EXIST_BRAND_BIND 品牌绑定，适用于品牌已申请成功的情况(当需要将已有品牌绑定到无品牌小程序时选择此项)。
	ApplyType     *string        `json:"apply_type,omitempty"`
	AuthorizeInfo *AuthorizeInfo `json:"authorize_info,omitempty"`
	// BrandId 是商户已有品牌Id(只有申请的资质类型为EXIST_BRAND_BIND 时需要填写)。
	BrandId *string `json:"brand_id,omitempty"`
	// BrandName 是品牌名称。
	BrandName             *string                `json:"brand_name,omitempty"`
	BrandRegistrationInfo *BrandRegistrationInfo `json:"brand_registration_info,omitempty"`
	// IdMaterials 是申请人身份证明材料(当前只支持图片类型，请调用alipay.open.mini.miniapp.brand.upload接口上传图片，可以上传多个身份证明材料)。
	IdMaterials *[]string `json:"id_materials,omitempty"`
}

// MiniMiniAppServiceConfigModifyRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniMiniAppServiceConfigModifyRequest struct {
	// HomeOpen 是云客服是否在小程序首页透出，true-开启，false-不开启。
	HomeOpen *bool `json:"home_open,omitempty"`
	// ServiceConfig 是客服方式，目前支持ANTCLOUD-云客服。
	ServiceConfig *string `json:"service_config,omitempty"`
}

// MiniMiniAppBrandUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniMiniAppBrandUploadRequest struct {
	FileContent *core.MultipartFile `json:"file_content,omitempty" multipart:"file_content,file"`
}

// MiniMiniAppBrandCancelRequest is the request type for alipay.open.mini.miniapp.brand.cancel.
type MiniMiniAppBrandCancelRequest struct{}

// MiniMiniAppBrandQueryRequest is the request type for alipay.open.mini.miniapp.brand.query.
type MiniMiniAppBrandQueryRequest struct{}
