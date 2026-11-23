package orderinfo

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor      core.Executor
	orderinfoSync core.Operation[OrderInfoSyncRequest, OrderInfoSyncResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	orderinfoSync, err := core.NewJSONOperation[OrderInfoSyncRequest, OrderInfoSyncResponse](core.OperationSpec{
		Method: "alipay.trade.orderinfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/orderinfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:      executor,
		orderinfoSync: orderinfoSync,
	}, nil
}

// OrderInfoSync 调用 alipay.trade.orderinfo.sync 同步订单信息
// （V3 POST /v3/alipay/trade/orderinfo/sync）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/f8b82729_alipay.trade.orderinfo.sync
// 相同 OutRequestNo 重复请求不会更新订单信息（网关幂等）。
func (s *Service) OrderInfoSync(ctx context.Context, req OrderInfoSyncRequest,
	opts ...core.CallOption) (*core.Response[OrderInfoSyncResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.orderinfoSync,
		core.Request[OrderInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
