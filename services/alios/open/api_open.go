// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package open

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service invokes V3 operations in this domain.
type Service struct{ executor core.Executor }

// NewService creates a service using the supplied executor.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

func execute[Req, Resp any](s *Service, ctx context.Context, method string, httpMethod core.HTTPMethod, path string, mode core.ParameterMode, req Req, opts []core.CallOption) (*core.Response[Resp], error) {
	op, err := core.NewJSONOperation[Req, Resp](core.OperationSpec{Method: method, HTTPMethod: httpMethod, Path: path, ParameterMode: mode, AuthRequirement: core.AuthRequirementAny})
	if err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, op, core.Request[Req]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AutoInfoQuery calls alios.open.auto.info.query.
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05igwf
func (s *Service) AutoInfoQuery(ctx context.Context, req AutoInfoQueryRequest, opts ...core.CallOption) (*core.Response[AutoInfoQueryResponse], error) {
	return execute[AutoInfoQueryRequest, AutoInfoQueryResponse](s, ctx, "alios.open.auto.info.query", core.HTTPMethodGet, "/v3/alios/open/auto/info/query", core.ParameterModeQuery, req, opts)
}
