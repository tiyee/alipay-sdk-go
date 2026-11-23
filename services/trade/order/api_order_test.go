package order_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade/order"
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

func newRecordingService(t *testing.T) (*order.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{}
	svc, err := order.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func TestNewServiceRejectsNilExecutor(t *testing.T) {
	if _, err := order.NewService(nil); err == nil {
		t.Fatal("NewService(nil) must fail")
	}
}

func ptr[T any](v T) *T { return &v }

func TestNewAPISpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*order.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
		wantGET    bool
	}{
		{
			name: "订单支付",
			call: func(s *order.Service, e *recordingExecutor) error {
				_, err := s.OrderPay(context.Background(), order.OrderPayRequest{
					TradeNo: "T1", TotalAmount: "10.00", OutRequestNo: "R1"})
				return err
			},
			wantMethod: "alipay.trade.order.pay",
			wantPath:   "/v3/alipay/trade/order/pay",
		},
		{
			name: "交易结算",
			call: func(s *order.Service, e *recordingExecutor) error {
				_, err := s.OrderSettle(context.Background(), order.OrderSettleRequest{
					TradeNo: "T1", OutRequestNo: "S1"})
				return err
			},
			wantMethod: "alipay.trade.order.settle",
			wantPath:   "/v3/alipay/trade/order/settle",
		},
		{
			name: "分账查询",
			call: func(s *order.Service, e *recordingExecutor) error {
				_, err := s.OrderSettleQuery(context.Background(), order.OrderSettleQueryRequest{
					SettleNo: ptr("S1")})
				return err
			},
			wantMethod: "alipay.trade.order.settle.query",
			wantPath:   "/v3/alipay/trade/order/settle/query",
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

// TestNewAPIUnknownEnumPreserved 验证未知分账状态原样保留。
func TestNewAPIUnknownEnumPreserved(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(`{"settle_no":"S1","royalty_detail_list":[{"state":"NEW_STATE","amount":"1.00"}]}`)
	resp, err := svc.OrderSettleQuery(context.Background(), order.OrderSettleQueryRequest{SettleNo: ptr("S1")})
	if err != nil {
		t.Fatalf("OrderSettleQuery() error = %v", err)
	}
	if len(resp.Data.RoyaltyDetailList) != 1 {
		t.Fatalf("details = %+v", resp.Data.RoyaltyDetailList)
	}
	state := resp.Data.RoyaltyDetailList[0].State
	if state == nil || *state != order.RoyaltyDetailState("NEW_STATE") {
		t.Fatalf("unknown state = %+v, want preserved", state)
	}
}

// TestNewAPIValidations 覆盖新接口的请求校验分支。
func TestNewAPIValidations(t *testing.T) {
	tests := []struct {
		name string
		call func(*order.Service) error
	}{
		{
			name: "订单支付缺少请求号",
			call: func(s *order.Service) error {
				_, err := s.OrderPay(context.Background(), order.OrderPayRequest{
					TradeNo: "T1", TotalAmount: "1.00"})
				return err
			},
		},
		{
			name: "分账查询缺少单据标识",
			call: func(s *order.Service) error {
				_, err := s.OrderSettleQuery(context.Background(), order.OrderSettleQueryRequest{
					OutRequestNo: ptr("R1")})
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
