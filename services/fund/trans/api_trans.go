package trans

import (
	"context"
	"strings"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor          core.Executor
	orderQuery        core.Operation[OrderQueryRequest, OrderQueryResponse]
	commonQuery       core.Operation[CommonQueryRequest, CommonQueryResponse]
	payeeBindQuery    core.Operation[PayeeBindQueryRequest, PayeeBindQueryResponse]
	uniTransfer       core.Operation[UniTransferRequest, UniTransferResponse]
	toaccountTransfer core.Operation[ToAccountTransferRequest, ToAccountTransferResponse]
	pay               core.Operation[PayRequest, PayResponse]
	refund            core.Operation[RefundRequest, RefundResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	orderQuery, err := core.NewJSONOperation[OrderQueryRequest, OrderQueryResponse](core.OperationSpec{
		Method: "alipay.fund.trans.order.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/trans/order/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	commonQuery, err := core.NewJSONOperation[CommonQueryRequest, CommonQueryResponse](core.OperationSpec{
		Method: "alipay.fund.trans.common.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/trans/common/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	payeeBindQuery, err := core.NewJSONOperation[PayeeBindQueryRequest, PayeeBindQueryResponse](core.OperationSpec{
		Method: "alipay.fund.trans.payee.bind.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/trans/payee/bind/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	uniTransfer, err := core.NewJSONOperation[UniTransferRequest, UniTransferResponse](core.OperationSpec{
		Method: "alipay.fund.trans.uni.transfer", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/trans/uni/transfer",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementCertificate,
	})
	if err != nil {
		return nil, err
	}
	toaccountTransfer, err := core.NewJSONOperation[ToAccountTransferRequest, ToAccountTransferResponse](core.OperationSpec{
		Method: "alipay.fund.trans.toaccount.transfer", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/trans/toaccount/transfer",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	pay, err := core.NewJSONOperation[PayRequest, PayResponse](core.OperationSpec{
		Method: "alipay.fund.trans.pay", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/trans/pay",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	refund, err := core.NewJSONOperation[RefundRequest, RefundResponse](core.OperationSpec{
		Method: "alipay.fund.trans.refund", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/trans/refund",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:          executor,
		orderQuery:        orderQuery,
		commonQuery:       commonQuery,
		payeeBindQuery:    payeeBindQuery,
		uniTransfer:       uniTransfer,
		toaccountTransfer: toaccountTransfer,
		pay:               pay,
		refund:            refund,
	}, nil
}

// validateParticipant 校验参与方标识与类型成对出现。
func validateParticipant(p *Participant) error {
	if (p.Identity == nil || *p.Identity == "") != (p.IdentityType == nil || *p.IdentityType == "") {
		return &core.ValidationError{Field: "identity",
			Reason: "identity and identity_type must be provided together"}
	}
	return nil
}

// amountLessThan 以十进制字符串比较金额 a < b（不含相等）；
// 仅用于本地校验提示，不做浮点运算。
func amountLessThan(a, b string) bool {
	ai, af, _ := strings.Cut(a, ".")
	bi, bf, _ := strings.Cut(b, ".")
	ai, af = normalizeDigits(ai, af, 2)
	bi, bf = normalizeDigits(bi, bf, 2)
	if len(ai) != len(bi) {
		return len(ai) < len(bi)
	}
	if ai != bi {
		return ai < bi
	}
	return af < bf
}

// normalizeDigits 将整数与小数部分补零对齐到固定小数位。
func normalizeDigits(intPart, fracPart string, fracLen int) (string, string) {
	for len(intPart) > 1 && intPart[0] == '0' {
		intPart = intPart[1:]
	}
	for len(fracPart) < fracLen {
		fracPart += "0"
	}
	return intPart, fracPart[:fracLen]
}

// OrderQuery 调用 alipay.fund.trans.order.query 查询转账订单
// （V3 GET /v3/alipay/fund/trans/order/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/f00381dd_alipay.fund.trans.order.query
// 查询成功不代表到账；DEALING/REFUND 等状态需按业务继续跟踪。
func (s *Service) OrderQuery(ctx context.Context, req OrderQueryRequest,
	opts ...core.CallOption) (*core.Response[OrderQueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.orderQuery,
		core.Request[OrderQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CommonQuery 调用 alipay.fund.trans.common.query 通用查询单据
// （V3 GET /v3/alipay/fund/trans/common/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/bcbe6351_alipay.fund.trans.common.query?scene=common&pathHash=36f664d9
func (s *Service) CommonQuery(ctx context.Context, req CommonQueryRequest,
	opts ...core.CallOption) (*core.Response[CommonQueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.commonQuery,
		core.Request[CommonQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PayeeBindQuery 调用 alipay.fund.trans.payee.bind.query 查询收款
// 账号绑定状态（V3 GET /v3/alipay/fund/trans/payee/bind/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/03fe4a0d_alipay.fund.trans.payee.bind.query?scene=common&pathHash=377d1f29
func (s *Service) PayeeBindQuery(ctx context.Context, req PayeeBindQueryRequest,
	opts ...core.CallOption) (*core.Response[PayeeBindQueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.payeeBindQuery,
		core.Request[PayeeBindQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// UniTransfer 调用 alipay.fund.trans.uni.transfer 发起单笔转账或
// 现金红包（V3 POST /v3/alipay/fund/trans/uni/transfer）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/26a22ad7_alipay.fund.trans.uni.transfer?scene=common%2Cca56bca529e64125a2786703c6192d41%2C66dd06f5a923403393b85de68d3c0055&pathHash=9b52afe2
// 资金写操作：SDK 不自动重试；超时或状态未知时以订单查询确认；
// 资金支出要求证书模式：https://opendocs.alipay.com/common/02kg66.md。
func (s *Service) UniTransfer(ctx context.Context, req UniTransferRequest,
	opts ...core.CallOption) (*core.Response[UniTransferResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.uniTransfer,
		core.Request[UniTransferRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ToAccountTransfer 调用 alipay.fund.trans.toaccount.transfer 转账到
// 支付宝账户（V3 POST /v3/alipay/fund/trans/toaccount/transfer）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/53a4ce8e_alipay.fund.trans.toaccount.transfer
// 金额必须 ≥0.1 元；写操作，SDK 不自动重试。
func (s *Service) ToAccountTransfer(ctx context.Context, req ToAccountTransferRequest,
	opts ...core.CallOption) (*core.Response[ToAccountTransferResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.toaccountTransfer,
		core.Request[ToAccountTransferRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// Pay 调用 alipay.fund.trans.pay 发起资金支付
// （V3 POST /v3/alipay/fund/trans/pay）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/pre-apis/05xggg
// status=UNKNOWN 时必须查询确认；写操作，SDK 不自动重试。
func (s *Service) Pay(ctx context.Context, req PayRequest,
	opts ...core.CallOption) (*core.Response[PayResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.pay,
		core.Request[PayRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// Refund 调用 alipay.fund.trans.refund 发起资金退回
// （V3 POST /v3/alipay/fund/trans/refund）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/0c6ff10a_alipay.fund.trans.refund
// 退回失败重新提交必须沿用原 OutRequestNo，防止重复退回。
func (s *Service) Refund(ctx context.Context, req RefundRequest,
	opts ...core.CallOption) (*core.Response[RefundResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.refund,
		core.Request[RefundRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
