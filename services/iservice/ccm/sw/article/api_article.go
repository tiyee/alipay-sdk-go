package article

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

// CcmSwArticleBatchQuery calls alipay.iservice.ccm.sw.article.batchquery (POST /v3/alipay/iservice/ccm/sw/article/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/031db27a_alipay.iservice.ccm.sw.article.batchquery
func (s *Service) CcmSwArticleBatchQuery(ctx context.Context, req CcmSwArticleBatchQueryRequest, opts ...core.CallOption) (*core.Response[CcmSwArticleBatchQueryResponse], error) {
	return core.ExecuteJSON[CcmSwArticleBatchQueryRequest, CcmSwArticleBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.article.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/sw/article/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwArticleBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwArticleCreate calls alipay.iservice.ccm.sw.article.create (PUT /v3/alipay/iservice/ccm/sw/article/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/7c58db8d_alipay.iservice.ccm.sw.article.create
func (s *Service) CcmSwArticleCreate(ctx context.Context, req CcmSwArticleCreateRequest, opts ...core.CallOption) (*core.Response[CcmSwArticleCreateResponse], error) {
	return core.ExecuteJSON[CcmSwArticleCreateRequest, CcmSwArticleCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.article.create", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/iservice/ccm/sw/article/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwArticleCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwArticleDelete calls alipay.iservice.ccm.sw.article.delete (DELETE /v3/alipay/iservice/ccm/sw/article/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/d2475255_alipay.iservice.ccm.sw.article.delete
func (s *Service) CcmSwArticleDelete(ctx context.Context, req CcmSwArticleDeleteRequest, opts ...core.CallOption) (*core.Response[CcmSwArticleDeleteResponse], error) {
	return core.ExecuteJSON[CcmSwArticleDeleteRequest, CcmSwArticleDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.article.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/iservice/ccm/sw/article/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwArticleDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwArticleGet calls alipay.iservice.ccm.sw.article.get (GET /v3/alipay/iservice/ccm/sw/article/get).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c2addca8_alipay.iservice.ccm.sw.article.get
func (s *Service) CcmSwArticleGet(ctx context.Context, req CcmSwArticleGetRequest, opts ...core.CallOption) (*core.Response[CcmSwArticleGetResponse], error) {
	return core.ExecuteJSON[CcmSwArticleGetRequest, CcmSwArticleGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.article.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/sw/article/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwArticleGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwArticleModify calls alipay.iservice.ccm.sw.article.modify (PATCH /v3/alipay/iservice/ccm/sw/article/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/cc0b0135_alipay.iservice.ccm.sw.article.modify
func (s *Service) CcmSwArticleModify(ctx context.Context, req CcmSwArticleModifyRequest, opts ...core.CallOption) (*core.Response[CcmSwArticleModifyResponse], error) {
	return core.ExecuteJSON[CcmSwArticleModifyRequest, CcmSwArticleModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.article.modify", HTTPMethod: core.HTTPMethodPatch, Path: "/v3/alipay/iservice/ccm/sw/article/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwArticleModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
