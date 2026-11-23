package twostage

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

// TwoStageCommonUse calls alipay.user.twostage.common.use (POST /v3/alipay/user/twostage/common/use).
// Official V3 contract: https://opendocs.alipay.com/apis/05xim7
func (s *Service) TwoStageCommonUse(ctx context.Context, req UserTwoStageCommonUseRequest, opts ...core.CallOption) (*core.Response[UserTwoStageCommonUseResponse], error) {
	return core.ExecuteJSON[UserTwoStageCommonUseRequest, UserTwoStageCommonUseResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.twostage.common.use", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/twostage/common/use",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserTwoStageCommonUseRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// TwoStageIndirectUse calls alipay.user.twostage.indirect.use (POST /v3/alipay/user/twostage/indirect/use).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05x9jk
func (s *Service) TwoStageIndirectUse(ctx context.Context, req UserTwoStageIndirectUseRequest, opts ...core.CallOption) (*core.Response[UserTwoStageIndirectUseResponse], error) {
	return core.ExecuteJSON[UserTwoStageIndirectUseRequest, UserTwoStageIndirectUseResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.twostage.indirect.use", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/twostage/indirect/use",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserTwoStageIndirectUseRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
