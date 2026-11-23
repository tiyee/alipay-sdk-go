// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package auth

// AuthAppAESSetRequest 是支付宝官方 V3 OpenAPI 模型。
type AuthAppAESSetRequest struct {
	// MerchantAppId 是商家应用appId。
	MerchantAppId *string `json:"merchant_app_id,omitempty"`
}

// AuthAppAESGetRequest is the request type for alipay.open.auth.app.aes.get.
type AuthAppAESGetRequest struct {
	// MerchantAppId 是商家应用appId
	MerchantAppId *string `json:"merchant_app_id,omitempty" query:"merchant_app_id,omitempty"`
}

// AuthTokenAppQueryRequest is the request type for alipay.open.auth.token.app.query.
type AuthTokenAppQueryRequest struct {
	// AppAuthToken 是应用授权令牌
	AppAuthToken *string `json:"app_auth_token,omitempty" query:"app_auth_token,omitempty"`
}

// AuthUserAuthRelationshipQueryRequest is the request type for alipay.open.auth.userauth.relationship.query.
type AuthUserAuthRelationshipQueryRequest struct {
	// UserId 是蚂蚁统一会员ID，如果未开启openid改造，则此值必传
	UserId *string `json:"user_id,omitempty" query:"user_id,omitempty"`
	// OpenId 是蚂蚁统一会员ID，蚂蚁统一会员ID，如果已完成openid改造，则此值必传
	OpenId *string `json:"open_id,omitempty" query:"open_id,omitempty"`
	// Scopes 是查询授权关系所关联的scope，半角逗号分隔的列表，且不允许有重复值
	Scopes *string `json:"scopes,omitempty" query:"scopes,omitempty"`
}
