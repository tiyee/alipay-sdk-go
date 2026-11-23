// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package issuerule

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

// InvoiceIssueRuleCreate calls alipay.ebpp.invoice.issuerule.create (POST /v3/alipay/ebpp/invoice/issuerule/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05xoz2
func (s *Service) InvoiceIssueRuleCreate(ctx context.Context, req InvoiceIssueRuleCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceIssueRuleCreateResponse], error) {
	return core.ExecuteJSON[InvoiceIssueRuleCreateRequest, InvoiceIssueRuleCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.issuerule.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/issuerule/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceIssueRuleCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceIssueRuleDelete calls alipay.ebpp.invoice.issuerule.delete (POST /v3/alipay/ebpp/invoice/issuerule/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05wol0
func (s *Service) InvoiceIssueRuleDelete(ctx context.Context, req InvoiceIssueRuleDeleteRequest, opts ...core.CallOption) (*core.Response[InvoiceIssueRuleDeleteResponse], error) {
	return core.ExecuteJSON[InvoiceIssueRuleDeleteRequest, InvoiceIssueRuleDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.issuerule.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/issuerule/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceIssueRuleDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceIssueRuleModify calls alipay.ebpp.invoice.issuerule.modify (POST /v3/alipay/ebpp/invoice/issuerule/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wyj4
func (s *Service) InvoiceIssueRuleModify(ctx context.Context, req InvoiceIssueRuleModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceIssueRuleModifyResponse], error) {
	return core.ExecuteJSON[InvoiceIssueRuleModifyRequest, InvoiceIssueRuleModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.issuerule.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/issuerule/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceIssueRuleModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
