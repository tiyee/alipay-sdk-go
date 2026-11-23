// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package category

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

// MiniCategoryQuery calls alipay.open.mini.category.query (GET /v3/alipay/open/mini/category/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/bb00b1a6_alipay.open.mini.category.query
func (s *Service) MiniCategoryQuery(ctx context.Context, req MiniCategoryQueryRequest, opts ...core.CallOption) (*core.Response[MiniCategoryQueryResponse], error) {
	return core.ExecuteJSON[MiniCategoryQueryRequest, MiniCategoryQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.category.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/category/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniCategoryQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniCategoryRequireQuery calls alipay.open.mini.category.require.query (POST /v3/alipay/open/mini/category/require/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/7f3621b8_alipay.open.mini.category.require.query
func (s *Service) MiniCategoryRequireQuery(ctx context.Context, req MiniCategoryRequireQueryRequest, opts ...core.CallOption) (*core.Response[MiniCategoryRequireQueryResponse], error) {
	return core.ExecuteJSON[MiniCategoryRequireQueryRequest, MiniCategoryRequireQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.category.require.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/category/require/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniCategoryRequireQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
