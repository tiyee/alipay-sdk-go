// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package inner

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

// MiniInnerExperienceCancel calls alipay.open.mini.inner.experience.cancel (POST /v3/alipay/open/mini/inner/experience/cancel).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xe27
func (s *Service) MiniInnerExperienceCancel(ctx context.Context, req MiniInnerExperienceCancelRequest, opts ...core.CallOption) (*core.Response[MiniInnerExperienceCancelResponse], error) {
	return core.ExecuteJSON[MiniInnerExperienceCancelRequest, MiniInnerExperienceCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.inner.experience.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/inner/experience/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerExperienceCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerMembersAdd calls alipay.open.mini.inner.members.add (POST /v3/alipay/open/mini/inner/members/add).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xhp2
func (s *Service) MiniInnerMembersAdd(ctx context.Context, req MiniInnerMembersAddRequest, opts ...core.CallOption) (*core.Response[MiniInnerMembersAddResponse], error) {
	return core.ExecuteJSON[MiniInnerMembersAddRequest, MiniInnerMembersAddResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.inner.members.add", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/inner/members/add",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerMembersAddRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerMembersDelete calls alipay.open.mini.inner.members.delete (DELETE /v3/alipay/open/mini/inner/members/delete).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xqzc
func (s *Service) MiniInnerMembersDelete(ctx context.Context, req MiniInnerMembersDeleteRequest, opts ...core.CallOption) (*core.Response[MiniInnerMembersDeleteResponse], error) {
	return core.ExecuteJSON[MiniInnerMembersDeleteRequest, MiniInnerMembersDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.inner.members.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/mini/inner/members/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerMembersDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerMembersQuery calls alipay.open.mini.inner.members.query (GET /v3/alipay/open/mini/inner/members/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xaad
func (s *Service) MiniInnerMembersQuery(ctx context.Context, req MiniInnerMembersQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerMembersQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerMembersQueryRequest, MiniInnerMembersQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.inner.members.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/inner/members/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerMembersQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
