// V3 REST 协议的签名、验签与错误解析。
//
// 支付宝 V3 [签名规则]与[验签规则]规定了签名原文的构造方式：
//   - 请求签名放在 Authorization 头，方案名为 ALIPAY-SHA256withRSA；
//   - 签名原文为多行串：authString、HTTP 方法、请求 URI（含 query）、
//     请求体（可为空行）、应用授权令牌（可省略），每行以 \n 结尾；
//   - 成功响应的签名在 alipay-signature 头，验签原文为
//     alipay-timestamp、alipay-nonce 与响应体三行，每行以 \n 结尾；
//   - 证书模式请求携带 alipay-root-cert-sn 头，响应以 alipay-sn 头
//     标识支付宝公钥证书 SN；
//   - 失败响应为非 2xx 状态加 JSON 错误体 {code, message, links?}。
//
// 详细规则记录于 docs/protocol-notes.md V3 章节。
//
// [签名规则]: https://opendocs.alipay.com/open-v3/054q58
// [验签规则]: https://opendocs.alipay.com/open-v3/054d0z
package protocol

import (
	"encoding/json"
	"fmt"
)

// V3 鉴权方案名与请求头名称。
const (
	// V3AuthScheme 是 Authorization 头中的签名方案名。
	V3AuthScheme = "ALIPAY-SHA256withRSA"

	// V3HeaderAuthorization 是 V3 请求鉴权头。
	V3HeaderAuthorization = "Authorization"
	// V3HeaderRootCertSN 是证书模式请求携带的支付宝根证书 SN 头。
	V3HeaderRootCertSN = "alipay-root-cert-sn"
	// V3HeaderRequestID 是 V3 请求标识头，值为随机 ID，不参与签名。
	V3HeaderRequestID = "alipay-request-id"
	// V3HeaderAppAuthToken 是 V3 代调用授权令牌头。
	V3HeaderAppAuthToken = "alipay-app-auth-token"
	// V3HeaderSignature 是成功响应的签名头。
	V3HeaderSignature = "alipay-signature"
	// V3HeaderCertSN 是响应携带的支付宝公钥证书 SN 头，
	// 证书模式下用于选择本地可信证书。
	V3HeaderCertSN = "alipay-sn"
	// V3HeaderTimestamp 是响应签名的时间戳头。
	V3HeaderTimestamp = "alipay-timestamp"
	// V3HeaderNonce 是响应签名的随机串头。
	V3HeaderNonce = "alipay-nonce"

	// V3PathPrefix 是官方 V3 REST 路由的公共前缀。
	V3PathPrefix = "/v3/"
)

// BuildV3AuthString 组装 Authorization 头中的 authString：
// app_id=...,app_cert_sn=...,nonce=...,timestamp=...。
// appCertSN 仅在证书模式提供；nonce 与 timestamp 由调用方生成。
func BuildV3AuthString(appID, appCertSN, nonce, timestamp string) string {
	authString := "app_id=" + appID
	if appCertSN != "" {
		authString += ",app_cert_sn=" + appCertSN
	}
	authString += ",nonce=" + nonce + ",timestamp=" + timestamp
	return authString
}

// BuildV3SignContent 组装 V3 请求签名原文。各行以 \n 结尾：
// authString、HTTP 方法、请求 URI（不含域名，含 query）、请求体
// （为空时是空行）、应用授权令牌（为空时整体省略）。
// 与官方实现一致，正文为空不代表省略该行。
func BuildV3SignContent(authString, method, uri, body, appAuthToken string) []byte {
	content := authString + "\n" +
		method + "\n" +
		uri + "\n" +
		body + "\n"
	if appAuthToken != "" {
		content += appAuthToken + "\n"
	}
	return []byte(content)
}

// BuildV3VerifyContent 组装 V3 响应验签原文：
// timestamp、nonce 与响应体三行，每行以 \n 结尾；响应体为空时是空行。
func BuildV3VerifyContent(timestamp, nonce string, body []byte) []byte {
	return []byte(timestamp + "\n" + nonce + "\n" + string(body) + "\n")
}

// V3Error 是 V3 REST 协议的错误响应体，字段依据官方 CommonErrorType：
// code 与 message 必有，links 可选。
type V3Error struct {
	// Code 是错误码，例如网关级 "invalid-signature" 或业务级
	// "ACQ.TRADE_NOT_EXIST"。未知取值按原值保留。
	Code string
	// Message 是错误描述原文。
	Message string
	// Links 是官方解决方案链接，可能为空。
	Links string
}

// ParseV3Error 从非 2xx 响应体解析错误 JSON。
// 顶层必须是 JSON 对象且 code、message 均存在，否则返回错误。
func ParseV3Error(body []byte) (V3Error, error) {
	var value struct {
		Code    *string `json:"code"`
		Message *string `json:"message"`
		Links   string  `json:"links"`
	}
	if err := json.Unmarshal(body, &value); err != nil {
		return V3Error{}, fmt.Errorf("decode error body: %w", err)
	}
	switch {
	case value.Code == nil:
		return V3Error{}, fmt.Errorf("error body is missing code")
	case value.Message == nil:
		return V3Error{}, fmt.Errorf("error body is missing message")
	}
	return V3Error{Code: *value.Code, Message: *value.Message, Links: value.Links}, nil
}
