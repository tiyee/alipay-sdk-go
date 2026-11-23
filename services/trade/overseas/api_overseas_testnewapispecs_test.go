package overseas_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade/overseas"
)

func TestNewAPISpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*overseas.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
		wantGET    bool
	}{
		{
			name: "跨境结算",
			call: func(s *overseas.Service, e *recordingExecutor) error {
				_, err := s.OverseasSettle(context.Background(), overseas.OverseasSettleRequest{
					TradeNo: "T1", OutRequestNo: "R1", Amount: "10.00", ForeignSettleCurrency: "USD"})
				return err
			},
			wantMethod: "alipay.trade.overseas.settle",
			wantPath:   "/v3/alipay/trade/overseas/settle",
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
