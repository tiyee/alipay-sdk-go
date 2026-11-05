// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package batchquery

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

// ActivityBatchQuery calls alipay.marketing.activity.batchquery (POST /v3/alipay/marketing/activity/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/05355ed4_alipay.marketing.activity.batchquery
func (s *Service) ActivityBatchQuery(ctx context.Context, req ActivityBatchQueryRequest, opts ...core.CallOption) (*core.Response[ActivityBatchQueryResponse], error) {
	return core.ExecuteJSON[ActivityBatchQueryRequest, ActivityBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
