// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package list

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

// InvoiceListExpenseSync calls alipay.ebpp.invoice.list.expense.sync (POST /v3/alipay/ebpp/invoice/list/expense/sync).
// Official V3 contract: https://opendocs.alipay.com/apis/05jve4
func (s *Service) InvoiceListExpenseSync(ctx context.Context, req InvoiceListExpenseSyncRequest, opts ...core.CallOption) (*core.Response[InvoiceListExpenseSyncResponse], error) {
	return core.ExecuteJSON[InvoiceListExpenseSyncRequest, InvoiceListExpenseSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.list.expense.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/list/expense/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceListExpenseSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
