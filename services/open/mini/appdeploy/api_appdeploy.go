// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package appdeploy

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

// MiniAppDeployByAppVersionQuery calls alipay.open.mini.appdeploy.byappversion.query (GET /v3/alipay/open/mini/appdeploy/byappversion/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xo5a
func (s *Service) MiniAppDeployByAppVersionQuery(ctx context.Context, req MiniAppDeployByAppVersionQueryRequest, opts ...core.CallOption) (*core.Response[MiniAppDeployByAppVersionQueryResponse], error) {
	return core.ExecuteJSON[MiniAppDeployByAppVersionQueryRequest, MiniAppDeployByAppVersionQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.appdeploy.byappversion.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/appdeploy/byappversion/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniAppDeployByAppVersionQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniAppDeployByDeployVersionQuery calls alipay.open.mini.appdeploy.bydeployversion.query (GET /v3/alipay/open/mini/appdeploy/bydeployversion/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xaay
func (s *Service) MiniAppDeployByDeployVersionQuery(ctx context.Context, req MiniAppDeployByDeployVersionQueryRequest, opts ...core.CallOption) (*core.Response[MiniAppDeployByDeployVersionQueryResponse], error) {
	return core.ExecuteJSON[MiniAppDeployByDeployVersionQueryRequest, MiniAppDeployByDeployVersionQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.appdeploy.bydeployversion.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/appdeploy/bydeployversion/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniAppDeployByDeployVersionQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
