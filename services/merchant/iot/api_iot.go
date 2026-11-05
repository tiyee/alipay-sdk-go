package iot

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

// IoTDeviceBind calls alipay.merchant.iot.device.bind (POST /v3/alipay/merchant/iot/device/bind).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c82d5f5f_alipay.merchant.iot.device.bind?scene=common&pathHash=0fb151a0
func (s *Service) IoTDeviceBind(ctx context.Context, req MerchantIoTDeviceBindRequest, opts ...core.CallOption) (*core.Response[MerchantIoTDeviceBindResponse], error) {
	return core.ExecuteJSON[MerchantIoTDeviceBindRequest, MerchantIoTDeviceBindResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.iot.device.bind", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/iot/device/bind",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantIoTDeviceBindRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// IoTDeviceQuery calls alipay.merchant.iot.device.query (GET /v3/alipay/merchant/iot/device/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/9b618feb_alipay.merchant.iot.device.query?scene=common&pathHash=9f128d89
func (s *Service) IoTDeviceQuery(ctx context.Context, req MerchantIoTDeviceQueryRequest, opts ...core.CallOption) (*core.Response[MerchantIoTDeviceQueryResponse], error) {
	return core.ExecuteJSON[MerchantIoTDeviceQueryRequest, MerchantIoTDeviceQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.iot.device.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/merchant/iot/device/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantIoTDeviceQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// IoTDeviceVerify calls alipay.merchant.iot.device.verify (POST /v3/alipay/merchant/iot/device/verify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4e0eced6_alipay.merchant.iot.device.verify?scene=common&pathHash=39ad307c
func (s *Service) IoTDeviceVerify(ctx context.Context, req MerchantIoTDeviceVerifyRequest, opts ...core.CallOption) (*core.Response[MerchantIoTDeviceVerifyResponse], error) {
	return core.ExecuteJSON[MerchantIoTDeviceVerifyRequest, MerchantIoTDeviceVerifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.iot.device.verify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/iot/device/verify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantIoTDeviceVerifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
