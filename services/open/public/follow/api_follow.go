// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package follow

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

// PublicFollowBatchQuery calls alipay.open.public.follow.batchquery (POST /v3/alipay/open/public/follow/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wcpt
func (s *Service) PublicFollowBatchQuery(ctx context.Context, req PublicFollowBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicFollowBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicFollowBatchQueryRequest, PublicFollowBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.follow.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/follow/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicFollowBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
