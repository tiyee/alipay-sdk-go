// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package doc

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

// DocTemplateCreate calls alipay.eco.doc.template.create (POST /v3/alipay/eco/doc/template/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05whak
func (s *Service) DocTemplateCreate(ctx context.Context, req DocTemplateCreateRequest, opts ...core.CallOption) (*core.Response[DocTemplateCreateResponse], error) {
	return core.ExecuteJSON[DocTemplateCreateRequest, DocTemplateCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.doc.template.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/doc/template/create",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DocTemplateCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
