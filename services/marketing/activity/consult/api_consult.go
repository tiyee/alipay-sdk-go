// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package consult

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

// ActivityConsult calls alipay.marketing.activity.consult (POST /v3/alipay/marketing/activity/consult).
// Official V3 contract: https://opendocs.alipay.com/open-v3/3c187ec5_alipay.marketing.activity.consult
func (s *Service) ActivityConsult(ctx context.Context, req ActivityConsultRequest, opts ...core.CallOption) (*core.Response[ActivityConsultResponse], error) {
	return core.ExecuteJSON[ActivityConsultRequest, ActivityConsultResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.consult", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/consult",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityConsultRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
