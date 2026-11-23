package open

// AuthTokenAppResponse 是换取应用授权令牌的业务响应。
// 令牌属于敏感凭证：调用方须安全保存，不得写入日志或错误信息。
type AuthTokenAppResponse struct {
	// AppAuthToken 是应用授权令牌；代调用时经 core.WithAppAuthToken 传入。
	AppAuthToken *string `json:"app_auth_token,omitempty"`
	// AppRefreshToken 是刷新令牌。
	AppRefreshToken *string `json:"app_refresh_token,omitempty"`
	// AuthAppID 是授权商户的 appid。
	AuthAppID *string `json:"auth_app_id,omitempty"`
	// UserID 是授权商户的 user_id。
	UserID *string `json:"user_id,omitempty"`
	// ReExpiresIn 是刷新令牌的有效时间，单位为秒。
	ReExpiresIn *string `json:"re_expires_in,omitempty"`
	// Tokens 是批量授权换码的访问令牌列表。
	Tokens []AppTokenExchangeElement `json:"tokens,omitempty"`
}
