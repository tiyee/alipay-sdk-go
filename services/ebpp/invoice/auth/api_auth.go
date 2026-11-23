// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package auth

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

// InvoiceAuthSign calls alipay.ebpp.invoice.auth.sign (POST /v3/alipay/ebpp/invoice/auth/sign).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05jtxa
func (s *Service) InvoiceAuthSign(ctx context.Context, req InvoiceAuthSignRequest, opts ...core.CallOption) (*core.Response[InvoiceAuthSignResponse], error) {
	return core.ExecuteJSON[InvoiceAuthSignRequest, InvoiceAuthSignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.auth.sign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/auth/sign",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceAuthSignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceAuthUnsign calls alipay.ebpp.invoice.auth.unsign (POST /v3/alipay/ebpp/invoice/auth/unsign).
// Official V3 contract: https://opendocs.alipay.com/apis/05jwd5
func (s *Service) InvoiceAuthUnsign(ctx context.Context, req InvoiceAuthUnsignRequest, opts ...core.CallOption) (*core.Response[InvoiceAuthUnsignResponse], error) {
	return core.ExecuteJSON[InvoiceAuthUnsignRequest, InvoiceAuthUnsignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.auth.unsign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/auth/unsign",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceAuthUnsignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
