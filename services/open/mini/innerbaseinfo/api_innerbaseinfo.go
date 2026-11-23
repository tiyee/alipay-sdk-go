// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package innerbaseinfo

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

// MiniInnerBaseInfoAppLogoUpload calls alipay.open.mini.innerbaseinfo.applogo.upload (POST /v3/alipay/open/mini/innerbaseinfo/applogo/upload).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xdzx
func (s *Service) MiniInnerBaseInfoAppLogoUpload(ctx context.Context, req MiniInnerBaseInfoAppLogoUploadRequest, opts ...core.CallOption) (*core.Response[MiniInnerBaseInfoAppLogoUploadResponse], error) {
	return core.ExecuteJSON[MiniInnerBaseInfoAppLogoUploadRequest, MiniInnerBaseInfoAppLogoUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerbaseinfo.applogo.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerbaseinfo/applogo/upload",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerBaseInfoAppLogoUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerBaseInfoQuery calls alipay.open.mini.innerbaseinfo.query (GET /v3/alipay/open/mini/innerbaseinfo/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xiqj
func (s *Service) MiniInnerBaseInfoQuery(ctx context.Context, req MiniInnerBaseInfoQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerBaseInfoQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerBaseInfoQueryRequest, MiniInnerBaseInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerbaseinfo.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/innerbaseinfo/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerBaseInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
