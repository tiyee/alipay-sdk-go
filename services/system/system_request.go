package system

import (
	"github.com/tiyee/alipay-sdk-go/core"
)

// OAuthTokenRequest 是 alipay.system.oauth.token（换取授权访问令牌）
// 的请求。grant_type=authorization_code 时 code 必填；
// grant_type=refresh_token 时 refresh_token 必填。
type OAuthTokenRequest struct {
	// GrantType 是授权方式。
	GrantType GrantType `json:"grant_type"`
	// Code 是授权码，用户对应用授权后得到；
	// grant_type 为 authorization_code 时必填。
	Code *string `json:"code,omitempty"`
	// RefreshToken 是刷新令牌，上次换取访问令牌时得到；
	// grant_type 为 refresh_token 时必填。
	RefreshToken *string `json:"refresh_token,omitempty"`
}

func (r OAuthTokenRequest) validate() error {
	if r.GrantType == "" {
		return &core.ValidationError{Field: "grant_type", Reason: "grant_type is required"}
	}
	if !r.GrantType.IsKnown() {
		return &core.ValidationError{Field: "grant_type",
			Reason: "grant_type must be authorization_code or refresh_token"}
	}
	switch r.GrantType {
	case GrantTypeAuthorizationCode:
		if r.Code == nil || *r.Code == "" {
			return &core.ValidationError{Field: "code",
				Reason: "code is required when grant_type is authorization_code"}
		}
	case GrantTypeRefreshToken:
		if r.RefreshToken == nil || *r.RefreshToken == "" {
			return &core.ValidationError{Field: "refresh_token",
				Reason: "refresh_token is required when grant_type is refresh_token"}
		}
	}
	return nil
}
