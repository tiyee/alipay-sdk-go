package trade_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestQueryRequestSerialization(t *testing.T) {
	tests := []struct {
		name string
		req  trade.QueryRequest
		want string
	}{
		{
			name: "仅商户订单号",
			req:  trade.QueryRequest{OutTradeNo: "20150320010101001"},
			want: `{"out_trade_no":"20150320010101001"}`,
		},
		{
			name: "交易号与查询选项",
			req: trade.QueryRequest{
				TradeNo:      "2026010122001400001234567890",
				QueryOptions: []trade.QueryOption{trade.QueryOptionFundBillList, trade.QueryOptionTradeSettleInfo},
			},
			want: `{"trade_no":"2026010122001400001234567890","query_options":["fund_bill_list","trade_settle_info"]}`,
		},
		{
			name: "中文与可选字段未传时省略",
			req:  trade.QueryRequest{OutTradeNo: "订单-001"},
			want: `{"out_trade_no":"订单-001"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, executor := newRecordingService(t)
			executor.result = verifiedResult(`{"code":"10000","msg":"Success"}`)
			if _, err := svc.Query(context.Background(), tt.req); err != nil {
				t.Fatalf("Query() error = %v", err)
			}
			if got := string(executor.req.Params.JSON); got != tt.want {
				t.Fatalf("biz_content = %s, want %s", got, tt.want)
			}
		})
	}
}
