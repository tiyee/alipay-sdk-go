package core

import (
	"bytes"
	"context"
	"html"
	"net/url"
	"slices"
	"time"

	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
	"github.com/tiyee/alipay-sdk-go/internal/protocol"
)

// ArtifactBuilder 构建页面跳转与唤端类接口的签名产物（支付 URL、
// HTML Form、App 订单串）。这类接口不走服务端 REST 调用：
// 产物按官方页面/唤端契约携带 method、version=1.0 与 biz_content，
// 由用户浏览器或支付宝客户端 SDK 提交给产物网关 gateway.do；
// 本类型只做参数组装与签名，不发起任何 HTTP 请求。
//
// 依据官方页面跳转与唤端说明（https://opendocs.alipay.com/common/0i2kxr.md、
// https://opendocs.alipay.com/common/0i2jez.md）与官方 V3 SDK
// GenericExecuteApi 的 pageExecute/sdkExecute 实现。
type ArtifactBuilder struct {
	settings Settings
	snapshot authSnapshot
	now      func() time.Time
}

// NewArtifactBuilder 以 Credentials 为第一个必选参数构建产物构建器。
// 产物网关必须显式选择：WithEnvironment 会按环境成对设置 REST 与
// 产物网关，也可用 option.WithArtifactGateway 单独指定；
// 证书模式凭证提供 app_cert_sn 与 alipay_root_cert_sn 公共参数。
// 构造过程不进行网络 I/O。
func NewArtifactBuilder(credentials auth.Credentials, opts ...ClientOption) (*ArtifactBuilder, error) {
	snapshot, err := bindCredentials(credentials)
	if err != nil {
		return nil, err
	}
	var settings Settings
	for _, opt := range opts {
		if opt == nil {
			return nil, &ConfigError{Field: "option", Reason: "nil option"}
		}
		if err := opt(&settings); err != nil {
			return nil, err
		}
	}
	if settings.Mode != auth.ModeUnset {
		return nil, &ConfigError{Field: "auth_mode", Reason: "auth mode is derived from credentials and cannot be set by options"}
	}
	settings.Mode = snapshot.mode
	settings.applyDefaults()
	if err := settings.validate(); err != nil {
		return nil, err
	}
	if err := settings.validateArtifactGateway(); err != nil {
		return nil, err
	}
	settings.Gateway = compat.CloneURL(settings.Gateway)
	settings.ArtifactGateway = compat.CloneURL(settings.ArtifactGateway)
	return &ArtifactBuilder{settings: settings, snapshot: snapshot, now: time.Now}, nil
}

// Settings 返回公共配置快照，语义与 [Client.Settings] 一致。
func (b *ArtifactBuilder) Settings() Settings {
	settings := b.settings
	settings.Gateway = compat.CloneURL(b.settings.Gateway)
	settings.ArtifactGateway = compat.CloneURL(b.settings.ArtifactGateway)
	return settings
}

// Mode 返回构造期绑定的认证模式。
func (b *ArtifactBuilder) Mode() auth.Mode { return b.snapshot.mode }

// Artifact 是一次产物构建的结果。Values 是签名后的全部参数
// （公共参数 + 业务参数 + sign），按名升序排列；公共参数子集另存
// 于 systemValues，业务参数子集另存于 BizParams。值在序列化为
// URL 或表单时按 rfc3986 规则编码；签名覆盖编码前的参数原值。
type Artifact struct {
	// Method 是接口方法名。
	Method string
	// Gateway 是产物提交的网关地址（例如 gateway.do）。
	Gateway string
	// Values 是参与提交的全部参数（含 sign）。
	Values url.Values
	// BizParams 是业务参数子集（biz_content、return_url、
	// app_auth_token 等），HTML Form 中作为隐藏域提交。
	BizParams url.Values

	system url.Values
}

// OrderString 返回唤端订单串：全部参数按名升序以 k=v 用 & 连接，
// 值做 URL 编码。产物交给支付宝客户端 SDK 使用。
func (a *Artifact) OrderString() string { return a.Values.Encode() }

// RedirectURL 返回页面跳转地址：网关地址拼接全部参数 query。
// 跳转完成不代表支付完成，支付结果以异步通知与查询接口为准。
func (a *Artifact) RedirectURL() string {
	return a.Gateway + "?" + a.Values.Encode()
}

// FormHTML 返回自动提交的 HTML 表单：公共参数位于 action 的 query，
// 业务参数作为隐藏域；属性中的特殊字符以 HTML 实体表示，浏览器提交时还原为签名原值。
// 输出末尾附带自动提交脚本；页面返回地址不作为支付完成的可信证据。
func (a *Artifact) FormHTML() string {
	action := a.Gateway + "?" + a.system.Encode()
	var sb bytes.Buffer
	sb.WriteString("<form name=\"punchout_form\" method=\"post\" action=\"")
	sb.WriteString(html.EscapeString(action))
	sb.WriteString("\">\n")
	for _, key := range sortedKeys(a.BizParams) {
		for _, value := range a.BizParams[key] {
			if key == "" || value == "" {
				continue
			}
			sb.WriteString("<input type=\"hidden\" name=\"")
			sb.WriteString(html.EscapeString(key))
			sb.WriteString("\" value=\"")
			sb.WriteString(html.EscapeString(value))
			sb.WriteString("\">\n")
		}
	}
	sb.WriteString("<input type=\"submit\" value=\"立即支付\" style=\"display:none\" >\n")
	sb.WriteString("</form>\n")
	sb.WriteString("<script>document.forms[0].submit();</script>")
	return sb.String()
}

// BuildArtifact 组装并签名页面/唤端产物。
//
// method 是官方接口方法名（如 alipay.trade.page.pay）；bizParams
// 是业务参数（如 biz_content、return_url），不得包含公共参数名，
// 空值参数不参与签名与产物。CallSettings 中的 NotifyURL 与
// AppAuthToken 分别映射为公共参数 notify_url 与业务参数
// app_auth_token。公共参数固定值依据官方产物实现：
// format=json、charset=UTF-8、version=1.0、sign_type=RSA2，
// timestamp 为 GMT+8 的 yyyy-MM-dd HH:mm:ss。
func (b *ArtifactBuilder) BuildArtifact(ctx context.Context, method string, bizParams url.Values, settings CallSettings) (*Artifact, error) {
	if ctx == nil {
		return nil, &ValidationError{Field: "ctx", Reason: "context must not be nil"}
	}
	if method == "" {
		return nil, &ValidationError{Field: "method", Reason: "method is required"}
	}
	biz := url.Values{}
	for key, values := range bizParams {
		// biz_content 是产物业务参数的载体，由调用方提供；
		// 其余受保护公共参数（含 notify_url，经 CallSettings 注入）不得覆盖。
		if key != protocol.ParamBizContent && (protocol.IsProtectedParam(key) || key == protocol.ParamSign) {
			return nil, &ValidationError{Field: key,
				Reason: "biz param overwrites a protected public parameter"}
		}
		if key == "" || len(values) > 1 {
			return nil, &ValidationError{Field: "biz_params", Reason: "expected a nonempty key with at most one value"}
		}
		for _, value := range values {
			if value != "" {
				biz.Add(key, value)
			}
		}
	}
	system := url.Values{}
	system.Set(protocol.ParamMethod, method)
	system.Set(protocol.ParamAppID, b.settings.AppID)
	system.Set(protocol.ParamFormat, protocol.FormatJSONArtifact)
	system.Set(protocol.ParamCharset, protocol.CharsetUTF8Artifact)
	system.Set(protocol.ParamSignType, b.snapshot.signer.Algorithm().String())
	system.Set(protocol.ParamTimestamp, b.now().In(gmtEight).Format(protocol.TimestampLayout))
	system.Set(protocol.ParamVersion, "1.0")
	if b.snapshot.mode == auth.ModeCertificate {
		system.Set(protocol.ParamAppCertSN, b.snapshot.identity.AppCertificateSN)
		system.Set(protocol.ParamAlipayRootCertSN, b.snapshot.identity.RootCertificateSN)
	}
	if settings.NotifyURL != "" {
		system.Set(protocol.ParamNotifyURL, settings.NotifyURL)
	}
	if settings.AppAuthToken != "" {
		biz.Set(protocol.ParamAppAuthToken, settings.AppAuthToken)
	}
	values := url.Values{}
	for key, vs := range system {
		values[key] = append([]string(nil), vs...)
	}
	for key, vs := range biz {
		values[key] = append([]string(nil), vs...)
	}
	signContent := protocol.BuildSignContent(values)
	signature, err := b.snapshot.signer.Sign(ctx, []byte(signContent))
	if err != nil {
		if _, ok := compat.AsType[*auth.CertValidityError](err); ok {
			return nil, &CertificateError{Reason: "artifact certificate validity check failed", Err: err}
		}
		return nil, &SignatureError{Stage: "request", Reason: "sign artifact", Err: err}
	}
	values.Set(protocol.ParamSign, signature.Value)
	system.Set(protocol.ParamSign, signature.Value)
	return &Artifact{
		Method:    method,
		Gateway:   b.settings.ArtifactGateway.String(),
		Values:    values,
		BizParams: biz,
		system:    system,
	}, nil
}

// gmtEight 是官方产物时间戳使用的固定时区（GMT+8）。
var gmtEight = time.FixedZone("GMT+8", 8*60*60)

// sortedKeys 返回升序参数名。
func sortedKeys(values url.Values) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}
