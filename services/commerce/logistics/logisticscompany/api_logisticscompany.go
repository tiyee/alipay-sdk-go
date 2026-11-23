package logisticscompany

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

// LogisticsLogisticsCompanyInstantDeliveryQuery calls alipay.commerce.logistics.logisticscompany.instantdelivery.query (GET /v3/alipay/commerce/logistics/logisticscompany/instantdelivery/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/068msd
func (s *Service) LogisticsLogisticsCompanyInstantDeliveryQuery(ctx context.Context, req LogisticsLogisticsCompanyInstantDeliveryQueryRequest, opts ...core.CallOption) (*core.Response[LogisticsLogisticsCompanyInstantDeliveryQueryResponse], error) {
	return core.ExecuteJSON[LogisticsLogisticsCompanyInstantDeliveryQueryRequest, LogisticsLogisticsCompanyInstantDeliveryQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.logistics.logisticscompany.instantdelivery.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/logistics/logisticscompany/instantdelivery/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[LogisticsLogisticsCompanyInstantDeliveryQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
