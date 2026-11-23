package core_test

import (
	"context"
	"errors"
	"log/slog"
	"math"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

func TestNewClientRejectsInvalidFinalConfiguration(t *testing.T) {
	var nilTransport *http.Transport
	var nilJar *cookiejar.Jar
	var nilHandler *slog.JSONHandler
	credentials := testutil.NewCredentials(t)
	for _, test := range []struct {
		name  string
		opt   core.ClientOption
		field string
	}{
		{"app id whitespace", option.WithAppID(" app "), "app_id"},
		{"app id newline", option.WithAppID("app\n"), "app_id"},
		{"app id auth delimiter", option.WithAppID("app,nonce=other"), "app_id"},
		{"app id form delimiter", option.WithAppID("app&other=value"), "app_id"},
		{"custom negative limit", func(s *core.Settings) error { s.MaxResponseBytes = -1; return nil }, "max_response_bytes"},
		{"overflowing limit", option.WithMaxResponseBytes(math.MaxInt64), "max_response_bytes"},
		{"custom overflowing limit", func(s *core.Settings) error { s.MaxResponseBytes = math.MaxInt64; return nil }, "max_response_bytes"},
		{"custom negative timeout", func(s *core.Settings) error { s.RequestTimeout = -time.Second; return nil }, "request_timeout"},
		{"custom auth override", func(s *core.Settings) error { s.Mode = auth.ModeCertificate; return nil }, "auth_mode"},
		{"http timeout", option.WithHTTPClient(&http.Client{Timeout: -time.Second}), "http_client"},
		{"typed nil transport", option.WithHTTPClient(&http.Client{Transport: nilTransport}), "http_client"},
		{"typed nil cookie jar", option.WithHTTPClient(&http.Client{Jar: nilJar}), "http_client"},
		{"zero logger", option.WithLogger(new(slog.Logger)), "logger"},
		{"typed nil log handler", option.WithLogHandler(nilHandler), "logger"},
		{"host missing", option.WithGateway("https://:443"), "gateway"},
		{"port out of range", option.WithGateway("https://gateway.example:65536"), "gateway"},
		{"port zero", option.WithGateway("https://gateway.example:0"), "gateway"},
		{"port empty", option.WithGateway("https://gateway.example:"), "gateway"},
		{"custom malformed host", func(s *core.Settings) error {
			s.Gateway = &url.URL{Scheme: "https", Host: "invalid host"}
			return nil
		}, "gateway"},
		{"custom malformed port", func(s *core.Settings) error {
			s.Gateway = &url.URL{Scheme: "https", Host: "gateway.example:bad"}
			return nil
		}, "gateway"},
		{"artifact credentials", option.WithArtifactGateway("https://user@gateway.example/gateway.do"), "artifact_gateway"},
		{"artifact query", option.WithArtifactGateway("https://gateway.example/gateway.do?unexpected=value"), "artifact_gateway"},
		{"artifact fragment", option.WithArtifactGateway("https://gateway.example/gateway.do#fragment"), "artifact_gateway"},
	} {
		t.Run(test.name, func(t *testing.T) {
			client, err := core.NewClient(credentials,
				option.WithAppID(testAppID), option.WithEnvironment(core.EnvironmentSandbox), test.opt)
			var configErr *core.ConfigError
			if client != nil || !errors.As(err, &configErr) || configErr.Field != test.field {
				t.Fatalf("NewClient() = %v, %v; want nil and ConfigError(%s)", client, err, test.field)
			}
		})
	}
}

func TestNewClientRejectsNilDefaultTransport(t *testing.T) {
	original := http.DefaultTransport
	t.Cleanup(func() { http.DefaultTransport = original })
	var typedNil *http.Transport
	for _, transport := range []http.RoundTripper{nil, typedNil} {
		http.DefaultTransport = transport
		client, err := core.NewClient(testutil.NewCredentials(t), option.WithAppID(testAppID), option.WithEnvironment(core.EnvironmentSandbox))
		var configErr *core.ConfigError
		if client != nil || !errors.As(err, &configErr) || configErr.Field != "http_client" {
			t.Fatalf("NewClient() = %v, %v; want nil and ConfigError(http_client)", client, err)
		}
	}
}

func TestNewClientRejectsMalformedCertificateIdentity(t *testing.T) {
	credentials := testutil.NewCredentials(t)
	for _, identity := range []auth.CertificateIdentity{
		{AppCertificateSN: "sn,nonce=other", RootCertificateSN: "root"},
		{AppCertificateSN: "sn", RootCertificateSN: "root\n"},
	} {
		client, err := core.NewClient(&staticCredentials{
			mode: auth.ModeCertificate, signer: credentials.Signer(), verifier: credentials.Verifier(), identity: &identity,
		}, option.WithAppID(testAppID), option.WithEnvironment(core.EnvironmentSandbox))
		var configErr *core.ConfigError
		if client != nil || !errors.As(err, &configErr) || configErr.Field != "certificate_identity" {
			t.Fatalf("NewClient() = %v, %v; want nil and ConfigError(certificate_identity)", client, err)
		}
	}
}

func TestNewClientPreservesDocumentedDefaults(t *testing.T) {
	client, err := core.NewClient(testutil.NewCredentials(t),
		option.WithAppID(testAppID), option.WithEnvironment(core.EnvironmentSandbox),
		option.WithHTTPClient(&http.Client{}), option.WithRequestTimeout(0),
		func(s *core.Settings) error { s.MaxResponseBytes = 0; return nil })
	if err != nil {
		t.Fatal(err)
	}
	settings := client.Settings()
	if settings.RequestTimeout != core.DefaultRequestTimeout || settings.MaxResponseBytes != core.DefaultMaxResponseBytes {
		t.Fatalf("documented defaults not applied: %+v", settings)
	}
}

func TestNewClientSnapshotsDefaultTransport(t *testing.T) {
	original := http.DefaultTransport
	t.Cleanup(func() { http.DefaultTransport = original })
	calls := 0
	http.DefaultTransport = contextTransport(func(*http.Request) (*http.Response, error) {
		calls++
		return nil, context.Canceled
	})
	supplied := &http.Client{}
	client := newTestClient(t, "https://gateway.example", option.WithHTTPClient(supplied))
	http.DefaultTransport = contextTransport(func(*http.Request) (*http.Response, error) {
		t.Error("client used an unvalidated replacement transport")
		return nil, context.Canceled
	})
	_, err := core.Execute(t.Context(), client, newQueryOperation(t), core.Request[queryRequest]{})
	if !errors.Is(err, context.Canceled) || calls != 1 {
		t.Fatalf("Execute() = %v, calls=%d; want the transport bound at construction", err, calls)
	}
	if supplied.Transport != nil {
		t.Fatal("constructor mutated caller's HTTP client")
	}
}

func TestArtifactBuilderRejectsInvalidConfiguration(t *testing.T) {
	credentials := testutil.NewCredentials(t)
	for _, test := range []struct {
		opt   core.ClientOption
		field string
	}{
		{option.WithArtifactGateway("https://gateway.example/gateway.do?x=1"), "artifact_gateway"},
		{option.WithArtifactGateway("https://:443/gateway.do"), "artifact_gateway"},
		{option.WithGateway("https://gateway.example:65536"), "gateway"},
		{func(s *core.Settings) error { s.MaxResponseBytes = -1; return nil }, "max_response_bytes"},
		{func(s *core.Settings) error { s.Mode = auth.ModePublicKey; return nil }, "auth_mode"},
	} {
		builder, err := core.NewArtifactBuilder(credentials,
			option.WithAppID(testAppID), option.WithEnvironment(core.EnvironmentSandbox), test.opt)
		var configErr *core.ConfigError
		if builder != nil || !errors.As(err, &configErr) || configErr.Field != test.field {
			t.Fatalf("NewArtifactBuilder() = %v, %v; want nil and ConfigError(%s)", builder, err, test.field)
		}
	}
}

func TestConstructorsSnapshotValidatedGateways(t *testing.T) {
	credentials := testutil.NewCredentials(t)
	rest := &url.URL{Scheme: "https", Host: "gateway.example"}
	artifact := &url.URL{Scheme: "https", Host: "gateway.example", Path: "/gateway.do"}
	opts := []core.ClientOption{option.WithAppID(testAppID), func(s *core.Settings) error {
		s.Gateway, s.ArtifactGateway = rest, artifact
		return nil
	}}
	client, err := core.NewClient(credentials, opts...)
	if err != nil {
		t.Fatal(err)
	}
	builder, err := core.NewArtifactBuilder(credentials, opts...)
	if err != nil {
		t.Fatal(err)
	}
	rest.Host, artifact.Host = "", ""
	for _, settings := range []core.Settings{client.Settings(), builder.Settings()} {
		if settings.Gateway.Host != "gateway.example" || settings.ArtifactGateway.Host != "gateway.example" {
			t.Fatal("caller mutated a validated gateway after construction")
		}
	}
	product, err := builder.BuildArtifact(t.Context(), "alipay.trade.page.pay", url.Values{"biz_content": {"{}"}}, core.CallSettings{})
	if err != nil || product.Gateway != "https://gateway.example/gateway.do" {
		t.Fatalf("BuildArtifact() = %+v, %v", product, err)
	}
}
