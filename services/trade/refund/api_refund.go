package refund

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装支付宝 V3 refund 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor    core.Executor
	refundApply core.Operation[RefundApplyRequest, RefundApplyResponse]
	refund      core.Operation[RefundRequest, RefundResponse]
	queryRefund core.Operation[QueryRefundRequest, QueryRefundResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	refundApply, err := core.NewJSONOperation[RefundApplyRequest, RefundApplyResponse](core.OperationSpec{
		Method: "alipay.trade.refund.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/refund/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	refund, err := core.NewJSONOperation[RefundRequest, RefundResponse](core.OperationSpec{
		Method: "alipay.trade.refund", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/refund",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	queryRefund, err := core.NewJSONOperation[QueryRefundRequest, QueryRefundResponse](core.OperationSpec{
		Method: "alipay.trade.fastpay.refund.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/fastpay/refund/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:    executor,
		refundApply: refundApply,
		refund:      refund,
		queryRefund: queryRefund,
	}, nil
}

// RefundApply 调用 alipay.trade.refund.apply 受理异步退款
// （V3 POST /v3/alipay/trade/refund/apply）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/apis/05wbvf
// RefundStatus 为 REFUND_PROCESSING 时退款仍在处理中，需通过
// QueryRefund 确认终态；SDK 不自动重试。
func (s *Service) RefundApply(ctx context.Context, req RefundApplyRequest,
	opts ...core.CallOption) (*core.Response[RefundApplyResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.refundApply,
		core.Request[RefundApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// Refund 调用 alipay.trade.refund 发起交易退款
// （V3 POST /v3/alipay/trade/refund）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/6b16d4a2_alipay.trade.refund?scene=common&pathHash=00ab12b7
// 请求成功仅代表退款请求受理：FundChange=Y 表示退款成功，
// N 或未返回时需通过 QueryRefund 进一步确认。
// 重试必须保持 OutRequestNo 与 RefundAmount 不变，防止重复退款；
// SDK 不自动重试退款。
func (s *Service) Refund(ctx context.Context, req RefundRequest,
	opts ...core.CallOption) (*core.Response[RefundResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.refund,
		core.Request[RefundRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// QueryRefund 调用 alipay.trade.fastpay.refund.query 查询退款结果
// （V3 POST /v3/alipay/trade/fastpay/refund/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/b9ef37bd_alipay.trade.fastpay.refund.query?scene=common&pathHash=f071344e
// 请求成功仅代表查询操作成功；仅当响应中 RefundStatus 为
// REFUND_SUCCESS 时表示退款成功，字段未返回表示退款请求未收到
// 或退款失败。官方建议退款发起后间隔 10 秒以上再查询。
func (s *Service) QueryRefund(ctx context.Context, req QueryRefundRequest,
	opts ...core.CallOption) (*core.Response[QueryRefundResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.queryRefund,
		core.Request[QueryRefundRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
