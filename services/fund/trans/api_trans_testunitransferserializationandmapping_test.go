package trans_test

import (
	"context"
	"testing"

	"github.com/tiyee/alipay-sdk-go/services/fund/trans"
)

func TestUniTransferSerializationAndMapping(t *testing.T) {
	svc, executor := newService(t)
	executor.result.Node = []byte(`{"order_id":"202611051100700015","out_biz_no":"T1","status":"SUCCESS","trans_date":"2026-11-05 12:00:00"}`)
	resp, err := svc.UniTransfer(context.Background(), trans.UniTransferRequest{
		OutBizNo: "T1", TransAmount: "10.00", ProductCode: "TRANS_ACCOUNT_NO_PWD",
		OrderTitle: ptr("转账标题"),
		PayeeInfo: &trans.Participant{
			Identity: ptr("2088101117955611"), IdentityType: ptr("ALIPAY_USER_ID"), Name: ptr("张三"),
		}})
	if err != nil {
		t.Fatalf("UniTransfer() error = %v", err)
	}
	body := string(executor.req.Params.JSON)
	for _, want := range []string{
		`"out_biz_no":"T1"`, `"trans_amount":"10.00"`, `"product_code":"TRANS_ACCOUNT_NO_PWD"`,
		`"payee_info":{"identity":"2088101117955611","identity_type":"ALIPAY_USER_ID","name":"张三"}`,
	} {
		if !contains(body, want) {
			t.Fatalf("body %s missing %s", body, want)
		}
	}
	if resp.Data.Status == nil || *resp.Data.Status != trans.TransStatusSuccess {
		t.Fatalf("status = %+v", resp.Data.Status)
	}
}
