package usertrade

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

// UserTradeBatchQuery calls alipay.fund.usertrade.batchquery (POST /v3/alipay/fund/usertrade/batchquery).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xggy
func (s *Service) UserTradeBatchQuery(ctx context.Context, req FundUserTradeBatchQueryRequest, opts ...core.CallOption) (*core.Response[FundUserTradeBatchQueryResponse], error) {
	return core.ExecuteJSON[FundUserTradeBatchQueryRequest, FundUserTradeBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.usertrade.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/usertrade/batchquery",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundUserTradeBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
