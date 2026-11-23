package trade_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

// TestNewAPIValidations 覆盖新接口的请求校验分支。
func TestNewAPIValidations(t *testing.T) {
	tests := []struct {
		name string
		call func(*trade.Service) error
	}{
		{
			name: "支付缺少付款码",
			call: func(s *trade.Service) error {
				_, err := s.Pay(context.Background(), trade.PayRequest{
					OutTradeNo: "T1", TotalAmount: "1.00", Subject: "标题",
					Scene: ptr(trade.PaySceneBarCode)})
				return err
			},
		},
		{
			name: "支付场景未知",
			call: func(s *trade.Service) error {
				_, err := s.Pay(context.Background(), trade.PayRequest{
					OutTradeNo: "T1", TotalAmount: "1.00", Subject: "标题",
					Scene: ptr(trade.PayScene("scan")), AuthCode: ptr("123")})
				return err
			},
		},
		{
			name: "创建缺少买家标识",
			call: func(s *trade.Service) error {
				_, err := s.Create(context.Background(), trade.CreateRequest{
					OutTradeNo: "T1", TotalAmount: "1.00", Subject: "标题"})
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
