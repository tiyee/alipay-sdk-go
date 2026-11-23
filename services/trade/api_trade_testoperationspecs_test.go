package trade_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestOperationSpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*trade.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
	}{
		{
			name: "查询",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"})
				return err
			},
			wantMethod: "alipay.trade.query",
			wantPath:   "/v3/alipay/trade/query",
		},
		{
			name: "预下单",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.PreCreate(context.Background(), trade.PreCreateRequest{
					OutTradeNo: "T1", TotalAmount: "88.88", Subject: "测试商品"})
				return err
			},
			wantMethod: "alipay.trade.precreate",
			wantPath:   "/v3/alipay/trade/precreate",
		},
		{
			name: "关单",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.Close(context.Background(), trade.CloseRequest{TradeNo: "T1"})
				return err
			},
			wantMethod: "alipay.trade.close",
			wantPath:   "/v3/alipay/trade/close",
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
			if spec.Method != tt.wantMethod {
				t.Errorf("Method = %q, want %q", spec.Method, tt.wantMethod)
			}
			if spec.HTTPMethod != core.HTTPMethodPost {
				t.Errorf("HTTPMethod = %v", spec.HTTPMethod)
			}
			if spec.Path != tt.wantPath {
				t.Errorf("Path = %q, want %q", spec.Path, tt.wantPath)
			}
			if spec.ParameterMode != core.ParameterModeJSONBody {
				t.Errorf("ParameterMode = %v", spec.ParameterMode)
			}
			if executor.req.Params.Mode != core.ParameterModeJSONBody {
				t.Errorf("params mode = %v", executor.req.Params.Mode)
			}
		})
	}
}
