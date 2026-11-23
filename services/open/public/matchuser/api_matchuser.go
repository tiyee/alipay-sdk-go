// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package matchuser

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

// PublicMatchUserLabelCreate calls alipay.open.public.matchuser.label.create (POST /v3/alipay/open/public/matchuser/label/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wycu
func (s *Service) PublicMatchUserLabelCreate(ctx context.Context, req PublicMatchUserLabelCreateRequest, opts ...core.CallOption) (*core.Response[PublicMatchUserLabelCreateResponse], error) {
	return core.ExecuteJSON[PublicMatchUserLabelCreateRequest, PublicMatchUserLabelCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.matchuser.label.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/matchuser/label/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMatchUserLabelCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicMatchUserLabelDelete calls alipay.open.public.matchuser.label.delete (POST /v3/alipay/open/public/matchuser/label/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05wz8y
func (s *Service) PublicMatchUserLabelDelete(ctx context.Context, req PublicMatchUserLabelDeleteRequest, opts ...core.CallOption) (*core.Response[PublicMatchUserLabelDeleteResponse], error) {
	return core.ExecuteJSON[PublicMatchUserLabelDeleteRequest, PublicMatchUserLabelDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.matchuser.label.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/matchuser/label/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicMatchUserLabelDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
