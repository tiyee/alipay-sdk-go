package order

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

// SearchBaseOrderModify calls alipay.open.search.baseorder.modify (POST /v3/alipay/open/search/baseorder/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/abbe8486_alipay.open.search.baseorder.modify
func (s *Service) SearchBaseOrderModify(ctx context.Context, req SearchBaseOrderModifyRequest, opts ...core.CallOption) (*core.Response[SearchBaseOrderModifyResponse], error) {
	return core.ExecuteJSON[SearchBaseOrderModifyRequest, SearchBaseOrderModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.baseorder.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/baseorder/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchBaseOrderModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchOrderDetailQuery calls alipay.open.search.orderdetail.query (POST /v3/alipay/open/search/orderdetail/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/d55fc0f0_alipay.open.search.orderdetail.query
func (s *Service) SearchOrderDetailQuery(ctx context.Context, req SearchOrderDetailQueryRequest, opts ...core.CallOption) (*core.Response[SearchOrderDetailQueryResponse], error) {
	return core.ExecuteJSON[SearchOrderDetailQueryRequest, SearchOrderDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.orderdetail.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/orderdetail/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchOrderDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchServiceOrderBatchQuery calls alipay.open.search.serviceorder.batchquery (POST /v3/alipay/open/search/serviceorder/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/13830d89_alipay.open.search.serviceorder.batchquery
func (s *Service) SearchServiceOrderBatchQuery(ctx context.Context, req SearchServiceOrderBatchQueryRequest, opts ...core.CallOption) (*core.Response[SearchServiceOrderBatchQueryResponse], error) {
	return core.ExecuteJSON[SearchServiceOrderBatchQueryRequest, SearchServiceOrderBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.serviceorder.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/serviceorder/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchServiceOrderBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
