package live

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

// LiveChannelQuery calls alipay.merchant.live.channel.query (GET /v3/alipay/merchant/live/channel/query).
// Official V3 contract: https://opendocs.alipay.com/apis/069r05
func (s *Service) LiveChannelQuery(ctx context.Context, req MerchantLiveChannelQueryRequest, opts ...core.CallOption) (*core.Response[MerchantLiveChannelQueryResponse], error) {
	return core.ExecuteJSON[MerchantLiveChannelQueryRequest, MerchantLiveChannelQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.live.channel.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/merchant/live/channel/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantLiveChannelQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// LiveItemPlanModify calls alipay.merchant.live.itemplan.modify (POST /v3/alipay/merchant/live/itemplan/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/069o8w
func (s *Service) LiveItemPlanModify(ctx context.Context, req MerchantLiveItemPlanModifyRequest, opts ...core.CallOption) (*core.Response[MerchantLiveItemPlanModifyResponse], error) {
	return core.ExecuteJSON[MerchantLiveItemPlanModifyRequest, MerchantLiveItemPlanModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.live.itemplan.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/live/itemplan/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantLiveItemPlanModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
