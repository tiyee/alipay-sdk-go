package api

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// AppAPIFieldApplyRequest 是支付宝官方 V3 OpenAPI 模型。
type AppAPIFieldApplyRequest struct {
	Data     *AppAPIFieldApply   `json:"data,omitempty" multipart:"data,json"`
	Picture1 *core.MultipartFile `json:"picture_1,omitempty" multipart:"picture_1,file"`
	Picture2 *core.MultipartFile `json:"picture_2,omitempty" multipart:"picture_2,file"`
	Picture3 *core.MultipartFile `json:"picture_3,omitempty" multipart:"picture_3,file"`
	Picture4 *core.MultipartFile `json:"picture_4,omitempty" multipart:"picture_4,file"`
	Picture5 *core.MultipartFile `json:"picture_5,omitempty" multipart:"picture_5,file"`
	Video    *core.MultipartFile `json:"video,omitempty" multipart:"video,file"`
}

// AppAPIFieldQueryRequest is the request type for alipay.open.app.api.field.query.
type AppAPIFieldQueryRequest struct{}

// AppAPIQueryRequest is the request type for alipay.open.app.api.query.
type AppAPIQueryRequest struct{}

// AppAPISceneQueryRequest is the request type for alipay.open.app.api.scene.query.
type AppAPISceneQueryRequest struct {
	// FieldName 是接口出参字段英文名。
	FieldName *string `json:"field_name,omitempty" query:"field_name,omitempty"`
	// ApiName 是接口英文名。
	ApiName *string `json:"api_name,omitempty" query:"api_name,omitempty"`
}
