//go:build go1.25

package core_test

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"testing/synctest"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/notify"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

func requireExpiredCertificate(t *testing.T, err error) {
	t.Helper()
	var certErr *core.CertificateError
	var validityErr *auth.CertValidityError
	if !errors.As(err, &certErr) || !errors.As(err, &validityErr) {
		t.Fatalf("error = %v, want CertificateError with CertValidityError cause", err)
	}
}

func TestExpiredCredentialsRejectEveryEntryPoint(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		suite := testutil.NewCertSuite(t)
		credentials := suite.NewCertificateCredentials(t)
		calls := 0
		httpClient := &http.Client{Transport: contextTransport(func(*http.Request) (*http.Response, error) {
			calls++
			return nil, errors.New("unexpected HTTP request")
		})}
		client, err := core.NewClient(credentials, option.WithAppID(testAppID), option.WithEnvironment(core.EnvironmentSandbox), option.WithHTTPClient(httpClient))
		if err != nil {
			t.Fatal(err)
		}
		builder, err := core.NewArtifactBuilder(credentials, option.WithAppID(testAppID), option.WithEnvironment(core.EnvironmentSandbox))
		if err != nil {
			t.Fatal(err)
		}
		handler, err := notify.NewHandler(credentials, notify.PolicyTradeStatusSync)
		if err != nil {
			t.Fatal(err)
		}
		form := url.Values{"notify_type": {"trade_status_sync"}, "out_trade_no": {"order"}}
		signature, err := testutil.SignWithAlipayKey([]byte(testutil.NotifySignContent(form)))
		if err != nil {
			t.Fatal(err)
		}
		form.Set("sign", signature)
		form.Set("sign_type", "RSA2")
		decoded := false
		decoder := notify.DecoderFunc[string](func(values url.Values) (string, error) {
			decoded = true
			return values.Get("out_trade_no"), nil
		})
		if _, err := notify.Parse(t.Context(), handler, form, decoder); err != nil {
			t.Fatal(err)
		}
		decoded = false
		time.Sleep(time.Until(suite.AlipayCert.NotAfter) + time.Second)
		var validityErr *auth.CertValidityError
		if err := credentials.Validate(); !errors.As(err, &validityErr) {
			t.Fatalf("Validate() = %v, want CertValidityError", err)
		}
		// 凭证构造时有效，Client 初始化时已失效，也必须拒绝创建实例。
		expiredClient, err := core.NewClient(credentials, option.WithAppID(testAppID), option.WithEnvironment(core.EnvironmentSandbox))
		var configErr *core.ConfigError
		if expiredClient != nil || !errors.As(err, &configErr) || !errors.As(err, &validityErr) {
			t.Fatalf("NewClient() = %v, %v; want nil and certificate configuration error", expiredClient, err)
		}
		response, err := core.Execute(t.Context(), client, newQueryOperation(t), core.Request[queryRequest]{Body: queryRequest{OutTradeNo: "order"}})
		requireExpiredCertificate(t, err)
		if calls != 0 || response == nil || response.Data != nil || response.Result != nil {
			t.Fatalf("expired credentials sent a request or produced data: calls=%d response=%+v", calls, response)
		}
		artifact, err := builder.BuildArtifact(t.Context(), "alipay.trade.page.pay", url.Values{"biz_content": {"{}"}}, core.CallSettings{})
		requireExpiredCertificate(t, err)
		if artifact != nil {
			t.Fatal("expired credentials produced an artifact")
		}
		result, err := notify.Parse(t.Context(), handler, form, decoder)
		requireExpiredCertificate(t, err)
		if decoded || result.Data != nil || result.Params != nil {
			t.Fatal("expired credentials produced notification data")
		}
		request, err := http.NewRequest(http.MethodPost, "https://merchant.example/notify", strings.NewReader(form.Encode()))
		if err != nil {
			t.Fatal(err)
		}
		result, err = notify.ParseRequest(t.Context(), handler, request, decoder)
		requireExpiredCertificate(t, err)
		if decoded || result.Data != nil || result.Params != nil {
			t.Fatal("expired credentials decoded the HTTP notification")
		}
	})
}

func TestCertificateExpiresWhileWaitingForResponse(t *testing.T) {
	for _, status := range []int{http.StatusOK, http.StatusBadRequest} {
		synctest.Test(t, func(t *testing.T) {
			suite := testutil.NewCertSuite(t)
			credentials := suite.NewCertificateCredentials(t)
			calls := 0
			httpClient := &http.Client{Transport: contextTransport(func(*http.Request) (*http.Response, error) {
				calls++
				time.Sleep(time.Until(suite.AlipayCert.NotAfter) + time.Second)
				body := `{"code":"invalid-parameter","message":"test","trade_no":"order"}`
				headers := testutil.V3ResponseHeaders(t, []byte(body), "1764000000012", "response-nonce", testutil.CertSN(t, suite.AlipayCert))
				return &http.Response{StatusCode: status, Header: headers, Body: io.NopCloser(strings.NewReader(body))}, nil
			})}
			client, err := core.NewClient(credentials, option.WithAppID(testAppID), option.WithEnvironment(core.EnvironmentSandbox), option.WithHTTPClient(httpClient), option.WithRequestTimeout(48*time.Hour))
			if err != nil {
				t.Fatal(err)
			}
			response, err := core.Execute(t.Context(), client, newQueryOperation(t), core.Request[queryRequest]{Body: queryRequest{OutTradeNo: "order"}})
			requireExpiredCertificate(t, err)
			if calls != 1 || response == nil || response.Data != nil || response.Result == nil || response.Result.SignatureVerified {
				t.Fatalf("unexpected response: calls=%d response=%+v", calls, response)
			}
		})
	}
}
