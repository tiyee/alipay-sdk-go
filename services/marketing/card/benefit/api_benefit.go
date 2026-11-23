package benefit

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

// CardBenefitCreate calls alipay.marketing.card.benefit.create (POST /v3/alipay/marketing/card/benefit/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2aa50202_alipay.marketing.card.benefit.create?scene=common&pathHash=c87e30d3
func (s *Service) CardBenefitCreate(ctx context.Context, req CardBenefitCreateRequest, opts ...core.CallOption) (*core.Response[CardBenefitCreateResponse], error) {
	return core.ExecuteJSON[CardBenefitCreateRequest, CardBenefitCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.benefit.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/benefit/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardBenefitCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardBenefitDelete calls alipay.marketing.card.benefit.delete (DELETE /v3/alipay/marketing/card/benefit/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/85008a68_alipay.marketing.card.benefit.delete?pathHash=112339fb
func (s *Service) CardBenefitDelete(ctx context.Context, req CardBenefitDeleteRequest, opts ...core.CallOption) (*core.Response[CardBenefitDeleteResponse], error) {
	return core.ExecuteJSON[CardBenefitDeleteRequest, CardBenefitDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.benefit.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/marketing/card/benefit/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardBenefitDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardBenefitModify calls alipay.marketing.card.benefit.modify (POST /v3/alipay/marketing/card/benefit/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/f8a32ad5_alipay.marketing.card.benefit.modify?pathHash=df326137
func (s *Service) CardBenefitModify(ctx context.Context, req CardBenefitModifyRequest, opts ...core.CallOption) (*core.Response[CardBenefitModifyResponse], error) {
	return core.ExecuteJSON[CardBenefitModifyRequest, CardBenefitModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.benefit.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/benefit/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardBenefitModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardBenefitQuery calls alipay.marketing.card.benefit.query (GET /v3/alipay/marketing/card/benefit/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/217a04aa_alipay.marketing.card.benefit.query?pathHash=c22b4480
func (s *Service) CardBenefitQuery(ctx context.Context, req CardBenefitQueryRequest, opts ...core.CallOption) (*core.Response[CardBenefitQueryResponse], error) {
	return core.ExecuteJSON[CardBenefitQueryRequest, CardBenefitQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.benefit.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/card/benefit/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardBenefitQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
