// Code generated from Alipay V3 OpenAPI 2026-09-23; DO NOT EDIT.
package fincloud

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

// GeneralSaaSFaceCertifyInitialize calls datadigital.fincloud.generalsaas.face.certify.initialize.
// Official V3 contract: https://opendocs.alipay.com/open-v3/645bdc9c_datadigital.fincloud.generalsaas.face.certify.initialize
func (s *Service) GeneralSaaSFaceCertifyInitialize(ctx context.Context, req GeneralSaaSFaceCertifyInitializeRequest, opts ...core.CallOption) (*core.Response[GeneralSaaSFaceCertifyInitializeResponse], error) {
	return execute[GeneralSaaSFaceCertifyInitializeRequest, GeneralSaaSFaceCertifyInitializeResponse](s, ctx, "datadigital.fincloud.generalsaas.face.certify.initialize", core.HTTPMethodPost, "/v3/datadigital/fincloud/generalsaas/face/certify/initialize", core.ParameterModeJSONBody, req, opts)
}

// GeneralSaaSFaceCertifyQuery calls datadigital.fincloud.generalsaas.face.certify.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/10a1ccfd_datadigital.fincloud.generalsaas.face.certify.query
func (s *Service) GeneralSaaSFaceCertifyQuery(ctx context.Context, req GeneralSaaSFaceCertifyQueryRequest, opts ...core.CallOption) (*core.Response[GeneralSaaSFaceCertifyQueryResponse], error) {
	return execute[GeneralSaaSFaceCertifyQueryRequest, GeneralSaaSFaceCertifyQueryResponse](s, ctx, "datadigital.fincloud.generalsaas.face.certify.query", core.HTTPMethodGet, "/v3/datadigital/fincloud/generalsaas/face/certify/query", core.ParameterModeQuery, req, opts)
}

// GeneralSaaSFaceCertifyVerify calls datadigital.fincloud.generalsaas.face.certify.verify.
// Official V3 contract: https://opendocs.alipay.com/open-v3/d3f37a88_datadigital.fincloud.generalsaas.face.certify.verify
func (s *Service) GeneralSaaSFaceCertifyVerify(ctx context.Context, req GeneralSaaSFaceCertifyVerifyRequest, opts ...core.CallOption) (*core.Response[GeneralSaaSFaceCertifyVerifyResponse], error) {
	return execute[GeneralSaaSFaceCertifyVerifyRequest, GeneralSaaSFaceCertifyVerifyResponse](s, ctx, "datadigital.fincloud.generalsaas.face.certify.verify", core.HTTPMethodPost, "/v3/datadigital/fincloud/generalsaas/face/certify/verify", core.ParameterModeJSONBody, req, opts)
}

// GeneralSaaSFaceCheckInitialize calls datadigital.fincloud.generalsaas.face.check.initialize.
// Official V3 contract: https://opendocs.alipay.com/open-v3/503aa13e_datadigital.fincloud.generalsaas.face.check.initialize
func (s *Service) GeneralSaaSFaceCheckInitialize(ctx context.Context, req GeneralSaaSFaceCheckInitializeRequest, opts ...core.CallOption) (*core.Response[GeneralSaaSFaceCheckInitializeResponse], error) {
	return execute[GeneralSaaSFaceCheckInitializeRequest, GeneralSaaSFaceCheckInitializeResponse](s, ctx, "datadigital.fincloud.generalsaas.face.check.initialize", core.HTTPMethodPost, "/v3/datadigital/fincloud/generalsaas/face/check/initialize", core.ParameterModeJSONBody, req, opts)
}

// GeneralSaaSFaceCheckQuery calls datadigital.fincloud.generalsaas.face.check.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/4b93abbc_datadigital.fincloud.generalsaas.face.check.query
func (s *Service) GeneralSaaSFaceCheckQuery(ctx context.Context, req GeneralSaaSFaceCheckQueryRequest, opts ...core.CallOption) (*core.Response[GeneralSaaSFaceCheckQueryResponse], error) {
	return execute[GeneralSaaSFaceCheckQueryRequest, GeneralSaaSFaceCheckQueryResponse](s, ctx, "datadigital.fincloud.generalsaas.face.check.query", core.HTTPMethodGet, "/v3/datadigital/fincloud/generalsaas/face/check/query", core.ParameterModeQuery, req, opts)
}

// GeneralSaaSFaceSourceCertify calls datadigital.fincloud.generalsaas.face.source.certify.
// Official V3 contract: https://opendocs.alipay.com/open-v3/21280890_datadigital.fincloud.generalsaas.face.source.certify
func (s *Service) GeneralSaaSFaceSourceCertify(ctx context.Context, req GeneralSaaSFaceSourceCertifyRequest, opts ...core.CallOption) (*core.Response[GeneralSaaSFaceSourceCertifyResponse], error) {
	return execute[GeneralSaaSFaceSourceCertifyRequest, GeneralSaaSFaceSourceCertifyResponse](s, ctx, "datadigital.fincloud.generalsaas.face.source.certify", core.HTTPMethodPost, "/v3/datadigital/fincloud/generalsaas/face/source/certify", core.ParameterModeMultipart, req, opts)
}

// GeneralSaaSFaceVerificationInitialize calls datadigital.fincloud.generalsaas.face.verification.initialize.
// Official V3 contract: https://opendocs.alipay.com/open-v3/34fb02dc_datadigital.fincloud.generalsaas.face.verification.initialize
func (s *Service) GeneralSaaSFaceVerificationInitialize(ctx context.Context, req GeneralSaaSFaceVerificationInitializeRequest, opts ...core.CallOption) (*core.Response[GeneralSaaSFaceVerificationInitializeResponse], error) {
	return execute[GeneralSaaSFaceVerificationInitializeRequest, GeneralSaaSFaceVerificationInitializeResponse](s, ctx, "datadigital.fincloud.generalsaas.face.verification.initialize", core.HTTPMethodPost, "/v3/datadigital/fincloud/generalsaas/face/verification/initialize", core.ParameterModeJSONBody, req, opts)
}

// GeneralSaaSFaceVerificationQuery calls datadigital.fincloud.generalsaas.face.verification.query.
// Official V3 contract: https://opendocs.alipay.com/open-v3/1ef1e2ed_datadigital.fincloud.generalsaas.face.verification.query
func (s *Service) GeneralSaaSFaceVerificationQuery(ctx context.Context, req GeneralSaaSFaceVerificationQueryRequest, opts ...core.CallOption) (*core.Response[GeneralSaaSFaceVerificationQueryResponse], error) {
	return execute[GeneralSaaSFaceVerificationQueryRequest, GeneralSaaSFaceVerificationQueryResponse](s, ctx, "datadigital.fincloud.generalsaas.face.verification.query", core.HTTPMethodGet, "/v3/datadigital/fincloud/generalsaas/face/verification/query", core.ParameterModeQuery, req, opts)
}

// GeneralSaaSOCRMobileInitialize calls datadigital.fincloud.generalsaas.ocr.mobile.initialize.
// Official V3 contract: https://opendocs.alipay.com/open-v3/74ea0784_datadigital.fincloud.generalsaas.ocr.mobile.initialize
func (s *Service) GeneralSaaSOCRMobileInitialize(ctx context.Context, req GeneralSaaSOCRMobileInitializeRequest, opts ...core.CallOption) (*core.Response[GeneralSaaSOCRMobileInitializeResponse], error) {
	return execute[GeneralSaaSOCRMobileInitializeRequest, GeneralSaaSOCRMobileInitializeResponse](s, ctx, "datadigital.fincloud.generalsaas.ocr.mobile.initialize", core.HTTPMethodPost, "/v3/datadigital/fincloud/generalsaas/ocr/mobile/initialize", core.ParameterModeJSONBody, req, opts)
}

// GeneralSaaSOCRServerDetect calls datadigital.fincloud.generalsaas.ocr.server.detect.
// Official V3 contract: https://opendocs.alipay.com/open-v3/a6f20b9d_datadigital.fincloud.generalsaas.ocr.server.detect
func (s *Service) GeneralSaaSOCRServerDetect(ctx context.Context, req GeneralSaaSOCRServerDetectRequest, opts ...core.CallOption) (*core.Response[GeneralSaaSOCRServerDetectResponse], error) {
	return execute[GeneralSaaSOCRServerDetectRequest, GeneralSaaSOCRServerDetectResponse](s, ctx, "datadigital.fincloud.generalsaas.ocr.server.detect", core.HTTPMethodPost, "/v3/datadigital/fincloud/generalsaas/ocr/server/detect", core.ParameterModeMultipart, req, opts)
}
