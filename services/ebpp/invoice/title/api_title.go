// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package title

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

// InvoiceTitleDynamicGet calls alipay.ebpp.invoice.title.dynamic.get (GET /v3/alipay/ebpp/invoice/title/dynamic/get).
// Official V3 contract: https://opendocs.alipay.com/apis/05v8k3
func (s *Service) InvoiceTitleDynamicGet(ctx context.Context, req InvoiceTitleDynamicGetRequest, opts ...core.CallOption) (*core.Response[InvoiceTitleDynamicGetResponse], error) {
	return core.ExecuteJSON[InvoiceTitleDynamicGetRequest, InvoiceTitleDynamicGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.title.dynamic.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/title/dynamic/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceTitleDynamicGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceTitleListGet calls alipay.ebpp.invoice.title.list.get (GET /v3/alipay/ebpp/invoice/title/list/get).
// Official V3 contract: https://opendocs.alipay.com/apis/06864n
func (s *Service) InvoiceTitleListGet(ctx context.Context, req InvoiceTitleListGetRequest, opts ...core.CallOption) (*core.Response[InvoiceTitleListGetResponse], error) {
	return core.ExecuteJSON[InvoiceTitleListGetRequest, InvoiceTitleListGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.title.list.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/title/list/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceTitleListGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
