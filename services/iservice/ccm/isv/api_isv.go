// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package isv

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

// CcmISVInitialize calls alipay.iservice.ccm.isv.initialize (POST /v3/alipay/iservice/ccm/isv/initialize).
// Official V3 contract: https://opendocs.alipay.com/open-v3/eb2e1712_alipay.iservice.ccm.isv.initialize?scene=common&pathHash=da5a3102
func (s *Service) CcmISVInitialize(ctx context.Context, req CcmISVInitializeRequest, opts ...core.CallOption) (*core.Response[CcmISVInitializeResponse], error) {
	return core.ExecuteJSON[CcmISVInitializeRequest, CcmISVInitializeResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.isv.initialize", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/isv/initialize",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmISVInitializeRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
