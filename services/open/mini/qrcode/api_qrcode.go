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

// MiniQRCodeBind calls alipay.open.mini.qrcode.bind (POST /v3/alipay/open/mini/qrcode/bind).
// Official V3 contract: https://opendocs.alipay.com/open-v3/01ebf6a3_alipay.open.mini.qrcode.bind
func (s *Service) MiniQRCodeBind(ctx context.Context, req MiniQRCodeBindRequest, opts ...core.CallOption) (*core.Response[MiniQRCodeBindResponse], error) {
	return core.ExecuteJSON[MiniQRCodeBindRequest, MiniQRCodeBindResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.qrcode.bind", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/qrcode/bind",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniQRCodeBindRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniQRCodeQuery calls alipay.open.mini.qrcode.query (GET /v3/alipay/open/mini/qrcode/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c89a2391_alipay.open.mini.qrcode.query
func (s *Service) MiniQRCodeQuery(ctx context.Context, req MiniQRCodeQueryRequest, opts ...core.CallOption) (*core.Response[MiniQRCodeQueryResponse], error) {
	return core.ExecuteJSON[MiniQRCodeQueryRequest, MiniQRCodeQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.qrcode.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/qrcode/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniQRCodeQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniQRCodeUnbind calls alipay.open.mini.qrcode.unbind (POST /v3/alipay/open/mini/qrcode/unbind).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8905c9d2_alipay.open.mini.qrcode.unbind
func (s *Service) MiniQRCodeUnbind(ctx context.Context, req MiniQRCodeUnbindRequest, opts ...core.CallOption) (*core.Response[MiniQRCodeUnbindResponse], error) {
	return core.ExecuteJSON[MiniQRCodeUnbindRequest, MiniQRCodeUnbindResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.qrcode.unbind", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/qrcode/unbind",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniQRCodeUnbindRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
