// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package instance

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

// CcmInstanceCreate calls alipay.iservice.ccm.instance.create (POST /v3/alipay/iservice/ccm/instance/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/d2fd8b94_alipay.iservice.ccm.instance.create
func (s *Service) CcmInstanceCreate(ctx context.Context, req CcmInstanceCreateRequest, opts ...core.CallOption) (*core.Response[CcmInstanceCreateResponse], error) {
	return core.ExecuteJSON[CcmInstanceCreateRequest, CcmInstanceCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.instance.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/instance/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmInstanceCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmInstanceGet calls alipay.iservice.ccm.instance.get (GET /v3/alipay/iservice/ccm/instance/get).
// Official V3 contract: https://opendocs.alipay.com/open-v3/e02604d3_alipay.iservice.ccm.instance.get
func (s *Service) CcmInstanceGet(ctx context.Context, req CcmInstanceGetRequest, opts ...core.CallOption) (*core.Response[CcmInstanceGetResponse], error) {
	return core.ExecuteJSON[CcmInstanceGetRequest, CcmInstanceGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.instance.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/instance/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmInstanceGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmInstanceQuery calls alipay.iservice.ccm.instance.query (GET /v3/alipay/iservice/ccm/instance/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/eb6f9aca_alipay.iservice.ccm.instance.query
func (s *Service) CcmInstanceQuery(ctx context.Context, req CcmInstanceQueryRequest, opts ...core.CallOption) (*core.Response[CcmInstanceQueryResponse], error) {
	return core.ExecuteJSON[CcmInstanceQueryRequest, CcmInstanceQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.instance.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/instance/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmInstanceQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
