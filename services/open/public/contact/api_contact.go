// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package contact

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

// PublicContactFollowBatchQuery calls alipay.open.public.contact.follow.batchquery (POST /v3/alipay/open/public/contact/follow/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wnsz
func (s *Service) PublicContactFollowBatchQuery(ctx context.Context, req PublicContactFollowBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicContactFollowBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicContactFollowBatchQueryRequest, PublicContactFollowBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.contact.follow.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/contact/follow/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicContactFollowBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
