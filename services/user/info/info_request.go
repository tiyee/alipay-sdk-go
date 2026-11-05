package info

// UserInfoShareRequest is the query request model for alipay.user.info.share.
type UserInfoShareRequest struct {
	// AuthToken 是用户授权令牌
	AuthToken *string `json:"auth_token,omitempty" query:"auth_token,omitempty"`
}
