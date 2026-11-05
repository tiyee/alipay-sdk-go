// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package query

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

// ActivityQuery calls alipay.marketing.activity.query (GET /v3/alipay/marketing/activity/{activity_id}).
// Official V3 contract: https://opendocs.alipay.com/open-v3/168d324a_alipay.marketing.activity.query
func (s *Service) ActivityQuery(ctx context.Context, req ActivityQueryRequest, opts ...core.CallOption) (*core.Response[ActivityQueryResponse], error) {
	return core.ExecuteJSON[ActivityQueryRequest, ActivityQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/activity/{activity_id}",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
