// Package protocol 实现页面/唤端产物与异步通知共用的参数规范化，
// 以及 V3 协议的签名串、验签串与错误解析。实现不依赖 Client，
// 由 core 包调用。
package protocol

import (
	"net/url"
	"regexp"
	"slices"
	"strings"
)

const notifyNameChars = "[A-Za-z0-9_]+"

var notifyName = regexp.MustCompile("^" + notifyNameChars + "$")
var notifyBoundary = regexp.MustCompile("&" + notifyNameChars + "=")

// ValidNotifyName 报告字段名能否按普通支付通知的字段边界规则识别。
func ValidNotifyName(name string) bool { return notifyName.MatchString(name) }

// HasNotifyBoundary 报告单次解码的字段值是否可能被解释为另一个通知字段。
func HasNotifyBoundary(value string) bool { return notifyBoundary.MatchString(value) }

// 网关协议的公共参数名与固定取值。
const (
	ParamAppID        = "app_id"
	ParamMethod       = "method"
	ParamFormat       = "format"
	ParamCharset      = "charset"
	ParamSignType     = "sign_type"
	ParamSign         = "sign"
	ParamTimestamp    = "timestamp"
	ParamVersion      = "version"
	ParamBizContent   = "biz_content"
	ParamNotifyURL    = "notify_url"
	ParamAppAuthToken = "app_auth_token"
	// ParamNotifyType 是异步通知的 notify_type 参数名。
	ParamNotifyType = "notify_type"

	// ParamAppCertSN 是证书模式请求公共参数 app_cert_sn（应用公钥证书 SN）。
	// 网关依据该参数区分公钥模式与证书模式，
	// 依据 https://opendocs.alipay.com/common/02kdne.md。
	ParamAppCertSN = "app_cert_sn"
	// ParamAlipayRootCertSN 是证书模式请求公共参数
	// alipay_root_cert_sn（支付宝根证书 SN）。
	ParamAlipayRootCertSN = "alipay_root_cert_sn"

	CharsetUTF8 = "utf-8"

	// TimestampLayout 是公共参数 timestamp 的格式。
	TimestampLayout = "2006-01-02 15:04:05"

	// 页面/唤端产物的公共参数固定值，依据官方 GenericExecuteApi
	// 实现与页面跳转说明（https://opendocs.alipay.com/common/0i2kxr.md），
	// 与网关请求的取值大小写不同，产物保持官方原样。
	FormatJSONArtifact  = "json"
	CharsetUTF8Artifact = "UTF-8"
)

// protectedParams 是调用方业务参数不得覆盖的公共参数集合。
// 覆盖尝试会被拒绝，而不是静默合并。
var protectedParams = map[string]struct{}{
	ParamAppID:            {},
	ParamMethod:           {},
	ParamFormat:           {},
	ParamCharset:          {},
	ParamSignType:         {},
	ParamSign:             {},
	ParamTimestamp:        {},
	ParamVersion:          {},
	ParamBizContent:       {},
	ParamNotifyURL:        {},
	ParamAppAuthToken:     {},
	ParamAppCertSN:        {},
	ParamAlipayRootCertSN: {},
}

// IsProtectedParam 报告参数名是否属于受保护的公共参数。
func IsProtectedParam(key string) bool {
	_, ok := protectedParams[key]
	return ok
}

// BuildSignContent 构造待签名串：剔除 sign 参数与空值后，
// 按参数名 ASCII 升序以 k=v 形式用 & 连接，值不进行 URL 编码。
// 调用方保证每个键只有一个值且值非空；本函数保持防御性检查。
func BuildSignContent(values url.Values) string {
	keys := make([]string, 0, len(values))
	for key, vs := range values {
		if key == ParamSign || len(vs) == 0 || vs[0] == "" {
			continue
		}
		keys = append(keys, key)
	}
	slices.Sort(keys)
	var sb strings.Builder
	for i, key := range keys {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(key)
		sb.WriteByte('=')
		sb.WriteString(values.Get(key))
	}
	return sb.String()
}

// BuildNotifySignContent 构造异步通知的待验签串：剔除 exclude 列出的
// 参数后，其余参数（含空值参数）按参数名 ASCII 升序以 k=v 用 & 连接，
// 值为已解码原文，不进行 URL 编码。
// 依据[官方验签文档](https://opendocs.alipay.com/common/02mse7.md)，
// 通知的待验签串仅剔除指定参数（普通支付通知为 sign 与 sign_type），
// 不剔除空值参数；这与请求待签名串剔除空值的规则不同。
// 重复参数应在调用前拒绝；本函数对多值键取首个值。
func BuildNotifySignContent(values url.Values, exclude ...string) string {
	excluded := make(map[string]struct{}, len(exclude))
	for _, key := range exclude {
		excluded[key] = struct{}{}
	}
	keys := make([]string, 0, len(values))
	for key, vs := range values {
		if _, ok := excluded[key]; ok || len(vs) == 0 {
			continue
		}
		keys = append(keys, key)
	}
	slices.Sort(keys)
	var sb strings.Builder
	for i, key := range keys {
		if i > 0 {
			sb.WriteByte('&')
		}
		sb.WriteString(key)
		sb.WriteByte('=')
		sb.WriteString(values.Get(key))
	}
	return sb.String()
}
