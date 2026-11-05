// Package system 封装支付宝 alipay.system 领域接口（V3 REST）。
package system

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// 接口名与路由，依据官方 V3 OpenAPI 快照固定。
const (
	methodOauthToken = "alipay.system.oauth.token"
	pathOauthToken   = "/v3/alipay/system/oauth/token"
)

// Service 是 alipay.system 领域的 API 封装，
// 通过注入 core.Executor（应为 *core.Client）调用统一执行流程。
type Service struct {
	executor core.Executor

	oauthToken core.Operation[OAuthTokenRequest, OAuthTokenResponse]
}

// NewService 基于执行器构建 system Service。执行器为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	oauthToken, err := core.NewJSONOperation[OAuthTokenRequest, OAuthTokenResponse](core.OperationSpec{
		Method: methodOauthToken, HTTPMethod: core.HTTPMethodPost, Path: pathOauthToken,
		ParameterMode:   core.ParameterModeJSONBody,
		AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{executor: executor, oauthToken: oauthToken}, nil
}

const (
	// GrantTypeAuthorizationCode 表示使用授权码换取访问令牌。
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	// GrantTypeRefreshToken 表示使用刷新令牌换取新的访问令牌。
	GrantTypeRefreshToken GrantType = "refresh_token"
)

// OAuthToken 调用 alipay.system.oauth.token 换取授权访问令牌
// （V3 POST /v3/alipay/system/oauth/token）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/a9cc8a2c_alipay.system.oauth.token
// 访问令牌与刷新令牌按调用方令牌管理策略保存与轮换，SDK 不缓存。
func (s *Service) OAuthToken(ctx context.Context, req OAuthTokenRequest,
	opts ...core.CallOption) (*core.Response[OAuthTokenResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.oauthToken,
		core.Request[OAuthTokenRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
