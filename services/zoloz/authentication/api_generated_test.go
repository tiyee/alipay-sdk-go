package authentication

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
		{"zoloz.authentication.customer.facemanage.create", "POST", "/v3/zoloz/authentication/customer/facemanage/create", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.CustomerFaceManageCreate(context.Background(), CustomerFaceManageCreateRequest{})
			return err
		}},
		{"zoloz.authentication.customer.facemanage.delete", "POST", "/v3/zoloz/authentication/customer/facemanage/delete", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.CustomerFaceManageDelete(context.Background(), CustomerFaceManageDeleteRequest{})
			return err
		}},
		{"zoloz.authentication.customer.ftoken.query", "POST", "/v3/zoloz/authentication/customer/ftoken/query", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.CustomerFTokenQuery(context.Background(), CustomerFTokenQueryRequest{})
			return err
		}},
		{"zoloz.authentication.customer.smilepay.initialize", "POST", "/v3/zoloz/authentication/customer/smilepay/initialize", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.CustomerSmilePayInitialize(context.Background(), CustomerSmilePayInitializeRequest{})
			return err
		}},
		{"zoloz.authentication.smilepay.initialize", "POST", "/v3/zoloz/authentication/smilepay/initialize", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.SmilePayInitialize(context.Background(), SmilePayInitializeRequest{})
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
