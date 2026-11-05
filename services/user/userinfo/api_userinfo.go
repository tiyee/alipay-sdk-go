package userinfo

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

// UserInfoShare calls alipay.user.userinfo.share (POST /v3/alipay/user/userinfo/share).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a7417208_alipay.user.userinfo.share
func (s *Service) UserInfoShare(ctx context.Context, req UserUserInfoShareRequest, opts ...core.CallOption) (*core.Response[UserUserInfoShareResponse], error) {
	return core.ExecuteJSON[UserUserInfoShareRequest, UserUserInfoShareResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.userinfo.share", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/userinfo/share",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserUserInfoShareRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
