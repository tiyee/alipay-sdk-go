package balance

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

// EcBalanceDownloadURLQuery calls alipay.commerce.ec.balance.downloadurl.query (GET /v3/alipay/commerce/ec/balance/downloadurl/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05j1bx
func (s *Service) EcBalanceDownloadURLQuery(ctx context.Context, req EcBalanceDownloadURLQueryRequest, opts ...core.CallOption) (*core.Response[EcBalanceDownloadURLQueryResponse], error) {
	return core.ExecuteJSON[EcBalanceDownloadURLQueryRequest, EcBalanceDownloadURLQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.balance.downloadurl.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/balance/downloadurl/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcBalanceDownloadURLQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcBalancePeriodModify calls alipay.commerce.ec.balance.period.modify (POST /v3/alipay/commerce/ec/balance/period/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05j82h
func (s *Service) EcBalancePeriodModify(ctx context.Context, req EcBalancePeriodModifyRequest, opts ...core.CallOption) (*core.Response[EcBalancePeriodModifyResponse], error) {
	return core.ExecuteJSON[EcBalancePeriodModifyRequest, EcBalancePeriodModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.balance.period.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/ec/balance/period/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcBalancePeriodModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
