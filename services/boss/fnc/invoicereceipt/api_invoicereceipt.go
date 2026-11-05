package invoicereceipt

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

// FncInvoiceReceiptBatchQuery calls alipay.boss.fnc.invoicereceipt.batchquery (POST /v3/alipay/boss/fnc/invoicereceipt/batchquery).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xiqt
func (s *Service) FncInvoiceReceiptBatchQuery(ctx context.Context, req BossFncInvoiceReceiptBatchQueryRequest, opts ...core.CallOption) (*core.Response[BossFncInvoiceReceiptBatchQueryResponse], error) {
	return core.ExecuteJSON[BossFncInvoiceReceiptBatchQueryRequest, BossFncInvoiceReceiptBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.boss.fnc.invoicereceipt.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/boss/fnc/invoicereceipt/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BossFncInvoiceReceiptBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// FncInvoiceReceiptCreate calls alipay.boss.fnc.invoicereceipt.create (POST /v3/alipay/boss/fnc/invoicereceipt/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xnh4
func (s *Service) FncInvoiceReceiptCreate(ctx context.Context, req BossFncInvoiceReceiptCreateRequest, opts ...core.CallOption) (*core.Response[BossFncInvoiceReceiptCreateResponse], error) {
	return core.ExecuteJSON[BossFncInvoiceReceiptCreateRequest, BossFncInvoiceReceiptCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.boss.fnc.invoicereceipt.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/boss/fnc/invoicereceipt/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BossFncInvoiceReceiptCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
