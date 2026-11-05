package core

import (
	"log/slog"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/tiyee/alipay-sdk-go/core/auth"
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

// V3 REST OpenAPI 的网关基地址（V3 相对路由拼接到该基地址之后），
// 生产与沙箱接入点见支付宝 V3 基本原则：
// https://opendocs.alipay.com/open-v3/054oog
const (
	// GatewayProduction 是生产环境 V3 网关基地址。
	GatewayProduction = "https://openapi.alipay.com"
	// GatewaySandbox 是支持 TLS 的支付宝沙箱 V3 REST 网关。
	GatewaySandbox = "https://openapi-sandbox.dl.alipaydev.com"
)

// 页面跳转与唤端产物的提交网关（gateway.do）。产物不走 V3 REST：
// 签名产物由用户浏览器或支付宝客户端 SDK 提交，因此 REST 地址与
// 产物网关是两项独立配置，WithEnvironment 会按环境成对设置。
const (
	// ArtifactGatewayProduction 是生产环境产物提交网关。
	ArtifactGatewayProduction = "https://openapi.alipay.com/gateway.do"
	// ArtifactGatewaySandbox 是沙箱环境产物提交网关。
	ArtifactGatewaySandbox = "https://openapi-sandbox.dl.alipaydev.com/gateway.do"
)

// 工程默认值，不是支付宝平台协议限制。
const (
	// DefaultRequestTimeout 是单次请求的默认超时。
	DefaultRequestTimeout = 30 * time.Second
	// DefaultMaxResponseBytes 是普通 JSON 响应的默认大小上限。
	DefaultMaxResponseBytes = int64(8 << 20)
)

// Settings 汇总 Client 的公共静态配置。认证模式、签名器与验签器
// 由 NewClient 的首参 Credentials 派生并绑定，公共选项无法覆盖；
// 字段由 core/option 中的选项填充，初始化后保持稳定，
// 调用方不应在并发调用时修改。
type Settings struct {
	// AppID 是支付宝开放平台应用 APPID。
	AppID string
	// Gateway 是 V3 REST 网关基地址；Client 必须显式选择环境或配置。
	Gateway *url.URL
	// ArtifactGateway 是页面跳转与唤端产物的提交网关（gateway.do）；
	// 仅 ArtifactBuilder 需要，与 REST 地址相互独立。
	ArtifactGateway *url.URL
	// Mode 是认证模式，由首参 Credentials 派生的只读信息。
	Mode auth.Mode
	// HTTPClient 是可选的自定义 HTTP 客户端；
	// SDK 复用其传输与超时设置，并在专用副本上禁止重定向。
	HTTPClient *http.Client
	// RequestTimeout 是单次请求超时；零值使用 DefaultRequestTimeout。
	RequestTimeout time.Duration
	// MaxResponseBytes 是响应体大小上限；零值使用 DefaultMaxResponseBytes。
	MaxResponseBytes int64
	// Logger 是结构化日志 Logger；nil 表示静默。
	Logger *slog.Logger
}

// ClientOption 修改 Client 的公共配置。选项实现位于 core/option；
// 选项不能设置认证模式、签名器或验签器。
type ClientOption func(*Settings) error

// authSnapshot 是构造期绑定的只读认证快照：模式、签名器、验签器
// 与证书身份。执行流程只读取快照，不回调 Credentials 的 getter，
// 调用方原地修改凭证配置不会切换已构造 Client 的认证行为。
type authSnapshot struct {
	mode     auth.Mode
	signer   auth.Signer
	verifier auth.Verifier
	identity auth.CertificateIdentity
}

// bindCredentials 校验凭证并冻结认证快照：
// 拒绝 nil 与 typed nil；要求已定义模式、非空算法对象与完整材料；
// 证书模式必须同时具备证书身份能力。不访问网络、不下载证书。
func bindCredentials(credentials auth.Credentials) (authSnapshot, error) {
	if credentials == nil || isNilCredentials(credentials) {
		return authSnapshot{}, &ConfigError{Field: "credentials", Reason: "credentials is required"}
	}
	if err := credentials.Validate(); err != nil {
		return authSnapshot{}, &ConfigError{Field: "credentials", Reason: "credentials validation failed", Err: err}
	}
	mode := credentials.Mode()
	switch mode {
	case auth.ModePublicKey, auth.ModeCertificate:
	default:
		return authSnapshot{}, &ConfigError{Field: "auth_mode", Reason: "unknown auth mode"}
	}
	signer := credentials.Signer()
	if signer == nil || isNilCredentials(signer) {
		return authSnapshot{}, &ConfigError{Field: "signer", Reason: "credentials returned a nil signer"}
	}
	verifier := credentials.Verifier()
	if verifier == nil || isNilCredentials(verifier) {
		return authSnapshot{}, &ConfigError{Field: "verifier", Reason: "credentials returned a nil verifier"}
	}
	if signer.Algorithm() != auth.SignatureAlgorithmRSA2 {
		return authSnapshot{}, &ConfigError{Field: "signer",
			Reason: "traditional gateway requests require the " + auth.SignatureAlgorithmRSA2.String() + " algorithm"}
	}
	snapshot := authSnapshot{mode: mode, signer: signer, verifier: verifier}
	if mode == auth.ModeCertificate {
		provider, ok := credentials.(auth.CertificateProvider)
		if !ok {
			return authSnapshot{}, &ConfigError{Field: "credentials",
				Reason: "certificate mode requires the certificate identity capability"}
		}
		identity := provider.CertificateIdentity()
		if identity.AppCertificateSN == "" || identity.RootCertificateSN == "" {
			return authSnapshot{}, &ConfigError{Field: "certificate_identity",
				Reason: "certificate identity is incomplete"}
		}
		snapshot.identity = identity
	}
	return snapshot, nil
}

// isNilCredentials 报告接口值是否为 nil 或持有 nil 指针（typed nil）。
func isNilCredentials(value any) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Func,
		reflect.Chan, reflect.Interface, reflect.UnsafePointer:
		return rv.IsNil()
	default:
		return false
	}
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
	if s.Mode == auth.ModeUnset {
		return &ConfigError{Field: "auth_mode", Reason: "auth mode is required"}
	}
	if s.RequestTimeout < 0 {
		return &ConfigError{Field: "request_timeout", Reason: "timeout must not be negative"}
	}
	return nil
}

// validateGateway 校验 V3 REST 网关基地址；仅 Client 构造路径使用。
func (s *Settings) validateGateway() error {
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
	return nil
}

// validateArtifactGateway 校验产物提交网关；仅 ArtifactBuilder 使用。
func (s *Settings) validateArtifactGateway() error {
	if s.ArtifactGateway == nil {
		return &ConfigError{Field: "artifact_gateway",
			Reason: "artifact gateway is required; set an environment or an explicit artifact gateway"}
	}
	if s.ArtifactGateway.Scheme != "http" && s.ArtifactGateway.Scheme != "https" {
		return &ConfigError{Field: "artifact_gateway", Reason: "artifact gateway scheme must be http or https"}
	}
	if s.ArtifactGateway.Host == "" {
		return &ConfigError{Field: "artifact_gateway", Reason: "artifact gateway host is required"}
	}
	return nil
}

// signedHTTPClient 保留调用方的传输和超时设置，并禁止重发签名请求。
func signedHTTPClient(configured *http.Client) *http.Client {
	var client http.Client
	if configured != nil {
		client = *configured
	} else if transport, ok := http.DefaultTransport.(*http.Transport); ok {
		client.Transport = transport.Clone()
	} else {
		client.Transport = http.DefaultTransport
	}
	client.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	return &client
}
