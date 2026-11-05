// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package cityfacilitator

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

// CityFacilitatorStationQuery calls alipay.commerce.cityfacilitator.station.query (GET /v3/alipay/commerce/cityfacilitator/station/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvuo
func (s *Service) CityFacilitatorStationQuery(ctx context.Context, req CityFacilitatorStationQueryRequest, opts ...core.CallOption) (*core.Response[CityFacilitatorStationQueryResponse], error) {
	return core.ExecuteJSON[CityFacilitatorStationQueryRequest, CityFacilitatorStationQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.cityfacilitator.station.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/cityfacilitator/station/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CityFacilitatorStationQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CityFacilitatorVoucherBatchQuery calls alipay.commerce.cityfacilitator.voucher.batchquery (POST /v3/alipay/commerce/cityfacilitator/voucher/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wol9
func (s *Service) CityFacilitatorVoucherBatchQuery(ctx context.Context, req CityFacilitatorVoucherBatchQueryRequest, opts ...core.CallOption) (*core.Response[CityFacilitatorVoucherBatchQueryResponse], error) {
	return core.ExecuteJSON[CityFacilitatorVoucherBatchQueryRequest, CityFacilitatorVoucherBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.cityfacilitator.voucher.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/cityfacilitator/voucher/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CityFacilitatorVoucherBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CityFacilitatorVoucherGenerate calls alipay.commerce.cityfacilitator.voucher.generate (POST /v3/alipay/commerce/cityfacilitator/voucher/generate).
// Official V3 contract: https://opendocs.alipay.com/apis/05wz9v
func (s *Service) CityFacilitatorVoucherGenerate(ctx context.Context, req CityFacilitatorVoucherGenerateRequest, opts ...core.CallOption) (*core.Response[CityFacilitatorVoucherGenerateResponse], error) {
	return core.ExecuteJSON[CityFacilitatorVoucherGenerateRequest, CityFacilitatorVoucherGenerateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.cityfacilitator.voucher.generate", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/cityfacilitator/voucher/generate",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CityFacilitatorVoucherGenerateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CityFacilitatorVoucherRefund calls alipay.commerce.cityfacilitator.voucher.refund (POST /v3/alipay/commerce/cityfacilitator/voucher/refund).
// Official V3 contract: https://opendocs.alipay.com/apis/05wujr
func (s *Service) CityFacilitatorVoucherRefund(ctx context.Context, req CityFacilitatorVoucherRefundRequest, opts ...core.CallOption) (*core.Response[CityFacilitatorVoucherRefundResponse], error) {
	return core.ExecuteJSON[CityFacilitatorVoucherRefundRequest, CityFacilitatorVoucherRefundResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.cityfacilitator.voucher.refund", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/cityfacilitator/voucher/refund",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CityFacilitatorVoucherRefundRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
