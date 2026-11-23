package api

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

// AppAPIFieldApply calls alipay.open.app.api.field.apply (POST /v3/alipay/open/app/api/field/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/db6692ea_alipay.open.app.api.field.apply
func (s *Service) AppAPIFieldApply(ctx context.Context, req AppAPIFieldApplyRequest, opts ...core.CallOption) (*core.Response[AppAPIFieldApplyResponse], error) {
	return core.ExecuteJSON[AppAPIFieldApplyRequest, AppAPIFieldApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.api.field.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/app/api/field/apply",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppAPIFieldApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppAPIFieldQuery calls alipay.open.app.api.field.query (GET /v3/alipay/open/app/api/field/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/0cfbef83_alipay.open.app.api.field.query
func (s *Service) AppAPIFieldQuery(ctx context.Context, req AppAPIFieldQueryRequest, opts ...core.CallOption) (*core.Response[AppAPIFieldQueryResponse], error) {
	return core.ExecuteJSON[AppAPIFieldQueryRequest, AppAPIFieldQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.api.field.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/app/api/field/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppAPIFieldQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppAPIQuery calls alipay.open.app.api.query (GET /v3/alipay/open/app/api/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/d091ad0b_alipay.open.app.api.query
func (s *Service) AppAPIQuery(ctx context.Context, req AppAPIQueryRequest, opts ...core.CallOption) (*core.Response[AppAPIQueryResponse], error) {
	return core.ExecuteJSON[AppAPIQueryRequest, AppAPIQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.api.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/app/api/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppAPIQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppAPISceneQuery calls alipay.open.app.api.scene.query (GET /v3/alipay/open/app/api/scene/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2bc4666c_alipay.open.app.api.scene.query
func (s *Service) AppAPISceneQuery(ctx context.Context, req AppAPISceneQueryRequest, opts ...core.CallOption) (*core.Response[AppAPISceneQueryResponse], error) {
	return core.ExecuteJSON[AppAPISceneQueryRequest, AppAPISceneQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.api.scene.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/app/api/scene/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppAPISceneQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
