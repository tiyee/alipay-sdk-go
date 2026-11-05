// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package templatemsg

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

// MiniTemplateMsgMaketingBatchQuery calls alipay.open.mini.templatemsg.maketing.batchquery (POST /v3/alipay/open/mini/templatemsg/maketing/batchquery).
// Official V3 contract: https://opendocs.alipay.com/open-v3/59c80879_alipay.open.mini.templatemsg.maketing.batchquery
func (s *Service) MiniTemplateMsgMaketingBatchQuery(ctx context.Context, req MiniTemplateMsgMaketingBatchQueryRequest, opts ...core.CallOption) (*core.Response[MiniTemplateMsgMaketingBatchQueryResponse], error) {
	return core.ExecuteJSON[MiniTemplateMsgMaketingBatchQueryRequest, MiniTemplateMsgMaketingBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.templatemsg.maketing.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/templatemsg/maketing/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniTemplateMsgMaketingBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// MiniTemplateMsgTinyPaySwitchConfirm calls alipay.open.mini.templatemsg.tinypayswitch.confirm (POST /v3/alipay/open/mini/templatemsg/tinypayswitch/confirm).
// Official V3 contract: https://opendocs.alipay.com/open-v3/653c0f4a_alipay.open.mini.templatemsg.tinypayswitch.confirm
func (s *Service) MiniTemplateMsgTinyPaySwitchConfirm(ctx context.Context, req MiniTemplateMsgTinyPaySwitchConfirmRequest, opts ...core.CallOption) (*core.Response[MiniTemplateMsgTinyPaySwitchConfirmResponse], error) {
	return core.ExecuteJSON[MiniTemplateMsgTinyPaySwitchConfirmRequest, MiniTemplateMsgTinyPaySwitchConfirmResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.mini.templatemsg.tinypayswitch.confirm", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/mini/templatemsg/tinypayswitch/confirm",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[MiniTemplateMsgTinyPaySwitchConfirmRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
