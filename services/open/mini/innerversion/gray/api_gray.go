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

// MiniInnerVersionGrayFinish calls alipay.open.mini.innerversion.gray.finish (POST /v3/alipay/open/mini/innerversion/gray/finish).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xge5
func (s *Service) MiniInnerVersionGrayFinish(ctx context.Context, req MiniInnerVersionGrayFinishRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionGrayFinishResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionGrayFinishRequest, MiniInnerVersionGrayFinishResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.gray.finish", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/gray/finish",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionGrayFinishRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionGrayPublish calls alipay.open.mini.innerversion.gray.publish (POST /v3/alipay/open/mini/innerversion/gray/publish).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xk4w
func (s *Service) MiniInnerVersionGrayPublish(ctx context.Context, req MiniInnerVersionGrayPublishRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionGrayPublishResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionGrayPublishRequest, MiniInnerVersionGrayPublishResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.gray.publish", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/gray/publish",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionGrayPublishRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
