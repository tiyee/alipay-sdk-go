package template

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

// PublicTemplateMessageDelete calls alipay.mobile.public.template.message.delete (POST /v3/alipay/mobile/public/template/message/delete).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xsen
func (s *Service) PublicTemplateMessageDelete(ctx context.Context, req MobilePublicTemplateMessageDeleteRequest, opts ...core.CallOption) (*core.Response[MobilePublicTemplateMessageDeleteResponse], error) {
	return core.ExecuteJSON[MobilePublicTemplateMessageDeleteRequest, MobilePublicTemplateMessageDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.public.template.message.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/mobile/public/template/message/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobilePublicTemplateMessageDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicTemplateMessageGet calls alipay.mobile.public.template.message.get (GET /v3/alipay/mobile/public/template/message/get).
// Official V3 contract: https://opendocs.alipay.com/apis/05wh9v
func (s *Service) PublicTemplateMessageGet(ctx context.Context, req MobilePublicTemplateMessageGetRequest, opts ...core.CallOption) (*core.Response[MobilePublicTemplateMessageGetResponse], error) {
	return core.ExecuteJSON[MobilePublicTemplateMessageGetRequest, MobilePublicTemplateMessageGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.public.template.message.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/mobile/public/template/message/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobilePublicTemplateMessageGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicTemplateMessageQuery calls alipay.mobile.public.template.message.query (POST /v3/alipay/mobile/public/template/message/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xdxr
func (s *Service) PublicTemplateMessageQuery(ctx context.Context, req MobilePublicTemplateMessageQueryRequest, opts ...core.CallOption) (*core.Response[MobilePublicTemplateMessageQueryResponse], error) {
	return core.ExecuteJSON[MobilePublicTemplateMessageQueryRequest, MobilePublicTemplateMessageQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.public.template.message.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/mobile/public/template/message/query",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobilePublicTemplateMessageQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
