// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package life

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service executes one alipay V3 service group.
type Service struct{ executor core.Executor }

// NewService constructs the service with a shared V3 executor.
// NewService constructs the service with a shared V3 executor.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

// PublicLifeAboardApply calls alipay.open.public.life.aboard.apply (POST /v3/alipay/open/public/life/aboard/apply).
// Official V3 contract: https://opendocs.alipay.com/apis/05wyga
func (s *Service) PublicLifeAboardApply(ctx context.Context, req PublicLifeAboardApplyRequest, opts ...core.CallOption) (*core.Response[PublicLifeAboardApplyResponse], error) {
	return core.ExecuteJSON[PublicLifeAboardApplyRequest, PublicLifeAboardApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.life.aboard.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/life/aboard/apply",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLifeAboardApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicLifeAgentCreateQuery calls alipay.open.public.life.agentcreate.query (GET /v3/alipay/open/public/life/agentcreate/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuo2
func (s *Service) PublicLifeAgentCreateQuery(ctx context.Context, req PublicLifeAgentCreateQueryRequest, opts ...core.CallOption) (*core.Response[PublicLifeAgentCreateQueryResponse], error) {
	return core.ExecuteJSON[PublicLifeAgentCreateQueryRequest, PublicLifeAgentCreateQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.life.agentcreate.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/public/life/agentcreate/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLifeAgentCreateQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicLifeDebarkApply calls alipay.open.public.life.debark.apply (POST /v3/alipay/open/public/life/debark/apply).
// Official V3 contract: https://opendocs.alipay.com/apis/05wnsi
func (s *Service) PublicLifeDebarkApply(ctx context.Context, req PublicLifeDebarkApplyRequest, opts ...core.CallOption) (*core.Response[PublicLifeDebarkApplyResponse], error) {
	return core.ExecuteJSON[PublicLifeDebarkApplyRequest, PublicLifeDebarkApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.life.debark.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/life/debark/apply",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLifeDebarkApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicLifeLabelBatchQuery calls alipay.open.public.life.label.batchquery (POST /v3/alipay/open/public/life/label/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wui2
func (s *Service) PublicLifeLabelBatchQuery(ctx context.Context, req PublicLifeLabelBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicLifeLabelBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicLifeLabelBatchQueryRequest, PublicLifeLabelBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.life.label.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/life/label/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLifeLabelBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicLifeLabelCreate calls alipay.open.public.life.label.create (POST /v3/alipay/open/public/life/label/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wh9j
func (s *Service) PublicLifeLabelCreate(ctx context.Context, req PublicLifeLabelCreateRequest, opts ...core.CallOption) (*core.Response[PublicLifeLabelCreateResponse], error) {
	return core.ExecuteJSON[PublicLifeLabelCreateRequest, PublicLifeLabelCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.life.label.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/life/label/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLifeLabelCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicLifeLabelDelete calls alipay.open.public.life.label.delete (DELETE /v3/alipay/open/public/life/label/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvsg
func (s *Service) PublicLifeLabelDelete(ctx context.Context, req PublicLifeLabelDeleteRequest, opts ...core.CallOption) (*core.Response[PublicLifeLabelDeleteResponse], error) {
	return core.ExecuteJSON[PublicLifeLabelDeleteRequest, PublicLifeLabelDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.life.label.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/public/life/label/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLifeLabelDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicLifeLabelModify calls alipay.open.public.life.label.modify (POST /v3/alipay/open/public/life/label/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05xhmq
func (s *Service) PublicLifeLabelModify(ctx context.Context, req PublicLifeLabelModifyRequest, opts ...core.CallOption) (*core.Response[PublicLifeLabelModifyResponse], error) {
	return core.ExecuteJSON[PublicLifeLabelModifyRequest, PublicLifeLabelModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.life.label.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/life/label/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLifeLabelModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicLifeMsgRecall calls alipay.open.public.life.msg.recall (POST /v3/alipay/open/public/life/msg/recall).
// Official V3 contract: https://opendocs.alipay.com/apis/05wieh
func (s *Service) PublicLifeMsgRecall(ctx context.Context, req PublicLifeMsgRecallRequest, opts ...core.CallOption) (*core.Response[PublicLifeMsgRecallResponse], error) {
	return core.ExecuteJSON[PublicLifeMsgRecallRequest, PublicLifeMsgRecallResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.life.msg.recall", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/life/msg/recall",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLifeMsgRecallRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicLifeMsgSend calls alipay.open.public.life.msg.send (POST /v3/alipay/open/public/life/msg/send).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wulv
func (s *Service) PublicLifeMsgSend(ctx context.Context, req PublicLifeMsgSendRequest, opts ...core.CallOption) (*core.Response[PublicLifeMsgSendResponse], error) {
	return core.ExecuteJSON[PublicLifeMsgSendRequest, PublicLifeMsgSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.life.msg.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/life/msg/send",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicLifeMsgSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
