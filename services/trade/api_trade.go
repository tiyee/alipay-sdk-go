package trade

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/internal/tradevalidate"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor  core.Executor
	pay       core.Operation[PayRequest, PayResponse]
	create    core.Operation[CreateRequest, CreateResponse]
	cancel    core.Operation[CancelRequest, CancelResponse]
	query     core.Operation[QueryRequest, QueryResponse]
	precreate core.Operation[PreCreateRequest, PreCreateResponse]
	close     core.Operation[CloseRequest, CloseResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	pay, err := core.NewJSONOperation[PayRequest, PayResponse](core.OperationSpec{
		Method: "alipay.trade.pay", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/pay",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	create, err := core.NewJSONOperation[CreateRequest, CreateResponse](core.OperationSpec{
		Method: "alipay.trade.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	cancel, err := core.NewJSONOperation[CancelRequest, CancelResponse](core.OperationSpec{
		Method: "alipay.trade.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	query, err := core.NewJSONOperation[QueryRequest, QueryResponse](core.OperationSpec{
		Method: "alipay.trade.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	precreate, err := core.NewJSONOperation[PreCreateRequest, PreCreateResponse](core.OperationSpec{
		Method: "alipay.trade.precreate", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/precreate",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	close, err := core.NewJSONOperation[CloseRequest, CloseResponse](core.OperationSpec{
		Method: "alipay.trade.close", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/close",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:  executor,
		pay:       pay,
		create:    create,
		cancel:    cancel,
		query:     query,
		precreate: precreate,
		close:     close,
	}, nil
}

// NewPaymentProducts 基于凭证构建页面/唤端产物构建入口。
// 构造过程不进行网络 I/O；配置不完整时返回 *core.ConfigError。
func NewPaymentProducts(credentials auth.Credentials, opts ...core.ClientOption) (*PaymentProducts, error) {
	builder, err := core.NewArtifactBuilder(credentials, opts...)
	if err != nil {
		return nil, err
	}
	return &PaymentProducts{builder: builder}, nil
}

// encodeArtifactParams 将 biz_content JSON 与 return_url 组装为
// 产物业务参数。
func encodeArtifactParams(biz any, returnURL *string) (url.Values, error) {
	raw, err := json.Marshal(biz)
	if err != nil {
		return nil, fmt.Errorf("alipay: encode artifact biz_content: %w", err)
	}
	params := url.Values{}
	params.Set("biz_content", string(raw))
	if returnURL != nil && *returnURL != "" {
		params.Set("return_url", *returnURL)
	}
	return params, nil
}

// validateGoodsDetail 校验商品明细的必填字段与金额格式。
func validateGoodsDetail(goods GoodsDetail, index int) error {
	if goods.GoodsID == "" || goods.GoodsName == "" || goods.Quantity <= 0 {
		return &core.ValidationError{Field: "goods_detail",
			Reason: fmt.Sprintf("goods_detail[%d] requires goods_id, goods_name and positive quantity", index)}
	}
	return tradevalidate.Amount(goods.Price, "goods_detail.price")
}

// Pay 调用 alipay.trade.pay 发起支付（V3 POST /v3/alipay/trade/pay）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/08c7f9f8_alipay.trade.pay
// 当面付条码/刷脸支付需传入 Scene 与 AuthCode；周期扣款场景经
// AgreementParams 传入协议号。写操作：SDK 不自动重试，超时后的
// 最终状态以查询接口为准。
func (s *Service) Pay(ctx context.Context, req PayRequest,
	opts ...core.CallOption) (*core.Response[PayResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.pay,
		core.Request[PayRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// Create 调用 alipay.trade.create 创建交易订单
// （V3 POST /v3/alipay/trade/create）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/bf4eae5e_alipay.trade.create
// 返回交易号仅表示订单创建成功；买家支付结果以异步通知或查询为准。
func (s *Service) Create(ctx context.Context, req CreateRequest,
	opts ...core.CallOption) (*core.Response[CreateResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.create,
		core.Request[CreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// Cancel 调用 alipay.trade.cancel 撤销交易
// （V3 POST /v3/alipay/trade/cancel）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/cd7d54d2_alipay.trade.cancel?scene=common&pathHash=cb83add6
// 未支付交易触发关闭，已支付交易触发退款；RetryFlag=Y 表示需要重试。
// 撤销是写操作，SDK 不自动重试。
func (s *Service) Cancel(ctx context.Context, req CancelRequest,
	opts ...core.CallOption) (*core.Response[CancelResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.cancel,
		core.Request[CancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// Query 调用 alipay.trade.query 查询交易状态（V3 POST /v3/alipay/trade/query）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/d33bf036_alipay.trade.query
// 请求成功不代表支付成功，业务结果以响应中的 TradeStatus 为准；
// 未生成真实交易时 TradeNo 与 BuyerLogonID 不返回。
func (s *Service) Query(ctx context.Context, req QueryRequest,
	opts ...core.CallOption) (*core.Response[QueryResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.query,
		core.Request[QueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PreCreate 调用 alipay.trade.precreate 生成当面付扫码支付二维码
// （V3 POST /v3/alipay/trade/precreate）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/fa0c2141_alipay.trade.precreate
// 二维码有效时间固定 2 小时；同一订单号不能重复申请二维码；
// qr_code 缺失时按业务响应处理，不代表预下单失败或成功。
func (s *Service) PreCreate(ctx context.Context, req PreCreateRequest,
	opts ...core.CallOption) (*core.Response[PreCreateResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.precreate,
		core.Request[PreCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// Close 调用 alipay.trade.close 关闭未付款交易
// （V3 POST /v3/alipay/trade/close）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/48ea518b_alipay.trade.close?scene=common&pathHash=2c6cdbea
// 仅"等待买家付款"状态的交易可关单；SDK 不依据本地状态自动发起关单。
func (s *Service) Close(ctx context.Context, req CloseRequest,
	opts ...core.CallOption) (*core.Response[CloseResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.close,
		core.Request[CloseRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
