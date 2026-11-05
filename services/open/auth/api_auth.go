// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package auth

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

// AuthAppAESGet calls alipay.open.auth.app.aes.get (GET /v3/alipay/open/auth/app/aes/get).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4d82ebbf_alipay.open.auth.app.aes.get
func (s *Service) AuthAppAESGet(ctx context.Context, req AuthAppAESGetRequest, opts ...core.CallOption) (*core.Response[AuthAppAESGetResponse], error) {
	return core.ExecuteJSON[AuthAppAESGetRequest, AuthAppAESGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.auth.app.aes.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/auth/app/aes/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AuthAppAESGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AuthAppAESSet calls alipay.open.auth.app.aes.set (POST /v3/alipay/open/auth/app/aes/set).
// Official V3 contract: https://opendocs.alipay.com/open-v3/306fd1a6_alipay.open.auth.app.aes.set
func (s *Service) AuthAppAESSet(ctx context.Context, req AuthAppAESSetRequest, opts ...core.CallOption) (*core.Response[AuthAppAESSetResponse], error) {
	return core.ExecuteJSON[AuthAppAESSetRequest, AuthAppAESSetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.auth.app.aes.set", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/auth/app/aes/set",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AuthAppAESSetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AuthTokenAppQuery calls alipay.open.auth.token.app.query (GET /v3/alipay/open/auth/token/app/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b6e7c6b3_alipay.open.auth.token.app.query
func (s *Service) AuthTokenAppQuery(ctx context.Context, req AuthTokenAppQueryRequest, opts ...core.CallOption) (*core.Response[AuthTokenAppQueryResponse], error) {
	return core.ExecuteJSON[AuthTokenAppQueryRequest, AuthTokenAppQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.auth.token.app.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/auth/token/app/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AuthTokenAppQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AuthUserAuthRelationshipQuery calls alipay.open.auth.userauth.relationship.query (GET /v3/alipay/open/auth/userauth/relationship/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/5ea1017e_alipay.open.auth.userauth.relationship.query
func (s *Service) AuthUserAuthRelationshipQuery(ctx context.Context, req AuthUserAuthRelationshipQueryRequest, opts ...core.CallOption) (*core.Response[AuthUserAuthRelationshipQueryResponse], error) {
	return core.ExecuteJSON[AuthUserAuthRelationshipQueryRequest, AuthUserAuthRelationshipQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.auth.userauth.relationship.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/auth/userauth/relationship/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AuthUserAuthRelationshipQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
