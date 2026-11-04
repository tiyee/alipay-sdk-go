// Package option 提供构建 core.Client 的配置选项。
// 本包依赖 core 与签名实现；core 不反向导入本包。
package option

import (
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/tiyee/go-alipay/core"
	"github.com/tiyee/go-alipay/core/auth/signers"
	"github.com/tiyee/go-alipay/core/auth/verifiers"
)

// WithAppID 设置应用 APPID。
func WithAppID(appID string) core.ClientOption {
	return func(s *core.Settings) error {
		s.AppID = appID
		return nil
	}
}

// WithEnvironment 按环境设置官方网关地址。
// 沙箱与生产互不回退；未在首次配置中选择的路径不会自动切换。
func WithEnvironment(env core.Environment) core.ClientOption {
	return func(s *core.Settings) error {
		switch env {
		case core.EnvironmentProduction:
			return WithGateway(core.GatewayProduction)(s)
		case core.EnvironmentSandbox:
			return WithGateway(core.GatewaySandbox)(s)
		default:
			return &core.ConfigError{Field: "environment",
				Reason: "unknown environment: " + string(env)}
		}
	}
}

// WithGateway 显式设置网关地址，用于测试或受控代理。
func WithGateway(rawURL string) core.ClientOption {
	return func(s *core.Settings) error {
		parsed, err := url.Parse(rawURL)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			return &core.ConfigError{Field: "gateway", Reason: "invalid gateway url"}
		}
		s.Gateway = parsed
		return nil
	}
}

// WithSigner 注入已构建的签名器。
func WithSigner(signer core.Signer) core.ClientOption {
	return func(s *core.Settings) error {
		if signer == nil {
			return &core.ConfigError{Field: "signer", Reason: "nil signer"}
		}
		s.Signer = signer
		return nil
	}
}

// WithVerifier 注入已构建的验签器，并将认证模式设置为公钥模式；
// 证书模式的验签器随 P5 的证书选项提供。
func WithVerifier(verifier core.Verifier) core.ClientOption {
	return func(s *core.Settings) error {
		if verifier == nil {
			return &core.ConfigError{Field: "verifier", Reason: "nil verifier"}
		}
		s.Verifier = verifier
		s.AuthMode = core.AuthModePublicKey
		return nil
	}
}

// WithPrivateKey 使用 PKCS#8 PEM 编码的应用私钥构建 RSA2 签名器。
func WithPrivateKey(privateKeyPEM string) core.ClientOption {
	return func(s *core.Settings) error {
		signer, err := signers.NewRSA2(privateKeyPEM)
		if err != nil {
			return err
		}
		s.Signer = signer
		return nil
	}
}

// WithAlipayPublicKey 使用 PEM 编码的支付宝公钥构建 RSA2 验签器，
// 并将认证模式设置为公钥模式。
func WithAlipayPublicKey(alipayPublicKeyPEM string) core.ClientOption {
	return func(s *core.Settings) error {
		verifier, err := verifiers.NewRSA2(alipayPublicKeyPEM)
		if err != nil {
			return err
		}
		s.Verifier = verifier
		s.AuthMode = core.AuthModePublicKey
		return nil
	}
}

// WithHTTPClient 注入自定义 HTTP 客户端；调用方管理其连接与超时策略。
// SDK 仍传递 ctx 并关闭所接收的响应体。
func WithHTTPClient(client *http.Client) core.ClientOption {
	return func(s *core.Settings) error {
		if client == nil {
			return &core.ConfigError{Field: "http_client", Reason: "nil http client"}
		}
		s.HTTPClient = client
		return nil
	}
}

// WithRequestTimeout 设置单次请求超时。
func WithRequestTimeout(d time.Duration) core.ClientOption {
	return func(s *core.Settings) error {
		if d < 0 {
			return &core.ConfigError{Field: "request_timeout", Reason: "negative duration"}
		}
		s.RequestTimeout = d
		return nil
	}
}

// WithMaxResponseBytes 设置响应体大小上限。
func WithMaxResponseBytes(n int64) core.ClientOption {
	return func(s *core.Settings) error {
		if n <= 0 {
			return &core.ConfigError{Field: "max_response_bytes", Reason: "must be positive"}
		}
		s.MaxResponseBytes = n
		return nil
	}
}

// WithLogger 注入结构化日志 Logger；显式传入 nil 视为配置错误。
// 未配置时 Client 静默，不读取全局默认 Logger。
func WithLogger(logger *slog.Logger) core.ClientOption {
	return func(s *core.Settings) error {
		if logger == nil {
			return &core.ConfigError{Field: "logger", Reason: "nil logger; omit the option for silence"}
		}
		s.Logger = logger
		return nil
	}
}

// WithLogHandler 通过任意 slog.Handler 实现注入日志；
// 显式传入 nil 视为配置错误。
func WithLogHandler(handler slog.Handler) core.ClientOption {
	return func(s *core.Settings) error {
		if handler == nil {
			return &core.ConfigError{Field: "log_handler", Reason: "nil handler; omit the option for silence"}
		}
		s.Logger = slog.New(handler)
		return nil
	}
}
