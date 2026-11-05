package settle_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
	"github.com/tiyee/alipay-sdk-go/services/trade/settle"
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

func newRecordingService(t *testing.T) (*settle.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{}
	svc, err := settle.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func TestNewServiceRejectsNilExecutor(t *testing.T) {
	if _, err := settle.NewService(nil); err == nil {
		t.Fatal("NewService(nil) must fail")
	}
}

func ptr[T any](v T) *T { return &v }

func TestNewAPISpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*settle.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
		wantGET    bool
	}{
		{
			name: "确认结算",
			call: func(s *settle.Service, e *recordingExecutor) error {
				_, err := s.SettleConfirm(context.Background(), settle.SettleConfirmRequest{
					TradeNo: "T1", OutRequestNo: "C1",
					SettleInfo: trade.SettleInfo{SettleDetailInfos: []trade.SettleDetailInfo{
						{TransInType: "userId", TransIn: ptr("2088001"), Amount: ptr("10.00")},
					}}})
				return err
			},
			wantMethod: "alipay.trade.settle.confirm",
			wantPath:   "/v3/alipay/trade/settle/confirm",
		},
		{
			name: "待结算查询",
			call: func(s *settle.Service, e *recordingExecutor) error {
				_, err := s.SettleReceivablesQuery(context.Background(), settle.SettleReceivablesQueryRequest{
					TradeNo: ptr("T1")})
				return err
			},
			wantMethod: "alipay.trade.settle.receivables.query",
			wantPath:   "/v3/alipay/trade/settle/receivables/query",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, executor := newRecordingService(t)
			executor.result = verifiedResult(`{}`)
			if err := tt.call(svc, executor); err != nil {
				t.Fatalf("call error = %v", err)
			}
			spec := executor.req.Spec
			if spec.Method != tt.wantMethod || spec.Path != tt.wantPath {
				t.Fatalf("spec = %s %s, want %s %s", spec.Method, spec.Path, tt.wantMethod, tt.wantPath)
			}
			wantMethod := core.HTTPMethodPost
			if tt.wantGET {
				wantMethod = core.HTTPMethodGet
			}
			if spec.HTTPMethod != wantMethod {
				t.Errorf("HTTPMethod = %v, want %v", spec.HTTPMethod, wantMethod)
			}
			wantMode := core.ParameterModeJSONBody
			if tt.wantGET {
				wantMode = core.ParameterModeQuery
			}
			if spec.ParameterMode != wantMode || executor.req.Params.Mode != wantMode {
				t.Errorf("parameter mode = %v / %v, want %v", spec.ParameterMode, executor.req.Params.Mode, wantMode)
			}
			if executor.result.HTTPStatus != 200 {
				t.Error("result must be passed through")
			}
		})
	}
}

// TestNewAPIValidations 覆盖新接口的请求校验分支。
func TestNewAPIValidations(t *testing.T) {
	tests := []struct {
		name string
		call func(*settle.Service) error
	}{
		{
			name: "确认结算缺少结算明细",
			call: func(s *settle.Service) error {
				_, err := s.SettleConfirm(context.Background(), settle.SettleConfirmRequest{
					TradeNo: "T1", OutRequestNo: "C1"})
				return err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, executor := newRecordingService(t)
			err := tt.call(svc)
			var valErr *core.ValidationError
			if !errors.As(err, &valErr) {
				t.Fatalf("error = %v, want *core.ValidationError", err)
			}
			if executor.called {
				t.Error("validation failure must not send request")
			}
		})
	}
}
