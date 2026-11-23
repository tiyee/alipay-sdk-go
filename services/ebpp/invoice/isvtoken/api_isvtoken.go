// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package isvtoken

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

// InvoiceISVTokenReimApply calls alipay.ebpp.invoice.isvtoken.reim.apply (POST /v3/alipay/ebpp/invoice/isvtoken/reim/apply).
// Official V3 contract: https://opendocs.alipay.com/apis/05xa7p
func (s *Service) InvoiceISVTokenReimApply(ctx context.Context, req InvoiceISVTokenReimApplyRequest, opts ...core.CallOption) (*core.Response[InvoiceISVTokenReimApplyResponse], error) {
	return core.ExecuteJSON[InvoiceISVTokenReimApplyRequest, InvoiceISVTokenReimApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.isvtoken.reim.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/isvtoken/reim/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceISVTokenReimApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
