package overseas_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade/overseas"
)

// recordingExecutor 记录 WireRequest，返回预置结果，
// 用于验证接口名、参数位置与响应映射（不发起网络请求）。
type recordingExecutor struct {
	req    core.WireRequest
	called bool

	result *core.APIResult
	err    error
}

func (e *recordingExecutor) Do(_ context.Context, req core.WireRequest) (*core.APIResult, error) {
	e.called = true
	e.req = req
	return e.result, e.err
}

func verifiedResult(node string) *core.APIResult {
	return &core.APIResult{
		HTTPStatus:        200,
		SignatureVerified: true,
		Node:              []byte(node),
	}
}

func errorResult(err error) *core.APIResult { return &core.APIResult{HTTPStatus: 200} }

func newRecordingService(t *testing.T) (*overseas.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{}
	svc, err := overseas.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func ptr[T any](v T) *T { return &v }

var _ = errorResult

// alipayPublicKeyFromCert 从证书 PEM 提取 RSA 公钥。
