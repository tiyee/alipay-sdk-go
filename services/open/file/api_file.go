// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package file

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

// FileUpload calls alipay.open.file.upload (POST /v3/alipay/open/file/upload).
// Official V3 contract: https://opendocs.alipay.com/open-v3/5aa91070_alipay.open.file.upload?scene=common&pathHash=e40a9831
func (s *Service) FileUpload(ctx context.Context, req FileUploadRequest, opts ...core.CallOption) (*core.Response[FileUploadResponse], error) {
	return core.ExecuteJSON[FileUploadRequest, FileUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.file.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/file/upload",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FileUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
