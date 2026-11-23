// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package resource

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service executes one alipay V3 service group.
type Service struct{ executor core.Executor }

// NewService constructs the service with a shared V3 executor.
// NewService constructs the service with a shared V3 executor.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

// MiniResourcePromotionSourceNotify calls alipay.open.mini.resource.promotionsource.notify (POST /v3/alipay/open/mini/resource/promotionsource/notify).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xhr9
func (s *Service) MiniResourcePromotionSourceNotify(ctx context.Context, req MiniResourcePromotionSourceNotifyRequest, opts ...core.CallOption) (*core.Response[MiniResourcePromotionSourceNotifyResponse], error) {
	return core.ExecuteJSON[MiniResourcePromotionSourceNotifyRequest, MiniResourcePromotionSourceNotifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.resource.promotionsource.notify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/resource/promotionsource/notify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniResourcePromotionSourceNotifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniResourceRecordNotify calls alipay.open.mini.resource.record.notify (POST /v3/alipay/open/mini/resource/record/notify).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xqyn
func (s *Service) MiniResourceRecordNotify(ctx context.Context, req MiniResourceRecordNotifyRequest, opts ...core.CallOption) (*core.Response[MiniResourceRecordNotifyResponse], error) {
	return core.ExecuteJSON[MiniResourceRecordNotifyRequest, MiniResourceRecordNotifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.resource.record.notify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/resource/record/notify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniResourceRecordNotifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
