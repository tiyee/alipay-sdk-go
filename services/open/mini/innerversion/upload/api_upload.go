package upload

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

// MiniInnerVersionNoBuildUpload calls alipay.open.mini.innerversion.nobuild.upload (POST /v3/alipay/open/mini/innerversion/nobuild/upload).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xiqc
func (s *Service) MiniInnerVersionNoBuildUpload(ctx context.Context, req MiniInnerVersionNoBuildUploadRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionNoBuildUploadResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionNoBuildUploadRequest, MiniInnerVersionNoBuildUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.nobuild.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/nobuild/upload",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionNoBuildUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionTemplatedUpload calls alipay.open.mini.innerversion.templated.upload (POST /v3/alipay/open/mini/innerversion/templated/upload).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xdzq
func (s *Service) MiniInnerVersionTemplatedUpload(ctx context.Context, req MiniInnerVersionTemplatedUploadRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionTemplatedUploadResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionTemplatedUploadRequest, MiniInnerVersionTemplatedUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.templated.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/templated/upload",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionTemplatedUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionUpload calls alipay.open.mini.innerversion.upload (POST /v3/alipay/open/mini/innerversion/upload).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xhot
func (s *Service) MiniInnerVersionUpload(ctx context.Context, req MiniInnerVersionUploadRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionUploadResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionUploadRequest, MiniInnerVersionUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/innerversion/upload",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniInnerVersionUploadStatusQuery calls alipay.open.mini.innerversion.uploadstatus.query (GET /v3/alipay/open/mini/innerversion/uploadstatus/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xskn
func (s *Service) MiniInnerVersionUploadStatusQuery(ctx context.Context, req MiniInnerVersionUploadStatusQueryRequest, opts ...core.CallOption) (*core.Response[MiniInnerVersionUploadStatusQueryResponse], error) {
	return core.ExecuteJSON[MiniInnerVersionUploadStatusQueryRequest, MiniInnerVersionUploadStatusQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.innerversion.uploadstatus.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/mini/innerversion/uploadstatus/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniInnerVersionUploadStatusQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
