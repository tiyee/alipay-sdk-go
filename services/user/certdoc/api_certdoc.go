package certdoc

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

// CertDocCertVerifyConsult calls alipay.user.certdoc.certverify.consult (GET /v3/alipay/user/certdoc/certverify/consult).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvw5
func (s *Service) CertDocCertVerifyConsult(ctx context.Context, req UserCertDocCertVerifyConsultRequest, opts ...core.CallOption) (*core.Response[UserCertDocCertVerifyConsultResponse], error) {
	return core.ExecuteJSON[UserCertDocCertVerifyConsultRequest, UserCertDocCertVerifyConsultResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.certdoc.certverify.consult", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/user/certdoc/certverify/consult",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserCertDocCertVerifyConsultRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CertDocCertVerifyPreConsult calls alipay.user.certdoc.certverify.preconsult (POST /v3/alipay/user/certdoc/certverify/preconsult).
// Official V3 contract: https://opendocs.alipay.com/apis/05wfau
func (s *Service) CertDocCertVerifyPreConsult(ctx context.Context, req UserCertDocCertVerifyPreConsultRequest, opts ...core.CallOption) (*core.Response[UserCertDocCertVerifyPreConsultResponse], error) {
	return core.ExecuteJSON[UserCertDocCertVerifyPreConsultRequest, UserCertDocCertVerifyPreConsultResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.certdoc.certverify.preconsult", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/certdoc/certverify/preconsult",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserCertDocCertVerifyPreConsultRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
