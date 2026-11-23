package open

// GrantType 是应用授权方式，取值依据官方 V3 文档。
type GrantType string

// AppTokenExchangeElement 是批量授权换码的单个令牌元素。
type AppTokenExchangeElement struct {
	// AppAuthToken 是应用授权令牌。
	AppAuthToken *string `json:"app_auth_token,omitempty"`
	// AppRefreshToken 是刷新令牌。
	AppRefreshToken *string `json:"app_refresh_token,omitempty"`
	// AuthAppID 是授权商户的 appid。
	AuthAppID *string `json:"auth_app_id,omitempty"`
	// UserID 是授权商户的 user_id。
	UserID *string `json:"user_id,omitempty"`
	// ExpiresIn 是过期时间，单位为秒。
	ExpiresIn *string `json:"expires_in,omitempty"`
	// ReExpiresIn 是刷新令牌过期时间，单位为秒。
	ReExpiresIn *string `json:"re_expires_in,omitempty"`
}

// IsKnown 报告授权方式是否属于文档已知值。
func (g GrantType) IsKnown() bool {
	return g == GrantTypeAuthorizationCode || g == GrantTypeRefreshToken
}
