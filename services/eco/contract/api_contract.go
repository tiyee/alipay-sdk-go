// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package contract

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

// ContractSignFlowsCreate calls alipay.eco.contract.signflows.create (POST /v3/alipay/eco/contract/signflows/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/0681tj
func (s *Service) ContractSignFlowsCreate(ctx context.Context, req ContractSignFlowsCreateRequest, opts ...core.CallOption) (*core.Response[ContractSignFlowsCreateResponse], error) {
	return core.ExecuteJSON[ContractSignFlowsCreateRequest, ContractSignFlowsCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.contract.signflows.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/contract/signflows/create",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ContractSignFlowsCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
