package qrcode

// MobilePublicQRCodeCreateResponse 是支付宝官方 V3 OpenAPI 模型。
type MobilePublicQRCodeCreateResponse struct {
	// Code 是返回结果码，如200，标识成功
	Code *int64 `json:"code,omitempty"`
	// CodeImg 是图片地址
	CodeImg *string `json:"code_img,omitempty"`
	// ExpireSecond 是码过期时间，单位：秒
	ExpireSecond *int64 `json:"expire_second,omitempty"`
	// Msg 是结果吗描述信息
	Msg *string `json:"msg,omitempty"`
}
