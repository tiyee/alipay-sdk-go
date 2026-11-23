// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package violation

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

// ViolationViolationDetailQuery calls alipay.open.violation.violationdetail.query (GET /v3/alipay/open/violation/violationdetail/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/cd81d87b_alipay.open.violation.violationdetail.query?scene=common&pathHash=e3dfda29
func (s *Service) ViolationViolationDetailQuery(ctx context.Context, req ViolationViolationDetailQueryRequest, opts ...core.CallOption) (*core.Response[ViolationViolationDetailQueryResponse], error) {
	return core.ExecuteJSON[ViolationViolationDetailQueryRequest, ViolationViolationDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.violation.violationdetail.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/violation/violationdetail/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ViolationViolationDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ViolationViolationEventBatchQuery calls alipay.open.violation.violationevent.batchquery (POST /v3/alipay/open/violation/violationevent/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/529a02eb_alipay.open.violation.violationevent.batchquery?scene=common&pathHash=67cab0ec
func (s *Service) ViolationViolationEventBatchQuery(ctx context.Context, req ViolationViolationEventBatchQueryRequest, opts ...core.CallOption) (*core.Response[ViolationViolationEventBatchQueryResponse], error) {
	return core.ExecuteJSON[ViolationViolationEventBatchQueryRequest, ViolationViolationEventBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.violation.violationevent.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/violation/violationevent/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ViolationViolationEventBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
