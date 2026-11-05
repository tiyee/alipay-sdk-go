// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expensescene

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

// InvoiceExpenseSceneMerchantQuery calls alipay.ebpp.invoice.expensescene.merchant.query (GET /v3/alipay/ebpp/invoice/expensescene/merchant/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05pe0a
func (s *Service) InvoiceExpenseSceneMerchantQuery(ctx context.Context, req InvoiceExpenseSceneMerchantQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseSceneMerchantQueryResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseSceneMerchantQueryRequest, InvoiceExpenseSceneMerchantQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expensescene.merchant.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/expensescene/merchant/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseSceneMerchantQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
