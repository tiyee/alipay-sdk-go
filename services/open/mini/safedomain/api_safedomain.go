// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package safedomain

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

// MiniSafeDomainCreate calls alipay.open.mini.safedomain.create (POST /v3/alipay/open/mini/safedomain/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/27ccb222_alipay.open.mini.safedomain.create
func (s *Service) MiniSafeDomainCreate(ctx context.Context, req MiniSafeDomainCreateRequest, opts ...core.CallOption) (*core.Response[MiniSafeDomainCreateResponse], error) {
	return core.ExecuteJSON[MiniSafeDomainCreateRequest, MiniSafeDomainCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.safedomain.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/safedomain/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniSafeDomainCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniSafeDomainDelete calls alipay.open.mini.safedomain.delete (DELETE /v3/alipay/open/mini/safedomain/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/d96d3a21_alipay.open.mini.safedomain.delete
func (s *Service) MiniSafeDomainDelete(ctx context.Context, req MiniSafeDomainDeleteRequest, opts ...core.CallOption) (*core.Response[MiniSafeDomainDeleteResponse], error) {
	return core.ExecuteJSON[MiniSafeDomainDeleteRequest, MiniSafeDomainDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.safedomain.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/mini/safedomain/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniSafeDomainDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
