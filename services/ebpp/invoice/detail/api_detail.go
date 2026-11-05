// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package detail

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

// InvoiceDetailOutputQuery calls alipay.ebpp.invoice.detail.output.query (POST /v3/alipay/ebpp/invoice/detail/output/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05jxvp
func (s *Service) InvoiceDetailOutputQuery(ctx context.Context, req InvoiceDetailOutputQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceDetailOutputQueryResponse], error) {
	return core.ExecuteJSON[InvoiceDetailOutputQueryRequest, InvoiceDetailOutputQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.detail.output.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/detail/output/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceDetailOutputQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
