package account

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

// InstantDeliveryAccountCreate calls alipay.open.instantdelivery.account.create (POST /v3/alipay/open/instantdelivery/account/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vlui
func (s *Service) InstantDeliveryAccountCreate(ctx context.Context, req InstantDeliveryAccountCreateRequest, opts ...core.CallOption) (*core.Response[InstantDeliveryAccountCreateResponse], error) {
	return core.ExecuteJSON[InstantDeliveryAccountCreateRequest, InstantDeliveryAccountCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.instantdelivery.account.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/instantdelivery/account/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InstantDeliveryAccountCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InstantDeliveryAccountQuery calls alipay.open.instantdelivery.account.query (GET /v3/alipay/open/instantdelivery/account/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vda7
func (s *Service) InstantDeliveryAccountQuery(ctx context.Context, req InstantDeliveryAccountQueryRequest, opts ...core.CallOption) (*core.Response[InstantDeliveryAccountQueryResponse], error) {
	return core.ExecuteJSON[InstantDeliveryAccountQueryRequest, InstantDeliveryAccountQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.instantdelivery.account.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/instantdelivery/account/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InstantDeliveryAccountQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
