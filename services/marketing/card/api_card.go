package card

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor core.Executor
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}

	return &Service{
		executor: executor,
	}, nil
}

// CardActivateFormQuery calls alipay.marketing.card.activateform.query (GET /v3/alipay/marketing/card/activateform/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a195d7c2_alipay.marketing.card.activateform.query
func (s *Service) CardActivateFormQuery(ctx context.Context, req CardActivateFormQueryRequest, opts ...core.CallOption) (*core.Response[CardActivateFormQueryResponse], error) {
	return core.ExecuteJSON[CardActivateFormQueryRequest, CardActivateFormQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.activateform.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/card/activateform/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardActivateFormQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardActivateURLApply calls alipay.marketing.card.activateurl.apply (POST /v3/alipay/marketing/card/activateurl/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/85f9492f_alipay.marketing.card.activateurl.apply
func (s *Service) CardActivateURLApply(ctx context.Context, req CardActivateURLApplyRequest, opts ...core.CallOption) (*core.Response[CardActivateURLApplyResponse], error) {
	return core.ExecuteJSON[CardActivateURLApplyRequest, CardActivateURLApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.activateurl.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/activateurl/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardActivateURLApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardConsumeSync calls alipay.marketing.card.consume.sync (POST /v3/alipay/marketing/card/consume/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e0f362b2_alipay.marketing.card.consume.sync
func (s *Service) CardConsumeSync(ctx context.Context, req CardConsumeSyncRequest, opts ...core.CallOption) (*core.Response[CardConsumeSyncResponse], error) {
	return core.ExecuteJSON[CardConsumeSyncRequest, CardConsumeSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.consume.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/consume/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardConsumeSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardDelete calls alipay.marketing.card.delete (DELETE /v3/alipay/marketing/card/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/fa4f5056_alipay.marketing.card.delete?scene=common&pathHash=b9e0e504
func (s *Service) CardDelete(ctx context.Context, req CardDeleteRequest, opts ...core.CallOption) (*core.Response[CardDeleteResponse], error) {
	return core.ExecuteJSON[CardDeleteRequest, CardDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/marketing/card/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardFormTemplateSet calls alipay.marketing.card.formtemplate.set (POST /v3/alipay/marketing/card/formtemplate/set).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c7503ca0_alipay.marketing.card.formtemplate.set?scene=common&pathHash=36809f41
func (s *Service) CardFormTemplateSet(ctx context.Context, req CardFormTemplateSetRequest, opts ...core.CallOption) (*core.Response[CardFormTemplateSetResponse], error) {
	return core.ExecuteJSON[CardFormTemplateSetRequest, CardFormTemplateSetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.formtemplate.set", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/formtemplate/set",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardFormTemplateSetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardMessageNotify calls alipay.marketing.card.message.notify (POST /v3/alipay/marketing/card/message/notify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b710cee9_alipay.marketing.card.message.notify?scene=common&pathHash=9eb33531
func (s *Service) CardMessageNotify(ctx context.Context, req CardMessageNotifyRequest, opts ...core.CallOption) (*core.Response[CardMessageNotifyResponse], error) {
	return core.ExecuteJSON[CardMessageNotifyRequest, CardMessageNotifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.message.notify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/message/notify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardMessageNotifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardOpen calls alipay.marketing.card.open (POST /v3/alipay/marketing/card/open).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8e9fe25b_alipay.marketing.card.open
func (s *Service) CardOpen(ctx context.Context, req CardOpenRequest, opts ...core.CallOption) (*core.Response[CardOpenResponse], error) {
	return core.ExecuteJSON[CardOpenRequest, CardOpenResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.open", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/open",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardOpenRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardQuery calls alipay.marketing.card.query (POST /v3/alipay/marketing/card/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b2494988_alipay.marketing.card.query?scene=common&pathHash=842cfce1
func (s *Service) CardQuery(ctx context.Context, req CardQueryRequest, opts ...core.CallOption) (*core.Response[CardQueryResponse], error) {
	return core.ExecuteJSON[CardQueryRequest, CardQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardUpdate calls alipay.marketing.card.update (POST /v3/alipay/marketing/card/update).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c83a3e12_alipay.marketing.card.update?scene=common&pathHash=dc8a305e
func (s *Service) CardUpdate(ctx context.Context, req CardUpdateRequest, opts ...core.CallOption) (*core.Response[CardUpdateResponse], error) {
	return core.ExecuteJSON[CardUpdateRequest, CardUpdateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.update", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/update",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardUpdateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
