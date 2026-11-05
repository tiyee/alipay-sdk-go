package accountbook

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

// AccountBookCreate calls alipay.fund.accountbook.create (POST /v3/alipay/fund/accountbook/create).
// Official V3 contract: https://opendocs.alipay.com/apis/0cjl4u
func (s *Service) AccountBookCreate(ctx context.Context, req FundAccountBookCreateRequest, opts ...core.CallOption) (*core.Response[FundAccountBookCreateResponse], error) {
	return core.ExecuteJSON[FundAccountBookCreateRequest, FundAccountBookCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.accountbook.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/accountbook/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAccountBookCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AccountBookNotifyQuery calls alipay.fund.accountbook.notify.query (GET /v3/alipay/fund/accountbook/notify/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/495b0416_alipay.fund.accountbook.notify.query
func (s *Service) AccountBookNotifyQuery(ctx context.Context, req FundAccountBookNotifyQueryRequest, opts ...core.CallOption) (*core.Response[FundAccountBookNotifyQueryResponse], error) {
	return core.ExecuteJSON[FundAccountBookNotifyQueryRequest, FundAccountBookNotifyQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.accountbook.notify.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/accountbook/notify/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAccountBookNotifyQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AccountBookNotifySubscribe calls alipay.fund.accountbook.notify.subscribe (POST /v3/alipay/fund/accountbook/notify/subscribe).
// Official V3 contract: https://opendocs.alipay.com/open-v3/37df5c87_alipay.fund.accountbook.notify.subscribe
func (s *Service) AccountBookNotifySubscribe(ctx context.Context, req FundAccountBookNotifySubscribeRequest, opts ...core.CallOption) (*core.Response[FundAccountBookNotifySubscribeResponse], error) {
	return core.ExecuteJSON[FundAccountBookNotifySubscribeRequest, FundAccountBookNotifySubscribeResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.accountbook.notify.subscribe", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/accountbook/notify/subscribe",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAccountBookNotifySubscribeRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AccountBookNotifyUnsubscribe calls alipay.fund.accountbook.notify.unsubscribe (POST /v3/alipay/fund/accountbook/notify/unsubscribe).
// Official V3 contract: https://opendocs.alipay.com/open-v3/1923216a_alipay.fund.accountbook.notify.unsubscribe
func (s *Service) AccountBookNotifyUnsubscribe(ctx context.Context, req FundAccountBookNotifyUnsubscribeRequest, opts ...core.CallOption) (*core.Response[FundAccountBookNotifyUnsubscribeResponse], error) {
	return core.ExecuteJSON[FundAccountBookNotifyUnsubscribeRequest, FundAccountBookNotifyUnsubscribeResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.accountbook.notify.unsubscribe", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/accountbook/notify/unsubscribe",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAccountBookNotifyUnsubscribeRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AccountBookQuery calls alipay.fund.accountbook.query (GET /v3/alipay/fund/accountbook/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/f072b3ba_alipay.fund.accountbook.query
func (s *Service) AccountBookQuery(ctx context.Context, req FundAccountBookQueryRequest, opts ...core.CallOption) (*core.Response[FundAccountBookQueryResponse], error) {
	return core.ExecuteJSON[FundAccountBookQueryRequest, FundAccountBookQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.accountbook.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/accountbook/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAccountBookQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
