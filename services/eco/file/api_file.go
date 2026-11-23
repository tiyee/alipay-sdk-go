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

// FilePathQuery calls alipay.eco.file.path.query (GET /v3/alipay/eco/file/path/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wghs
func (s *Service) FilePathQuery(ctx context.Context, req FilePathQueryRequest, opts ...core.CallOption) (*core.Response[FilePathQueryResponse], error) {
	return core.ExecuteJSON[FilePathQueryRequest, FilePathQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.file.path.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/eco/file/path/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FilePathQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
