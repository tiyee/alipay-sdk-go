// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package isv

// CcmISVInitializeResponse 是支付宝官方 V3 OpenAPI 模型。
type CcmISVInitializeResponse struct {
	// CcmPubKey 是Base64编码CCM公钥：CCM公私钥对由CCM自动生成，用于Iframe spi接口安全认证
	CcmPubKey *string `json:"ccm_pub_key,omitempty"`
}
