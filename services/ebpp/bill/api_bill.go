// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package bill

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

// BillAdd calls alipay.ebpp.bill.add (POST /v3/alipay/ebpp/bill/add).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b9117832_alipay.ebpp.bill.add
func (s *Service) BillAdd(ctx context.Context, req BillAddRequest, opts ...core.CallOption) (*core.Response[BillAddResponse], error) {
	return core.ExecuteJSON[BillAddRequest, BillAddResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.bill.add", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/bill/add",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BillAddRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// BillGet calls alipay.ebpp.bill.get (GET /v3/alipay/ebpp/bill/get).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b6ddabc9_alipay.ebpp.bill.get
func (s *Service) BillGet(ctx context.Context, req BillGetRequest, opts ...core.CallOption) (*core.Response[BillGetResponse], error) {
	return core.ExecuteJSON[BillGetRequest, BillGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.bill.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/bill/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BillGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
