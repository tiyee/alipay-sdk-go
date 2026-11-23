package bizrisks

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
		{"grandsecurity.bizrisks.fact.check", "POST", "/v3/stream/grandsecurity/bizrisks/fact/check", core.ParameterModeJSONBody, func(s *Service) error { _, err := s.FactCheck(context.Background(), FactCheckRequest{}); return err }}}
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
