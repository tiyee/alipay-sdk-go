// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package educate

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

// EducateCampusInstitutionsAdd calls alipay.commerce.educate.campus.institutions.add (POST /v3/alipay/commerce/educate/campus/institutions/add).
// Official V3 contract: https://opendocs.alipay.com/open-v3/2f573ce0_alipay.commerce.educate.campus.institutions.add
func (s *Service) EducateCampusInstitutionsAdd(ctx context.Context, req EducateCampusInstitutionsAddRequest, opts ...core.CallOption) (*core.Response[EducateCampusInstitutionsAddResponse], error) {
	return core.ExecuteJSON[EducateCampusInstitutionsAddRequest, EducateCampusInstitutionsAddResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.educate.campus.institutions.add", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/educate/campus/institutions/add",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EducateCampusInstitutionsAddRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EducateCampusInstitutionsQuery calls alipay.commerce.educate.campus.institutions.query (GET /v3/alipay/commerce/educate/campus/institutions/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b340d586_alipay.commerce.educate.campus.institutions.query
func (s *Service) EducateCampusInstitutionsQuery(ctx context.Context, req EducateCampusInstitutionsQueryRequest, opts ...core.CallOption) (*core.Response[EducateCampusInstitutionsQueryResponse], error) {
	return core.ExecuteJSON[EducateCampusInstitutionsQueryRequest, EducateCampusInstitutionsQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.educate.campus.institutions.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/educate/campus/institutions/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EducateCampusInstitutionsQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// EducateSchoolCardOrderSync calls alipay.commerce.educate.schoolcard.order.sync (POST /v3/alipay/commerce/educate/schoolcard/order/sync).
// Official V3 contract: https://opendocs.alipay.com/open-v3/b4497dd8_alipay.commerce.educate.schoolcard.order.sync
func (s *Service) EducateSchoolCardOrderSync(ctx context.Context, req EducateSchoolCardOrderSyncRequest, opts ...core.CallOption) (*core.Response[EducateSchoolCardOrderSyncResponse], error) {
	return core.ExecuteJSON[EducateSchoolCardOrderSyncRequest, EducateSchoolCardOrderSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.educate.schoolcard.order.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/educate/schoolcard/order/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[EducateSchoolCardOrderSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
