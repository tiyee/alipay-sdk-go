package bill_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/data/bill"
)

func TestBillDownloadURLMapping(t *testing.T) {
	svc, executor := newService(t)
	executor.result.Node = []byte(`{"bill_download_url":"https://download.alipay.com/bill.zip","bill_file_code":"EMPTY_DATA_WITH_BILL_FILE"}`)
	resp, err := svc.BillDownloadURLQuery(context.Background(), bill.BillDownloadURLRequest{
		BillType: "trade", BillDate: "2026-11-04"})
	if err != nil {
		t.Fatalf("BillDownloadURLQuery() error = %v", err)
	}
	if resp.Data.BillDownloadURL == nil || *resp.Data.BillDownloadURL != "https://download.alipay.com/bill.zip" {
		t.Fatalf("download url missing: %+v", resp.Data)
	}
	if resp.Data.BillFileCode == nil || *resp.Data.BillFileCode != bill.BillFileCodeEmptyDataWithFile {
		t.Fatalf("file code = %+v", resp.Data.BillFileCode)
	}
}
