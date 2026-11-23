// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package sycn

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

// InvoiceSycn calls alipay.ebpp.invoice.sycn (POST /v3/alipay/ebpp/invoice/sycn).
// Official V3 contract: https://opendocs.alipay.com/apis/05la3z
func (s *Service) InvoiceSycn(ctx context.Context, req InvoiceSycnRequest, opts ...core.CallOption) (*core.Response[InvoiceSycnResponse], error) {
	return core.ExecuteJSON[InvoiceSycnRequest, InvoiceSycnResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.sycn", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/sycn",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceSycnRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
