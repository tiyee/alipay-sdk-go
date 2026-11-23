// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package user

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

// PublicUserDataBatchQuery calls alipay.open.public.user.data.batchquery (POST /v3/alipay/open/public/user/data/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wum9
func (s *Service) PublicUserDataBatchQuery(ctx context.Context, req PublicUserDataBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicUserDataBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicUserDataBatchQueryRequest, PublicUserDataBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.user.data.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/user/data/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicUserDataBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicUserFollowQuery calls alipay.open.public.user.follow.query (GET /v3/alipay/open/public/user/follow/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wcq8
func (s *Service) PublicUserFollowQuery(ctx context.Context, req PublicUserFollowQueryRequest, opts ...core.CallOption) (*core.Response[PublicUserFollowQueryResponse], error) {
	return core.ExecuteJSON[PublicUserFollowQueryRequest, PublicUserFollowQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.user.follow.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/public/user/follow/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicUserFollowQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
