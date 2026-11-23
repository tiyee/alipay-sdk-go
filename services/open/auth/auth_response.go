// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package auth

// AuthAppAESGetResponse 是支付宝官方 V3 OpenAPI 模型。
type AuthAppAESGetResponse struct {
	// AesKey 是商家应用AES KEY密文，需要ISV使用三方应用配置的AES密钥内容进行解密。详情参见 <a href="https://opendocs.alipay.com/isv/grefvl/getaes">应用AES密钥管理</a>。。
	AesKey *string `json:"aes_key,omitempty"`
}

// AuthAppAESSetResponse 是支付宝官方 V3 OpenAPI 模型。
type AuthAppAESSetResponse struct {
	// AesKey 是商家应用AES KEY密文，需要ISV使用三方应用配置的AES密钥内容进行解密。详情参见 <a href="https://opendocs.alipay.com/isv/grefvl/getaes">应用AES密钥管理</a>。。
	AesKey *string `json:"aes_key,omitempty"`
}

// AuthTokenAppQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AuthTokenAppQueryResponse struct {
	// AuthAppId 是授权商户的appid。
	AuthAppId *string `json:"auth_app_id,omitempty"`
	// AuthEnd 是授权失效时间。
	AuthEnd *string `json:"auth_end,omitempty"`
	// AuthMethods 是当前app_auth_token的授权接口列表。
	AuthMethods *[]string `json:"auth_methods,omitempty"`
	// AuthStart 是授权生效时间。
	AuthStart *string `json:"auth_start,omitempty"`
	// ExpiresIn 是应用授权令牌失效时间，单位到秒。
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	// IsByAppAuth 是当商家选择全权委托授权时值为true，未选择时为false。。
	IsByAppAuth *bool `json:"is_by_app_auth,omitempty"`
	// Status 是valid：有效状态；invalid：无效状态。
	Status *string `json:"status,omitempty"`
	// UserId 是授权商户的user_id。
	UserId *string `json:"user_id,omitempty"`
}

// AuthUserAuthRelationshipQueryResponse 是支付宝官方 V3 OpenAPI 模型。
type AuthUserAuthRelationshipQueryResponse struct {
	// QueryDetail 是授权结果查询结果，可以转换为为Map<String,Boolean>格式。
	QueryDetail *string `json:"query_detail,omitempty"`
}
