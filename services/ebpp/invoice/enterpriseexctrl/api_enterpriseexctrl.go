// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterpriseexctrl

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

// InvoiceEnterpriseExCtrlEmployerTitleBatchQuery calls alipay.ebpp.invoice.enterpriseexctrl.employertitle.batchquery (POST /v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvun
func (s *Service) InvoiceEnterpriseExCtrlEmployerTitleBatchQuery(ctx context.Context, req InvoiceEnterpriseExCtrlEmployerTitleBatchQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseExCtrlEmployerTitleBatchQueryResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseExCtrlEmployerTitleBatchQueryRequest, InvoiceEnterpriseExCtrlEmployerTitleBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseexctrl.employertitle.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseExCtrlEmployerTitleBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceEnterpriseExCtrlEmployerTitleCreate calls alipay.ebpp.invoice.enterpriseexctrl.employertitle.create (POST /v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wydt
func (s *Service) InvoiceEnterpriseExCtrlEmployerTitleCreate(ctx context.Context, req InvoiceEnterpriseExCtrlEmployerTitleCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseExCtrlEmployerTitleCreateResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseExCtrlEmployerTitleCreateRequest, InvoiceEnterpriseExCtrlEmployerTitleCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseexctrl.employertitle.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseExCtrlEmployerTitleCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceEnterpriseExCtrlEmployerTitleModify calls alipay.ebpp.invoice.enterpriseexctrl.employertitle.modify (PUT /v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle).
// Official V3 contract: https://opendocs.alipay.com/apis/05lhsq
func (s *Service) InvoiceEnterpriseExCtrlEmployerTitleModify(ctx context.Context, req InvoiceEnterpriseExCtrlEmployerTitleModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseExCtrlEmployerTitleModifyResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseExCtrlEmployerTitleModifyRequest, InvoiceEnterpriseExCtrlEmployerTitleModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseexctrl.employertitle.modify", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseExCtrlEmployerTitleModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceEnterpriseExCtrlEmployerTitleQuery calls alipay.ebpp.invoice.enterpriseexctrl.employertitle.query (GET /v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05x1fx
func (s *Service) InvoiceEnterpriseExCtrlEmployerTitleQuery(ctx context.Context, req InvoiceEnterpriseExCtrlEmployerTitleQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseExCtrlEmployerTitleQueryResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseExCtrlEmployerTitleQueryRequest, InvoiceEnterpriseExCtrlEmployerTitleQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseexctrl.employertitle.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/enterpriseexctrl/employertitle/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseExCtrlEmployerTitleQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
