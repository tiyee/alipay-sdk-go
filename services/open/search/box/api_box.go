package box

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

// SearchBoxApply calls alipay.open.search.box.apply (POST /v3/alipay/open/search/box/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/622c6b11_alipay.open.search.box.apply?scene=common&pathHash=d1a9edf5
func (s *Service) SearchBoxApply(ctx context.Context, req SearchBoxApplyRequest, opts ...core.CallOption) (*core.Response[SearchBoxApplyResponse], error) {
	return core.ExecuteJSON[SearchBoxApplyRequest, SearchBoxApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.box.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/box/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchBoxApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchBoxBatchQuery calls alipay.open.search.box.batchquery (POST /v3/alipay/open/search/box/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/fd1480ed_alipay.open.search.box.batchquery?scene=common&pathHash=4ca4364b
func (s *Service) SearchBoxBatchQuery(ctx context.Context, req SearchBoxBatchQueryRequest, opts ...core.CallOption) (*core.Response[SearchBoxBatchQueryResponse], error) {
	return core.ExecuteJSON[SearchBoxBatchQueryRequest, SearchBoxBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.box.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/box/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchBoxBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchBoxConsult calls alipay.open.search.box.consult (POST /v3/alipay/open/search/box/consult).
// Official V3 contract: https://opendocs.alipay.com/open-v3/28d19e54_alipay.open.search.box.consult
func (s *Service) SearchBoxConsult(ctx context.Context, req SearchBoxConsultRequest, opts ...core.CallOption) (*core.Response[SearchBoxConsultResponse], error) {
	return core.ExecuteJSON[SearchBoxConsultRequest, SearchBoxConsultResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.box.consult", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/box/consult",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchBoxConsultRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchBoxModify calls alipay.open.search.box.modify (POST /v3/alipay/open/search/box/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/28bd3996_alipay.open.search.box.modify?scene=common&pathHash=f07276ab
func (s *Service) SearchBoxModify(ctx context.Context, req SearchBoxModifyRequest, opts ...core.CallOption) (*core.Response[SearchBoxModifyResponse], error) {
	return core.ExecuteJSON[SearchBoxModifyRequest, SearchBoxModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.box.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/box/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchBoxModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchBoxOffline calls alipay.open.search.box.offline (POST /v3/alipay/open/search/box/offline).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4d9cc1dc_alipay.open.search.box.offline
func (s *Service) SearchBoxOffline(ctx context.Context, req SearchBoxOfflineRequest, opts ...core.CallOption) (*core.Response[SearchBoxOfflineResponse], error) {
	return core.ExecuteJSON[SearchBoxOfflineRequest, SearchBoxOfflineResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.box.offline", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/box/offline",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchBoxOfflineRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchBoxOnline calls alipay.open.search.box.online (POST /v3/alipay/open/search/box/online).
// Official V3 contract: https://opendocs.alipay.com/open-v3/00d18795_alipay.open.search.box.online
func (s *Service) SearchBoxOnline(ctx context.Context, req SearchBoxOnlineRequest, opts ...core.CallOption) (*core.Response[SearchBoxOnlineResponse], error) {
	return core.ExecuteJSON[SearchBoxOnlineRequest, SearchBoxOnlineResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.box.online", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/box/online",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchBoxOnlineRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchBoxQuery calls alipay.open.search.box.query (GET /v3/alipay/open/search/box/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/97b73d37_alipay.open.search.box.query?scene=common&pathHash=6f2e052c
func (s *Service) SearchBoxQuery(ctx context.Context, req SearchBoxQueryRequest, opts ...core.CallOption) (*core.Response[SearchBoxQueryResponse], error) {
	return core.ExecuteJSON[SearchBoxQueryRequest, SearchBoxQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.box.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/search/box/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchBoxQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchBoxActivityApply calls alipay.open.search.boxactivity.apply (POST /v3/alipay/open/search/boxactivity/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/33107e1b_alipay.open.search.boxactivity.apply
func (s *Service) SearchBoxActivityApply(ctx context.Context, req SearchBoxActivityApplyRequest, opts ...core.CallOption) (*core.Response[SearchBoxActivityApplyResponse], error) {
	return core.ExecuteJSON[SearchBoxActivityApplyRequest, SearchBoxActivityApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.boxactivity.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/search/boxactivity/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchBoxActivityApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SearchBoxActivityQuery calls alipay.open.search.boxactivity.query (GET /v3/alipay/open/search/boxactivity/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/5b195961_alipay.open.search.boxactivity.query
func (s *Service) SearchBoxActivityQuery(ctx context.Context, req SearchBoxActivityQueryRequest, opts ...core.CallOption) (*core.Response[SearchBoxActivityQueryResponse], error) {
	return core.ExecuteJSON[SearchBoxActivityQueryRequest, SearchBoxActivityQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.search.boxactivity.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/search/boxactivity/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SearchBoxActivityQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
