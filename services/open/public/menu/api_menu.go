// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package menu

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

// PublicMenuBatchQuery calls alipay.open.public.menu.batchquery (POST /v3/alipay/open/public/menu/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wol5
func (s *Service) PublicMenuBatchQuery(ctx context.Context, req PublicMenuBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicMenuBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicMenuBatchQueryRequest, PublicMenuBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.menu.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/menu/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMenuBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMenuCreate calls alipay.open.public.menu.create (POST /v3/alipay/open/public/menu/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvxt
func (s *Service) PublicMenuCreate(ctx context.Context, req PublicMenuCreateRequest, opts ...core.CallOption) (*core.Response[PublicMenuCreateResponse], error) {
	return core.ExecuteJSON[PublicMenuCreateRequest, PublicMenuCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.menu.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/menu/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMenuCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMenuDataBatchQuery calls alipay.open.public.menu.data.batchquery (POST /v3/alipay/open/public/menu/data/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wyh5
func (s *Service) PublicMenuDataBatchQuery(ctx context.Context, req PublicMenuDataBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicMenuDataBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicMenuDataBatchQueryRequest, PublicMenuDataBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.menu.data.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/menu/data/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMenuDataBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMenuModify calls alipay.open.public.menu.modify (POST /v3/alipay/open/public/menu/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05xgci
func (s *Service) PublicMenuModify(ctx context.Context, req PublicMenuModifyRequest, opts ...core.CallOption) (*core.Response[PublicMenuModifyResponse], error) {
	return core.ExecuteJSON[PublicMenuModifyRequest, PublicMenuModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.menu.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/menu/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMenuModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMenuQuery calls alipay.open.public.menu.query (GET /v3/alipay/open/public/menu/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wuni
func (s *Service) PublicMenuQuery(ctx context.Context, req PublicMenuQueryRequest, opts ...core.CallOption) (*core.Response[PublicMenuQueryResponse], error) {
	return core.ExecuteJSON[PublicMenuQueryRequest, PublicMenuQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.menu.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/public/menu/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMenuQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
