// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package info

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

// PublicInfoModify calls alipay.open.public.info.modify (POST /v3/alipay/open/public/info/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wz9l
func (s *Service) PublicInfoModify(ctx context.Context, req PublicInfoModifyRequest, opts ...core.CallOption) (*core.Response[PublicInfoModifyResponse], error) {
	return core.ExecuteJSON[PublicInfoModifyRequest, PublicInfoModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.info.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/info/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicInfoModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicInfoQuery calls alipay.open.public.info.query (GET /v3/alipay/open/public/info/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wygh
func (s *Service) PublicInfoQuery(ctx context.Context, req PublicInfoQueryRequest, opts ...core.CallOption) (*core.Response[PublicInfoQueryResponse], error) {
	return core.ExecuteJSON[PublicInfoQueryRequest, PublicInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.info.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/open/public/info/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
