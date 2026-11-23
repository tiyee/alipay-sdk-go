package order

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装支付宝 V3 order 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor         core.Executor
	orderPay         core.Operation[OrderPayRequest, OrderPayResponse]
	orderSettle      core.Operation[OrderSettleRequest, OrderSettleResponse]
	orderSettleQuery core.Operation[OrderSettleQueryRequest, OrderSettleQueryResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	orderPay, err := core.NewJSONOperation[OrderPayRequest, OrderPayResponse](core.OperationSpec{
		Method: "alipay.trade.order.pay", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/order/pay",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	orderSettle, err := core.NewJSONOperation[OrderSettleRequest, OrderSettleResponse](core.OperationSpec{
		Method: "alipay.trade.order.settle", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/order/settle",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	orderSettleQuery, err := core.NewJSONOperation[OrderSettleQueryRequest, OrderSettleQueryResponse](core.OperationSpec{
		Method: "alipay.trade.order.settle.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/trade/order/settle/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:         executor,
		orderPay:         orderPay,
		orderSettle:      orderSettle,
		orderSettleQuery: orderSettleQuery,
	}, nil
}

// OrderPay 调用 alipay.trade.order.pay 对已存在订单发起支付
// （V3 POST /v3/alipay/trade/order/pay）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/c2751e3b_alipay.trade.order.pay
// OutRequestNo 用于幂等控制；异步支付以异步通知或查询确认结果。
func (s *Service) OrderPay(ctx context.Context, req OrderPayRequest,
	opts ...core.CallOption) (*core.Response[OrderPayResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.orderPay,
		core.Request[OrderPayRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// OrderSettle 调用 alipay.trade.order.settle 发起交易分账/结算
// （V3 POST /v3/alipay/trade/order/settle）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/34402227_alipay.trade.order.settle
// OutRequestNo 用于幂等；分账结果以 SettleNo 配合查询接口确认。
func (s *Service) OrderSettle(ctx context.Context, req OrderSettleRequest,
	opts ...core.CallOption) (*core.Response[OrderSettleResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.orderSettle,
		core.Request[OrderSettleRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// OrderSettleQuery 调用 alipay.trade.order.settle.query 查询分账结果
// （V3 GET /v3/alipay/trade/order/settle/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/3af46ab7_alipay.trade.order.settle.query
// settle_no 或（out_request_no + trade_no）二选一。
func (s *Service) OrderSettleQuery(ctx context.Context, req OrderSettleQueryRequest,
	opts ...core.CallOption) (*core.Response[OrderSettleQueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.orderSettleQuery,
		core.Request[OrderSettleQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
