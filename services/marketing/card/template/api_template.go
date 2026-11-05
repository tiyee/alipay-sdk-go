package template

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

// CardTemplateBatchQuery calls alipay.marketing.card.template.batchquery (POST /v3/alipay/marketing/card/template/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/54dcbcfa_alipay.marketing.card.template.batchquery?pathHash=030a7dec
func (s *Service) CardTemplateBatchQuery(ctx context.Context, req CardTemplateBatchQueryRequest, opts ...core.CallOption) (*core.Response[CardTemplateBatchQueryResponse], error) {
	return core.ExecuteJSON[CardTemplateBatchQueryRequest, CardTemplateBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.template.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/template/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardTemplateBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardTemplateCreate calls alipay.marketing.card.template.create (POST /v3/alipay/marketing/card/template/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/18fdd0bd_alipay.marketing.card.template.create?scene=1c74e13b68cc4a58912fb6f0c40b62af%2C1529c3090a7b4db99164667a0b22dcaa&pathHash=d206dedb
func (s *Service) CardTemplateCreate(ctx context.Context, req CardTemplateCreateRequest, opts ...core.CallOption) (*core.Response[CardTemplateCreateResponse], error) {
	return core.ExecuteJSON[CardTemplateCreateRequest, CardTemplateCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.template.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/template/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardTemplateCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardTemplateModify calls alipay.marketing.card.template.modify (POST /v3/alipay/marketing/card/template/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/fabfc518_alipay.marketing.card.template.modify?scene=692a39a5ad9e45d4899177a7ec9cf68e%2Cc76bbe273e3b471bbba690c624b9ed86&pathHash=e43d6caf
func (s *Service) CardTemplateModify(ctx context.Context, req CardTemplateModifyRequest, opts ...core.CallOption) (*core.Response[CardTemplateModifyResponse], error) {
	return core.ExecuteJSON[CardTemplateModifyRequest, CardTemplateModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.template.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/card/template/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardTemplateModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CardTemplateQuery calls alipay.marketing.card.template.query (GET /v3/alipay/marketing/card/template/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/0bd1a249_alipay.marketing.card.template.query?scene=871b0111e48c4490ad0a7dbf2a465b36%2C1eca10e1486044f9b3a0d9171111f810&pathHash=84fe74c8
func (s *Service) CardTemplateQuery(ctx context.Context, req CardTemplateQueryRequest, opts ...core.CallOption) (*core.Response[CardTemplateQueryResponse], error) {
	return core.ExecuteJSON[CardTemplateQueryRequest, CardTemplateQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.card.template.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/marketing/card/template/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CardTemplateQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
