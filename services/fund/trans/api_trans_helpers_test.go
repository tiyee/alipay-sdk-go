package trans_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/fund/trans"
)

type recordingExecutor struct {
	req    core.WireRequest
	called bool
	result *core.APIResult
}

func (e *recordingExecutor) Do(_ context.Context, req core.WireRequest) (*core.APIResult, error) {
	e.called = true
	e.req = req
	return e.result, nil
}

func newService(t *testing.T) (*trans.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{result: &core.APIResult{
		HTTPStatus:        200,
		SignatureVerified: true,
		Node:              []byte(`{}`),
	}}
	svc, err := trans.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func ptr[T any](v T) *T { return &v }

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(sub) == 0 ||
		(func() bool {
			for i := 0; i+len(sub) <= len(s); i++ {
				if s[i:i+len(sub)] == sub {
					return true
				}
			}
			return false
		})())
}
