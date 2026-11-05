// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package tips

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

// MiniTipsDeliveryBatchQuery calls alipay.open.mini.tips.delivery.batchquery (POST /v3/alipay/open/mini/tips/delivery/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/75394e37_alipay.open.mini.tips.delivery.batchquery
func (s *Service) MiniTipsDeliveryBatchQuery(ctx context.Context, req MiniTipsDeliveryBatchQueryRequest, opts ...core.CallOption) (*core.Response[MiniTipsDeliveryBatchQueryResponse], error) {
	return core.ExecuteJSON[MiniTipsDeliveryBatchQueryRequest, MiniTipsDeliveryBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.tips.delivery.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/tips/delivery/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniTipsDeliveryBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniTipsDeliveryCreate calls alipay.open.mini.tips.delivery.create (POST /v3/alipay/open/mini/tips/delivery/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/ca9b74a9_alipay.open.mini.tips.delivery.create
func (s *Service) MiniTipsDeliveryCreate(ctx context.Context, req MiniTipsDeliveryCreateRequest, opts ...core.CallOption) (*core.Response[MiniTipsDeliveryCreateResponse], error) {
	return core.ExecuteJSON[MiniTipsDeliveryCreateRequest, MiniTipsDeliveryCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.tips.delivery.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/tips/delivery/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniTipsDeliveryCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniTipsDeliveryModify calls alipay.open.mini.tips.delivery.modify (POST /v3/alipay/open/mini/tips/delivery/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a4b55863_alipay.open.mini.tips.delivery.modify
func (s *Service) MiniTipsDeliveryModify(ctx context.Context, req MiniTipsDeliveryModifyRequest, opts ...core.CallOption) (*core.Response[MiniTipsDeliveryModifyResponse], error) {
	return core.ExecuteJSON[MiniTipsDeliveryModifyRequest, MiniTipsDeliveryModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.tips.delivery.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/tips/delivery/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniTipsDeliveryModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniTipsDeliveryQuery calls alipay.open.mini.tips.delivery.query (GET /v3/alipay/open/mini/tips/delivery/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/f59fb71e_alipay.open.mini.tips.delivery.query
func (s *Service) MiniTipsDeliveryQuery(ctx context.Context, req MiniTipsDeliveryQueryRequest, opts ...core.CallOption) (*core.Response[MiniTipsDeliveryQueryResponse], error) {
	return core.ExecuteJSON[MiniTipsDeliveryQueryRequest, MiniTipsDeliveryQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.tips.delivery.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/tips/delivery/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniTipsDeliveryQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniTipsStatisticQuery calls alipay.open.mini.tips.statistic.query (GET /v3/alipay/open/mini/tips/statistic/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/87364a7f_alipay.open.mini.tips.statistic.query
func (s *Service) MiniTipsStatisticQuery(ctx context.Context, req MiniTipsStatisticQueryRequest, opts ...core.CallOption) (*core.Response[MiniTipsStatisticQueryResponse], error) {
	return core.ExecuteJSON[MiniTipsStatisticQueryRequest, MiniTipsStatisticQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.tips.statistic.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/tips/statistic/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniTipsStatisticQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
