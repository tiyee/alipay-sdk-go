// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package order

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

// InvoiceOrderQuery calls alipay.ebpp.invoice.order.query (GET /v3/alipay/ebpp/invoice/order/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05v5uk
func (s *Service) InvoiceOrderQuery(ctx context.Context, req InvoiceOrderQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceOrderQueryResponse], error) {
	return core.ExecuteJSON[InvoiceOrderQueryRequest, InvoiceOrderQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.order.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/order/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceOrderQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
