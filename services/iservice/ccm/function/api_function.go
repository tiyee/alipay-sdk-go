// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package function

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

// CcmFunctionQuery calls alipay.iservice.ccm.function.query (GET /v3/alipay/iservice/ccm/function/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/45b96512_alipay.iservice.ccm.function.query
func (s *Service) CcmFunctionQuery(ctx context.Context, req CcmFunctionQueryRequest, opts ...core.CallOption) (*core.Response[CcmFunctionQueryResponse], error) {
	return core.ExecuteJSON[CcmFunctionQueryRequest, CcmFunctionQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.function.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/function/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmFunctionQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
