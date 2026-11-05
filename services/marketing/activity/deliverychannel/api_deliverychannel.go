// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package deliverychannel

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

// ActivityDeliveryChannelQuery calls alipay.marketing.activity.deliverychannel.query (POST /v3/alipay/marketing/activity/deliverychannel/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05v7cm
func (s *Service) ActivityDeliveryChannelQuery(ctx context.Context, req ActivityDeliveryChannelQueryRequest, opts ...core.CallOption) (*core.Response[ActivityDeliveryChannelQueryResponse], error) {
	return core.ExecuteJSON[ActivityDeliveryChannelQueryRequest, ActivityDeliveryChannelQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.marketing.activity.deliverychannel.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/marketing/activity/deliverychannel/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[ActivityDeliveryChannelQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
