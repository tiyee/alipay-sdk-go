// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package individual

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

// MiniIndividualBusinessCertify calls alipay.open.mini.individual.business.certify (POST /v3/alipay/open/mini/individual/business/certify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/7ba79277_alipay.open.mini.individual.business.certify
func (s *Service) MiniIndividualBusinessCertify(ctx context.Context, req MiniIndividualBusinessCertifyRequest, opts ...core.CallOption) (*core.Response[MiniIndividualBusinessCertifyResponse], error) {
	return core.ExecuteJSON[MiniIndividualBusinessCertifyRequest, MiniIndividualBusinessCertifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.individual.business.certify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/individual/business/certify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniIndividualBusinessCertifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
