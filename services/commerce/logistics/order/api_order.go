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

// LogisticsOrderInstantDeliveryCancel calls alipay.commerce.logistics.order.instantdelivery.cancel (POST /v3/alipay/commerce/logistics/order/instantdelivery/cancel).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/068eft
func (s *Service) LogisticsOrderInstantDeliveryCancel(ctx context.Context, req LogisticsOrderInstantDeliveryCancelRequest, opts ...core.CallOption) (*core.Response[LogisticsOrderInstantDeliveryCancelResponse], error) {
	return core.ExecuteJSON[LogisticsOrderInstantDeliveryCancelRequest, LogisticsOrderInstantDeliveryCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.logistics.order.instantdelivery.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/logistics/order/instantdelivery/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[LogisticsOrderInstantDeliveryCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// LogisticsOrderInstantDeliveryCreate calls alipay.commerce.logistics.order.instantdelivery.create (POST /v3/alipay/commerce/logistics/order/instantdelivery/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05lo9w
func (s *Service) LogisticsOrderInstantDeliveryCreate(ctx context.Context, req LogisticsOrderInstantDeliveryCreateRequest, opts ...core.CallOption) (*core.Response[LogisticsOrderInstantDeliveryCreateResponse], error) {
	return core.ExecuteJSON[LogisticsOrderInstantDeliveryCreateRequest, LogisticsOrderInstantDeliveryCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.logistics.order.instantdelivery.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/logistics/order/instantdelivery/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[LogisticsOrderInstantDeliveryCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// LogisticsOrderInstantDeliveryPreCreate calls alipay.commerce.logistics.order.instantdelivery.precreate (POST /v3/alipay/commerce/logistics/order/instantdelivery/precreate).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05l560
func (s *Service) LogisticsOrderInstantDeliveryPreCreate(ctx context.Context, req LogisticsOrderInstantDeliveryPreCreateRequest, opts ...core.CallOption) (*core.Response[LogisticsOrderInstantDeliveryPreCreateResponse], error) {
	return core.ExecuteJSON[LogisticsOrderInstantDeliveryPreCreateRequest, LogisticsOrderInstantDeliveryPreCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.logistics.order.instantdelivery.precreate", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/logistics/order/instantdelivery/precreate",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[LogisticsOrderInstantDeliveryPreCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// LogisticsOrderIstdCancelPreConsult calls alipay.commerce.logistics.order.istdcancel.preconsult (POST /v3/alipay/commerce/logistics/order/istdcancel/preconsult).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/068d52
func (s *Service) LogisticsOrderIstdCancelPreConsult(ctx context.Context, req LogisticsOrderIstdCancelPreConsultRequest, opts ...core.CallOption) (*core.Response[LogisticsOrderIstdCancelPreConsultResponse], error) {
	return core.ExecuteJSON[LogisticsOrderIstdCancelPreConsultRequest, LogisticsOrderIstdCancelPreConsultResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.logistics.order.istdcancel.preconsult", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/logistics/order/istdcancel/preconsult",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[LogisticsOrderIstdCancelPreConsultRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// LogisticsOrderIstdRetryCreate calls alipay.commerce.logistics.order.istdretry.create (POST /v3/alipay/commerce/logistics/order/istdretry/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05l561
func (s *Service) LogisticsOrderIstdRetryCreate(ctx context.Context, req LogisticsOrderIstdRetryCreateRequest, opts ...core.CallOption) (*core.Response[LogisticsOrderIstdRetryCreateResponse], error) {
	return core.ExecuteJSON[LogisticsOrderIstdRetryCreateRequest, LogisticsOrderIstdRetryCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.logistics.order.istdretry.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/logistics/order/istdretry/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[LogisticsOrderIstdRetryCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
