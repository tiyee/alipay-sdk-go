package bill_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/services/data/bill"
)

func TestBillSpecs(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*bill.Service, *recordingExecutor) error
		wantMethod string
		wantPath   string
		wantQuery  string
	}{
		{
			name: "对账单下载地址",
			call: func(s *bill.Service, e *recordingExecutor) error {
				_, err := s.BillDownloadURLQuery(context.Background(), bill.BillDownloadURLRequest{
					BillType: "trade", BillDate: "2026-11-04"})
				return err
			},
			wantMethod: "alipay.data.dataservice.bill.downloadurl.query",
			wantPath:   "/v3/alipay/data/dataservice/bill/downloadurl/query",
			wantQuery:  "bill_date=2026-11-04&bill_type=trade",
		},
		{
			name: "当前余额",
			call: func(s *bill.Service, e *recordingExecutor) error {
				_, err := s.BillBalance(context.Background(), bill.BillBalanceRequest{
					BillUserID: ptr("2088001")})
				return err
			},
			wantMethod: "alipay.data.bill.balance.query",
			wantPath:   "/v3/alipay/data/bill/balance/query",
			wantQuery:  "bill_user_id=2088001",
		},
		{
			name: "历史余额",
			call: func(s *bill.Service, e *recordingExecutor) error {
				_, err := s.BillBalanceHis(context.Background(), bill.BillBalanceHisRequest{
					BizDate: ptr("2026-11-04")})
				return err
			},
			wantMethod: "alipay.data.bill.balancehis.query",
			wantPath:   "/v3/alipay/data/bill/balancehis/query",
			wantQuery:  "biz_date=2026-11-04",
		},
		{
			name: "账务明细",
			call: func(s *bill.Service, e *recordingExecutor) error {
				_, err := s.BillAccountLog(context.Background(), bill.BillAccountLogRequest{
					AlipayOrderNo: ptr("2026110522001400001"), PageNo: ptr("1")})
				return err
			},
			wantMethod: "alipay.data.bill.accountlog.query",
			wantPath:   "/v3/alipay/data/bill/accountlog/query",
			wantQuery:  "alipay_order_no=2026110522001400001&page_no=1",
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
			if spec.HTTPMethod != core.HTTPMethodGet ||
				spec.ParameterMode != core.ParameterModeQuery {
				t.Fatalf("unexpected protocol metadata: %+v", spec)
			}
			if got := executor.req.Params.Form.Encode(); got != tt.wantQuery {
				t.Fatalf("query = %q, want %q", got, tt.wantQuery)
			}
		})
	}
}
