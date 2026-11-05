// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package material

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

// MaterialCreate calls alipay.marketing.material.create (POST /v3/alipay/marketing/material/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05xk1p
func (s *Service) MaterialCreate(ctx context.Context, req MaterialCreateRequest, opts ...core.CallOption) (*core.Response[MaterialCreateResponse], error) {
	return core.ExecuteJSON[MaterialCreateRequest, MaterialCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.material.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/material/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MaterialCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MaterialImageUpload calls alipay.marketing.material.image.upload (POST /v3/alipay/marketing/material/image).
// Official V3 contract: https://opendocs.alipay.com/open-v3/574939dd_alipay.marketing.material.image.upload
func (s *Service) MaterialImageUpload(ctx context.Context, req MaterialImageUploadRequest, opts ...core.CallOption) (*core.Response[MaterialImageUploadResponse], error) {
	return core.ExecuteJSON[MaterialImageUploadRequest, MaterialImageUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.material.image.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/material/image",
		ParameterMode: core.ParameterModeMultipart, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MaterialImageUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
