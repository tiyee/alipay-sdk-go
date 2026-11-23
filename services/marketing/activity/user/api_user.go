// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package user

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

// ActivityUserBatchQueryVoucher calls alipay.marketing.activity.user.batchqueryvoucher (GET /v3/alipay/marketing/activity/ordervoucher/user/batchqueryvoucher).
// Official V3 contract: https://opendocs.alipay.com/open-v3/63264abf_alipay.marketing.activity.user.batchqueryvoucher
func (s *Service) ActivityUserBatchQueryVoucher(ctx context.Context, req ActivityUserBatchQueryVoucherRequest, opts ...core.CallOption) (*core.Response[ActivityUserBatchQueryVoucherResponse], error) {
	return core.ExecuteJSON[ActivityUserBatchQueryVoucherRequest, ActivityUserBatchQueryVoucherResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.user.batchqueryvoucher", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/activity/ordervoucher/user/batchqueryvoucher",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityUserBatchQueryVoucherRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityUserQueryVoucher calls alipay.marketing.activity.user.queryvoucher (GET /v3/alipay/marketing/activity/ordervoucher/user/voucher).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e75be5aa_alipay.marketing.activity.user.queryvoucher
func (s *Service) ActivityUserQueryVoucher(ctx context.Context, req ActivityUserQueryVoucherRequest, opts ...core.CallOption) (*core.Response[ActivityUserQueryVoucherResponse], error) {
	return core.ExecuteJSON[ActivityUserQueryVoucherRequest, ActivityUserQueryVoucherResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.user.queryvoucher", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/activity/ordervoucher/user/voucher",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityUserQueryVoucherRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
