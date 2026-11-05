package refund_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade/refund"
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

func newRecordingService(t *testing.T) (*refund.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{}
	svc, err := refund.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func TestNewServiceRejectsNilExecutor(t *testing.T) {
	if _, err := refund.NewService(nil); err == nil {
		t.Fatal("NewService(nil) must fail")
	}
}

func ptr[T any](v T) *T { return &v }

func TestOperationSpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*refund.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
	}{
		{
			name: "退款",
			call: func(s *refund.Service, e *recordingExecutor) error {
				_, err := s.Refund(context.Background(), refund.RefundRequest{
					OutTradeNo: "T1", RefundAmount: "10.00"})
				return err
			},
			wantMethod: "alipay.trade.refund",
			wantPath:   "/v3/alipay/trade/refund",
		},
		{
			name: "退款查询",
			call: func(s *refund.Service, e *recordingExecutor) error {
				_, err := s.QueryRefund(context.Background(), refund.QueryRefundRequest{
					OutRequestNo: "R1", TradeNo: "T1"})
				return err
			},
			wantMethod: "alipay.trade.fastpay.refund.query",
			wantPath:   "/v3/alipay/trade/fastpay/refund/query",
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

func TestRefundResponseFundChangeSemantics(t *testing.T) {
	// code=10000 仅代表请求成功；SDK 不把 fund_change=N 判定为失败，
	// 语义交由调用方依据文档判断。
	for _, tt := range []struct {
		fundChange string
		want       refund.FundChange
	}{
		{fundChange: "Y", want: refund.FundChangeYes},
		{fundChange: "N", want: refund.FundChangeNo},
	} {
		node := fmt.Sprintf(`{"code":"10000","msg":"Success","trade_no":"T1","refund_fee":"12.33","fund_change":%q}`, tt.fundChange)
		svc, executor := newRecordingService(t)
		executor.result = verifiedResult(node)
		resp, err := svc.Refund(context.Background(), refund.RefundRequest{
			OutTradeNo: "T1", RefundAmount: "12.33", OutRequestNo: ptr("R1")})
		if err != nil {
			t.Fatalf("Refund() error = %v", err)
		}
		if resp.Data.FundChange == nil || *resp.Data.FundChange != tt.want {
			t.Fatalf("FundChange = %+v, want %q", resp.Data.FundChange, tt.want)
		}
		if executor.req.Params.Mode != core.ParameterModeJSONBody {
			t.Fatal("refund must use the v3 json body")
		}
	}
}

func TestQueryRefundResponseMapping(t *testing.T) {
	node := `{"code":"10000","msg":"Success","trade_no":"T1","out_trade_no":"O1",` +
		`"out_request_no":"R1","refund_status":"REFUND_SUCCESS",` +
		`"refund_amount":"12.33","total_amount":"100.20"}`
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(node)
	resp, err := svc.QueryRefund(context.Background(), refund.QueryRefundRequest{
		OutRequestNo: "R1", TradeNo: "T1"})
	if err != nil {
		t.Fatalf("QueryRefund() error = %v", err)
	}
	if resp.Data.RefundStatus == nil || *resp.Data.RefundStatus != refund.RefundStatusSuccess {
		t.Fatalf("RefundStatus = %+v", resp.Data.RefundStatus)
	}
}

func TestServiceValidations(t *testing.T) {
	tests := []struct {
		name string
		call func(*refund.Service) error
	}{
		{
			name: "退款缺少金额",
			call: func(s *refund.Service) error {
				_, err := s.Refund(context.Background(), refund.RefundRequest{TradeNo: "T1"})
				return err
			},
		},
		{
			name: "退款查询缺少请求号",
			call: func(s *refund.Service) error {
				_, err := s.QueryRefund(context.Background(), refund.QueryRefundRequest{TradeNo: "T1"})
				return err
			},
		},
		{
			name: "金额超过两位小数",
			call: func(s *refund.Service) error {
				_, err := s.Refund(context.Background(), refund.RefundRequest{
					TradeNo: "T1", RefundAmount: "1.234"})
				return err
			},
		},
		{
			name: "金额为零",
			call: func(s *refund.Service) error {
				_, err := s.Refund(context.Background(), refund.RefundRequest{
					TradeNo: "T1", RefundAmount: "0.00"})
				return err
			},
		},
		{
			name: "金额为负",
			call: func(s *refund.Service) error {
				_, err := s.Refund(context.Background(), refund.RefundRequest{
					TradeNo: "T1", RefundAmount: "-1.00"})
				return err
			},
		},
		{
			name: "金额非数字",
			call: func(s *refund.Service) error {
				_, err := s.Refund(context.Background(), refund.RefundRequest{
					TradeNo: "T1", RefundAmount: "12.3a"})
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

func TestNewAPISpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*refund.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
		wantGET    bool
	}{
		{
			name: "异步退款受理",
			call: func(s *refund.Service, e *recordingExecutor) error {
				_, err := s.RefundApply(context.Background(), refund.RefundApplyRequest{
					TradeNo: "T1", RefundAmount: "1.00", OutRequestNo: ptr("R1")})
				return err
			},
			wantMethod: "alipay.trade.refund.apply",
			wantPath:   "/v3/alipay/trade/refund/apply",
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

func TestRefundApplyResponseMapping(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(`{"trade_no":"T1","out_request_no":"R1","refund_amount":"1.00","refund_status":"REFUND_PROCESSING"}`)
	resp2, err := svc.RefundApply(context.Background(), refund.RefundApplyRequest{
		TradeNo: "T1", RefundAmount: "1.00", OutRequestNo: ptr("R1")})
	if err != nil {
		t.Fatalf("RefundApply() error = %v", err)
	}
	if resp2.Data.RefundStatus == nil || *resp2.Data.RefundStatus != refund.RefundApplyStatusProcessing {
		t.Fatalf("RefundStatus = %+v", resp2.Data.RefundStatus)
	}
}

// TestNewAPIValidations 覆盖新接口的请求校验分支。
func TestNewAPIValidations(t *testing.T) {
	tests := []struct {
		name string
		call func(*refund.Service) error
	}{
		{
			name: "异步退款缺少请求号",
			call: func(s *refund.Service) error {
				_, err := s.RefundApply(context.Background(), refund.RefundApplyRequest{
					TradeNo: "T1", RefundAmount: "1.00"})
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
