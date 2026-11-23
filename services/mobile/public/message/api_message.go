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

// PublicMessageCustomSend calls alipay.mobile.public.message.custom.send (POST /v3/alipay/mobile/public/message/custom/send).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuk7
func (s *Service) PublicMessageCustomSend(ctx context.Context, req MobilePublicMessageCustomSendRequest, opts ...core.CallOption) (*core.Response[MobilePublicMessageCustomSendResponse], error) {
	return core.ExecuteJSON[MobilePublicMessageCustomSendRequest, MobilePublicMessageCustomSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.public.message.custom.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/mobile/public/message/custom/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobilePublicMessageCustomSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMessageSingleSend calls alipay.mobile.public.message.single.send (POST /v3/alipay/mobile/public/message/single/send).
// Official V3 contract: https://opendocs.alipay.com/apis/06bf6f
func (s *Service) PublicMessageSingleSend(ctx context.Context, req MobilePublicMessageSingleSendRequest, opts ...core.CallOption) (*core.Response[MobilePublicMessageSingleSendResponse], error) {
	return core.ExecuteJSON[MobilePublicMessageSingleSendRequest, MobilePublicMessageSingleSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.public.message.single.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/mobile/public/message/single/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobilePublicMessageSingleSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMessageTotalSend calls alipay.mobile.public.message.total.send (POST /v3/alipay/mobile/public/message/total/send).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wyco
func (s *Service) PublicMessageTotalSend(ctx context.Context, req MobilePublicMessageTotalSendRequest, opts ...core.CallOption) (*core.Response[MobilePublicMessageTotalSendResponse], error) {
	return core.ExecuteJSON[MobilePublicMessageTotalSendRequest, MobilePublicMessageTotalSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.public.message.total.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/mobile/public/message/total/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobilePublicMessageTotalSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
