// Package testutil 提供仅供测试使用的专用虚构密钥与协议断言辅助。
// 所有密钥均为本仓库测试专用的虚构材料，不对应任何真实商户或应用，
// 严禁用于生产环境。
package testutil

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core/auth/signers"
	"github.com/tiyee/alipay-sdk-go/core/auth/verifiers"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// 专用虚构测试材料。app 私钥模拟商户，alipay 私钥模拟支付宝网关，
// 两个密钥对相互独立，避免同一密钥自我证明。
const (
	// AppPrivateKeyPEM 是模拟商户应用的 PKCS#8 测试私钥。
	AppPrivateKeyPEM = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC5YUs80rBV8zTn
1pWcXu/z0SjgGmHJC6MsVpbLhXeER5blh/MMa5AI12Q+v2Z88wN77Q6vgD4zkOfm
gbVFbHIh9sVOfzRG7BfA86lAunqFY2JMYhaB4pgN9vaqWftZF4TpGDKojVcOy+QN
27CCkzugh4yQDTELTjQ3ntvOdpVaNcFjCq4gxVLhhOuzNNVBzm339kmkYOjDnJby
24McqaUc7jGzN7QsODytxchsQQiybcLJlHmYNomRU4yewuNK02Vib46/R9ipFtQ9
Uy4KlfLg+QM8wR6eLUUHTb8HP21xIiEEhHENhQpIF/8BpSa49icq/oAhS9uhBeRA
PV45KWXVAgMBAAECggEADtw3QUVsYD9yNsFzVckgeA2Ets42l3yo1GWOifsFdSOd
+4FV4OGAdsz5WrRSHWMHnPgITNY9ZjnP9SDhrJKAZauX7m5w/NxxHQULRMaOCsJI
P3I/dqL46GmrJUW4SdCCVYMckHlgCK4cdmXtdNIjnf8vq95613qp4emImjjUSRpS
EsUW7Frtm+jlTGFgj1owhR3IWZ0acCawaqwcZrscjvMPQaL8Xxsmn3qM42dAYaG7
66JYBc6JFTzKnm56o3tmMydBEXLRLX09yxveKarXTX+lFQAbuui2ahpLvc4Xt/UP
pRq4asZUtCz50wLvpqTPSFdJ9NYGDBvMBrwwW7HnowKBgQDkWQzoYgilMFxFgri5
GDutjn5iH8s2pv2/6tEhJTncbNQ41UTQt3me5T7tJRJBFAeMnH0UL9/kdi97nZai
IdjvH+krE8TSve8MTjDehXrrFClcq7QplZsAoCLTyvxtzAC+2M/D9qojFsA/hZkh
YBvqPtaScooackzuvVgrVFn7kwKBgQDP1DXf/RG2i1gy7k08NxVghZ6GKV2U0CIp
lfaxQcwQOD7mJTQpQaaU4kWXCOZUYR3MEwIeOx1eb2A0nWtCHLiYZGsjiDy/yDIr
w30tgU30RHHE/bU7pD2Cz12B81iiIX+lQ1iZ0SjxvgpK5mkcS3QthELo+WMN/3ym
WUtUxDKJ9wKBgFdnEn68T2zTJu5gS85lTSpwWUW7v6/KmWPlQGj8c7pAWbbrQuzU
3ZOJUd25i9QnE4CpwWKBtKQg7KnXNvoWFI/q2ffoGDVGCALgxunVk6f5Tp+vBUXJ
gsvpf4srrhVkuYBTQqtSder6YVmvDF7vB6BQrxbdCql2yKW1CdOD9PFLAoGAd6OD
DBlj/ADFWjRik9pl3PWupy7i50rVj7j7lVj6PxAR7WzXVTSsz6JeSy7dYQ+zFoqO
upZvNrDI4iEI4Lsl7NxJ7N99kS2O7DCSPALmvEo2cB9ZGCdZHZeADdHnYRYFpD5b
n9iPSHKM/EcdlSM+kwFDf1gpFKLRv/VBb1+udEsCgYEAnELCldU0C515OivSujPD
J6Lyk5mrGxGESalnxKKU5uDyy3w5rBF8FsaRaAafN9l0uvCOjQjQRn8nIztmlPwa
DMZKYi3aqYFZBXf3X2UYJg8EBJz5e2ov4eWTjgVA2Hzqyk2LV4l9ppFD5Qo+ioGu
AUQQvZXiTGK2/oPx/c0nb6k=
-----END PRIVATE KEY-----`

	// AppPublicKeyPEM 是与 AppPrivateKeyPEM 配对的测试公钥。
	AppPublicKeyPEM = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuWFLPNKwVfM059aVnF7v
89Eo4BphyQujLFaWy4V3hEeW5YfzDGuQCNdkPr9mfPMDe+0Or4A+M5Dn5oG1RWxy
IfbFTn80RuwXwPOpQLp6hWNiTGIWgeKYDfb2qln7WReE6RgyqI1XDsvkDduwgpM7
oIeMkA0xC040N57bznaVWjXBYwquIMVS4YTrszTVQc5t9/ZJpGDow5yW8tuDHKml
HO4xsze0LDg8rcXIbEEIsm3CyZR5mDaJkVOMnsLjStNlYm+Ov0fYqRbUPVMuCpXy
4PkDPMEeni1FB02/Bz9tcSIhBIRxDYUKSBf/AaUmuPYnKv6AIUvboQXkQD1eOSll
1QIDAQAB
-----END PUBLIC KEY-----`

	// AlipayPrivateKeyPEM 是模拟支付宝网关的 PKCS#8 测试私钥，
	// 仅用于在测试中构造网关签名响应。
	AlipayPrivateKeyPEM = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC+5TNSTwndQyo3
Uw8LVAzLYxsS2Iu4A4ObWMKhovnBbeIbVTS4ckpNQsXVY+LYRjaBvLPbrhnG13HW
7N/67K4iwhuTnV+tcjPU+OL15Xt47e3LnRp5gLMB1YA+73Q79d50JIMSvzB0spQN
CdVzDogb/Tw38YfP3QiufglkIUvNWWSOPrtsxo9/XULjAnTL0e5cXNGtLdrbT5+4
ehpKfr74tyyseAQvD/9sQBsLz5300c2oSrcoP7FWkSprqhffQbXXmjOKb/PfQoTW
xaCECSHpxqVytVUXu2Xr9IYAr5VEpHdgVEcWQ2rsGS60C6TD7mi9CEeGHgo05dq1
m5/H1SUfAgMBAAECggEAOEhGbeWc0G49k6s6+8mPriWttUxAQYRTuKkeVV7CSWv2
5r0j+p/D1Wyy/HvjRYXKC01SMBpwQiaCS25yiT5gByJ40/hp3a8WphutevdqhJiN
7GMJC7rMzGUg1y1tdUJEELU1KgFRrY57RS7Zjn9B1RLAPvRI6RnPtfa2GR4SBdxr
A+vvHpJPFSLjimUnfL7lXV7jVI160O0v3Ddl468xpyAUZXXTt66cdOeRwzE7V18s
0O265q/bq+Iqah863CQkCOWejE0vzKEgNpQ2R1og0CMlPtEBfXz88YNR40H9VcaE
Vg2IJ+nccz0QZLswiOpDa0JGhgSQGmI4XS+Jko1TUQKBgQDxXBMOjJSUKBF7uipj
SX4/+8WfgdEQt35597iIBvmOgqmD7Pn9UKmYiWOnrt/L4Owd8fFkzgURxAqM5gZf
xUiZDOKdNuX6fTAoTofSVwrYujQDYfEICWUvGw8fTzTYhDr2fad56xXjHyoyPysb
QSuW2zuzkDGCXxcOKHScCK+bcQKBgQDKeX8Hk7X6Qw8QH6IuMa7WjP9LuHHiKUaD
y1NEXBUbSQtJj7vo3eEC2CEXKUDX3bdMwXADNIhgNO6MWQdlkdrZkSLDnk/ljVR+
tGupT2NwNOfuL98lSH4ghl47RqLXF+VbwTeG4m2u1+qm7Mwrk0xWh1w3vr3VZqjI
8Twz4o/hjwKBgCvRZXxMtWS4aYQm5i8kVpFTKiutK8F3LrUB7jRz37aprhDPYwtE
K6YPu1HkA6lF6lf9Zq/AX1c8WhX+ZMXaUTXICxQztUjmlWmb6DidqNBz7aF3lIS6
/+2vYe1zOcrdJr89TLxrm4uiSszYfwDxsfIvVwVgdjyO+uaNHa5+PTNhAoGAMVH/
G31jqp6DIM5T1yTLDgAyf+Nh+lFLl/cDRXeYvJjCYxzcDRFfKi4zgUzPOsYLTGsi
OeWmAqliIlXdc2RgBvdPnI8Hr0cJ+gNA5nEsK7z3chG/LXnHH20brxy5teNgP0WF
nRmy1hBKXHb6FPNAeSNn1zxjSYREQQ5iN9FMcwkCgYB0eVpTw6+RSTWa4IQYJ9ig
gQbOS7XwsLWUIlrbc5qC7IVKzLmC1VMqZ1B6lW6gkiAXlsJPBMBt++o0kRCrfwOl
MCFEbIzcpVjS+WCKBLPPrUymVVhV4Tw7x4Ru5MWkclwawKJrWVKYWhI+PQ8z3GWU
Z1Ey7NwYNAJtZwPpwgS8tA==
-----END PRIVATE KEY-----`

	// AlipayPublicKeyPEM 是与 AlipayPrivateKeyPEM 配对的测试公钥，
	// 模拟商户配置的支付宝公钥。
	AlipayPublicKeyPEM = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvuUzUk8J3UMqN1MPC1QM
y2MbEtiLuAODm1jCoaL5wW3iG1U0uHJKTULF1WPi2EY2gbyz264Zxtdx1uzf+uyu
IsIbk51frXIz1Pji9eV7eO3ty50aeYCzAdWAPu90O/XedCSDEr8wdLKUDQnVcw6I
G/08N/GHz90Irn4JZCFLzVlkjj67bMaPf11C4wJ0y9HuXFzRrS3a20+fuHoaSn6+
+LcsrHgELw//bEAbC8+d9NHNqEq3KD+xVpEqa6oX30G115ozim/z30KE1sWghAkh
6calcrVVF7tl6/SGAK+VRKR3YFRHFkNq7BkutAukw+5ovQhHhh4KNOXatZufx9Ul
HwIDAQAB
-----END PUBLIC KEY-----`
)

// KnownAnswerSignContent 是独立签名向量对应的待签名串。
const KnownAnswerSignContent = "app_id=2021000000000001&method=alipay.trade.query&timestamp=2026-01-01 00:00:00&version=1.0"

// KnownAnswerSignature 是 openssl dgst -sha256 -sign 对
// KnownAnswerSignContent 与 AppPrivateKeyPEM 独立生成的 base64 签名，
// 用于交叉验证本仓库签名实现。
const KnownAnswerSignature = "RTibFKBCo/OpbcsjlvxGmvqhV6hkvUsfzIKuUw17mXrAC8siGeo0mKpqTQOFWV48B4OjR63alMwlWGLL6BGUgV+HXmyXT930MdfdNfrMac/wRCMqNWVBB2a0yJNwIo7tNNIccdUZqJsiwr5pAaTBueHqc0isyKltJj/YAPyHSxoSjdDa+rYRmSLweuaew1CTy2uC34H5lWVUo7s2yJN/xI1wisD57rRLgw5vQoXOUAM8Jb3hz+YhfWZCJ+DwfMKeLMt7MqKeGuKKlD/uYU6S44OVD4kQP0saIKrcMjNXa9CinVNaA3ITsNOzkkn31mgbrGdMCT2UIHgc4SnYJSnTnw=="

// failOrFail 统一处理测试辅助失败：tb 为 nil（Example 等无测试上下文
// 场景）时直接 panic，否则走 t.Fatalf。
func failOrFail(tb testing.TB, msg string, err error) {
	if err == nil {
		return
	}
	if tb == nil {
		panic(msg + ": " + err.Error())
	}
	tb.Helper()
	tb.Fatalf("%s: %v", msg, err)
}

// NewAppCryptoSigner 返回解析后的测试应用私钥（crypto.Signer）。
// tb 允许为 nil（Example 函数等无测试上下文场景）。
func NewAppCryptoSigner(tb testing.TB) *rsa.PrivateKey {
	key, err := utils.ParsePrivateKeyPKCS8(utils.TrimKeyText(AppPrivateKeyPEM))
	failOrFail(tb, "parse app private key", err)
	return key
}

// NewAppSigner 返回使用测试应用私钥的签名器。
func NewAppSigner(tb testing.TB) *signers.RSA2 {
	signer, err := signers.NewRSA2(NewAppCryptoSigner(tb))
	if err != nil {
		tb.Fatalf("build app signer: %v", err)
	}
	return signer
}

// NewAlipayVerifier 返回使用测试支付宝公钥的验签器。
func NewAlipayVerifier(tb testing.TB) *verifiers.RSA2 {
	verifier, err := verifiers.NewRSA2FromPEM(AlipayPublicKeyPEM)
	if err != nil {
		tb.Fatalf("build alipay verifier: %v", err)
	}
	return verifier
}

// NewAlipayCryptoSigner 返回解析后的测试网关私钥（crypto.Signer），
// 模拟支付宝网关持有的私钥。
func NewAlipayCryptoSigner(tb testing.TB) *rsa.PrivateKey {
	key, err := utils.ParsePrivateKeyPKCS8(utils.TrimKeyText(AlipayPrivateKeyPEM))
	failOrFail(tb, "parse alipay private key", err)
	return key
}

// SignWithAlipayKey 使用模拟网关私钥对待验签数据生成 base64 签名。
func SignWithAlipayKey(data []byte) (string, error) {
	key, err := utils.ParsePrivateKeyPKCS8(utils.TrimKeyText(AlipayPrivateKeyPEM))
	if err != nil {
		return "", err
	}
	signer, err := signers.NewRSA2(key)
	if err != nil {
		return "", err
	}
	return signer.Sign(data)
}

// SignForm 按协议规则对待发送表单重新构造待签名串（测试辅助，
// 独立于 internal/protocol 实现，用于服务端交叉核对）。
func SignForm(form url.Values) string {
	keys := make([]string, 0, len(form))
	for key, vs := range form {
		if key == "sign" || len(vs) == 0 || vs[0] == "" {
			continue
		}
		keys = append(keys, key)
	}
	slices.Sort(keys)
	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		parts = append(parts, key+"="+form.Get(key))
	}
	return strings.Join(parts, "&")
}

// NotifySignContent 独立构造异步通知的待验签串（测试辅助，
// 独立于 internal/protocol 实现，用于交叉核对）：
// 剔除 sign 与 sign_type，保留空值参数，按参数名升序 k=v 用 & 连接。
func NotifySignContent(form url.Values) string {
	keys := make([]string, 0, len(form))
	for key, vs := range form {
		if key == "sign" || key == "sign_type" || len(vs) == 0 {
			continue
		}
		keys = append(keys, key)
	}
	slices.Sort(keys)
	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		parts = append(parts, key+"="+form.Get(key))
	}
	return strings.Join(parts, "&")
}

// SignNotify 使用模拟网关私钥对通知待验签串签名，
// 返回填充 sign 与 sign_type 后的表单副本。
func SignNotify(tb testing.TB, form url.Values) url.Values {
	tb.Helper()
	sign, err := SignWithAlipayKey([]byte(NotifySignContent(form)))
	if err != nil {
		tb.Fatalf("sign notify form: %v", err)
	}
	out := compat.CloneValues(form)
	out.Set("sign_type", "RSA2")
	out.Set("sign", sign)
	return out
}

// VerifyWithAppKey 使用测试应用公钥校验签名，模拟网关侧对请求签名的核对。
func VerifyWithAppKey(data []byte, sign string) error {
	publicKey, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(AppPublicKeyPEM))
	if err != nil {
		return err
	}
	signature, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	digest := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, digest[:], signature)
}

// SignWithAppKey 使用模拟商户私钥对待签名数据生成 base64 签名。
func SignWithAppKey(data []byte) (string, error) {
	key, err := utils.ParsePrivateKeyPKCS8(utils.TrimKeyText(AppPrivateKeyPEM))
	if err != nil {
		return "", err
	}
	signer, err := signers.NewRSA2(key)
	if err != nil {
		return "", err
	}
	return signer.Sign(data)
}

// V3ResponseHeaders 构造带 V3 响应签名头的响应头副本：
// 验签原文为 timestamp、nonce 与 body 三行（每行以 \n 结尾），
// 由模拟网关私钥签名。certSN 非空时追加 alipay-sn 头（证书模式）。
func V3ResponseHeaders(tb testing.TB, body []byte, timestamp, nonce, certSN string) http.Header {
	tb.Helper()
	content := timestamp + "\n" + nonce + "\n" + string(body) + "\n"
	sign, err := SignWithAlipayKey([]byte(content))
	if err != nil {
		tb.Fatalf("sign v3 response: %v", err)
	}
	header := http.Header{}
	header.Set("alipay-signature", sign)
	header.Set("alipay-timestamp", timestamp)
	header.Set("alipay-nonce", nonce)
	if certSN != "" {
		header.Set("alipay-sn", certSN)
	}
	return header
}

// V3ResponseHeadersWithKey 使用指定私钥构造带 V3 响应签名头的响应头，
// certSN 非空时追加 alipay-sn 头；供证书轮换等场景使用。
func V3ResponseHeadersWithKey(tb testing.TB, body []byte, timestamp, nonce string, key *rsa.PrivateKey, certSN string) http.Header {
	tb.Helper()
	content := timestamp + "\n" + nonce + "\n" + string(body) + "\n"
	sign, err := SignWithKey([]byte(content), key)
	if err != nil {
		tb.Fatalf("sign v3 response with key: %v", err)
	}
	header := http.Header{}
	header.Set("alipay-signature", sign)
	header.Set("alipay-timestamp", timestamp)
	header.Set("alipay-nonce", nonce)
	if certSN != "" {
		header.Set("alipay-sn", certSN)
	}
	return header
}

// V3ResponseHeadersWithAppKey 构造带请求方向签名的响应头，
// 供测试错钥场景使用：签名来自商户私钥而非网关私钥。
func V3ResponseHeadersWithAppKey(tb testing.TB, body []byte, timestamp, nonce string) http.Header {
	tb.Helper()
	content := timestamp + "\n" + nonce + "\n" + string(body) + "\n"
	sign, err := SignWithAppKey([]byte(content))
	if err != nil {
		tb.Fatalf("sign v3 response with app key: %v", err)
	}
	header := http.Header{}
	header.Set("alipay-signature", sign)
	header.Set("alipay-timestamp", timestamp)
	header.Set("alipay-nonce", nonce)
	return header
}

// ReadAllBody 读取并关闭请求体，供测试网关侧核对签名使用。
func ReadAllBody(tb testing.TB, r *http.Request) []byte {
	tb.Helper()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		tb.Fatalf("read request body: %v", err)
	}
	_ = r.Body.Close()
	return body
}

// VerifyV3Authorization 在测试网关侧独立重建 V3 请求签名原文并核对：
// 原文 = authString、HTTP 方法、请求 URI、请求体（空则空行）、
// 可选的代调用令牌，每行以 \n 结尾。实现独立于 internal/protocol，
// 用于交叉核对 SDK 的签名行为。
func VerifyV3Authorization(tb testing.TB, r *http.Request, body []byte, appAuthToken string) {
	tb.Helper()
	const scheme = "ALIPAY-SHA256withRSA"
	raw := r.Header.Get("Authorization")
	schemeGot, rest, ok := strings.Cut(raw, " ")
	if !ok || schemeGot != scheme {
		tb.Fatalf("unexpected Authorization scheme: %q", raw)
	}
	authString := rest
	for _, part := range strings.Split(rest, ",") {
		if strings.HasPrefix(part, "sign=") {
			authString = strings.TrimSuffix(rest, ","+part)
			signature := strings.TrimPrefix(part, "sign=")
			content := authString + "\n" + r.Method + "\n" + r.RequestURI + "\n" + string(body) + "\n"
			if appAuthToken != "" {
				content += appAuthToken + "\n"
			}
			if err := VerifyWithAppKey([]byte(content), signature); err != nil {
				tb.Fatalf("request signature mismatch: %v\ncontent: %q", err, content)
			}
			return
		}
	}
	tb.Fatalf("Authorization header has no sign field: %q", raw)
}
