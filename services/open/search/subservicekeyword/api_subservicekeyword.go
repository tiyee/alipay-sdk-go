package subservicekeyword

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

// SearchSubServiceKeywordApply calls alipay.open.search.subservicekeyword.apply (POST /v3/alipay/open/search/subservicekeyword/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/183cd891_alipay.open.search.subservicekeyword.apply
func (s *Service) SearchSubServiceKeywordApply(ctx context.Context, req SearchSubServiceKeywordApplyRequest, opts ...core.CallOption) (*core.Response[SearchSubServiceKeywordApplyResponse], error) {
	return core.ExecuteJSON[SearchSubServiceKeywordApplyRequest, SearchSubServiceKeywordApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.subservicekeyword.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/subservicekeyword/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchSubServiceKeywordApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchSubServiceKeywordBatchQuery calls alipay.open.search.subservicekeyword.batchquery (GET /v3/alipay/open/search/subservicekeyword/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/6ee3720b_alipay.open.search.subservicekeyword.batchquery
func (s *Service) SearchSubServiceKeywordBatchQuery(ctx context.Context, req SearchSubServiceKeywordBatchQueryRequest, opts ...core.CallOption) (*core.Response[SearchSubServiceKeywordBatchQueryResponse], error) {
	return core.ExecuteJSON[SearchSubServiceKeywordBatchQueryRequest, SearchSubServiceKeywordBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.subservicekeyword.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/search/subservicekeyword/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchSubServiceKeywordBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchSubServiceKeywordDelete calls alipay.open.search.subservicekeyword.delete (POST /v3/alipay/open/search/subservicekeyword/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/bf575dfc_alipay.open.search.subservicekeyword.delete
func (s *Service) SearchSubServiceKeywordDelete(ctx context.Context, req SearchSubServiceKeywordDeleteRequest, opts ...core.CallOption) (*core.Response[SearchSubServiceKeywordDeleteResponse], error) {
	return core.ExecuteJSON[SearchSubServiceKeywordDeleteRequest, SearchSubServiceKeywordDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.subservicekeyword.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/subservicekeyword/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchSubServiceKeywordDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchSubServiceKeywordQueryStatus calls alipay.open.search.subservicekeyword.querystatus (GET /v3/alipay/open/search/subservicekeyword/querystatus).
// Official V3 contract: https://opendocs.alipay.com/open-v3/7acef1a0_alipay.open.search.subservicekeyword.querystatus
func (s *Service) SearchSubServiceKeywordQueryStatus(ctx context.Context, req SearchSubServiceKeywordQueryStatusRequest, opts ...core.CallOption) (*core.Response[SearchSubServiceKeywordQueryStatusResponse], error) {
	return core.ExecuteJSON[SearchSubServiceKeywordQueryStatusRequest, SearchSubServiceKeywordQueryStatusResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.subservicekeyword.querystatus", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/search/subservicekeyword/querystatus",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchSubServiceKeywordQueryStatusRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
