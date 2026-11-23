// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package token

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

// InvoiceTokenBatchQuery calls alipay.ebpp.invoice.token.batchquery (GET /v3/alipay/ebpp/invoice/token/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05khyy
func (s *Service) InvoiceTokenBatchQuery(ctx context.Context, req InvoiceTokenBatchQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceTokenBatchQueryResponse], error) {
	return core.ExecuteJSON[InvoiceTokenBatchQueryRequest, InvoiceTokenBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.token.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/token/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceTokenBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
