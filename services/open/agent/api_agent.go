package agent

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

// AgentCancel calls alipay.open.agent.cancel (POST /v3/alipay/open/agent/cancel).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8e66406f_alipay.open.agent.cancel
func (s *Service) AgentCancel(ctx context.Context, req AgentCancelRequest, opts ...core.CallOption) (*core.Response[AgentCancelResponse], error) {
	return core.ExecuteJSON[AgentCancelRequest, AgentCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/agent/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgentConfirm calls alipay.open.agent.confirm (POST /v3/alipay/open/agent/confirm).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e853dfa1_alipay.open.agent.confirm
func (s *Service) AgentConfirm(ctx context.Context, req AgentConfirmRequest, opts ...core.CallOption) (*core.Response[AgentConfirmResponse], error) {
	return core.ExecuteJSON[AgentConfirmRequest, AgentConfirmResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.confirm", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/agent/confirm",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentConfirmRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgentCreate calls alipay.open.agent.create (POST /v3/alipay/open/agent/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b4c46445_alipay.open.agent.create
func (s *Service) AgentCreate(ctx context.Context, req AgentCreateRequest, opts ...core.CallOption) (*core.Response[AgentCreateResponse], error) {
	return core.ExecuteJSON[AgentCreateRequest, AgentCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/agent/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgentMiniCreate calls alipay.open.agent.mini.create (POST /v3/alipay/open/agent/mini/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e39a93e8_alipay.open.agent.mini.create
func (s *Service) AgentMiniCreate(ctx context.Context, req AgentMiniCreateRequest, opts ...core.CallOption) (*core.Response[AgentMiniCreateResponse], error) {
	return core.ExecuteJSON[AgentMiniCreateRequest, AgentMiniCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.mini.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/agent/mini/create",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentMiniCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgentOrderQuery calls alipay.open.agent.order.query (GET /v3/alipay/open/agent/order/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/ceae4d66_alipay.open.agent.order.query
func (s *Service) AgentOrderQuery(ctx context.Context, req AgentOrderQueryRequest, opts ...core.CallOption) (*core.Response[AgentOrderQueryResponse], error) {
	return core.ExecuteJSON[AgentOrderQueryRequest, AgentOrderQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.order.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/agent/order/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentOrderQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
