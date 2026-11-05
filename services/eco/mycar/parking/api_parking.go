package parking

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

// MyCarParkingChargeInfoSync calls alipay.eco.mycar.parking.chargeinfo.sync (POST /v3/alipay/eco/mycar/parking/chargeinfo/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/18d2ca24_alipay.eco.mycar.parking.chargeinfo.sync
func (s *Service) MyCarParkingChargeInfoSync(ctx context.Context, req MyCarParkingChargeInfoSyncRequest, opts ...core.CallOption) (*core.Response[MyCarParkingChargeInfoSyncResponse], error) {
	return core.ExecuteJSON[MyCarParkingChargeInfoSyncRequest, MyCarParkingChargeInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.chargeinfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/chargeinfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingChargeInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingChargingInfoSync calls alipay.eco.mycar.parking.charginginfo.sync (POST /v3/alipay/eco/mycar/parking/charginginfo/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/91cababc_alipay.eco.mycar.parking.charginginfo.sync
func (s *Service) MyCarParkingChargingInfoSync(ctx context.Context, req MyCarParkingChargingInfoSyncRequest, opts ...core.CallOption) (*core.Response[MyCarParkingChargingInfoSyncResponse], error) {
	return core.ExecuteJSON[MyCarParkingChargingInfoSyncRequest, MyCarParkingChargingInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.charginginfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/charginginfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingChargingInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingConfigQuery calls alipay.eco.mycar.parking.config.query (GET /v3/alipay/eco/mycar/parking/config/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/735dfdcb_alipay.eco.mycar.parking.config.query
func (s *Service) MyCarParkingConfigQuery(ctx context.Context, req MyCarParkingConfigQueryRequest, opts ...core.CallOption) (*core.Response[MyCarParkingConfigQueryResponse], error) {
	return core.ExecuteJSON[MyCarParkingConfigQueryRequest, MyCarParkingConfigQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.config.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/eco/mycar/parking/config/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingConfigQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingConfigSet calls alipay.eco.mycar.parking.config.set (POST /v3/alipay/eco/mycar/parking/config/set).
// Official V3 contract: https://opendocs.alipay.com/open-v3/17f00a51_alipay.eco.mycar.parking.config.set
func (s *Service) MyCarParkingConfigSet(ctx context.Context, req MyCarParkingConfigSetRequest, opts ...core.CallOption) (*core.Response[MyCarParkingConfigSetResponse], error) {
	return core.ExecuteJSON[MyCarParkingConfigSetRequest, MyCarParkingConfigSetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.config.set", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/config/set",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingConfigSetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingEnterInfoSync calls alipay.eco.mycar.parking.enterinfo.sync (POST /v3/alipay/eco/mycar/parking/enterinfo/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/eb761e07_alipay.eco.mycar.parking.enterinfo.sync
func (s *Service) MyCarParkingEnterInfoSync(ctx context.Context, req MyCarParkingEnterInfoSyncRequest, opts ...core.CallOption) (*core.Response[MyCarParkingEnterInfoSyncResponse], error) {
	return core.ExecuteJSON[MyCarParkingEnterInfoSyncRequest, MyCarParkingEnterInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.enterinfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/enterinfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingEnterInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingExitInfoSync calls alipay.eco.mycar.parking.exitinfo.sync (POST /v3/alipay/eco/mycar/parking/exitinfo/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4b64c632_alipay.eco.mycar.parking.exitinfo.sync
func (s *Service) MyCarParkingExitInfoSync(ctx context.Context, req MyCarParkingExitInfoSyncRequest, opts ...core.CallOption) (*core.Response[MyCarParkingExitInfoSyncResponse], error) {
	return core.ExecuteJSON[MyCarParkingExitInfoSyncRequest, MyCarParkingExitInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.exitinfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/exitinfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingExitInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingOrderSync calls alipay.eco.mycar.parking.order.sync (POST /v3/alipay/eco/mycar/parking/order/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e22974fc_alipay.eco.mycar.parking.order.sync
func (s *Service) MyCarParkingOrderSync(ctx context.Context, req MyCarParkingOrderSyncRequest, opts ...core.CallOption) (*core.Response[MyCarParkingOrderSyncResponse], error) {
	return core.ExecuteJSON[MyCarParkingOrderSyncRequest, MyCarParkingOrderSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.order.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/order/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingOrderSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingOrderUpdate calls alipay.eco.mycar.parking.order.update (POST /v3/alipay/eco/mycar/parking/order/update).
// Official V3 contract: https://opendocs.alipay.com/open-v3/fe91ddf6_alipay.eco.mycar.parking.order.update
func (s *Service) MyCarParkingOrderUpdate(ctx context.Context, req MyCarParkingOrderUpdateRequest, opts ...core.CallOption) (*core.Response[MyCarParkingOrderUpdateResponse], error) {
	return core.ExecuteJSON[MyCarParkingOrderUpdateRequest, MyCarParkingOrderUpdateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.order.update", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/order/update",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingOrderUpdateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingOvertimeChargingInfoSync calls alipay.eco.mycar.parking.overtimecharginginfo.sync (POST /v3/alipay/eco/mycar/parking/overtimecharginginfo/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/3e99daa8_alipay.eco.mycar.parking.overtimecharginginfo.sync
func (s *Service) MyCarParkingOvertimeChargingInfoSync(ctx context.Context, req MyCarParkingOvertimeChargingInfoSyncRequest, opts ...core.CallOption) (*core.Response[MyCarParkingOvertimeChargingInfoSyncResponse], error) {
	return core.ExecuteJSON[MyCarParkingOvertimeChargingInfoSyncRequest, MyCarParkingOvertimeChargingInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.overtimecharginginfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/overtimecharginginfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingOvertimeChargingInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingParkingLotBizTransfer calls alipay.eco.mycar.parking.parkinglotbiz.transfer (POST /v3/alipay/eco/mycar/parking/parkinglotbiz/transfer).
// Official V3 contract: https://opendocs.alipay.com/open-v3/cb924074_alipay.eco.mycar.parking.parkinglotbiz.transfer
func (s *Service) MyCarParkingParkingLotBizTransfer(ctx context.Context, req MyCarParkingParkingLotBizTransferRequest, opts ...core.CallOption) (*core.Response[MyCarParkingParkingLotBizTransferResponse], error) {
	return core.ExecuteJSON[MyCarParkingParkingLotBizTransferRequest, MyCarParkingParkingLotBizTransferResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.parkinglotbiz.transfer", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/parkinglotbiz/transfer",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingParkingLotBizTransferRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingParkingLotInfoCreate calls alipay.eco.mycar.parking.parkinglotinfo.create (POST /v3/alipay/eco/mycar/parking/parkinglotinfo/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/fc05d317_alipay.eco.mycar.parking.parkinglotinfo.create
func (s *Service) MyCarParkingParkingLotInfoCreate(ctx context.Context, req MyCarParkingParkingLotInfoCreateRequest, opts ...core.CallOption) (*core.Response[MyCarParkingParkingLotInfoCreateResponse], error) {
	return core.ExecuteJSON[MyCarParkingParkingLotInfoCreateRequest, MyCarParkingParkingLotInfoCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.parkinglotinfo.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/parkinglotinfo/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingParkingLotInfoCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingParkingLotInfoQuery calls alipay.eco.mycar.parking.parkinglotinfo.query (GET /v3/alipay/eco/mycar/parking/parkinglotinfo/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4772d6e6_alipay.eco.mycar.parking.parkinglotinfo.query
func (s *Service) MyCarParkingParkingLotInfoQuery(ctx context.Context, req MyCarParkingParkingLotInfoQueryRequest, opts ...core.CallOption) (*core.Response[MyCarParkingParkingLotInfoQueryResponse], error) {
	return core.ExecuteJSON[MyCarParkingParkingLotInfoQueryRequest, MyCarParkingParkingLotInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.parkinglotinfo.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/eco/mycar/parking/parkinglotinfo/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingParkingLotInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingParkingLotInfoUpdate calls alipay.eco.mycar.parking.parkinglotinfo.update (POST /v3/alipay/eco/mycar/parking/parkinglotinfo/update).
// Official V3 contract: https://opendocs.alipay.com/open-v3/223553d2_alipay.eco.mycar.parking.parkinglotinfo.update
func (s *Service) MyCarParkingParkingLotInfoUpdate(ctx context.Context, req MyCarParkingParkingLotInfoUpdateRequest, opts ...core.CallOption) (*core.Response[MyCarParkingParkingLotInfoUpdateResponse], error) {
	return core.ExecuteJSON[MyCarParkingParkingLotInfoUpdateRequest, MyCarParkingParkingLotInfoUpdateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.parkinglotinfo.update", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/parkinglotinfo/update",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingParkingLotInfoUpdateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingPaymentInfoSync calls alipay.eco.mycar.parking.paymentinfo.sync (POST /v3/alipay/eco/mycar/parking/paymentinfo/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/def30621_alipay.eco.mycar.parking.paymentinfo.sync
func (s *Service) MyCarParkingPaymentInfoSync(ctx context.Context, req MyCarParkingPaymentInfoSyncRequest, opts ...core.CallOption) (*core.Response[MyCarParkingPaymentInfoSyncResponse], error) {
	return core.ExecuteJSON[MyCarParkingPaymentInfoSyncRequest, MyCarParkingPaymentInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.paymentinfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/paymentinfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingPaymentInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingSpaceInfoSync calls alipay.eco.mycar.parking.spaceinfo.sync (POST /v3/alipay/eco/mycar/parking/spaceinfo/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/cab695b7_alipay.eco.mycar.parking.spaceinfo.sync
func (s *Service) MyCarParkingSpaceInfoSync(ctx context.Context, req MyCarParkingSpaceInfoSyncRequest, opts ...core.CallOption) (*core.Response[MyCarParkingSpaceInfoSyncResponse], error) {
	return core.ExecuteJSON[MyCarParkingSpaceInfoSyncRequest, MyCarParkingSpaceInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.spaceinfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/mycar/parking/spaceinfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingSpaceInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MyCarParkingVehicleQuery calls alipay.eco.mycar.parking.vehicle.query (GET /v3/alipay/eco/mycar/parking/vehicle/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c04327bf_alipay.eco.mycar.parking.vehicle.query
func (s *Service) MyCarParkingVehicleQuery(ctx context.Context, req MyCarParkingVehicleQueryRequest, opts ...core.CallOption) (*core.Response[MyCarParkingVehicleQueryResponse], error) {
	return core.ExecuteJSON[MyCarParkingVehicleQueryRequest, MyCarParkingVehicleQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.mycar.parking.vehicle.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/eco/mycar/parking/vehicle/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MyCarParkingVehicleQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
