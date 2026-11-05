// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package institution

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

// InvoiceInstitutionCreate calls alipay.ebpp.invoice.institution.create (POST /v3/alipay/ebpp/invoice/institution/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvvq
func (s *Service) InvoiceInstitutionCreate(ctx context.Context, req InvoiceInstitutionCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceInstitutionCreateResponse], error) {
	return core.ExecuteJSON[InvoiceInstitutionCreateRequest, InvoiceInstitutionCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.institution.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/institution/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceInstitutionCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceInstitutionDelete calls alipay.ebpp.invoice.institution.delete (DELETE /v3/alipay/ebpp/invoice/institution/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05wunr
func (s *Service) InvoiceInstitutionDelete(ctx context.Context, req InvoiceInstitutionDeleteRequest, opts ...core.CallOption) (*core.Response[InvoiceInstitutionDeleteResponse], error) {
	return core.ExecuteJSON[InvoiceInstitutionDeleteRequest, InvoiceInstitutionDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.institution.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/ebpp/invoice/institution/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceInstitutionDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceInstitutionDetailInfoQuery calls alipay.ebpp.invoice.institution.detailinfo.query (GET /v3/alipay/ebpp/invoice/institution/detailinfo/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05j9jg
func (s *Service) InvoiceInstitutionDetailInfoQuery(ctx context.Context, req InvoiceInstitutionDetailInfoQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceInstitutionDetailInfoQueryResponse], error) {
	return core.ExecuteJSON[InvoiceInstitutionDetailInfoQueryRequest, InvoiceInstitutionDetailInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.institution.detailinfo.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/institution/detailinfo/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceInstitutionDetailInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceInstitutionExpenseRuleCreate calls alipay.ebpp.invoice.institution.expenserule.create (POST /v3/alipay/ebpp/invoice/institution/expenserule/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05waoz
func (s *Service) InvoiceInstitutionExpenseRuleCreate(ctx context.Context, req InvoiceInstitutionExpenseRuleCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceInstitutionExpenseRuleCreateResponse], error) {
	return core.ExecuteJSON[InvoiceInstitutionExpenseRuleCreateRequest, InvoiceInstitutionExpenseRuleCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.institution.expenserule.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/institution/expenserule/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceInstitutionExpenseRuleCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceInstitutionExpenseRuleDelete calls alipay.ebpp.invoice.institution.expenserule.delete (POST /v3/alipay/ebpp/invoice/institution/expenserule/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05xa73
func (s *Service) InvoiceInstitutionExpenseRuleDelete(ctx context.Context, req InvoiceInstitutionExpenseRuleDeleteRequest, opts ...core.CallOption) (*core.Response[InvoiceInstitutionExpenseRuleDeleteResponse], error) {
	return core.ExecuteJSON[InvoiceInstitutionExpenseRuleDeleteRequest, InvoiceInstitutionExpenseRuleDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.institution.expenserule.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/institution/expenserule/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceInstitutionExpenseRuleDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceInstitutionExpenseRuleModify calls alipay.ebpp.invoice.institution.expenserule.modify (POST /v3/alipay/ebpp/invoice/institution/expenserule/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuq0
func (s *Service) InvoiceInstitutionExpenseRuleModify(ctx context.Context, req InvoiceInstitutionExpenseRuleModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceInstitutionExpenseRuleModifyResponse], error) {
	return core.ExecuteJSON[InvoiceInstitutionExpenseRuleModifyRequest, InvoiceInstitutionExpenseRuleModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.institution.expenserule.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/institution/expenserule/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceInstitutionExpenseRuleModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceInstitutionModify calls alipay.ebpp.invoice.institution.modify (POST /v3/alipay/ebpp/invoice/institution/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvqu
func (s *Service) InvoiceInstitutionModify(ctx context.Context, req InvoiceInstitutionModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceInstitutionModifyResponse], error) {
	return core.ExecuteJSON[InvoiceInstitutionModifyRequest, InvoiceInstitutionModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.institution.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/institution/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceInstitutionModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceInstitutionPageInfoQuery calls alipay.ebpp.invoice.institution.pageinfo.query (GET /v3/alipay/ebpp/invoice/institution/pageinfo/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05xo4v
func (s *Service) InvoiceInstitutionPageInfoQuery(ctx context.Context, req InvoiceInstitutionPageInfoQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceInstitutionPageInfoQueryResponse], error) {
	return core.ExecuteJSON[InvoiceInstitutionPageInfoQueryRequest, InvoiceInstitutionPageInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.institution.pageinfo.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/institution/pageinfo/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceInstitutionPageInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceInstitutionScopeModify calls alipay.ebpp.invoice.institution.scope.modify (POST /v3/alipay/ebpp/invoice/institution/scope/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05ip0w
func (s *Service) InvoiceInstitutionScopeModify(ctx context.Context, req InvoiceInstitutionScopeModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceInstitutionScopeModifyResponse], error) {
	return core.ExecuteJSON[InvoiceInstitutionScopeModifyRequest, InvoiceInstitutionScopeModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.institution.scope.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/institution/scope/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceInstitutionScopeModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceInstitutionScopePageInfoQuery calls alipay.ebpp.invoice.institution.scopepageinfo.query (GET /v3/alipay/ebpp/invoice/institution/scopepageinfo/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05j1c2
func (s *Service) InvoiceInstitutionScopePageInfoQuery(ctx context.Context, req InvoiceInstitutionScopePageInfoQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceInstitutionScopePageInfoQueryResponse], error) {
	return core.ExecuteJSON[InvoiceInstitutionScopePageInfoQueryRequest, InvoiceInstitutionScopePageInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.institution.scopepageinfo.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/institution/scopepageinfo/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceInstitutionScopePageInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
