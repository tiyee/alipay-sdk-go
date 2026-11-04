package core_test

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strings"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/tiyee/go-alipay/core"
	"github.com/tiyee/go-alipay/core/option"
	"github.com/tiyee/go-alipay/internal/testutil"
)

// 测试目标：本地 httptest 服务独立校验 SDK 请求（服务端签名核对），
// 并返回独立构造的签名响应；正确报文通过，篡改报文失败。

const (
	testAppID       = "2021000000000001"
	testMethod      = "alipay.trade.query"
	testRespKey     = "alipay_trade_query_response"
	testFormMethod  = "alipay.trade.custom.form" // Form 模式扩展接口
	testFormRespKey = "alipay_trade_custom_form_response"
)

type queryRequest struct {
	OutTradeNo string  `json:"out_trade_no,omitempty"`
	TradeNo    *string `json:"trade_no,omitempty"`
	Amount     string  `json:"amount,omitempty"`
	Subject    string  `json:"subject,omitempty"`
}

type queryResponse struct {
	Code        string `json:"code"`
	TradeNo     string `json:"trade_no"`
	TradeStatus string `json:"trade_status"`
	TotalAmount string `json:"total_amount"`
}

func newTestOperation(t *testing.T) core.Operation[queryRequest, queryResponse] {
	t.Helper()
	op, err := core.NewJSONOperation[queryRequest, queryResponse](core.OperationSpec{
		Method:        testMethod,
		Version:       "1.0",
		ResponseKey:   testRespKey,
		ParameterMode: core.ParameterModeBizJSON,
	})
	if err != nil {
		t.Fatalf("NewJSONOperation() error = %v", err)
	}
	return op
}

// testFormRequest 演示直接参数模式：业务字段由调用方编码为表单参数。
type testFormRequest struct {
	OutTradeNo string
	Overrides  url.Values // 测试用：允许注入受保护参数名
}

func testFormEncoder() core.Encoder[testFormRequest] {
	return core.EncoderFunc[testFormRequest](func(value testFormRequest) (core.EncodedParams, error) {
		if value.OutTradeNo == "" {
			return core.EncodedParams{}, errors.New("out_trade_no is required")
		}
		form := url.Values{"out_trade_no": {value.OutTradeNo}}
		for key, vs := range value.Overrides {
			form[key] = vs
		}
		return core.EncodedParams{Mode: core.ParameterModeForm, Form: form}, nil
	})
}

func newTestFormOperation(t *testing.T) core.Operation[testFormRequest, queryResponse] {
	t.Helper()
	op, err := core.NewOperation[testFormRequest, queryResponse](core.OperationSpec{
		Method:        testFormMethod,
		Version:       "1.0",
		ResponseKey:   testFormRespKey,
		ParameterMode: core.ParameterModeForm,
	}, testFormEncoder(), core.JSONDecoder[queryResponse]())
	if err != nil {
		t.Fatalf("NewOperation() error = %v", err)
	}
	return op
}

// newGateway 启动测试网关：每个请求先做服务端签名核对，
// 再由 respond 决定响应。核对失败会记录为测试错误。
func newGateway(t *testing.T, respond func(form url.Values) (int, string)) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Errorf("gateway parse form: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		form := r.PostForm
		signContent := testutil.SignForm(form)
		if err := testutil.VerifyWithAppKey([]byte(signContent), form.Get("sign")); err != nil {
			t.Errorf("gateway: request signature invalid: %v\nsign content: %s", err, signContent)
		}
		status, body := respond(form)
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
}

// signedResponse 构造带网关签名的同步响应。
func signedResponse(t *testing.T, responseKey, node string) string {
	t.Helper()
	sign, err := testutil.SignWithAlipayKey([]byte(node))
	if err != nil {
		t.Fatalf("sign response: %v", err)
	}
	return `{"` + responseKey + `":` + node + `,"sign":"` + sign + `"}`
}

func successNode(extra string) string {
	return `{"code":"10000","msg":"Success"` + extra + `}`
}

func newTestClient(t *testing.T, gatewayURL string, opts ...core.ClientOption) *core.Client {
	t.Helper()
	base := []core.ClientOption{
		option.WithAppID(testAppID),
		option.WithGateway(gatewayURL),
		option.WithPrivateKey(testutil.AppPrivateKeyPEM),
		option.WithAlipayPublicKey(testutil.AlipayPublicKeyPEM),
	}
	client, err := core.NewClient(append(base, opts...)...)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	return client
}

func executeQuery(t *testing.T, client *core.Client, op core.Operation[queryRequest, queryResponse]) (*core.Response[queryResponse], error) {
	t.Helper()
	return client.Execute(context.Background(), op, core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "20150320010101001", Amount: "88.88", Subject: "测试商品"},
	})
}

func TestClientSuccessAndServerSideVerification(t *testing.T) {
	var gotForm url.Values
	gateway := newGateway(t, func(form url.Values) (int, string) {
		gotForm = form
		return http.StatusOK, signedResponse(t, testRespKey,
			`{"code":"10000","msg":"Success","trade_no":"2026010122001400001234567890","trade_status":"TRADE_SUCCESS","total_amount":"88.88"}`)
	})
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	op := newTestOperation(t)

	resp, err := executeQuery(t, client, op)
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if resp.Data == nil {
		t.Fatal("Data is nil")
	}
	if resp.Data.TradeStatus != "TRADE_SUCCESS" || resp.Data.TotalAmount != "88.88" {
		t.Fatalf("Data = %+v", resp.Data)
	}
	if !resp.Result.SignatureVerified {
		t.Fatal("SignatureVerified = false")
	}
	if resp.Result.HTTPStatus != http.StatusOK {
		t.Fatalf("HTTPStatus = %d", resp.Result.HTTPStatus)
	}

	// 服务端核对公共参数与业务参数。
	want := map[string]string{
		"app_id":      testAppID,
		"method":      testMethod,
		"format":      "JSON",
		"charset":     "utf-8",
		"sign_type":   "RSA2",
		"version":     "1.0",
		"biz_content": `{"out_trade_no":"20150320010101001","amount":"88.88","subject":"测试商品"}`,
	}
	for key, wantValue := range want {
		if got := gotForm.Get(key); got != wantValue {
			t.Errorf("form[%q] = %q, want %q", key, got, wantValue)
		}
	}
	ts, err := time.ParseInLocation("2006-01-02 15:04:05", gotForm.Get("timestamp"), time.Local)
	if err != nil {
		t.Fatalf("timestamp format invalid: %q", gotForm.Get("timestamp"))
	}
	if drift := time.Since(ts); drift < -time.Minute || drift > time.Minute {
		t.Errorf("timestamp = %v, drift %v", ts, drift)
	}
	if gotForm.Get("notify_url") != "" || gotForm.Get("app_auth_token") != "" {
		t.Error("call settings should be absent when unset")
	}
}

func TestClientRawResponseNodeSurvivesForVerification(t *testing.T) {
	// 节点含 \/ 转义、空格与中文，网关按原样签名；
	// 客户端若重新序列化将导致验签失败。
	node := `{"code":"10000" , "msg":"Success","url":"https:\/\/example.com\/pay","name":"手机\/商品"}`
	gateway := newGateway(t, func(url.Values) (int, string) {
		return http.StatusOK, signedResponse(t, testRespKey, node)
	})
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	op := newTestOperation(t)

	resp, err := executeQuery(t, client, op)
	if err != nil {
		t.Fatalf("Execute() error = %v (raw node verification failed)", err)
	}
	if string(resp.Result.Node) != node {
		t.Fatalf("Node = %q, want original %q", resp.Result.Node, node)
	}
}

func TestClientBusinessErrorIsSignedAndClassified(t *testing.T) {
	node := `{"code":"40004","msg":"Business Failed","sub_code":"ACQ.TRADE_NOT_EXIST","sub_msg":"交易不存在"}`
	gateway := newGateway(t, func(url.Values) (int, string) {
		return http.StatusOK, signedResponse(t, testRespKey, node)
	})
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	op := newTestOperation(t)

	resp, err := executeQuery(t, client, op)
	if err == nil {
		t.Fatal("expect APIError")
	}
	if resp == nil || resp.Data != nil {
		t.Fatalf("Data must be nil on failure, resp = %+v", resp)
	}
	if resp.Result == nil || resp.Result.HTTPStatus != http.StatusOK {
		t.Fatal("Result with HTTP metadata expected")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error type = %T, want *core.APIError", err)
	}
	if !apiErr.Verified {
		t.Error("signed business error should be Verified")
	}
	if apiErr.Code != core.ResponseCode("40004") || apiErr.SubCode != "ACQ.TRADE_NOT_EXIST" {
		t.Fatalf("APIError = %+v", apiErr)
	}
}

func TestClientErrorResponseIsUnverified(t *testing.T) {
	body := `{"error_response":{"code":"40002","msg":"Invalid Arguments","sub_code":"isv.invalid-signature","sub_msg":"验签出错"}}`
	gateway := newGateway(t, func(url.Values) (int, string) {
		return http.StatusOK, body
	})
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	op := newTestOperation(t)

	resp, err := executeQuery(t, client, op)
	if err == nil {
		t.Fatal("expect APIError for error_response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error type = %T, want *core.APIError", err)
	}
	if apiErr.Verified {
		t.Error("error_response is unsigned; Verified must be false")
	}
	if apiErr.Code != core.ResponseCode("40002") {
		t.Fatalf("Code = %q", apiErr.Code)
	}
	if resp == nil || resp.Result == nil || resp.Result.HTTPStatus != http.StatusOK {
		t.Fatal("HTTP metadata should be preserved")
	}
}

func TestClientRejectsMissingSignatures(t *testing.T) {
	node := `{"code":"10000","msg":"Success","trade_no":"2026010122001400001234567890"}`
	tests := []struct {
		name string
		body string
	}{
		{name: "缺失签名", body: `{"` + testRespKey + `":` + node + `}`},
		{name: "空签名", body: `{"` + testRespKey + `":` + node + `,"sign":""}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gateway := newGateway(t, func(url.Values) (int, string) {
				return http.StatusOK, tt.body
			})
			defer gateway.Close()
			client := newTestClient(t, gateway.URL)
			op := newTestOperation(t)
			resp, err := executeQuery(t, client, op)
			if err == nil {
				t.Fatal("expect signature error")
			}
			var sigErr *core.SignatureError
			if !errors.As(err, &sigErr) {
				t.Fatalf("error type = %T, want *core.SignatureError", err)
			}
			if resp == nil || resp.Data != nil {
				t.Fatal("untrusted payload must not produce Data")
			}
		})
	}
}

func TestClientTamperDetection(t *testing.T) {
	// 签名对应原始 node，报文被替换后验签必须失败。
	node := `{"code":"10000","msg":"Success","trade_status":"TRADE_SUCCESS"}`
	tampered := strings.Replace(node, "TRADE_SUCCESS", "TRADE_FINISHED", 1)
	sign, err := testutil.SignWithAlipayKey([]byte(node))
	if err != nil {
		t.Fatalf("sign: %v", err)
	}
	body := `{"` + testRespKey + `":` + tampered + `,"sign":"` + sign + `"}`
	gateway := newGateway(t, func(url.Values) (int, string) {
		return http.StatusOK, body
	})
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	op := newTestOperation(t)
	resp, err := executeQuery(t, client, op)
	if err == nil {
		t.Fatal("tampered payload must fail verification")
	}
	if !errors.As(err, new(*core.SignatureError)) {
		t.Fatalf("error type = %T, want *core.SignatureError", err)
	}
	if resp != nil && resp.Data != nil {
		t.Fatal("tampered payload must not produce Data")
	}
}

func TestClientHTTPFailures(t *testing.T) {
	t.Run("HTTP 500", func(t *testing.T) {
		gateway := newGateway(t, func(url.Values) (int, string) {
			return http.StatusBadGateway, "<html>Bad Gateway</html>"
		})
		defer gateway.Close()
		client := newTestClient(t, gateway.URL)
		op := newTestOperation(t)
		resp, err := executeQuery(t, client, op)
		var httpErr *core.HTTPError
		if !errors.As(err, &httpErr) {
			t.Fatalf("error type = %T, want *core.HTTPError", err)
		}
		if httpErr.Status != http.StatusBadGateway {
			t.Fatalf("Status = %d", httpErr.Status)
		}
		if resp == nil || resp.Result == nil || resp.Result.HTTPStatus != http.StatusBadGateway {
			t.Fatal("HTTP metadata should be preserved")
		}
	})
	t.Run("重定向不跟随", func(t *testing.T) {
		gateway := newGateway(t, func(url.Values) (int, string) {
			return http.StatusFound, ""
		})
		defer gateway.Close()
		client := newTestClient(t, gateway.URL)
		op := newTestOperation(t)
		_, err := executeQuery(t, client, op)
		var httpErr *core.HTTPError
		if !errors.As(err, &httpErr) || httpErr.Status != http.StatusFound {
			t.Fatalf("error = %v, want 302 HTTPError", err)
		}
	})
	t.Run("HTTP 错误携带有界片段且 Error 不含片段", func(t *testing.T) {
		long := strings.Repeat("<div>error</div> ", 100)
		gateway := newGateway(t, func(url.Values) (int, string) {
			return http.StatusServiceUnavailable, long
		})
		defer gateway.Close()
		client := newTestClient(t, gateway.URL)
		op := newTestOperation(t)
		_, err := executeQuery(t, client, op)
		var httpErr *core.HTTPError
		if !errors.As(err, &httpErr) {
			t.Fatalf("error type = %T", err)
		}
		if httpErr.BodySnippet == "" || len(httpErr.BodySnippet) > 512 {
			t.Fatalf("BodySnippet = %q", httpErr.BodySnippet)
		}
		if strings.Contains(httpErr.Error(), "div") {
			t.Error("Error() should not include body snippet")
		}
	})
}

func TestClientResponseSizeLimit(t *testing.T) {
	gateway := newGateway(t, func(url.Values) (int, string) {
		return http.StatusOK, strings.Repeat("x", 4096)
	})
	defer gateway.Close()
	client := newTestClient(t, gateway.URL, option.WithMaxResponseBytes(1024))
	op := newTestOperation(t)
	_, err := executeQuery(t, client, op)
	var protoErr *core.ProtocolError
	if !errors.As(err, &protoErr) {
		t.Fatalf("error type = %T, want *core.ProtocolError", err)
	}
	if !strings.Contains(protoErr.Reason, "size limit") {
		t.Fatalf("Reason = %q", protoErr.Reason)
	}
}

func TestClientTimeoutWithVirtualTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		gateway := httptest.NewTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(2 * time.Hour) // 虚拟时间，立即流逝
		}))
		client := newTestClient(t, "http://alipay.test/gateway.do",
			option.WithHTTPClient(gateway.Client()), option.WithRequestTimeout(time.Minute))
		op := newTestOperation(t)
		start := time.Now()
		_, err := executeQuery(t, client, op)
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("error = %v, want DeadlineExceeded", err)
		}
		var transportErr *core.TransportError
		if !errors.As(err, &transportErr) {
			t.Fatalf("error type = %T, want *core.TransportError", err)
		}
		if elapsed := time.Since(start); elapsed < time.Minute {
			t.Fatalf("timeout returned too early: %v", elapsed)
		}
	})
}

func TestClientContextCancellation(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		gateway := httptest.NewTestServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Hour)
		}))
		client := newTestClient(t, "http://alipay.test/gateway.do",
			option.WithHTTPClient(gateway.Client()), option.WithRequestTimeout(2*time.Hour))
		op := newTestOperation(t)
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(time.Second)
			cancel()
		}()
		start := time.Now()
		_, err := client.Execute(ctx, op, core.Request[queryRequest]{})
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context.Canceled", err)
		}
		var transportErr *core.TransportError
		if !errors.As(err, &transportErr) || transportErr.Kind() != core.ErrorKindCanceled {
			t.Fatalf("error = %v, want canceled transport error", err)
		}
		if elapsed := time.Since(start); elapsed < time.Second {
			t.Fatalf("cancel returned too early: %v", elapsed)
		}
	})
}

func TestClientRejectsProtectedParamOverride(t *testing.T) {
	gateway := newGateway(t, func(url.Values) (int, string) {
		return http.StatusOK, signedResponse(t, testFormRespKey, successNode(""))
	})
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	op := newTestFormOperation(t)
	_, err := client.Execute(context.Background(), op, core.Request[testFormRequest]{
		Body: testFormRequest{OutTradeNo: "T001",
			Overrides: url.Values{"app_id": {"evil"}, "sign": {"evil"}, "biz_content": {"evil"}}},
	})
	var valErr *core.ValidationError
	if !errors.As(err, &valErr) {
		t.Fatalf("error = %v, want protected param ValidationError", err)
	}
	protected := map[string]bool{"app_id": true, "sign": true, "biz_content": true}
	if !protected[valErr.Field] {
		t.Fatalf("field = %q, want one of the protected params", valErr.Field)
	}
}

func TestClientDoesNotModifyCallerParams(t *testing.T) {
	// 编码器返回的表单即调用方输入；Client 不得向其添加公共参数或签名。
	callerForm := url.Values{"out_trade_no": {"T001"}}
	frozen := maps.Clone(callerForm)
	op, err := core.NewOperation[testFormRequest, queryResponse](core.OperationSpec{
		Method:        testFormMethod,
		Version:       "1.0",
		ResponseKey:   testFormRespKey,
		ParameterMode: core.ParameterModeForm,
	}, core.EncoderFunc[testFormRequest](func(testFormRequest) (core.EncodedParams, error) {
		return core.EncodedParams{Mode: core.ParameterModeForm, Form: callerForm}, nil
	}), core.JSONDecoder[queryResponse]())
	if err != nil {
		t.Fatalf("NewOperation() error = %v", err)
	}
	gateway := newGateway(t, func(url.Values) (int, string) {
		return http.StatusOK, signedResponse(t, testFormRespKey, successNode(""))
	})
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	if _, err := client.Execute(context.Background(), op, core.Request[testFormRequest]{
		Body: testFormRequest{OutTradeNo: "T001"}}); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if !maps.EqualFunc(callerForm, frozen, slices.Equal) {
		t.Fatalf("caller params modified: %v", callerForm)
	}
}

func TestClientCallSettingsIsolation(t *testing.T) {
	var mu sync.Mutex
	seen := map[string][]string{}
	gateway := newGateway(t, func(form url.Values) (int, string) {
		mu.Lock()
		defer mu.Unlock()
		seen[form.Get("app_auth_token")] = append(seen[form.Get("app_auth_token")], form.Get("notify_url"))
		return http.StatusOK, signedResponse(t, testRespKey, successNode(""))
	})
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	op := newTestOperation(t)

	var wg sync.WaitGroup
	tokens := []string{"token-A", "token-B", "token-C"}
	for _, token := range tokens {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range 5 {
				_, err := client.Execute(context.Background(), op, core.Request[queryRequest]{
					Body:    queryRequest{OutTradeNo: fmt.Sprintf("T-%s-%d", token, i)},
					Options: core.NewCallSettings(core.WithAppAuthToken(token), core.WithNotifyURL("https://notify.example/"+token)),
				})
				if err != nil {
					t.Errorf("Execute() error = %v", err)
				}
			}
		}()
	}
	wg.Wait()
	for _, token := range tokens {
		urls := seen[token]
		if len(urls) != 5 {
			t.Fatalf("token %s: %d requests, want 5", token, len(urls))
		}
		for _, u := range urls {
			if u != "https://notify.example/"+token {
				t.Fatalf("token %s polluted with notify url %q", token, u)
			}
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
		Node:              []byte(`{"code":"10000","msg":"Success","trade_no":"T123","trade_status":"TRADE_SUCCESS","total_amount":"1.00"}`),
		Header:            http.Header{"X-Test": {"1"}},
	}}
	op := newTestOperation(t)
	resp, err := core.Execute(context.Background(), executor, op,
		core.Request[queryRequest]{Body: queryRequest{OutTradeNo: "T1"}})
	if err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
	if !executor.called {
		t.Fatal("executor not called")
	}
	if executor.req.Spec.Method != testMethod {
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
	op := newTestOperation(t)
	t.Run("执行器错误时保留元信息", func(t *testing.T) {
		executor := &mockExecutor{err: errors.New("boom"), result: &core.APIResult{HTTPStatus: 599}}
		resp, err := core.Execute(context.Background(), executor, op, core.Request[queryRequest]{})
		if err == nil || resp == nil || resp.Result == nil || resp.Result.HTTPStatus != 599 {
			t.Fatalf("resp = %+v, err = %v", resp, err)
		}
	})
	t.Run("未验证结果不得进入解码", func(t *testing.T) {
		executor := &mockExecutor{result: &core.APIResult{
			HTTPStatus: 200, SignatureVerified: false, Node: []byte(`{"code":"10000"}`),
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
			HTTPStatus: 200, SignatureVerified: true, Node: []byte(`{"code":10000}`), // code 类型错误
		}}
		_, err := core.Execute(context.Background(), executor, op, core.Request[queryRequest]{})
		if !errors.As(err, new(*core.ProtocolError)) {
			t.Fatalf("error = %v, want ProtocolError", err)
		}
	})
}

func TestClientExecuteMatchesPackageExecute(t *testing.T) {
	gateway := newGateway(t, func(url.Values) (int, string) {
		return http.StatusOK, signedResponse(t, testRespKey, successNode(`,"trade_no":"T1"`))
	})
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)
	op := newTestOperation(t)
	req := core.Request[queryRequest]{Body: queryRequest{OutTradeNo: "T1"}}

	viaMethod, err1 := client.Execute(context.Background(), op, req)
	viaPackage, err2 := core.Execute(context.Background(), client, op, req)
	if err1 != nil || err2 != nil {
		t.Fatalf("errors: %v / %v", err1, err2)
	}
	if viaMethod.Data.TradeNo != viaPackage.Data.TradeNo {
		t.Fatal("two entry points diverge")
	}
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
		gateway := newGateway(t, func(url.Values) (int, string) {
			return http.StatusOK, signedResponse(t, testRespKey, successNode(`,"trade_no":"T1"`))
		})
		defer gateway.Close()
		client := newTestClient(t, gateway.URL, option.WithLogHandler(handler))
		op := newTestOperation(t)
		if _, err := executeQuery(t, client, op); err != nil {
			t.Fatal(err)
		}
		r := handler.find(t, slog.LevelDebug, "alipay request succeeded")
		if r == nil {
			t.Fatal("success log record not found")
		}
		for key, want := range map[string]string{
			"sdk": "alipay", "operation": testMethod, "code": "10000",
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
	t.Run("业务拒绝为 Warn 且含 sub_code", func(t *testing.T) {
		handler := &captureHandler{}
		node := `{"code":"40004","msg":"Business Failed","sub_code":"ACQ.TRADE_NOT_EXIST"}`
		gateway := newGateway(t, func(url.Values) (int, string) {
			return http.StatusOK, signedResponse(t, testRespKey, node)
		})
		defer gateway.Close()
		client := newTestClient(t, gateway.URL, option.WithLogHandler(handler))
		op := newTestOperation(t)
		_, _ = executeQuery(t, client, op)
		r := handler.find(t, slog.LevelWarn, "alipay request failed")
		if r == nil {
			t.Fatal("warn record not found")
		}
		if v, _ := handler.attrValue(t, r, "sub_code"); v.String() != "ACQ.TRADE_NOT_EXIST" {
			t.Fatalf("sub_code attr = %v", v)
		}
	})
	t.Run("安全失败为 Error 且 error_kind 为 signature", func(t *testing.T) {
		handler := &captureHandler{}
		gateway := newGateway(t, func(url.Values) (int, string) {
			return http.StatusOK, `{"` + testRespKey + `":{"code":"10000","msg":"Success"},"sign":""}`
		})
		defer gateway.Close()
		client := newTestClient(t, gateway.URL, option.WithLogHandler(handler))
		op := newTestOperation(t)
		_, _ = executeQuery(t, client, op)
		r := handler.find(t, slog.LevelError, "alipay request failed")
		if r == nil {
			t.Fatal("error record not found")
		}
		if v, _ := handler.attrValue(t, r, "error_kind"); v.String() != "signature" {
			t.Fatalf("error_kind = %v", v)
		}
	})
	t.Run("未配置日志时静默", func(t *testing.T) {
		gateway := newGateway(t, func(url.Values) (int, string) {
			return http.StatusOK, signedResponse(t, testRespKey, successNode(""))
		})
		defer gateway.Close()
		client := newTestClient(t, gateway.URL) // 不注入日志
		op := newTestOperation(t)
		if _, err := executeQuery(t, client, op); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Handler 错误不改变 API 结果", func(t *testing.T) {
		handler := &captureHandler{errOn: map[string]bool{"alipay request succeeded": true}}
		gateway := newGateway(t, func(url.Values) (int, string) {
			return http.StatusOK, signedResponse(t, testRespKey, successNode(`,"trade_no":"T9"`))
		})
		defer gateway.Close()
		client := newTestClient(t, gateway.URL, option.WithLogHandler(handler))
		op := newTestOperation(t)
		resp, err := executeQuery(t, client, op)
		if err != nil || resp.Data.TradeNo != "T9" {
			t.Fatalf("handler failure must not affect result: %v, %+v", err, resp)
		}
	})
}
