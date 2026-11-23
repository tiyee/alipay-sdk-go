package enterprisepay

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

// EnterprisePayGroupAdd calls alipay.fund.enterprisepay.group.add (POST /v3/alipay/fund/enterprisepay/group/add).
// Official V3 contract: https://opendocs.alipay.com/apis/05wmyu
func (s *Service) EnterprisePayGroupAdd(ctx context.Context, req FundEnterprisePayGroupAddRequest, opts ...core.CallOption) (*core.Response[FundEnterprisePayGroupAddResponse], error) {
	return core.ExecuteJSON[FundEnterprisePayGroupAddRequest, FundEnterprisePayGroupAddResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.enterprisepay.group.add", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/enterprisepay/group/add",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundEnterprisePayGroupAddRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EnterprisePayGroupDelete calls alipay.fund.enterprisepay.group.delete (DELETE /v3/alipay/fund/enterprisepay/group/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvyb
func (s *Service) EnterprisePayGroupDelete(ctx context.Context, req FundEnterprisePayGroupDeleteRequest, opts ...core.CallOption) (*core.Response[FundEnterprisePayGroupDeleteResponse], error) {
	return core.ExecuteJSON[FundEnterprisePayGroupDeleteRequest, FundEnterprisePayGroupDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.enterprisepay.group.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/fund/enterprisepay/group/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundEnterprisePayGroupDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EnterprisePayGroupModify calls alipay.fund.enterprisepay.group.modify (POST /v3/alipay/fund/enterprisepay/group/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wmzb
func (s *Service) EnterprisePayGroupModify(ctx context.Context, req FundEnterprisePayGroupModifyRequest, opts ...core.CallOption) (*core.Response[FundEnterprisePayGroupModifyResponse], error) {
	return core.ExecuteJSON[FundEnterprisePayGroupModifyRequest, FundEnterprisePayGroupModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.enterprisepay.group.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/enterprisepay/group/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundEnterprisePayGroupModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EnterprisePayGroupQuery calls alipay.fund.enterprisepay.group.query (GET /v3/alipay/fund/enterprisepay/group/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wcpq
func (s *Service) EnterprisePayGroupQuery(ctx context.Context, req FundEnterprisePayGroupQueryRequest, opts ...core.CallOption) (*core.Response[FundEnterprisePayGroupQueryResponse], error) {
	return core.ExecuteJSON[FundEnterprisePayGroupQueryRequest, FundEnterprisePayGroupQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.enterprisepay.group.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/enterprisepay/group/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundEnterprisePayGroupQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EnterprisePayMemberModify calls alipay.fund.enterprisepay.member.modify (POST /v3/alipay/fund/enterprisepay/member/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvv1
func (s *Service) EnterprisePayMemberModify(ctx context.Context, req FundEnterprisePayMemberModifyRequest, opts ...core.CallOption) (*core.Response[FundEnterprisePayMemberModifyResponse], error) {
	return core.ExecuteJSON[FundEnterprisePayMemberModifyRequest, FundEnterprisePayMemberModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.enterprisepay.member.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/enterprisepay/member/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundEnterprisePayMemberModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EnterprisePaySign calls alipay.fund.enterprisepay.sign (POST /v3/alipay/fund/enterprisepay/sign).
// Official V3 contract: https://opendocs.alipay.com/apis/05udr5
func (s *Service) EnterprisePaySign(ctx context.Context, req FundEnterprisePaySignRequest, opts ...core.CallOption) (*core.Response[FundEnterprisePaySignResponse], error) {
	return core.ExecuteJSON[FundEnterprisePaySignRequest, FundEnterprisePaySignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.enterprisepay.sign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/enterprisepay/sign",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundEnterprisePaySignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EnterprisePayUnsign calls alipay.fund.enterprisepay.unsign (POST /v3/alipay/fund/enterprisepay/unsign).
// Official V3 contract: https://opendocs.alipay.com/apis/05wbw2
func (s *Service) EnterprisePayUnsign(ctx context.Context, req FundEnterprisePayUnsignRequest, opts ...core.CallOption) (*core.Response[FundEnterprisePayUnsignResponse], error) {
	return core.ExecuteJSON[FundEnterprisePayUnsignRequest, FundEnterprisePayUnsignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.enterprisepay.unsign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/enterprisepay/unsign",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundEnterprisePayUnsignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
