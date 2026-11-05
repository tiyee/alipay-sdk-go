package innerversion

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

// MiniInnerVersionBackDevPublish calls alipay.open.mini.innerversion.backdev.publish (POST /v3/alipay/open/mini/innerversion/backdev/publish).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xqxi
func (s *Service) MiniInnerVersionBackDevPublish(ctx context.Context, req MiniInnerVersionBackDevPublishRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionBackDevPublishResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionBackDevPublishRequest, MiniInnerVersionBackDevPublishResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.backdev.publish", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/backdev/publish",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionBackDevPublishRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionConditionBatchQuery calls alipay.open.mini.innerversion.condition.batchquery (POST /v3/alipay/open/mini/innerversion/condition/batchquery).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xk66
func (s *Service) MiniInnerVersionConditionBatchQuery(ctx context.Context, req MiniInnerVersionConditionBatchQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionConditionBatchQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionConditionBatchQueryRequest, MiniInnerVersionConditionBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.condition.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/condition/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionConditionBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionContentRollback calls alipay.open.mini.innerversion.content.rollback (POST /v3/alipay/open/mini/innerversion/content/rollback).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xp3p
func (s *Service) MiniInnerVersionContentRollback(ctx context.Context, req MiniInnerVersionContentRollbackRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionContentRollbackResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionContentRollbackRequest, MiniInnerVersionContentRollbackResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.content.rollback", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/content/rollback",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionContentRollbackRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionInfoDelete calls alipay.open.mini.innerversion.info.delete (DELETE /v3/alipay/open/mini/innerversion/info/delete).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xe18
func (s *Service) MiniInnerVersionInfoDelete(ctx context.Context, req MiniInnerVersionInfoDeleteRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionInfoDeleteResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionInfoDeleteRequest, MiniInnerVersionInfoDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.info.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/mini/innerversion/info/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionInfoDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionInfoQuery calls alipay.open.mini.innerversion.info.query (GET /v3/alipay/open/mini/innerversion/info/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xdzi
func (s *Service) MiniInnerVersionInfoQuery(ctx context.Context, req MiniInnerVersionInfoQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionInfoQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionInfoQueryRequest, MiniInnerVersionInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.info.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/innerversion/info/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionInstantiationQuery calls alipay.open.mini.innerversion.instantiation.query (GET /v3/alipay/open/mini/innerversion/instantiation/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xdyz
func (s *Service) MiniInnerVersionInstantiationQuery(ctx context.Context, req MiniInnerVersionInstantiationQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionInstantiationQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionInstantiationQueryRequest, MiniInnerVersionInstantiationQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.instantiation.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/innerversion/instantiation/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionInstantiationQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionLastQuery calls alipay.open.mini.innerversion.last.query (GET /v3/alipay/open/mini/innerversion/last/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xhs8
func (s *Service) MiniInnerVersionLastQuery(ctx context.Context, req MiniInnerVersionLastQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionLastQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionLastQueryRequest, MiniInnerVersionLastQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.last.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/innerversion/last/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionLastQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionModelForOfflineQuery calls alipay.open.mini.innerversion.modelforoffline.query (GET /v3/alipay/open/mini/innerversion/modelforoffline/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xqwp
func (s *Service) MiniInnerVersionModelForOfflineQuery(ctx context.Context, req MiniInnerVersionModelForOfflineQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionModelForOfflineQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionModelForOfflineQueryRequest, MiniInnerVersionModelForOfflineQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.modelforoffline.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/innerversion/modelforoffline/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionModelForOfflineQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionOfflinePublish calls alipay.open.mini.innerversion.offline.publish (POST /v3/alipay/open/mini/innerversion/offline/publish).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xo7b
func (s *Service) MiniInnerVersionOfflinePublish(ctx context.Context, req MiniInnerVersionOfflinePublishRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionOfflinePublishResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionOfflinePublishRequest, MiniInnerVersionOfflinePublishResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.offline.publish", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/offline/publish",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionOfflinePublishRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionOnline calls alipay.open.mini.innerversion.online (POST /v3/alipay/open/mini/innerversion/online).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xhq6
func (s *Service) MiniInnerVersionOnline(ctx context.Context, req MiniInnerVersionOnlineRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionOnlineResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionOnlineRequest, MiniInnerVersionOnlineResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.online", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/online",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionOnlineRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionOnlinePublish calls alipay.open.mini.innerversion.online.publish (POST /v3/alipay/open/mini/innerversion/online/publish).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xgf5
func (s *Service) MiniInnerVersionOnlinePublish(ctx context.Context, req MiniInnerVersionOnlinePublishRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionOnlinePublishResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionOnlinePublishRequest, MiniInnerVersionOnlinePublishResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.online.publish", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/online/publish",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionOnlinePublishRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionPackageCreate calls alipay.open.mini.innerversion.package.create (POST /v3/alipay/open/mini/innerversion/package/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05x9fo
func (s *Service) MiniInnerVersionPackageCreate(ctx context.Context, req MiniInnerVersionPackageCreateRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionPackageCreateResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionPackageCreateRequest, MiniInnerVersionPackageCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.package.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/package/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionPackageCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionSync calls alipay.open.mini.innerversion.sync (POST /v3/alipay/open/mini/innerversion/sync).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xlnt
func (s *Service) MiniInnerVersionSync(ctx context.Context, req MiniInnerVersionSyncRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionSyncResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionSyncRequest, MiniInnerVersionSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
