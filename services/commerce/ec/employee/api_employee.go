package employee

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

// EcEmployeeAdd calls alipay.commerce.ec.employee.add (POST /v3/alipay/commerce/ec/employee/add).
// Official V3 contract: https://opendocs.alipay.com/apis/05lo9i
func (s *Service) EcEmployeeAdd(ctx context.Context, req EcEmployeeAddRequest, opts ...core.CallOption) (*core.Response[EcEmployeeAddResponse], error) {
	return core.ExecuteJSON[EcEmployeeAddRequest, EcEmployeeAddResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.employee.add", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/ec/employee/add",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEmployeeAddRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEmployeeDelete calls alipay.commerce.ec.employee.delete (POST /v3/alipay/commerce/ec/employee/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05j06m
func (s *Service) EcEmployeeDelete(ctx context.Context, req EcEmployeeDeleteRequest, opts ...core.CallOption) (*core.Response[EcEmployeeDeleteResponse], error) {
	return core.ExecuteJSON[EcEmployeeDeleteRequest, EcEmployeeDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.employee.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/ec/employee/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEmployeeDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEmployeeIDListQuery calls alipay.commerce.ec.employee.idlist.query (GET /v3/alipay/commerce/ec/employee/idlist/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05udr4
func (s *Service) EcEmployeeIDListQuery(ctx context.Context, req EcEmployeeIDListQueryRequest, opts ...core.CallOption) (*core.Response[EcEmployeeIDListQueryResponse], error) {
	return core.ExecuteJSON[EcEmployeeIDListQueryRequest, EcEmployeeIDListQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.employee.idlist.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/employee/idlist/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEmployeeIDListQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEmployeeInfoModify calls alipay.commerce.ec.employee.info.modify (PUT /v3/alipay/commerce/ec/employee/info).
// Official V3 contract: https://opendocs.alipay.com/apis/05j1c0
func (s *Service) EcEmployeeInfoModify(ctx context.Context, req EcEmployeeInfoModifyRequest, opts ...core.CallOption) (*core.Response[EcEmployeeInfoModifyResponse], error) {
	return core.ExecuteJSON[EcEmployeeInfoModifyRequest, EcEmployeeInfoModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.employee.info.modify", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/commerce/ec/employee/info",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEmployeeInfoModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEmployeeInfoQuery calls alipay.commerce.ec.employee.info.query (GET /v3/alipay/commerce/ec/employee/info/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05uhw6
func (s *Service) EcEmployeeInfoQuery(ctx context.Context, req EcEmployeeInfoQueryRequest, opts ...core.CallOption) (*core.Response[EcEmployeeInfoQueryResponse], error) {
	return core.ExecuteJSON[EcEmployeeInfoQueryRequest, EcEmployeeInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.employee.info.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/employee/info/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEmployeeInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEmployeeInviteQuery calls alipay.commerce.ec.employee.invite.query (GET /v3/alipay/commerce/ec/employee/invite).
// Official V3 contract: https://opendocs.alipay.com/apis/05j9je
func (s *Service) EcEmployeeInviteQuery(ctx context.Context, req EcEmployeeInviteQueryRequest, opts ...core.CallOption) (*core.Response[EcEmployeeInviteQueryResponse], error) {
	return core.ExecuteJSON[EcEmployeeInviteQueryRequest, EcEmployeeInviteQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.employee.invite.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/employee/invite",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEmployeeInviteQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEmployeeTitleCreate calls alipay.commerce.ec.employee.title.create (POST /v3/alipay/commerce/ec/employee/title).
// Official V3 contract: https://opendocs.alipay.com/apis/05ip0u
func (s *Service) EcEmployeeTitleCreate(ctx context.Context, req EcEmployeeTitleCreateRequest, opts ...core.CallOption) (*core.Response[EcEmployeeTitleCreateResponse], error) {
	return core.ExecuteJSON[EcEmployeeTitleCreateRequest, EcEmployeeTitleCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.employee.title.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/ec/employee/title",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEmployeeTitleCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEmployeeTitleDelete calls alipay.commerce.ec.employee.title.delete (POST /v3/alipay/commerce/ec/employee/title/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05imon
func (s *Service) EcEmployeeTitleDelete(ctx context.Context, req EcEmployeeTitleDeleteRequest, opts ...core.CallOption) (*core.Response[EcEmployeeTitleDeleteResponse], error) {
	return core.ExecuteJSON[EcEmployeeTitleDeleteRequest, EcEmployeeTitleDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.employee.title.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/ec/employee/title/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEmployeeTitleDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEmployeeTitleModify calls alipay.commerce.ec.employee.title.modify (PUT /v3/alipay/commerce/ec/employee/title).
// Official V3 contract: https://opendocs.alipay.com/apis/05j4cg
func (s *Service) EcEmployeeTitleModify(ctx context.Context, req EcEmployeeTitleModifyRequest, opts ...core.CallOption) (*core.Response[EcEmployeeTitleModifyResponse], error) {
	return core.ExecuteJSON[EcEmployeeTitleModifyRequest, EcEmployeeTitleModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.employee.title.modify", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/commerce/ec/employee/title",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEmployeeTitleModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
