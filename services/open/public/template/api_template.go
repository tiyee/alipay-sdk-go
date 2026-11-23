// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package template

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

// PublicTemplateMessageGet calls alipay.open.public.template.message.get (GET /v3/alipay/open/public/template/message/get).
// Official V3 contract: https://opendocs.alipay.com/apis/05wbw6
func (s *Service) PublicTemplateMessageGet(ctx context.Context, req PublicTemplateMessageGetRequest, opts ...core.CallOption) (*core.Response[PublicTemplateMessageGetResponse], error) {
	return core.ExecuteJSON[PublicTemplateMessageGetRequest, PublicTemplateMessageGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.template.message.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/public/template/message/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicTemplateMessageGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicTemplateMessageIndustryModify calls alipay.open.public.template.message.industry.modify (POST /v3/alipay/open/public/template/message/industry/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuhh
func (s *Service) PublicTemplateMessageIndustryModify(ctx context.Context, req PublicTemplateMessageIndustryModifyRequest, opts ...core.CallOption) (*core.Response[PublicTemplateMessageIndustryModifyResponse], error) {
	return core.ExecuteJSON[PublicTemplateMessageIndustryModifyRequest, PublicTemplateMessageIndustryModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.template.message.industry.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/template/message/industry/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicTemplateMessageIndustryModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
