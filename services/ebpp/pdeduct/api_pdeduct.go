package pdeduct

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

// PdeductAsyncPay calls alipay.ebpp.pdeduct.async.pay (POST /v3/alipay/ebpp/pdeduct/async/pay).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xmda
func (s *Service) PdeductAsyncPay(ctx context.Context, req PdeductAsyncPayRequest, opts ...core.CallOption) (*core.Response[PdeductAsyncPayResponse], error) {
	return core.ExecuteJSON[PdeductAsyncPayRequest, PdeductAsyncPayResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.pdeduct.async.pay", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/pdeduct/async/pay",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PdeductAsyncPayRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PdeductBillPayStatus calls alipay.ebpp.pdeduct.bill.pay.status (GET /v3/alipay/ebpp/pdeduct/bill/pay/status).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c7416677_alipay.ebpp.pdeduct.bill.pay.status
func (s *Service) PdeductBillPayStatus(ctx context.Context, req PdeductBillPayStatusRequest, opts ...core.CallOption) (*core.Response[PdeductBillPayStatusResponse], error) {
	return core.ExecuteJSON[PdeductBillPayStatusRequest, PdeductBillPayStatusResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.pdeduct.bill.pay.status", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/pdeduct/bill/pay/status",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PdeductBillPayStatusRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PdeductPay calls alipay.ebpp.pdeduct.pay (POST /v3/alipay/ebpp/pdeduct/pay).
// Official V3 contract: https://opendocs.alipay.com/open-v3/fea30a1b_alipay.ebpp.pdeduct.pay
func (s *Service) PdeductPay(ctx context.Context, req PdeductPayRequest, opts ...core.CallOption) (*core.Response[PdeductPayResponse], error) {
	return core.ExecuteJSON[PdeductPayRequest, PdeductPayResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.pdeduct.pay", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/pdeduct/pay",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PdeductPayRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
