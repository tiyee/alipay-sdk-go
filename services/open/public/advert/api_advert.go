// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package advert

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

// PublicAdvertBatchQuery calls alipay.open.public.advert.batchquery (POST /v3/alipay/open/public/advert/batchquery).
// Official V3 contract: https://opendocs.alipay.com/apis/05wggg
func (s *Service) PublicAdvertBatchQuery(ctx context.Context, req PublicAdvertBatchQueryRequest, opts ...core.CallOption) (*core.Response[PublicAdvertBatchQueryResponse], error) {
	return core.ExecuteJSON[PublicAdvertBatchQueryRequest, PublicAdvertBatchQueryResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.advert.batchquery", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/advert/batchquery",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicAdvertBatchQueryRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicAdvertCreate calls alipay.open.public.advert.create (POST /v3/alipay/open/public/advert/create).
// Official V3 contract: https://opendocs.alipay.com/apis/05xmcx
func (s *Service) PublicAdvertCreate(ctx context.Context, req PublicAdvertCreateRequest, opts ...core.CallOption) (*core.Response[PublicAdvertCreateResponse], error) {
	return core.ExecuteJSON[PublicAdvertCreateRequest, PublicAdvertCreateResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.advert.create", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/advert/create",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicAdvertCreateRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicAdvertDelete calls alipay.open.public.advert.delete (DELETE /v3/alipay/open/public/advert/delete).
// Official V3 contract: https://opendocs.alipay.com/apis/05wvvm
func (s *Service) PublicAdvertDelete(ctx context.Context, req PublicAdvertDeleteRequest, opts ...core.CallOption) (*core.Response[PublicAdvertDeleteResponse], error) {
	return core.ExecuteJSON[PublicAdvertDeleteRequest, PublicAdvertDeleteResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.advert.delete", HTTPMethod: core.HTTPMethodDelete, Path: "/v3/alipay/open/public/advert/delete",
		ParameterMode: core.ParameterModeQuery, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicAdvertDeleteRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}

// PublicAdvertModify calls alipay.open.public.advert.modify (POST /v3/alipay/open/public/advert/modify).
// Official V3 contract: https://opendocs.alipay.com/apis/05x1fw
func (s *Service) PublicAdvertModify(ctx context.Context, req PublicAdvertModifyRequest, opts ...core.CallOption) (*core.Response[PublicAdvertModifyResponse], error) {
	return core.ExecuteJSON[PublicAdvertModifyRequest, PublicAdvertModifyResponse](ctx, s.executor, core.OperationSpec{
		Method: "alipay.open.public.advert.modify", HTTPMethod: core.HTTPMethodPost, Path: "/v3/alipay/open/public/advert/modify",
		ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
	}, core.Request[PublicAdvertModifyRequest]{Body: req, Options: core.NewCallSettings(opts...)})
}
