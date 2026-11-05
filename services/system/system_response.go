package system

// OAuthTokenResponse 是换取授权访问令牌的业务响应。
// 令牌属于敏感凭证：调用方须安全保存，不得写入日志或错误信息。
type OAuthTokenResponse struct {
	// AccessToken 是访问令牌；通过该令牌调用需要授权的接口。
	AccessToken string `json:"access_token"`
	// RefreshToken 是刷新令牌；通过该令牌可以刷新 access_token。
	RefreshToken *string `json:"refresh_token,omitempty"`
	// ExpiresIn 是访问令牌的有效时间，单位是秒。
	ExpiresIn *string `json:"expires_in,omitempty"`
	// ReExpiresIn 是刷新令牌的有效时间，单位是秒。
	ReExpiresIn *string `json:"re_expires_in,omitempty"`
	// AuthStart 是授权 token 开始时间，作为有效期计算的起点。
	AuthStart *string `json:"auth_start,omitempty"`
	// AuthTokenType 是令牌类型；permanent 表示永久有效，
	// 非永久令牌不返回该字段。
	AuthTokenType *string `json:"auth_token_type,omitempty"`
	// UserID 是支付宝用户的唯一标识（2088 开头 16 位数字）。
	UserID *string `json:"user_id,omitempty"`
	// OpenID 是支付宝用户在该应用下的唯一标识。
	OpenID *string `json:"open_id,omitempty"`
	// UnionID 是用户在开放平台的唯一标识符；配置应用分组后返回。
	UnionID *string `json:"union_id,omitempty"`
}
