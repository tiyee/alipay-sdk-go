// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package group

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

// PublicGroupBatchQuery calls alipay.open.public.group.batchquery (POST /v3/alipay/open/public/group/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wydu
func (s *Service) PublicGroupBatchQuery(ctx context.Context, req PublicGroupBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicGroupBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicGroupBatchQueryRequest, PublicGroupBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.group.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/group/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicGroupBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicGroupCreate calls alipay.open.public.group.create (POST /v3/alipay/open/public/group/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05wukz
func (s *Service) PublicGroupCreate(ctx context.Context, req PublicGroupCreateRequest, opts ...core.CallOption) (*core.Response[PublicGroupCreateResponse], error) {
	return core.ExecuteJSON[PublicGroupCreateRequest, PublicGroupCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.group.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/group/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicGroupCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicGroupCrowdQuery calls alipay.open.public.group.crowd.query (POST /v3/alipay/open/public/group/crowd/query).
// Official V3 contract: https://opendocs.alipay.com/apis/05wyfg
func (s *Service) PublicGroupCrowdQuery(ctx context.Context, req PublicGroupCrowdQueryRequest, opts ...core.CallOption) (*core.Response[PublicGroupCrowdQueryResponse], error) {
	return core.ExecuteJSON[PublicGroupCrowdQueryRequest, PublicGroupCrowdQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.group.crowd.query", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/group/crowd/query",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicGroupCrowdQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicGroupDelete calls alipay.open.public.group.delete (DELETE /v3/alipay/open/public/group/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05wuht
func (s *Service) PublicGroupDelete(ctx context.Context, req PublicGroupDeleteRequest, opts ...core.CallOption) (*core.Response[PublicGroupDeleteResponse], error) {
	return core.ExecuteJSON[PublicGroupDeleteRequest, PublicGroupDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.group.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/public/group/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicGroupDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicGroupModify calls alipay.open.public.group.modify (POST /v3/alipay/open/public/group/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05wcqq
func (s *Service) PublicGroupModify(ctx context.Context, req PublicGroupModifyRequest, opts ...core.CallOption) (*core.Response[PublicGroupModifyResponse], error) {
	return core.ExecuteJSON[PublicGroupModifyRequest, PublicGroupModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.group.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/group/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicGroupModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
