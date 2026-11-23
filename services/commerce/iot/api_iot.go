// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package iot

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

// IoTDeviceReportUpload calls alipay.commerce.iot.device.report.upload (POST /v3/alipay/commerce/iot/device/report/upload).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05xk46
func (s *Service) IoTDeviceReportUpload(ctx context.Context, req IoTDeviceReportUploadRequest, opts ...core.CallOption) (*core.Response[IoTDeviceReportUploadResponse], error) {
	return core.ExecuteJSON[IoTDeviceReportUploadRequest, IoTDeviceReportUploadResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.iot.device.report.upload", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/iot/device/report/upload",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[IoTDeviceReportUploadRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// IoTSdarttoolMessageQuery calls alipay.commerce.iot.sdarttool.message.query (GET /v3/alipay/commerce/iot/sdarttool/message/query).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/06bpqm
func (s *Service) IoTSdarttoolMessageQuery(ctx context.Context, req IoTSdarttoolMessageQueryRequest, opts ...core.CallOption) (*core.Response[IoTSdarttoolMessageQueryResponse], error) {
	return core.ExecuteJSON[IoTSdarttoolMessageQueryRequest, IoTSdarttoolMessageQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.iot.sdarttool.message.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/alipay/commerce/iot/sdarttool/message/query",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[IoTSdarttoolMessageQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// IoTSdarttoolMessageSend calls alipay.commerce.iot.sdarttool.message.send (POST /v3/alipay/commerce/iot/sdarttool/message/send).
// Official V3 contract: https://opendocs.alipay.com/pre-apis/06bzwt
func (s *Service) IoTSdarttoolMessageSend(ctx context.Context, req IoTSdarttoolMessageSendRequest, opts ...core.CallOption) (*core.Response[IoTSdarttoolMessageSendResponse], error) {
	return core.ExecuteJSON[IoTSdarttoolMessageSendRequest, IoTSdarttoolMessageSendResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.commerce.iot.sdarttool.message.send", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/commerce/iot/sdarttool/message/send",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[IoTSdarttoolMessageSendRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
