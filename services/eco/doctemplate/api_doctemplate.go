// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package doctemplate

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

// DocTemplateSettingURLQuery calls alipay.eco.doctemplate.settingurl.query (GET /v3/alipay/eco/doctemplate/settingurl/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05wf9z
func (s *Service) DocTemplateSettingURLQuery(ctx context.Context, req DocTemplateSettingURLQueryRequest, opts ...core.CallOption) (*core.Response[DocTemplateSettingURLQueryResponse], error) {
	return core.ExecuteJSON[DocTemplateSettingURLQueryRequest, DocTemplateSettingURLQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.eco.doctemplate.settingurl.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/eco/doctemplate/settingurl/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[DocTemplateSettingURLQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
