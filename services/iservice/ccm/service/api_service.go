// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package service

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

// CcmServiceBuy calls alipay.iservice.ccm.service.buy (POST /v3/alipay/iservice/ccm/service/buy).
// Official V3 contract: https://opendocs.alipay.com/open-v3/705fa9ad_alipay.iservice.ccm.service.buy?scene=common&pathHash=42f013fd
func (s *Service) CcmServiceBuy(ctx context.Context, req CcmServiceBuyRequest, opts ...core.CallOption) (*core.Response[CcmServiceBuyResponse], error) {
	return core.ExecuteJSON[CcmServiceBuyRequest, CcmServiceBuyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.service.buy", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/service/buy",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmServiceBuyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmServiceClose calls alipay.iservice.ccm.service.close (POST /v3/alipay/iservice/ccm/service/close).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2202303e_alipay.iservice.ccm.service.close?scene=common&pathHash=4e8520fc
func (s *Service) CcmServiceClose(ctx context.Context, req CcmServiceCloseRequest, opts ...core.CallOption) (*core.Response[CcmServiceCloseResponse], error) {
	return core.ExecuteJSON[CcmServiceCloseRequest, CcmServiceCloseResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.service.close", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/service/close",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmServiceCloseRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmServiceInitialize calls alipay.iservice.ccm.service.initialize (POST /v3/alipay/iservice/ccm/service/initialize).
// Official V3 contract: https://opendocs.alipay.com/open-v3/ab388310_alipay.iservice.ccm.service.initialize?scene=common&pathHash=3dce471e
func (s *Service) CcmServiceInitialize(ctx context.Context, req CcmServiceInitializeRequest, opts ...core.CallOption) (*core.Response[CcmServiceInitializeResponse], error) {
	return core.ExecuteJSON[CcmServiceInitializeRequest, CcmServiceInitializeResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.service.initialize", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/service/initialize",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmServiceInitializeRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
