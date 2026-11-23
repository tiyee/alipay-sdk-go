package members

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

// AppMembersCreate calls alipay.open.app.members.create (POST /v3/alipay/open/app/members/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8b6e0190_alipay.open.app.members.create
func (s *Service) AppMembersCreate(ctx context.Context, req AppMembersCreateRequest, opts ...core.CallOption) (*core.Response[AppMembersCreateResponse], error) {
	return core.ExecuteJSON[AppMembersCreateRequest, AppMembersCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.members.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/app/members/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppMembersCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppMembersDelete calls alipay.open.app.members.delete (DELETE /v3/alipay/open/app/members/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/26ef7be4_alipay.open.app.members.delete
func (s *Service) AppMembersDelete(ctx context.Context, req AppMembersDeleteRequest, opts ...core.CallOption) (*core.Response[AppMembersDeleteResponse], error) {
	return core.ExecuteJSON[AppMembersDeleteRequest, AppMembersDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.members.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/app/members/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppMembersDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AppMembersQuery calls alipay.open.app.members.query (GET /v3/alipay/open/app/members/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/6a8b8874_alipay.open.app.members.query
func (s *Service) AppMembersQuery(ctx context.Context, req AppMembersQueryRequest, opts ...core.CallOption) (*core.Response[AppMembersQueryResponse], error) {
	return core.ExecuteJSON[AppMembersQueryRequest, AppMembersQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.app.members.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/app/members/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AppMembersQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
