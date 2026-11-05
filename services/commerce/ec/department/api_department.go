package department

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

// EcDepartmentCreate calls alipay.commerce.ec.department.create (POST /v3/alipay/commerce/ec/department).
// Official V3 contract: https://opendocs.alipay.com/apis/05itkw
func (s *Service) EcDepartmentCreate(ctx context.Context, req EcDepartmentCreateRequest, opts ...core.CallOption) (*core.Response[EcDepartmentCreateResponse], error) {
	return core.ExecuteJSON[EcDepartmentCreateRequest, EcDepartmentCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.department.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/ec/department",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcDepartmentCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcDepartmentDelete calls alipay.commerce.ec.department.delete (POST /v3/alipay/commerce/ec/department/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05jgvm
func (s *Service) EcDepartmentDelete(ctx context.Context, req EcDepartmentDeleteRequest, opts ...core.CallOption) (*core.Response[EcDepartmentDeleteResponse], error) {
	return core.ExecuteJSON[EcDepartmentDeleteRequest, EcDepartmentDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.department.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/ec/department/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcDepartmentDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcDepartmentInfoModify calls alipay.commerce.ec.department.info.modify (PUT /v3/alipay/commerce/ec/department).
// Official V3 contract: https://opendocs.alipay.com/apis/05iwe0
func (s *Service) EcDepartmentInfoModify(ctx context.Context, req EcDepartmentInfoModifyRequest, opts ...core.CallOption) (*core.Response[EcDepartmentInfoModifyResponse], error) {
	return core.ExecuteJSON[EcDepartmentInfoModifyRequest, EcDepartmentInfoModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.department.info.modify", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/commerce/ec/department",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcDepartmentInfoModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcDepartmentInfoQuery calls alipay.commerce.ec.department.info.query (GET /v3/alipay/commerce/ec/department/info/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05uggp
func (s *Service) EcDepartmentInfoQuery(ctx context.Context, req EcDepartmentInfoQueryRequest, opts ...core.CallOption) (*core.Response[EcDepartmentInfoQueryResponse], error) {
	return core.ExecuteJSON[EcDepartmentInfoQueryRequest, EcDepartmentInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.department.info.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/department/info/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcDepartmentInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcDepartmentSubListQuery calls alipay.commerce.ec.department.sublist.query (GET /v3/alipay/commerce/ec/department/sublist/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05uipk
func (s *Service) EcDepartmentSubListQuery(ctx context.Context, req EcDepartmentSubListQueryRequest, opts ...core.CallOption) (*core.Response[EcDepartmentSubListQueryResponse], error) {
	return core.ExecuteJSON[EcDepartmentSubListQueryRequest, EcDepartmentSubListQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.department.sublist.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/department/sublist/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcDepartmentSubListQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
