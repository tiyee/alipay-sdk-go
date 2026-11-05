// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package signflows

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

// SignFlowsDetailQuery calls alipay.eco.signflows.detail.query (GET /v3/alipay/eco/signflows/detail/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wylm
func (s *Service) SignFlowsDetailQuery(ctx context.Context, req SignFlowsDetailQueryRequest, opts ...core.CallOption) (*core.Response[SignFlowsDetailQueryResponse], error) {
	return core.ExecuteJSON[SignFlowsDetailQueryRequest, SignFlowsDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.signflows.detail.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/eco/signflows/detail/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SignFlowsDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SignFlowsURLQuery calls alipay.eco.signflows.url.query (GET /v3/alipay/eco/signflows/url/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wyc8
func (s *Service) SignFlowsURLQuery(ctx context.Context, req SignFlowsURLQueryRequest, opts ...core.CallOption) (*core.Response[SignFlowsURLQueryResponse], error) {
	return core.ExecuteJSON[SignFlowsURLQueryRequest, SignFlowsURLQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.signflows.url.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/eco/signflows/url/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SignFlowsURLQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
