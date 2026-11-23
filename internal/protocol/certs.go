// 本文件实现支付宝开放平台证书模式的公钥证书序列号（SN）计算规则。
// 依据：
//   - 官方常见问题：https://opendocs.alipay.com/common/02kdne.md
//     （SN = MD5(签发机构名称 + 证书内置序列号)，具体格式以官方 SDK 源码为准）；
//   - 以下格式细节已与官方 SDK 的 AntCertificationUtil 核对：
//     getCertSN 使用 X500Principal.getName()（RFC2253）拼接
//     BigInteger 十进制序列号后取 MD5，十六进制补零至 32 位；
//     getRootCertSN 仅保留 RSA 签名算法（OID 前缀 1.2.840.113549.1.1）的
//     根证书，逐张计算 SN 后按文件顺序用 "_" 拼接，不做整体 MD5。
//   - RFC2253 名称格式化与 OpenJDK X500Name/RDN/AVA 实现对齐：
//     RDN 逆序输出、多值 RDN 用 "+" 连接、标准关键字表
//     （CN/C/L/ST/O/OU/STREET/DC/UID），其余属性输出点分 OID；
//     值中 ",=+<>#;\"\'\\" 等特殊字符加反斜杠转义（详见 escapeRFC2253Value）。
package protocol

import (
	"crypto/md5"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"strings"
	"unicode/utf16"
)

// rsaSignatureOIDPrefix 是官方 getRootCertSN 使用的签名算法 OID 前缀，
// 覆盖 RSA PKCS#1 家族（SHA1/SHA256/SHA384/SHA512WithRSA 等）。
const rsaSignatureOIDPrefix = "1.2.840.113549.1.1"

// rfc2253Keywords 是 RFC2253 输出使用的标准属性类型关键字，
// 与 OpenJDK AVAKeyword 的 RFC2253 合规注册表一致。
// 其余属性类型按 RFC2253 输出点分 OID（不带前缀）。
var rfc2253Keywords = map[string]string{
	"2.5.4.3":                    "CN",
	"2.5.4.6":                    "C",
	"2.5.4.7":                    "L",
	"2.5.4.8":                    "ST",
	"2.5.4.9":                    "STREET",
	"2.5.4.10":                   "O",
	"2.5.4.11":                   "OU",
	"0.9.2342.19200300.100.1.25": "DC",
	"0.9.2342.19200300.100.1.1":  "UID",
}

// RFC2253Name 将 X.509 名称（RDN 序列）格式化为 Java
// X500Principal.getName()（RFC2253）的等价形式：
// 从序列最后一个 RDN 开始逆序输出，RDN 之间用 "," 连接，
// 多值 RDN 内部按原始顺序用 "+" 连接。
func RFC2253Name(seq pkix.RDNSequence) string {
	if len(seq) == 0 {
		return ""
	}
	parts := make([]string, 0, len(seq))
	for i := len(seq) - 1; i >= 0; i-- {
		rdn := seq[i]
		avas := make([]string, 0, len(rdn))
		for _, atv := range rdn {
			avas = append(avas, formatRFC2253AttributeTypeAndValue(atv))
		}
		parts = append(parts, strings.Join(avas, "+"))
	}
	return strings.Join(parts, ",")
}

// rfc2253IssuerName 从证书原始 DER 中解析签发机构名称并格式化为 RFC2253。
// 使用 RawIssuer 原始字节而非 pkix.Name，避免重新编码丢失属性顺序与类型。
func rfc2253IssuerName(cert *x509.Certificate) (string, error) {
	var seq pkix.RDNSequence
	if _, err := asn1.Unmarshal(cert.RawIssuer, &seq); err != nil {
		return "", fmt.Errorf("parse issuer rdn sequence: %w", err)
	}
	return RFC2253Name(seq), nil
}

func formatRFC2253AttributeTypeAndValue(atv pkix.AttributeTypeAndValue) string {
	keyword, ok := rfc2253Keywords[atv.Type.String()]
	if !ok {
		keyword = atv.Type.String()
	}
	return keyword + "=" + formatRFC2253Value(atv.Value)
}

// asRawValue 将属性值规整为 asn1.RawValue：解析证书时值已是 RawValue；
// 手工构造的序列可能携带具体类型，先编码再解析回 TLV。
func asRawValue(value any) (asn1.RawValue, error) {
	if raw, ok := value.(asn1.RawValue); ok {
		return raw, nil
	}
	der, err := asn1.Marshal(value)
	if err != nil {
		return asn1.RawValue{}, err
	}
	var raw asn1.RawValue
	if _, err := asn1.Unmarshal(der, &raw); err != nil {
		return asn1.RawValue{}, err
	}
	return raw, nil
}

// rfc2253StringTags 是 Java AVA.isDerString 认可的字符串类型标签：
// UTF8String、PrintableString、TeletexString、IA5String、BMPString。
// 其余类型的值按 "#"+DER 十六进制输出。
var rfc2253StringTags = map[int]struct{}{
	12: {}, // UTF8String
	19: {}, // PrintableString
	20: {}, // TeletexString
	22: {}, // IA5String
	30: {}, // BMPString
}

// formatRFC2253Value 按 RFC2253 第 2.4 节与 Java AVA 实现编码属性值：
// 字符串类型按对应字符集解码后转义；其他类型输出 "#"+DER 十六进制。
func formatRFC2253Value(value any) string {
	raw, err := asRawValue(value)
	if err != nil {
		// 无法编码的值按空字符串处理，避免输出未定义内容。
		return ""
	}
	if _, ok := rfc2253StringTags[raw.Tag]; !ok {
		return "#" + hex.EncodeToString(raw.FullBytes)
	}
	var text string
	switch raw.Tag {
	case 20: // TeletexString（T61），Java 使用 ISO-8859-1 解码
		runes := make([]rune, len(raw.Bytes))
		for i, b := range raw.Bytes {
			runes[i] = rune(b)
		}
		text = string(runes)
	case 30: // BMPString，Java 使用 UTF-16BE 解码
		if len(raw.Bytes)%2 != 0 {
			// 长度异常时保留原始字节，避免伪造可解码文本。
			return "#" + hex.EncodeToString(raw.FullBytes)
		}
		units := make([]uint16, len(raw.Bytes)/2)
		for i := range units {
			units[i] = uint16(raw.Bytes[2*i])<<8 | uint16(raw.Bytes[2*i+1])
		}
		text = string(utf16.Decode(units))
	default: // UTF8String / PrintableString / IA5String，字节即原文
		text = string(raw.Bytes)
	}
	return escapeRFC2253Value(text)
}

// escapeRFC2253Value 转义属性值中的协议转义集合（与 OpenJDK AVA 的
// escapees 一致：逗号、等号、加号、尖括号、井号、分号、双引号、反斜杠），
// 空字符输出 "\00"，随后对首尾空格与 \r 加反斜杠。
// 其余字符（含中文等非 ASCII）按原文保留。
func escapeRFC2253Value(text string) string {
	var b strings.Builder
	b.Grow(len(text))
	for _, r := range text {
		switch r {
		case ',', '=', '+', '<', '>', '#', ';', '"', '\\':
			b.WriteByte('\\')
			b.WriteRune(r)
		case 0:
			b.WriteString("\\00")
		default:
			b.WriteRune(r)
		}
	}
	escaped := b.String()
	// Java 在转义后再处理首尾空白（' ' 与 '\r'），逐字符加反斜杠。
	lead, trail := 0, len(escaped)-1
	for lead < len(escaped) && (escaped[lead] == ' ' || escaped[lead] == '\r') {
		lead++
	}
	for trail >= 0 && (escaped[trail] == ' ' || escaped[trail] == '\r') {
		trail--
	}
	if lead == 0 && trail == len(escaped)-1 {
		return escaped
	}
	b.Reset()
	for i := 0; i < len(escaped); i++ {
		if i < lead || i > trail {
			b.WriteByte('\\')
		}
		b.WriteByte(escaped[i])
	}
	return b.String()
}

// CertSN 按支付宝开放平台规则计算公钥证书序列号：
// 对「签发机构 RFC2253 名称 + 证书内置序列号十进制字符串」拼接后取 MD5，
// 以小写十六进制输出（等价于官方 SDK 的 BigInteger 十六进制 + 补零至 32 位）。
// 注意：该 SN 与 X.509 证书内置序列号不同。
func CertSN(cert *x509.Certificate) (string, error) {
	issuer, err := rfc2253IssuerName(cert)
	if err != nil {
		return "", err
	}
	sum := md5.Sum([]byte(issuer + cert.SerialNumber.String()))
	return hex.EncodeToString(sum[:]), nil
}

// RootCertSN 按官方规则计算支付宝根证书序列号：
// 仅保留 RSA 家族签名算法（OID 前缀 1.2.840.113549.1.1，与官方
// getRootCertSN 一致）的证书，逐张计算 CertSN 后按传入顺序用 "_" 拼接；
// 不做整体 MD5。没有 RSA 签名算法证书时返回空串。
func RootCertSN(certs []*x509.Certificate) (string, error) {
	sns := make([]string, 0, len(certs))
	for _, cert := range certs {
		rsaSigned, err := IsRSASignedCertificate(cert.Raw)
		if err != nil {
			return "", err
		}
		if !rsaSigned {
			continue
		}
		sn, err := CertSN(cert)
		if err != nil {
			return "", err
		}
		sns = append(sns, sn)
	}
	return strings.Join(sns, "_"), nil
}

// IsRSASignedCertificate 根据外层签名算法 OID 识别参与根证书 SN 的证书。
// 该操作可用于过滤标准库不支持的国密根证书，不赋予证书信任。
func IsRSASignedCertificate(der []byte) (bool, error) {
	oid, err := signatureAlgorithmOID(der)
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(oid.String(), rsaSignatureOIDPrefix+"."), nil
}

// signatureAlgorithmOID 从证书原始 DER 的外层 Certificate 结构中
// 提取 signatureAlgorithm 的算法 OID。Go 的 x509.Certificate 未导出
// 该 OID，直接解析原始字节避免按枚举名猜测。
func signatureAlgorithmOID(der []byte) (asn1.ObjectIdentifier, error) {
	var outer struct {
		TBS                asn1.RawValue
		SignatureAlgorithm pkix.AlgorithmIdentifier
		SignatureValue     asn1.BitString
	}
	rest, err := asn1.Unmarshal(der, &outer)
	if err != nil {
		return nil, fmt.Errorf("parse certificate signature algorithm: %w", err)
	}
	if len(rest) != 0 {
		return nil, fmt.Errorf("trailing data in certificate")
	}
	return outer.SignatureAlgorithm.Algorithm, nil
}
