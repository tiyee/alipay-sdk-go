// Package open 提供支付宝 alipay.open 领域共享模型与应用授权令牌接口（V3 REST）。
// 其他 open 子域接口位于本包下的业务子包。
package open

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// 接口名与路由，依据官方 V3 OpenAPI 快照固定。
const (
	methodAuthTokenApp = "alipay.open.auth.token.app"
	pathAuthTokenApp   = "/v3/alipay/open/auth/token/app"
)

// Service 是 alipay.open 领域（应用授权）的 API 封装，
// 通过注入 core.Executor（应为 *core.Client）调用统一执行流程。
type Service struct {
	executor core.Executor

	authTokenApp core.Operation[AuthTokenAppRequest, AuthTokenAppResponse]
}

// NewService 基于执行器构建 open Service。执行器为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	authTokenApp, err := core.NewJSONOperation[AuthTokenAppRequest, AuthTokenAppResponse](core.OperationSpec{
		Method: methodAuthTokenApp, HTTPMethod: core.HTTPMethodPost, Path: pathAuthTokenApp,
		ParameterMode:   core.ParameterModeJSONBody,
		AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{executor: executor, authTokenApp: authTokenApp}, nil
}

const (
	// GrantTypeAuthorizationCode 表示使用应用授权码换取应用授权令牌。
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	// GrantTypeRefreshToken 表示使用刷新令牌换取新的应用授权令牌。
	GrantTypeRefreshToken GrantType = "refresh_token"
)

// AuthTokenApp 调用 alipay.open.auth.token.app 换取应用授权令牌
// （V3 POST /v3/alipay/open/auth/token/app）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/4b00cc9c_alipay.open.auth.token.app
// 令牌按调用方授权管理策略保存；代调用经 core.WithAppAuthToken
// 传递，不修改共享 Client 配置。
func (s *Service) AuthTokenApp(ctx context.Context, req AuthTokenAppRequest,
	opts ...core.CallOption) (*core.Response[AuthTokenAppResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.authTokenApp,
		core.Request[AuthTokenAppRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
