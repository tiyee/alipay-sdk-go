// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package voucherpackage

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

// ActivityVoucherPackageBatchQuery calls alipay.marketing.activity.voucherpackage.batchquery (POST /v3/alipay/marketing/activity/voucherpackage/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05velf
func (s *Service) ActivityVoucherPackageBatchQuery(ctx context.Context, req ActivityVoucherPackageBatchQueryRequest, opts ...core.CallOption) (*core.Response[ActivityVoucherPackageBatchQueryResponse], error) {
	return core.ExecuteJSON[ActivityVoucherPackageBatchQueryRequest, ActivityVoucherPackageBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.voucherpackage.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/voucherpackage/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityVoucherPackageBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityVoucherPackageConsult calls alipay.marketing.activity.voucherpackage.consult (POST /v3/alipay/marketing/activity/voucherpackage/consult).
// Official V3 contract: https://opendocs.alipay.com/apis/05vc1l
func (s *Service) ActivityVoucherPackageConsult(ctx context.Context, req ActivityVoucherPackageConsultRequest, opts ...core.CallOption) (*core.Response[ActivityVoucherPackageConsultResponse], error) {
	return core.ExecuteJSON[ActivityVoucherPackageConsultRequest, ActivityVoucherPackageConsultResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.voucherpackage.consult", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/voucherpackage/consult",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityVoucherPackageConsultRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityVoucherPackageQuery calls alipay.marketing.activity.voucherpackage.query (GET /v3/alipay/marketing/activity/voucherpackage/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05v9l4
func (s *Service) ActivityVoucherPackageQuery(ctx context.Context, req ActivityVoucherPackageQueryRequest, opts ...core.CallOption) (*core.Response[ActivityVoucherPackageQueryResponse], error) {
	return core.ExecuteJSON[ActivityVoucherPackageQueryRequest, ActivityVoucherPackageQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.voucherpackage.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/activity/voucherpackage/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityVoucherPackageQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
