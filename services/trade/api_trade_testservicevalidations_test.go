package trade_test

import (
	"context"
	"errors"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestServiceValidations(t *testing.T) {
	tests := []struct {
		name string
		call func(*trade.Service) error
	}{
		{
			name: "查询缺少订单标识",
			call: func(s *trade.Service) error {
				_, err := s.Query(context.Background(), trade.QueryRequest{})
				return err
			},
		},
		{
			name: "预下单缺少金额",
			call: func(s *trade.Service) error {
				_, err := s.PreCreate(context.Background(), trade.PreCreateRequest{
					OutTradeNo: "T1", Subject: "标题"})
				return err
			},
		},
		{
			name: "预下单标题含特殊字符",
			call: func(s *trade.Service) error {
				_, err := s.PreCreate(context.Background(), trade.PreCreateRequest{
					OutTradeNo: "T1", TotalAmount: "1.00", Subject: "a/b"})
				return err
			},
		},
		{
			name: "预下单订单号含非法字符",
			call: func(s *trade.Service) error {
				_, err := s.PreCreate(context.Background(), trade.PreCreateRequest{
					OutTradeNo: "T-1", TotalAmount: "1.00", Subject: "标题"})
				return err
			},
		},
		{
			name: "关单缺少订单标识",
			call: func(s *trade.Service) error {
				_, err := s.Close(context.Background(), trade.CloseRequest{})
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
