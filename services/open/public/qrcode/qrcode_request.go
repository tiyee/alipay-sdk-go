// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package qrcode

// PublicQRCodeCreateRequest 是支付宝官方 V3 OpenAPI 模型。
type PublicQRCodeCreateRequest struct {
	CodeInfo *CodeInfo `json:"code_info,omitempty"`
	// CodeType 是二维码类型，目前只支持两种类型： TEMP：临时的（默认）； PERM：永久的。
	CodeType *string `json:"code_type,omitempty"`
	// ExpireSecond 是临时码过期时间，以秒为单位，最大不超过1800秒； 永久码置空。
	ExpireSecond *string `json:"expire_second,omitempty"`
	// ShowLogo 是二维码中间是否显示服务窗logo，Y：显示；N：不显示（默认）。
	ShowLogo *string `json:"show_logo,omitempty"`
}
