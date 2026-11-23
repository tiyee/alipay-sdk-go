package advance

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service 封装同一业务组的支付宝 V3 接口，复用注入执行器的通信与鉴权能力。
// 零值不可用；请通过 NewService 构建。
type Service struct {
	executor       core.Executor
	advanceConsult core.Operation[AdvanceConsultRequest, AdvanceConsultResponse]
}

// NewService 使用共享的 V3 执行器构建 Service；executor 为 nil 时返回错误。
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	advanceConsult, err := core.NewJSONOperation[AdvanceConsultRequest, AdvanceConsultResponse](core.OperationSpec{
		Method: "alipay.trade.advance.consult", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/trade/advance/consult",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return &Service{
		executor:       executor,
		advanceConsult: advanceConsult,
	}, nil
}

// AdvanceConsult 调用 alipay.trade.advance.consult 进行订单咨询
// （V3 GET /v3/alipay/trade/advance/consult）。
//
// 契约：官方 V3 文档
// https://opendocs.alipay.com/open-v3/11cb1f6e_alipay.trade.advance.consult
// 垫资咨询场景 AlipayUserID 必选；风险预评估场景需 ConsultScene 与
// AgreementNo 等参数。
func (s *Service) AdvanceConsult(ctx context.Context, req AdvanceConsultRequest,
	opts ...core.CallOption) (*core.Response[AdvanceConsultResponse], error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, s.advanceConsult,
		core.Request[AdvanceConsultRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
