package trade_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestNewAPISpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*trade.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
		wantGET    bool
	}{
		{
			name: "支付",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.Pay(context.Background(), trade.PayRequest{
					OutTradeNo: "T1", TotalAmount: "10.00", Subject: "标题",
					Scene: ptr(trade.PaySceneBarCode), AuthCode: ptr("28763443825664394")})
				return err
			},
			wantMethod: "alipay.trade.pay",
			wantPath:   "/v3/alipay/trade/pay",
		},
		{
			name: "创建订单",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.Create(context.Background(), trade.CreateRequest{
					OutTradeNo: "T1", TotalAmount: "10.00", Subject: "标题", BuyerID: ptr("2088001")})
				return err
			},
			wantMethod: "alipay.trade.create",
			wantPath:   "/v3/alipay/trade/create",
		},
		{
			name: "撤销",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.Cancel(context.Background(), trade.CancelRequest{TradeNo: "T1"})
				return err
			},
			wantMethod: "alipay.trade.cancel",
			wantPath:   "/v3/alipay/trade/cancel",
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
