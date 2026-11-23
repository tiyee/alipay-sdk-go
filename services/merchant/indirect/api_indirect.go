package indirect

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

// IndirectAuthOrderClose calls alipay.merchant.indirect.authorder.close (PUT /v3/alipay/merchant/indirect/authorder/close).
// Official V3 contract: https://opendocs.alipay.com/apis/063knt
func (s *Service) IndirectAuthOrderClose(ctx context.Context, req MerchantIndirectAuthOrderCloseRequest, opts ...core.CallOption) (*core.Response[MerchantIndirectAuthOrderCloseResponse], error) {
	return core.ExecuteJSON[MerchantIndirectAuthOrderCloseRequest, MerchantIndirectAuthOrderCloseResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.indirect.authorder.close", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/merchant/indirect/authorder/close",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantIndirectAuthOrderCloseRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// IndirectAuthOrderCreate calls alipay.merchant.indirect.authorder.create (POST /v3/alipay/merchant/indirect/authorder/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05gz1v
func (s *Service) IndirectAuthOrderCreate(ctx context.Context, req MerchantIndirectAuthOrderCreateRequest, opts ...core.CallOption) (*core.Response[MerchantIndirectAuthOrderCreateResponse], error) {
	return core.ExecuteJSON[MerchantIndirectAuthOrderCreateRequest, MerchantIndirectAuthOrderCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.indirect.authorder.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/indirect/authorder/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantIndirectAuthOrderCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// IndirectAuthOrderQueryStatus calls alipay.merchant.indirect.authorder.querystatus (POST /v3/alipay/merchant/indirect/authorder/querystatus).
// Official V3 contract: https://opendocs.alipay.com/apis/063luk
func (s *Service) IndirectAuthOrderQueryStatus(ctx context.Context, req MerchantIndirectAuthOrderQueryStatusRequest, opts ...core.CallOption) (*core.Response[MerchantIndirectAuthOrderQueryStatusResponse], error) {
	return core.ExecuteJSON[MerchantIndirectAuthOrderQueryStatusRequest, MerchantIndirectAuthOrderQueryStatusResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.indirect.authorder.querystatus", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/merchant/indirect/authorder/querystatus",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantIndirectAuthOrderQueryStatusRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// IndirectSMIDBindQuery calls alipay.merchant.indirect.smidbind.query (GET /v3/alipay/merchant/indirect/smidbind/query).
// Official V3 contract: https://opendocs.alipay.com/apis/063ec9
func (s *Service) IndirectSMIDBindQuery(ctx context.Context, req MerchantIndirectSMIDBindQueryRequest, opts ...core.CallOption) (*core.Response[MerchantIndirectSMIDBindQueryResponse], error) {
	return core.ExecuteJSON[MerchantIndirectSMIDBindQueryRequest, MerchantIndirectSMIDBindQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.merchant.indirect.smidbind.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/merchant/indirect/smidbind/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MerchantIndirectSMIDBindQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
