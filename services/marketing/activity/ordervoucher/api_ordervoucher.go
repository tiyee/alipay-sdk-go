// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ordervoucher

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

// ActivityOrderVoucherAppend calls alipay.marketing.activity.ordervoucher.append (PATCH /v3/alipay/marketing/ordervoucher/activity/{activity_id}/append).
// Official V3 contract: https://opendocs.alipay.com/apis/05vjnq
func (s *Service) ActivityOrderVoucherAppend(ctx context.Context, req ActivityOrderVoucherAppendRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherAppendResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherAppendRequest, ActivityOrderVoucherAppendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.append", HTTPMethod: core.HTTPMethodPatch, Path: "/v3/alipay/marketing/ordervoucher/activity/{activity_id}/append",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherAppendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityOrderVoucherAssociate calls alipay.marketing.activity.ordervoucher.associate (POST /v3/alipay/marketing/activity/ordervoucher/associate).
// Official V3 contract: https://opendocs.alipay.com/apis/05v8k7
func (s *Service) ActivityOrderVoucherAssociate(ctx context.Context, req ActivityOrderVoucherAssociateRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherAssociateResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherAssociateRequest, ActivityOrderVoucherAssociateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.associate", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/ordervoucher/associate",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherAssociateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityOrderVoucherCodeCount calls alipay.marketing.activity.ordervoucher.codecount (GET /v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/codecount).
// Official V3 contract: https://opendocs.alipay.com/open-v3/1b98a84b_alipay.marketing.activity.ordervoucher.codecount
func (s *Service) ActivityOrderVoucherCodeCount(ctx context.Context, req ActivityOrderVoucherCodeCountRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherCodeCountResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherCodeCountRequest, ActivityOrderVoucherCodeCountResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.codecount", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/codecount",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherCodeCountRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityOrderVoucherCodeDeposit calls alipay.marketing.activity.ordervoucher.codedeposit (POST /v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/codedeposit).
// Official V3 contract: https://opendocs.alipay.com/apis/05u431
func (s *Service) ActivityOrderVoucherCodeDeposit(ctx context.Context, req ActivityOrderVoucherCodeDepositRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherCodeDepositResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherCodeDepositRequest, ActivityOrderVoucherCodeDepositResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.codedeposit", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/codedeposit",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherCodeDepositRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityOrderVoucherCreate calls alipay.marketing.activity.ordervoucher.create (POST /v3/alipay/marketing/ordervoucher/activity).
// Official V3 contract: https://opendocs.alipay.com/apis/05qgfh
func (s *Service) ActivityOrderVoucherCreate(ctx context.Context, req ActivityOrderVoucherCreateRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherCreateResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherCreateRequest, ActivityOrderVoucherCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/ordervoucher/activity",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityOrderVoucherDisassociate calls alipay.marketing.activity.ordervoucher.disassociate (POST /v3/alipay/marketing/activity/ordervoucher/disassociate).
// Official V3 contract: https://opendocs.alipay.com/apis/05vluj
func (s *Service) ActivityOrderVoucherDisassociate(ctx context.Context, req ActivityOrderVoucherDisassociateRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherDisassociateResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherDisassociateRequest, ActivityOrderVoucherDisassociateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.disassociate", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/ordervoucher/disassociate",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherDisassociateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityOrderVoucherModify calls alipay.marketing.activity.ordervoucher.modify (PATCH /v3/alipay/marketing/ordervoucher/activity/{activity_id}).
// Official V3 contract: https://opendocs.alipay.com/apis/05vfpx
func (s *Service) ActivityOrderVoucherModify(ctx context.Context, req ActivityOrderVoucherModifyRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherModifyResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherModifyRequest, ActivityOrderVoucherModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.modify", HTTPMethod: core.HTTPMethodPatch, Path: "/v3/alipay/marketing/ordervoucher/activity/{activity_id}",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityOrderVoucherQuery calls alipay.marketing.activity.ordervoucher.query (GET /v3/alipay/marketing/activity/ordervoucher/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/ab283f3f_alipay.marketing.activity.ordervoucher.query
func (s *Service) ActivityOrderVoucherQuery(ctx context.Context, req ActivityOrderVoucherQueryRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherQueryResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherQueryRequest, ActivityOrderVoucherQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/activity/ordervoucher/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityOrderVoucherRefund calls alipay.marketing.activity.ordervoucher.refund (POST /v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/{voucher_code}/refund).
// Official V3 contract: https://opendocs.alipay.com/apis/0636fj
func (s *Service) ActivityOrderVoucherRefund(ctx context.Context, req ActivityOrderVoucherRefundRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherRefundResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherRefundRequest, ActivityOrderVoucherRefundResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.refund", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/{voucher_code}/refund",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherRefundRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityOrderVoucherStop calls alipay.marketing.activity.ordervoucher.stop (PATCH /v3/alipay/marketing/ordervoucher/activity/{activity_id}/stop).
// Official V3 contract: https://opendocs.alipay.com/apis/0630yz
func (s *Service) ActivityOrderVoucherStop(ctx context.Context, req ActivityOrderVoucherStopRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherStopResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherStopRequest, ActivityOrderVoucherStopResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.stop", HTTPMethod: core.HTTPMethodPatch, Path: "/v3/alipay/marketing/ordervoucher/activity/{activity_id}/stop",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherStopRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityOrderVoucherUse calls alipay.marketing.activity.ordervoucher.use (POST /v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/{voucher_code}/use).
// Official V3 contract: https://opendocs.alipay.com/apis/0639qt
func (s *Service) ActivityOrderVoucherUse(ctx context.Context, req ActivityOrderVoucherUseRequest, opts ...core.CallOption) (*core.Response[ActivityOrderVoucherUseResponse], error) {
	return core.ExecuteJSON[ActivityOrderVoucherUseRequest, ActivityOrderVoucherUseResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.ordervoucher.use", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/ordervoucher/activity/{activity_id}/voucher/{voucher_code}/use",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityOrderVoucherUseRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
