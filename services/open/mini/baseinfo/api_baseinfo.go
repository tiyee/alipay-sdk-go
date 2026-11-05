// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package baseinfo

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

// MiniBaseInfoModify calls alipay.open.mini.baseinfo.modify (POST /v3/alipay/open/mini/baseinfo/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/762743d3_alipay.open.mini.baseinfo.modify
func (s *Service) MiniBaseInfoModify(ctx context.Context, req MiniBaseInfoModifyRequest, opts ...core.CallOption) (*core.Response[MiniBaseInfoModifyResponse], error) {
	return core.ExecuteJSON[MiniBaseInfoModifyRequest, MiniBaseInfoModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.baseinfo.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/baseinfo/modify",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniBaseInfoModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniBaseInfoNameCheck calls alipay.open.mini.baseinfo.name.check (GET /v3/alipay/open/mini/baseinfo/name/check).
// Official V3 contract: https://opendocs.alipay.com/open-v3/bf9816ac_alipay.open.mini.baseinfo.name.check
func (s *Service) MiniBaseInfoNameCheck(ctx context.Context, req MiniBaseInfoNameCheckRequest, opts ...core.CallOption) (*core.Response[MiniBaseInfoNameCheckResponse], error) {
	return core.ExecuteJSON[MiniBaseInfoNameCheckRequest, MiniBaseInfoNameCheckResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.baseinfo.name.check", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/baseinfo/name/check",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniBaseInfoNameCheckRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniBaseInfoQuery calls alipay.open.mini.baseinfo.query (GET /v3/alipay/open/mini/baseinfo/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c5995fc4_alipay.open.mini.baseinfo.query
func (s *Service) MiniBaseInfoQuery(ctx context.Context, req MiniBaseInfoQueryRequest, opts ...core.CallOption) (*core.Response[MiniBaseInfoQueryResponse], error) {
	return core.ExecuteJSON[MiniBaseInfoQueryRequest, MiniBaseInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.baseinfo.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/baseinfo/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniBaseInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
