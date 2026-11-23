package audit

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

// MiniInnerVersionAuditCancel calls alipay.open.mini.innerversion.audit.cancel (POST /v3/alipay/open/mini/innerversion/audit/cancel).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xaab
func (s *Service) MiniInnerVersionAuditCancel(ctx context.Context, req MiniInnerVersionAuditCancelRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionAuditCancelResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionAuditCancelRequest, MiniInnerVersionAuditCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.audit.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/audit/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionAuditCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionAuditSubmit calls alipay.open.mini.innerversion.audit.submit (POST /v3/alipay/open/mini/innerversion/audit/submit).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xnee
func (s *Service) MiniInnerVersionAuditSubmit(ctx context.Context, req MiniInnerVersionAuditSubmitRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionAuditSubmitResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionAuditSubmitRequest, MiniInnerVersionAuditSubmitResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.audit.submit", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/audit/submit",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionAuditSubmitRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
