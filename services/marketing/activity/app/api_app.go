// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package app

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

// ActivityAppBatchQuery calls alipay.marketing.activity.app.batchquery (GET /v3/alipay/marketing/activity/{activity_id}/app/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c8f333a8_alipay.marketing.activity.app.batchquery
func (s *Service) ActivityAppBatchQuery(ctx context.Context, req ActivityAppBatchQueryRequest, opts ...core.CallOption) (*core.Response[ActivityAppBatchQueryResponse], error) {
	return core.ExecuteJSON[ActivityAppBatchQueryRequest, ActivityAppBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.app.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/activity/{activity_id}/app/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityAppBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
