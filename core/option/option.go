// Package option 提供构建 core.Client 的公共配置选项。
// 认证凭证是 NewClient 的第一个必选参数（core/auth.Credentials），
// 本包选项不设置认证模式、签名器或验签器；
// 公共选项无法覆盖首参确定的认证策略。
package option

import (
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/tiyee/alipay-sdk-go/core"
)

// WithAppID 设置应用 APPID。
func WithAppID(appID string) core.ClientOption {
	return func(s *core.Settings) error {
		s.AppID = appID
		return nil
	}
}

// WithEnvironment 按环境成对设置官方 REST 网关与产物提交网关。
// 沙箱与生产互不回退；未在首次配置中选择的路径不会自动切换。
func WithEnvironment(env core.Environment) core.ClientOption {
	return func(s *core.Settings) error {
		switch env {
		case core.EnvironmentProduction:
			return withOfficialGateways(core.GatewayProduction, core.ArtifactGatewayProduction)(s)
		case core.EnvironmentSandbox:
			return withOfficialGateways(core.GatewaySandbox, core.ArtifactGatewaySandbox)(s)
		default:
			return &core.ConfigError{Field: "environment",
				Reason: "unknown environment: " + string(env)}
		}
	}
}

// withOfficialGateways 按官方地址成对设置 REST 网关与产物网关。
func withOfficialGateways(gateway, artifactGateway string) core.ClientOption {
	return func(s *core.Settings) error {
		if err := WithGateway(gateway)(s); err != nil {
			return err
		}
		return WithArtifactGateway(artifactGateway)(s)
	}
}

// WithGateway 显式设置 V3 REST 网关基地址，用于测试或受控代理；
// 不影响产物提交网关。
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

// WithArtifactGateway 显式设置页面跳转与唤端产物的提交网关
// （gateway.do），用于测试或受控代理；不影响 REST 网关。
func WithArtifactGateway(rawURL string) core.ClientOption {
	return func(s *core.Settings) error {
		parsed, err := url.Parse(rawURL)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			return &core.ConfigError{Field: "artifact_gateway", Reason: "invalid artifact gateway url"}
		}
		s.ArtifactGateway = parsed
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
			return &core.ConfigError{Field: "logger",
				Reason: "nil logger; omit the option for silence"}
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
			return &core.ConfigError{Field: "log_handler",
				Reason: "nil handler; omit the option for silence"}
		}
		s.Logger = slog.New(handler)
		return nil
	}
}
