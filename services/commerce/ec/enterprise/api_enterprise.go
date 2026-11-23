package enterprise

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

// EcEnterpriseAddressAdd calls alipay.commerce.ec.enterprise.address.add (POST /v3/alipay/commerce/ec/enterprise/address).
// Official V3 contract: https://opendocs.alipay.com/apis/05j04z
func (s *Service) EcEnterpriseAddressAdd(ctx context.Context, req EcEnterpriseAddressAddRequest, opts ...core.CallOption) (*core.Response[EcEnterpriseAddressAddResponse], error) {
	return core.ExecuteJSON[EcEnterpriseAddressAddRequest, EcEnterpriseAddressAddResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.enterprise.address.add", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/ec/enterprise/address",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEnterpriseAddressAddRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEnterpriseAddressModify calls alipay.commerce.ec.enterprise.address.modify (PUT /v3/alipay/commerce/ec/enterprise/address).
// Official V3 contract: https://opendocs.alipay.com/apis/05j1bb
func (s *Service) EcEnterpriseAddressModify(ctx context.Context, req EcEnterpriseAddressModifyRequest, opts ...core.CallOption) (*core.Response[EcEnterpriseAddressModifyResponse], error) {
	return core.ExecuteJSON[EcEnterpriseAddressModifyRequest, EcEnterpriseAddressModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.enterprise.address.modify", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/commerce/ec/enterprise/address",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEnterpriseAddressModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEnterpriseAddressQuery calls alipay.commerce.ec.enterprise.address.query (GET /v3/alipay/commerce/ec/enterprise/address).
// Official V3 contract: https://opendocs.alipay.com/apis/05iz2p
func (s *Service) EcEnterpriseAddressQuery(ctx context.Context, req EcEnterpriseAddressQueryRequest, opts ...core.CallOption) (*core.Response[EcEnterpriseAddressQueryResponse], error) {
	return core.ExecuteJSON[EcEnterpriseAddressQueryRequest, EcEnterpriseAddressQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.enterprise.address.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/enterprise/address",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEnterpriseAddressQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEnterpriseAgreementQuery calls alipay.commerce.ec.enterprise.agreement.query (GET /v3/alipay/commerce/ec/enterprise/agreement).
// Official V3 contract: https://opendocs.alipay.com/apis/05j2cw
func (s *Service) EcEnterpriseAgreementQuery(ctx context.Context, req EcEnterpriseAgreementQueryRequest, opts ...core.CallOption) (*core.Response[EcEnterpriseAgreementQueryResponse], error) {
	return core.ExecuteJSON[EcEnterpriseAgreementQueryRequest, EcEnterpriseAgreementQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.enterprise.agreement.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/enterprise/agreement",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEnterpriseAgreementQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEnterpriseCreate calls alipay.commerce.ec.enterprise.create (POST /v3/alipay/commerce/ec/enterprise/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05l7n7
func (s *Service) EcEnterpriseCreate(ctx context.Context, req EcEnterpriseCreateRequest, opts ...core.CallOption) (*core.Response[EcEnterpriseCreateResponse], error) {
	return core.ExecuteJSON[EcEnterpriseCreateRequest, EcEnterpriseCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.enterprise.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/ec/enterprise/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEnterpriseCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEnterpriseDelete calls alipay.commerce.ec.enterprise.delete (POST /v3/alipay/commerce/ec/enterprise/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05j82m
func (s *Service) EcEnterpriseDelete(ctx context.Context, req EcEnterpriseDeleteRequest, opts ...core.CallOption) (*core.Response[EcEnterpriseDeleteResponse], error) {
	return core.ExecuteJSON[EcEnterpriseDeleteRequest, EcEnterpriseDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.enterprise.delete", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/ec/enterprise/delete",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEnterpriseDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEnterpriseInfoModify calls alipay.commerce.ec.enterprise.info.modify (PUT /v3/alipay/commerce/ec/enterprise).
// Official V3 contract: https://opendocs.alipay.com/apis/05ip0v
func (s *Service) EcEnterpriseInfoModify(ctx context.Context, req EcEnterpriseInfoModifyRequest, opts ...core.CallOption) (*core.Response[EcEnterpriseInfoModifyResponse], error) {
	return core.ExecuteJSON[EcEnterpriseInfoModifyRequest, EcEnterpriseInfoModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.enterprise.info.modify", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/commerce/ec/enterprise",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEnterpriseInfoModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEnterpriseInfoQuery calls alipay.commerce.ec.enterprise.info.query (GET /v3/alipay/commerce/ec/enterprise/info).
// Official V3 contract: https://opendocs.alipay.com/apis/05iwc6
func (s *Service) EcEnterpriseInfoQuery(ctx context.Context, req EcEnterpriseInfoQueryRequest, opts ...core.CallOption) (*core.Response[EcEnterpriseInfoQueryResponse], error) {
	return core.ExecuteJSON[EcEnterpriseInfoQueryRequest, EcEnterpriseInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.enterprise.info.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/enterprise/info",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEnterpriseInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EcEnterpriseUnsign calls alipay.commerce.ec.enterprise.unsign (GET /v3/alipay/commerce/ec/enterprise/unsignurl).
// Official V3 contract: https://opendocs.alipay.com/apis/05jach
func (s *Service) EcEnterpriseUnsign(ctx context.Context, req EcEnterpriseUnsignRequest, opts ...core.CallOption) (*core.Response[EcEnterpriseUnsignResponse], error) {
	return core.ExecuteJSON[EcEnterpriseUnsignRequest, EcEnterpriseUnsignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.ec.enterprise.unsign", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/ec/enterprise/unsignurl",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EcEnterpriseUnsignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
