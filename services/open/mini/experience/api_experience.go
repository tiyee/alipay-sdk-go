// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package experience

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

// MiniExperienceCancel calls alipay.open.mini.experience.cancel (POST /v3/alipay/open/mini/experience/cancel).
// Official V3 contract: https://opendocs.alipay.com/open-v3/d4e75c29_alipay.open.mini.experience.cancel
func (s *Service) MiniExperienceCancel(ctx context.Context, req MiniExperienceCancelRequest, opts ...core.CallOption) (*core.Response[MiniExperienceCancelResponse], error) {
	return core.ExecuteJSON[MiniExperienceCancelRequest, MiniExperienceCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.experience.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/experience/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniExperienceCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniExperienceCreate calls alipay.open.mini.experience.create (POST /v3/alipay/open/mini/experience/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/287a2040_alipay.open.mini.experience.create
func (s *Service) MiniExperienceCreate(ctx context.Context, req MiniExperienceCreateRequest, opts ...core.CallOption) (*core.Response[MiniExperienceCreateResponse], error) {
	return core.ExecuteJSON[MiniExperienceCreateRequest, MiniExperienceCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.experience.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/experience/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniExperienceCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniExperienceQuery calls alipay.open.mini.experience.query (GET /v3/alipay/open/mini/experience/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/99c5f4ae_alipay.open.mini.experience.query
func (s *Service) MiniExperienceQuery(ctx context.Context, req MiniExperienceQueryRequest, opts ...core.CallOption) (*core.Response[MiniExperienceQueryResponse], error) {
	return core.ExecuteJSON[MiniExperienceQueryRequest, MiniExperienceQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.experience.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/experience/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniExperienceQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
