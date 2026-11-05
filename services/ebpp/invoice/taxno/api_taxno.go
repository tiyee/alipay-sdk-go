// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package taxno

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

// InvoiceTaxNoBatchQuery calls alipay.ebpp.invoice.taxno.batchquery (POST /v3/alipay/ebpp/invoice/taxno/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05jrk9
func (s *Service) InvoiceTaxNoBatchQuery(ctx context.Context, req InvoiceTaxNoBatchQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceTaxNoBatchQueryResponse], error) {
	return core.ExecuteJSON[InvoiceTaxNoBatchQueryRequest, InvoiceTaxNoBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.taxno.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/taxno/batchquery",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceTaxNoBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
