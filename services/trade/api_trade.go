package trade

import (
	"context"
	"fmt"
	"strings"

	"github.com/tiyee/go-alipay/core"
)

// 传统网关接口名、版本与响应节点名，依据各接口官方文档固定。
const (
	methodQuery       = "alipay.trade.query"
	methodPrecreate   = "alipay.trade.precreate"
	methodRefund      = "alipay.trade.refund"
	methodQueryRefund = "alipay.trade.fastpay.refund.query"
	methodClose       = "alipay.trade.close"

	version10 = "1.0"

	respKeyQuery       = "alipay_trade_query_response"
	respKeyPrecreate   = "alipay_trade_precreate_response"
	respKeyRefund      = "alipay_trade_refund_response"
	respKeyQueryRefund = "alipay_trade_fastpay_refund_query_response"
	respKeyClose       = "alipay_trade_close_response"
)

// Service 是交易领域的 API 封装，通过注入 core.Executor（通常为
// *core.Client）调用统一执行流程。Service 不实现 HTTP、签名或验签。
// 零值不可用；通过 NewService 构建。
type Service struct {
	executor core.Executor

	query       core.Operation[QueryRequest, QueryResponse]
	precreate   core.Operation[PrecreateRequest, PrecreateResponse]
	refund      core.Operation[RefundRequest, RefundResponse]
	queryRefund core.Operation[QueryRefundRequest, QueryRefundResponse]
	close       core.Operation[CloseRequest, CloseResponse]
}

// NewService 基于执行器构建交易 Service。执行器为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	query, err := core.NewJSONOperation[QueryRequest, QueryResponse](core.OperationSpec{
		Method: methodQuery, Version: version10, ResponseKey: respKeyQuery,
		ParameterMode: core.ParameterModeBizJSON,
	})
	if err != nil {
		return nil, err
	}
	precreate, err := core.NewJSONOperation[PrecreateRequest, PrecreateResponse](core.OperationSpec{
		Method: methodPrecreate, Version: version10, ResponseKey: respKeyPrecreate,
		ParameterMode: core.ParameterModeBizJSON,
	})
	if err != nil {
		return nil, err
	}
	refund, err := core.NewJSONOperation[RefundRequest, RefundResponse](core.OperationSpec{
		Method: methodRefund, Version: version10, ResponseKey: respKeyRefund,
		ParameterMode: core.ParameterModeBizJSON,
	})
	if err != nil {
		return nil, err
	}
	queryRefund, err := core.NewJSONOperation[QueryRefundRequest, QueryRefundResponse](core.OperationSpec{
		Method: methodQueryRefund, Version: version10, ResponseKey: respKeyQueryRefund,
		ParameterMode: core.ParameterModeBizJSON,
	})
	if err != nil {
		return nil, err
	}
	closeOp, err := core.NewJSONOperation[CloseRequest, CloseResponse](core.OperationSpec{
		Method: methodClose, Version: version10, ResponseKey: respKeyClose,
		ParameterMode: core.ParameterModeBizJSON,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:    executor,
		query:       query,
		precreate:   precreate,
		refund:      refund,
		queryRefund: queryRefund,
		close:       closeOp,
	}, nil
}

// Query 调用 alipay.trade.query 查询交易状态。
//
// 契约：https://opendocs.alipay.com/open/02ekfq.md。
// 调用成功（code=10000）不代表支付成功，业务结果以响应中的
// TradeStatus 为准；未生成真实交易时 TradeNo 与 BuyerLogonID 不返回。
func (s *Service) Query(ctx context.Context, req QueryRequest,
	opts ...core.CallOption) (*core.Response[QueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.query,
		core.Request[QueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// Precreate 调用 alipay.trade.precreate 生成当面付扫码支付二维码。
//
// 契约：https://opendocs.alipay.com/open/02np92。
// 二维码有效时间固定 2 小时；同一订单号不能重复申请二维码；
// qr_code 缺失时按业务响应处理，不代表预下单失败或成功。
func (s *Service) Precreate(ctx context.Context, req PrecreateRequest,
	opts ...core.CallOption) (*core.Response[PrecreateResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.precreate,
		core.Request[PrecreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// Refund 调用 alipay.trade.refund 发起交易退款。
//
// 契约：https://opendocs.alipay.com/open/02ekfs.md。
// code=10000 仅代表退款请求成功：FundChange=Y 表示退款成功，
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

// QueryRefund 调用 alipay.trade.fastpay.refund.query 查询退款结果。
//
// 契约：https://opendocs.alipay.com/open/02ekft.md。
// code=10000 仅代表查询操作成功；仅当响应中 RefundStatus 为
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

// Close 调用 alipay.trade.close 关闭未付款交易。
//
// 契约：https://opendocs.alipay.com/open/02ivbu.md。
// 仅"等待买家付款"状态的交易可关单；SDK 不依据本地状态自动发起关单。
func (s *Service) Close(ctx context.Context, req CloseRequest,
	opts ...core.CallOption) (*core.Response[CloseResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.close,
		core.Request[CloseRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// tradeIdentifier 校验 trade_no / out_trade_no 二选一约束：
// 不能同时为空；文档注明同时存在时网关优先使用 trade_no。
func tradeIdentifier(tradeNo, outTradeNo, field string) error {
	if tradeNo == "" && outTradeNo == "" {
		return &core.ValidationError{Field: field,
			Reason: "trade_no and out_trade_no cannot both be empty"}
	}
	return nil
}

func (r QueryRequest) validate() error {
	return tradeIdentifier(r.TradeNo, r.OutTradeNo, "out_trade_no")
}

func (r QueryRefundRequest) validate() error {
	if r.OutRequestNo == "" {
		return &core.ValidationError{Field: "out_request_no", Reason: "out_request_no is required"}
	}
	return tradeIdentifier(r.TradeNo, r.OutTradeNo, "out_trade_no")
}

func (r CloseRequest) validate() error {
	return tradeIdentifier(r.TradeNo, r.OutTradeNo, "out_trade_no")
}

func (r RefundRequest) validate() error {
	if err := validAmount(r.RefundAmount, "refund_amount"); err != nil {
		return err
	}
	return tradeIdentifier(r.TradeNo, r.OutTradeNo, "out_trade_no")
}

func (r PrecreateRequest) validate() error {
	if r.OutTradeNo == "" {
		return &core.ValidationError{Field: "out_trade_no", Reason: "out_trade_no is required"}
	}
	for _, ch := range r.OutTradeNo {
		if !(ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' || ch == '_') {
			return &core.ValidationError{Field: "out_trade_no",
				Reason: "out_trade_no only supports letters, digits and underscore"}
		}
	}
	if err := validAmount(r.TotalAmount, "total_amount"); err != nil {
		return err
	}
	if r.Subject == "" {
		return &core.ValidationError{Field: "subject", Reason: "subject is required"}
	}
	if strings.ContainsAny(r.Subject, "/=&") {
		return &core.ValidationError{Field: "subject",
			Reason: `subject must not contain "/", "=" or "&"`}
	}
	if r.DiscountableAmount != nil {
		if err := validAmount(*r.DiscountableAmount, "discountable_amount"); err != nil {
			return err
		}
	}
	if r.UndiscountableAmount != nil {
		if err := validAmount(*r.UndiscountableAmount, "undiscountable_amount"); err != nil {
			return err
		}
	}
	for i, goods := range r.GoodsDetail {
		if goods.GoodsID == "" || goods.GoodsName == "" || goods.Quantity <= 0 {
			return &core.ValidationError{Field: "goods_detail",
				Reason: fmt.Sprintf("goods_detail[%d] requires goods_id, goods_name and positive quantity", i)}
		}
		if err := validAmount(goods.Price, "goods_detail.price"); err != nil {
			return err
		}
	}
	return nil
}

// validAmount 校验金额为精确的十进制字符串：非负、至多两位小数、
// 至少含一个非零数字。金额不得使用浮点数建模或计算。
func validAmount(amount, field string) error {
	if amount == "" {
		return &core.ValidationError{Field: field, Reason: "amount is required"}
	}
	intPart, fracPart, found := strings.Cut(amount, ".")
	if intPart == "" && fracPart == "" {
		return &core.ValidationError{Field: field, Reason: "amount is not a decimal string"}
	}
	if len(intPart) > 1 && intPart[0] == '0' {
		return &core.ValidationError{Field: field, Reason: "amount has leading zeros"}
	}
	for _, part := range [2]string{intPart, fracPart} {
		for _, ch := range part {
			if ch < '0' || ch > '9' {
				return &core.ValidationError{Field: field, Reason: "amount is not a decimal string"}
			}
		}
	}
	if found && len(fracPart) > 2 {
		return &core.ValidationError{Field: field, Reason: "amount supports at most two decimal places"}
	}
	nonZero := strings.ContainsFunc(intPart, func(ch rune) bool { return ch != '0' }) ||
		strings.ContainsFunc(fracPart, func(ch rune) bool { return ch != '0' })
	if !nonZero {
		return &core.ValidationError{Field: field, Reason: "amount must be greater than zero"}
	}
	return nil
}
