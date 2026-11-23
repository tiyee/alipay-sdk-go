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

// LogisticsInvoiceIstdWaybillCreate calls alipay.commerce.logistics.invoice.istdwaybill.create (POST /v3/alipay/commerce/logistics/invoice/istdwaybill/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/068d50
func (s *Service) LogisticsInvoiceIstdWaybillCreate(ctx context.Context, req LogisticsInvoiceIstdWaybillCreateRequest, opts ...core.CallOption) (*core.Response[LogisticsInvoiceIstdWaybillCreateResponse], error) {
	return core.ExecuteJSON[LogisticsInvoiceIstdWaybillCreateRequest, LogisticsInvoiceIstdWaybillCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.logistics.invoice.istdwaybill.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/logistics/invoice/istdwaybill/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[LogisticsInvoiceIstdWaybillCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// LogisticsInvoiceIstdWaybillQuery calls alipay.commerce.logistics.invoice.istdwaybill.query (GET /v3/alipay/commerce/logistics/invoice/istdwaybill/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/068lff
func (s *Service) LogisticsInvoiceIstdWaybillQuery(ctx context.Context, req LogisticsInvoiceIstdWaybillQueryRequest, opts ...core.CallOption) (*core.Response[LogisticsInvoiceIstdWaybillQueryResponse], error) {
	return core.ExecuteJSON[LogisticsInvoiceIstdWaybillQueryRequest, LogisticsInvoiceIstdWaybillQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.logistics.invoice.istdwaybill.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/logistics/invoice/istdwaybill/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[LogisticsInvoiceIstdWaybillQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
