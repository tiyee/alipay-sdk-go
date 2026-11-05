// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package fee

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

// FeeAdjustApply calls alipay.open.fee.adjust.apply (POST /v3/alipay/open/fee/adjust/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/fe77ed62_alipay.open.fee.adjust.apply
func (s *Service) FeeAdjustApply(ctx context.Context, req FeeAdjustApplyRequest, opts ...core.CallOption) (*core.Response[FeeAdjustApplyResponse], error) {
	return core.ExecuteJSON[FeeAdjustApplyRequest, FeeAdjustApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.fee.adjust.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/fee/adjust/apply",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FeeAdjustApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
