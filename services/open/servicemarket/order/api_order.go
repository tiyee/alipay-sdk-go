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

// ServiceMarketOrderAccept calls alipay.open.servicemarket.order.accept (POST /v3/alipay/open/servicemarket/order/accept).
// Official V3 contract: https://opendocs.alipay.com/open-v3/added906_alipay.open.servicemarket.order.accept
func (s *Service) ServiceMarketOrderAccept(ctx context.Context, req ServiceMarketOrderAcceptRequest, opts ...core.CallOption) (*core.Response[ServiceMarketOrderAcceptResponse], error) {
	return core.ExecuteJSON[ServiceMarketOrderAcceptRequest, ServiceMarketOrderAcceptResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.servicemarket.order.accept", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/servicemarket/order/accept",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ServiceMarketOrderAcceptRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ServiceMarketOrderCreate calls alipay.open.servicemarket.order.create (POST /v3/alipay/open/servicemarket/order/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xo6s
func (s *Service) ServiceMarketOrderCreate(ctx context.Context, req ServiceMarketOrderCreateRequest, opts ...core.CallOption) (*core.Response[ServiceMarketOrderCreateResponse], error) {
	return core.ExecuteJSON[ServiceMarketOrderCreateRequest, ServiceMarketOrderCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.servicemarket.order.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/servicemarket/order/create",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ServiceMarketOrderCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ServiceMarketOrderItemCancel calls alipay.open.servicemarket.order.item.cancel (POST /v3/alipay/open/servicemarket/order/item/cancel).
// Official V3 contract: https://opendocs.alipay.com/open-v3/97757082_alipay.open.servicemarket.order.item.cancel
func (s *Service) ServiceMarketOrderItemCancel(ctx context.Context, req ServiceMarketOrderItemCancelRequest, opts ...core.CallOption) (*core.Response[ServiceMarketOrderItemCancelResponse], error) {
	return core.ExecuteJSON[ServiceMarketOrderItemCancelRequest, ServiceMarketOrderItemCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.servicemarket.order.item.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/servicemarket/order/item/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ServiceMarketOrderItemCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ServiceMarketOrderItemComplete calls alipay.open.servicemarket.order.item.complete (POST /v3/alipay/open/servicemarket/order/item/complete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e31fe80d_alipay.open.servicemarket.order.item.complete
func (s *Service) ServiceMarketOrderItemComplete(ctx context.Context, req ServiceMarketOrderItemCompleteRequest, opts ...core.CallOption) (*core.Response[ServiceMarketOrderItemCompleteResponse], error) {
	return core.ExecuteJSON[ServiceMarketOrderItemCompleteRequest, ServiceMarketOrderItemCompleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.servicemarket.order.item.complete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/servicemarket/order/item/complete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ServiceMarketOrderItemCompleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ServiceMarketOrderItemConfirm calls alipay.open.servicemarket.order.item.confirm (POST /v3/alipay/open/servicemarket/order/item/confirm).
// Official V3 contract: https://opendocs.alipay.com/open-v3/13c55349_alipay.open.servicemarket.order.item.confirm
func (s *Service) ServiceMarketOrderItemConfirm(ctx context.Context, req ServiceMarketOrderItemConfirmRequest, opts ...core.CallOption) (*core.Response[ServiceMarketOrderItemConfirmResponse], error) {
	return core.ExecuteJSON[ServiceMarketOrderItemConfirmRequest, ServiceMarketOrderItemConfirmResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.servicemarket.order.item.confirm", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/servicemarket/order/item/confirm",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ServiceMarketOrderItemConfirmRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ServiceMarketOrderNotify calls alipay.open.servicemarket.order.notify (POST /v3/alipay/open/servicemarket/order/notify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2233a7ba_alipay.open.servicemarket.order.notify
func (s *Service) ServiceMarketOrderNotify(ctx context.Context, req ServiceMarketOrderNotifyRequest, opts ...core.CallOption) (*core.Response[ServiceMarketOrderNotifyResponse], error) {
	return core.ExecuteJSON[ServiceMarketOrderNotifyRequest, ServiceMarketOrderNotifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.servicemarket.order.notify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/servicemarket/order/notify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ServiceMarketOrderNotifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ServiceMarketOrderQuery calls alipay.open.servicemarket.order.query (GET /v3/alipay/open/servicemarket/order/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a59abd06_alipay.open.servicemarket.order.query
func (s *Service) ServiceMarketOrderQuery(ctx context.Context, req ServiceMarketOrderQueryRequest, opts ...core.CallOption) (*core.Response[ServiceMarketOrderQueryResponse], error) {
	return core.ExecuteJSON[ServiceMarketOrderQueryRequest, ServiceMarketOrderQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.servicemarket.order.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/servicemarket/order/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ServiceMarketOrderQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ServiceMarketOrderReject calls alipay.open.servicemarket.order.reject (POST /v3/alipay/open/servicemarket/order/reject).
// Official V3 contract: https://opendocs.alipay.com/open-v3/f5d452ab_alipay.open.servicemarket.order.reject
func (s *Service) ServiceMarketOrderReject(ctx context.Context, req ServiceMarketOrderRejectRequest, opts ...core.CallOption) (*core.Response[ServiceMarketOrderRejectResponse], error) {
	return core.ExecuteJSON[ServiceMarketOrderRejectRequest, ServiceMarketOrderRejectResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.servicemarket.order.reject", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/servicemarket/order/reject",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ServiceMarketOrderRejectRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
