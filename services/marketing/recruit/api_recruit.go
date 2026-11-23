// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package recruit

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

// RecruitEnrollClose calls alipay.marketing.recruit.enroll.close (POST /v3/alipay/marketing/recruit/enroll/close).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xe2w
func (s *Service) RecruitEnrollClose(ctx context.Context, req RecruitEnrollCloseRequest, opts ...core.CallOption) (*core.Response[RecruitEnrollCloseResponse], error) {
	return core.ExecuteJSON[RecruitEnrollCloseRequest, RecruitEnrollCloseResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.recruit.enroll.close", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/recruit/enroll/close",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[RecruitEnrollCloseRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// RecruitEnrollCreate calls alipay.marketing.recruit.enroll.create (POST /v3/alipay/marketing/recruit/enroll/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xqxy
func (s *Service) RecruitEnrollCreate(ctx context.Context, req RecruitEnrollCreateRequest, opts ...core.CallOption) (*core.Response[RecruitEnrollCreateResponse], error) {
	return core.ExecuteJSON[RecruitEnrollCreateRequest, RecruitEnrollCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.recruit.enroll.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/recruit/enroll/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[RecruitEnrollCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// RecruitEnrollQuery calls alipay.marketing.recruit.enroll.query (GET /v3/alipay/marketing/recruit/enroll/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xk7o
func (s *Service) RecruitEnrollQuery(ctx context.Context, req RecruitEnrollQueryRequest, opts ...core.CallOption) (*core.Response[RecruitEnrollQueryResponse], error) {
	return core.ExecuteJSON[RecruitEnrollQueryRequest, RecruitEnrollQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.recruit.enroll.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/recruit/enroll/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[RecruitEnrollQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// RecruitPlanQuery calls alipay.marketing.recruit.plan.query (POST /v3/alipay/marketing/recruit/plan/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xhns
func (s *Service) RecruitPlanQuery(ctx context.Context, req RecruitPlanQueryRequest, opts ...core.CallOption) (*core.Response[RecruitPlanQueryResponse], error) {
	return core.ExecuteJSON[RecruitPlanQueryRequest, RecruitPlanQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.recruit.plan.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/recruit/plan/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[RecruitPlanQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// RecruitPlanListQuery calls alipay.marketing.recruit.planlist.query (POST /v3/alipay/marketing/recruit/planlist/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xk3f
func (s *Service) RecruitPlanListQuery(ctx context.Context, req RecruitPlanListQueryRequest, opts ...core.CallOption) (*core.Response[RecruitPlanListQueryResponse], error) {
	return core.ExecuteJSON[RecruitPlanListQueryRequest, RecruitPlanListQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.recruit.planlist.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/recruit/planlist/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[RecruitPlanListQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
