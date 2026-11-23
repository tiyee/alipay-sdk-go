// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package account

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

// PublicAccountCreate calls alipay.open.public.account.create (POST /v3/alipay/open/public/account/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wyjj
func (s *Service) PublicAccountCreate(ctx context.Context, req PublicAccountCreateRequest, opts ...core.CallOption) (*core.Response[PublicAccountCreateResponse], error) {
	return core.ExecuteJSON[PublicAccountCreateRequest, PublicAccountCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.account.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/account/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicAccountCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicAccountDelete calls alipay.open.public.account.delete (DELETE /v3/alipay/open/public/account/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuq3
func (s *Service) PublicAccountDelete(ctx context.Context, req PublicAccountDeleteRequest, opts ...core.CallOption) (*core.Response[PublicAccountDeleteResponse], error) {
	return core.ExecuteJSON[PublicAccountDeleteRequest, PublicAccountDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.account.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/public/account/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicAccountDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicAccountQuery calls alipay.open.public.account.query (GET /v3/alipay/open/public/account/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wupq
func (s *Service) PublicAccountQuery(ctx context.Context, req PublicAccountQueryRequest, opts ...core.CallOption) (*core.Response[PublicAccountQueryResponse], error) {
	return core.ExecuteJSON[PublicAccountQueryRequest, PublicAccountQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.account.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/public/account/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicAccountQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
