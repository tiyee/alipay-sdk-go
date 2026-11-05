package core_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
	"github.com/tiyee/alipay-sdk-go/internal/testutil"
)

// 测试目标：Credentials 作为 NewClient 第一个必选参数的构造边界——
// nil 与 typed nil、校验失败、未知模式、nil 算法对象、算法不一致、
// 模式与证书能力矛盾均在构造阶段失败；外部包可以实现自定义
// Credentials；AuthRequirement 在统一 Do 路径校验。
//
// 本文件所有网关地址均为虚构的 V3 基地址，仅用于构造期校验，
// 不发起网络请求。

// staticCredentials 是测试用最小 Credentials 实现（外部实现可行性验证）。
type staticCredentials struct {
	mode     auth.Mode
	signer   auth.Signer
	verifier auth.Verifier
	identity *auth.CertificateIdentity
}

func (c *staticCredentials) Mode() auth.Mode         { return c.mode }
func (c *staticCredentials) Signer() auth.Signer     { return c.signer }
func (c *staticCredentials) Verifier() auth.Verifier { return c.verifier }
func (c *staticCredentials) Validate() error         { return nil }

func (c *staticCredentials) CertificateIdentity() auth.CertificateIdentity {
	if c.identity == nil {
		return auth.CertificateIdentity{}
	}
	return *c.identity
}

func TestNewClientCredentialsBoundaries(t *testing.T) {
	credentials := testutil.NewCredentials(t)

	t.Run("nil credentials", func(t *testing.T) {
		_, err := core.NewClient(nil, option.WithAppID("app"),
			option.WithGateway("https://gateway.example.com"))
		var cfgErr *core.ConfigError
		if !errors.As(err, &cfgErr) || cfgErr.Field != "credentials" {
			t.Fatalf("error = %v, want ConfigError(credentials)", err)
		}
	})

	t.Run("typed nil credentials", func(t *testing.T) {
		var typedNil *testutilCredentials
		_, err := core.NewClient(typedNil)
		var cfgErr *core.ConfigError
		if !errors.As(err, &cfgErr) || cfgErr.Field != "credentials" {
			t.Fatalf("error = %v, want ConfigError(credentials)", err)
		}
	})

	t.Run("validation failure", func(t *testing.T) {
		_, err := core.NewClient(failingCredentials{})
		var cfgErr *core.ConfigError
		if !errors.As(err, &cfgErr) || cfgErr.Field != "credentials" {
			t.Fatalf("error = %v, want ConfigError(credentials)", err)
		}
		if !errors.Is(err, errInvalidCredentials) {
			t.Fatalf("error should keep the underlying cause: %v", err)
		}
	})

	t.Run("unknown mode", func(t *testing.T) {
		_, err := core.NewClient(&staticCredentials{mode: auth.ModeUnset,
			signer: credentials.Signer(), verifier: credentials.Verifier()})
		var cfgErr *core.ConfigError
		if !errors.As(err, &cfgErr) || cfgErr.Field != "auth_mode" {
			t.Fatalf("error = %v, want ConfigError(auth_mode)", err)
		}
	})

	t.Run("nil signer getter", func(t *testing.T) {
		_, err := core.NewClient(&staticCredentials{mode: auth.ModePublicKey,
			signer: nil, verifier: credentials.Verifier()})
		var cfgErr *core.ConfigError
		if !errors.As(err, &cfgErr) || cfgErr.Field != "signer" {
			t.Fatalf("error = %v, want ConfigError(signer)", err)
		}
	})

	t.Run("nil verifier getter", func(t *testing.T) {
		_, err := core.NewClient(&staticCredentials{mode: auth.ModePublicKey,
			signer: credentials.Signer(), verifier: nil})
		var cfgErr *core.ConfigError
		if !errors.As(err, &cfgErr) || cfgErr.Field != "verifier" {
			t.Fatalf("error = %v, want ConfigError(verifier)", err)
		}
	})

	t.Run("non-rsa2 signer algorithm", func(t *testing.T) {
		_, err := core.NewClient(&staticCredentials{mode: auth.ModePublicKey,
			signer: rsa1Signer{}, verifier: credentials.Verifier()})
		var cfgErr *core.ConfigError
		if !errors.As(err, &cfgErr) || cfgErr.Field != "signer" {
			t.Fatalf("error = %v, want ConfigError(signer)", err)
		}
	})

	t.Run("certificate mode without identity capability", func(t *testing.T) {
		// 仅实现 Credentials、不实现 CertificateProvider 的凭证。
		_, err := core.NewClient(&noIdentityCredentials{mode: auth.ModeCertificate,
			signer: credentials.Signer(), verifier: credentials.Verifier()})
		var cfgErr *core.ConfigError
		if !errors.As(err, &cfgErr) || cfgErr.Field != "credentials" {
			t.Fatalf("error = %v, want ConfigError(credentials)", err)
		}
	})

	t.Run("certificate mode with incomplete identity", func(t *testing.T) {
		_, err := core.NewClient(&staticCredentials{mode: auth.ModeCertificate,
			signer: credentials.Signer(), verifier: credentials.Verifier(),
			identity: &auth.CertificateIdentity{AppCertificateSN: "sn"}})
		var cfgErr *core.ConfigError
		if !errors.As(err, &cfgErr) || cfgErr.Field != "certificate_identity" {
			t.Fatalf("error = %v, want ConfigError(certificate_identity)", err)
		}
	})
}

// noIdentityCredentials 只实现 Credentials，用于模式/能力矛盾用例。
type noIdentityCredentials struct {
	mode     auth.Mode
	signer   auth.Signer
	verifier auth.Verifier
}

func (c *noIdentityCredentials) Mode() auth.Mode         { return c.mode }
func (c *noIdentityCredentials) Signer() auth.Signer     { return c.signer }
func (c *noIdentityCredentials) Verifier() auth.Verifier { return c.verifier }
func (c *noIdentityCredentials) Validate() error         { return nil }

// testutilCredentials 仅用于 typed nil 检测。
type testutilCredentials struct{}

func (*testutilCredentials) Mode() auth.Mode         { return auth.ModePublicKey }
func (*testutilCredentials) Signer() auth.Signer     { return nil }
func (*testutilCredentials) Verifier() auth.Verifier { return nil }
func (*testutilCredentials) Validate() error         { return nil }

var errInvalidCredentials = errors.New("broken credentials")

type failingCredentials struct{}

func (failingCredentials) Mode() auth.Mode         { return auth.ModePublicKey }
func (failingCredentials) Signer() auth.Signer     { return nil }
func (failingCredentials) Verifier() auth.Verifier { return nil }
func (failingCredentials) Validate() error         { return errInvalidCredentials }

type rsa1Signer struct{}

func (rsa1Signer) Algorithm() auth.SignatureAlgorithm { return "RSA1" }
func (rsa1Signer) Sign(context.Context, []byte) (auth.Signature, error) {
	return auth.Signature{}, nil
}

func TestCustomCredentialsImplementContract(t *testing.T) {
	credentials := testutil.NewCredentials(t)
	client, err := core.NewClient(&staticCredentials{mode: auth.ModePublicKey,
		signer: credentials.Signer(), verifier: credentials.Verifier()},
		option.WithAppID(testAppID),
		option.WithGateway("https://gateway.example.com"),
	)
	if err != nil {
		t.Fatalf("NewClient() with custom credentials error = %v", err)
	}
	if client.Mode() != auth.ModePublicKey {
		t.Fatalf("Mode() = %v", client.Mode())
	}
}

func TestClientSettingsSnapshotIsReadOnly(t *testing.T) {
	credentials := testutil.NewCredentials(t)
	client, err := core.NewClient(credentials,
		option.WithAppID(testAppID),
		option.WithGateway("https://gateway.example.com"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	s := client.Settings()
	// 快照不含签名器、验签器或凭证；模式为只读派生信息。
	if s.Mode != auth.ModePublicKey {
		t.Fatalf("Mode = %v", s.Mode)
	}
	// 修改快照不影响 Client。
	before := client.Settings().Gateway.String()
	s.Gateway.Host = "mutated.example.com"
	s.ArtifactGateway = &url.URL{Host: "mutated.example.com"}
	if after := client.Settings().Gateway.String(); after != before {
		t.Fatalf("Settings() write-back detected: %q -> %q", before, after)
	}
}

func TestAuthRequirementCheckedOnUnifiedDoPath(t *testing.T) {
	// 公钥 Client 调用声明证书要求的 Operation：请求不发送。
	var called bool
	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	}))
	defer gateway.Close()
	client := newTestClient(t, gateway.URL)

	// 声明证书要求的 Operation。
	certOnlyOp, err := core.NewJSONOperation[queryRequest, queryResponse](core.OperationSpec{
		Method:          "alipay.trade.query",
		HTTPMethod:      core.HTTPMethodPost,
		Path:            testV3Path,
		ParameterMode:   core.ParameterModeJSONBody,
		AuthRequirement: core.AuthRequirementCertificate,
	})
	if err != nil {
		t.Fatalf("NewJSONOperation() error = %v", err)
	}
	_, err = core.Execute(context.Background(), client, certOnlyOp, core.Request[queryRequest]{
		Body: queryRequest{OutTradeNo: "20150320010101001"},
	})
	valErr, ok := compat.AsType[*core.ValidationError](err)
	if !ok || valErr.Field != "auth_requirement" {
		t.Fatalf("error = %v, want ValidationError(auth_requirement)", err)
	}
	if called {
		t.Fatal("request must not be sent on auth requirement mismatch")
	}
}
