// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ecorder

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

// InvoiceEcOrderOrderQuery calls alipay.ebpp.invoice.ecorder.order.query (GET /v3/alipay/ebpp/invoice/ecorder/order/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05imol
func (s *Service) InvoiceEcOrderOrderQuery(ctx context.Context, req InvoiceEcOrderOrderQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceEcOrderOrderQueryResponse], error) {
	return core.ExecuteJSON[InvoiceEcOrderOrderQueryRequest, InvoiceEcOrderOrderQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.ecorder.order.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/ecorder/order/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEcOrderOrderQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
