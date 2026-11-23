package sign

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

// PdeductSignAdd calls alipay.ebpp.pdeduct.sign.add (POST /v3/alipay/ebpp/pdeduct/sign/add).
// Official V3 contract: https://opendocs.alipay.com/open-v3/084db95a_alipay.ebpp.pdeduct.sign.add
func (s *Service) PdeductSignAdd(ctx context.Context, req PdeductSignAddRequest, opts ...core.CallOption) (*core.Response[PdeductSignAddResponse], error) {
	return core.ExecuteJSON[PdeductSignAddRequest, PdeductSignAddResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.pdeduct.sign.add", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/pdeduct/sign/add",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PdeductSignAddRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PdeductSignCancel calls alipay.ebpp.pdeduct.sign.cancel (POST /v3/alipay/ebpp/pdeduct/sign/cancel).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2ad13974_alipay.ebpp.pdeduct.sign.cancel
func (s *Service) PdeductSignCancel(ctx context.Context, req PdeductSignCancelRequest, opts ...core.CallOption) (*core.Response[PdeductSignCancelResponse], error) {
	return core.ExecuteJSON[PdeductSignCancelRequest, PdeductSignCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.pdeduct.sign.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/pdeduct/sign/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PdeductSignCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PdeductSignQuery calls alipay.ebpp.pdeduct.sign.query (GET /v3/alipay/ebpp/pdeduct/sign/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05x9ew
func (s *Service) PdeductSignQuery(ctx context.Context, req PdeductSignQueryRequest, opts ...core.CallOption) (*core.Response[PdeductSignQueryResponse], error) {
	return core.ExecuteJSON[PdeductSignQueryRequest, PdeductSignQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.pdeduct.sign.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/ebpp/pdeduct/sign/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PdeductSignQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PdeductSignValidate calls alipay.ebpp.pdeduct.sign.validate (POST /v3/alipay/ebpp/pdeduct/sign/validate).
// Official V3 contract: https://opendocs.alipay.com/apis/05wukm
func (s *Service) PdeductSignValidate(ctx context.Context, req PdeductSignValidateRequest, opts ...core.CallOption) (*core.Response[PdeductSignValidateResponse], error) {
	return core.ExecuteJSON[PdeductSignValidateRequest, PdeductSignValidateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.pdeduct.sign.validate", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/pdeduct/sign/validate",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PdeductSignValidateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
