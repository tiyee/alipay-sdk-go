package invoice

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

// FncInvoiceApply calls alipay.boss.fnc.invoice.apply (POST /v3/alipay/boss/fnc/invoice/apply).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xo5f
func (s *Service) FncInvoiceApply(ctx context.Context, req BossFncInvoiceApplyRequest, opts ...core.CallOption) (*core.Response[BossFncInvoiceApplyResponse], error) {
	return core.ExecuteJSON[BossFncInvoiceApplyRequest, BossFncInvoiceApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.boss.fnc.invoice.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/boss/fnc/invoice/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BossFncInvoiceApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// FncInvoiceBatchQuery calls alipay.boss.fnc.invoice.batchquery (POST /v3/alipay/boss/fnc/invoice/batchquery).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xdyx
func (s *Service) FncInvoiceBatchQuery(ctx context.Context, req BossFncInvoiceBatchQueryRequest, opts ...core.CallOption) (*core.Response[BossFncInvoiceBatchQueryResponse], error) {
	return core.ExecuteJSON[BossFncInvoiceBatchQueryRequest, BossFncInvoiceBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.boss.fnc.invoice.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/boss/fnc/invoice/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BossFncInvoiceBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// FncInvoiceQuery calls alipay.boss.fnc.invoice.query (GET /v3/alipay/boss/fnc/invoice/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xhqn
func (s *Service) FncInvoiceQuery(ctx context.Context, req BossFncInvoiceQueryRequest, opts ...core.CallOption) (*core.Response[BossFncInvoiceQueryResponse], error) {
	return core.ExecuteJSON[BossFncInvoiceQueryRequest, BossFncInvoiceQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.boss.fnc.invoice.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/boss/fnc/invoice/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BossFncInvoiceQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
