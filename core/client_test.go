package core_test

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

// V3 REST 协议测试：本地 httptest 服务独立复算 V3 请求签名
// （从收到的 Authorization 头与方法、URI、请求体重建签名原文），
// 并返回带响应签名头的 JSON；正确报文通过，篡改与缺签失败。

const (
	testAppID   = "2021000000000001"
	testV3Path  = "/v3/alipay/trade/query"
	testV3Query = "/v3/alipay/data/dataservice/bill/downloadurl/query"
)

type queryRequest struct {
	OutTradeNo string  `json:"out_trade_no"`
	TradeNo    *string `json:"trade_no,omitempty"`
}

type queryResponse struct {
	TradeNo     string `json:"trade_no"`
	TradeStatus string `json:"trade_status"`
}

type billRequest struct {
	BillType string  `query:"bill_type"`
	BillDate string  `query:"bill_date,omitempty"`
	Flag     *string `query:"flag,omitempty"`
}

type billResponse struct {
	BillDownloadURL string `json:"bill_download_url"`
}

func newQueryOperation(t *testing.T) core.Operation[queryRequest, queryResponse] {
	t.Helper()
	op, err := core.NewJSONOperation[queryRequest, queryResponse](core.OperationSpec{
		Method:          "alipay.trade.query",
		HTTPMethod:      core.HTTPMethodPost,
		Path:            testV3Path,
		ParameterMode:   core.ParameterModeJSONBody,
		AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		t.Fatalf("NewJSONOperation() error = %v", err)
	}
	return op
}

func newBillOperation(t *testing.T) core.Operation[billRequest, billResponse] {
	t.Helper()
	op, err := core.NewJSONOperation[billRequest, billResponse](core.OperationSpec{
		Method:          "alipay.data.dataservice.bill.downloadurl.query",
		HTTPMethod:      core.HTTPMethodGet,
		Path:            testV3Query,
		ParameterMode:   core.ParameterModeQuery,
		AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		t.Fatalf("NewJSONOperation() error = %v", err)
	}
	return op
}

func newTestClient(t *testing.T, gateway string, opts ...core.ClientOption) *core.Client {
	t.Helper()
	opts = append([]core.ClientOption{option.WithAppID(testAppID), option.WithGateway(gateway)}, opts...)
	client, err := core.NewClient(testutil.NewCredentials(t), opts...)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	return client
}

func TestClientSuccessAndServerSideVerification(t *testing.T) {
	var gotAuthorization, gotRequestID string
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := testutil.ReadAllBody(t, r)
		testutil.VerifyV3Authorization(t, r, body, "")
		gotAuthorization = r.Header.Get("Authorization")
		gotRequestID = r.Header.Get("alipay-request-id")
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("content type = %q", ct)
		}
		if gotRequestID == "" {
			t.Error("alipay-request-id header is required")
		}
		node := `{"trade_no":"2026110522001400001","trade_status":"TRADE_SUCCESS"}`
		for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000000", "resp-nonce-1", "") {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(node))
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	op := newQueryOperation(t)
	resp, err := core.Execute(context.Background(), client, op, core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050001"},
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if !resp.Result.SignatureVerified {
		t.Error("result must be signature verified")
	}
	if resp.Data == nil || resp.Data.TradeNo != "2026110522001400001" || resp.Data.TradeStatus != "TRADE_SUCCESS" {
		t.Fatalf("unexpected data: %+v", resp.Data)
	}
	if !strings.Contains(gotAuthorization, "app_id="+testAppID+",") || strings.Contains(gotAuthorization, "app_cert_sn") {
		t.Errorf("public key mode Authorization must carry app_id without cert sn: %q", gotAuthorization)
	}
	if resp.Result.HTTPStatus != http.StatusOK {
		t.Errorf("http status = %d", resp.Result.HTTPStatus)
	}
}

func TestClientGETQuerySigning(t *testing.T) {
	var gotURI string
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		body := testutil.ReadAllBody(t, r)
		if len(body) != 0 {
			t.Errorf("GET body = %q, want empty", body)
		}
		testutil.VerifyV3Authorization(t, r, body, "")
		gotURI = r.RequestURI
		node := `{"bill_download_url":"https://example.com/bill.zip"}`
		for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000001", "resp-nonce-2", "") {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(node))
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	resp, err := core.Execute(context.Background(), client, newBillOperation(t), core.Request[billRequest]{
		Body: billRequest{BillType: "trade", BillDate: "2026-11-04"},
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if resp.Data == nil || resp.Data.BillDownloadURL != "https://example.com/bill.zip" {
		t.Fatalf("unexpected data: %+v", resp.Data)
	}
	// query 键序确定（Encode 按键排序），签名与传输使用同一份 URI。
	if gotURI != testV3Query+"?bill_date=2026-11-04&bill_type=trade" {
		t.Fatalf("request uri = %q", gotURI)
	}
}

func TestClientTamperDetection(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 签名覆盖原始报文，返回体被篡改：验签必须失败。
		signed := []byte(`{"trade_no":"honest-value","trade_status":"TRADE_SUCCESS"}`)
		tampered := `{"trade_no":"tampered-value","trade_status":"TRADE_SUCCESS"}`
		for key, value := range testutil.V3ResponseHeaders(t, signed, "1764000000002", "resp-nonce-3", "") {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(tampered))
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050002"},
	})
	var sigErr *core.SignatureError
	if !errors.As(err, &sigErr) {
		t.Fatalf("expected *core.SignatureError, got %v", err)
	}
}

func TestClientRejectsMissingResponseSignature(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"trade_no":"x","trade_status":"TRADE_SUCCESS"}`))
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050003"},
	})
	var sigErr *core.SignatureError
	if !errors.As(err, &sigErr) {
		t.Fatalf("expected *core.SignatureError for missing alipay-signature, got %v", err)
	}
}

func TestClientWrongKeyFails(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		node := `{"trade_no":"x","trade_status":"TRADE_SUCCESS"}`
		// 响应签名来自商户私钥而非网关私钥：公钥模式验签必须失败。
		for key, value := range testutil.V3ResponseHeadersWithAppKey(t, []byte(node), "1764000000003", "resp-nonce-4") {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(node))
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050004"},
	})
	var sigErr *core.SignatureError
	if !errors.As(err, &sigErr) {
		t.Fatalf("expected *core.SignatureError for wrong key, got %v", err)
	}
}

func TestClientBusinessErrorSignedAndClassified(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := `{"code":"ACQ.TRADE_NOT_EXIST","message":"交易不存在","links":"https://opendocs.alipay.com/common/02km9f"}`
		for key, value := range testutil.V3ResponseHeaders(t, []byte(body), "1764000000004", "resp-nonce-5", "") {
			w.Header()[key] = value
		}
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(body))
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	resp, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050005"},
	})
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *core.APIError, got %v", err)
	}
	if apiErr.Code != "ACQ.TRADE_NOT_EXIST" || apiErr.Msg != "交易不存在" || !apiErr.Verified {
		t.Fatalf("unexpected api error: %+v", apiErr)
	}
	if apiErr.Links == "" {
		t.Error("links must be preserved")
	}
	if resp.Result == nil || resp.Result.HTTPStatus != http.StatusBadRequest || !resp.Result.SignatureVerified {
		t.Fatalf("result metadata missing: %+v", resp.Result)
	}
}

func TestClientUnsignedErrorIsUnverified(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"code":"insufficient-isv-permissions","message":"无权限"}`))
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	resp, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050006"},
	})
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *core.APIError, got %v", err)
	}
	if apiErr.Verified || resp.Result == nil || resp.Result.SignatureVerified {
		t.Error("unsigned error must not be marked verified")
	}
	if apiErr.Code != "insufficient-isv-permissions" {
		t.Fatalf("code = %q", apiErr.Code)
	}
}

func TestClientUnparseableErrorBody(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`upstreamoops`))
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "202611050007"},
	})
	var httpErr *core.HTTPError
	if !errors.As(err, &httpErr) || httpErr.Status != http.StatusInternalServerError {
		t.Fatalf("expected *core.HTTPError, got %v", err)
	}
}

func TestClientAppAuthTokenIsHeaderAndSigned(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := testutil.ReadAllBody(t, r)
		testutil.VerifyV3Authorization(t, r, body, "app-auth-token-1")
		if got := r.Header.Get("alipay-app-auth-token"); got != "app-auth-token-1" {
			t.Errorf("app auth token header = %q", got)
		}
		node := `{"trade_no":"x","trade_status":"TRADE_SUCCESS"}`
		for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000005", "resp-nonce-6", "") {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(node))
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body:    queryRequest{OutTradeNo: "202611050008"},
		Options: core.NewCallSettings(core.WithAppAuthToken("app-auth-token-1")),
	})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
}

func TestClientNotifyURLOptionRejected(t *testing.T) {
	var called bool
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body:    queryRequest{OutTradeNo: "202611050009"},
		Options: core.NewCallSettings(core.WithNotifyURL("https://merchant.example.com/notify")),
	})
	var valErr *core.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("expected *core.ValidationError, got %v", err)
	}
	if called {
		t.Error("request must not be sent when notify_url option is set")
	}
}

func TestClientSpecValidation(t *testing.T) {
	tests := []struct {
		name   string
		spec   core.OperationSpec
		reason string
	}{
		{
			name: "missing method",
			spec: core.OperationSpec{
				HTTPMethod: core.HTTPMethodPost, Path: testV3Path,
				ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
			},
			reason: "method",
		},
		{
			name: "unknown http method",
			spec: core.OperationSpec{
				Method: "x", HTTPMethod: core.HTTPMethod("FETCH"), Path: testV3Path,
				ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
			},
			reason: "http_method",
		},
		{
			name: "GET with json body",
			spec: core.OperationSpec{
				Method: "x", HTTPMethod: core.HTTPMethodGet, Path: testV3Path,
				ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
			},
			reason: "parameter_mode",
		},
		{
			name: "path without prefix",
			spec: core.OperationSpec{
				Method: "x", HTTPMethod: core.HTTPMethodPost, Path: "/alipay/trade/query",
				ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
			},
			reason: "path",
		},
		{
			name: "path carries query",
			spec: core.OperationSpec{
				Method: "x", HTTPMethod: core.HTTPMethodPost, Path: testV3Path + "?a=1",
				ParameterMode: core.ParameterModeJSONBody, AuthRequirement: core.AuthRequirementAny,
			},
			reason: "path",
		},
		{
			name: "unknown parameter mode",
			spec: core.OperationSpec{
				Method: "x", HTTPMethod: core.HTTPMethodPost, Path: testV3Path,
				ParameterMode: core.ParameterMode(9), AuthRequirement: core.AuthRequirementAny,
			},
			reason: "parameter_mode",
		},
		{
			name: "missing auth requirement",
			spec: core.OperationSpec{
				Method: "x", HTTPMethod: core.HTTPMethodPost, Path: testV3Path,
				ParameterMode: core.ParameterModeJSONBody,
			},
			reason: "auth_requirement",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := core.NewJSONOperation[queryRequest, queryResponse](tt.spec)
			var valErr *core.ValidationError
			if !errors.As(err, &valErr) || valErr.Field != tt.reason {
				t.Fatalf("expected validation error on %s, got %v", tt.reason, err)
			}
		})
	}
}

func TestJSONOperationInfersParameterMode(t *testing.T) {
	// GET 未声明 ParameterMode 时推断为 query；POST 推断为 JSON body。
	getOp, err := core.NewJSONOperation[billRequest, billResponse](core.OperationSpec{
		Method:     "alipay.data.dataservice.bill.downloadurl.query",
		HTTPMethod: core.HTTPMethodGet, Path: testV3Query,
		AuthRequirement: core.AuthRequirementAny,
	})
	if err != nil {
		t.Fatalf("NewJSONOperation() error = %v", err)
	}
	if got := getOp.Spec().ParameterMode; got != core.ParameterModeQuery {
		t.Fatalf("get parameter mode = %v, want query", got)
	}
	postOp := newQueryOperation(t)
	if got := postOp.Spec().ParameterMode; got != core.ParameterModeJSONBody {
		t.Fatalf("post parameter mode = %v, want json_body", got)
	}
}

func TestClientDoesNotModifyCallerParams(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		node := `{"bill_download_url":"u"}`
		for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "1764000000009", "resp-nonce-a", "") {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(node))
	}))
	defer gateway.Close()

	client := newTestClient(t, gateway.URL)
	op := newBillOperation(t)
	body := billRequest{BillType: "trade"}
	encoder := core.QueryEncoder[billRequest]()
	params, err := encoder.Encode(body)
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}
	before := compat.CloneValues(params.Form)
	_, err = core.Execute(context.Background(), client, op, core.Request[billRequest]{Body: body})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	for key, values := range before {
		if got := params.Form[key]; len(got) != len(values) {
			t.Fatalf("caller params modified for key %q", key)
		}
	}
}

func TestClientResponseSizeLimit(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(strings.Repeat("x", 4096)))
	}))
	defer gateway.Close()
	client := newTestClient(t, gateway.URL, option.WithMaxResponseBytes(1024))
	_, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "size-1"},
	})
	var protoErr *core.ProtocolError
	if !errors.As(err, &protoErr) {
		t.Fatalf("error type = %T, want *core.ProtocolError", err)
	}
	if !strings.Contains(protoErr.Reason, "size limit") {
		t.Fatalf("Reason = %q", protoErr.Reason)
	}
}

func TestClientCallSettingsIsolation(t *testing.T) {
	var mu sync.Mutex
	seen := map[string]int{}
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := testutil.ReadAllBody(t, r)
		token := r.Header.Get("alipay-app-auth-token")
		if token == "" {
			t.Error("app auth token header is required")
			return
		}
		// 签名原文必须包含本次调用的 app_auth_token，且与头一致。
		testutil.VerifyV3Authorization(t, r, body, token)
		mu.Lock()
		seen[token]++
		mu.Unlock()
		node := `{"trade_no":"x","trade_status":"TRADE_SUCCESS"}`
		for key, value := range testutil.V3ResponseHeaders(t, []byte(node), "176400000000a", "resp-nonce-b", "") {
			w.Header()[key] = value
		}
		_, _ = w.Write([]byte(node))
	}))
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	op := newQueryOperation(t)

	var wg sync.WaitGroup
	tokens := []string{"token-A", "token-B", "token-C"}
	for _, token := range tokens {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range 5 {
				_, err := core.Execute(context.Background(), client, op, core.Request[queryRequest]{
					Body:    queryRequest{OutTradeNo: fmt.Sprintf("T-%s-%d", token, i)},
					Options: core.NewCallSettings(core.WithAppAuthToken(token)),
				})
				if err != nil {
					t.Errorf("Execute() error = %v", err)
				}
			}
		}()
	}
	wg.Wait()
	for _, token := range tokens {
		if seen[token] != 5 {
			t.Fatalf("token %s: %d requests, want 5", token, seen[token])
		}
	}
}

// mockExecutor 实现非泛型 Executor 接口，模拟领域包注入的测试执行器。
type mockExecutor struct {
	result *core.APIResult
	err    error
	called bool
	req    core.WireRequest
}

func (m *mockExecutor) Do(_ context.Context, req core.WireRequest) (*core.APIResult, error) {
	m.called = true
	m.req = req
	return m.result, m.err
}

func TestPackageExecuteWithMockExecutor(t *testing.T) {
	executor := &mockExecutor{result: &core.APIResult{
		HTTPStatus:        200,
		SignatureVerified: true,
		Node:              []byte(`{"trade_no":"T123","trade_status":"TRADE_SUCCESS"}`),
		Header:            http.Header{"X-Test": {"1"}},
	}}
	op := newQueryOperation(t)
	resp, err := core.Execute(context.Background(), executor, op,
		core.Request[queryRequest]{Body: queryRequest{OutTradeNo: "T1"}})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if !executor.called {
		t.Fatal("executor not called")
	}
	if executor.req.Spec.Method != "alipay.trade.query" {
		t.Fatalf("Spec.Method = %q", executor.req.Spec.Method)
	}
	if resp.Data.TradeNo != "T123" {
		t.Fatalf("Data = %+v", resp.Data)
	}
	if executor.req.Settings.NotifyURL != "" {
		t.Fatal("unexpected notify url")
	}
}

func TestPackageExecuteFailureSemantics(t *testing.T) {
	op := newQueryOperation(t)
	t.Run("执行器错误时保留元信息", func(t *testing.T) {
		executor := &mockExecutor{err: errors.New("boom"), result: &core.APIResult{HTTPStatus: 599}}
		resp, err := core.Execute(context.Background(), executor, op, core.Request[queryRequest]{})
		if err == nil || resp == nil || resp.Result == nil || resp.Result.HTTPStatus != 599 {
			t.Fatalf("resp = %+v, err = %v", resp, err)
		}
	})
	t.Run("未验证结果不得进入解码", func(t *testing.T) {
		executor := &mockExecutor{result: &core.APIResult{
			HTTPStatus: 200, SignatureVerified: false, Node: []byte(`{"trade_no":"x"}`),
		}}
		resp, err := core.Execute(context.Background(), executor, op, core.Request[queryRequest]{})
		if !errors.As(err, new(*core.SignatureError)) {
			t.Fatalf("error = %v, want SignatureError", err)
		}
		if resp == nil || resp.Data != nil {
			t.Fatal("Data must be nil")
		}
	})
	t.Run("零值 Operation 失败", func(t *testing.T) {
		var zero core.Operation[queryRequest, queryResponse]
		executor := &mockExecutor{}
		if _, err := core.Execute(context.Background(), executor, zero, core.Request[queryRequest]{}); err == nil {
			t.Fatal("zero operation must fail before request")
		}
		if executor.called {
			t.Fatal("executor must not be called")
		}
	})
	t.Run("nil 执行器失败", func(t *testing.T) {
		if _, err := core.Execute[queryRequest, queryResponse](context.Background(), nil, op, core.Request[queryRequest]{}); err == nil {
			t.Fatal("nil executor must fail")
		}
	})
	t.Run("解码失败返回协议错误", func(t *testing.T) {
		executor := &mockExecutor{result: &core.APIResult{
			HTTPStatus: 200, SignatureVerified: true, Node: []byte(`{"trade_no":123456}`), // 类型不匹配
		}}
		_, err := core.Execute(context.Background(), executor, op, core.Request[queryRequest]{})
		if !errors.As(err, new(*core.ProtocolError)) {
			t.Fatalf("error = %v, want ProtocolError", err)
		}
	})
}

// captureHandler 捕获 slog 记录，用于校验级别与字段。
type captureHandler struct {
	mu      sync.Mutex
	records []slog.Record
	errOn   map[string]bool // 命中消息时返回错误
}

func (h *captureHandler) Enabled(context.Context, slog.Level) bool { return true }

func (h *captureHandler) Handle(_ context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.errOn[r.Message] {
		return errors.New("handler failure")
	}
	h.records = append(h.records, r.Clone())
	return nil
}

func (h *captureHandler) WithAttrs([]slog.Attr) slog.Handler { return h }
func (h *captureHandler) WithGroup(string) slog.Handler      { return h }

func (h *captureHandler) find(t *testing.T, level slog.Level, msg string) *slog.Record {
	t.Helper()
	h.mu.Lock()
	defer h.mu.Unlock()
	for i := range h.records {
		r := &h.records[i]
		if r.Level == level && r.Message == msg {
			return r
		}
	}
	return nil
}

func (h *captureHandler) attrValue(t *testing.T, r *slog.Record, key string) (slog.Value, bool) {
	t.Helper()
	var found slog.Value
	var ok bool
	r.Attrs(func(a slog.Attr) bool {
		if a.Key == key {
			found, ok = a.Value, true
			return false
		}
		return true
	})
	return found, ok
}

func TestClientLogging(t *testing.T) {
	t.Run("成功为 Debug 且字段完整", func(t *testing.T) {
		handler := &captureHandler{}
		gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			node := []byte(`{"trade_no":"T1","trade_status":"TRADE_SUCCESS"}`)
			for key, value := range testutil.V3ResponseHeaders(t, node, "176400000000d", "resp-nonce-e", "") {
				w.Header()[key] = value
			}
			_, _ = w.Write(node)
		}))
		defer gateway.Close()
		client := newTestClient(t, gateway.URL, option.WithLogHandler(handler))
		if _, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
			Body: queryRequest{OutTradeNo: "T1"},
		}); err != nil {
			t.Fatal(err)
		}
		r := handler.find(t, slog.LevelDebug, "alipay request succeeded")
		if r == nil {
			t.Fatal("success log record not found")
		}
		for key, want := range map[string]string{
			"sdk": "alipay", "operation": "alipay.trade.query",
		} {
			if v, ok := handler.attrValue(t, r, key); !ok || v.String() != want {
				t.Errorf("attr %q = %v (present=%v), want %q", key, v, ok, want)
			}
		}
		if v, ok := handler.attrValue(t, r, "duration"); !ok || v.Duration() <= 0 {
			t.Errorf("attr duration = %v (present=%v)", v, ok)
		}
		if v, ok := handler.attrValue(t, r, "http_status"); !ok || v.Int64() != 200 {
			t.Errorf("attr http_status = %v", v)
		}
	})
	t.Run("业务拒绝为 Warn 且含 code", func(t *testing.T) {
		handler := &captureHandler{}
		body := `{"code":"ACQ.TRADE_NOT_EXIST","message":"交易不存在"}`
		gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for key, value := range testutil.V3ResponseHeaders(t, []byte(body), "176400000000e", "resp-nonce-f", "") {
				w.Header()[key] = value
			}
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(body))
		}))
		defer gateway.Close()
		client := newTestClient(t, gateway.URL, option.WithLogHandler(handler))
		_, _ = core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
			Body: queryRequest{OutTradeNo: "log-2"},
		})
		r := handler.find(t, slog.LevelWarn, "alipay request failed")
		if r == nil {
			t.Fatal("warn record not found")
		}
		if v, _ := handler.attrValue(t, r, "code"); v.String() != "ACQ.TRADE_NOT_EXIST" {
			t.Fatalf("code attr = %v", v)
		}
	})
	t.Run("安全失败为 Error 且 error_kind 为 signature", func(t *testing.T) {
		handler := &captureHandler{}
		gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"trade_no":"x"}`)) // 缺少签名头
		}))
		defer gateway.Close()
		client := newTestClient(t, gateway.URL, option.WithLogHandler(handler))
		_, _ = core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
			Body: queryRequest{OutTradeNo: "log-3"},
		})
		r := handler.find(t, slog.LevelError, "alipay request failed")
		if r == nil {
			t.Fatal("error record not found")
		}
		if v, _ := handler.attrValue(t, r, "error_kind"); v.String() != "signature" {
			t.Fatalf("error_kind = %v", v)
		}
	})
	t.Run("未配置日志时静默", func(t *testing.T) {
		gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			node := []byte(`{"trade_no":"T1","trade_status":"TRADE_SUCCESS"}`)
			for key, value := range testutil.V3ResponseHeaders(t, node, "1764000000010", "resp-nonce-g", "") {
				w.Header()[key] = value
			}
			_, _ = w.Write(node)
		}))
		defer gateway.Close()
		client := newTestClient(t, gateway.URL) // 不注入日志
		if _, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
			Body: queryRequest{OutTradeNo: "T1"},
		}); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Handler 错误不改变 API 结果", func(t *testing.T) {
		handler := &captureHandler{errOn: map[string]bool{"alipay request succeeded": true}}
		gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			node := []byte(`{"trade_no":"T9","trade_status":"TRADE_SUCCESS"}`)
			for key, value := range testutil.V3ResponseHeaders(t, node, "1764000000011", "resp-nonce-h", "") {
				w.Header()[key] = value
			}
			_, _ = w.Write(node)
		}))
		defer gateway.Close()
		client := newTestClient(t, gateway.URL, option.WithLogHandler(handler))
		resp, err := core.Execute(context.Background(), client, newQueryOperation(t), core.Request[queryRequest]{
			Body: queryRequest{OutTradeNo: "T9"},
		})
		if err != nil || resp.Data.TradeNo != "T9" {
			t.Fatalf("handler failure must not affect result: %v, %+v", err, resp)
		}
	})
}

func TestSignedURIWithRepeatedQueryAndGatewayPrefix(t *testing.T) {
	type req struct {
		Items []string `query:"item"`
		ID    string   `path:"id"`
	}
	op, err := core.NewJSONOperation[req, struct{}](core.OperationSpec{Method: "alipay.test.query", HTTPMethod: core.HTTPMethodGet, Path: "/v3/items/{id}", AuthRequirement: core.AuthRequirementAny})
	if err != nil {
		t.Fatal(err)
	}
	for _, base := range []string{"/", "/proxy%2Fprefix/"} {
		t.Run(base, func(t *testing.T) {
			gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if want := strings.TrimSuffix(base, "/") + "/v3/items/a%2Fb?item=b%2B&item=&item=a%26"; r.RequestURI != want {
					t.Errorf("URI %q, want %q", r.RequestURI, want)
				}
				testutil.VerifyV3Authorization(t, r, nil, "")
				for key, value := range testutil.V3ResponseHeaders(t, []byte("{}"), "1764000000001", "response-nonce", "") {
					w.Header()[key] = value
				}
				fmt.Fprint(w, "{}")
			}))
			defer gateway.Close()
			_, err := core.Execute(t.Context(), newTestClient(t, gateway.URL+base), op, core.Request[req]{Body: req{ID: "a/b", Items: []string{"b+", "", "a&"}}})
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestRejectsNonBaseGateway(t *testing.T) {
	for _, url := range []string{"https://gateway.test?x=1", "https://gateway.test?", "https://gateway.test/path#fragment", "https://user@gateway.test"} {
		_, err := core.NewClient(testutil.NewCredentials(t), option.WithAppID(testAppID), option.WithGateway(url))
		var cfg *core.ConfigError
		if !errors.As(err, &cfg) || cfg.Field != "gateway" {
			t.Fatalf("invalid gateway: %v", err)
		}
	}
}

func TestSignedNullCannotPopulateResponseStruct(t *testing.T) {
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		node := []byte(" null\n")
		for key, value := range testutil.V3ResponseHeaders(t, node, "1764000000001", "response-nonce", "") {
			w.Header()[key] = value
		}
		w.Write(node)
	}))
	defer gateway.Close()
	resp, err := core.Execute(t.Context(), newTestClient(t, gateway.URL), newQueryOperation(t), core.Request[queryRequest]{})
	var protocolErr *core.ProtocolError
	if !errors.As(err, &protocolErr) || resp == nil || resp.Result == nil || !resp.Result.SignatureVerified || resp.Data != nil {
		t.Fatalf("unexpected response %+v, err %v", resp, err)
	}
	if value, err := core.JSONDecoder[*queryResponse]().Decode([]byte("null")); err != nil || value != nil {
		t.Fatalf("nullable type: %v, %v", value, err)
	}
}
