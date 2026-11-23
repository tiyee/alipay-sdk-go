package isv

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

// SpISVBillSync calls alipay.open.sp.isv.bill.sync (POST /v3/alipay/open/sp/isv/bill/sync).
// Official V3 contract: https://opendocs.alipay.com/apis/05vidn
func (s *Service) SpISVBillSync(ctx context.Context, req SpISVBillSyncRequest, opts ...core.CallOption) (*core.Response[SpISVBillSyncResponse], error) {
	return core.ExecuteJSON[SpISVBillSyncRequest, SpISVBillSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.isv.bill.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/sp/isv/bill/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpISVBillSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SpISVSignAuthCreate calls alipay.open.sp.isv.signauth.create (POST /v3/alipay/open/sp/isv/signauth/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e4f4b975_alipay.open.sp.isv.signauth.create
func (s *Service) SpISVSignAuthCreate(ctx context.Context, req SpISVSignAuthCreateRequest, opts ...core.CallOption) (*core.Response[SpISVSignAuthCreateResponse], error) {
	return core.ExecuteJSON[SpISVSignAuthCreateRequest, SpISVSignAuthCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.isv.signauth.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/sp/isv/signauth/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpISVSignAuthCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// SpISVSignAuthQuery calls alipay.open.sp.isv.signauth.query (GET /v3/alipay/open/sp/isv/signauth/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a322ed32_alipay.open.sp.isv.signauth.query
func (s *Service) SpISVSignAuthQuery(ctx context.Context, req SpISVSignAuthQueryRequest, opts ...core.CallOption) (*core.Response[SpISVSignAuthQueryResponse], error) {
	return core.ExecuteJSON[SpISVSignAuthQueryRequest, SpISVSignAuthQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.sp.isv.signauth.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/sp/isv/signauth/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SpISVSignAuthQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
