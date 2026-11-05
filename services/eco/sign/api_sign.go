// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package sign

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

// SignFlowCancel calls alipay.eco.sign.flow.cancel (POST /v3/alipay/eco/sign/flow/cancel).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wvs6
func (s *Service) SignFlowCancel(ctx context.Context, req SignFlowCancelRequest, opts ...core.CallOption) (*core.Response[SignFlowCancelResponse], error) {
	return core.ExecuteJSON[SignFlowCancelRequest, SignFlowCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.sign.flow.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/sign/flow/cancel",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SignFlowCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SignFlowCreate calls alipay.eco.sign.flow.create (POST /v3/alipay/eco/sign/flow/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xsf2
func (s *Service) SignFlowCreate(ctx context.Context, req SignFlowCreateRequest, opts ...core.CallOption) (*core.Response[SignFlowCreateResponse], error) {
	return core.ExecuteJSON[SignFlowCreateRequest, SignFlowCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.sign.flow.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/sign/flow/create",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SignFlowCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SignFlowFinish calls alipay.eco.sign.flow.finish (POST /v3/alipay/eco/sign/flow/finish).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wuhl
func (s *Service) SignFlowFinish(ctx context.Context, req SignFlowFinishRequest, opts ...core.CallOption) (*core.Response[SignFlowFinishResponse], error) {
	return core.ExecuteJSON[SignFlowFinishRequest, SignFlowFinishResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.sign.flow.finish", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/sign/flow/finish",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SignFlowFinishRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SignFlowQuery calls alipay.eco.sign.flow.query (GET /v3/alipay/eco/sign/flow/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wvxl
func (s *Service) SignFlowQuery(ctx context.Context, req SignFlowQueryRequest, opts ...core.CallOption) (*core.Response[SignFlowQueryResponse], error) {
	return core.ExecuteJSON[SignFlowQueryRequest, SignFlowQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.sign.flow.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/eco/sign/flow/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SignFlowQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
