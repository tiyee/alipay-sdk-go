package certify

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

// CertifyOpenInitialize calls alipay.user.certify.open.initialize (POST /v3/alipay/user/certify/open/initialize).
// Official V3 contract: https://opendocs.alipay.com/open-v3/013ee3e1_alipay.user.certify.open.initialize
func (s *Service) CertifyOpenInitialize(ctx context.Context, req UserCertifyOpenInitializeRequest, opts ...core.CallOption) (*core.Response[UserCertifyOpenInitializeResponse], error) {
	return core.ExecuteJSON[UserCertifyOpenInitializeRequest, UserCertifyOpenInitializeResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.certify.open.initialize", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/certify/open/initialize",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserCertifyOpenInitializeRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CertifyOpenQuery calls alipay.user.certify.open.query (GET /v3/alipay/user/certify/open/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/c166c117_alipay.user.certify.open.query
func (s *Service) CertifyOpenQuery(ctx context.Context, req UserCertifyOpenQueryRequest, opts ...core.CallOption) (*core.Response[UserCertifyOpenQueryResponse], error) {
	return core.ExecuteJSON[UserCertifyOpenQueryRequest, UserCertifyOpenQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.certify.open.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/user/certify/open/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserCertifyOpenQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
