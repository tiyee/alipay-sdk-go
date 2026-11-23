// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package edu

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

// EduKtBillingModify calls alipay.eco.edu.kt.billing.modify (POST /v3/alipay/eco/edu/kt/billing/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/1406f415_alipay.eco.edu.kt.billing.modify
func (s *Service) EduKtBillingModify(ctx context.Context, req EduKtBillingModifyRequest, opts ...core.CallOption) (*core.Response[EduKtBillingModifyResponse], error) {
	return core.ExecuteJSON[EduKtBillingModifyRequest, EduKtBillingModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.edu.kt.billing.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/edu/kt/billing/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EduKtBillingModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EduKtBillingQuery calls alipay.eco.edu.kt.billing.query (GET /v3/alipay/eco/edu/kt/billing/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/d9852f73_alipay.eco.edu.kt.billing.query
func (s *Service) EduKtBillingQuery(ctx context.Context, req EduKtBillingQueryRequest, opts ...core.CallOption) (*core.Response[EduKtBillingQueryResponse], error) {
	return core.ExecuteJSON[EduKtBillingQueryRequest, EduKtBillingQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.edu.kt.billing.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/eco/edu/kt/billing/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EduKtBillingQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EduKtBillingSend calls alipay.eco.edu.kt.billing.send (POST /v3/alipay/eco/edu/kt/billing/send).
// Official V3 contract: https://opendocs.alipay.com/open-v3/7c14b5a1_alipay.eco.edu.kt.billing.send
func (s *Service) EduKtBillingSend(ctx context.Context, req EduKtBillingSendRequest, opts ...core.CallOption) (*core.Response[EduKtBillingSendResponse], error) {
	return core.ExecuteJSON[EduKtBillingSendRequest, EduKtBillingSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.edu.kt.billing.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/edu/kt/billing/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EduKtBillingSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EduKtParentQuery calls alipay.eco.edu.kt.parent.query (GET /v3/alipay/eco/edu/kt/parent/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05x0kr
func (s *Service) EduKtParentQuery(ctx context.Context, req EduKtParentQueryRequest, opts ...core.CallOption) (*core.Response[EduKtParentQueryResponse], error) {
	return core.ExecuteJSON[EduKtParentQueryRequest, EduKtParentQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.edu.kt.parent.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/eco/edu/kt/parent/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EduKtParentQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EduKtSchoolInfoModify calls alipay.eco.edu.kt.schoolinfo.modify (POST /v3/alipay/eco/edu/kt/schoolinfo/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/a319c0c9_alipay.eco.edu.kt.schoolinfo.modify
func (s *Service) EduKtSchoolInfoModify(ctx context.Context, req EduKtSchoolInfoModifyRequest, opts ...core.CallOption) (*core.Response[EduKtSchoolInfoModifyResponse], error) {
	return core.ExecuteJSON[EduKtSchoolInfoModifyRequest, EduKtSchoolInfoModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.edu.kt.schoolinfo.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/edu/kt/schoolinfo/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EduKtSchoolInfoModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EduKtStudentModify calls alipay.eco.edu.kt.student.modify (POST /v3/alipay/eco/edu/kt/student/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wun3
func (s *Service) EduKtStudentModify(ctx context.Context, req EduKtStudentModifyRequest, opts ...core.CallOption) (*core.Response[EduKtStudentModifyResponse], error) {
	return core.ExecuteJSON[EduKtStudentModifyRequest, EduKtStudentModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.edu.kt.student.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/eco/edu/kt/student/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EduKtStudentModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
