package account

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

// AccountQuery calls alipay.fund.account.query (GET /v3/alipay/fund/account/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/77e2b925_alipay.fund.account.query
func (s *Service) AccountQuery(ctx context.Context, req FundAccountQueryRequest, opts ...core.CallOption) (*core.Response[FundAccountQueryResponse], error) {
	return core.ExecuteJSON[FundAccountQueryRequest, FundAccountQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.account.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/account/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAccountQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
