// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package personalized

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

// PublicPersonalizedExtensionBatchQuery calls alipay.open.public.personalized.extension.batchquery (POST /v3/alipay/open/public/personalized/extension/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wcq6
func (s *Service) PublicPersonalizedExtensionBatchQuery(ctx context.Context, req PublicPersonalizedExtensionBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicPersonalizedExtensionBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicPersonalizedExtensionBatchQueryRequest, PublicPersonalizedExtensionBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.personalized.extension.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/personalized/extension/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicPersonalizedExtensionBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicPersonalizedExtensionCreate calls alipay.open.public.personalized.extension.create (POST /v3/alipay/open/public/personalized/extension/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wcpn
func (s *Service) PublicPersonalizedExtensionCreate(ctx context.Context, req PublicPersonalizedExtensionCreateRequest, opts ...core.CallOption) (*core.Response[PublicPersonalizedExtensionCreateResponse], error) {
	return core.ExecuteJSON[PublicPersonalizedExtensionCreateRequest, PublicPersonalizedExtensionCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.personalized.extension.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/personalized/extension/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicPersonalizedExtensionCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicPersonalizedExtensionDelete calls alipay.open.public.personalized.extension.delete (DELETE /v3/alipay/open/public/personalized/extension/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05wydg
func (s *Service) PublicPersonalizedExtensionDelete(ctx context.Context, req PublicPersonalizedExtensionDeleteRequest, opts ...core.CallOption) (*core.Response[PublicPersonalizedExtensionDeleteResponse], error) {
	return core.ExecuteJSON[PublicPersonalizedExtensionDeleteRequest, PublicPersonalizedExtensionDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.personalized.extension.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/public/personalized/extension/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicPersonalizedExtensionDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicPersonalizedMenuCreate calls alipay.open.public.personalized.menu.create (POST /v3/alipay/open/public/personalized/menu/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05x9eg
func (s *Service) PublicPersonalizedMenuCreate(ctx context.Context, req PublicPersonalizedMenuCreateRequest, opts ...core.CallOption) (*core.Response[PublicPersonalizedMenuCreateResponse], error) {
	return core.ExecuteJSON[PublicPersonalizedMenuCreateRequest, PublicPersonalizedMenuCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.personalized.menu.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/personalized/menu/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicPersonalizedMenuCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicPersonalizedMenuDelete calls alipay.open.public.personalized.menu.delete (DELETE /v3/alipay/open/public/personalized/menu/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05xa6t
func (s *Service) PublicPersonalizedMenuDelete(ctx context.Context, req PublicPersonalizedMenuDeleteRequest, opts ...core.CallOption) (*core.Response[PublicPersonalizedMenuDeleteResponse], error) {
	return core.ExecuteJSON[PublicPersonalizedMenuDeleteRequest, PublicPersonalizedMenuDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.personalized.menu.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/public/personalized/menu/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicPersonalizedMenuDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
