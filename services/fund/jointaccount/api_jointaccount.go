package jointaccount

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

// JointAccountBillQuery calls alipay.fund.jointaccount.bill.query (GET /v3/alipay/fund/jointaccount/bill/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuhp
func (s *Service) JointAccountBillQuery(ctx context.Context, req FundJointAccountBillQueryRequest, opts ...core.CallOption) (*core.Response[FundJointAccountBillQueryResponse], error) {
	return core.ExecuteJSON[FundJointAccountBillQueryRequest, FundJointAccountBillQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.bill.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/jointaccount/bill/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountBillQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountDetailQuery calls alipay.fund.jointaccount.detail.query (GET /v3/alipay/fund/jointaccount/detail/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05imoi
func (s *Service) JointAccountDetailQuery(ctx context.Context, req FundJointAccountDetailQueryRequest, opts ...core.CallOption) (*core.Response[FundJointAccountDetailQueryResponse], error) {
	return core.ExecuteJSON[FundJointAccountDetailQueryRequest, FundJointAccountDetailQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.detail.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/jointaccount/detail/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountDetailQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountListQuery calls alipay.fund.jointaccount.list.query (GET /v3/alipay/fund/jointaccount/list/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wcqw
func (s *Service) JointAccountListQuery(ctx context.Context, req FundJointAccountListQueryRequest, opts ...core.CallOption) (*core.Response[FundJointAccountListQueryResponse], error) {
	return core.ExecuteJSON[FundJointAccountListQueryRequest, FundJointAccountListQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.list.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/jointaccount/list/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountListQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountMemberBatchQuery calls alipay.fund.jointaccount.member.batchquery (POST /v3/alipay/fund/jointaccount/member/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wfae
func (s *Service) JointAccountMemberBatchQuery(ctx context.Context, req FundJointAccountMemberBatchQueryRequest, opts ...core.CallOption) (*core.Response[FundJointAccountMemberBatchQueryResponse], error) {
	return core.ExecuteJSON[FundJointAccountMemberBatchQueryRequest, FundJointAccountMemberBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.member.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/jointaccount/member/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountMemberBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountMemberBind calls alipay.fund.jointaccount.member.bind (POST /v3/alipay/fund/jointaccount/member/bind).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvt3
func (s *Service) JointAccountMemberBind(ctx context.Context, req FundJointAccountMemberBindRequest, opts ...core.CallOption) (*core.Response[FundJointAccountMemberBindResponse], error) {
	return core.ExecuteJSON[FundJointAccountMemberBindRequest, FundJointAccountMemberBindResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.member.bind", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/jointaccount/member/bind",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountMemberBindRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountMemberConsult calls alipay.fund.jointaccount.member.consult (POST /v3/alipay/fund/jointaccount/member/consult).
// Official V3 contract: https://opendocs.alipay.com/apis/05wupc
func (s *Service) JointAccountMemberConsult(ctx context.Context, req FundJointAccountMemberConsultRequest, opts ...core.CallOption) (*core.Response[FundJointAccountMemberConsultResponse], error) {
	return core.ExecuteJSON[FundJointAccountMemberConsultRequest, FundJointAccountMemberConsultResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.member.consult", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/jointaccount/member/consult",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountMemberConsultRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountMemberQuery calls alipay.fund.jointaccount.member.query (GET /v3/alipay/fund/jointaccount/member/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wujm
func (s *Service) JointAccountMemberQuery(ctx context.Context, req FundJointAccountMemberQueryRequest, opts ...core.CallOption) (*core.Response[FundJointAccountMemberQueryResponse], error) {
	return core.ExecuteJSON[FundJointAccountMemberQueryRequest, FundJointAccountMemberQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.member.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/jointaccount/member/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountMemberQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountMemberUnbind calls alipay.fund.jointaccount.member.unbind (POST /v3/alipay/fund/jointaccount/member/unbind).
// Official V3 contract: https://opendocs.alipay.com/apis/05x1fk
func (s *Service) JointAccountMemberUnbind(ctx context.Context, req FundJointAccountMemberUnbindRequest, opts ...core.CallOption) (*core.Response[FundJointAccountMemberUnbindResponse], error) {
	return core.ExecuteJSON[FundJointAccountMemberUnbindRequest, FundJointAccountMemberUnbindResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.member.unbind", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/jointaccount/member/unbind",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountMemberUnbindRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountMemberRuleModify calls alipay.fund.jointaccount.memberrule.modify (POST /v3/alipay/fund/jointaccount/memberrule/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wygg
func (s *Service) JointAccountMemberRuleModify(ctx context.Context, req FundJointAccountMemberRuleModifyRequest, opts ...core.CallOption) (*core.Response[FundJointAccountMemberRuleModifyResponse], error) {
	return core.ExecuteJSON[FundJointAccountMemberRuleModifyRequest, FundJointAccountMemberRuleModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.memberrule.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/jointaccount/memberrule/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountMemberRuleModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountQuotaQuery calls alipay.fund.jointaccount.quota.query (GET /v3/alipay/fund/jointaccount/quota/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuhk
func (s *Service) JointAccountQuotaQuery(ctx context.Context, req FundJointAccountQuotaQueryRequest, opts ...core.CallOption) (*core.Response[FundJointAccountQuotaQueryResponse], error) {
	return core.ExecuteJSON[FundJointAccountQuotaQueryRequest, FundJointAccountQuotaQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.quota.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/jointaccount/quota/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountQuotaQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountRuleModify calls alipay.fund.jointaccount.rule.modify (POST /v3/alipay/fund/jointaccount/rule/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wghc
func (s *Service) JointAccountRuleModify(ctx context.Context, req FundJointAccountRuleModifyRequest, opts ...core.CallOption) (*core.Response[FundJointAccountRuleModifyResponse], error) {
	return core.ExecuteJSON[FundJointAccountRuleModifyRequest, FundJointAccountRuleModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.rule.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/jointaccount/rule/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountRuleModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountSignQuery calls alipay.fund.jointaccount.sign.query (GET /v3/alipay/fund/jointaccount/sign/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wulj
func (s *Service) JointAccountSignQuery(ctx context.Context, req FundJointAccountSignQueryRequest, opts ...core.CallOption) (*core.Response[FundJointAccountSignQueryResponse], error) {
	return core.ExecuteJSON[FundJointAccountSignQueryRequest, FundJointAccountSignQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.sign.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/jointaccount/sign/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountSignQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// JointAccountTradeQuery calls alipay.fund.jointaccount.trade.query (GET /v3/alipay/fund/jointaccount/trade/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wykp
func (s *Service) JointAccountTradeQuery(ctx context.Context, req FundJointAccountTradeQueryRequest, opts ...core.CallOption) (*core.Response[FundJointAccountTradeQueryResponse], error) {
	return core.ExecuteJSON[FundJointAccountTradeQueryRequest, FundJointAccountTradeQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.jointaccount.trade.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/fund/jointaccount/trade/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundJointAccountTradeQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
