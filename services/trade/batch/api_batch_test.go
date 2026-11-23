package batch_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade/batch"
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

func newRecordingService(t *testing.T) (*batch.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{}
	svc, err := batch.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func TestNewServiceRejectsNilExecutor(t *testing.T) {
	if _, err := batch.NewService(nil); err == nil {
		t.Fatal("NewService(nil) must fail")
	}
}

func ptr[T any](v T) *T { return &v }

func TestNewAPISpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*batch.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
		wantGET    bool
	}{
		{
			name: "批量结算",
			call: func(s *batch.Service, e *recordingExecutor) error {
				_, err := s.BatchSettle(context.Background(), batch.BatchSettleRequest{
					OutRequestNo: "B1",
					SettleClauses: []batch.SettleClause{
						{SettleAccountType: ptr("defaultSettle"), Amount: ptr("10.00"), Currency: ptr("CNY")},
					}})
				return err
			},
			wantMethod: "alipay.trade.batch.settle",
			wantPath:   "/v3/alipay/trade/batch/settle",
		},
		{
			name: "批量结算查询",
			call: func(s *batch.Service, e *recordingExecutor) error {
				_, err := s.BatchSettleQuery(context.Background(), batch.BatchSettleQueryRequest{
					SettleNo: ptr("S1")})
				return err
			},
			wantMethod: "alipay.trade.batch.settle.query",
			wantPath:   "/v3/alipay/trade/batch/settle/query",
		},
		{
			name: "请款分账",
			call: func(s *batch.Service, e *recordingExecutor) error {
				_, err := s.BatchTransfer(context.Background(), batch.BatchTransferRequest{
					OutRequestNo:      "R1",
					RoyaltyParameters: []batch.BatchRoyaltyClause{{Amount: ptr("1.00")}}})
				return err
			},
			wantMethod: "alipay.trade.batch.transfer",
			wantPath:   "/v3/alipay/trade/batch/transfer",
		},
		{
			name: "请款分账查询",
			call: func(s *batch.Service, e *recordingExecutor) error {
				_, err := s.BatchTransferQuery(context.Background(), batch.BatchTransferQueryRequest{
					SettleNo: ptr("S1")})
				return err
			},
			wantMethod: "alipay.trade.batch.transfer.query",
			wantPath:   "/v3/alipay/trade/batch/transfer/query",
			wantGET:    true,
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
		call func(*batch.Service) error
	}{
		{
			name: "批量结算缺少条款",
			call: func(s *batch.Service) error {
				_, err := s.BatchSettle(context.Background(), batch.BatchSettleRequest{
					OutRequestNo: "B1"})
				return err
			},
		},
		{
			name: "请款分账缺少条款",
			call: func(s *batch.Service) error {
				_, err := s.BatchTransfer(context.Background(), batch.BatchTransferRequest{
					OutRequestNo: "R1"})
				return err
			},
		},
		{
			name: "请款分账查询缺少单据",
			call: func(s *batch.Service) error {
				_, err := s.BatchTransferQuery(context.Background(), batch.BatchTransferQueryRequest{})
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
