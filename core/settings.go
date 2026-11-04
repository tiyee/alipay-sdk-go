package core

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tiyee/go-alipay/core/auth"
)

// Environment 标识支付宝接入环境。环境必须显式选择，
// 沙箱与生产不能自动回退或混用。
type Environment string

const (
	// EnvironmentProduction 是生产环境。
	EnvironmentProduction Environment = "production"
	// EnvironmentSandbox 是沙箱环境，网关依据官方沙箱接入说明。
	EnvironmentSandbox Environment = "sandbox"
)

// String 实现 fmt.Stringer。
func (e Environment) String() string { return string(e) }

// AuthMode 标识 Client 的认证模式。模式由装载的凭证决定：
// 公钥模式使用支付宝公钥验签；证书模式按响应 SN 选择可信证书（P5）。
type AuthMode uint8

const (
	// AuthModeUnset 表示尚未选择认证模式，Client 初始化失败。
	AuthModeUnset AuthMode = iota
	// AuthModePublicKey 表示公钥模式。
	AuthModePublicKey
	// AuthModeCertificate 表示本地证书模式。
	AuthModeCertificate
)

// String 实现 fmt.Stringer。
func (m AuthMode) String() string {
	switch m {
	case AuthModePublicKey:
		return "public_key"
	case AuthModeCertificate:
		return "certificate"
	default:
		return "unset"
	}
}

// 传统网关 OpenAPI 的网关地址，依据官方沙箱接入说明
// （https://opendocs.alipay.com/common/097jyh.md）与签名规范维护。
const (
	// GatewayProduction 是生产环境传统网关地址。
	GatewayProduction = "https://openapi.alipay.com/gateway.do"
	// GatewaySandbox 是沙箱环境传统网关地址。
	GatewaySandbox = "https://openapi-sandbox.dl.alipaydev.com/gateway.do"
)

// 工程默认值，不是支付宝平台协议限制。
const (
	// DefaultRequestTimeout 是单次请求的默认超时。
	DefaultRequestTimeout = 30 * time.Second
	// DefaultMaxResponseBytes 是普通 JSON 响应的默认大小上限。
	DefaultMaxResponseBytes = int64(8 << 20)
)

// Settings 汇总 Client 的静态配置。字段由 core/option 中的选项填充；
// Client 初始化后配置保持稳定，调用方不应在并发调用时修改。
type Settings struct {
	// AppID 是支付宝开放平台应用 APPID。
	AppID string
	// Gateway 是网关地址；必须显式选择环境或配置网关。
	Gateway *url.URL
	// Signer 生成请求签名。
	Signer auth.Signer
	// Verifier 校验响应签名。
	Verifier auth.Verifier
	// AuthMode 是当前认证模式，由凭证选项设置；未选择时初始化失败。
	AuthMode AuthMode
	// HTTPClient 是可选的自定义 HTTP 客户端；
	// 注入时由调用方管理连接与超时策略，SDK 仍传递 ctx 并关闭响应体。
	HTTPClient *http.Client
	// RequestTimeout 是单次请求超时；零值使用 DefaultRequestTimeout。
	RequestTimeout time.Duration
	// MaxResponseBytes 是响应体大小上限；零值使用 DefaultMaxResponseBytes。
	MaxResponseBytes int64
	// Logger 是结构化日志 Logger；nil 表示静默。
	Logger *slog.Logger
}

// ClientOption 修改 Client 配置。选项实现位于 core/option。
type ClientOption func(*Settings) error

// Client 是支付宝传统网关 OpenAPI 的通用客户端，负责公共参数、
// 编码、签名、HTTP 执行、原始响应提取、验签与错误分类。
// Client 不解释商户订单状态，不包含业务字段。
// 初始化后可并发复用；每次调用复制参数并生成独立的时间戳与签名。
type Client struct {
	settings Settings
	http     *http.Client
	logger   *slog.Logger
	now      func() time.Time
}

// NewClient 应用选项并校验配置；配置不完整时返回 *ConfigError。
// 构造过程不进行网络 I/O。
func NewClient(opts ...ClientOption) (*Client, error) {
	var settings Settings
	for _, opt := range opts {
		if opt == nil {
			return nil, &ConfigError{Field: "option", Reason: "nil option"}
		}
		if err := opt(&settings); err != nil {
			return nil, err
		}
	}
	settings.applyDefaults()
	if err := settings.validate(); err != nil {
		return nil, err
	}
	httpClient := settings.HTTPClient
	if httpClient == nil {
		httpClient = defaultHTTPClient()
	}
	return &Client{
		settings: settings,
		http:     httpClient,
		logger:   settings.Logger,
		now:      time.Now,
	}, nil
}

// Settings 返回配置的快照，用于诊断；修改快照不影响 Client。
func (c *Client) Settings() Settings {
	c.settings.Gateway = cloneURL(c.settings.Gateway)
	return c.settings
}

func (s *Settings) applyDefaults() {
	if s.RequestTimeout == 0 {
		s.RequestTimeout = DefaultRequestTimeout
	}
	if s.MaxResponseBytes <= 0 {
		s.MaxResponseBytes = DefaultMaxResponseBytes
	}
	if s.Logger == nil {
		s.Logger = slog.New(slog.DiscardHandler)
	}
}

func (s *Settings) validate() error {
	if strings.TrimSpace(s.AppID) == "" {
		return &ConfigError{Field: "app_id", Reason: "app id is required"}
	}
	if s.Gateway == nil {
		return &ConfigError{Field: "gateway",
			Reason: "gateway is required; set an environment or an explicit gateway"}
	}
	if s.Gateway.Scheme != "http" && s.Gateway.Scheme != "https" {
		return &ConfigError{Field: "gateway", Reason: "gateway scheme must be http or https"}
	}
	if s.Gateway.Host == "" {
		return &ConfigError{Field: "gateway", Reason: "gateway host is required"}
	}
	if s.Signer == nil {
		return &ConfigError{Field: "signer", Reason: "signer is required"}
	}
	if s.Verifier == nil {
		return &ConfigError{Field: "verifier", Reason: "verifier is required"}
	}
	if s.AuthMode == AuthModeUnset {
		return &ConfigError{Field: "auth_mode",
			Reason: "auth mode is required; configure a public key or certificate"}
	}
	if s.RequestTimeout < 0 {
		return &ConfigError{Field: "request_timeout", Reason: "timeout must not be negative"}
	}
	return nil
}

// defaultHTTPClient 返回拒绝自动重定向的专用客户端，
// 防止携带凭证的请求发往其他地址；不修改全局默认客户端。
func defaultHTTPClient() *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	return &http.Client{
		Transport: transport,
		CheckRedirect: func(*http.Request, []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func cloneURL(u *url.URL) *url.URL {
	if u == nil {
		return nil
	}
	return u.Clone()
}

func invalidConfigf(field, format string, args ...any) *ConfigError {
	return &ConfigError{Field: field, Reason: fmt.Sprintf(format, args...)}
}
