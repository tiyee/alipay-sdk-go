// Package monitor provides the V3 heartbeat service.
package monitor

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service invokes monitor operations using a shared executor.
type Service struct{ executor core.Executor }

// NewService constructs a monitor service.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

// HeartbeatSyn calls monitor.heartbeat.syn.
// Official V3 contract: https://opendocs.alipay.com/apis/05vc16
func (s *Service) HeartbeatSyn(ctx context.Context, req HeartbeatSynRequest, opts ...core.CallOption) (*core.Response[HeartbeatSynResponse], error) {
	op, err := core.NewJSONOperation[HeartbeatSynRequest, HeartbeatSynResponse](core.OperationSpec{
		Method: "monitor.heartbeat.syn", HTTPMethod: core.HTTPMethodPost,
		Path: "/v3/monitor/heartbeat/syn", ParameterMode: core.ParameterModeJSONBody,
		AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, op, core.Request[HeartbeatSynRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
