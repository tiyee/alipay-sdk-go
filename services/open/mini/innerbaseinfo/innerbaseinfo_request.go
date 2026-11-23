// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package innerbaseinfo

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// MiniInnerBaseInfoAppLogoUploadRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniInnerBaseInfoAppLogoUploadRequest struct {
	AppLogo *core.MultipartFile `json:"app_logo,omitempty" multipart:"app_logo,file"`
}

// MiniInnerBaseInfoQueryRequest is the request type for alipay.open.mini.innerbaseinfo.query.
type MiniInnerBaseInfoQueryRequest struct {
	// MiniAppId 是小程序ID，mini_app_id 和 mini_app_name 两个需要有其中一个必填，当填了mini_app_id时只使用id去进行查询。
	MiniAppId *string `json:"mini_app_id,omitempty" query:"mini_app_id,omitempty"`
	// InstCode 是租户code，alipay or taobao
	InstCode *string `json:"inst_code,omitempty" query:"inst_code,omitempty"`
	// MiniAppName 是小程序name，mini_app_id 和 mini_app_name 两个需要有其中一个必填，当填了mini_app_id时只使用id去进行查询。
	MiniAppName *string `json:"mini_app_name,omitempty" query:"mini_app_name,omitempty"`
	// AppSubType 是小程序类型，TINYAPP_TEMPLATE，TINYAPP_NORMAL，TINYAPP_PLUGIN，使用mini_app_name查询的时候，该字段要求必传。
	AppSubType *string `json:"app_sub_type,omitempty" query:"app_sub_type,omitempty"`
}
