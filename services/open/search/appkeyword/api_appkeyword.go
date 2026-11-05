package appkeyword

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

// SearchAppKeywordApply calls alipay.open.search.appkeyword.apply (POST /v3/alipay/open/search/appkeyword/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/5ed663e2_alipay.open.search.appkeyword.apply
func (s *Service) SearchAppKeywordApply(ctx context.Context, req SearchAppKeywordApplyRequest, opts ...core.CallOption) (*core.Response[SearchAppKeywordApplyResponse], error) {
	return core.ExecuteJSON[SearchAppKeywordApplyRequest, SearchAppKeywordApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.appkeyword.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/appkeyword/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchAppKeywordApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchAppKeywordBatchQuery calls alipay.open.search.appkeyword.batchquery (GET /v3/alipay/open/search/appkeyword/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/0e3cdc41_alipay.open.search.appkeyword.batchquery
func (s *Service) SearchAppKeywordBatchQuery(ctx context.Context, req SearchAppKeywordBatchQueryRequest, opts ...core.CallOption) (*core.Response[SearchAppKeywordBatchQueryResponse], error) {
	return core.ExecuteJSON[SearchAppKeywordBatchQueryRequest, SearchAppKeywordBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.appkeyword.batchquery", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/search/appkeyword/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchAppKeywordBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchAppKeywordDelete calls alipay.open.search.appkeyword.delete (POST /v3/alipay/open/search/appkeyword/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e7b09aa2_alipay.open.search.appkeyword.delete
func (s *Service) SearchAppKeywordDelete(ctx context.Context, req SearchAppKeywordDeleteRequest, opts ...core.CallOption) (*core.Response[SearchAppKeywordDeleteResponse], error) {
	return core.ExecuteJSON[SearchAppKeywordDeleteRequest, SearchAppKeywordDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.appkeyword.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/appkeyword/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchAppKeywordDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchAppKeywordQueryStatus calls alipay.open.search.appkeyword.querystatus (GET /v3/alipay/open/search/appkeyword/querystatus).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2028e38b_alipay.open.search.appkeyword.querystatus
func (s *Service) SearchAppKeywordQueryStatus(ctx context.Context, req SearchAppKeywordQueryStatusRequest, opts ...core.CallOption) (*core.Response[SearchAppKeywordQueryStatusResponse], error) {
	return core.ExecuteJSON[SearchAppKeywordQueryStatusRequest, SearchAppKeywordQueryStatusResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.appkeyword.querystatus", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/search/appkeyword/querystatus",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchAppKeywordQueryStatusRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchAppKeywordQuotaQuery calls alipay.open.search.appkeywordquota.query (GET /v3/alipay/open/search/appkeywordquota/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/764cf4b6_alipay.open.search.appkeywordquota.query
func (s *Service) SearchAppKeywordQuotaQuery(ctx context.Context, req SearchAppKeywordQuotaQueryRequest, opts ...core.CallOption) (*core.Response[SearchAppKeywordQuotaQueryResponse], error) {
	return core.ExecuteJSON[SearchAppKeywordQuotaQueryRequest, SearchAppKeywordQuotaQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.appkeywordquota.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/search/appkeywordquota/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchAppKeywordQuotaQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
