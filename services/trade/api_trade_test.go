package trade_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tiyee/go-alipay/core"
	"github.com/tiyee/go-alipay/core/option"
	"github.com/tiyee/go-alipay/internal/testutil"
	"github.com/tiyee/go-alipay/services/trade"
)

// recordingExecutor 记录 WireRequest，返回预置结果，
// 用于验证接口名、参数位置与响应映射（不发起网络请求）。
type recordingExecutor struct {
	req    core.WireRequest
	called bool

	result *core.APIResult
	err    error
}

func (e *recordingExecutor) Do(_ context.Context, req core.WireRequest) (*core.APIResult, error) {
	e.called = true
	e.req = req
	return e.result, e.err
}

func verifiedResult(node string) *core.APIResult {
	return &core.APIResult{
		HTTPStatus:        200,
		SignatureVerified: true,
		Node:              []byte(node),
	}
}

func errorResult(err error) *core.APIResult { return &core.APIResult{HTTPStatus: 200} }

func newRecordingService(t *testing.T) (*trade.Service, *recordingExecutor) {
	t.Helper()
	executor := &recordingExecutor{}
	svc, err := trade.NewService(executor)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	return svc, executor
}

func TestNewServiceRejectsNilExecutor(t *testing.T) {
	if _, err := trade.NewService(nil); err == nil {
		t.Fatal("NewService(nil) must fail")
	}
}

func TestOperationSpecs(t *testing.T) {
	tests := []struct {
		name        string
		call        func(*trade.Service, *recordingExecutor) error
		wantMethod  string
		wantRespKey string
	}{
		{
			name: "查询",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"})
				return err
			},
			wantMethod:  "alipay.trade.query",
			wantRespKey: "alipay_trade_query_response",
		},
		{
			name: "预下单",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.Precreate(context.Background(), trade.PrecreateRequest{
					OutTradeNo: "T1", TotalAmount: "88.88", Subject: "测试商品"})
				return err
			},
			wantMethod:  "alipay.trade.precreate",
			wantRespKey: "alipay_trade_precreate_response",
		},
		{
			name: "退款",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.Refund(context.Background(), trade.RefundRequest{
					OutTradeNo: "T1", RefundAmount: "10.00"})
				return err
			},
			wantMethod:  "alipay.trade.refund",
			wantRespKey: "alipay_trade_refund_response",
		},
		{
			name: "退款查询",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.QueryRefund(context.Background(), trade.QueryRefundRequest{
					OutRequestNo: "R1", TradeNo: "T1"})
				return err
			},
			wantMethod:  "alipay.trade.fastpay.refund.query",
			wantRespKey: "alipay_trade_fastpay_refund_query_response",
		},
		{
			name: "关单",
			call: func(s *trade.Service, e *recordingExecutor) error {
				_, err := s.Close(context.Background(), trade.CloseRequest{TradeNo: "T1"})
				return err
			},
			wantMethod:  "alipay.trade.close",
			wantRespKey: "alipay_trade_close_response",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, executor := newRecordingService(t)
			executor.result = verifiedResult(`{"code":"10000","msg":"Success"}`)
			if err := tt.call(svc, executor); err != nil {
				t.Fatalf("call error = %v", err)
			}
			spec := executor.req.Spec
			if spec.Method != tt.wantMethod {
				t.Errorf("Method = %q, want %q", spec.Method, tt.wantMethod)
			}
			if spec.ResponseKey != tt.wantRespKey {
				t.Errorf("ResponseKey = %q, want %q", spec.ResponseKey, tt.wantRespKey)
			}
			if spec.Version != "1.0" {
				t.Errorf("Version = %q", spec.Version)
			}
			if spec.ParameterMode != core.ParameterModeBizJSON {
				t.Errorf("ParameterMode = %v", spec.ParameterMode)
			}
			if len(spec.AcceptedCodes) != 1 || spec.AcceptedCodes[0] != core.ResponseCodeSuccess {
				t.Errorf("AcceptedCodes = %v", spec.AcceptedCodes)
			}
			if executor.req.Params.Mode != core.ParameterModeBizJSON {
				t.Errorf("params mode = %v", executor.req.Params.Mode)
			}
		})
	}
}

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

func TestQueryResponseMapping(t *testing.T) {
	// 响应示例依据查询文档正常示例。
	node := `{"code":"10000","msg":"Success",` +
		`"trade_no":"2013112011001004330000121536","out_trade_no":"6823789339978248",` +
		`"buyer_logon_id":"159****5620","trade_status":"TRADE_CLOSED",` +
		`"total_amount":"88.88","send_pay_date":"2014-11-27 15:45:57",` +
		`"buyer_pay_amount":"8.88","point_amount":"10","receipt_amount":"15.25",` +
		`"store_id":"NJ_S_001","terminal_id":"NJ_T_001",` +
		`"fund_bill_list":[{"fund_channel":"ALIPAYACCOUNT","amount":"10","real_amount":"11.21"}],` +
		`"store_name":"证大五道口店","buyer_user_id":"2088101117955611",` +
		`"buyer_user_type":"PRIVATE","mdiscount_amount":"88.88"}`
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(node)
	resp, err := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"})
	if err != nil {
		t.Fatalf("Query() error = %v", err)
	}
	data := resp.Data
	if data.TradeStatus != trade.TradeStatusClosed {
		t.Errorf("TradeStatus = %q", data.TradeStatus)
	}
	if data.TotalAmount != "88.88" || data.ReceiptAmount == nil || *data.ReceiptAmount != "15.25" {
		t.Errorf("amounts = %+v", data)
	}
	if len(data.FundBillList) != 1 || data.FundBillList[0].FundChannel != "ALIPAYACCOUNT" {
		t.Errorf("FundBillList = %+v", data.FundBillList)
	}
	if data.BuyerUserType == nil || *data.BuyerUserType != trade.BuyerUserTypePrivate {
		t.Errorf("BuyerUserType = %+v", data.BuyerUserType)
	}
	if data.StoreName == nil || *data.StoreName != "证大五道口店" {
		t.Errorf("StoreName = %+v", data.StoreName)
	}
	if data.TradeStatus.IsKnown() != true {
		t.Error("known status expected")
	}
}

func TestQueryResponseUnknownEnumPreserved(t *testing.T) {
	node := `{"code":"10000","msg":"Success","trade_status":"FUTURE_NEW_STATUS","total_amount":"1.00"}`
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(node)
	resp, err := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"})
	if err != nil {
		t.Fatalf("Query() error = %v", err)
	}
	if resp.Data.TradeStatus != trade.TradeStatus("FUTURE_NEW_STATUS") {
		t.Fatalf("unknown status = %q, want preserved", resp.Data.TradeStatus)
	}
	if resp.Data.TradeStatus.IsKnown() {
		t.Error("unknown status must not be reported as known")
	}
}

func TestPrecreateRequestAndResponse(t *testing.T) {
	qr := "https://qr.alipay.com/bavh4wjlxf12tper3a"
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(fmt.Sprintf(
		`{"code":"10000","msg":"Success","out_trade_no":"20150320010101001","qr_code":%q}`, qr))
	resp, err := svc.Precreate(context.Background(), trade.PrecreateRequest{
		OutTradeNo:  "20150320010101001",
		TotalAmount: "88.88",
		Subject:     "Iphone6 16G",
		GoodsDetail: []trade.GoodsDetail{{
			GoodsID: "apple-01", GoodsName: "ipad", Quantity: 1, Price: "2000",
		}},
	})
	if err != nil {
		t.Fatalf("Precreate() error = %v", err)
	}
	if resp.Data.QRCode == nil || *resp.Data.QRCode != qr {
		t.Fatalf("QRCode = %+v", resp.Data.QRCode)
	}
	if resp.Data.OutTradeNo != "20150320010101001" {
		t.Fatalf("OutTradeNo = %q", resp.Data.OutTradeNo)
	}
	var biz map[string]json.RawMessage
	if err := json.Unmarshal(executor.req.Params.JSON, &biz); err != nil {
		t.Fatal(err)
	}
	if got := string(biz["total_amount"]); got != `"88.88"` {
		t.Errorf("total_amount = %s, want exact decimal string", got)
	}
	if got := string(biz["subject"]); got != `"Iphone6 16G"` {
		t.Errorf("subject = %s", got)
	}
	goods := `[{"goods_id":"apple-01","goods_name":"ipad","quantity":1,"price":"2000"}]`
	if got := string(biz["goods_detail"]); got != goods {
		t.Errorf("goods_detail = %s, want %s", got, goods)
	}
	if _, ok := biz["product_code"]; ok {
		t.Error("unset optional fields must be omitted")
	}
}

func TestRefundResponseFundChangeSemantics(t *testing.T) {
	// code=10000 仅代表请求成功；SDK 不把 fund_change=N 判定为失败，
	// 语义交由调用方依据文档判断。
	for _, tt := range []struct {
		fundChange string
		want       trade.FundChange
	}{
		{fundChange: "Y", want: trade.FundChangeYes},
		{fundChange: "N", want: trade.FundChangeNo},
	} {
		node := fmt.Sprintf(`{"code":"10000","msg":"Success","trade_no":"T1","refund_fee":"12.33","fund_change":%q}`, tt.fundChange)
		svc, executor := newRecordingService(t)
		executor.result = verifiedResult(node)
		resp, err := svc.Refund(context.Background(), trade.RefundRequest{
			OutTradeNo: "T1", RefundAmount: "12.33", OutRequestNo: ptr("R1")})
		if err != nil {
			t.Fatalf("Refund() error = %v", err)
		}
		if resp.Data.FundChange == nil || *resp.Data.FundChange != tt.want {
			t.Fatalf("FundChange = %+v, want %q", resp.Data.FundChange, tt.want)
		}
		if executor.req.Params.Mode != core.ParameterModeBizJSON {
			t.Fatal("refund must use biz_content")
		}
	}
}

func TestQueryRefundResponseMapping(t *testing.T) {
	node := `{"code":"10000","msg":"Success","trade_no":"T1","out_trade_no":"O1",` +
		`"out_request_no":"R1","refund_status":"REFUND_SUCCESS",` +
		`"refund_amount":"12.33","total_amount":"100.20"}`
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(node)
	resp, err := svc.QueryRefund(context.Background(), trade.QueryRefundRequest{
		OutRequestNo: "R1", TradeNo: "T1"})
	if err != nil {
		t.Fatalf("QueryRefund() error = %v", err)
	}
	if resp.Data.RefundStatus == nil || *resp.Data.RefundStatus != trade.RefundStatusSuccess {
		t.Fatalf("RefundStatus = %+v", resp.Data.RefundStatus)
	}
}

func TestCloseResponseMapping(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(
		`{"code":"10000","msg":"Success","trade_no":"T1","out_trade_no":"O1"}`)
	resp, err := svc.Close(context.Background(), trade.CloseRequest{TradeNo: "T1"})
	if err != nil {
		t.Fatalf("Close() error = %v", err)
	}
	if resp.Data.TradeNo == nil || *resp.Data.TradeNo != "T1" {
		t.Fatalf("TradeNo = %+v", resp.Data.TradeNo)
	}
}

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
				_, err := s.Precreate(context.Background(), trade.PrecreateRequest{
					OutTradeNo: "T1", Subject: "标题"})
				return err
			},
		},
		{
			name: "预下单标题含特殊字符",
			call: func(s *trade.Service) error {
				_, err := s.Precreate(context.Background(), trade.PrecreateRequest{
					OutTradeNo: "T1", TotalAmount: "1.00", Subject: "a/b"})
				return err
			},
		},
		{
			name: "预下单订单号含非法字符",
			call: func(s *trade.Service) error {
				_, err := s.Precreate(context.Background(), trade.PrecreateRequest{
					OutTradeNo: "T-1", TotalAmount: "1.00", Subject: "标题"})
				return err
			},
		},
		{
			name: "退款缺少金额",
			call: func(s *trade.Service) error {
				_, err := s.Refund(context.Background(), trade.RefundRequest{TradeNo: "T1"})
				return err
			},
		},
		{
			name: "退款查询缺少请求号",
			call: func(s *trade.Service) error {
				_, err := s.QueryRefund(context.Background(), trade.QueryRefundRequest{TradeNo: "T1"})
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
		{
			name: "金额超过两位小数",
			call: func(s *trade.Service) error {
				_, err := s.Refund(context.Background(), trade.RefundRequest{
					TradeNo: "T1", RefundAmount: "1.234"})
				return err
			},
		},
		{
			name: "金额为零",
			call: func(s *trade.Service) error {
				_, err := s.Refund(context.Background(), trade.RefundRequest{
					TradeNo: "T1", RefundAmount: "0.00"})
				return err
			},
		},
		{
			name: "金额为负",
			call: func(s *trade.Service) error {
				_, err := s.Refund(context.Background(), trade.RefundRequest{
					TradeNo: "T1", RefundAmount: "-1.00"})
				return err
			},
		},
		{
			name: "金额非数字",
			call: func(s *trade.Service) error {
				_, err := s.Refund(context.Background(), trade.RefundRequest{
					TradeNo: "T1", RefundAmount: "12.3a"})
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

func TestAmountSerializationBoundaries(t *testing.T) {
	tests := []string{"0.01", "88.88", "100", "100000000.00", "12.3"}
	for _, amount := range tests {
		t.Run(amount, func(t *testing.T) {
			svc, executor := newRecordingService(t)
			executor.result = verifiedResult(`{"code":"10000","msg":"Success"}`)
			_, err := svc.Precreate(context.Background(), trade.PrecreateRequest{
				OutTradeNo: "T1", TotalAmount: amount, Subject: "标题"})
			if err != nil {
				t.Fatalf("Precreate() error = %v", err)
			}
			var biz struct {
				TotalAmount string `json:"total_amount"`
			}
			if err := json.Unmarshal(executor.req.Params.JSON, &biz); err != nil {
				t.Fatal(err)
			}
			if biz.TotalAmount != amount {
				t.Fatalf("total_amount = %q, want %q (exact decimal passthrough)", biz.TotalAmount, amount)
			}
		})
	}
}

func TestServiceEndToEndWithRealClient(t *testing.T) {
	// 通过真实 Client 与本地网关验证完整链路：签名请求、验签响应、
	// 业务错误分类。业务错误分类发生在 Client 执行层。
	node := `{"code":"40004","msg":"Business Failed","sub_code":"ACQ.TRADE_NOT_EXIST","sub_msg":"交易不存在"}`
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Errorf("parse form: %v", err)
		}
		if r.PostForm.Get("method") != "alipay.trade.query" {
			t.Errorf("method = %q", r.PostForm.Get("method"))
		}
		if err := testutil.VerifyWithAppKey([]byte(testutil.SignForm(r.PostForm)), r.PostForm.Get("sign")); err != nil {
			t.Errorf("request signature invalid: %v", err)
		}
		sign, err := testutil.SignWithAlipayKey([]byte(node))
		if err != nil {
			t.Fatal(err)
		}
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		_, _ = w.Write([]byte(`{"alipay_trade_query_response":` + node + `,"sign":"` + sign + `"}`))
	}))
	defer gateway.Close()

	client, err := core.NewClient(
		option.WithAppID("2021000000000001"),
		option.WithGateway(gateway.URL),
		option.WithPrivateKey(testutil.AppPrivateKeyPEM),
		option.WithAlipayPublicKey(testutil.AlipayPublicKeyPEM),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	svc, err := trade.NewService(client)
	if err != nil {
		t.Fatalf("NewService() error = %v", err)
	}
	resp, err := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"})
	if err == nil {
		t.Fatal("expect APIError")
	}
	if resp == nil || resp.Data != nil {
		t.Fatal("no trusted data on business failure")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) || apiErr.SubCode != "ACQ.TRADE_NOT_EXIST" || !apiErr.Verified {
		t.Fatalf("error = %v", err)
	}
}

func TestServiceCallOptions(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(`{"code":"10000","msg":"Success"}`)
	_, err := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"},
		core.WithAppAuthToken("auth-token-1"), core.WithNotifyURL("https://notify.example/1"))
	if err != nil {
		t.Fatalf("Query() error = %v", err)
	}
	if executor.req.Settings.AppAuthToken != "auth-token-1" {
		t.Errorf("AppAuthToken = %q", executor.req.Settings.AppAuthToken)
	}
	if executor.req.Settings.NotifyURL != "https://notify.example/1" {
		t.Errorf("NotifyURL = %q", executor.req.Settings.NotifyURL)
	}
}

func TestServiceDecodeFailureIsProtocolError(t *testing.T) {
	svc, executor := newRecordingService(t)
	executor.result = verifiedResult(`{"code":"10000","total_amount":1.00}`)
	_, err := svc.Query(context.Background(), trade.QueryRequest{TradeNo: "T1"})
	if !errors.As(err, new(*core.ProtocolError)) {
		t.Fatalf("error = %v, want ProtocolError", err)
	}
	if !strings.Contains(fmt.Sprint(err), "decode") {
		t.Errorf("error = %v", err)
	}
}

func ptr[T any](v T) *T { return &v }

var _ = errorResult
