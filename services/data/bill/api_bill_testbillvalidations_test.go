package bill_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/data/bill"
)

func TestBillValidations(t *testing.T) {
	svc, executor := newService(t)
	if _, err := svc.BillDownloadURLQuery(context.Background(), bill.BillDownloadURLRequest{
		BillType: "trade"}); err == nil {
		t.Fatal("bill_date is required")
	}
	if _, err := svc.BillBalanceHis(context.Background(), bill.BillBalanceHisRequest{
		BizDate: ptr("2026-11-04"), BizMonth: ptr("2026-11")}); err == nil {
		t.Fatal("biz_date and biz_month are mutually exclusive")
	}
	if _, err := svc.BillAccountLog(context.Background(), bill.BillAccountLogRequest{}); err == nil {
		t.Fatal("a query condition is required")
	}
	if _, err := svc.BillAccountLog(context.Background(), bill.BillAccountLogRequest{
		AlipayOrderNo: ptr("A1"), MerchantOrderNo: ptr("M1")}); err == nil {
		t.Fatal("order numbers are mutually exclusive")
	}
	if executor.called {
		t.Error("validation failures must not send request")
	}
	if _, err := svc.BillAccountLog(context.Background(), bill.BillAccountLogRequest{
		MerchantOrderNo: ptr("M1")}); err != nil {
		t.Fatalf("exact order query should pass: %v", err)
	}
}
