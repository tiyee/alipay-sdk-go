// Package customapi_test 验收外部扩展：customapi 包仅导入公开包
// （core、option、utils）即可定义 V3 自定义 Operation、编解码器与
// Service，不修改 SDK、不注册全局接口表、扩展代码不导入 internal/。
// 测试自身使用 internal/testutil 的专用虚构密钥与协议断言辅助，
// 通过 httptest 模拟 V3 网关，不访问支付宝。
package customapi_test

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/examples/customapi"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// newStubGateway 启动本地 V3 网关：独立核对请求签名，返回带
// alipay-signature 响应签名头的 JSON。
func newStubGateway(t *testing.T,
	respond func(r *http.Request, body []byte) (int, string)) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := testutil.ReadAllBody(t, r)
		testutil.VerifyV3Authorization(t, r, body, "")
		status, respBody := respond(r, body)
		for key, value := range testutil.V3ResponseHeaders(t, []byte(respBody),
			"1764000000300", "customapi-nonce", "") {
			w.Header()[key] = value
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(respBody))
	}))
}

// newClient 基于测试专用虚构密钥构建指向模拟网关的 Client。
func newClient(t *testing.T, gatewayURL string) *core.Client {
	t.Helper()
	credentials, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner:       testutil.NewAppCryptoSigner(t),
		AlipayPublicKey: mustAlipayPublicKey(t),
	})
	if err != nil {
		t.Fatalf("build credentials: %v", err)
	}
	client, err := core.NewClient(
		credentials,
		option.WithAppID("2021000000000001"),
		option.WithGateway(gatewayURL),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	return client
}

// mustAlipayPublicKey 解析测试专用的支付宝公钥。
func mustAlipayPublicKey(t *testing.T) *rsa.PublicKey {
	t.Helper()
	key, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(testutil.AlipayPublicKeyPEM))
	if err != nil {
		t.Fatalf("parse alipay public key: %v", err)
	}
	return key
}

func TestCustomQueryAPIEndToEnd(t *testing.T) {
	node := `{"bill_download_url":"https://download.alipay.com/bill.zip","bill_file_code":"EMPTY_DATA_WITH_BILL_FILE"}`
	gateway := newStubGateway(t, func(r *http.Request, _ []byte) (int, string) {
		if r.URL.Path != "/v3/alipay/data/dataservice/bill/downloadurl/query" {
			t.Errorf("path = %q", r.URL.Path)
		}
		// V3 GET 参数在 query；默认 QueryEncoder 按 query 标签输出。
		if got := r.URL.Query().Get("bill_type"); got != "trade" {
			t.Errorf("bill_type = %q", got)
		}
		if got := r.URL.Query().Get("bill_date"); got != "2026-11-04" {
			t.Errorf("bill_date = %q", got)
		}
		if got := r.URL.Query().Get("smid"); got != "" {
			t.Errorf("smid = %q, want omitted", got)
		}
		return http.StatusOK, node
	})
	defer gateway.Close()
	client := newClient(t, gateway.URL)

	// 通过用户自定义 Service 调用。
	svc, err := customapi.NewBillService(client)
	if err != nil {
		t.Fatalf("NewBillService() error = %v", err)
	}
	resp, err := svc.Query(context.Background(), customapi.BillRequest{
		BillType: customapi.BillTypeTrade, BillDate: "2026-11-04"})
	if err != nil {
		t.Fatalf("Query() error = %v", err)
	}
	if resp.Data.BillDownloadURL == nil || *resp.Data.BillDownloadURL != "https://download.alipay.com/bill.zip" {
		t.Fatalf("Data = %+v", resp.Data)
	}
	if resp.Data.BillFileCode == nil || *resp.Data.BillFileCode != customapi.BillFileCodeEmptyData {
		t.Fatalf("BillFileCode = %+v", resp.Data.BillFileCode)
	}

	// 通过 Client 泛型方法直接调用（同一 Operation 契约）。
	resp2, err := customapi.QueryBill(context.Background(), client,
		customapi.BillRequest{BillType: customapi.BillTypeTrade, BillDate: "2026-11-04"})
	if err != nil || resp2.Data.BillDownloadURL == nil {
		t.Fatalf("QueryBill() error = %v", err)
	}
}

// TestCustomDecoderNotCalledOnBadSignature 验证验签失败时解码器不接触
// 响应数据：篡改报文不能作为可信业务结果。
func TestCustomDecoderNotCalledOnBadSignature(t *testing.T) {
	node := `{"bill_download_url":"https://example.com/x.zip"}`
	tampered := strings.Replace(node, "example.com", "attacker.example", 1)
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// 响应签名对应原始 node，返回的却是被篡改的报文。
		for key, value := range testutil.V3ResponseHeaders(t, []byte(node),
			"1764000000300", "tamper-nonce", "") {
			w.Header()[key] = value
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(tampered))
	}))
	defer gateway.Close()
	client := newClient(t, gateway.URL)

	decoder := &countingDecoder{inner: core.JSONDecoder[customapi.BillResponse]()}
	op, err := core.NewOperation[customapi.BillRequest, customapi.BillResponse](
		core.OperationSpec{
			Method:          "alipay.data.dataservice.bill.downloadurl.query",
			HTTPMethod:      core.HTTPMethodGet,
			Path:            "/v3/alipay/data/dataservice/bill/downloadurl/query",
			ParameterMode:   core.ParameterModeQuery,
			AuthRequirement: core.AuthRequirementAny,
		}, core.QueryEncoder[customapi.BillRequest](), decoder)
	if err != nil {
		t.Fatal(err)
	}
	_, err = core.Execute(context.Background(), client, op,
		core.Request[customapi.BillRequest]{Body: customapi.BillRequest{
			BillType: customapi.BillTypeTrade, BillDate: "2026-11-04"}})
	var sigErr *core.SignatureError
	if !errors.As(err, &sigErr) {
		t.Fatalf("error = %v, want SignatureError", err)
	}
	if decoder.calls != 0 {
		t.Fatalf("custom decoder called %d times on bad signature; must not be called", decoder.calls)
	}
}

type countingDecoder struct {
	inner core.Decoder[customapi.BillResponse]
	calls int
}

func (d *countingDecoder) Decode(payload json.RawMessage) (customapi.BillResponse, error) {
	d.calls++
	return d.inner.Decode(payload)
}

// TestCustomEncoderOperation 验证自定义 query 编码器的接入与必填校验。
func TestCustomEncoderOperation(t *testing.T) {
	var gotQuery url.Values
	gateway := newStubGateway(t, func(r *http.Request, _ []byte) (int, string) {
		if r.URL.Path != "/v3/custom/demo/query" {
			t.Errorf("path = %q", r.URL.Path)
		}
		gotQuery = r.URL.Query()
		return http.StatusOK, `{"trade_no":"T123"}`
	})
	defer gateway.Close()
	client := newClient(t, gateway.URL)

	op, err := customapi.NewDirectParamsOperation()
	if err != nil {
		t.Fatal(err)
	}
	resp, err := core.Execute(context.Background(), client, op,
		core.Request[customapi.DirectParamsRequest]{
			Body: customapi.DirectParamsRequest{OutTradeNo: "T001", Extra: "中文&值"}})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if resp.Data.TradeNo != "T123" {
		t.Fatalf("Data = %+v", resp.Data)
	}
	if gotQuery.Get("out_trade_no") != "T001" {
		t.Errorf("out_trade_no = %q", gotQuery.Get("out_trade_no"))
	}
	if got := gotQuery.Get("extra"); got != "中文&值" {
		t.Errorf("extra = %q, want url-decoded original value", got)
	}
	if got := gotQuery.Get("app_id"); got != "" {
		t.Errorf("app_id = %q, want absent from query; public params live in the Authorization header", got)
	}

	// 编码器必填校验：缺 out_trade_no 时请求不发送。
	_, err = core.Execute(context.Background(), client, op,
		core.Request[customapi.DirectParamsRequest]{})
	var valErr *core.ValidationError
	if !errors.As(err, &valErr) || valErr.Field != "out_trade_no" {
		t.Fatalf("error = %v, want out_trade_no validation error", err)
	}
}

// ExampleNewBillService 演示用户包中自定义 Service 的构造与调用。
func ExampleNewBillService() {
	// executor 通常为 *core.Client；这里仅示意构造流程。
	var executor core.Executor
	svc, err := customapi.NewBillService(executor)
	if err == nil {
		_ = svc
	}
	fmt.Println("executor is required:", err != nil)
	// Output:
	// executor is required: true
}
