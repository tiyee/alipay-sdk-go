package blueseaactivity

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

// SpBlueSeaActivityCreate calls alipay.open.sp.blueseaactivity.create (POST /v3/alipay/open/sp/blueseaactivity/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a974b7ae_alipay.open.sp.blueseaactivity.create
func (s *Service) SpBlueSeaActivityCreate(ctx context.Context, req SpBlueSeaActivityCreateRequest, opts ...core.CallOption) (*core.Response[SpBlueSeaActivityCreateResponse], error) {
	return core.ExecuteJSON[SpBlueSeaActivityCreateRequest, SpBlueSeaActivityCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.blueseaactivity.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/sp/blueseaactivity/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpBlueSeaActivityCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SpBlueSeaActivityModify calls alipay.open.sp.blueseaactivity.modify (POST /v3/alipay/open/sp/blueseaactivity/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8cbcb343_alipay.open.sp.blueseaactivity.modify
func (s *Service) SpBlueSeaActivityModify(ctx context.Context, req SpBlueSeaActivityModifyRequest, opts ...core.CallOption) (*core.Response[SpBlueSeaActivityModifyResponse], error) {
	return core.ExecuteJSON[SpBlueSeaActivityModifyRequest, SpBlueSeaActivityModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.blueseaactivity.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/sp/blueseaactivity/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpBlueSeaActivityModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SpBlueSeaActivityQuery calls alipay.open.sp.blueseaactivity.query (GET /v3/alipay/open/sp/blueseaactivity/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4ed0241a_alipay.open.sp.blueseaactivity.query
func (s *Service) SpBlueSeaActivityQuery(ctx context.Context, req SpBlueSeaActivityQueryRequest, opts ...core.CallOption) (*core.Response[SpBlueSeaActivityQueryResponse], error) {
	return core.ExecuteJSON[SpBlueSeaActivityQueryRequest, SpBlueSeaActivityQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.blueseaactivity.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/sp/blueseaactivity/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpBlueSeaActivityQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
