package oppor

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

// SpOpporDetailQuery calls alipay.open.sp.oppor.detail.query (GET /v3/alipay/open/sp/oppor/detail/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05vn6l
func (s *Service) SpOpporDetailQuery(ctx context.Context, req SpOpporDetailQueryRequest, opts ...core.CallOption) (*core.Response[SpOpporDetailQueryResponse], error) {
	return core.ExecuteJSON[SpOpporDetailQueryRequest, SpOpporDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.oppor.detail.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/sp/oppor/detail/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpOpporDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SpOpporPageQuery calls alipay.open.sp.oppor.page.query (GET /v3/alipay/open/sp/oppor/page/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05vdai
func (s *Service) SpOpporPageQuery(ctx context.Context, req SpOpporPageQueryRequest, opts ...core.CallOption) (*core.Response[SpOpporPageQueryResponse], error) {
	return core.ExecuteJSON[SpOpporPageQueryRequest, SpOpporPageQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.oppor.page.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/sp/oppor/page/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpOpporPageQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
