package alipaypoint

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

// AlipayPointBudgetLibQuery calls alipay.user.alipaypoint.budgetlib.query (GET /v3/alipay/user/alipaypoint/budgetlib/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wmz4
func (s *Service) AlipayPointBudgetLibQuery(ctx context.Context, req UserAlipayPointBudgetLibQueryRequest, opts ...core.CallOption) (*core.Response[UserAlipayPointBudgetLibQueryResponse], error) {
	return core.ExecuteJSON[UserAlipayPointBudgetLibQueryRequest, UserAlipayPointBudgetLibQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.alipaypoint.budgetlib.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/user/alipaypoint/budgetlib/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserAlipayPointBudgetLibQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AlipayPointSend calls alipay.user.alipaypoint.send (POST /v3/alipay/user/alipaypoint/send).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvr1
func (s *Service) AlipayPointSend(ctx context.Context, req UserAlipayPointSendRequest, opts ...core.CallOption) (*core.Response[UserAlipayPointSendResponse], error) {
	return core.ExecuteJSON[UserAlipayPointSendRequest, UserAlipayPointSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.alipaypoint.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/alipaypoint/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserAlipayPointSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
