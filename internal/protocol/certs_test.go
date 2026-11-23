package protocol

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"math/big"
	"strings"
	"testing"
	"time"
)

// 证书 PEM 与预期 SN 为独立已知答案向量：证书由 openssl CLI 生成（固定序列号），
// 预期 SN 由 openssl -nameopt RFC2253、-serial 与系统 md5 命令独立计算，
// 不经过本包实现，用于交叉验证 SN 规则（issuer RFC2253 + 十进制序列号 + MD5）。
const (
	knownRoot1PEM = `-----BEGIN CERTIFICATE-----
MIIDUzCCAjugAwIBAgIBETANBgkqhkiG9w0BAQsFADBSMQswCQYDVQQGEwJDTjEb
MBkGA1UECgwSQWxpcGF5IEdvIFNESyBUZXN0MSYwJAYDVQQDDB1BbGlwYXkgR28g
U0RLIFRlc3QgUm9vdCBDQSBSMTAgFw0yNjExMDQxOTQxMDVaGA8yMTI2MTAxMTE5
NDEwNVowUjELMAkGA1UEBhMCQ04xGzAZBgNVBAoMEkFsaXBheSBHbyBTREsgVGVz
dDEmMCQGA1UEAwwdQWxpcGF5IEdvIFNESyBUZXN0IFJvb3QgQ0EgUjEwggEiMA0G
CSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCYcI1INI/ZsfFHKfdiyDTLm+JlBLOf
elpa5kT9Zjn9DiA1eAS3MLfFuUe27GNBz5lRJZS/THj8gifwpjiukzG2fqAwn2ho
TvyNcegKphQPvfZlDQlnuOyzWkdAf8V0X5kiNpGIoEZf9plEhyMWYlmkYDKo/XYy
rZg0uZ4r95pOwfwnT9HPOV0k6alz+27QGYWDw76mGg157GQamoqCjniV6d0BQKd8
XNjXiI4UTuSssPr0uUAaK4ffPgSbDEhi8PLfeTSarqXDCoAhO3udz0/ot9D2aKtd
0zEYQiGtLgoQbV6X73xjj5H0b58eUJALhsEz7V7ptFwGd7MwMsbnKLPvAgMBAAGj
MjAwMB0GA1UdDgQWBBTrrAMDWC8yo2nUtatJq2qfSklGEjAPBgNVHRMBAf8EBTAD
AQH/MA0GCSqGSIb3DQEBCwUAA4IBAQBctKW7nTV4CtYjkBNztpDO6IeOpKI4t7c7
KcM/Ns2pstL4guLj2bnDe0MzlLni7dVKrPweYFZhDNrlBZKHM/U032yoJ1Z857bp
dCQq+F/KNV5WuOAOTVn8xphvI4pTHN2TV70Q4tUgklyGN5a3+kIChRIUjkU3b5Re
Brtjuu+IapGCuvINwAZBY0WjzCxRtz9stVjlRMLnNzFQEbg74dKpgotpKVa/EbDF
KGq887ovVrA+PL/goxJAEXMagrALvdtbnh0f0Q0qluUpua/zSXO8CnehxQtVcyIf
cxGXHhIJVXpuPKbEN8wHFFXHMcwt+JJm/rcF1kXba6ZopctPVaN+
-----END CERTIFICATE-----`

	knownRoot2PEM = `-----BEGIN CERTIFICATE-----
MIIDUzCCAjugAwIBAgIBIjANBgkqhkiG9w0BAQsFADBSMQswCQYDVQQGEwJDTjEb
MBkGA1UECgwSQWxpcGF5IEdvIFNESyBUZXN0MSYwJAYDVQQDDB1BbGlwYXkgR28g
U0RLIFRlc3QgUm9vdCBDQSBSMjAgFw0yNjExMDQxOTQxMDVaGA8yMTI2MTAxMTE5
NDEwNVowUjELMAkGA1UEBhMCQ04xGzAZBgNVBAoMEkFsaXBheSBHbyBTREsgVGVz
dDEmMCQGA1UEAwwdQWxpcGF5IEdvIFNESyBUZXN0IFJvb3QgQ0EgUjIwggEiMA0G
CSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQC4iwBKFCRvcvhBTj4RHdvg9MAtYO4O
noBaYScJ8k3Q3Q0ZEtMZRcyzTAVXgeD8meOQsFn1LNG2t5oGfKzIQfpr4CFylEZT
m7t0srgrffc7rpUowWX56iX7bYU16DMV5Wyt1QDlQ+o30z2rNflxUpT7KpEKo30B
0qLl1QtRbhZGeYfsblYqt2ca/k/9bUpr17h2aBmEIfnZ174r2SdujiU2lMGRua19
cWlQZ416jYGpu7skuTLKRFB7IEZFtqk4nfX1EijwtFym2iaFGnlEnhf0Yf7lv+r/
WBOLA+WlHd5P/aUSWYdAxDyXHWAjs80sJCmMIDYLs6E/FhrOWkGTiPPzAgMBAAGj
MjAwMB0GA1UdDgQWBBR2GdAPTN5iEYrixWB954dECkHluTAPBgNVHRMBAf8EBTAD
AQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCeEZLujzf7e4c9R8sOMG17w3MPpSjzfk5c
qUpIBRk0WpBkhA28O/MjtDizqjqw/4fMJXNLCnLILPPmNbzNIokKN+SAsGHEuGuv
ykMrnRSJES5urvjQ+npaUM/TW5+K1YUAPAXel2/W56SRfqWmKtKzwdncLhODyyZd
56S+sQCniElxcTUwJOIykwXdRSNcV4Rf00rs3kLv9zRFh9Z9wfIVrvCfk4f3SjE/
YH+w6Wrp05pr6HnY6sIb03bZgjHTd/UnhbRFLWAN1WLcCCX7I2/YvyURYhOHagXX
jYWG2CXN+qDizSHcNqPl7krQ/FeoEKYlx05Ox6owXfk96TcLa+V9
-----END CERTIFICATE-----`

	knownAlipayCertPEM = `-----BEGIN CERTIFICATE-----
MIIDZDCCAkygAwIBAgIBMzANBgkqhkiG9w0BAQsFADBSMQswCQYDVQQGEwJDTjEb
MBkGA1UECgwSQWxpcGF5IEdvIFNESyBUZXN0MSYwJAYDVQQDDB1BbGlwYXkgR28g
U0RLIFRlc3QgUm9vdCBDQSBSMTAgFw0yNjExMDQxOTQxMDVaGA8yMTI2MTAxMTE5
NDEwNVowUzELMAkGA1UEBhMCQ04xGzAZBgNVBAoMEkFsaXBheSBHbyBTREsgVGVz
dDEnMCUGA1UEAwweQWxpcGF5IEdvIFNESyBUZXN0IFB1YmxpYyBDZXJ0MIIBIjAN
BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1MWJyScDaNy2Z6frFLXTm6EOu8QN
koINT+G4cyEViwi6h+zfVaiuYGtXl3THGOhlmRVIKjqPjbkp41BZ4H/+Qxz3i4p+
kWB6Apt3R0hP2ozHWwgVSmYQ3HxAJs9u1SP7TkfP/rMex4/9+uYRA+AxsiJE/+aB
GQ7WtqnFWVHB2V1a6yFZEByq2SY4l4hcj2HlpTsom807VfDmC6bK3uluZ3M3GnvL
tHDa/0UQ5SESyvW3Qx2dAGcscCJ1YpFKl8SUfLH46TYXzRwknGTJAxEI5tSOxcU+
V7gftJY4ZhJn33Pm2R7YzUfmyyDDKCqL6ayn0M7j1dczfnq8l1BAWB8alwIDAQAB
o0IwQDAdBgNVHQ4EFgQU7/qSUdYGISSbdOjjo9ox5h3zjP8wHwYDVR0jBBgwFoAU
66wDA1gvMqNp1LWrSatqn0pJRhIwDQYJKoZIhvcNAQELBQADggEBAE0klWMQcW0x
7rVc5NIebZ/97SL+dcj3nboXo1g/xrFB2QzD2XP60gParSdmN7MX3jGMYCqSDnY1
zp1SAv5ME+8JgcWW/+RnbKGwuMLdU+8yAnhIt5z6cE71EIknp1cy5/bWbRatydEY
+zEQk+9B3Kqbri5JPWL6L7m0FFwLXzPSGkglNt6tPO9nEpAkj8lmhsAyt9Oabj4k
xuYdnwszCPWx1JN/NOGSmFDGLlpdPBJ4OwzUjUGCDNpB2BvxUzweckdO1sO1xeMI
dXuW9LBZtiHlssQRyPLHT/lo29Y2OSCMhctsnG5SWhPp/qY6UP0Wzu8RTQpmIxzA
oYghhiMA5v0=
-----END CERTIFICATE-----`

	knownAppCertPEM = `-----BEGIN CERTIFICATE-----
MIIDYTCCAkmgAwIBAgIBRDANBgkqhkiG9w0BAQsFADBSMQswCQYDVQQGEwJDTjEb
MBkGA1UECgwSQWxpcGF5IEdvIFNESyBUZXN0MSYwJAYDVQQDDB1BbGlwYXkgR28g
U0RLIFRlc3QgUm9vdCBDQSBSMTAgFw0yNjExMDQxOTQxMDVaGA8yMTI2MTAxMTE5
NDEwNVowUDELMAkGA1UEBhMCQ04xGzAZBgNVBAoMEkFsaXBheSBHbyBTREsgVGVz
dDEkMCIGA1UEAwwbQWxpcGF5IEdvIFNESyBUZXN0IEFwcCBDZXJ0MIIBIjANBgkq
hkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvVxNDtSZ9dSYxMPKTAbt68mjkhI1kBld
o7aImQX2wP7eyC8qorQepCbqgRVGYUFgQpZmLH/bElLsbW1/QW2dbJLmxmaexMzD
26EnW8AnAepJiMRyBaoAoky91k7BarOjP1Qs203L9NrBK1zPTaew/bOTVouKe+Xj
iMwo976ppqq8sRYl803YQHO90MAjTJXXZ2/q6u0a92lzDw1jPT0u7h6Ew3hJwoEE
xirO4hKRBFzlw21EOK+tu0H9jyC+ldS9JhM27exvM/wPTr6nI2dxH03fpJVK4jP+
0nUBm9YrlcUCizOns/72WfwBKrZZ51PCg3f47u7ROtguZntBx/V2cwIDAQABo0Iw
QDAdBgNVHQ4EFgQUivrtJqjMbCSWzhwNXDLfst9QMVkwHwYDVR0jBBgwFoAU66wD
A1gvMqNp1LWrSatqn0pJRhIwDQYJKoZIhvcNAQELBQADggEBAEE3XdkJnxYsYn4m
WnJrDeYnRxSR4mKy0eGH1UZE4BQCnCrc4oNyGtjJR9i2ZVhxCANNmyZDCfGH244Y
2t/k8Y4d7FJpGO3MN8akBoSjd2n9rirDMMPz9Ha4+kArQFFbc0+LvLS/YECJcjQ/
LwNuH/Moog8HzIZJTIP5kKOslTpOlRdhKj7eLdwgvTd2PdJv6EMsewF4FvDOpjhE
9WwTFQk9RdKhS4ozEC80i5Szc/w6JLvc+OKUVe3Kwb+Ul2bb3PQ+aWnaDU56+fWX
4i28l51f3936ZewqqSIRfTSNvqYh0FMM9V9zWeE2ewfMIxf4PgZthCExQjIMbpsm
PH43Deg=
-----END CERTIFICATE-----`
)

// 独立预期值：SN = MD5(issuer RFC2253 + 内置序列号十进制)，openssl 交叉计算。
const (
	wantRoot1SN  = "b6ae3dff5b08b28504175777c09e7cec"
	wantRoot2SN  = "946f69fb1774eda04b72ebba65c19618"
	wantAlipaySN = "e6cee1c7a680f747e32a4d633b208766"
	wantAppSN    = "3cd9e46fe2bbfe9c07861f72fe69b1b7"
)

// knownIssuerRFC2253 是四张测试证书共同的签发机构 RFC2253 形式
// （openssl -nameopt RFC2253 独立输出）。
const knownIssuerRFC2253 = "CN=Alipay Go SDK Test Root CA R1,O=Alipay Go SDK Test,C=CN"

func parseTestCert(t *testing.T, pemData string) *x509.Certificate {
	t.Helper()
	block, _ := pem.Decode([]byte(pemData))
	if block == nil || block.Type != "CERTIFICATE" {
		t.Fatalf("no certificate block")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatalf("parse certificate: %v", err)
	}
	return cert
}

func TestCertSNMatchesIndependentKnownAnswer(t *testing.T) {
	cases := []struct {
		name string
		pem  string
		want string
	}{
		{"root1", knownRoot1PEM, wantRoot1SN},
		{"root2", knownRoot2PEM, wantRoot2SN},
		{"alipay", knownAlipayCertPEM, wantAlipaySN},
		{"app", knownAppCertPEM, wantAppSN},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cert := parseTestCert(t, tc.pem)
			sn, err := CertSN(cert)
			if err != nil {
				t.Fatalf("CertSN() error = %v", err)
			}
			if sn != tc.want {
				t.Fatalf("CertSN() = %q, want %q", sn, tc.want)
			}
		})
	}
}

func TestCertSNIssuerFormatMatchesOpenSSL(t *testing.T) {
	cert := parseTestCert(t, knownRoot1PEM)
	issuer, err := rfc2253IssuerName(cert)
	if err != nil {
		t.Fatalf("rfc2253IssuerName() error = %v", err)
	}
	if issuer != knownIssuerRFC2253 {
		t.Fatalf("issuer = %q, want %q", issuer, knownIssuerRFC2253)
	}
}

func TestRootCertSNJoinsInFileOrder(t *testing.T) {
	root1 := parseTestCert(t, knownRoot1PEM)
	root2 := parseTestCert(t, knownRoot2PEM)

	sn, err := RootCertSN([]*x509.Certificate{root1, root2})
	if err != nil {
		t.Fatalf("RootCertSN() error = %v", err)
	}
	if sn != wantRoot1SN+"_"+wantRoot2SN {
		t.Fatalf("RootCertSN() = %q", sn)
	}

	// 顺序跟随输入（与官方 SDK 按文件顺序拼接一致）。
	reversed, err := RootCertSN([]*x509.Certificate{root2, root1})
	if err != nil {
		t.Fatalf("RootCertSN() reversed error = %v", err)
	}
	if reversed != wantRoot2SN+"_"+wantRoot1SN {
		t.Fatalf("reversed RootCertSN() = %q", reversed)
	}

	// 单张根证书不加下划线。
	single, err := RootCertSN([]*x509.Certificate{root1})
	if err != nil {
		t.Fatalf("RootCertSN() single error = %v", err)
	}
	if single != wantRoot1SN {
		t.Fatalf("single RootCertSN() = %q, want %q", single, wantRoot1SN)
	}

	// 空集合返回空串。
	empty, err := RootCertSN(nil)
	if err != nil {
		t.Fatalf("RootCertSN(nil) error = %v", err)
	}
	if empty != "" {
		t.Fatalf("RootCertSN(nil) = %q, want empty", empty)
	}
}

func TestRootCertSNSkipsNonRSASignatureCerts(t *testing.T) {
	// 官方 getRootCertSN 仅保留 RSA 家族签名算法的证书；
	// ECDSA 签名算法（OID 1.2.840.10045.4.*）不在前缀内，应被跳过。
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("generate ecdsa key: %v", err)
	}
	template := &x509.Certificate{
		SerialNumber: big.NewInt(77),
		Subject:      pkix.Name{CommonName: "ECDSA Root", Organization: []string{"Alipay Go SDK Test"}, Country: []string{"CN"}},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("create ecdsa certificate: %v", err)
	}
	ecdsaCert, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatalf("parse ecdsa certificate: %v", err)
	}
	oid, err := signatureAlgorithmOID(ecdsaCert.Raw)
	if err != nil {
		t.Fatalf("signatureAlgorithmOID() error = %v", err)
	}
	if strings.HasPrefix(oid.String(), "1.2.840.113549.1.1.") {
		t.Fatalf("expected a non-RSA signature OID, got %q", oid.String())
	}

	root1 := parseTestCert(t, knownRoot1PEM)
	sn, err := RootCertSN([]*x509.Certificate{ecdsaCert, root1})
	if err != nil {
		t.Fatalf("RootCertSN() error = %v", err)
	}
	if sn != wantRoot1SN {
		t.Fatalf("RootCertSN() = %q, want only the RSA root SN %q", sn, wantRoot1SN)
	}
}

func TestSignatureAlgorithmOID(t *testing.T) {
	cert := parseTestCert(t, knownRoot1PEM)
	oid, err := signatureAlgorithmOID(cert.Raw)
	if err != nil {
		t.Fatalf("signatureAlgorithmOID() error = %v", err)
	}
	if got := oid.String(); got != "1.2.840.113549.1.1.11" { // sha256WithRSAEncryption
		t.Fatalf("signature algorithm OID = %q", got)
	}
}

// buildSelfSigned 以固定序列号构造自签 RSA 证书，用于 RFC2253 边界用例。
func buildSelfSigned(t *testing.T, subject pkix.Name, serial int64) *x509.Certificate {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	template := &x509.Certificate{
		SerialNumber: big.NewInt(serial),
		Subject:      subject,
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("create certificate: %v", err)
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatalf("parse certificate: %v", err)
	}
	return cert
}

func TestRFC2253NameEscapesSpecialCharacters(t *testing.T) {
	// Go 按 C, O, OU, CN 顺序编码 subject；RFC2253 逆序输出为 CN, OU, O, C。
	cert := buildSelfSigned(t, pkix.Name{
		CommonName:         `a,b+c<d>e;f"g\h#i`,
		Organization:       []string{" 支付宝 "},
		Country:            []string{"CN"},
		OrganizationalUnit: []string{"测试单元"},
	}, 100)

	got, err := rfc2253IssuerName(cert)
	if err != nil {
		t.Fatalf("rfc2253IssuerName() error = %v", err)
	}
	// CN 值中的协议特殊字符逐个加反斜杠；组织名首尾空格加反斜杠；
	// 中文等非 ASCII 原样保留。
	want := `CN=a\,b\+c\<d\>e\;f\"g\\h\#i,OU=测试单元,O=\ 支付宝\ ,C=CN`
	if got != want {
		t.Fatalf("issuer = %q, want %q", got, want)
	}
	again, err := rfc2253IssuerName(cert)
	if err != nil {
		t.Fatalf("rfc2253IssuerName() second error = %v", err)
	}
	if again != want {
		t.Fatalf("issuer not stable: %q", again)
	}
}

func TestRFC2253NameUnknownOIDUsesDottedForm(t *testing.T) {
	// EmailAddress（1.2.840.113549.1.9.1）不在 RFC2253 关键字表内，
	// 按 Java 行为输出点分 OID，不带 "OID." 前缀。
	cert := buildSelfSigned(t, pkix.Name{
		CommonName: "cn",
		ExtraNames: []pkix.AttributeTypeAndValue{
			{Type: asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}, Value: "dev@example.com"},
		},
	}, 101)

	got, err := rfc2253IssuerName(cert)
	if err != nil {
		t.Fatalf("rfc2253IssuerName() error = %v", err)
	}
	// ExtraNames 编码在 CommonName 之后，逆序输出时位于最前。
	want := "1.2.840.113549.1.9.1=dev@example.com,CN=cn"
	if got != want {
		t.Fatalf("issuer = %q, want %q", got, want)
	}
}

func TestRFC2253NameMultiValuedRDNAndHexValue(t *testing.T) {
	// 直接构造 RDN 序列：多值 RDN 按原始顺序用 "+" 连接；
	// 非字符串类型的值输出 "#"+DER 十六进制（如 BOOLEAN 01 01 FF）。
	boolDER, err := asn1.Marshal(true)
	if err != nil {
		t.Fatalf("marshal boolean: %v", err)
	}
	seq := pkix.RDNSequence{
		{{Type: asn1.ObjectIdentifier{2, 5, 4, 6}, Value: "CN"}}, // C，最先编码
		{
			{Type: asn1.ObjectIdentifier{2, 5, 4, 10}, Value: "O Org"},
			{Type: asn1.ObjectIdentifier{2, 5, 4, 3}, Value: "CN Name"},
		},
		{{Type: asn1.ObjectIdentifier{2, 5, 4, 11}, Value: asn1.RawValue{
			Tag: 1, Bytes: boolDER[2:], FullBytes: boolDER,
		}}}, // OU，布尔值
	}
	got := RFC2253Name(seq)
	want := `OU=#0101ff,O=O Org+CN=CN Name,C=CN`
	if got != want {
		t.Fatalf("RFC2253Name() = %q, want %q", got, want)
	}
	if RFC2253Name(nil) != "" {
		t.Fatal("empty sequence should produce empty string")
	}
}
