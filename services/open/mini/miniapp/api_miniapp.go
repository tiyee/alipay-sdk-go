// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package miniapp

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

// MiniMiniAppBrandCancel calls alipay.open.mini.miniapp.brand.cancel (POST /v3/alipay/open/mini/miniapp/brand/cancel).
// Official V3 contract: https://opendocs.alipay.com/open-v3/ac02eafe_alipay.open.mini.miniapp.brand.cancel
func (s *Service) MiniMiniAppBrandCancel(ctx context.Context, req MiniMiniAppBrandCancelRequest, opts ...core.CallOption) (*core.Response[MiniMiniAppBrandCancelResponse], error) {
	return core.ExecuteJSON[MiniMiniAppBrandCancelRequest, MiniMiniAppBrandCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.miniapp.brand.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/miniapp/brand/cancel",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniMiniAppBrandCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniMiniAppBrandQuery calls alipay.open.mini.miniapp.brand.query (GET /v3/alipay/open/mini/miniapp/brand/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/6187351c_alipay.open.mini.miniapp.brand.query
func (s *Service) MiniMiniAppBrandQuery(ctx context.Context, req MiniMiniAppBrandQueryRequest, opts ...core.CallOption) (*core.Response[MiniMiniAppBrandQueryResponse], error) {
	return core.ExecuteJSON[MiniMiniAppBrandQueryRequest, MiniMiniAppBrandQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.miniapp.brand.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/miniapp/brand/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniMiniAppBrandQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniMiniAppBrandSubmit calls alipay.open.mini.miniapp.brand.submit (POST /v3/alipay/open/mini/miniapp/brand/submit).
// Official V3 contract: https://opendocs.alipay.com/open-v3/24509d98_alipay.open.mini.miniapp.brand.submit
func (s *Service) MiniMiniAppBrandSubmit(ctx context.Context, req MiniMiniAppBrandSubmitRequest, opts ...core.CallOption) (*core.Response[MiniMiniAppBrandSubmitResponse], error) {
	return core.ExecuteJSON[MiniMiniAppBrandSubmitRequest, MiniMiniAppBrandSubmitResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.miniapp.brand.submit", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/miniapp/brand/submit",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniMiniAppBrandSubmitRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniMiniAppBrandUpload calls alipay.open.mini.miniapp.brand.upload (POST /v3/alipay/open/mini/miniapp/brand/upload).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e77e0097_alipay.open.mini.miniapp.brand.upload
func (s *Service) MiniMiniAppBrandUpload(ctx context.Context, req MiniMiniAppBrandUploadRequest, opts ...core.CallOption) (*core.Response[MiniMiniAppBrandUploadResponse], error) {
	return core.ExecuteJSON[MiniMiniAppBrandUploadRequest, MiniMiniAppBrandUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.miniapp.brand.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/miniapp/brand/upload",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniMiniAppBrandUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniMiniAppServiceConfigModify calls alipay.open.mini.miniapp.serviceconfig.modify (POST /v3/alipay/open/mini/miniapp/serviceconfig/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/95e9abe6_alipay.open.mini.miniapp.serviceconfig.modify
func (s *Service) MiniMiniAppServiceConfigModify(ctx context.Context, req MiniMiniAppServiceConfigModifyRequest, opts ...core.CallOption) (*core.Response[MiniMiniAppServiceConfigModifyResponse], error) {
	return core.ExecuteJSON[MiniMiniAppServiceConfigModifyRequest, MiniMiniAppServiceConfigModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.miniapp.serviceconfig.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/miniapp/serviceconfig/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniMiniAppServiceConfigModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
