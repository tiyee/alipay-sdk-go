package royalty

import (
	"context"
	"fmt"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装支付宝 V3 royalty 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor                  core.Executor
	royaltyRateQuery          core.Operation[RoyaltyRateQueryRequest, RoyaltyRateQueryResponse]
	royaltyRelationBind       core.Operation[RoyaltyRelationBindRequest, RoyaltyRelationBindResponse]
	royaltyRelationUnbind     core.Operation[RoyaltyRelationUnbindRequest, RoyaltyRelationUnbindResponse]
	royaltyRelationBatchquery core.Operation[RoyaltyRelationBatchQueryRequest, RoyaltyRelationBatchQueryResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	royaltyRateQuery, err := core.NewJSONOperation[RoyaltyRateQueryRequest, RoyaltyRateQueryResponse](core.OperationSpec{
		Method: "alipay.trade.royalty.rate.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/trade/royalty/rate/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	royaltyRelationBind, err := core.NewJSONOperation[RoyaltyRelationBindRequest, RoyaltyRelationBindResponse](core.OperationSpec{
		Method: "alipay.trade.royalty.relation.bind", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/royalty/relation/bind",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	royaltyRelationUnbind, err := core.NewJSONOperation[RoyaltyRelationUnbindRequest, RoyaltyRelationUnbindResponse](core.OperationSpec{
		Method: "alipay.trade.royalty.relation.unbind", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/royalty/relation/unbind",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	royaltyRelationBatchquery, err := core.NewJSONOperation[RoyaltyRelationBatchQueryRequest, RoyaltyRelationBatchQueryResponse](core.OperationSpec{
		Method: "alipay.trade.royalty.relation.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/royalty/relation/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:                  executor,
		royaltyRateQuery:          royaltyRateQuery,
		royaltyRelationBind:       royaltyRelationBind,
		royaltyRelationUnbind:     royaltyRelationUnbind,
		royaltyRelationBatchquery: royaltyRelationBatchquery,
	}, nil
}

// RoyaltyRateQuery 调用 alipay.trade.royalty.rate.query 查询分账比例
// （V3 GET /v3/alipay/trade/royalty/rate/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/d9aece9f_alipay.trade.royalty.rate.query
func (s *Service) RoyaltyRateQuery(ctx context.Context, req RoyaltyRateQueryRequest,
	opts ...core.CallOption) (*core.Response[RoyaltyRateQueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.royaltyRateQuery,
		core.Request[RoyaltyRateQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// RoyaltyRelationBind 调用 alipay.trade.royalty.relation.bind 绑定
// 分账关系（V3 POST /v3/alipay/trade/royalty/relation/bind）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/e701637d_alipay.trade.royalty.relation.bind
// 绑定结果以 ResultCode 为准。
func (s *Service) RoyaltyRelationBind(ctx context.Context, req RoyaltyRelationBindRequest,
	opts ...core.CallOption) (*core.Response[RoyaltyRelationBindResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.royaltyRelationBind,
		core.Request[RoyaltyRelationBindRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// RoyaltyRelationUnbind 调用 alipay.trade.royalty.relation.unbind 解绑
// 分账关系（V3 POST /v3/alipay/trade/royalty/relation/unbind）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/302cd1f8_alipay.trade.royalty.relation.unbind
func (s *Service) RoyaltyRelationUnbind(ctx context.Context, req RoyaltyRelationUnbindRequest,
	opts ...core.CallOption) (*core.Response[RoyaltyRelationUnbindResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.royaltyRelationUnbind,
		core.Request[RoyaltyRelationUnbindRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// RoyaltyRelationBatchQuery 调用 alipay.trade.royalty.relation.batchquery
// 查询分账关系（V3 POST /v3/alipay/trade/royalty/relation/batchquery）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/1a11475c_alipay.trade.royalty.relation.batchquery
func (s *Service) RoyaltyRelationBatchQuery(ctx context.Context, req RoyaltyRelationBatchQueryRequest,
	opts ...core.CallOption) (*core.Response[RoyaltyRelationBatchQueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.royaltyRelationBatchquery,
		core.Request[RoyaltyRelationBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// validateReceiverList 校验分账接收方列表：非空、最多 20 个、类型合法。
func validateReceiverList(list []RoyaltyEntity) error {
	if len(list) == 0 {
		return &core.ValidationError{Field: "receiver_list", Reason: "receiver_list is required"}
	}
	if len(list) > 20 {
		return &core.ValidationError{Field: "receiver_list",
			Reason: "receiver_list supports at most 20 entries"}
	}
	for i, entity := range list {
		if !entity.Type.IsKnown() {
			return &core.ValidationError{Field: "receiver_list.type",
				Reason: "receiver type must be userId, loginName or openId"}
		}
		if entity.Type != RoyaltyEntityTypeOpenID && (entity.Account == nil || *entity.Account == "") {
			return &core.ValidationError{Field: "receiver_list.account",
				Reason: fmt.Sprintf("receiver_list[%d].account is required for the declared type", i)}
		}
	}
	return nil
}
