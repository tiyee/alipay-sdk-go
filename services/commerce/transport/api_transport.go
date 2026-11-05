// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package transport

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service executes one alipay V3 service group.
type Service struct{ executor core.Executor }

// NewService constructs the service with a shared V3 executor.
// NewService constructs the service with a shared V3 executor.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

// TransportChargerChargerBindInfoSync calls alipay.commerce.transport.charger.chargerbindinfo.sync (POST /v3/alipay/commerce/transport/charger/chargerbindinfo/sync).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xqw6
func (s *Service) TransportChargerChargerBindInfoSync(ctx context.Context, req TransportChargerChargerBindInfoSyncRequest, opts ...core.CallOption) (*core.Response[TransportChargerChargerBindInfoSyncResponse], error) {
	return core.ExecuteJSON[TransportChargerChargerBindInfoSyncRequest, TransportChargerChargerBindInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.transport.charger.chargerbindinfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/transport/charger/chargerbindinfo/sync",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[TransportChargerChargerBindInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// TransportParkingEnterInfoSync calls alipay.commerce.transport.parking.enterinfo.sync (POST /v3/alipay/commerce/transport/parking/enterinfo/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e38b4d57_alipay.commerce.transport.parking.enterinfo.sync
func (s *Service) TransportParkingEnterInfoSync(ctx context.Context, req TransportParkingEnterInfoSyncRequest, opts ...core.CallOption) (*core.Response[TransportParkingEnterInfoSyncResponse], error) {
	return core.ExecuteJSON[TransportParkingEnterInfoSyncRequest, TransportParkingEnterInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.transport.parking.enterinfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/transport/parking/enterinfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[TransportParkingEnterInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// TransportParkingExitInfoSync calls alipay.commerce.transport.parking.exitinfo.sync (POST /v3/alipay/commerce/transport/parking/exitinfo/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/9986eded_alipay.commerce.transport.parking.exitinfo.sync
func (s *Service) TransportParkingExitInfoSync(ctx context.Context, req TransportParkingExitInfoSyncRequest, opts ...core.CallOption) (*core.Response[TransportParkingExitInfoSyncResponse], error) {
	return core.ExecuteJSON[TransportParkingExitInfoSyncRequest, TransportParkingExitInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.transport.parking.exitinfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/transport/parking/exitinfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[TransportParkingExitInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// TransportParkingPaymentInfoSync calls alipay.commerce.transport.parking.paymentinfo.sync (POST /v3/alipay/commerce/transport/parking/paymentinfo/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/12f22c8f_alipay.commerce.transport.parking.paymentinfo.sync
func (s *Service) TransportParkingPaymentInfoSync(ctx context.Context, req TransportParkingPaymentInfoSyncRequest, opts ...core.CallOption) (*core.Response[TransportParkingPaymentInfoSyncResponse], error) {
	return core.ExecuteJSON[TransportParkingPaymentInfoSyncRequest, TransportParkingPaymentInfoSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.transport.parking.paymentinfo.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/transport/parking/paymentinfo/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[TransportParkingPaymentInfoSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
