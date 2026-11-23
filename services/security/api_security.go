// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package security

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service executes one alipay V3 service group.
type Service struct{ executor core.Executor }

// NewService constructs the service with a shared V3 executor.
// NewService constructs the service with a shared V3 executor.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

// ProdSssQuery calls alipay.security.prod.sss.query (POST /v3/alipay/security/prod/sss/query/{bbb}).
// Official V3 contract: https://opendocs.alipay.com/apis/0557cf
func (s *Service) ProdSssQuery(ctx context.Context, req SecurityProdSssQueryRequest, opts ...core.CallOption) (*core.Response[SecurityProdSssQueryResponse], error) {
	return core.ExecuteJSON[SecurityProdSssQueryRequest, SecurityProdSssQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.security.prod.sss.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/security/prod/sss/query/{bbb}",
		ParameterMode: core.ParameterModeJSONBodyWithQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SecurityProdSssQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// RiskCustomerRiskSend calls alipay.security.risk.customerrisk.send (POST /v3/alipay/security/risk/customerrisk/send).
// Official V3 contract: https://opendocs.alipay.com/open-v3/cfacc9de_alipay.security.risk.customerrisk.send
func (s *Service) RiskCustomerRiskSend(ctx context.Context, req SecurityRiskCustomerRiskSendRequest, opts ...core.CallOption) (*core.Response[SecurityRiskCustomerRiskSendResponse], error) {
	return core.ExecuteJSON[SecurityRiskCustomerRiskSendRequest, SecurityRiskCustomerRiskSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.security.risk.customerrisk.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/security/risk/customerrisk/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SecurityRiskCustomerRiskSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// RiskVerifyIdentityMiniAppConfirm calls alipay.security.risk.verifyidentity.miniapp.confirm (POST /v3/alipay/security/risk/verifyidentity/miniapp/confirm).
// Official V3 contract: https://opendocs.alipay.com/apis/05vidu
func (s *Service) RiskVerifyIdentityMiniAppConfirm(ctx context.Context, req SecurityRiskVerifyIdentityMiniAppConfirmRequest, opts ...core.CallOption) (*core.Response[SecurityRiskVerifyIdentityMiniAppConfirmResponse], error) {
	return core.ExecuteJSON[SecurityRiskVerifyIdentityMiniAppConfirmRequest, SecurityRiskVerifyIdentityMiniAppConfirmResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.security.risk.verifyidentity.miniapp.confirm", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/security/risk/verifyidentity/miniapp/confirm",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SecurityRiskVerifyIdentityMiniAppConfirmRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
