// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchantlist

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

// InvoiceMerchantListEnterApply calls alipay.ebpp.invoice.merchantlist.enter.apply (POST /v3/alipay/ebpp/invoice/merchantlist/enter/apply).
// Official V3 contract: https://opendocs.alipay.com/apis/05vxtl
func (s *Service) InvoiceMerchantListEnterApply(ctx context.Context, req InvoiceMerchantListEnterApplyRequest, opts ...core.CallOption) (*core.Response[InvoiceMerchantListEnterApplyResponse], error) {
	return core.ExecuteJSON[InvoiceMerchantListEnterApplyRequest, InvoiceMerchantListEnterApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.merchantlist.enter.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/merchantlist/enter/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceMerchantListEnterApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
