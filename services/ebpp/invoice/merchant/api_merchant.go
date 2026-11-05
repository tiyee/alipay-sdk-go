// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

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

// InvoiceMerchantEnterStatusQuery calls alipay.ebpp.invoice.merchant.enterstatus.query (GET /v3/alipay/ebpp/invoice/merchant/enterstatus/query).
// Official V3 contract: https://opendocs.alipay.com/apis/06x4dr
func (s *Service) InvoiceMerchantEnterStatusQuery(ctx context.Context, req InvoiceMerchantEnterStatusQueryRequest, opts ...core.CallOption) (*core.Response[InvoiceMerchantEnterStatusQueryResponse], error) {
	return core.ExecuteJSON[InvoiceMerchantEnterStatusQueryRequest, InvoiceMerchantEnterStatusQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.merchant.enterstatus.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/invoice/merchant/enterstatus/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceMerchantEnterStatusQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
