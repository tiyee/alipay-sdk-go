// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package operation

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

// OperationActivityMerchantModify calls alipay.commerce.operation.activity.merchant.modify (POST /v3/alipay/commerce/operation/activity/merchant/modify).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/063rs2
func (s *Service) OperationActivityMerchantModify(ctx context.Context, req OperationActivityMerchantModifyRequest, opts ...core.CallOption) (*core.Response[OperationActivityMerchantModifyResponse], error) {
	return core.ExecuteJSON[OperationActivityMerchantModifyRequest, OperationActivityMerchantModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.operation.activity.merchant.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/operation/activity/merchant/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[OperationActivityMerchantModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// OperationActivityMerchantSign calls alipay.commerce.operation.activity.merchant.sign (POST /v3/alipay/commerce/operation/activity/merchant/sign).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/06347n
func (s *Service) OperationActivityMerchantSign(ctx context.Context, req OperationActivityMerchantSignRequest, opts ...core.CallOption) (*core.Response[OperationActivityMerchantSignResponse], error) {
	return core.ExecuteJSON[OperationActivityMerchantSignRequest, OperationActivityMerchantSignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.operation.activity.merchant.sign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/operation/activity/merchant/sign",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[OperationActivityMerchantSignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// OperationActivityMerchantUnsign calls alipay.commerce.operation.activity.merchant.unsign (POST /v3/alipay/commerce/operation/activity/merchant/unsign).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/06336j
func (s *Service) OperationActivityMerchantUnsign(ctx context.Context, req OperationActivityMerchantUnsignRequest, opts ...core.CallOption) (*core.Response[OperationActivityMerchantUnsignResponse], error) {
	return core.ExecuteJSON[OperationActivityMerchantUnsignRequest, OperationActivityMerchantUnsignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.operation.activity.merchant.unsign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/operation/activity/merchant/unsign",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[OperationActivityMerchantUnsignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// OperationBankActivityQuery calls alipay.commerce.operation.bank.activity.query (GET /v3/alipay/commerce/operation/bank/activity/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/063ecd
func (s *Service) OperationBankActivityQuery(ctx context.Context, req OperationBankActivityQueryRequest, opts ...core.CallOption) (*core.Response[OperationBankActivityQueryResponse], error) {
	return core.ExecuteJSON[OperationBankActivityQueryRequest, OperationBankActivityQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.operation.bank.activity.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/operation/bank/activity/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[OperationBankActivityQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
