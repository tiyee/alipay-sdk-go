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

// MiniTemplateMarketingCreate calls alipay.open.mini.template.marketing.create (POST /v3/alipay/open/mini/template/marketing/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/5c89642c_alipay.open.mini.template.marketing.create
func (s *Service) MiniTemplateMarketingCreate(ctx context.Context, req MiniTemplateMarketingCreateRequest, opts ...core.CallOption) (*core.Response[MiniTemplateMarketingCreateResponse], error) {
	return core.ExecuteJSON[MiniTemplateMarketingCreateRequest, MiniTemplateMarketingCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.template.marketing.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/template/marketing/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniTemplateMarketingCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniTemplateUsageQuery calls alipay.open.mini.template.usage.query (GET /v3/alipay/open/mini/template/usage/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/883d84eb_alipay.open.mini.template.usage.query
func (s *Service) MiniTemplateUsageQuery(ctx context.Context, req MiniTemplateUsageQueryRequest, opts ...core.CallOption) (*core.Response[MiniTemplateUsageQueryResponse], error) {
	return core.ExecuteJSON[MiniTemplateUsageQueryRequest, MiniTemplateUsageQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.template.usage.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/template/usage/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniTemplateUsageQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
