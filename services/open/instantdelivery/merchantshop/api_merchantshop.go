package merchantshop

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

// InstantDeliveryMerchantShopBatchQuery calls alipay.open.instantdelivery.merchantshop.batchquery (POST /v3/alipay/open/instantdelivery/merchantshop/batchquery).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vn66
func (s *Service) InstantDeliveryMerchantShopBatchQuery(ctx context.Context, req InstantDeliveryMerchantShopBatchQueryRequest, opts ...core.CallOption) (*core.Response[InstantDeliveryMerchantShopBatchQueryResponse], error) {
	return core.ExecuteJSON[InstantDeliveryMerchantShopBatchQueryRequest, InstantDeliveryMerchantShopBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.instantdelivery.merchantshop.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/instantdelivery/merchantshop/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InstantDeliveryMerchantShopBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InstantDeliveryMerchantShopCreate calls alipay.open.instantdelivery.merchantshop.create (POST /v3/alipay/open/instantdelivery/merchantshop/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vluh
func (s *Service) InstantDeliveryMerchantShopCreate(ctx context.Context, req InstantDeliveryMerchantShopCreateRequest, opts ...core.CallOption) (*core.Response[InstantDeliveryMerchantShopCreateResponse], error) {
	return core.ExecuteJSON[InstantDeliveryMerchantShopCreateRequest, InstantDeliveryMerchantShopCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.instantdelivery.merchantshop.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/instantdelivery/merchantshop/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InstantDeliveryMerchantShopCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InstantDeliveryMerchantShopModify calls alipay.open.instantdelivery.merchantshop.modify (POST /v3/alipay/open/instantdelivery/merchantshop/modify).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/068gkz
func (s *Service) InstantDeliveryMerchantShopModify(ctx context.Context, req InstantDeliveryMerchantShopModifyRequest, opts ...core.CallOption) (*core.Response[InstantDeliveryMerchantShopModifyResponse], error) {
	return core.ExecuteJSON[InstantDeliveryMerchantShopModifyRequest, InstantDeliveryMerchantShopModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.instantdelivery.merchantshop.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/instantdelivery/merchantshop/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InstantDeliveryMerchantShopModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InstantDeliveryMerchantShopQuery calls alipay.open.instantdelivery.merchantshop.query (GET /v3/alipay/open/instantdelivery/merchantshop/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/068ja3
func (s *Service) InstantDeliveryMerchantShopQuery(ctx context.Context, req InstantDeliveryMerchantShopQueryRequest, opts ...core.CallOption) (*core.Response[InstantDeliveryMerchantShopQueryResponse], error) {
	return core.ExecuteJSON[InstantDeliveryMerchantShopQueryRequest, InstantDeliveryMerchantShopQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.instantdelivery.merchantshop.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/instantdelivery/merchantshop/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InstantDeliveryMerchantShopQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
