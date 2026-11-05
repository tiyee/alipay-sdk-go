package notify_test

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"slices"
	"strings"
	"sync"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/notify"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// 测试目标：Handler 以 Credentials 为构造首参并绑定认证快照；
// 通知表单只解码一次，特殊字符（+、%2B、%26、中文、空值）在传输编码
// 与待验签串之间保持一致；缺签、错签、错公钥、重复参数与未支持策略
// 全部拒绝；日志不包含通知参数值与签名。

const (
	testAppID  = "2021000000000001"
	testNotify = "trade_status_sync"
)

// testNotification 是测试用的最小领域通知模型。
type testNotification struct {
	OutTradeNo  string
	TradeStatus string
	Subject     string
	Body        string
}

// testDecoder 映射少量字段，验证泛型 Decoder 契约。
var testDecoder = notify.DecoderFunc[testNotification](func(values url.Values) (testNotification, error) {
	return testNotification{
		OutTradeNo:  values.Get("out_trade_no"),
		TradeStatus: values.Get("trade_status"),
		Subject:     values.Get("subject"),
		Body:        values.Get("body"),
	}, nil
})

// newHandler 返回公钥模式的测试 Handler。
func newHandler(t *testing.T, opts ...notify.ParseOption) *notify.Handler {
	t.Helper()
	handler, err := notify.NewHandler(testutil.NewCredentials(t), notify.PolicyTradeStatusSync, opts...)
	if err != nil {
		t.Fatalf("NewHandler() error = %v", err)
	}
	return handler
}

// notifySignContent 独立构造通知待验签串（复用 internal/testutil 的
// 独立实现，与 internal/protocol、core/notify 互为交叉核对）：
// 剔除 sign 与 sign_type，保留空值参数，按参数名升序 k=v 用 & 连接。
func notifySignContent(form url.Values) string {
	return testutil.NotifySignContent(form)
}

// signContentDropEmpty 按“剔除空值”的请求签名规则构造待验签串，
// 仅用于验证通知验签不采用该规则。
func signContentDropEmpty(form url.Values) string {
	keys := make([]string, 0, len(form))
	for key, vs := range form {
		if key == "sign" || key == "sign_type" || len(vs) == 0 || vs[0] == "" {
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

// baseForm 返回满足普通支付通知策略的最小参数集。
func baseForm() url.Values {
	return url.Values{
		"notify_time":    {"2026-11-05 10:00:00"},
		"notify_type":    {testNotify},
		"notify_id":      {"ac05099524730693a8b330c5ecf72da9786"},
		"trade_no":       {"2026110522001400001234567890"},
		"app_id":         {testAppID},
		"auth_app_id":    {testAppID},
		"out_trade_no":   {"ORDER-20261105-001"},
		"trade_status":   {"TRADE_SUCCESS"},
		"total_amount":   {"88.88"},
		"receipt_amount": {"88.88"},
	}
}

// signedForm 以网关视角对待验签串签名并填充 sign 与 sign_type。
func signedForm(t *testing.T, form url.Values) url.Values {
	t.Helper()
	sign, err := testutil.SignWithAlipayKey([]byte(notifySignContent(form)))
	if err != nil {
		t.Fatalf("sign form: %v", err)
	}
	out := compat.CloneValues(form)
	out.Set("sign_type", "RSA2")
	out.Set("sign", sign)
	return out
}

func TestParseVerifiesSignedNotification(t *testing.T) {
	ctx := context.Background()
	source := baseForm()
	form := signedForm(t, source)

	got, err := notify.Parse(ctx, newHandler(t), form, testDecoder)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if got.Data == nil {
		t.Fatal("data is nil")
	}
	if got.Data.OutTradeNo != "ORDER-20261105-001" || got.Data.TradeStatus != "TRADE_SUCCESS" {
		t.Fatalf("unexpected data: %+v", *got.Data)
	}
	if _, ok := got.Params["sign"]; ok {
		t.Error("params should not contain sign")
	}
	if _, ok := got.Params["sign_type"]; ok {
		t.Error("params should not contain sign_type")
	}
	if got.Params.Get("out_trade_no") != "ORDER-20261105-001" {
		t.Errorf("params out_trade_no = %q", got.Params.Get("out_trade_no"))
	}
	// Params 必须是副本：修改返回值不得影响调用方输入。
	got.Params.Set("out_trade_no", "MUTATED")
	if source.Get("out_trade_no") != "ORDER-20261105-001" {
		t.Error("parse mutated caller values")
	}
}

func TestParseSingleDecodeSpecialCharacters(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name    string
		subject string
		body    string
	}{
		{name: "space as plus", subject: "测试 单据", body: "备注 A"},
		{name: "literal plus", subject: "a+b", body: "C++"},
		{name: "ampersand and equals", subject: "a&b=c", body: "x=1&y=2"},
		{name: "chinese utf-8", subject: "支付宝交易", body: "中文内容"},
		{name: "mixed", subject: "测试 a+b&c=d e", body: "x=1&y=2 单据+&="},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			form := baseForm()
			form.Set("subject", tc.subject)
			form.Set("body", tc.body)
			// 模拟传输：按表单编码发送，接收端只解码一次再验签。
			received, err := url.ParseQuery(form.Encode())
			if err != nil {
				t.Fatalf("parse query: %v", err)
			}
			signed := signedForm(t, received)

			got, err := notify.Parse(ctx, newHandler(t), signed, testDecoder)
			if err != nil {
				t.Fatalf("parse: %v", err)
			}
			if got.Data.Subject != tc.subject {
				t.Errorf("subject = %q, want %q", got.Data.Subject, tc.subject)
			}
			if got.Data.Body != tc.body {
				t.Errorf("body = %q, want %q", got.Data.Body, tc.body)
			}
		})
	}
}

// 空值参数参与待验签串：与请求签名剔除空值的规则不同。
// 若按请求规则把空值剔除后再构造待验签串，验签必须失败。
func TestParseEmptyValueStaysInSignContent(t *testing.T) {
	ctx := context.Background()

	form := baseForm()
	form.Set("body", "")
	signed := signedForm(t, form)
	if _, err := notify.Parse(ctx, newHandler(t), signed, testDecoder); err != nil {
		t.Fatalf("empty value should participate in sign content: %v", err)
	}

	// 以“剔除空值”的错误规则重签，验签必须失败。
	wrong := baseForm()
	wrong.Set("body", "")
	sign, err := testutil.SignWithAlipayKey([]byte(signContentDropEmpty(wrong)))
	if err != nil {
		t.Fatal(err)
	}
	wrong.Set("sign_type", "RSA2")
	wrong.Set("sign", sign)
	if _, err := notify.Parse(ctx, newHandler(t), wrong, testDecoder); err == nil {
		t.Fatal("sign content built without empty values must fail verification")
	}
}

func TestParseRejectsDuplicateParams(t *testing.T) {
	ctx := context.Background()
	form := signedForm(t, baseForm())
	form["out_trade_no"] = []string{"ORDER-A", "ORDER-B"}

	_, err := notify.Parse(ctx, newHandler(t), form, testDecoder)
	protoErr, ok := compat.AsType[*core.ProtocolError](err)
	if !ok || !strings.Contains(protoErr.Reason, "duplicate") {
		t.Fatalf("want duplicate param protocol error, got %v", err)
	}
}

func TestParseSignatureFailures(t *testing.T) {
	ctx := context.Background()
	cases := []struct {
		name   string
		mutate func(t *testing.T) url.Values
	}{
		{
			name: "missing sign",
			mutate: func(t *testing.T) url.Values {
				form := signedForm(t, baseForm())
				form.Del("sign")
				return form
			},
		},
		{
			name: "tampered value",
			mutate: func(t *testing.T) url.Values {
				form := signedForm(t, baseForm())
				form.Set("total_amount", "0.01")
				return form
			},
		},
		{
			name: "tampered sign",
			mutate: func(t *testing.T) url.Values {
				form := signedForm(t, baseForm())
				other := baseForm()
				other.Set("out_trade_no", "OTHER")
				otherSigned := signedForm(t, other)
				form.Set("sign", otherSigned.Get("sign"))
				return form
			},
		},
		{
			name: "malformed sign",
			mutate: func(t *testing.T) url.Values {
				form := signedForm(t, baseForm())
				form.Set("sign", "not-base64!!")
				return form
			},
		},
		{
			name: "unsupported sign type",
			mutate: func(t *testing.T) url.Values {
				form := signedForm(t, baseForm())
				form.Set("sign_type", "RSA1")
				return form
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := notify.Parse(ctx, newHandler(t), tc.mutate(t), testDecoder)
			sigErr, ok := compat.AsType[*core.SignatureError](err)
			if !ok {
				t.Fatalf("want *core.SignatureError, got %v", err)
			}
			if sigErr.Stage != "notify" {
				t.Errorf("stage = %q, want notify", sigErr.Stage)
			}
		})
	}

	t.Run("wrong public key", func(t *testing.T) {
		// 以应用公钥充当支付宝公钥：验签必然失败。
		alipayPublicKey := appPublicKey(t)
		wrongCredentials, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
			AppSigner:       testutil.NewAppCryptoSigner(t),
			AlipayPublicKey: alipayPublicKey,
		})
		if err != nil {
			t.Fatal(err)
		}
		wrongHandler, err := notify.NewHandler(wrongCredentials, notify.PolicyTradeStatusSync)
		if err != nil {
			t.Fatal(err)
		}
		_, err = notify.Parse(ctx, wrongHandler, signedForm(t, baseForm()), testDecoder)
		if _, ok := compat.AsType[*core.SignatureError](err); !ok {
			t.Fatalf("want signature error with wrong public key, got %v", err)
		}
	})

	t.Run("certificate mode via credentials verifier", func(t *testing.T) {
		// 证书模式复用同一 Handler：单张可信证书 + 通知契约无标识时，
		// 按官方 rsaCertCheckV1 行为使用该证书公钥验签。
		suite := testutil.NewCertSuite(t)
		certHandler, err := notify.NewHandler(suite.NewCertificateCredentials(t), notify.PolicyTradeStatusSync)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := notify.Parse(ctx, certHandler, signedForm(t, baseForm()), testDecoder); err != nil {
			t.Fatalf("certificate mode parse: %v", err)
		}
	})
}

func TestNewHandlerRejections(t *testing.T) {
	signed := signedForm(t, baseForm())

	t.Run("nil credentials", func(t *testing.T) {
		_, err := notify.NewHandler(nil, notify.PolicyTradeStatusSync)
		if _, ok := compat.AsType[*core.ValidationError](err); !ok {
			t.Fatalf("want validation error, got %v", err)
		}
	})
	t.Run("zero policy", func(t *testing.T) {
		_, err := notify.NewHandler(testutil.NewCredentials(t), notify.Policy(0))
		if _, ok := compat.AsType[*core.ValidationError](err); !ok {
			t.Fatalf("want validation error, got %v", err)
		}
	})
	t.Run("unimplemented policy value", func(t *testing.T) {
		_, err := notify.NewHandler(testutil.NewCredentials(t), notify.Policy(9))
		if _, ok := compat.AsType[*core.ValidationError](err); !ok {
			t.Fatalf("want validation error, got %v", err)
		}
	})
	t.Run("credentials without verifier", func(t *testing.T) {
		_, err := notify.NewHandler(brokenCredentials{}, notify.PolicyTradeStatusSync)
		if _, ok := compat.AsType[*core.ValidationError](err); !ok {
			t.Fatalf("want validation error, got %v", err)
		}
	})
	t.Run("nil option", func(t *testing.T) {
		_, err := notify.NewHandler(testutil.NewCredentials(t), notify.PolicyTradeStatusSync, nil)
		if _, ok := compat.AsType[*core.ValidationError](err); !ok {
			t.Fatalf("want validation error, got %v", err)
		}
	})
	t.Run("invalid max form bytes", func(t *testing.T) {
		_, err := notify.NewHandler(testutil.NewCredentials(t), notify.PolicyTradeStatusSync,
			notify.WithMaxFormBytes(0))
		if _, ok := compat.AsType[*core.ValidationError](err); !ok {
			t.Fatalf("want validation error, got %v", err)
		}
	})

	// 以下为 Parse 调用期校验。
	handler := newHandler(t)
	t.Run("nil decoder", func(t *testing.T) {
		_, err := notify.Parse[testNotification](context.Background(), handler, signed, nil)
		if _, ok := compat.AsType[*core.ValidationError](err); !ok {
			t.Fatalf("want validation error, got %v", err)
		}
	})
	t.Run("nil ctx", func(t *testing.T) {
		// 本用例专门验证 Parse 对 nil ctx 的显式拒绝，属于刻意的越界输入。
		//lint:ignore SA1012 testing the nil-context guard on purpose
		_, err := notify.Parse(nil, handler, signed, testDecoder)
		if _, ok := compat.AsType[*core.ValidationError](err); !ok {
			t.Fatalf("want validation error, got %v", err)
		}
	})
	t.Run("other notify type", func(t *testing.T) {
		form := compat.CloneValues(signed)
		form.Set("notify_type", "other_notify")
		if _, err := notify.Parse(context.Background(), handler, form, testDecoder); err == nil {
			t.Fatal("mismatched notify type must fail")
		}
	})
	t.Run("missing notify type", func(t *testing.T) {
		form := compat.CloneValues(signed)
		form.Del("notify_type")
		if _, err := notify.Parse(context.Background(), handler, form, testDecoder); err == nil {
			t.Fatal("missing notify type must fail")
		}
	})
	t.Run("unsupported charset", func(t *testing.T) {
		form := compat.CloneValues(signed)
		form.Set("charset", "gbk")
		if _, err := notify.Parse(context.Background(), handler, form, testDecoder); err == nil {
			t.Fatal("unsupported charset must fail")
		}
	})
}

// brokenCredentials 的 Verifier 返回 nil，用于构造期拒绝路径。
type brokenCredentials struct{}

func (brokenCredentials) Mode() auth.Mode         { return auth.ModePublicKey }
func (brokenCredentials) Signer() auth.Signer     { return nil }
func (brokenCredentials) Verifier() auth.Verifier { return nil }
func (brokenCredentials) Validate() error         { return nil }

// appPublicKey 返回解析后的测试应用公钥（作为错误的支付宝公钥使用）。
func appPublicKey(t *testing.T) *rsa.PublicKey {
	t.Helper()
	key, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(testutil.AppPublicKeyPEM))
	if err != nil {
		t.Fatalf("parse app public key: %v", err)
	}
	return key
}

// 领域解码器失败时错误被包装为 ProtocolError，但通过原因链仍可匹配。
func TestParseWrapsDecoderError(t *testing.T) {
	ctx := context.Background()
	decoder := notify.DecoderFunc[testNotification](func(url.Values) (testNotification, error) {
		return testNotification{}, &core.ValidationError{Field: "out_trade_no",
			Reason: "required notify param is missing"}
	})
	_, err := notify.Parse(ctx, newHandler(t), signedForm(t, baseForm()), decoder)
	if _, ok := compat.AsType[*core.ProtocolError](err); !ok {
		t.Fatalf("want protocol error, got %v", err)
	}
	valErr, ok := compat.AsType[*core.ValidationError](err)
	if !ok || valErr.Field != "out_trade_no" {
		t.Fatalf("want wrapped validation error, got %v", err)
	}
}

func TestParseRequest(t *testing.T) {
	ctx := context.Background()

	post := func(t *testing.T, body string) *http.Request {
		t.Helper()
		req := httptest.NewRequest(http.MethodPost, "/alipay/notify", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		return req
	}

	t.Run("verified", func(t *testing.T) {
		raw := signedForm(t, baseForm()).Encode()
		got, err := notify.ParseRequest(ctx, newHandler(t), post(t, raw), testDecoder)
		if err != nil {
			t.Fatalf("parse request: %v", err)
		}
		if got.Data.OutTradeNo != "ORDER-20261105-001" {
			t.Errorf("out_trade_no = %q", got.Data.OutTradeNo)
		}
	})

	t.Run("rejects GET", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/alipay/notify?x=1", nil)
		_, err := notify.ParseRequest(ctx, newHandler(t), req, testDecoder)
		if _, ok := compat.AsType[*core.ProtocolError](err); !ok {
			t.Fatalf("want protocol error, got %v", err)
		}
	})

	t.Run("size limit", func(t *testing.T) {
		raw := signedForm(t, baseForm()).Encode()
		handler, err := notify.NewHandler(testutil.NewCredentials(t),
			notify.PolicyTradeStatusSync, notify.WithMaxFormBytes(16))
		if err != nil {
			t.Fatal(err)
		}
		_, err = notify.ParseRequest(ctx, handler, post(t, raw), testDecoder)
		if _, ok := compat.AsType[*core.ProtocolError](err); !ok {
			t.Fatalf("want protocol error, got %v", err)
		}
	})

	t.Run("malformed form", func(t *testing.T) {
		_, err := notify.ParseRequest(ctx, newHandler(t), post(t, "a=%zz&b=c"), testDecoder)
		if _, ok := compat.AsType[*core.ProtocolError](err); !ok {
			t.Fatalf("want protocol error, got %v", err)
		}
	})

	t.Run("read failure", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/alipay/notify", nil)
		req.Body = io.NopCloser(io.MultiReader(strings.NewReader("a=1"), errReader{}))
		_, err := notify.ParseRequest(ctx, newHandler(t), req, testDecoder)
		if _, ok := compat.AsType[*core.TransportError](err); !ok {
			t.Fatalf("want transport error, got %v", err)
		}
	})

	t.Run("nil request", func(t *testing.T) {
		_, err := notify.ParseRequest(ctx, newHandler(t), nil, testDecoder)
		if _, ok := compat.AsType[*core.ValidationError](err); !ok {
			t.Fatalf("want validation error, got %v", err)
		}
	})
}

type errReader struct{}

func (errReader) Read([]byte) (int, error) { return 0, errors.New("boom") }

// captureHandler 记录日志记录，用于断言级别、字段与脱敏。
type captureHandler struct {
	mu      sync.Mutex
	records []slog.Record
}

func (h *captureHandler) Enabled(context.Context, slog.Level) bool { return true }

func (h *captureHandler) Handle(_ context.Context, r slog.Record) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.records = append(h.records, r.Clone())
	return nil
}

func (h *captureHandler) WithAttrs([]slog.Attr) slog.Handler { return h }
func (h *captureHandler) WithGroup(string) slog.Handler      { return h }

func (h *captureHandler) text() string {
	h.mu.Lock()
	defer h.mu.Unlock()
	var sb strings.Builder
	for _, r := range h.records {
		sb.WriteString(r.Message)
		r.Attrs(func(a slog.Attr) bool {
			fmt.Fprintf(&sb, " %s=%v", a.Key, a.Value)
			return true
		})
		sb.WriteString("\n")
	}
	return sb.String()
}

func TestParseLogsFilteredFieldsOnly(t *testing.T) {
	ctx := context.Background()

	t.Run("success logs debug without values", func(t *testing.T) {
		capture := &captureHandler{}
		form := baseForm()
		form.Set("subject", "机密标题-不应出现在日志")
		if _, err := notify.Parse(ctx, newHandler(t, notify.WithLogger(slog.New(capture))),
			signedForm(t, form), testDecoder); err != nil {
			t.Fatalf("parse: %v", err)
		}
		text := capture.text()
		if want := "alipay notify verified sdk=alipay operation=notify\n"; text != want {
			t.Errorf("success log = %q, want %q", text, want)
		}
	})

	t.Run("failure logs error with error kind", func(t *testing.T) {
		capture := &captureHandler{}
		form := signedForm(t, baseForm())
		form.Set("total_amount", "0.01")
		_, err := notify.Parse(ctx, newHandler(t, notify.WithLogger(slog.New(capture))), form, testDecoder)
		if err == nil {
			t.Fatal("want failure")
		}
		text := capture.text()
		if want := "alipay notify failed sdk=alipay operation=notify error_kind=signature\n"; text != want {
			t.Errorf("failure log = %q, want %q", text, want)
		}
	})

	t.Run("nil logger option rejected", func(t *testing.T) {
		_, err := notify.NewHandler(testutil.NewCredentials(t), notify.PolicyTradeStatusSync,
			notify.WithLogger(nil))
		if _, ok := compat.AsType[*core.ValidationError](err); !ok {
			t.Fatalf("want validation error, got %v", err)
		}
	})
}

// 共享 Handler 的并发解析必须安全且互不干扰。
func TestParseConcurrent(t *testing.T) {
	ctx := context.Background()
	handler := newHandler(t)
	const workers = 16
	var wg sync.WaitGroup
	errs := make([]error, workers)
	wg.Add(workers)
	for i := range workers {
		go func() {
			defer wg.Done()
			form := baseForm()
			form.Set("out_trade_no", fmt.Sprintf("ORDER-%02d", i))
			got, err := notify.Parse(ctx, handler, signedForm(t, form), testDecoder)
			switch {
			case err != nil:
				errs[i] = err
			case got.Data == nil || got.Data.OutTradeNo != fmt.Sprintf("ORDER-%02d", i):
				errs[i] = fmt.Errorf("worker %d got %+v", i, *got.Data)
			}
		}()
	}
	wg.Wait()
	for i, err := range errs {
		if err != nil {
			t.Fatalf("worker %d: %v", i, err)
		}
	}
}

func TestPolicyString(t *testing.T) {
	if got := notify.PolicyTradeStatusSync.String(); got != testNotify {
		t.Errorf("policy string = %q, want %q", got, testNotify)
	}
	if got := notify.Policy(9).String(); got != "Policy(9)" {
		t.Errorf("unknown policy string = %q", got)
	}
}

func TestDefaultMaxFormBytes(t *testing.T) {
	if notify.DefaultMaxFormBytes <= 0 || notify.DefaultMaxFormBytes > 1<<20 {
		t.Fatalf("unexpected default max form bytes: %d", notify.DefaultMaxFormBytes)
	}
}
