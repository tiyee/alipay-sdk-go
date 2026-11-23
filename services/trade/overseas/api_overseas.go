package overseas

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor       core.Executor
	overseasSettle core.Operation[OverseasSettleRequest, OverseasSettleResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	overseasSettle, err := core.NewJSONOperation[OverseasSettleRequest, OverseasSettleResponse](core.OperationSpec{
		Method: "alipay.trade.overseas.settle", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/trade/overseas/settle",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:       executor,
		overseasSettle: overseasSettle,
	}, nil
}

// OverseasSettle 调用 alipay.trade.overseas.settle 发起跨境结算
// （V3 POST /v3/alipay/trade/overseas/settle）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/apis/05v7c4
// OutRequestNo 用于幂等控制；写操作，SDK 不自动重试。
func (s *Service) OverseasSettle(ctx context.Context, req OverseasSettleRequest,
	opts ...core.CallOption) (*core.Response[OverseasSettleResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.overseasSettle,
		core.Request[OverseasSettleRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
