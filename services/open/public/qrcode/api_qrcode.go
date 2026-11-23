// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package qrcode

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

// PublicQRCodeCreate calls alipay.open.public.qrcode.create (POST /v3/alipay/open/public/qrcode/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuhi
func (s *Service) PublicQRCodeCreate(ctx context.Context, req PublicQRCodeCreateRequest, opts ...core.CallOption) (*core.Response[PublicQRCodeCreateResponse], error) {
	return core.ExecuteJSON[PublicQRCodeCreateRequest, PublicQRCodeCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.qrcode.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/qrcode/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicQRCodeCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
