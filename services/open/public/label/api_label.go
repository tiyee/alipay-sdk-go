// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package label

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

// PublicLabelCreate calls alipay.open.public.label.create (POST /v3/alipay/open/public/label/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wupo
func (s *Service) PublicLabelCreate(ctx context.Context, req PublicLabelCreateRequest, opts ...core.CallOption) (*core.Response[PublicLabelCreateResponse], error) {
	return core.ExecuteJSON[PublicLabelCreateRequest, PublicLabelCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.label.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/label/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLabelCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicLabelQuery calls alipay.open.public.label.query (GET /v3/alipay/open/public/label/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xq3d
func (s *Service) PublicLabelQuery(ctx context.Context, req PublicLabelQueryRequest, opts ...core.CallOption) (*core.Response[PublicLabelQueryResponse], error) {
	return core.ExecuteJSON[PublicLabelQueryRequest, PublicLabelQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.label.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/public/label/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLabelQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicLabelUserQuery calls alipay.open.public.label.user.query (GET /v3/alipay/open/public/label/user/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wum4
func (s *Service) PublicLabelUserQuery(ctx context.Context, req PublicLabelUserQueryRequest, opts ...core.CallOption) (*core.Response[PublicLabelUserQueryResponse], error) {
	return core.ExecuteJSON[PublicLabelUserQueryRequest, PublicLabelUserQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.label.user.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/public/label/user/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLabelUserQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
