package anchor

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

// AnchorInfoQuery calls alipay.merchant.anchor.info.query (GET /v3/alipay/merchant/anchor/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05yqsw
func (s *Service) AnchorInfoQuery(ctx context.Context, req MerchantAnchorInfoQueryRequest, opts ...core.CallOption) (*core.Response[MerchantAnchorInfoQueryResponse], error) {
	return core.ExecuteJSON[MerchantAnchorInfoQueryRequest, MerchantAnchorInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.anchor.info.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/merchant/anchor/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantAnchorInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
