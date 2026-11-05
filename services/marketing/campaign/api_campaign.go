// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package campaign

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

// CampaignCashCreate calls alipay.marketing.campaign.cash.create (POST /v3/alipay/marketing/campaign/cash/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e572d9c1_alipay.marketing.campaign.cash.create?scene=common&pathHash=5197b580
func (s *Service) CampaignCashCreate(ctx context.Context, req CampaignCashCreateRequest, opts ...core.CallOption) (*core.Response[CampaignCashCreateResponse], error) {
	return core.ExecuteJSON[CampaignCashCreateRequest, CampaignCashCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.campaign.cash.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/campaign/cash/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CampaignCashCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CampaignCashDetailQuery calls alipay.marketing.campaign.cash.detail.query (GET /v3/alipay/marketing/campaign/cash/detail/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/7fede5af_alipay.marketing.campaign.cash.detail.query?scene=common&pathHash=430db29a
func (s *Service) CampaignCashDetailQuery(ctx context.Context, req CampaignCashDetailQueryRequest, opts ...core.CallOption) (*core.Response[CampaignCashDetailQueryResponse], error) {
	return core.ExecuteJSON[CampaignCashDetailQueryRequest, CampaignCashDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.campaign.cash.detail.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/campaign/cash/detail/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CampaignCashDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CampaignCashListQuery calls alipay.marketing.campaign.cash.list.query (GET /v3/alipay/marketing/campaign/cash/list/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b0871b97_alipay.marketing.campaign.cash.list.query?scene=common&pathHash=e06176c5
func (s *Service) CampaignCashListQuery(ctx context.Context, req CampaignCashListQueryRequest, opts ...core.CallOption) (*core.Response[CampaignCashListQueryResponse], error) {
	return core.ExecuteJSON[CampaignCashListQueryRequest, CampaignCashListQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.campaign.cash.list.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/campaign/cash/list/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CampaignCashListQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CampaignCashStatusModify calls alipay.marketing.campaign.cash.status.modify (POST /v3/alipay/marketing/campaign/cash/status/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4f41b60d_alipay.marketing.campaign.cash.status.modify?scene=common&pathHash=7b368c26
func (s *Service) CampaignCashStatusModify(ctx context.Context, req CampaignCashStatusModifyRequest, opts ...core.CallOption) (*core.Response[CampaignCashStatusModifyResponse], error) {
	return core.ExecuteJSON[CampaignCashStatusModifyRequest, CampaignCashStatusModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.campaign.cash.status.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/campaign/cash/status/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CampaignCashStatusModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CampaignCashTrigger calls alipay.marketing.campaign.cash.trigger (POST /v3/alipay/marketing/campaign/cash/trigger).
// Official V3 contract: https://opendocs.alipay.com/open-v3/06823eee_alipay.marketing.campaign.cash.trigger?scene=common&pathHash=e770d208
func (s *Service) CampaignCashTrigger(ctx context.Context, req CampaignCashTriggerRequest, opts ...core.CallOption) (*core.Response[CampaignCashTriggerResponse], error) {
	return core.ExecuteJSON[CampaignCashTriggerRequest, CampaignCashTriggerResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.campaign.cash.trigger", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/campaign/cash/trigger",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CampaignCashTriggerRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CampaignOrderVoucherConsult calls alipay.marketing.campaign.order.voucher.consult (POST /v3/alipay/marketing/campaign/order/voucher/consult).
// Official V3 contract: https://opendocs.alipay.com/open-v3/bb374dd7_alipay.marketing.campaign.order.voucher.consult
func (s *Service) CampaignOrderVoucherConsult(ctx context.Context, req CampaignOrderVoucherConsultRequest, opts ...core.CallOption) (*core.Response[CampaignOrderVoucherConsultResponse], error) {
	return core.ExecuteJSON[CampaignOrderVoucherConsultRequest, CampaignOrderVoucherConsultResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.campaign.order.voucher.consult", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/campaign/order/voucher/consult",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CampaignOrderVoucherConsultRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
