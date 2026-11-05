// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package sync

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

// InvoiceSyncSimpleSend calls alipay.ebpp.invoice.sync.simple.send (POST /v3/alipay/ebpp/invoice/sync/simple/send).
// Official V3 contract: https://opendocs.alipay.com/apis/05kxx8
func (s *Service) InvoiceSyncSimpleSend(ctx context.Context, req InvoiceSyncSimpleSendRequest, opts ...core.CallOption) (*core.Response[InvoiceSyncSimpleSendResponse], error) {
	return core.ExecuteJSON[InvoiceSyncSimpleSendRequest, InvoiceSyncSimpleSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.sync.simple.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/sync/simple/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceSyncSimpleSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
