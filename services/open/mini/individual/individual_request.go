// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package individual

// MiniIndividualBusinessCertifyRequest 是支付宝官方 V3 OpenAPI 模型。
type MiniIndividualBusinessCertifyRequest struct {
	// LicenseNo 是营业执照注册号即营业执照的编号。。
	LicenseNo *string `json:"license_no,omitempty"`
	// LicensePic 是营业执照图片的Base64编码字符串，图片大小不能超过2M。
	LicensePic *string `json:"license_pic,omitempty"`
}
