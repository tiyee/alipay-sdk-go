package trans_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/fund/trans"
)

func TestTransferSpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*trans.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
		wantGET    bool
	}{
		{
			name: "单笔转账",
			call: func(s *trans.Service, e *recordingExecutor) error {
				_, err := s.UniTransfer(context.Background(), trans.UniTransferRequest{
					OutBizNo: "T20261105001", TransAmount: "10.00", ProductCode: "TRANS_ACCOUNT_NO_PWD",
					PayeeInfo: &trans.Participant{
						Identity: ptr("2088101117955611"), IdentityType: ptr("ALIPAY_USER_ID"),
					}})
				return err
			},
			wantMethod: "alipay.fund.trans.uni.transfer",
			wantPath:   "/v3/alipay/fund/trans/uni/transfer",
		},
		{
			name: "转账到支付宝账户",
			call: func(s *trans.Service, e *recordingExecutor) error {
				_, err := s.ToAccountTransfer(context.Background(), trans.ToAccountTransferRequest{
					OutBizNo: "T20261105001", TransAmount: "1.00",
					PayeeAccount: "2088101117955611", PayeeType: trans.PayeeTypeUserID})
				return err
			},
			wantMethod: "alipay.fund.trans.toaccount.transfer",
			wantPath:   "/v3/alipay/fund/trans/toaccount/transfer",
		},
		{
			name: "资金退回",
			call: func(s *trans.Service, e *recordingExecutor) error {
				_, err := s.Refund(context.Background(), trans.RefundRequest{
					OrderID: "O1", OutRequestNo: "R1", RefundAmount: "1.00", BizScene: "REFUND"})
				return err
			},
			wantMethod: "alipay.fund.trans.refund",
			wantPath:   "/v3/alipay/fund/trans/refund",
		},
		{
			name: "订单查询",
			call: func(s *trans.Service, e *recordingExecutor) error {
				_, err := s.OrderQuery(context.Background(), trans.OrderQueryRequest{
					OrderID: ptr("O1")})
				return err
			},
			wantMethod: "alipay.fund.trans.order.query",
			wantPath:   "/v3/alipay/fund/trans/order/query",
			wantGET:    true,
		},
		{
			name: "通用查询",
			call: func(s *trans.Service, e *recordingExecutor) error {
				_, err := s.CommonQuery(context.Background(), trans.CommonQueryRequest{
					OrderID: ptr("O1")})
				return err
			},
			wantMethod: "alipay.fund.trans.common.query",
			wantPath:   "/v3/alipay/fund/trans/common/query",
			wantGET:    true,
		},
		{
			name: "收款账号绑定查询",
			call: func(s *trans.Service, e *recordingExecutor) error {
				_, err := s.PayeeBindQuery(context.Background(), trans.PayeeBindQueryRequest{
					Identity: "110101199001011234"})
				return err
			},
			wantMethod: "alipay.fund.trans.payee.bind.query",
			wantPath:   "/v3/alipay/fund/trans/payee/bind/query",
			wantGET:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, executor := newService(t)
			if err := tt.call(svc, executor); err != nil {
				t.Fatalf("call error = %v", err)
			}
			spec := executor.req.Spec
			if spec.Method != tt.wantMethod || spec.Path != tt.wantPath {
				t.Fatalf("spec = %s %s, want %s %s", spec.Method, spec.Path, tt.wantMethod, tt.wantPath)
			}
			wantMode := core.ParameterModeJSONBody
			if tt.wantGET {
				wantMode = core.ParameterModeQuery
			}
			if spec.ParameterMode != wantMode {
				t.Errorf("parameter mode = %v, want %v", spec.ParameterMode, wantMode)
			}
		})
	}
}
