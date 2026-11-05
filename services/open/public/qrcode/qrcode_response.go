// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package qrcode

// PublicQRCodeCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type PublicQRCodeCreateResponse struct {
	// CodeImg 是二维码图片地址，可跳转到实际图片。
	CodeImg *string `json:"code_img,omitempty"`
	// ExpireSecond 是二维码有效时间，单位（秒）。永久码暂时忽略该信息。
	ExpireSecond *string `json:"expire_second,omitempty"`
}
