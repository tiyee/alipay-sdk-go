package trans_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/fund/trans"
)

func TestFundValidations(t *testing.T) {
	tests := []struct {
		name string
		call func(*trans.Service) error
	}{
		{
			name: "转账缺少收款方",
			call: func(s *trans.Service) error {
				_, err := s.UniTransfer(context.Background(), trans.UniTransferRequest{
					OutBizNo: "T1", TransAmount: "1.00", ProductCode: "TRANS_ACCOUNT_NO_PWD"})
				return err
			},
		},
		{
			name: "转账金额为零",
			call: func(s *trans.Service) error {
				_, err := s.UniTransfer(context.Background(), trans.UniTransferRequest{
					OutBizNo: "T1", TransAmount: "0.00", ProductCode: "TRANS_ACCOUNT_NO_PWD",
					PayeeInfo: &trans.Participant{Identity: ptr("1"), IdentityType: ptr("ALIPAY_USER_ID")}})
				return err
			},
		},
		{
			name: "到账转账金额低于下限",
			call: func(s *trans.Service) error {
				_, err := s.ToAccountTransfer(context.Background(), trans.ToAccountTransferRequest{
					OutBizNo: "T1", TransAmount: "0.01",
					PayeeAccount: "2088001", PayeeType: trans.PayeeTypeUserID})
				return err
			},
		},
		{
			name: "到账转账未知账户类型",
			call: func(s *trans.Service) error {
				_, err := s.ToAccountTransfer(context.Background(), trans.ToAccountTransferRequest{
					OutBizNo: "T1", TransAmount: "1.00",
					PayeeAccount: "x", PayeeType: trans.PayeeAccountType("BANK")})
				return err
			},
		},
		{
			name: "退回缺少请求号",
			call: func(s *trans.Service) error {
				_, err := s.Refund(context.Background(), trans.RefundRequest{
					OrderID: "O1", RefundAmount: "1.00"})
				return err
			},
		},
		{
			name: "订单查询缺少单据",
			call: func(s *trans.Service) error {
				_, err := s.OrderQuery(context.Background(), trans.OrderQueryRequest{})
				return err
			},
		},
		{
			name: "通用查询缺产品码",
			call: func(s *trans.Service) error {
				_, err := s.CommonQuery(context.Background(), trans.CommonQueryRequest{
					OutBizNo: ptr("T1")})
				return err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, executor := newService(t)
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
