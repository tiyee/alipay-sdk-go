package beta

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor core.Executor
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}

	return &Service{
		executor: executor,
	}, nil
}

// MiniInnerVersionBetaConfirm calls alipay.open.mini.innerversion.beta.confirm (POST /v3/alipay/open/mini/innerversion/beta/confirm).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05x9jh
func (s *Service) MiniInnerVersionBetaConfirm(ctx context.Context, req MiniInnerVersionBetaConfirmRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionBetaConfirmResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionBetaConfirmRequest, MiniInnerVersionBetaConfirmResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.beta.confirm", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/beta/confirm",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionBetaConfirmRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionBetaInfoQuery calls alipay.open.mini.innerversion.betainfo.query (GET /v3/alipay/open/mini/innerversion/betainfo/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xp3o
func (s *Service) MiniInnerVersionBetaInfoQuery(ctx context.Context, req MiniInnerVersionBetaInfoQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionBetaInfoQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionBetaInfoQueryRequest, MiniInnerVersionBetaInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.betainfo.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/innerversion/betainfo/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionBetaInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
