// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package merchant

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service invokes V3 operations in this domain.
type Service struct{ executor core.Executor }

// NewService creates a service using the supplied executor.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

func execute[Req, Resp any](s *Service, ctx context.Context, method string, httpMethod core.HTTPMethod, path string, mode core.ParameterMode, req Req, opts []core.CallOption) (*core.Response[Resp], error) {
	op, err := core.NewJSONOperation[Req, Resp](core.OperationSpec{Method: method, HTTPMethod: httpMethod, Path: path, ParameterMode: mode, AuthRequirement: core.AuthRequirementAny})
	if err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, op, core.Request[Req]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SubsidiariesApply calls zhima.merchant.subsidiaries.apply.
// Official V3 contract: https://opendocs.alipay.com/apis/05ygh0
func (s *Service) SubsidiariesApply(ctx context.Context, req SubsidiariesApplyRequest, opts ...core.CallOption) (*core.Response[SubsidiariesApplyResponse], error) {
	return execute[SubsidiariesApplyRequest, SubsidiariesApplyResponse](s, ctx, "zhima.merchant.subsidiaries.apply", core.HTTPMethodPost, "/v3/zhima/merchant/subsidiaries/apply", core.ParameterModeJSONBody, req, opts)
}

// SubsidiariesClose calls zhima.merchant.subsidiaries.close.
// Official V3 contract: https://opendocs.alipay.com/apis/05yopn
func (s *Service) SubsidiariesClose(ctx context.Context, req SubsidiariesCloseRequest, opts ...core.CallOption) (*core.Response[SubsidiariesCloseResponse], error) {
	return execute[SubsidiariesCloseRequest, SubsidiariesCloseResponse](s, ctx, "zhima.merchant.subsidiaries.close", core.HTTPMethodPost, "/v3/zhima/merchant/subsidiaries/close", core.ParameterModeJSONBody, req, opts)
}

// ZmGoCumulateQuery calls zhima.merchant.zmgo.cumulate.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/62aec983_zhima.merchant.zmgo.cumulate.query
func (s *Service) ZmGoCumulateQuery(ctx context.Context, req ZmGoCumulateQueryRequest, opts ...core.CallOption) (*core.Response[ZmGoCumulateQueryResponse], error) {
	return execute[ZmGoCumulateQueryRequest, ZmGoCumulateQueryResponse](s, ctx, "zhima.merchant.zmgo.cumulate.query", core.HTTPMethodGet, "/v3/zhima/merchant/zmgo/cumulate/query", core.ParameterModeQuery, req, opts)
}

// ZmGoCumulateSync calls zhima.merchant.zmgo.cumulate.sync.
// Official V3 contract: https://opendocs.alipay.com/open-v3/153c048c_zhima.merchant.zmgo.cumulate.sync
func (s *Service) ZmGoCumulateSync(ctx context.Context, req ZmGoCumulateSyncRequest, opts ...core.CallOption) (*core.Response[ZmGoCumulateSyncResponse], error) {
	return execute[ZmGoCumulateSyncRequest, ZmGoCumulateSyncResponse](s, ctx, "zhima.merchant.zmgo.cumulate.sync", core.HTTPMethodPost, "/v3/zhima/merchant/zmgo/cumulate/sync", core.ParameterModeJSONBody, req, opts)
}

// ZmGoTemplateCreate calls zhima.merchant.zmgo.template.create.
// Official V3 contract: https://opendocs.alipay.com/open-v3/b0eba80a_zhima.merchant.zmgo.template.create
func (s *Service) ZmGoTemplateCreate(ctx context.Context, req ZmGoTemplateCreateRequest, opts ...core.CallOption) (*core.Response[ZmGoTemplateCreateResponse], error) {
	return execute[ZmGoTemplateCreateRequest, ZmGoTemplateCreateResponse](s, ctx, "zhima.merchant.zmgo.template.create", core.HTTPMethodPost, "/v3/zhima/merchant/zmgo/template/create", core.ParameterModeJSONBody, req, opts)
}

// ZmGoTemplateQuery calls zhima.merchant.zmgo.template.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/1cbec0d3_zhima.merchant.zmgo.template.query
func (s *Service) ZmGoTemplateQuery(ctx context.Context, req ZmGoTemplateQueryRequest, opts ...core.CallOption) (*core.Response[ZmGoTemplateQueryResponse], error) {
	return execute[ZmGoTemplateQueryRequest, ZmGoTemplateQueryResponse](s, ctx, "zhima.merchant.zmgo.template.query", core.HTTPMethodGet, "/v3/zhima/merchant/zmgo/template/query", core.ParameterModeQuery, req, opts)
}
