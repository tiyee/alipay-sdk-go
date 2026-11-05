package order

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

// OrderSync calls alipay.merchant.order.sync (POST /v3/alipay/merchant/order/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/be2462f6_alipay.merchant.order.sync
func (s *Service) OrderSync(ctx context.Context, req MerchantOrderSyncRequest, opts ...core.CallOption) (*core.Response[MerchantOrderSyncResponse], error) {
	return core.ExecuteJSON[MerchantOrderSyncRequest, MerchantOrderSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.order.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/order/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantOrderSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
