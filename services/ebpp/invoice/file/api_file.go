// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package file

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

// InvoiceFileOutputQuery calls alipay.ebpp.invoice.file.output.query (GET /v3/alipay/ebpp/invoice/file/output/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05l1fc
func (s *Service) InvoiceFileOutputQuery(ctx context.Context, req InvoiceFileOutputQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceFileOutputQueryResponse], error) {
	return core.ExecuteJSON[InvoiceFileOutputQueryRequest, InvoiceFileOutputQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.file.output.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/file/output/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceFileOutputQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
