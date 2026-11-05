package agreement

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

// AgreementQuotaModify calls alipay.fund.agreement.quota.modify (POST /v3/alipay/fund/agreement/quota/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05xo4q
func (s *Service) AgreementQuotaModify(ctx context.Context, req FundAgreementQuotaModifyRequest, opts ...core.CallOption) (*core.Response[FundAgreementQuotaModifyResponse], error) {
	return core.ExecuteJSON[FundAgreementQuotaModifyRequest, FundAgreementQuotaModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.agreement.quota.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/agreement/quota/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAgreementQuotaModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgreementQuotaQuery calls alipay.fund.agreement.quota.query (POST /v3/alipay/fund/agreement/quota/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05ximq
func (s *Service) AgreementQuotaQuery(ctx context.Context, req FundAgreementQuotaQueryRequest, opts ...core.CallOption) (*core.Response[FundAgreementQuotaQueryResponse], error) {
	return core.ExecuteJSON[FundAgreementQuotaQueryRequest, FundAgreementQuotaQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.fund.agreement.quota.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/fund/agreement/quota/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[FundAgreementQuotaQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
