package tree

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

// CcmSwTreeBatchQuery calls alipay.iservice.ccm.sw.tree.batchquery (GET /v3/alipay/iservice/ccm/sw/tree/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/7ce6cb6e_alipay.iservice.ccm.sw.tree.batchquery
func (s *Service) CcmSwTreeBatchQuery(ctx context.Context, req CcmSwTreeBatchQueryRequest, opts ...core.CallOption) (*core.Response[CcmSwTreeBatchQueryResponse], error) {
	return core.ExecuteJSON[CcmSwTreeBatchQueryRequest, CcmSwTreeBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.tree.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/sw/tree/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwTreeBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwTreeCreate calls alipay.iservice.ccm.sw.tree.create (PUT /v3/alipay/iservice/ccm/sw/tree/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c2bb833d_alipay.iservice.ccm.sw.tree.create
func (s *Service) CcmSwTreeCreate(ctx context.Context, req CcmSwTreeCreateRequest, opts ...core.CallOption) (*core.Response[CcmSwTreeCreateResponse], error) {
	return core.ExecuteJSON[CcmSwTreeCreateRequest, CcmSwTreeCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.tree.create", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/iservice/ccm/sw/tree/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwTreeCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwTreeDelete calls alipay.iservice.ccm.sw.tree.delete (DELETE /v3/alipay/iservice/ccm/sw/tree/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/10547684_alipay.iservice.ccm.sw.tree.delete
func (s *Service) CcmSwTreeDelete(ctx context.Context, req CcmSwTreeDeleteRequest, opts ...core.CallOption) (*core.Response[CcmSwTreeDeleteResponse], error) {
	return core.ExecuteJSON[CcmSwTreeDeleteRequest, CcmSwTreeDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.tree.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/iservice/ccm/sw/tree/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwTreeDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwTreeGet calls alipay.iservice.ccm.sw.tree.get (GET /v3/alipay/iservice/ccm/sw/tree/get).
// Official V3 contract: https://opendocs.alipay.com/open-v3/930c8404_alipay.iservice.ccm.sw.tree.get
func (s *Service) CcmSwTreeGet(ctx context.Context, req CcmSwTreeGetRequest, opts ...core.CallOption) (*core.Response[CcmSwTreeGetResponse], error) {
	return core.ExecuteJSON[CcmSwTreeGetRequest, CcmSwTreeGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.tree.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/sw/tree/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwTreeGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwTreeModify calls alipay.iservice.ccm.sw.tree.modify (PATCH /v3/alipay/iservice/ccm/sw/tree/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b89ce966_alipay.iservice.ccm.sw.tree.modify
func (s *Service) CcmSwTreeModify(ctx context.Context, req CcmSwTreeModifyRequest, opts ...core.CallOption) (*core.Response[CcmSwTreeModifyResponse], error) {
	return core.ExecuteJSON[CcmSwTreeModifyRequest, CcmSwTreeModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.tree.modify", HTTPMethod: core.HTTPMethodPatch, Path: "/v3/alipay/iservice/ccm/sw/tree/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwTreeModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmSwTreePublish calls alipay.iservice.ccm.sw.tree.publish (PUT /v3/alipay/iservice/ccm/sw/tree/publish).
// Official V3 contract: https://opendocs.alipay.com/open-v3/3aa3bbe6_alipay.iservice.ccm.sw.tree.publish
func (s *Service) CcmSwTreePublish(ctx context.Context, req CcmSwTreePublishRequest, opts ...core.CallOption) (*core.Response[CcmSwTreePublishResponse], error) {
	return core.ExecuteJSON[CcmSwTreePublishRequest, CcmSwTreePublishResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.sw.tree.publish", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/iservice/ccm/sw/tree/publish",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmSwTreePublishRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
