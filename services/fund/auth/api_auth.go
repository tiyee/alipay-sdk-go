package auth

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor core.Executor
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}

	return &Service{
		executor: executor,
	}, nil
}

// AuthOperationCancel calls alipay.fund.auth.operation.cancel (POST /v3/alipay/fund/auth/operation/cancel).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2fecf986_alipay.fund.auth.operation.cancel
func (s *Service) AuthOperationCancel(ctx context.Context, req FundAuthOperationCancelRequest, opts ...core.CallOption) (*core.Response[FundAuthOperationCancelResponse], error) {
	return core.ExecuteJSON[FundAuthOperationCancelRequest, FundAuthOperationCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.auth.operation.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/auth/operation/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAuthOperationCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AuthOperationDetailQuery calls alipay.fund.auth.operation.detail.query (POST /v3/alipay/fund/auth/operation/detail/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/984254a0_alipay.fund.auth.operation.detail.query
func (s *Service) AuthOperationDetailQuery(ctx context.Context, req FundAuthOperationDetailQueryRequest, opts ...core.CallOption) (*core.Response[FundAuthOperationDetailQueryResponse], error) {
	return core.ExecuteJSON[FundAuthOperationDetailQueryRequest, FundAuthOperationDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.auth.operation.detail.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/auth/operation/detail/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAuthOperationDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AuthOrderAppFreeze calls alipay.fund.auth.order.app.freeze (POST /v3/alipay/fund/auth/order/app/freeze).
// The POST route comes from the OpenAPI snapshot. Alipay documents App order string generation via sdkExecute: https://opendocs.alipay.com/open-v3/46bf87fe_alipay.fund.auth.order.app.freeze
func (s *Service) AuthOrderAppFreeze(ctx context.Context, req FundAuthOrderAppFreezeRequest, opts ...core.CallOption) (*core.Response[FundAuthOrderAppFreezeResponse], error) {
	return core.ExecuteJSON[FundAuthOrderAppFreezeRequest, FundAuthOrderAppFreezeResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.auth.order.app.freeze", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/auth/order/app/freeze",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAuthOrderAppFreezeRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AuthOrderFreeze calls alipay.fund.auth.order.freeze (POST /v3/alipay/fund/auth/order/freeze).
// Official V3 contract: https://opendocs.alipay.com/open-v3/1ffb104d_alipay.fund.auth.order.freeze
func (s *Service) AuthOrderFreeze(ctx context.Context, req FundAuthOrderFreezeRequest, opts ...core.CallOption) (*core.Response[FundAuthOrderFreezeResponse], error) {
	return core.ExecuteJSON[FundAuthOrderFreezeRequest, FundAuthOrderFreezeResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.auth.order.freeze", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/auth/order/freeze",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAuthOrderFreezeRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AuthOrderUnfreeze calls alipay.fund.auth.order.unfreeze (POST /v3/alipay/fund/auth/order/unfreeze).
// Official V3 contract: https://opendocs.alipay.com/open-v3/cb3b0ae2_alipay.fund.auth.order.unfreeze
func (s *Service) AuthOrderUnfreeze(ctx context.Context, req FundAuthOrderUnfreezeRequest, opts ...core.CallOption) (*core.Response[FundAuthOrderUnfreezeResponse], error) {
	return core.ExecuteJSON[FundAuthOrderUnfreezeRequest, FundAuthOrderUnfreezeResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.auth.order.unfreeze", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/auth/order/unfreeze",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAuthOrderUnfreezeRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AuthOrderVoucherCreate calls alipay.fund.auth.order.voucher.create (POST /v3/alipay/fund/auth/order/voucher/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/6f5ccbfe_alipay.fund.auth.order.voucher.create
func (s *Service) AuthOrderVoucherCreate(ctx context.Context, req FundAuthOrderVoucherCreateRequest, opts ...core.CallOption) (*core.Response[FundAuthOrderVoucherCreateResponse], error) {
	return core.ExecuteJSON[FundAuthOrderVoucherCreateRequest, FundAuthOrderVoucherCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.auth.order.voucher.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/auth/order/voucher/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAuthOrderVoucherCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
