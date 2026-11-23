// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package setting

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

// PublicSettingCategoryQuery calls alipay.open.public.setting.category.query (GET /v3/alipay/open/public/setting/category/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05xdxp
func (s *Service) PublicSettingCategoryQuery(ctx context.Context, req PublicSettingCategoryQueryRequest, opts ...core.CallOption) (*core.Response[PublicSettingCategoryQueryResponse], error) {
	return core.ExecuteJSON[PublicSettingCategoryQueryRequest, PublicSettingCategoryQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.setting.category.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/public/setting/category/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicSettingCategoryQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
