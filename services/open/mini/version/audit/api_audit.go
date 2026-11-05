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

// MiniVersionAuditApply calls alipay.open.mini.version.audit.apply (POST /v3/alipay/open/mini/version/audit/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/dd9959d2_alipay.open.mini.version.audit.apply
func (s *Service) MiniVersionAuditApply(ctx context.Context, req MiniVersionAuditApplyRequest, opts ...core.CallOption) (*core.Response[MiniVersionAuditApplyResponse], error) {
	return core.ExecuteJSON[MiniVersionAuditApplyRequest, MiniVersionAuditApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.audit.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/version/audit/apply",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionAuditApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniVersionAuditCancel calls alipay.open.mini.version.audit.cancel (POST /v3/alipay/open/mini/version/audit/cancel).
// Official V3 contract: https://opendocs.alipay.com/open-v3/9b50e79c_alipay.open.mini.version.audit.cancel
func (s *Service) MiniVersionAuditCancel(ctx context.Context, req MiniVersionAuditCancelRequest, opts ...core.CallOption) (*core.Response[MiniVersionAuditCancelResponse], error) {
	return core.ExecuteJSON[MiniVersionAuditCancelRequest, MiniVersionAuditCancelResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.version.audit.cancel", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/version/audit/cancel",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniVersionAuditCancelRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
