package gray

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

// MiniVersionGrayCancel calls alipay.open.mini.version.gray.cancel (POST /v3/alipay/open/mini/version/gray/cancel).
// Official V3 contract: https://opendocs.alipay.com/open-v3/d539c567_alipay.open.mini.version.gray.cancel
func (s *Service) MiniVersionGrayCancel(ctx context.Context, req MiniVersionGrayCancelRequest, opts ...core.CallOption) (*core.Response[MiniVersionGrayCancelResponse], error) {
	return core.ExecuteJSON[MiniVersionGrayCancelRequest, MiniVersionGrayCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.gray.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/version/gray/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionGrayCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniVersionGrayOnline calls alipay.open.mini.version.gray.online (POST /v3/alipay/open/mini/version/gray/online).
// Official V3 contract: https://opendocs.alipay.com/open-v3/6ab395d9_alipay.open.mini.version.gray.online
func (s *Service) MiniVersionGrayOnline(ctx context.Context, req MiniVersionGrayOnlineRequest, opts ...core.CallOption) (*core.Response[MiniVersionGrayOnlineResponse], error) {
	return core.ExecuteJSON[MiniVersionGrayOnlineRequest, MiniVersionGrayOnlineResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.gray.online", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/version/gray/online",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionGrayOnlineRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
