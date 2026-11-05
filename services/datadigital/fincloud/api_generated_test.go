package fincloud

import (
	"context"
	"github.com/tiyee/alipay-sdk-go/core"
	"testing"
)

type recordingExecutor struct{ req core.WireRequest }

func (e *recordingExecutor) Do(_ context.Context, r core.WireRequest) (*core.APIResult, error) {
	e.req = r
	return &core.APIResult{HTTPStatus: 200, SignatureVerified: true, Node: []byte(`{}`)}, nil
}
func TestOperationSpecs(t *testing.T) {
	cases := []struct {
		operation, method, path string
		mode                    core.ParameterMode
		call                    func(*Service) error
	}{
		{"datadigital.fincloud.generalsaas.face.certify.initialize", "POST", "/v3/datadigital/fincloud/generalsaas/face/certify/initialize", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.GeneralSaaSFaceCertifyInitialize(context.Background(), GeneralSaaSFaceCertifyInitializeRequest{})
			return err
		}},
		{"datadigital.fincloud.generalsaas.face.certify.query", "GET", "/v3/datadigital/fincloud/generalsaas/face/certify/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.GeneralSaaSFaceCertifyQuery(context.Background(), GeneralSaaSFaceCertifyQueryRequest{})
			return err
		}},
		{"datadigital.fincloud.generalsaas.face.certify.verify", "POST", "/v3/datadigital/fincloud/generalsaas/face/certify/verify", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.GeneralSaaSFaceCertifyVerify(context.Background(), GeneralSaaSFaceCertifyVerifyRequest{})
			return err
		}},
		{"datadigital.fincloud.generalsaas.face.check.initialize", "POST", "/v3/datadigital/fincloud/generalsaas/face/check/initialize", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.GeneralSaaSFaceCheckInitialize(context.Background(), GeneralSaaSFaceCheckInitializeRequest{})
			return err
		}},
		{"datadigital.fincloud.generalsaas.face.check.query", "GET", "/v3/datadigital/fincloud/generalsaas/face/check/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.GeneralSaaSFaceCheckQuery(context.Background(), GeneralSaaSFaceCheckQueryRequest{})
			return err
		}},
		{"datadigital.fincloud.generalsaas.face.source.certify", "POST", "/v3/datadigital/fincloud/generalsaas/face/source/certify", core.ParameterModeMultipart, func(s *Service) error {
			_, err := s.GeneralSaaSFaceSourceCertify(context.Background(), GeneralSaaSFaceSourceCertifyRequest{})
			return err
		}},
		{"datadigital.fincloud.generalsaas.face.verification.initialize", "POST", "/v3/datadigital/fincloud/generalsaas/face/verification/initialize", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.GeneralSaaSFaceVerificationInitialize(context.Background(), GeneralSaaSFaceVerificationInitializeRequest{})
			return err
		}},
		{"datadigital.fincloud.generalsaas.face.verification.query", "GET", "/v3/datadigital/fincloud/generalsaas/face/verification/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.GeneralSaaSFaceVerificationQuery(context.Background(), GeneralSaaSFaceVerificationQueryRequest{})
			return err
		}},
		{"datadigital.fincloud.generalsaas.ocr.mobile.initialize", "POST", "/v3/datadigital/fincloud/generalsaas/ocr/mobile/initialize", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.GeneralSaaSOCRMobileInitialize(context.Background(), GeneralSaaSOCRMobileInitializeRequest{})
			return err
		}},
		{"datadigital.fincloud.generalsaas.ocr.server.detect", "POST", "/v3/datadigital/fincloud/generalsaas/ocr/server/detect", core.ParameterModeMultipart, func(s *Service) error {
			_, err := s.GeneralSaaSOCRServerDetect(context.Background(), GeneralSaaSOCRServerDetectRequest{})
			return err
		}}}
	for _, tc := range cases {
		t.Run(tc.operation, func(t *testing.T) {
			e := &recordingExecutor{}
			s, err := NewService(e)
			if err != nil {
				t.Fatal(err)
			}
			if err = tc.call(s); err != nil {
				t.Fatal(err)
			}
			if e.req.Spec.Method != tc.operation || e.req.Spec.HTTPMethod.String() != tc.method || e.req.Spec.Path != tc.path || e.req.Spec.ParameterMode != tc.mode || e.req.Params.Mode != tc.mode {
				t.Fatalf("unexpected spec: %+v", e.req.Spec)
			}
		})
	}
}
