// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package shop

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

// ActivityShopBatchQuery calls alipay.marketing.activity.shop.batchquery (GET /v3/alipay/marketing/activity/{activity_id}/shop/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/ebd50e50_alipay.marketing.activity.shop.batchquery
func (s *Service) ActivityShopBatchQuery(ctx context.Context, req ActivityShopBatchQueryRequest, opts ...core.CallOption) (*core.Response[ActivityShopBatchQueryResponse], error) {
	return core.ExecuteJSON[ActivityShopBatchQueryRequest, ActivityShopBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.shop.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/activity/{activity_id}/shop/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityShopBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
