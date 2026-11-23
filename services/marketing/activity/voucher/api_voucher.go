// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package voucher

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

// ActivityVoucherAppend calls alipay.marketing.activity.voucher.append (POST /v3/alipay/marketing/activity/voucher/append).
// Official V3 contract: https://opendocs.alipay.com/apis/05vsc4
func (s *Service) ActivityVoucherAppend(ctx context.Context, req ActivityVoucherAppendRequest, opts ...core.CallOption) (*core.Response[ActivityVoucherAppendResponse], error) {
	return core.ExecuteJSON[ActivityVoucherAppendRequest, ActivityVoucherAppendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.voucher.append", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/voucher/append",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityVoucherAppendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityVoucherCreate calls alipay.marketing.activity.voucher.create (POST /v3/alipay/marketing/activity/voucher/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05vb41
func (s *Service) ActivityVoucherCreate(ctx context.Context, req ActivityVoucherCreateRequest, opts ...core.CallOption) (*core.Response[ActivityVoucherCreateResponse], error) {
	return core.ExecuteJSON[ActivityVoucherCreateRequest, ActivityVoucherCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.voucher.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/voucher/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityVoucherCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityVoucherModify calls alipay.marketing.activity.voucher.modify (POST /v3/alipay/marketing/activity/voucher/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05v8k9
func (s *Service) ActivityVoucherModify(ctx context.Context, req ActivityVoucherModifyRequest, opts ...core.CallOption) (*core.Response[ActivityVoucherModifyResponse], error) {
	return core.ExecuteJSON[ActivityVoucherModifyRequest, ActivityVoucherModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.voucher.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/voucher/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityVoucherModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityVoucherPublish calls alipay.marketing.activity.voucher.publish (POST /v3/alipay/marketing/activity/voucher/publish).
// Official V3 contract: https://opendocs.alipay.com/apis/05vsc3
func (s *Service) ActivityVoucherPublish(ctx context.Context, req ActivityVoucherPublishRequest, opts ...core.CallOption) (*core.Response[ActivityVoucherPublishResponse], error) {
	return core.ExecuteJSON[ActivityVoucherPublishRequest, ActivityVoucherPublishResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.voucher.publish", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/voucher/publish",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityVoucherPublishRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityVoucherQuery calls alipay.marketing.activity.voucher.query (GET /v3/alipay/marketing/activity/voucher/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05veld
func (s *Service) ActivityVoucherQuery(ctx context.Context, req ActivityVoucherQueryRequest, opts ...core.CallOption) (*core.Response[ActivityVoucherQueryResponse], error) {
	return core.ExecuteJSON[ActivityVoucherQueryRequest, ActivityVoucherQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.voucher.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/activity/voucher/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityVoucherQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityVoucherSend calls alipay.marketing.activity.voucher.send (POST /v3/alipay/marketing/activity/voucher/send).
// Official V3 contract: https://opendocs.alipay.com/apis/05vidi
func (s *Service) ActivityVoucherSend(ctx context.Context, req ActivityVoucherSendRequest, opts ...core.CallOption) (*core.Response[ActivityVoucherSendResponse], error) {
	return core.ExecuteJSON[ActivityVoucherSendRequest, ActivityVoucherSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.voucher.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/voucher/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityVoucherSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// ActivityVoucherStop calls alipay.marketing.activity.voucher.stop (POST /v3/alipay/marketing/activity/voucher/stop).
// Official V3 contract: https://opendocs.alipay.com/apis/05vc1k
func (s *Service) ActivityVoucherStop(ctx context.Context, req ActivityVoucherStopRequest, opts ...core.CallOption) (*core.Response[ActivityVoucherStopResponse], error) {
	return core.ExecuteJSON[ActivityVoucherStopRequest, ActivityVoucherStopResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.voucher.stop", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/voucher/stop",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityVoucherStopRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
