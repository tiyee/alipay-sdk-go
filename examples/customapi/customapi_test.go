package customapi_test

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json/jsontext"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strings"
	"testing"

	"github.com/tiyee/go-alipay/core"
	"github.com/tiyee/go-alipay/core/option"
	"github.com/tiyee/go-alipay/examples/customapi"
)

// 外部扩展验收：仅导入公开包即可完成 JSON 与直接参数两类扩展，
// 不修改 SDK、不注册全局接口表、不导入 internal/。

// generateKeyPair 在测试内生成专用密钥对，返回 PKCS#8 私钥与 PKIX 公钥 PEM。
func generateKeyPair(t *testing.T) (privatePEM, publicPEM string) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	privateDER, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	publicDER, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		t.Fatal(err)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privateDER})),
		string(pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: publicDER}))
}

// signData 用网关侧私钥按 RSA2 生成 base64 签名（独立于 SDK 实现）。
func signData(t *testing.T, privateKeyPEM string, data []byte) string {
	t.Helper()
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		t.Fatal("bad gateway key")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		t.Fatal(err)
	}
	rsaKey := key.(*rsa.PrivateKey)
	digest := sha256.Sum256(data)
	sig, err := rsa.SignPKCS1v15(rand.Reader, rsaKey, crypto.SHA256, digest[:])
	if err != nil {
		t.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(sig)
}

// signContent 独立构造待签名串（排除 sign 与空值、按 key 升序）。
func signContent(form url.Values) string {
	keys := make([]string, 0, len(form))
	for key, vs := range form {
		if key == "sign" || len(vs) == 0 || vs[0] == "" {
			continue
		}
		keys = append(keys, key)
	}
	slices.Sort(keys)
	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		parts = append(parts, key+"="+form.Get(key))
	}
	return strings.Join(parts, "&")
}

// newStubGateway 启动本地网关：核对请求签名并返回用户构造的节点响应。
func newStubGateway(t *testing.T, gatewayPrivatePEM, appPublicPEM string,
	respond func(form url.Values) (int, string)) *httptest.Server {
	t.Helper()
	block, _ := pem.Decode([]byte(appPublicPEM))
	if block == nil {
		t.Fatal("bad app public key")
	}
	appKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		t.Fatal(err)
	}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Errorf("parse form: %v", err)
		}
		form := r.PostForm
		content := signContent(form)
		digest := sha256.Sum256([]byte(content))
		sig, err := base64.StdEncoding.DecodeString(form.Get("sign"))
		if err != nil || rsa.VerifyPKCS1v15(appKey.(*rsa.PublicKey), crypto.SHA256, digest[:], sig) != nil {
			t.Errorf("gateway: request signature invalid: %v", err)
		}
		status, body := respond(form)
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
}

// newClient 基于测试密钥构建 Client。
func newClient(t *testing.T, gatewayURL, appPrivatePEM, gatewayPublicPEM string) *core.Client {
	t.Helper()
	client, err := core.NewClient(
		option.WithAppID("2021000000000001"),
		option.WithGateway(gatewayURL),
		option.WithPrivateKey(appPrivatePEM),
		option.WithAlipayPublicKey(gatewayPublicPEM),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	return client
}

func signedBody(t *testing.T, gatewayPrivatePEM, responseKey, node string) string {
	t.Helper()
	return `{"` + responseKey + `":` + node + `,"sign":"` +
		signData(t, gatewayPrivatePEM, []byte(node)) + `"}`
}

func TestCustomJSONAPIEndToEnd(t *testing.T) {
	appPrivate, appPublic := generateKeyPair(t)
	gatewayPrivate, gatewayPublic := generateKeyPair(t)

	node := `{"code":"10000","msg":"Success",` +
		`"bill_download_url":"https://download.alipay.com/bill.zip","bill_file_code":"EMPTY_DATA_WITH_BILL_FILE"}`
	gateway := newStubGateway(t, gatewayPrivate, appPublic, func(form url.Values) (int, string) {
		if form.Get("method") != "alipay.data.dataservice.bill.downloadurl.query" {
			t.Errorf("method = %q", form.Get("method"))
		}
		if form.Get("biz_content") != `{"bill_type":"trade","bill_date":"2026-11-04"}` {
			t.Errorf("biz_content = %q", form.Get("biz_content"))
		}
		return http.StatusOK, signedBody(t, gatewayPrivate,
			"alipay_data_dataservice_bill_downloadurl_query_response", node)
	})
	defer gateway.Close()
	client := newClient(t, gateway.URL, appPrivate, gatewayPublic)

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

func TestCustomDecoderNotCalledOnBadSignature(t *testing.T) {
	appPrivate, appPublic := generateKeyPair(t)
	gatewayPrivate, gatewayPublic := generateKeyPair(t)

	node := `{"code":"10000","msg":"Success","bill_download_url":"https://example.com/x.zip"}`
	// 签名对应原始 node，但返回被篡改的节点。
	tampered := strings.Replace(node, "example.com", "attacker.example", 1)
	gateway := newStubGateway(t, gatewayPrivate, appPublic, func(url.Values) (int, string) {
		return http.StatusOK, `{"alipay_data_dataservice_bill_downloadurl_query_response":` +
			tampered + `,"sign":"` + signData(t, gatewayPrivate, []byte(node)) + `"}`
	})
	defer gateway.Close()
	client := newClient(t, gateway.URL, appPrivate, gatewayPublic)

	decoder := &countingDecoder{inner: core.JSONDecoder[customapi.BillResponse]()}
	op, err := core.NewOperation[customapi.BillRequest, customapi.BillResponse](
		core.OperationSpec{
			Method:        "alipay.data.dataservice.bill.downloadurl.query",
			Version:       "1.0",
			ResponseKey:   "alipay_data_dataservice_bill_downloadurl_query_response",
			ParameterMode: core.ParameterModeBizJSON,
		}, core.JSONEncoder[customapi.BillRequest](), decoder)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Execute(context.Background(), op,
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

func (d *countingDecoder) Decode(payload jsontext.Value) (customapi.BillResponse, error) {
	d.calls++
	return d.inner.Decode(payload)
}

func TestCustomFormAPIAndProtectedParams(t *testing.T) {
	appPrivate, appPublic := generateKeyPair(t)
	gatewayPrivate, gatewayPublic := generateKeyPair(t)

	// custom.demo.form.api 是机制演示占位方法名，不是真实支付宝接口。
	const demoResponseKey = "custom_demo_form_api_response"
	var gotForm url.Values
	gateway := newStubGateway(t, gatewayPrivate, appPublic, func(form url.Values) (int, string) {
		gotForm = form
		return http.StatusOK, signedBody(t, gatewayPrivate, demoResponseKey,
			`{"code":"10000","msg":"Success","trade_no":"T123"}`)
	})
	defer gateway.Close()
	client := newClient(t, gateway.URL, appPrivate, gatewayPublic)

	op, err := core.NewOperation[customapi.FormRequest, customapi.FormResponse](
		core.OperationSpec{
			Method:        "custom.demo.form.api",
			Version:       "1.0",
			ResponseKey:   demoResponseKey,
			ParameterMode: core.ParameterModeForm,
		}, customapi.FormEncoder{}, core.JSONDecoder[customapi.FormResponse]())
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Execute(context.Background(), op, core.Request[customapi.FormRequest]{
		Body: customapi.FormRequest{OutTradeNo: "T001", Extra: "中文&值"}})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if resp.Data.TradeNo != "T123" {
		t.Fatalf("Data = %+v", resp.Data)
	}
	if gotForm.Get("out_trade_no") != "T001" {
		t.Errorf("out_trade_no = %q", gotForm.Get("out_trade_no"))
	}
	if got := gotForm.Get("extra"); got != "中文&值" {
		t.Errorf("extra = %q, want url-decoded original value", got)
	}
	if gotForm.Get("biz_content") != "" {
		t.Error("form mode must not wrap params into biz_content")
	}

	// 受保护参数覆盖被拒绝。
	_, err = client.Execute(context.Background(), op, core.Request[customapi.FormRequest]{
		Body: customapi.FormRequest{OutTradeNo: "T002", Extra: "x"}})
	// 先通过合法编码确认机制，再演示覆盖：
	coverOp, err2 := core.NewOperation[customapi.FormRequest, customapi.FormResponse](
		core.OperationSpec{
			Method:        "custom.demo.form.api",
			Version:       "1.0",
			ResponseKey:   demoResponseKey,
			ParameterMode: core.ParameterModeForm,
		}, coverEncoder{}, core.JSONDecoder[customapi.FormResponse]())
	if err2 != nil {
		t.Fatal(err2)
	}
	_, err = client.Execute(context.Background(), coverOp, core.Request[customapi.FormRequest]{
		Body: customapi.FormRequest{OutTradeNo: "T003"}})
	var valErr *core.ValidationError
	if !errors.As(err, &valErr) || valErr.Field != "app_id" {
		t.Fatalf("error = %v, want protected param rejection", err)
	}
}

// coverEncoder 故意生成包含受保护参数的表单，验证 Client 的拦截。
type coverEncoder struct{}

func (coverEncoder) Encode(customapi.FormRequest) (core.EncodedParams, error) {
	return core.EncodedParams{Mode: core.ParameterModeForm,
		Form: url.Values{"app_id": {"evil"}}}, nil
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
