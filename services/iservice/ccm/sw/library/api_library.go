package library

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

// CcmSwLibraryBatchQuery calls alipay.iservice.ccm.sw.library.batchquery (GET /v3/alipay/iservice/ccm/sw/library/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/04e0096f_alipay.iservice.ccm.sw.library.batchquery
func (s *Service) CcmSwLibraryBatchQuery(ctx context.Context, req CcmSwLibraryBatchQueryRequest, opts ...core.CallOption) (*core.Response[CcmSwLibraryBatchQueryResponse], error) {
	return core.ExecuteJSON[CcmSwLibraryBatchQueryRequest, CcmSwLibraryBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.library.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/sw/library/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwLibraryBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwLibraryCreate calls alipay.iservice.ccm.sw.library.create (PUT /v3/alipay/iservice/ccm/sw/library/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/404aca95_alipay.iservice.ccm.sw.library.create
func (s *Service) CcmSwLibraryCreate(ctx context.Context, req CcmSwLibraryCreateRequest, opts ...core.CallOption) (*core.Response[CcmSwLibraryCreateResponse], error) {
	return core.ExecuteJSON[CcmSwLibraryCreateRequest, CcmSwLibraryCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.library.create", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/iservice/ccm/sw/library/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwLibraryCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwLibraryDelete calls alipay.iservice.ccm.sw.library.delete (DELETE /v3/alipay/iservice/ccm/sw/library/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/cd16a3aa_alipay.iservice.ccm.sw.library.delete
func (s *Service) CcmSwLibraryDelete(ctx context.Context, req CcmSwLibraryDeleteRequest, opts ...core.CallOption) (*core.Response[CcmSwLibraryDeleteResponse], error) {
	return core.ExecuteJSON[CcmSwLibraryDeleteRequest, CcmSwLibraryDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.library.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/iservice/ccm/sw/library/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwLibraryDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwLibraryModify calls alipay.iservice.ccm.sw.library.modify (PATCH /v3/alipay/iservice/ccm/sw/library/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/10893431_alipay.iservice.ccm.sw.library.modify
func (s *Service) CcmSwLibraryModify(ctx context.Context, req CcmSwLibraryModifyRequest, opts ...core.CallOption) (*core.Response[CcmSwLibraryModifyResponse], error) {
	return core.ExecuteJSON[CcmSwLibraryModifyRequest, CcmSwLibraryModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.library.modify", HTTPMethod: core.HTTPMethodPatch, Path: "/v3/alipay/iservice/ccm/sw/library/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwLibraryModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
