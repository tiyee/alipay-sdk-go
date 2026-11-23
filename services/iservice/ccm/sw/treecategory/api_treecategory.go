package treecategory

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

// CcmSwTreeCategoryBatchQuery calls alipay.iservice.ccm.sw.treecategory.batchquery (GET /v3/alipay/iservice/ccm/sw/treecategory/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/24848e83_alipay.iservice.ccm.sw.treecategory.batchquery
func (s *Service) CcmSwTreeCategoryBatchQuery(ctx context.Context, req CcmSwTreeCategoryBatchQueryRequest, opts ...core.CallOption) (*core.Response[CcmSwTreeCategoryBatchQueryResponse], error) {
	if (req.PageNum == nil) != (req.PageSize == nil) {
		return nil, &core.ValidationError{Field: "page_num", Reason: "page_num and page_size must be provided together"}
	}
	return core.ExecuteJSON[CcmSwTreeCategoryBatchQueryRequest, CcmSwTreeCategoryBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.treecategory.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/sw/treecategory/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwTreeCategoryBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwTreeCategoryCreate calls alipay.iservice.ccm.sw.treecategory.create (PUT /v3/alipay/iservice/ccm/sw/treecategory/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/08f8ea8c_alipay.iservice.ccm.sw.treecategory.create
func (s *Service) CcmSwTreeCategoryCreate(ctx context.Context, req CcmSwTreeCategoryCreateRequest, opts ...core.CallOption) (*core.Response[CcmSwTreeCategoryCreateResponse], error) {
	return core.ExecuteJSON[CcmSwTreeCategoryCreateRequest, CcmSwTreeCategoryCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.treecategory.create", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/iservice/ccm/sw/treecategory/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwTreeCategoryCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwTreeCategoryDelete calls alipay.iservice.ccm.sw.treecategory.delete (DELETE /v3/alipay/iservice/ccm/sw/treecategory/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/cf84306f_alipay.iservice.ccm.sw.treecategory.delete
func (s *Service) CcmSwTreeCategoryDelete(ctx context.Context, req CcmSwTreeCategoryDeleteRequest, opts ...core.CallOption) (*core.Response[CcmSwTreeCategoryDeleteResponse], error) {
	return core.ExecuteJSON[CcmSwTreeCategoryDeleteRequest, CcmSwTreeCategoryDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.treecategory.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/iservice/ccm/sw/treecategory/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwTreeCategoryDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwTreeCategoryModify calls alipay.iservice.ccm.sw.treecategory.modify (PATCH /v3/alipay/iservice/ccm/sw/treecategory/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/6fea3439_alipay.iservice.ccm.sw.treecategory.modify
func (s *Service) CcmSwTreeCategoryModify(ctx context.Context, req CcmSwTreeCategoryModifyRequest, opts ...core.CallOption) (*core.Response[CcmSwTreeCategoryModifyResponse], error) {
	return core.ExecuteJSON[CcmSwTreeCategoryModifyRequest, CcmSwTreeCategoryModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.treecategory.modify", HTTPMethod: core.HTTPMethodPatch, Path: "/v3/alipay/iservice/ccm/sw/treecategory/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwTreeCategoryModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
