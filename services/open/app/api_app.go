package app

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

// AppMessageTemplateSubscribeQuery calls alipay.open.app.messagetemplate.subscribe.query (POST /v3/alipay/open/app/messagetemplate/subscribe/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/af42509c_alipay.open.app.messagetemplate.subscribe.query
func (s *Service) AppMessageTemplateSubscribeQuery(ctx context.Context, req AppMessageTemplateSubscribeQueryRequest, opts ...core.CallOption) (*core.Response[AppMessageTemplateSubscribeQueryResponse], error) {
	return core.ExecuteJSON[AppMessageTemplateSubscribeQueryRequest, AppMessageTemplateSubscribeQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.messagetemplate.subscribe.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/app/messagetemplate/subscribe/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppMessageTemplateSubscribeQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppMiniTemplateMessageSend calls alipay.open.app.mini.templatemessage.send (POST /v3/alipay/open/app/mini/templatemessage/send).
// Official V3 contract: https://opendocs.alipay.com/open-v3/0136e987_alipay.open.app.mini.templatemessage.send
func (s *Service) AppMiniTemplateMessageSend(ctx context.Context, req AppMiniTemplateMessageSendRequest, opts ...core.CallOption) (*core.Response[AppMiniTemplateMessageSendResponse], error) {
	return core.ExecuteJSON[AppMiniTemplateMessageSendRequest, AppMiniTemplateMessageSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.mini.templatemessage.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/app/mini/templatemessage/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppMiniTemplateMessageSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppQRCodeCreate calls alipay.open.app.qrcode.create (POST /v3/alipay/open/app/qrcode/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/76c2627e_alipay.open.app.qrcode.create
func (s *Service) AppQRCodeCreate(ctx context.Context, req AppQRCodeCreateRequest, opts ...core.CallOption) (*core.Response[AppQRCodeCreateResponse], error) {
	return core.ExecuteJSON[AppQRCodeCreateRequest, AppQRCodeCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.qrcode.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/app/qrcode/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppQRCodeCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
