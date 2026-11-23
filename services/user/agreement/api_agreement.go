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

// AgreementExecutionPlanModify calls alipay.user.agreement.executionplan.modify (POST /v3/alipay/user/agreement/executionplan/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/21227669_alipay.user.agreement.executionplan.modify?scene=common&pathHash=89395d7c
func (s *Service) AgreementExecutionPlanModify(ctx context.Context, req UserAgreementExecutionPlanModifyRequest, opts ...core.CallOption) (*core.Response[UserAgreementExecutionPlanModifyResponse], error) {
	return core.ExecuteJSON[UserAgreementExecutionPlanModifyRequest, UserAgreementExecutionPlanModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.agreement.executionplan.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/agreement/executionplan/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserAgreementExecutionPlanModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgreementMigrate calls alipay.user.agreement.migrate (POST /v3/alipay/user/agreement/migrate).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a611f0a3_alipay.user.agreement.migrate
func (s *Service) AgreementMigrate(ctx context.Context, req UserAgreementMigrateRequest, opts ...core.CallOption) (*core.Response[UserAgreementMigrateResponse], error) {
	return core.ExecuteJSON[UserAgreementMigrateRequest, UserAgreementMigrateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.agreement.migrate", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/agreement/migrate",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserAgreementMigrateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgreementPermissionCreate calls alipay.user.agreement.permission.create (POST /v3/alipay/user/agreement/permission/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/87401d30_alipay.user.agreement.permission.create
func (s *Service) AgreementPermissionCreate(ctx context.Context, req UserAgreementPermissionCreateRequest, opts ...core.CallOption) (*core.Response[UserAgreementPermissionCreateResponse], error) {
	return core.ExecuteJSON[UserAgreementPermissionCreateRequest, UserAgreementPermissionCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.agreement.permission.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/agreement/permission/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserAgreementPermissionCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgreementQuery calls alipay.user.agreement.query (GET /v3/alipay/user/agreement/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/4dd54af2_alipay.user.agreement.query
func (s *Service) AgreementQuery(ctx context.Context, req UserAgreementQueryRequest, opts ...core.CallOption) (*core.Response[UserAgreementQueryResponse], error) {
	return core.ExecuteJSON[UserAgreementQueryRequest, UserAgreementQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.agreement.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/user/agreement/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserAgreementQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgreementTransfer calls alipay.user.agreement.transfer (POST /v3/alipay/user/agreement/transfer).
// Official V3 contract: https://opendocs.alipay.com/open-v3/360a356d_alipay.user.agreement.transfer
func (s *Service) AgreementTransfer(ctx context.Context, req UserAgreementTransferRequest, opts ...core.CallOption) (*core.Response[UserAgreementTransferResponse], error) {
	return core.ExecuteJSON[UserAgreementTransferRequest, UserAgreementTransferResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.agreement.transfer", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/agreement/transfer",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserAgreementTransferRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgreementUnsign calls alipay.user.agreement.unsign (POST /v3/alipay/user/agreement/unsign).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b16bea7c_alipay.user.agreement.unsign
func (s *Service) AgreementUnsign(ctx context.Context, req UserAgreementUnsignRequest, opts ...core.CallOption) (*core.Response[UserAgreementUnsignResponse], error) {
	return core.ExecuteJSON[UserAgreementUnsignRequest, UserAgreementUnsignResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.agreement.unsign", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/agreement/unsign",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserAgreementUnsignRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgreementUserVerifyApply calls alipay.user.agreement.userverify.apply (POST /v3/alipay/user/agreement/userverify/apply).
// Official V3 contract: https://opendocs.alipay.com/open-v3/df28f556_alipay.user.agreement.userverify.apply
func (s *Service) AgreementUserVerifyApply(ctx context.Context, req UserAgreementUserVerifyApplyRequest, opts ...core.CallOption) (*core.Response[UserAgreementUserVerifyApplyResponse], error) {
	return core.ExecuteJSON[UserAgreementUserVerifyApplyRequest, UserAgreementUserVerifyApplyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.agreement.userverify.apply", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/user/agreement/userverify/apply",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserAgreementUserVerifyApplyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// AgreementUserVerifyQuery calls alipay.user.agreement.userverify.query (GET /v3/alipay/user/agreement/userverify/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a514c337_alipay.user.agreement.userverify.query
func (s *Service) AgreementUserVerifyQuery(ctx context.Context, req UserAgreementUserVerifyQueryRequest, opts ...core.CallOption) (*core.Response[UserAgreementUserVerifyQueryResponse], error) {
	return core.ExecuteJSON[UserAgreementUserVerifyQueryRequest, UserAgreementUserVerifyQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.user.agreement.userverify.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/user/agreement/userverify/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[UserAgreementUserVerifyQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
