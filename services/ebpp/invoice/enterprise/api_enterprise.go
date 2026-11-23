// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package enterprise

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

// InvoiceEnterpriseMerchantRelationCreate calls alipay.ebpp.invoice.enterprise.merchantrelation.create (POST /v3/alipay/ebpp/invoice/enterprise/merchantrelation/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05psjg
func (s *Service) InvoiceEnterpriseMerchantRelationCreate(ctx context.Context, req InvoiceEnterpriseMerchantRelationCreateRequest, opts ...core.CallOption) (*core.Response[InvoiceEnterpriseMerchantRelationCreateResponse], error) {
	return core.ExecuteJSON[InvoiceEnterpriseMerchantRelationCreateRequest, InvoiceEnterpriseMerchantRelationCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.enterprise.merchantrelation.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/enterprise/merchantrelation/create",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceEnterpriseMerchantRelationCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
