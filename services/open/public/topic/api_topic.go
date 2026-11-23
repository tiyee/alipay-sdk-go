// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package topic

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

// PublicTopicBatchQuery calls alipay.open.public.topic.batchquery (POST /v3/alipay/open/public/topic/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvr7
func (s *Service) PublicTopicBatchQuery(ctx context.Context, req PublicTopicBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicTopicBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicTopicBatchQueryRequest, PublicTopicBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.topic.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/topic/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicTopicBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicTopicCreate calls alipay.open.public.topic.create (POST /v3/alipay/open/public/topic/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuj6
func (s *Service) PublicTopicCreate(ctx context.Context, req PublicTopicCreateRequest, opts ...core.CallOption) (*core.Response[PublicTopicCreateResponse], error) {
	return core.ExecuteJSON[PublicTopicCreateRequest, PublicTopicCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.topic.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/topic/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicTopicCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicTopicDelete calls alipay.open.public.topic.delete (DELETE /v3/alipay/open/public/topic/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05wyj8
func (s *Service) PublicTopicDelete(ctx context.Context, req PublicTopicDeleteRequest, opts ...core.CallOption) (*core.Response[PublicTopicDeleteResponse], error) {
	return core.ExecuteJSON[PublicTopicDeleteRequest, PublicTopicDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.topic.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/public/topic/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicTopicDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicTopicModify calls alipay.open.public.topic.modify (POST /v3/alipay/open/public/topic/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuk1
func (s *Service) PublicTopicModify(ctx context.Context, req PublicTopicModifyRequest, opts ...core.CallOption) (*core.Response[PublicTopicModifyResponse], error) {
	return core.ExecuteJSON[PublicTopicModifyRequest, PublicTopicModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.topic.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/topic/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicTopicModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
