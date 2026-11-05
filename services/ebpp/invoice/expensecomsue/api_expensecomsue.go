// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expensecomsue

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

// InvoiceExpenseComsueOutsourceNotify calls alipay.ebpp.invoice.expensecomsue.outsource.notify (POST /v3/alipay/ebpp/invoice/expensecomsue/outsource/notify).
// Official V3 contract: https://opendocs.alipay.com/apis/05jgvn
func (s *Service) InvoiceExpenseComsueOutsourceNotify(ctx context.Context, req InvoiceExpenseComsueOutsourceNotifyRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseComsueOutsourceNotifyResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseComsueOutsourceNotifyRequest, InvoiceExpenseComsueOutsourceNotifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expensecomsue.outsource.notify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expensecomsue/outsource/notify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseComsueOutsourceNotifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
