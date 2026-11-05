// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package delivery

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

// ActivityDeliveryCreate calls alipay.marketing.activity.delivery.create (POST /v3/alipay/marketing/delivery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/43363a7a_alipay.marketing.activity.delivery.create
func (s *Service) ActivityDeliveryCreate(ctx context.Context, req ActivityDeliveryCreateRequest, opts ...core.CallOption) (*core.Response[ActivityDeliveryCreateResponse], error) {
	return core.ExecuteJSON[ActivityDeliveryCreateRequest, ActivityDeliveryCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.delivery.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/delivery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityDeliveryCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityDeliveryQuery calls alipay.marketing.activity.delivery.query (POST /v3/alipay/marketing/delivery/{delivery_id}/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/cab7f6c3_alipay.marketing.activity.delivery.query
func (s *Service) ActivityDeliveryQuery(ctx context.Context, req ActivityDeliveryQueryRequest, opts ...core.CallOption) (*core.Response[ActivityDeliveryQueryResponse], error) {
	return core.ExecuteJSON[ActivityDeliveryQueryRequest, ActivityDeliveryQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.delivery.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/delivery/{delivery_id}/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityDeliveryQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityDeliveryStop calls alipay.marketing.activity.delivery.stop (PATCH /v3/alipay/marketing/delivery/{delivery_id}/stop).
// Official V3 contract: https://opendocs.alipay.com/open-v3/dbc77a2d_alipay.marketing.activity.delivery.stop
func (s *Service) ActivityDeliveryStop(ctx context.Context, req ActivityDeliveryStopRequest, opts ...core.CallOption) (*core.Response[ActivityDeliveryStopResponse], error) {
	return core.ExecuteJSON[ActivityDeliveryStopRequest, ActivityDeliveryStopResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.delivery.stop", HTTPMethod: core.HTTPMethodPatch, Path: "/v3/alipay/marketing/delivery/{delivery_id}/stop",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityDeliveryStopRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
