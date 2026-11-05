package project

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor core.Executor
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}

	return &Service{
		executor: executor,
	}, nil
}

// InvoiceExpenseRulesProjectEmployeeModify calls alipay.ebpp.invoice.expenserules.projectemployee.modify (POST /v3/alipay/ebpp/invoice/expenserules/projectemployee/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvx5
func (s *Service) InvoiceExpenseRulesProjectEmployeeModify(ctx context.Context, req InvoiceExpenseRulesProjectEmployeeModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesProjectEmployeeModifyResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesProjectEmployeeModifyRequest, InvoiceExpenseRulesProjectEmployeeModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.projectemployee.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expenserules/projectemployee/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesProjectEmployeeModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseRulesProjectInfoModify calls alipay.ebpp.invoice.expenserules.projectinfo.modify (POST /v3/alipay/ebpp/invoice/expenserules/projectinfo/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuq6
func (s *Service) InvoiceExpenseRulesProjectInfoModify(ctx context.Context, req InvoiceExpenseRulesProjectInfoModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesProjectInfoModifyResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesProjectInfoModifyRequest, InvoiceExpenseRulesProjectInfoModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.projectinfo.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expenserules/projectinfo/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesProjectInfoModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseRulesProjectRuleCreate calls alipay.ebpp.invoice.expenserules.projectrule.create (POST /v3/alipay/ebpp/invoice/expenserules/projectrule/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05isio
func (s *Service) InvoiceExpenseRulesProjectRuleCreate(ctx context.Context, req InvoiceExpenseRulesProjectRuleCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesProjectRuleCreateResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesProjectRuleCreateRequest, InvoiceExpenseRulesProjectRuleCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.projectrule.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expenserules/projectrule/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesProjectRuleCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseRulesProjectRuleQuery calls alipay.ebpp.invoice.expenserules.projectrule.query (GET /v3/alipay/ebpp/invoice/expenserules/projectrule/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05ip0x
func (s *Service) InvoiceExpenseRulesProjectRuleQuery(ctx context.Context, req InvoiceExpenseRulesProjectRuleQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesProjectRuleQueryResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesProjectRuleQueryRequest, InvoiceExpenseRulesProjectRuleQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.projectrule.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/expenserules/projectrule/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesProjectRuleQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseRulesProjectRulesModify calls alipay.ebpp.invoice.expenserules.projectrules.modify (POST /v3/alipay/ebpp/invoice/expenserules/projectrules/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05xnct
func (s *Service) InvoiceExpenseRulesProjectRulesModify(ctx context.Context, req InvoiceExpenseRulesProjectRulesModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesProjectRulesModifyResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesProjectRulesModifyRequest, InvoiceExpenseRulesProjectRulesModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.projectrules.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expenserules/projectrules/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesProjectRulesModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
