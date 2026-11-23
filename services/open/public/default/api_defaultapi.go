// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package defaultapi

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

// PublicDefaultExtensionCreate calls alipay.open.public.default.extension.create (POST /v3/alipay/open/public/default/extension/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvst
func (s *Service) PublicDefaultExtensionCreate(ctx context.Context, req PublicDefaultExtensionCreateRequest, opts ...core.CallOption) (*core.Response[PublicDefaultExtensionCreateResponse], error) {
	return core.ExecuteJSON[PublicDefaultExtensionCreateRequest, PublicDefaultExtensionCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.default.extension.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/default/extension/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicDefaultExtensionCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
