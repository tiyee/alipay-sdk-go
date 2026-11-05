// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package expensecontrol

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

// InvoiceExpenseControlAggregationCreate calls alipay.ebpp.invoice.expensecontrol.aggregation.create (POST /v3/alipay/ebpp/invoice/expensecontrol/aggregation/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05xsf0
func (s *Service) InvoiceExpenseControlAggregationCreate(ctx context.Context, req InvoiceExpenseControlAggregationCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseControlAggregationCreateResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseControlAggregationCreateRequest, InvoiceExpenseControlAggregationCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expensecontrol.aggregation.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expensecontrol/aggregation/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseControlAggregationCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseControlIssueBatchCancel calls alipay.ebpp.invoice.expensecontrol.issuebatch.cancel (POST /v3/alipay/ebpp/invoice/expensecontrol/issuebatch/cancel).
// Official V3 contract: https://opendocs.alipay.com/apis/05wumt
func (s *Service) InvoiceExpenseControlIssueBatchCancel(ctx context.Context, req InvoiceExpenseControlIssueBatchCancelRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseControlIssueBatchCancelResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseControlIssueBatchCancelRequest, InvoiceExpenseControlIssueBatchCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expensecontrol.issuebatch.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expensecontrol/issuebatch/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseControlIssueBatchCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseControlIssueBatchCreate calls alipay.ebpp.invoice.expensecontrol.issuebatch.create (POST /v3/alipay/ebpp/invoice/expensecontrol/issuebatch/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05j9jf
func (s *Service) InvoiceExpenseControlIssueBatchCreate(ctx context.Context, req InvoiceExpenseControlIssueBatchCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseControlIssueBatchCreateResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseControlIssueBatchCreateRequest, InvoiceExpenseControlIssueBatchCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expensecontrol.issuebatch.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expensecontrol/issuebatch/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseControlIssueBatchCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseControlQuotaCreate calls alipay.ebpp.invoice.expensecontrol.quota.create (POST /v3/alipay/ebpp/invoice/expensecontrol/quota/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05iz32
func (s *Service) InvoiceExpenseControlQuotaCreate(ctx context.Context, req InvoiceExpenseControlQuotaCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseControlQuotaCreateResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseControlQuotaCreateRequest, InvoiceExpenseControlQuotaCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expensecontrol.quota.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expensecontrol/quota/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseControlQuotaCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseControlQuotaModify calls alipay.ebpp.invoice.expensecontrol.quota.modify (POST /v3/alipay/ebpp/invoice/expensecontrol/quota/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuoa
func (s *Service) InvoiceExpenseControlQuotaModify(ctx context.Context, req InvoiceExpenseControlQuotaModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseControlQuotaModifyResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseControlQuotaModifyRequest, InvoiceExpenseControlQuotaModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expensecontrol.quota.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expensecontrol/quota/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseControlQuotaModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseControlQuotaQuery calls alipay.ebpp.invoice.expensecontrol.quota.query (POST /v3/alipay/ebpp/invoice/expensecontrol/quota/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05imop
func (s *Service) InvoiceExpenseControlQuotaQuery(ctx context.Context, req InvoiceExpenseControlQuotaQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseControlQuotaQueryResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseControlQuotaQueryRequest, InvoiceExpenseControlQuotaQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expensecontrol.quota.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expensecontrol/quota/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseControlQuotaQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
