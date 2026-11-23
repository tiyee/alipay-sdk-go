package sandbox

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/examples/requestfile"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/services/trade"
)

func TestSelected(t *testing.T) {
	for _, tt := range []struct {
		name    string
		enabled string
		api     string
		env     string
		want    bool
		wantErr bool
	}{
		{name: "default is offline"},
		{name: "explicitly disabled", enabled: "0", api: "alipay.trade.query"},
		{name: "invalid toggle", enabled: "true", wantErr: true},
		{name: "requires API selection", enabled: "1", wantErr: true},
		{name: "different API skipped", enabled: "1", api: "alipay.trade.refund"},
		{name: "selected", enabled: "1", api: "alipay.trade.query", want: true},
		{name: "explicit sandbox", enabled: "1", api: "alipay.trade.query", env: "sandbox", want: true},
		{name: "production rejected", enabled: "1", api: "alipay.trade.query", env: "production", wantErr: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			values := map[string]string{"ALIPAY_SANDBOX": tt.enabled, "ALIPAY_SANDBOX_API": tt.api, "ALIPAY_ENV": tt.env}
			got, err := selected(func(name string) string { return values[name] }, "alipay.trade.query")
			if got != tt.want || (err != nil) != tt.wantErr {
				t.Fatalf("selected() = %t, %v; want enabled=%t, error=%t", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestRunDisabledDoesNotCreateService(t *testing.T) {
	t.Setenv("ALIPAY_SANDBOX", "")
	called := false
	t.Run("offline", func(t *testing.T) {
		Run(t, "alipay.trade.query", func(core.Executor) (*trade.Service, error) {
			called = true
			return nil, fmt.Errorf("must not be called")
		}, (*trade.Service).Query)
	})
	if called {
		t.Fatal("disabled sandbox attempted to create Service")
	}
}

func TestDecodeRequest(t *testing.T) {
	for _, tt := range []struct {
		name string
		body string
		bad  bool
	}{
		{name: "valid", body: `{"out_trade_no":"SANDBOX-ORDER"}`},
		{name: "unknown field", body: `{"out_trade_no":"SANDBOX-ORDER","typo":1}`, bad: true},
		{name: "wrong type", body: `{"out_trade_no":1}`, bad: true},
		{name: "null", body: `null`, bad: true},
		{name: "array", body: `[]`, bad: true},
		{name: "duplicate", body: `{"out_trade_no":"a","out_trade_no":"b"}`, bad: true},
		{name: "trailing data", body: `{} {}`, bad: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			req, err := requestfile.Decode[trade.QueryRequest]([]byte(tt.body))
			if (err != nil) != tt.bad {
				t.Fatalf("decodeRequest() error = %v; want error=%t", err, tt.bad)
			}
			if !tt.bad && req.OutTradeNo != "SANDBOX-ORDER" {
				t.Fatalf("OutTradeNo = %q", req.OutTradeNo)
			}
		})
	}
	type mixedRequest struct {
		ID    string  `json:"-" path:"id"`
		Token *string `json:"-" query:"auth_token"`
		Title string  `json:"title"`
	}
	mixed, err := requestfile.Decode[mixedRequest]([]byte(`{"id":"a/b","auth_token":"offline","title":"账单"}`))
	if err != nil || mixed.ID != "a/b" || mixed.Token == nil || *mixed.Token != "offline" || mixed.Title != "账单" {
		t.Fatalf("mixed request = %+v, %v", mixed, err)
	}
	if _, err := requestfile.Decode[mixedRequest]([]byte(`{"id":"a","id":"b"}`)); err == nil {
		t.Fatal("duplicate path parameter accepted")
	}
	// 没有 json 标签的 query/path 字段按 Go 字段名读取；multipart 文件
	// 按请求模型中的 Content []byte 接受 JSON base64 字符串。
	type queryRequest struct {
		PageSize int64 `query:"page_size"`
	}
	if req, err := requestfile.Decode[queryRequest]([]byte(`{"PageSize":10}`)); err != nil || req.PageSize != 10 {
		t.Fatalf("query request = %+v, %v", req, err)
	}
	if req, err := requestfile.Decode[core.MultipartFile]([]byte(`{"Filename":"test.txt","Content":"aGVsbG8="}`)); err != nil || string(req.Content) != "hello" {
		t.Fatalf("multipart request error = %v", err)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestCertificateClient(t *testing.T) {
	suite := testutil.NewCertSuite(t)
	dir := t.TempDir()
	write := func(name, content string) string {
		t.Helper()
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(content), 0600); err != nil {
			t.Fatal(err)
		}
		return path
	}
	values := map[string]string{
		"ALIPAY_APP_ID":          testutil.TestAppID,
		"ALIPAY_APP_PRIVATE_KEY": write("private.pem", testutil.AppPrivateKeyPEM),
		"ALIPAY_APP_CERT":        write("app.crt", suite.AppCertPEM),
		"ALIPAY_ROOT_CERT":       write("root.crt", suite.RootPEM),
		"ALIPAY_CERT":            write("alipay.crt", suite.AlipayCertPEM),
	}
	secondCert, _ := suite.NewSecondAlipayCert(t)
	secondPEM := testutil.CertToPEM(t, secondCert)
	secondPath := write("alipay-2.crt", secondPEM)
	bundlePath := write("alipay-bundle.crt", suite.AlipayCertPEM+secondPEM)
	wrongKey := write("wrong-private.pem", testutil.AlipayPrivateKeyPEM)
	untrustedRoot := write("untrusted-root.crt", suite.SecondRootPEM)
	for _, tt := range []struct {
		name    string
		field   string
		value   string
		wantErr bool
	}{
		{name: "single certificate"},
		{name: "certificate bundle", field: "ALIPAY_CERT", value: bundlePath},
		{name: "multiple files", field: "ALIPAY_CERT", value: values["ALIPAY_CERT"] + ", " + secondPath},
		{name: "missing APPID", field: "ALIPAY_APP_ID", wantErr: true},
		{name: "missing cert", field: "ALIPAY_CERT", wantErr: true},
		{name: "missing app cert", field: "ALIPAY_APP_CERT", wantErr: true},
		{name: "empty cert path", field: "ALIPAY_CERT", value: values["ALIPAY_CERT"] + ",", wantErr: true},
		{name: "unreadable cert", field: "ALIPAY_CERT", value: filepath.Join(dir, "missing"), wantErr: true},
		{name: "invalid cert", field: "ALIPAY_CERT", value: values["ALIPAY_APP_PRIVATE_KEY"], wantErr: true},
		{name: "mismatched private key", field: "ALIPAY_APP_PRIVATE_KEY", value: wrongKey, wantErr: true},
		{name: "untrusted certificate", field: "ALIPAY_ROOT_CERT", value: untrustedRoot, wantErr: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			getenv := func(name string) string {
				if name == tt.field {
					return tt.value
				}
				return values[name]
			}
			for _, tamper := range []bool{false, true} {
				calls := 0
				httpClient := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
					calls++
					if r.URL.String() != core.GatewaySandbox+"/v3/alipay/trade/query" {
						t.Fatalf("unexpected URL %s", r.URL)
					}
					body := testutil.ReadAllBody(t, r)
					r.RequestURI = r.URL.RequestURI()
					testutil.VerifyV3Authorization(t, r, body, "")
					if !strings.Contains(r.Header.Get("Authorization"), "app_cert_sn="+testutil.CertSN(t, suite.AppCert)) ||
						r.Header.Get("alipay-root-cert-sn") != testutil.RootCertSNOf(t, suite.Root) {
						t.Fatal("missing or incorrect certificate identity")
					}
					var request trade.QueryRequest
					if err := json.Unmarshal(body, &request); err != nil || request.OutTradeNo != "SANDBOX-ORDER" {
						t.Fatal("request parameters not preserved")
					}
					node := `{"trade_status":"WAIT_BUYER_PAY","total_amount":"0.01"}`
					headers := testutil.V3ResponseHeaders(t, []byte(node), "1764000000200", "sandbox-test", testutil.CertSN(t, suite.AlipayCert))
					if tamper {
						node = strings.ReplaceAll(node, "0.01", "0.02")
					}
					return &http.Response{StatusCode: http.StatusOK, Header: headers, Body: io.NopCloser(strings.NewReader(node))}, nil
				})}
				client, err := newClient(getenv, httpClient)
				if (err != nil) != tt.wantErr {
					t.Fatalf("newClient() = %v; want error=%t", err, tt.wantErr)
				}
				if tt.wantErr {
					if calls != 0 {
						t.Fatal("invalid configuration attempted HTTP")
					}
					break
				}
				service, err := trade.NewService(client)
				if err != nil {
					t.Fatal(err)
				}
				resp, err := service.Query(context.Background(), trade.QueryRequest{OutTradeNo: "SANDBOX-ORDER"})
				if tamper {
					if _, ok := compat.AsType[*core.SignatureError](err); !ok {
						t.Fatalf("tampered response error = %T", err)
					}
				} else if err != nil || !resp.Result.SignatureVerified || resp.Data.TradeStatus != trade.TradeStatusWaitBuyerPay {
					t.Fatalf("valid certificate response error = %v", err)
				}
				if calls != 1 {
					t.Fatalf("HTTP calls = %d; want 1", calls)
				}
			}
		})
	}
}
