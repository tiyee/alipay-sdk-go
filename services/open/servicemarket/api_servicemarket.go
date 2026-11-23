package servicemarket

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

// ServiceMarketCommodityShopOffline calls alipay.open.servicemarket.commodity.shop.offline (POST /v3/alipay/open/servicemarket/commodity/shop/offline).
// Official V3 contract: https://opendocs.alipay.com/open-v3/f56986eb_alipay.open.servicemarket.commodity.shop.offline
func (s *Service) ServiceMarketCommodityShopOffline(ctx context.Context, req ServiceMarketCommodityShopOfflineRequest, opts ...core.CallOption) (*core.Response[ServiceMarketCommodityShopOfflineResponse], error) {
	return core.ExecuteJSON[ServiceMarketCommodityShopOfflineRequest, ServiceMarketCommodityShopOfflineResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.servicemarket.commodity.shop.offline", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/servicemarket/commodity/shop/offline",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ServiceMarketCommodityShopOfflineRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ServiceMarketPromotionTaskRelationQuery calls alipay.open.servicemarket.promotiontask.relation.query (GET /v3/alipay/open/servicemarket/promotiontask/relation/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/76898459_alipay.open.servicemarket.promotiontask.relation.query
func (s *Service) ServiceMarketPromotionTaskRelationQuery(ctx context.Context, req ServiceMarketPromotionTaskRelationQueryRequest, opts ...core.CallOption) (*core.Response[ServiceMarketPromotionTaskRelationQueryResponse], error) {
	return core.ExecuteJSON[ServiceMarketPromotionTaskRelationQueryRequest, ServiceMarketPromotionTaskRelationQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.servicemarket.promotiontask.relation.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/servicemarket/promotiontask/relation/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ServiceMarketPromotionTaskRelationQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
