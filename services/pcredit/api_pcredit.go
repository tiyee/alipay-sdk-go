// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package pcredit

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

// HuabeiAuthAccumulationQuery calls alipay.pcredit.huabei.auth.accumulation.query (GET /v3/alipay/pcredit/huabei/auth/accumulation/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05ge5u
func (s *Service) HuabeiAuthAccumulationQuery(ctx context.Context, req PcreditHuabeiAuthAccumulationQueryRequest, opts ...core.CallOption) (*core.Response[PcreditHuabeiAuthAccumulationQueryResponse], error) {
	return core.ExecuteJSON[PcreditHuabeiAuthAccumulationQueryRequest, PcreditHuabeiAuthAccumulationQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.pcredit.huabei.auth.accumulation.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/pcredit/huabei/auth/accumulation/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PcreditHuabeiAuthAccumulationQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// HuabeiAuthAgreementClose calls alipay.pcredit.huabei.auth.agreement.close (PUT /v3/alipay/pcredit/huabei/auth/agreement/close).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05gb9o
func (s *Service) HuabeiAuthAgreementClose(ctx context.Context, req PcreditHuabeiAuthAgreementCloseRequest, opts ...core.CallOption) (*core.Response[PcreditHuabeiAuthAgreementCloseResponse], error) {
	return core.ExecuteJSON[PcreditHuabeiAuthAgreementCloseRequest, PcreditHuabeiAuthAgreementCloseResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.pcredit.huabei.auth.agreement.close", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/pcredit/huabei/auth/agreement/close",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PcreditHuabeiAuthAgreementCloseRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// HuabeiAuthAgreementQuery calls alipay.pcredit.huabei.auth.agreement.query (GET /v3/alipay/pcredit/huabei/auth/agreement/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05fign
func (s *Service) HuabeiAuthAgreementQuery(ctx context.Context, req PcreditHuabeiAuthAgreementQueryRequest, opts ...core.CallOption) (*core.Response[PcreditHuabeiAuthAgreementQueryResponse], error) {
	return core.ExecuteJSON[PcreditHuabeiAuthAgreementQueryRequest, PcreditHuabeiAuthAgreementQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.pcredit.huabei.auth.agreement.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/pcredit/huabei/auth/agreement/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PcreditHuabeiAuthAgreementQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// HuabeiAuthOrderQuery calls alipay.pcredit.huabei.auth.order.query (GET /v3/alipay/pcredit/huabei/auth/order/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05h84d
func (s *Service) HuabeiAuthOrderQuery(ctx context.Context, req PcreditHuabeiAuthOrderQueryRequest, opts ...core.CallOption) (*core.Response[PcreditHuabeiAuthOrderQueryResponse], error) {
	return core.ExecuteJSON[PcreditHuabeiAuthOrderQueryRequest, PcreditHuabeiAuthOrderQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.pcredit.huabei.auth.order.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/pcredit/huabei/auth/order/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PcreditHuabeiAuthOrderQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// HuabeiAuthOrderUnfreeze calls alipay.pcredit.huabei.auth.order.unfreeze (PUT /v3/alipay/pcredit/huabei/auth/order/unfreeze).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05gcoh
func (s *Service) HuabeiAuthOrderUnfreeze(ctx context.Context, req PcreditHuabeiAuthOrderUnfreezeRequest, opts ...core.CallOption) (*core.Response[PcreditHuabeiAuthOrderUnfreezeResponse], error) {
	return core.ExecuteJSON[PcreditHuabeiAuthOrderUnfreezeRequest, PcreditHuabeiAuthOrderUnfreezeResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.pcredit.huabei.auth.order.unfreeze", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/pcredit/huabei/auth/order/unfreeze",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PcreditHuabeiAuthOrderUnfreezeRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// HuabeiAuthSettleApply calls alipay.pcredit.huabei.auth.settle.apply (PUT /v3/alipay/pcredit/huabei/auth/settle/apply).
// Official V3 contract: https://opendocs.alipay.com/apis/05ghz8
func (s *Service) HuabeiAuthSettleApply(ctx context.Context, req PcreditHuabeiAuthSettleApplyRequest, opts ...core.CallOption) (*core.Response[PcreditHuabeiAuthSettleApplyResponse], error) {
	return core.ExecuteJSON[PcreditHuabeiAuthSettleApplyRequest, PcreditHuabeiAuthSettleApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.pcredit.huabei.auth.settle.apply", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/pcredit/huabei/auth/settle/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PcreditHuabeiAuthSettleApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
