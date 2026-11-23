package userinvoiceinfo

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

// FncUserInvoiceInfoCreate calls alipay.boss.fnc.userinvoiceinfo.create (POST /v3/alipay/boss/fnc/userinvoiceinfo/create).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xgfe
func (s *Service) FncUserInvoiceInfoCreate(ctx context.Context, req BossFncUserInvoiceInfoCreateRequest, opts ...core.CallOption) (*core.Response[BossFncUserInvoiceInfoCreateResponse], error) {
	return core.ExecuteJSON[BossFncUserInvoiceInfoCreateRequest, BossFncUserInvoiceInfoCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.boss.fnc.userinvoiceinfo.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/boss/fnc/userinvoiceinfo/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BossFncUserInvoiceInfoCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// FncUserInvoiceInfoModify calls alipay.boss.fnc.userinvoiceinfo.modify (POST /v3/alipay/boss/fnc/userinvoiceinfo/modify).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xshh
func (s *Service) FncUserInvoiceInfoModify(ctx context.Context, req BossFncUserInvoiceInfoModifyRequest, opts ...core.CallOption) (*core.Response[BossFncUserInvoiceInfoModifyResponse], error) {
	return core.ExecuteJSON[BossFncUserInvoiceInfoModifyRequest, BossFncUserInvoiceInfoModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.boss.fnc.userinvoiceinfo.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/boss/fnc/userinvoiceinfo/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BossFncUserInvoiceInfoModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// FncUserInvoiceInfoQuery calls alipay.boss.fnc.userinvoiceinfo.query (GET /v3/alipay/boss/fnc/userinvoiceinfo/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xmie
func (s *Service) FncUserInvoiceInfoQuery(ctx context.Context, req BossFncUserInvoiceInfoQueryRequest, opts ...core.CallOption) (*core.Response[BossFncUserInvoiceInfoQueryResponse], error) {
	return core.ExecuteJSON[BossFncUserInvoiceInfoQueryRequest, BossFncUserInvoiceInfoQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.boss.fnc.userinvoiceinfo.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/boss/fnc/userinvoiceinfo/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[BossFncUserInvoiceInfoQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
