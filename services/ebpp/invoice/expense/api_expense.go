// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expense

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

// InvoiceExpenseProgressSync calls alipay.ebpp.invoice.expense.progress.sync (POST /v3/alipay/ebpp/invoice/expense/progress/sync).
// Official V3 contract: https://opendocs.alipay.com/apis/05wx2n
func (s *Service) InvoiceExpenseProgressSync(ctx context.Context, req InvoiceExpenseProgressSyncRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseProgressSyncResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseProgressSyncRequest, InvoiceExpenseProgressSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expense.progress.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expense/progress/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseProgressSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
