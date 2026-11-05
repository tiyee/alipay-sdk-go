// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package shortlink

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

// PublicShortLinkCreate calls alipay.open.public.shortlink.create (POST /v3/alipay/open/public/shortlink/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05xin2
func (s *Service) PublicShortLinkCreate(ctx context.Context, req PublicShortLinkCreateRequest, opts ...core.CallOption) (*core.Response[PublicShortLinkCreateResponse], error) {
	return core.ExecuteJSON[PublicShortLinkCreateRequest, PublicShortLinkCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.shortlink.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/shortlink/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicShortLinkCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
