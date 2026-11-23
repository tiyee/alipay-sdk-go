// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package social

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

// BaseContentLibStandardContentBatchQuery calls alipay.social.base.contentlib.standardcontent.batchquery (POST /v3/alipay/social/base/contentlib/standardcontent/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wye1
func (s *Service) BaseContentLibStandardContentBatchQuery(ctx context.Context, req SocialBaseContentLibStandardContentBatchQueryRequest, opts ...core.CallOption) (*core.Response[SocialBaseContentLibStandardContentBatchQueryResponse], error) {
	return core.ExecuteJSON[SocialBaseContentLibStandardContentBatchQueryRequest, SocialBaseContentLibStandardContentBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.social.base.contentlib.standardcontent.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/social/base/contentlib/standardcontent/batchquery",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[SocialBaseContentLibStandardContentBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
