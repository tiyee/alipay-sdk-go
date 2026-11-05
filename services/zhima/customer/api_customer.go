// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package customer

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

// JobWorthAuthenticationQuery calls zhima.customer.jobworth.authentication.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/d25db704_zhima.customer.jobworth.authentication.query
func (s *Service) JobWorthAuthenticationQuery(ctx context.Context, req JobWorthAuthenticationQueryRequest, opts ...core.CallOption) (*core.Response[JobWorthAuthenticationQueryResponse], error) {
	return execute[JobWorthAuthenticationQueryRequest, JobWorthAuthenticationQueryResponse](s, ctx, "zhima.customer.jobworth.authentication.query", core.HTTPMethodGet, "/v3/zhima/customer/jobworth/authentication", core.ParameterModeQuery, req, opts)
}

// JobWorthCloudResumeQuery calls zhima.customer.jobworth.cloudresume.query.
// Official V3 contract: https://opendocs.alipay.com/apis/05wukc
func (s *Service) JobWorthCloudResumeQuery(ctx context.Context, req JobWorthCloudResumeQueryRequest, opts ...core.CallOption) (*core.Response[JobWorthCloudResumeQueryResponse], error) {
	return execute[JobWorthCloudResumeQueryRequest, JobWorthCloudResumeQueryResponse](s, ctx, "zhima.customer.jobworth.cloudresume.query", core.HTTPMethodGet, "/v3/zhima/customer/jobworth/cloudresume/query", core.ParameterModeQuery, req, opts)
}

// JobWorthInfoQuery calls zhima.customer.jobworth.info.query.
// Official V3 contract: https://opendocs.alipay.com/apis/05ttl7
func (s *Service) JobWorthInfoQuery(ctx context.Context, req JobWorthInfoQueryRequest, opts ...core.CallOption) (*core.Response[JobWorthInfoQueryResponse], error) {
	return execute[JobWorthInfoQueryRequest, JobWorthInfoQueryResponse](s, ctx, "zhima.customer.jobworth.info.query", core.HTTPMethodGet, "/v3/zhima/customer/jobworth/info/query", core.ParameterModeQuery, req, opts)
}

// JobWorthJobDataAdd calls zhima.customer.jobworth.jobdata.add.
// Official V3 contract: https://opendocs.alipay.com/apis/05wyeb
func (s *Service) JobWorthJobDataAdd(ctx context.Context, req JobWorthJobDataAddRequest, opts ...core.CallOption) (*core.Response[JobWorthJobDataAddResponse], error) {
	return execute[JobWorthJobDataAddRequest, JobWorthJobDataAddResponse](s, ctx, "zhima.customer.jobworth.jobdata.add", core.HTTPMethodPost, "/v3/zhima/customer/jobworth/jobdata/add", core.ParameterModeJSONBody, req, opts)
}

// JobWorthPictureUpload calls zhima.customer.jobworth.picture.upload.
// Official V3 contract: https://opendocs.alipay.com/apis/05wulq
func (s *Service) JobWorthPictureUpload(ctx context.Context, req JobWorthPictureUploadRequest, opts ...core.CallOption) (*core.Response[JobWorthPictureUploadResponse], error) {
	return execute[JobWorthPictureUploadRequest, JobWorthPictureUploadResponse](s, ctx, "zhima.customer.jobworth.picture.upload", core.HTTPMethodPost, "/v3/zhima/customer/jobworth/picture/upload", core.ParameterModeJSONBody, req, opts)
}

// JobWorthPositionAdd calls zhima.customer.jobworth.position.add.
// Official V3 contract: https://opendocs.alipay.com/apis/05xa7n
func (s *Service) JobWorthPositionAdd(ctx context.Context, req JobWorthPositionAddRequest, opts ...core.CallOption) (*core.Response[JobWorthPositionAddResponse], error) {
	return execute[JobWorthPositionAddRequest, JobWorthPositionAddResponse](s, ctx, "zhima.customer.jobworth.position.add", core.HTTPMethodPost, "/v3/zhima/customer/jobworth/position/add", core.ParameterModeJSONBody, req, opts)
}
