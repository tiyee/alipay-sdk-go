// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package ols

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

// CcmOlsChatRecordQuery calls alipay.iservice.ccm.ols.chatrecord.query (GET /v3/alipay/iservice/ccm/ols/chatrecord/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wokm
func (s *Service) CcmOlsChatRecordQuery(ctx context.Context, req CcmOlsChatRecordQueryRequest, opts ...core.CallOption) (*core.Response[CcmOlsChatRecordQueryResponse], error) {
	return core.ExecuteJSON[CcmOlsChatRecordQueryRequest, CcmOlsChatRecordQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.ols.chatrecord.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/iservice/ccm/ols/chatrecord/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmOlsChatRecordQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
