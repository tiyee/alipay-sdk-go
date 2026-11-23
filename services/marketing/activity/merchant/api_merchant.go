// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

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

// ActivityMerchantBatchQuery calls alipay.marketing.activity.merchant.batchquery (GET /v3/alipay/marketing/activity/{activity_id}/merchant/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/30e4ba35_alipay.marketing.activity.merchant.batchquery
func (s *Service) ActivityMerchantBatchQuery(ctx context.Context, req ActivityMerchantBatchQueryRequest, opts ...core.CallOption) (*core.Response[ActivityMerchantBatchQueryResponse], error) {
	return core.ExecuteJSON[ActivityMerchantBatchQueryRequest, ActivityMerchantBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.merchant.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/activity/{activity_id}/merchant/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityMerchantBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
