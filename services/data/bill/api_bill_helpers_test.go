package bill_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/data/bill"
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

func newService(t *testing.T) (*bill.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{result: &core.APIResult{
		HTTPStatus:        200,
		SignatureVerified: true,
		Node:              []byte(`{}`),
	}}
	svc, err := bill.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func ptr[T any](v T) *T { return &v }
