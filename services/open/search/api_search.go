package search

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

// SearchAbilityBatchQuery calls alipay.open.search.ability.batchquery (POST /v3/alipay/open/search/ability/batchquery).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wggi
func (s *Service) SearchAbilityBatchQuery(ctx context.Context, req SearchAbilityBatchQueryRequest, opts ...core.CallOption) (*core.Response[SearchAbilityBatchQueryResponse], error) {
	return core.ExecuteJSON[SearchAbilityBatchQueryRequest, SearchAbilityBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.ability.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/ability/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchAbilityBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
