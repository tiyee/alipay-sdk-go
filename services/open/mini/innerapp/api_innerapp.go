// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package innerapp

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

// MiniInnerAppCreate calls alipay.open.mini.innerapp.create (POST /v3/alipay/open/mini/innerapp/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05x9i7
func (s *Service) MiniInnerAppCreate(ctx context.Context, req MiniInnerAppCreateRequest, opts ...core.CallOption) (*core.Response[MiniInnerAppCreateResponse], error) {
	return core.ExecuteJSON[MiniInnerAppCreateRequest, MiniInnerAppCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerapp.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerapp/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerAppCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerAppPluginOrder calls alipay.open.mini.innerapp.plugin.order (POST /v3/alipay/open/mini/innerapp/plugin/order).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xng4
func (s *Service) MiniInnerAppPluginOrder(ctx context.Context, req MiniInnerAppPluginOrderRequest, opts ...core.CallOption) (*core.Response[MiniInnerAppPluginOrderResponse], error) {
	return core.ExecuteJSON[MiniInnerAppPluginOrderRequest, MiniInnerAppPluginOrderResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerapp.plugin.order", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerapp/plugin/order",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerAppPluginOrderRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerAppPluginRelationQuery calls alipay.open.mini.innerapp.pluginrelation.query (POST /v3/alipay/open/mini/innerapp/pluginrelation/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xhsh
func (s *Service) MiniInnerAppPluginRelationQuery(ctx context.Context, req MiniInnerAppPluginRelationQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerAppPluginRelationQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerAppPluginRelationQueryRequest, MiniInnerAppPluginRelationQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerapp.pluginrelation.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerapp/pluginrelation/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerAppPluginRelationQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerAppPluginServicePublish calls alipay.open.mini.innerapp.pluginservice.publish (POST /v3/alipay/open/mini/innerapp/pluginservice/publish).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xe0v
func (s *Service) MiniInnerAppPluginServicePublish(ctx context.Context, req MiniInnerAppPluginServicePublishRequest, opts ...core.CallOption) (*core.Response[MiniInnerAppPluginServicePublishResponse], error) {
	return core.ExecuteJSON[MiniInnerAppPluginServicePublishRequest, MiniInnerAppPluginServicePublishResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerapp.pluginservice.publish", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerapp/pluginservice/publish",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerAppPluginServicePublishRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerAppPluginSyncModeModify calls alipay.open.mini.innerapp.pluginsyncmode.modify (POST /v3/alipay/open/mini/innerapp/pluginsyncmode/modify).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xo5c
func (s *Service) MiniInnerAppPluginSyncModeModify(ctx context.Context, req MiniInnerAppPluginSyncModeModifyRequest, opts ...core.CallOption) (*core.Response[MiniInnerAppPluginSyncModeModifyResponse], error) {
	return core.ExecuteJSON[MiniInnerAppPluginSyncModeModifyRequest, MiniInnerAppPluginSyncModeModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerapp.pluginsyncmode.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerapp/pluginsyncmode/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerAppPluginSyncModeModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerAppPluginUseConfigCancel calls alipay.open.mini.innerapp.pluginuseconfig.cancel (POST /v3/alipay/open/mini/innerapp/pluginuseconfig/cancel).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xgfv
func (s *Service) MiniInnerAppPluginUseConfigCancel(ctx context.Context, req MiniInnerAppPluginUseConfigCancelRequest, opts ...core.CallOption) (*core.Response[MiniInnerAppPluginUseConfigCancelResponse], error) {
	return core.ExecuteJSON[MiniInnerAppPluginUseConfigCancelRequest, MiniInnerAppPluginUseConfigCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerapp.pluginuseconfig.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerapp/pluginuseconfig/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerAppPluginUseConfigCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerAppPluginUseConfigOnline calls alipay.open.mini.innerapp.pluginuseconfig.online (POST /v3/alipay/open/mini/innerapp/pluginuseconfig/online).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xacn
func (s *Service) MiniInnerAppPluginUseConfigOnline(ctx context.Context, req MiniInnerAppPluginUseConfigOnlineRequest, opts ...core.CallOption) (*core.Response[MiniInnerAppPluginUseConfigOnlineResponse], error) {
	return core.ExecuteJSON[MiniInnerAppPluginUseConfigOnlineRequest, MiniInnerAppPluginUseConfigOnlineResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerapp.pluginuseconfig.online", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerapp/pluginuseconfig/online",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerAppPluginUseConfigOnlineRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerAppPluginUseConfigUpgrade calls alipay.open.mini.innerapp.pluginuseconfig.upgrade (POST /v3/alipay/open/mini/innerapp/pluginuseconfig/upgrade).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xo8d
func (s *Service) MiniInnerAppPluginUseConfigUpgrade(ctx context.Context, req MiniInnerAppPluginUseConfigUpgradeRequest, opts ...core.CallOption) (*core.Response[MiniInnerAppPluginUseConfigUpgradeResponse], error) {
	return core.ExecuteJSON[MiniInnerAppPluginUseConfigUpgradeRequest, MiniInnerAppPluginUseConfigUpgradeResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerapp.pluginuseconfig.upgrade", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerapp/pluginuseconfig/upgrade",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerAppPluginUseConfigUpgradeRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerAppServicePublish calls alipay.open.mini.innerapp.service.publish (POST /v3/alipay/open/mini/innerapp/service/publish).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xndu
func (s *Service) MiniInnerAppServicePublish(ctx context.Context, req MiniInnerAppServicePublishRequest, opts ...core.CallOption) (*core.Response[MiniInnerAppServicePublishResponse], error) {
	return core.ExecuteJSON[MiniInnerAppServicePublishRequest, MiniInnerAppServicePublishResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerapp.service.publish", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerapp/service/publish",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerAppServicePublishRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerAppServiceQuery calls alipay.open.mini.innerapp.service.query (GET /v3/alipay/open/mini/innerapp/service/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xk3b
func (s *Service) MiniInnerAppServiceQuery(ctx context.Context, req MiniInnerAppServiceQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerAppServiceQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerAppServiceQueryRequest, MiniInnerAppServiceQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerapp.service.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/innerapp/service/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerAppServiceQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
