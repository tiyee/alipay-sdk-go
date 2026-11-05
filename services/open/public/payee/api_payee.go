// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package payee

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

// PublicPayeeBindCreate calls alipay.open.public.payee.bind.create (POST /v3/alipay/open/public/payee/bind/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvrh
func (s *Service) PublicPayeeBindCreate(ctx context.Context, req PublicPayeeBindCreateRequest, opts ...core.CallOption) (*core.Response[PublicPayeeBindCreateResponse], error) {
	return core.ExecuteJSON[PublicPayeeBindCreateRequest, PublicPayeeBindCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.payee.bind.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/payee/bind/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicPayeeBindCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicPayeeBindDelete calls alipay.open.public.payee.bind.delete (DELETE /v3/alipay/open/public/payee/bind/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05xq3f
func (s *Service) PublicPayeeBindDelete(ctx context.Context, req PublicPayeeBindDeleteRequest, opts ...core.CallOption) (*core.Response[PublicPayeeBindDeleteResponse], error) {
	return core.ExecuteJSON[PublicPayeeBindDeleteRequest, PublicPayeeBindDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.payee.bind.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/public/payee/bind/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicPayeeBindDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
