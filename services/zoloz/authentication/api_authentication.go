// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package authentication

import (
	"context"

	"github.com/tiyee/alipay-sdk-go/core"
)

// Service invokes V3 operations in this domain.
type Service struct{ executor core.Executor }

// NewService creates a service using the supplied executor.
func NewService(executor core.Executor) (*Service, error) {
	if executor == nil {
		return nil, &core.ValidationError{Field: "executor", Reason: "executor is required"}
	}
	return &Service{executor: executor}, nil
}

func execute[Req, Resp any](s *Service, ctx context.Context, method string, httpMethod core.HTTPMethod, path string, mode core.ParameterMode, req Req, opts []core.CallOption) (*core.Response[Resp], error) {
	op, err := core.NewJSONOperation[Req, Resp](core.OperationSpec{Method: method, HTTPMethod: httpMethod, Path: path, ParameterMode: mode, AuthRequirement: core.AuthRequirementAny})
	if err != nil {
		return nil, err
	}
	return core.Execute(ctx, s.executor, op, core.Request[Req]{Body: req, Options: core.NewCallSettings(opts...)})
}

// CustomerFaceManageCreate calls zoloz.authentication.customer.facemanage.create.
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vda2
func (s *Service) CustomerFaceManageCreate(ctx context.Context, req CustomerFaceManageCreateRequest, opts ...core.CallOption) (*core.Response[CustomerFaceManageCreateResponse], error) {
	return execute[CustomerFaceManageCreateRequest, CustomerFaceManageCreateResponse](s, ctx, "zoloz.authentication.customer.facemanage.create", core.HTTPMethodPost, "/v3/zoloz/authentication/customer/facemanage/create", core.ParameterModeJSONBody, req, opts)
}

// CustomerFaceManageDelete calls zoloz.authentication.customer.facemanage.delete.
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vh16
func (s *Service) CustomerFaceManageDelete(ctx context.Context, req CustomerFaceManageDeleteRequest, opts ...core.CallOption) (*core.Response[CustomerFaceManageDeleteResponse], error) {
	return execute[CustomerFaceManageDeleteRequest, CustomerFaceManageDeleteResponse](s, ctx, "zoloz.authentication.customer.facemanage.delete", core.HTTPMethodPost, "/v3/zoloz/authentication/customer/facemanage/delete", core.ParameterModeJSONBody, req, opts)
}

// CustomerFTokenQuery calls zoloz.authentication.customer.ftoken.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/ebd5a135_zoloz.authentication.customer.ftoken.query
func (s *Service) CustomerFTokenQuery(ctx context.Context, req CustomerFTokenQueryRequest, opts ...core.CallOption) (*core.Response[CustomerFTokenQueryResponse], error) {
	return execute[CustomerFTokenQueryRequest, CustomerFTokenQueryResponse](s, ctx, "zoloz.authentication.customer.ftoken.query", core.HTTPMethodPost, "/v3/zoloz/authentication/customer/ftoken/query", core.ParameterModeJSONBody, req, opts)
}

// CustomerSmilePayInitialize calls zoloz.authentication.customer.smilepay.initialize.
// Official V3 contract: https://opendocs.alipay.com/pre-apis/05vfpu
func (s *Service) CustomerSmilePayInitialize(ctx context.Context, req CustomerSmilePayInitializeRequest, opts ...core.CallOption) (*core.Response[CustomerSmilePayInitializeResponse], error) {
	return execute[CustomerSmilePayInitializeRequest, CustomerSmilePayInitializeResponse](s, ctx, "zoloz.authentication.customer.smilepay.initialize", core.HTTPMethodPost, "/v3/zoloz/authentication/customer/smilepay/initialize", core.ParameterModeJSONBody, req, opts)
}

// SmilePayInitialize calls zoloz.authentication.smilepay.initialize.
// Official V3 contract: https://opendocs.alipay.com/open-v3/1177dbeb_zoloz.authentication.smilepay.initialize
func (s *Service) SmilePayInitialize(ctx context.Context, req SmilePayInitializeRequest, opts ...core.CallOption) (*core.Response[SmilePayInitializeResponse], error) {
	return execute[SmilePayInitializeRequest, SmilePayInitializeResponse](s, ctx, "zoloz.authentication.smilepay.initialize", core.HTTPMethodPost, "/v3/zoloz/authentication/smilepay/initialize", core.ParameterModeJSONBody, req, opts)
}
