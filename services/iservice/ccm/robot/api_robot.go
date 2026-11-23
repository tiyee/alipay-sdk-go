// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package robot

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

// CcmRobotAvatarBaseQuery calls alipay.iservice.ccm.robot.avatarbase.query (POST /v3/alipay/iservice/ccm/robot/avatarbase/query).
// Official V3 contract: https://opendocs.alipay.com/open-v3/6d82b0cb_alipay.iservice.ccm.robot.avatarbase.query
func (s *Service) CcmRobotAvatarBaseQuery(ctx context.Context, req CcmRobotAvatarBaseQueryRequest, opts ...core.CallOption) (*core.Response[CcmRobotAvatarBaseQueryResponse], error) {
	return core.ExecuteJSON[CcmRobotAvatarBaseQueryRequest, CcmRobotAvatarBaseQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.iservice.ccm.robot.avatarbase.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/iservice/ccm/robot/avatarbase/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[CcmRobotAvatarBaseQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
