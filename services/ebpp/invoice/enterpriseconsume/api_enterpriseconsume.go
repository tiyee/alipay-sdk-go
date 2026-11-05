// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterpriseconsume

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

// InvoiceEnterpriseConsumeConsumeBatchQuery calls alipay.ebpp.invoice.enterpriseconsume.consume.batchquery (POST /v3/alipay/ebpp/invoice/enterpriseconsume/consume/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05jgvh
func (s *Service) InvoiceEnterpriseConsumeConsumeBatchQuery(ctx context.Context, req InvoiceEnterpriseConsumeConsumeBatchQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseConsumeConsumeBatchQueryResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseConsumeConsumeBatchQueryRequest, InvoiceEnterpriseConsumeConsumeBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseconsume.consume.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/enterpriseconsume/consume/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseConsumeConsumeBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceEnterpriseConsumeDetailQuery calls alipay.ebpp.invoice.enterpriseconsume.detail.query (GET /v3/alipay/ebpp/invoice/enterpriseconsume/detail/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05j4ce
func (s *Service) InvoiceEnterpriseConsumeDetailQuery(ctx context.Context, req InvoiceEnterpriseConsumeDetailQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseConsumeDetailQueryResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseConsumeDetailQueryRequest, InvoiceEnterpriseConsumeDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseconsume.detail.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/enterpriseconsume/detail/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseConsumeDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceEnterpriseConsumeEnterpriseOpenRuleCreate calls alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.create (POST /v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvsl
func (s *Service) InvoiceEnterpriseConsumeEnterpriseOpenRuleCreate(ctx context.Context, req InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateRequest, InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseConsumeEnterpriseOpenRuleCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceEnterpriseConsumeEnterpriseOpenRuleModify calls alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.modify (POST /v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wzal
func (s *Service) InvoiceEnterpriseConsumeEnterpriseOpenRuleModify(ctx context.Context, req InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyRequest, InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseConsumeEnterpriseOpenRuleModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceEnterpriseConsumeEnterpriseOpenRuleQuery calls alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.query (GET /v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wye3
func (s *Service) InvoiceEnterpriseConsumeEnterpriseOpenRuleQuery(ctx context.Context, req InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryRequest, InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseconsume.enterpriseopenrule.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/enterpriseconsume/enterpriseopenrule/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseConsumeEnterpriseOpenRuleQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceEnterpriseConsumeOpenRuleRelationCreate calls alipay.ebpp.invoice.enterpriseconsume.openrulerelation.create (POST /v3/alipay/ebpp/invoice/enterpriseconsume/openrulerelation/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05ucs6
func (s *Service) InvoiceEnterpriseConsumeOpenRuleRelationCreate(ctx context.Context, req InvoiceEnterpriseConsumeOpenRuleRelationCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseConsumeOpenRuleRelationCreateResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseConsumeOpenRuleRelationCreateRequest, InvoiceEnterpriseConsumeOpenRuleRelationCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseconsume.openrulerelation.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/enterpriseconsume/openrulerelation/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseConsumeOpenRuleRelationCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceEnterpriseConsumeRelateDetailQuery calls alipay.ebpp.invoice.enterpriseconsume.relatedetail.query (GET /v3/alipay/ebpp/invoice/enterpriseconsume/relatedetail/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05ip0t
func (s *Service) InvoiceEnterpriseConsumeRelateDetailQuery(ctx context.Context, req InvoiceEnterpriseConsumeRelateDetailQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseConsumeRelateDetailQueryResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseConsumeRelateDetailQueryRequest, InvoiceEnterpriseConsumeRelateDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseconsume.relatedetail.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/enterpriseconsume/relatedetail/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseConsumeRelateDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceEnterpriseConsumeSummaryInvoiceDetailQuery calls alipay.ebpp.invoice.enterpriseconsume.summaryinvoicedetail.query (GET /v3/alipay/ebpp/invoice/enterpriseconsume/summaryinvoicedetail/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05imom
func (s *Service) InvoiceEnterpriseConsumeSummaryInvoiceDetailQuery(ctx context.Context, req InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryRequest, InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterpriseconsume.summaryinvoicedetail.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/enterpriseconsume/summaryinvoicedetail/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseConsumeSummaryInvoiceDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
