package option_test

import (
	"testing"
	"time"

	"github.com/tiyee/go-alipay/core"
	"github.com/tiyee/go-alipay/core/option"
	"github.com/tiyee/go-alipay/internal/testutil"
)

// 测试密钥为仓库专用虚构材料。
const (
	validPrivatePEM = testutil.AppPrivateKeyPEM
	validPublicPEM  = testutil.AlipayPublicKeyPEM
)

func TestNewClientAppliesOptionsAndDefaults(t *testing.T) {
	client, err := core.NewClient(
		option.WithAppID("2021000000000001"),
		option.WithEnvironment(core.EnvironmentProduction),
		option.WithPrivateKey(validPrivatePEM),
		option.WithAlipayPublicKey(validPublicPEM),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	s := client.Settings()
	if s.AppID != "2021000000000001" {
		t.Errorf("AppID = %q", s.AppID)
	}
	if s.Gateway == nil || s.Gateway.String() != core.GatewayProduction {
		t.Errorf("Gateway = %v, want %s", s.Gateway, core.GatewayProduction)
	}
	if s.Signer == nil || s.Signer.Algorithm() != "RSA2" {
		t.Error("Signer not configured")
	}
	if s.Verifier == nil || s.Verifier.Algorithm() != "RSA2" {
		t.Error("Verifier not configured")
	}
	if s.RequestTimeout != core.DefaultRequestTimeout {
		t.Errorf("RequestTimeout = %v", s.RequestTimeout)
	}
	if s.MaxResponseBytes != core.DefaultMaxResponseBytes {
		t.Errorf("MaxResponseBytes = %v", s.MaxResponseBytes)
	}
	if s.Logger == nil {
		t.Error("Logger must default to a silent logger, not nil")
	}
}

func TestWithEnvironmentSandbox(t *testing.T) {
	client, err := core.NewClient(
		option.WithAppID("app"),
		option.WithEnvironment(core.EnvironmentSandbox),
		option.WithPrivateKey(validPrivatePEM),
		option.WithAlipayPublicKey(validPublicPEM),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	if got := client.Settings().Gateway.String(); got != core.GatewaySandbox {
		t.Fatalf("Gateway = %q, want %q", got, core.GatewaySandbox)
	}
}

func TestNewClientValidationFailures(t *testing.T) {
	tests := []struct {
		name      string
		opts      []core.ClientOption
		wantField string
	}{
		{
			name:      "缺少 APPID",
			opts:      []core.ClientOption{option.WithEnvironment(core.EnvironmentProduction), option.WithPrivateKey(validPrivatePEM), option.WithAlipayPublicKey(validPublicPEM)},
			wantField: "app_id",
		},
		{
			name:      "缺少网关",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithPrivateKey(validPrivatePEM), option.WithAlipayPublicKey(validPublicPEM)},
			wantField: "gateway",
		},
		{
			name:      "缺少私钥",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithEnvironment(core.EnvironmentProduction), option.WithAlipayPublicKey(validPublicPEM)},
			wantField: "signer",
		},
		{
			name:      "缺少验签公钥",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithEnvironment(core.EnvironmentProduction), option.WithPrivateKey(validPrivatePEM)},
			wantField: "verifier",
		},
		{
			name:      "未知环境",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithEnvironment("staging")},
			wantField: "environment",
		},
		{
			name:      "非法网关地址",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithGateway("://bad")},
			wantField: "gateway",
		},
		{
			name:      "非法网关协议",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithGateway("ftp://example.com")},
			wantField: "gateway",
		},
		{
			name:      "非法私钥",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithEnvironment(core.EnvironmentProduction), option.WithPrivateKey("not a pem"), option.WithAlipayPublicKey(validPublicPEM)},
			wantField: "",
		},
		{
			name:      "负超时",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithEnvironment(core.EnvironmentProduction), option.WithPrivateKey(validPrivatePEM), option.WithAlipayPublicKey(validPublicPEM), option.WithRequestTimeout(-time.Second)},
			wantField: "request_timeout",
		},
		{
			name:      "非法响应上限",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithEnvironment(core.EnvironmentProduction), option.WithPrivateKey(validPrivatePEM), option.WithAlipayPublicKey(validPublicPEM), option.WithMaxResponseBytes(0)},
			wantField: "max_response_bytes",
		},
		{
			name:      "显式 nil Logger",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithEnvironment(core.EnvironmentProduction), option.WithPrivateKey(validPrivatePEM), option.WithAlipayPublicKey(validPublicPEM), option.WithLogger(nil)},
			wantField: "logger",
		},
		{
			name:      "显式 nil Handler",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithEnvironment(core.EnvironmentProduction), option.WithPrivateKey(validPrivatePEM), option.WithAlipayPublicKey(validPublicPEM), option.WithLogHandler(nil)},
			wantField: "log_handler",
		},
		{
			name:      "nil HTTP Client",
			opts:      []core.ClientOption{option.WithAppID("app"), option.WithEnvironment(core.EnvironmentProduction), option.WithPrivateKey(validPrivatePEM), option.WithAlipayPublicKey(validPublicPEM), option.WithHTTPClient(nil)},
			wantField: "http_client",
		},
		{
			name:      "nil 选项",
			opts:      []core.ClientOption{nil},
			wantField: "option",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := core.NewClient(tt.opts...)
			if err == nil {
				t.Fatal("NewClient() expect error")
			}
			if client != nil {
				t.Fatal("client must be nil on error")
			}
			if tt.wantField != "" {
				cfgErr, ok := err.(*core.ConfigError)
				if !ok || cfgErr.Field != tt.wantField {
					t.Fatalf("error = %v (%T), want ConfigError field %q", err, err, tt.wantField)
				}
			}
		})
	}
}

func TestWithGatewayExplicit(t *testing.T) {
	client, err := core.NewClient(
		option.WithAppID("app"),
		option.WithGateway("http://127.0.0.1:8080/gateway.do"),
		option.WithPrivateKey(validPrivatePEM),
		option.WithAlipayPublicKey(validPublicPEM),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	if got := client.Settings().Gateway.String(); got != "http://127.0.0.1:8080/gateway.do" {
		t.Fatalf("Gateway = %q", got)
	}
}

func TestEnvironmentConstants(t *testing.T) {
	if core.EnvironmentProduction.String() != "production" ||
		core.EnvironmentSandbox.String() != "sandbox" {
		t.Fatal("environment string values changed")
	}
}
