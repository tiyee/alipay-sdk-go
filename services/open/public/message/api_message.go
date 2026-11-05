// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package message

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service executes one alipay V3 service group.
type Service struct{ executor core.Executor }

// NewService constructs the service with a shared V3 executor.
// NewService constructs the service with a shared V3 executor.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

// PublicMessageContentCreate calls alipay.open.public.message.content.create (POST /v3/alipay/open/public/message/content/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05xhmr
func (s *Service) PublicMessageContentCreate(ctx context.Context, req PublicMessageContentCreateRequest, opts ...core.CallOption) (*core.Response[PublicMessageContentCreateResponse], error) {
	return core.ExecuteJSON[PublicMessageContentCreateRequest, PublicMessageContentCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.message.content.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/message/content/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMessageContentCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMessageContentModify calls alipay.open.public.message.content.modify (POST /v3/alipay/open/public/message/content/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuj5
func (s *Service) PublicMessageContentModify(ctx context.Context, req PublicMessageContentModifyRequest, opts ...core.CallOption) (*core.Response[PublicMessageContentModifyResponse], error) {
	return core.ExecuteJSON[PublicMessageContentModifyRequest, PublicMessageContentModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.message.content.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/message/content/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMessageContentModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMessageCustomSend calls alipay.open.public.message.custom.send (POST /v3/alipay/open/public/message/custom/send).
// Official V3 contract: https://opendocs.alipay.com/apis/05wcpx
func (s *Service) PublicMessageCustomSend(ctx context.Context, req PublicMessageCustomSendRequest, opts ...core.CallOption) (*core.Response[PublicMessageCustomSendResponse], error) {
	return core.ExecuteJSON[PublicMessageCustomSendRequest, PublicMessageCustomSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.message.custom.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/message/custom/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMessageCustomSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMessageGroupSend calls alipay.open.public.message.group.send (POST /v3/alipay/open/public/message/group/send).
// Official V3 contract: https://opendocs.alipay.com/apis/05wyfq
func (s *Service) PublicMessageGroupSend(ctx context.Context, req PublicMessageGroupSendRequest, opts ...core.CallOption) (*core.Response[PublicMessageGroupSendResponse], error) {
	return core.ExecuteJSON[PublicMessageGroupSendRequest, PublicMessageGroupSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.message.group.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/message/group/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMessageGroupSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMessageLabelSend calls alipay.open.public.message.label.send (POST /v3/alipay/open/public/message/label/send).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xlk0
func (s *Service) PublicMessageLabelSend(ctx context.Context, req PublicMessageLabelSendRequest, opts ...core.CallOption) (*core.Response[PublicMessageLabelSendResponse], error) {
	return core.ExecuteJSON[PublicMessageLabelSendRequest, PublicMessageLabelSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.message.label.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/message/label/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMessageLabelSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMessageSingleSend calls alipay.open.public.message.single.send (POST /v3/alipay/open/public/message/single/send).
// Official V3 contract: https://opendocs.alipay.com/apis/06bpqn
func (s *Service) PublicMessageSingleSend(ctx context.Context, req PublicMessageSingleSendRequest, opts ...core.CallOption) (*core.Response[PublicMessageSingleSendResponse], error) {
	return core.ExecuteJSON[PublicMessageSingleSendRequest, PublicMessageSingleSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.message.single.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/message/single/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMessageSingleSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMessageTotalSend calls alipay.open.public.message.total.send (POST /v3/alipay/open/public/message/total/send).
// Official V3 contract: https://opendocs.alipay.com/apis/05wnte
func (s *Service) PublicMessageTotalSend(ctx context.Context, req PublicMessageTotalSendRequest, opts ...core.CallOption) (*core.Response[PublicMessageTotalSendResponse], error) {
	return core.ExecuteJSON[PublicMessageTotalSendRequest, PublicMessageTotalSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.message.total.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/message/total/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMessageTotalSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
