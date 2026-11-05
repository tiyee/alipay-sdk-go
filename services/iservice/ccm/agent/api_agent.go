// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package agent

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

// CcmAgentCreate calls alipay.iservice.ccm.agent.create (POST /v3/alipay/iservice/ccm/agent/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/d09eb15f_alipay.iservice.ccm.agent.create
func (s *Service) CcmAgentCreate(ctx context.Context, req CcmAgentCreateRequest, opts ...core.CallOption) (*core.Response[CcmAgentCreateResponse], error) {
	return core.ExecuteJSON[CcmAgentCreateRequest, CcmAgentCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.agent.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/agent/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmAgentCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmAgentDelete calls alipay.iservice.ccm.agent.delete (DELETE /v3/alipay/iservice/ccm/agent/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/dc156fee_alipay.iservice.ccm.agent.delete
func (s *Service) CcmAgentDelete(ctx context.Context, req CcmAgentDeleteRequest, opts ...core.CallOption) (*core.Response[CcmAgentDeleteResponse], error) {
	return core.ExecuteJSON[CcmAgentDeleteRequest, CcmAgentDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.agent.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/iservice/ccm/agent/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmAgentDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmAgentGet calls alipay.iservice.ccm.agent.get (GET /v3/alipay/iservice/ccm/agent/get).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4a14056c_alipay.iservice.ccm.agent.get
func (s *Service) CcmAgentGet(ctx context.Context, req CcmAgentGetRequest, opts ...core.CallOption) (*core.Response[CcmAgentGetResponse], error) {
	return core.ExecuteJSON[CcmAgentGetRequest, CcmAgentGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.agent.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/agent/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmAgentGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmAgentModify calls alipay.iservice.ccm.agent.modify (POST /v3/alipay/iservice/ccm/agent/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e5964f19_alipay.iservice.ccm.agent.modify
func (s *Service) CcmAgentModify(ctx context.Context, req CcmAgentModifyRequest, opts ...core.CallOption) (*core.Response[CcmAgentModifyResponse], error) {
	return core.ExecuteJSON[CcmAgentModifyRequest, CcmAgentModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.agent.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/agent/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmAgentModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmAgentQuery calls alipay.iservice.ccm.agent.query (POST /v3/alipay/iservice/ccm/agent/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/992aed93_alipay.iservice.ccm.agent.query
func (s *Service) CcmAgentQuery(ctx context.Context, req CcmAgentQueryRequest, opts ...core.CallOption) (*core.Response[CcmAgentQueryResponse], error) {
	return core.ExecuteJSON[CcmAgentQueryRequest, CcmAgentQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.agent.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/agent/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmAgentQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmAgentScheduleLogQuery calls alipay.iservice.ccm.agent.schedulelog.query (POST /v3/alipay/iservice/ccm/agent/schedulelog/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/be76c956_alipay.iservice.ccm.agent.schedulelog.query
func (s *Service) CcmAgentScheduleLogQuery(ctx context.Context, req CcmAgentScheduleLogQueryRequest, opts ...core.CallOption) (*core.Response[CcmAgentScheduleLogQueryResponse], error) {
	return core.ExecuteJSON[CcmAgentScheduleLogQueryRequest, CcmAgentScheduleLogQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.agent.schedulelog.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/agent/schedulelog/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmAgentScheduleLogQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
