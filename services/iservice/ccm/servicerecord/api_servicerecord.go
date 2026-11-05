// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package servicerecord

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

// CcmServiceRecordCreate calls alipay.iservice.ccm.servicerecord.create (POST /v3/alipay/iservice/ccm/servicerecord/create).
// Official V3 contract: https://opendocs.alipay.com/open-v3/44b1d498_alipay.iservice.ccm.servicerecord.create
func (s *Service) CcmServiceRecordCreate(ctx context.Context, req CcmServiceRecordCreateRequest, opts ...core.CallOption) (*core.Response[CcmServiceRecordCreateResponse], error) {
	return core.ExecuteJSON[CcmServiceRecordCreateRequest, CcmServiceRecordCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.servicerecord.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/servicerecord/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmServiceRecordCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CcmServiceRecordModify calls alipay.iservice.ccm.servicerecord.modify (POST /v3/alipay/iservice/ccm/servicerecord/modify).
// Official V3 contract: https://opendocs.alipay.com/open-v3/fefe64f8_alipay.iservice.ccm.servicerecord.modify
func (s *Service) CcmServiceRecordModify(ctx context.Context, req CcmServiceRecordModifyRequest, opts ...core.CallOption) (*core.Response[CcmServiceRecordModifyResponse], error) {
	return core.ExecuteJSON[CcmServiceRecordModifyRequest, CcmServiceRecordModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.servicerecord.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/servicerecord/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmServiceRecordModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
