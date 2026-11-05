package service

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

// AppServiceApply calls alipay.open.app.service.apply (POST /v3/alipay/open/app/service/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8b18b82d_alipay.open.app.service.apply
func (s *Service) AppServiceApply(ctx context.Context, req AppServiceApplyRequest, opts ...core.CallOption) (*core.Response[AppServiceApplyResponse], error) {
	return core.ExecuteJSON[AppServiceApplyRequest, AppServiceApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.service.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/app/service/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppServiceApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppServiceDelete calls alipay.open.app.service.delete (DELETE /v3/alipay/open/app/service/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/5b7ade20_alipay.open.app.service.delete
func (s *Service) AppServiceDelete(ctx context.Context, req AppServiceDeleteRequest, opts ...core.CallOption) (*core.Response[AppServiceDeleteResponse], error) {
	return core.ExecuteJSON[AppServiceDeleteRequest, AppServiceDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.service.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/app/service/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppServiceDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppServiceListQuery calls alipay.open.app.service.list.query (GET /v3/alipay/open/app/service/list/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/77d4e708_alipay.open.app.service.list.query
func (s *Service) AppServiceListQuery(ctx context.Context, req AppServiceListQueryRequest, opts ...core.CallOption) (*core.Response[AppServiceListQueryResponse], error) {
	return core.ExecuteJSON[AppServiceListQueryRequest, AppServiceListQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.service.list.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/app/service/list/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppServiceListQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppServiceQuery calls alipay.open.app.service.query (GET /v3/alipay/open/app/service/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/86469046_alipay.open.app.service.query
func (s *Service) AppServiceQuery(ctx context.Context, req AppServiceQueryRequest, opts ...core.CallOption) (*core.Response[AppServiceQueryResponse], error) {
	return core.ExecuteJSON[AppServiceQueryRequest, AppServiceQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.service.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/app/service/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppServiceQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppServiceSchemaQuery calls alipay.open.app.service.schema.query (GET /v3/alipay/open/app/service/schema/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/afe544b5_alipay.open.app.service.schema.query
func (s *Service) AppServiceSchemaQuery(ctx context.Context, req AppServiceSchemaQueryRequest, opts ...core.CallOption) (*core.Response[AppServiceSchemaQueryResponse], error) {
	return core.ExecuteJSON[AppServiceSchemaQueryRequest, AppServiceSchemaQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.service.schema.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/app/service/schema/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppServiceSchemaQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
