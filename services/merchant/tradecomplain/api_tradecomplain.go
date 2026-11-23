package tradecomplain

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

// TradeComplainBatchQuery calls alipay.merchant.tradecomplain.batchquery (POST /v3/alipay/merchant/tradecomplain/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/152a9215_alipay.merchant.tradecomplain.batchquery
func (s *Service) TradeComplainBatchQuery(ctx context.Context, req MerchantTradeComplainBatchQueryRequest, opts ...core.CallOption) (*core.Response[MerchantTradeComplainBatchQueryResponse], error) {
	return core.ExecuteJSON[MerchantTradeComplainBatchQueryRequest, MerchantTradeComplainBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.tradecomplain.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/tradecomplain/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantTradeComplainBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// TradeComplainFeedbackSubmit calls alipay.merchant.tradecomplain.feedback.submit (POST /v3/alipay/merchant/tradecomplain/feedback).
// Official V3 contract: https://opendocs.alipay.com/open-v3/da1eaf4f_alipay.merchant.tradecomplain.feedback.submit
func (s *Service) TradeComplainFeedbackSubmit(ctx context.Context, req MerchantTradeComplainFeedbackSubmitRequest, opts ...core.CallOption) (*core.Response[MerchantTradeComplainFeedbackSubmitResponse], error) {
	return core.ExecuteJSON[MerchantTradeComplainFeedbackSubmitRequest, MerchantTradeComplainFeedbackSubmitResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.tradecomplain.feedback.submit", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/tradecomplain/feedback",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantTradeComplainFeedbackSubmitRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// TradeComplainQuery calls alipay.merchant.tradecomplain.query (GET /v3/alipay/merchant/tradecomplain/{complain_event_id}).
// Official V3 contract: https://opendocs.alipay.com/open-v3/cecf0505_alipay.merchant.tradecomplain.query
func (s *Service) TradeComplainQuery(ctx context.Context, req MerchantTradeComplainQueryRequest, opts ...core.CallOption) (*core.Response[MerchantTradeComplainQueryResponse], error) {
	return core.ExecuteJSON[MerchantTradeComplainQueryRequest, MerchantTradeComplainQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.tradecomplain.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/merchant/tradecomplain/{complain_event_id}",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantTradeComplainQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// TradeComplainReplySubmit calls alipay.merchant.tradecomplain.reply.submit (POST /v3/alipay/merchant/tradecomplain/reply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/5851abbd_alipay.merchant.tradecomplain.reply.submit
func (s *Service) TradeComplainReplySubmit(ctx context.Context, req MerchantTradeComplainReplySubmitRequest, opts ...core.CallOption) (*core.Response[MerchantTradeComplainReplySubmitResponse], error) {
	return core.ExecuteJSON[MerchantTradeComplainReplySubmitRequest, MerchantTradeComplainReplySubmitResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.tradecomplain.reply.submit", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/tradecomplain/reply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantTradeComplainReplySubmitRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// TradeComplainSupplementSubmit calls alipay.merchant.tradecomplain.supplement.submit (POST /v3/alipay/merchant/tradecomplain/supplement).
// Official V3 contract: https://opendocs.alipay.com/open-v3/bb0a25ee_alipay.merchant.tradecomplain.supplement.submit
func (s *Service) TradeComplainSupplementSubmit(ctx context.Context, req MerchantTradeComplainSupplementSubmitRequest, opts ...core.CallOption) (*core.Response[MerchantTradeComplainSupplementSubmitResponse], error) {
	return core.ExecuteJSON[MerchantTradeComplainSupplementSubmitRequest, MerchantTradeComplainSupplementSubmitResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.tradecomplain.supplement.submit", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/tradecomplain/supplement",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantTradeComplainSupplementSubmitRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
