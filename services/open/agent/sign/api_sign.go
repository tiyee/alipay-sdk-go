package sign

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

// AgentCommonSign calls alipay.open.agent.common.sign (POST /v3/alipay/open/agent/common/sign).
// Official V3 contract: https://opendocs.alipay.com/open-v3/f0ab46ad_alipay.open.agent.common.sign
func (s *Service) AgentCommonSign(ctx context.Context, req AgentCommonSignRequest, opts ...core.CallOption) (*core.Response[AgentCommonSignResponse], error) {
	return core.ExecuteJSON[AgentCommonSignRequest, AgentCommonSignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.common.sign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/agent/common/sign",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentCommonSignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgentCommonSignConfirm calls alipay.open.agent.commonsign.confirm (POST /v3/alipay/open/agent/commonsign/confirm).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2f809b83_alipay.open.agent.commonsign.confirm
func (s *Service) AgentCommonSignConfirm(ctx context.Context, req AgentCommonSignConfirmRequest, opts ...core.CallOption) (*core.Response[AgentCommonSignConfirmResponse], error) {
	return core.ExecuteJSON[AgentCommonSignConfirmRequest, AgentCommonSignConfirmResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.commonsign.confirm", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/agent/commonsign/confirm",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentCommonSignConfirmRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgentFaceToFaceSign calls alipay.open.agent.facetoface.sign (POST /v3/alipay/open/agent/facetoface/sign).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e1eb4236_alipay.open.agent.facetoface.sign
func (s *Service) AgentFaceToFaceSign(ctx context.Context, req AgentFaceToFaceSignRequest, opts ...core.CallOption) (*core.Response[AgentFaceToFaceSignResponse], error) {
	return core.ExecuteJSON[AgentFaceToFaceSignRequest, AgentFaceToFaceSignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.facetoface.sign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/agent/facetoface/sign",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentFaceToFaceSignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgentMobilePaySign calls alipay.open.agent.mobilepay.sign (POST /v3/alipay/open/agent/mobilepay/sign).
// Official V3 contract: https://opendocs.alipay.com/open-v3/db21f6cb_alipay.open.agent.mobilepay.sign
func (s *Service) AgentMobilePaySign(ctx context.Context, req AgentMobilePaySignRequest, opts ...core.CallOption) (*core.Response[AgentMobilePaySignResponse], error) {
	return core.ExecuteJSON[AgentMobilePaySignRequest, AgentMobilePaySignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.mobilepay.sign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/agent/mobilepay/sign",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentMobilePaySignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgentOfflinePaymentSign calls alipay.open.agent.offlinepayment.sign (POST /v3/alipay/open/agent/offlinepayment/sign).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8b80b506_alipay.open.agent.offlinepayment.sign?scene=common&pathHash=255833d8
func (s *Service) AgentOfflinePaymentSign(ctx context.Context, req AgentOfflinePaymentSignRequest, opts ...core.CallOption) (*core.Response[AgentOfflinePaymentSignResponse], error) {
	return core.ExecuteJSON[AgentOfflinePaymentSignRequest, AgentOfflinePaymentSignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.offlinepayment.sign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/agent/offlinepayment/sign",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentOfflinePaymentSignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgentSignStatusQuery calls alipay.open.agent.signstatus.query (POST /v3/alipay/open/agent/signstatus/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/031e92b9_alipay.open.agent.signstatus.query
func (s *Service) AgentSignStatusQuery(ctx context.Context, req AgentSignStatusQueryRequest, opts ...core.CallOption) (*core.Response[AgentSignStatusQueryResponse], error) {
	return core.ExecuteJSON[AgentSignStatusQueryRequest, AgentSignStatusQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.agent.signstatus.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/agent/signstatus/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[AgentSignStatusQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
