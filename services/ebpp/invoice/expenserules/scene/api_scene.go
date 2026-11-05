package scene

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

// InvoiceExpenseRulesSceneRuleCreate calls alipay.ebpp.invoice.expenserules.scenerule.create (POST /v3/alipay/ebpp/invoice/expenserules/scenerule/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05j82p
func (s *Service) InvoiceExpenseRulesSceneRuleCreate(ctx context.Context, req InvoiceExpenseRulesSceneRuleCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesSceneRuleCreateResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesSceneRuleCreateRequest, InvoiceExpenseRulesSceneRuleCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.scenerule.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expenserules/scenerule/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesSceneRuleCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseRulesSceneRuleModify calls alipay.ebpp.invoice.expenserules.scenerule.modify (POST /v3/alipay/ebpp/invoice/expenserules/scenerule/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05xk1w
func (s *Service) InvoiceExpenseRulesSceneRuleModify(ctx context.Context, req InvoiceExpenseRulesSceneRuleModifyRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesSceneRuleModifyResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesSceneRuleModifyRequest, InvoiceExpenseRulesSceneRuleModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.scenerule.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/expenserules/scenerule/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesSceneRuleModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceExpenseRulesSceneRuleQuery calls alipay.ebpp.invoice.expenserules.scenerule.query (GET /v3/alipay/ebpp/invoice/expenserules/scenerule/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05x1fs
func (s *Service) InvoiceExpenseRulesSceneRuleQuery(ctx context.Context, req InvoiceExpenseRulesSceneRuleQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceExpenseRulesSceneRuleQueryResponse], error) {
	return core.ExecuteJSON[InvoiceExpenseRulesSceneRuleQueryRequest, InvoiceExpenseRulesSceneRuleQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.expenserules.scenerule.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/expenserules/scenerule/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceExpenseRulesSceneRuleQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
