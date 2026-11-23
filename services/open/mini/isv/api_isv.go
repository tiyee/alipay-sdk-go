// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package isv

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

// MiniISVCreate calls alipay.open.mini.isv.create (POST /v3/alipay/open/mini/isv/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/719ad502_alipay.open.mini.isv.create
func (s *Service) MiniISVCreate(ctx context.Context, req MiniISVCreateRequest, opts ...core.CallOption) (*core.Response[MiniISVCreateResponse], error) {
	return core.ExecuteJSON[MiniISVCreateRequest, MiniISVCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.isv.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/isv/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniISVCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniISVFastRegisterCreate calls alipay.open.mini.isv.fastregister.create (POST /v3/alipay/open/mini/isv/fastregister/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/58d4f210_alipay.open.mini.isv.fastregister.create
func (s *Service) MiniISVFastRegisterCreate(ctx context.Context, req MiniISVFastRegisterCreateRequest, opts ...core.CallOption) (*core.Response[MiniISVFastRegisterCreateResponse], error) {
	return core.ExecuteJSON[MiniISVFastRegisterCreateRequest, MiniISVFastRegisterCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.isv.fastregister.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/isv/fastregister/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniISVFastRegisterCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniISVFastRegisterQuery calls alipay.open.mini.isv.fastregister.query (GET /v3/alipay/open/mini/isv/fastregister/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/033e1882_alipay.open.mini.isv.fastregister.query
func (s *Service) MiniISVFastRegisterQuery(ctx context.Context, req MiniISVFastRegisterQueryRequest, opts ...core.CallOption) (*core.Response[MiniISVFastRegisterQueryResponse], error) {
	return core.ExecuteJSON[MiniISVFastRegisterQueryRequest, MiniISVFastRegisterQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.isv.fastregister.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/isv/fastregister/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniISVFastRegisterQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniISVQuery calls alipay.open.mini.isv.query (GET /v3/alipay/open/mini/isv/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/f6bafc97_alipay.open.mini.isv.query
func (s *Service) MiniISVQuery(ctx context.Context, req MiniISVQueryRequest, opts ...core.CallOption) (*core.Response[MiniISVQueryResponse], error) {
	return core.ExecuteJSON[MiniISVQueryRequest, MiniISVQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.isv.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/isv/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniISVQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
