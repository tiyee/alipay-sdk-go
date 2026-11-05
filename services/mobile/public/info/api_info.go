package info

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

// PublicInfoQuery calls alipay.mobile.public.info.query (GET /v3/alipay/mobile/public/info/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuox
func (s *Service) PublicInfoQuery(ctx context.Context, req MobilePublicInfoQueryRequest, opts ...core.CallOption) (*core.Response[MobilePublicInfoQueryResponse], error) {
	return core.ExecuteJSON[MobilePublicInfoQueryRequest, MobilePublicInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.mobile.public.info.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/mobile/public/info/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MobilePublicInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
