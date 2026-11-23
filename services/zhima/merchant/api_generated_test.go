package merchant

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
		{"zhima.merchant.subsidiaries.apply", "POST", "/v3/zhima/merchant/subsidiaries/apply", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.SubsidiariesApply(context.Background(), SubsidiariesApplyRequest{})
			return err
		}},
		{"zhima.merchant.subsidiaries.close", "POST", "/v3/zhima/merchant/subsidiaries/close", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.SubsidiariesClose(context.Background(), SubsidiariesCloseRequest{})
			return err
		}},
		{"zhima.merchant.zmgo.cumulate.query", "GET", "/v3/zhima/merchant/zmgo/cumulate/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ZmGoCumulateQuery(context.Background(), ZmGoCumulateQueryRequest{})
			return err
		}},
		{"zhima.merchant.zmgo.cumulate.sync", "POST", "/v3/zhima/merchant/zmgo/cumulate/sync", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ZmGoCumulateSync(context.Background(), ZmGoCumulateSyncRequest{})
			return err
		}},
		{"zhima.merchant.zmgo.template.create", "POST", "/v3/zhima/merchant/zmgo/template/create", core.ParameterModeJSONBody, func(s *Service) error {
			_, err := s.ZmGoTemplateCreate(context.Background(), ZmGoTemplateCreateRequest{})
			return err
		}},
		{"zhima.merchant.zmgo.template.query", "GET", "/v3/zhima/merchant/zmgo/template/query", core.ParameterModeQuery, func(s *Service) error {
			_, err := s.ZmGoTemplateQuery(context.Background(), ZmGoTemplateQueryRequest{})
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
