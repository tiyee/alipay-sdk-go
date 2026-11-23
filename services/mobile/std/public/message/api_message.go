package message

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

// StdPublicMessageCustomSend calls alipay.mobile.std.public.message.custom.send (POST /v3/alipay/mobile/std/public/message/custom/send).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wult
func (s *Service) StdPublicMessageCustomSend(ctx context.Context, req MobileStdPublicMessageCustomSendRequest, opts ...core.CallOption) (*core.Response[MobileStdPublicMessageCustomSendResponse], error) {
	return core.ExecuteJSON[MobileStdPublicMessageCustomSendRequest, MobileStdPublicMessageCustomSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.std.public.message.custom.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/mobile/std/public/message/custom/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobileStdPublicMessageCustomSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
