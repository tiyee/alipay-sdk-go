// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package einvpackage

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

// InvoiceEInvPackageQuery calls alipay.ebpp.invoice.einvpackage.query (GET /v3/alipay/ebpp/invoice/einvpackage/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvv7
func (s *Service) InvoiceEInvPackageQuery(ctx context.Context, req InvoiceEInvPackageQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceEInvPackageQueryResponse], error) {
	return core.ExecuteJSON[InvoiceEInvPackageQueryRequest, InvoiceEInvPackageQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.einvpackage.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/einvpackage/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEInvPackageQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
