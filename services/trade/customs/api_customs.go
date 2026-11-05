package customs

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装支付宝 V3 customs 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor       core.Executor
	customsDeclare core.Operation[CustomsDeclareRequest, CustomsDeclareResponse]
	customsQuery   core.Operation[CustomsQueryRequest, CustomsQueryResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	customsDeclare, err := core.NewJSONOperation[CustomsDeclareRequest, CustomsDeclareResponse](core.OperationSpec{
		Method: "alipay.trade.customs.declare", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/customs/declare",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	customsQuery, err := core.NewJSONOperation[CustomsQueryRequest, CustomsQueryResponse](core.OperationSpec{
		Method: "alipay.trade.customs.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/trade/customs/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:       executor,
		customsDeclare: customsDeclare,
		customsQuery:   customsQuery,
	}, nil
}

// CustomsDeclare 调用 alipay.trade.customs.declare 报关
// （V3 POST /v3/alipay/trade/customs/declare）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/pre-apis/05wyhb
// 报关为异步流程：受理成功不代表报关完成，以 CustomsQuery 查询为准。
func (s *Service) CustomsDeclare(ctx context.Context, req CustomsDeclareRequest,
	opts ...core.CallOption) (*core.Response[CustomsDeclareResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.customsDeclare,
		core.Request[CustomsDeclareRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CustomsQuery 调用 alipay.trade.customs.query 查询报关详细信息
// （V3 GET /v3/alipay/trade/customs/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/pre-apis/05wyjm
// 单次最多查询 10 个报关请求号。
func (s *Service) CustomsQuery(ctx context.Context, req CustomsQueryRequest,
	opts ...core.CallOption) (*core.Response[CustomsQueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.customsQuery,
		core.Request[CustomsQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
