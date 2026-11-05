// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package role

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

// CcmRoleCreate calls alipay.iservice.ccm.role.create (POST /v3/alipay/iservice/ccm/role/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/632638c3_alipay.iservice.ccm.role.create
func (s *Service) CcmRoleCreate(ctx context.Context, req CcmRoleCreateRequest, opts ...core.CallOption) (*core.Response[CcmRoleCreateResponse], error) {
	return core.ExecuteJSON[CcmRoleCreateRequest, CcmRoleCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.role.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/role/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmRoleCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmRoleDelete calls alipay.iservice.ccm.role.delete (DELETE /v3/alipay/iservice/ccm/role/delete).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8a5f6702_alipay.iservice.ccm.role.delete
func (s *Service) CcmRoleDelete(ctx context.Context, req CcmRoleDeleteRequest, opts ...core.CallOption) (*core.Response[CcmRoleDeleteResponse], error) {
	return core.ExecuteJSON[CcmRoleDeleteRequest, CcmRoleDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.role.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/iservice/ccm/role/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmRoleDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmRoleGet calls alipay.iservice.ccm.role.get (GET /v3/alipay/iservice/ccm/role/get).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvxp
func (s *Service) CcmRoleGet(ctx context.Context, req CcmRoleGetRequest, opts ...core.CallOption) (*core.Response[CcmRoleGetResponse], error) {
	return core.ExecuteJSON[CcmRoleGetRequest, CcmRoleGetResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.role.get", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/role/get",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmRoleGetRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmRoleModify calls alipay.iservice.ccm.role.modify (POST /v3/alipay/iservice/ccm/role/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/8e38792c_alipay.iservice.ccm.role.modify
func (s *Service) CcmRoleModify(ctx context.Context, req CcmRoleModifyRequest, opts ...core.CallOption) (*core.Response[CcmRoleModifyResponse], error) {
	return core.ExecuteJSON[CcmRoleModifyRequest, CcmRoleModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.role.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/role/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmRoleModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmRolePageQuery calls alipay.iservice.ccm.role.page.query (GET /v3/alipay/iservice/ccm/role/page/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/0c530d84_alipay.iservice.ccm.role.page.query
func (s *Service) CcmRolePageQuery(ctx context.Context, req CcmRolePageQueryRequest, opts ...core.CallOption) (*core.Response[CcmRolePageQueryResponse], error) {
	return core.ExecuteJSON[CcmRolePageQueryRequest, CcmRolePageQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.role.page.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/role/page/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmRolePageQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
