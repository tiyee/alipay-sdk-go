// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package isv

// CcmISVInitializeRequest 是支付宝官方 V3 OpenAPI 模型。
type CcmISVInitializeRequest struct {
	// IsvPubKey 是用于接入CCM iframe接口安全认证的 ISV 公钥，采用 base64 编码，秘钥对生成算法为 RSA2048。
	IsvPubKey *string `json:"isv_pub_key,omitempty"`
}
