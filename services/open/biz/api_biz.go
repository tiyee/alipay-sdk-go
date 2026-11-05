// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package biz

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

// BizCreate calls alipay.open.biz.create (POST /v3/alipay/open/biz/create).
// Official V3 contract: https://opendocs.alipay.com/apis/0i01un
func (s *Service) BizCreate(ctx context.Context, req BizCreateRequest, opts ...core.CallOption) (*core.Response[BizCreateResponse], error) {
	return core.ExecuteJSON[BizCreateRequest, BizCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.biz.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/biz/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BizCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
