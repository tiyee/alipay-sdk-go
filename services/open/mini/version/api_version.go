package version

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

// MiniVersionAuditedCancel calls alipay.open.mini.version.audited.cancel (POST /v3/alipay/open/mini/version/audited/cancel).
// Official V3 contract: https://opendocs.alipay.com/open-v3/7a284018_alipay.open.mini.version.audited.cancel
func (s *Service) MiniVersionAuditedCancel(ctx context.Context, req MiniVersionAuditedCancelRequest, opts ...core.CallOption) (*core.Response[MiniVersionAuditedCancelResponse], error) {
	return core.ExecuteJSON[MiniVersionAuditedCancelRequest, MiniVersionAuditedCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.audited.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/version/audited/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionAuditedCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniVersionBuildQuery calls alipay.open.mini.version.build.query (GET /v3/alipay/open/mini/version/build/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/dda696d3_alipay.open.mini.version.build.query
func (s *Service) MiniVersionBuildQuery(ctx context.Context, req MiniVersionBuildQueryRequest, opts ...core.CallOption) (*core.Response[MiniVersionBuildQueryResponse], error) {
	return core.ExecuteJSON[MiniVersionBuildQueryRequest, MiniVersionBuildQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.build.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/version/build/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionBuildQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniVersionDelete calls alipay.open.mini.version.delete (POST /v3/alipay/open/mini/version/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/eef61c13_alipay.open.mini.version.delete
func (s *Service) MiniVersionDelete(ctx context.Context, req MiniVersionDeleteRequest, opts ...core.CallOption) (*core.Response[MiniVersionDeleteResponse], error) {
	return core.ExecuteJSON[MiniVersionDeleteRequest, MiniVersionDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/version/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniVersionDetailQuery calls alipay.open.mini.version.detail.query (GET /v3/alipay/open/mini/version/detail/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/048e5e71_alipay.open.mini.version.detail.query
func (s *Service) MiniVersionDetailQuery(ctx context.Context, req MiniVersionDetailQueryRequest, opts ...core.CallOption) (*core.Response[MiniVersionDetailQueryResponse], error) {
	return core.ExecuteJSON[MiniVersionDetailQueryRequest, MiniVersionDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.detail.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/version/detail/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniVersionListQuery calls alipay.open.mini.version.list.query (GET /v3/alipay/open/mini/version/list/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/ac91bcd8_alipay.open.mini.version.list.query
func (s *Service) MiniVersionListQuery(ctx context.Context, req MiniVersionListQueryRequest, opts ...core.CallOption) (*core.Response[MiniVersionListQueryResponse], error) {
	return core.ExecuteJSON[MiniVersionListQueryRequest, MiniVersionListQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.list.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/version/list/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionListQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniVersionOffline calls alipay.open.mini.version.offline (POST /v3/alipay/open/mini/version/offline).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c5ffdf9c_alipay.open.mini.version.offline
func (s *Service) MiniVersionOffline(ctx context.Context, req MiniVersionOfflineRequest, opts ...core.CallOption) (*core.Response[MiniVersionOfflineResponse], error) {
	return core.ExecuteJSON[MiniVersionOfflineRequest, MiniVersionOfflineResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.offline", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/version/offline",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionOfflineRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniVersionOnline calls alipay.open.mini.version.online (POST /v3/alipay/open/mini/version/online).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4d9b2578_alipay.open.mini.version.online
func (s *Service) MiniVersionOnline(ctx context.Context, req MiniVersionOnlineRequest, opts ...core.CallOption) (*core.Response[MiniVersionOnlineResponse], error) {
	return core.ExecuteJSON[MiniVersionOnlineRequest, MiniVersionOnlineResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.online", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/version/online",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionOnlineRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniVersionRollback calls alipay.open.mini.version.rollback (POST /v3/alipay/open/mini/version/rollback).
// Official V3 contract: https://opendocs.alipay.com/open-v3/f1c39e2b_alipay.open.mini.version.rollback
func (s *Service) MiniVersionRollback(ctx context.Context, req MiniVersionRollbackRequest, opts ...core.CallOption) (*core.Response[MiniVersionRollbackResponse], error) {
	return core.ExecuteJSON[MiniVersionRollbackRequest, MiniVersionRollbackResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.rollback", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/version/rollback",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionRollbackRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniVersionUpload calls alipay.open.mini.version.upload (POST /v3/alipay/open/mini/version/upload).
// Official V3 contract: https://opendocs.alipay.com/open-v3/9d32147b_alipay.open.mini.version.upload
func (s *Service) MiniVersionUpload(ctx context.Context, req MiniVersionUploadRequest, opts ...core.CallOption) (*core.Response[MiniVersionUploadResponse], error) {
	return core.ExecuteJSON[MiniVersionUploadRequest, MiniVersionUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/version/upload",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
