// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package apply

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

// InvoiceApplyInvUpload calls alipay.ebpp.invoice.apply.inv.upload (PUT /v3/alipay/ebpp/invoice/apply/inv/upload).
// Official V3 contract: https://opendocs.alipay.com/apis/05ujox
func (s *Service) InvoiceApplyInvUpload(ctx context.Context, req InvoiceApplyInvUploadRequest, opts ...core.CallOption) (*core.Response[InvoiceApplyInvUploadResponse], error) {
	return core.ExecuteJSON[InvoiceApplyInvUploadRequest, InvoiceApplyInvUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.apply.inv.upload", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/ebpp/invoice/apply/inv/upload",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceApplyInvUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceApplyResultSync calls alipay.ebpp.invoice.apply.result.sync (POST /v3/alipay/ebpp/invoice/apply/result/sync).
// Official V3 contract: https://opendocs.alipay.com/apis/05vwzk
func (s *Service) InvoiceApplyResultSync(ctx context.Context, req InvoiceApplyResultSyncRequest, opts ...core.CallOption) (*core.Response[InvoiceApplyResultSyncResponse], error) {
	return core.ExecuteJSON[InvoiceApplyResultSyncRequest, InvoiceApplyResultSyncResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.apply.result.sync", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/ebpp/invoice/apply/result/sync",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceApplyResultSyncRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// InvoiceApplyStatusNotify calls alipay.ebpp.invoice.apply.status.notify (PUT /v3/alipay/ebpp/invoice/apply/status/notify).
// Official V3 contract: https://opendocs.alipay.com/apis/05uo9u
func (s *Service) InvoiceApplyStatusNotify(ctx context.Context, req InvoiceApplyStatusNotifyRequest, opts ...core.CallOption) (*core.Response[InvoiceApplyStatusNotifyResponse], error) {
	return core.ExecuteJSON[InvoiceApplyStatusNotifyRequest, InvoiceApplyStatusNotifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.ebpp.invoice.apply.status.notify", HTTPMethod: core.HTTPMethodPut, Path: "/v3/alipay/ebpp/invoice/apply/status/notify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[InvoiceApplyStatusNotifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
