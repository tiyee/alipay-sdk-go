package shop

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

// MarketShopBatchQuery calls alipay.offline.market.shop.batchquery (POST /v3/alipay/offline/market/shop/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05xf6f
func (s *Service) MarketShopBatchQuery(ctx context.Context, req OfflineMarketShopBatchQueryRequest, opts ...core.CallOption) (*core.Response[OfflineMarketShopBatchQueryResponse], error) {
	return core.ExecuteJSON[OfflineMarketShopBatchQueryRequest, OfflineMarketShopBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.offline.market.shop.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/offline/market/shop/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[OfflineMarketShopBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MarketShopCategoryQuery calls alipay.offline.market.shop.category.query (GET /v3/alipay/offline/market/shop/category/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05xdx8
func (s *Service) MarketShopCategoryQuery(ctx context.Context, req OfflineMarketShopCategoryQueryRequest, opts ...core.CallOption) (*core.Response[OfflineMarketShopCategoryQueryResponse], error) {
	return core.ExecuteJSON[OfflineMarketShopCategoryQueryRequest, OfflineMarketShopCategoryQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.offline.market.shop.category.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/offline/market/shop/category/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[OfflineMarketShopCategoryQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MarketShopCreate calls alipay.offline.market.shop.create (POST /v3/alipay/offline/market/shop/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05woku
func (s *Service) MarketShopCreate(ctx context.Context, req OfflineMarketShopCreateRequest, opts ...core.CallOption) (*core.Response[OfflineMarketShopCreateResponse], error) {
	return core.ExecuteJSON[OfflineMarketShopCreateRequest, OfflineMarketShopCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.offline.market.shop.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/offline/market/shop/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[OfflineMarketShopCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MarketShopModify calls alipay.offline.market.shop.modify (POST /v3/alipay/offline/market/shop/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wx1o
func (s *Service) MarketShopModify(ctx context.Context, req OfflineMarketShopModifyRequest, opts ...core.CallOption) (*core.Response[OfflineMarketShopModifyResponse], error) {
	return core.ExecuteJSON[OfflineMarketShopModifyRequest, OfflineMarketShopModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.offline.market.shop.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/offline/market/shop/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[OfflineMarketShopModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MarketShopQueryDetail calls alipay.offline.market.shop.querydetail (GET /v3/alipay/offline/market/shop/querydetail).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuly
func (s *Service) MarketShopQueryDetail(ctx context.Context, req OfflineMarketShopQueryDetailRequest, opts ...core.CallOption) (*core.Response[OfflineMarketShopQueryDetailResponse], error) {
	return core.ExecuteJSON[OfflineMarketShopQueryDetailRequest, OfflineMarketShopQueryDetailResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.offline.market.shop.querydetail", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/offline/market/shop/querydetail",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[OfflineMarketShopQueryDetailRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MarketShopSummaryBatchQuery calls alipay.offline.market.shop.summary.batchquery (POST /v3/alipay/offline/market/shop/summary/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wx2m
func (s *Service) MarketShopSummaryBatchQuery(ctx context.Context, req OfflineMarketShopSummaryBatchQueryRequest, opts ...core.CallOption) (*core.Response[OfflineMarketShopSummaryBatchQueryResponse], error) {
	return core.ExecuteJSON[OfflineMarketShopSummaryBatchQueryRequest, OfflineMarketShopSummaryBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.offline.market.shop.summary.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/offline/market/shop/summary/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[OfflineMarketShopSummaryBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
