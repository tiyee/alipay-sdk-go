package bill_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/data/bill"
)

func TestBillAccountLogMapping(t *testing.T) {
	svc, executor := newService(t)
	executor.result.Node = []byte(`{"detail_list":[{"account_log_id":"L1","trans_amount":"-10.00","direction":"支出","balance":"100.00"}],"total_size":"1","page_no":"1","page_size":"2000"}`)
	resp, err := svc.BillAccountLog(context.Background(), bill.BillAccountLogRequest{
		StartTime: ptr("2026-11-01 00:00:00"), EndTime: ptr("2026-11-02 00:00:00")})
	if err != nil {
		t.Fatalf("BillAccountLog() error = %v", err)
	}
	if len(resp.Data.DetailList) != 1 || *resp.Data.DetailList[0].AccountLogID != "L1" {
		t.Fatalf("details = %+v", resp.Data.DetailList)
	}
}
