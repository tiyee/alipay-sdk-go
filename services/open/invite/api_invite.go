// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package invite

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

// InviteOrderQuery calls alipay.open.invite.order.query (GET /v3/alipay/open/invite/order/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a5f3aaab_alipay.open.invite.order.query
func (s *Service) InviteOrderQuery(ctx context.Context, req InviteOrderQueryRequest, opts ...core.CallOption) (*core.Response[InviteOrderQueryResponse], error) {
	return core.ExecuteJSON[InviteOrderQueryRequest, InviteOrderQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.invite.order.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/invite/order/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InviteOrderQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
