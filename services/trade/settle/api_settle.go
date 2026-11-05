package settle

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装支付宝 V3 settle 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor               core.Executor
	settleConfirm          core.Operation[SettleConfirmRequest, SettleConfirmResponse]
	settleReceivablesQuery core.Operation[SettleReceivablesQueryRequest, SettleReceivablesQueryResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	settleConfirm, err := core.NewJSONOperation[SettleConfirmRequest, SettleConfirmResponse](core.OperationSpec{
		Method: "alipay.trade.settle.confirm", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/settle/confirm",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	settleReceivablesQuery, err := core.NewJSONOperation[SettleReceivablesQueryRequest, SettleReceivablesQueryResponse](core.OperationSpec{
		Method: "alipay.trade.settle.receivables.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/settle/receivables/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:               executor,
		settleConfirm:          settleConfirm,
		settleReceivablesQuery: settleReceivablesQuery,
	}, nil
}

// SettleConfirm 调用 alipay.trade.settle.confirm 确认结算
// （V3 POST /v3/alipay/trade/settle/confirm）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/apis/0hp5md
// OutRequestNo 用于幂等控制；写操作，SDK 不自动重试。
func (s *Service) SettleConfirm(ctx context.Context, req SettleConfirmRequest,
	opts ...core.CallOption) (*core.Response[SettleConfirmResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.settleConfirm,
		core.Request[SettleConfirmRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SettleReceivablesQuery 调用 alipay.trade.settle.receivables.query
// 查询待结算金额（V3 POST /v3/alipay/trade/settle/receivables/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/pre-apis/05vlu7
func (s *Service) SettleReceivablesQuery(ctx context.Context, req SettleReceivablesQueryRequest,
	opts ...core.CallOption) (*core.Response[SettleReceivablesQueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.settleReceivablesQuery,
		core.Request[SettleReceivablesQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
