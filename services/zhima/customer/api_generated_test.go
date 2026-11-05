package customer

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
		{"zhima.customer.jobworth.authentication.query", "GET", "/v3/zhima/customer/jobworth/authentication", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.JobWorthAuthenticationQuery(context.Background(), JobWorthAuthenticationQueryRequest{})
			return err
		}},
		{"zhima.customer.jobworth.cloudresume.query", "GET", "/v3/zhima/customer/jobworth/cloudresume/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.JobWorthCloudResumeQuery(context.Background(), JobWorthCloudResumeQueryRequest{})
			return err
		}},
		{"zhima.customer.jobworth.info.query", "GET", "/v3/zhima/customer/jobworth/info/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.JobWorthInfoQuery(context.Background(), JobWorthInfoQueryRequest{})
			return err
		}},
		{"zhima.customer.jobworth.jobdata.add", "POST", "/v3/zhima/customer/jobworth/jobdata/add", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.JobWorthJobDataAdd(context.Background(), JobWorthJobDataAddRequest{})
			return err
		}},
		{"zhima.customer.jobworth.picture.upload", "POST", "/v3/zhima/customer/jobworth/picture/upload", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.JobWorthPictureUpload(context.Background(), JobWorthPictureUploadRequest{})
			return err
		}},
		{"zhima.customer.jobworth.position.add", "POST", "/v3/zhima/customer/jobworth/position/add", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.JobWorthPositionAdd(context.Background(), JobWorthPositionAddRequest{})
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
