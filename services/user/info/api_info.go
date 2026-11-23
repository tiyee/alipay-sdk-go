package info

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

// InfoShare calls alipay.user.info.share (POST /v3/alipay/user/info/share).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2f4a8991_alipay.user.info.share
func (s *Service) InfoShare(ctx context.Context, req UserInfoShareRequest, opts ...core.CallOption) (*core.Response[UserInfoShareResponse], error) {
	return core.ExecuteJSON[UserInfoShareRequest, UserInfoShareResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.info.share", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/info/share",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserInfoShareRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
