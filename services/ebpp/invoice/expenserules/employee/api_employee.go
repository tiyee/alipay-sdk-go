package employee

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

// InvoiceExpenseRulesEmployeeQuery calls alipay.ebpp.invoice.expenserules.employee.query (GET /v3/alipay/ebpp/invoice/expenserules/employee/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05j4cj
func (s *Service) InvoiceExpenseRulesEmployeeQuery(ctx context.Context, req InvoiceExpenseRulesEmployeeQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesEmployeeQueryResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesEmployeeQueryRequest, InvoiceExpenseRulesEmployeeQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.employee.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/expenserules/employee/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesEmployeeQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseRulesEmployeeRulesQuery calls alipay.ebpp.invoice.expenserules.employeerules.query (GET /v3/alipay/ebpp/invoice/expenserules/employeerules/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05j4ci
func (s *Service) InvoiceExpenseRulesEmployeeRulesQuery(ctx context.Context, req InvoiceExpenseRulesEmployeeRulesQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesEmployeeRulesQueryResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesEmployeeRulesQueryRequest, InvoiceExpenseRulesEmployeeRulesQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.employeerules.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/expenserules/employeerules/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesEmployeeRulesQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseRulesGroupEmployeeModify calls alipay.ebpp.invoice.expenserules.groupemployee.modify (POST /v3/alipay/ebpp/invoice/expenserules/groupemployee/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05j82o
func (s *Service) InvoiceExpenseRulesGroupEmployeeModify(ctx context.Context, req InvoiceExpenseRulesGroupEmployeeModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesGroupEmployeeModifyResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesGroupEmployeeModifyRequest, InvoiceExpenseRulesGroupEmployeeModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.groupemployee.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expenserules/groupemployee/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesGroupEmployeeModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
