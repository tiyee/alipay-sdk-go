package auth

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"time"

	"github.com/tiyee/alipay-sdk-go/internal/protocol"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// CertificateConfig 是本地证书模式的凭证配置。全部材料为解析后的
// X.509 证书（可经 utils.ParseCertificatePEM / ParseCertificatesPEM
// 从 PEM 文本解析，根证书包使用 utils.ParseRootCertificatesPEM）；应用私钥以标准库 crypto.Signer 传入。
// 配置不提供“失败后改用公钥”的备用项。
type CertificateConfig struct {
	// AppSigner 是应用私钥对应的签名者，必须与应用证书公钥匹配。
	AppSigner crypto.Signer
	// AppCertificate 是应用公钥证书（单张）。
	AppCertificate *x509.Certificate
	// AlipayCertificates 至少包含一张平台验签叶证书；可附带签发它的中间 CA。
	// 轮换时可同时配置新旧叶证书，只有叶证书会按响应标识选为验签密钥。
	AlipayCertificates []*x509.Certificate
	// RootCertificates 是支付宝根证书集合（可包含多张）。
	RootCertificates []*x509.Certificate
}

// CertificateCredentials 是本地证书模式的 Credentials 实现：
// 与应用证书匹配的应用签名器 + 按证书标识选择可信证书的验签器，
// 并实现 CertificateProvider 提供请求公共参数所需身份。
// 实例构造后内部状态不变，可并发复用。
type CertificateCredentials struct {
	signer   Signer
	verifier Verifier
	identity CertificateIdentity
}

// NewCertificateCredentials 构造证书凭证。装配时执行以下校验，
// 失败返回错误且不降级为其他认证模式：
//   - 证书可解析且为 RSA 密钥；
//   - 应用证书、支付宝证书与每张根证书在当前时刻处于有效期内；
//   - 应用私钥与应用证书公钥匹配；
//   - 每张根证书为自签发 CA；
//   - 每张平台验签叶证书可通过本地提供的中间 CA 或直接链至配置的根证书；
//   - 根证书集合中至少存在一张 RSA 签名算法证书，否则无法计算
//     alipay_root_cert_sn（官方规则仅统计 RSA 家族签名算法的证书）。
//
// 构造不下载证书、不访问网络；更新证书时用新配置重建凭证与 Client。
func NewCertificateCredentials(config CertificateConfig) (*CertificateCredentials, error) {
	if isNilInterface(config.AppSigner) {
		return nil, nilCredentialsErr("app signer")
	}
	appSigner := config.AppSigner
	if config.AppCertificate == nil {
		return nil, nilCredentialsErr("app certificate")
	}
	if len(config.AlipayCertificates) == 0 {
		return nil, nilCredentialsErr("alipay certificates")
	}
	if len(config.RootCertificates) == 0 {
		return nil, nilCredentialsErr("alipay root certificates")
	}
	now := time.Now()

	if _, ok := config.AppCertificate.PublicKey.(*rsa.PublicKey); !ok {
		return nil, fmt.Errorf("alipay/auth: app certificate public key is not RSA")
	}
	roots := x509.NewCertPool()
	for _, root := range config.RootCertificates {
		if root == nil {
			return nil, nilCredentialsErr("alipay root certificate")
		}
		if err := utils.CheckCertValidity(root, now); err != nil {
			return nil, fmt.Errorf("alipay/auth: check alipay root certificate validity: %w", err)
		}
		if !root.IsCA || !root.BasicConstraintsValid || (root.KeyUsage != 0 && root.KeyUsage&x509.KeyUsageCertSign == 0) {
			return nil, fmt.Errorf("alipay/auth: alipay root certificate is not a CA certificate")
		}
		if !bytes.Equal(root.RawIssuer, root.RawSubject) {
			return nil, fmt.Errorf("alipay/auth: root certificate is not self-issued")
		}
		// 对显式配置的信任锚检查自签名；平台叶证书单独验证完整证书链。
		if err := root.CheckSignature(root.SignatureAlgorithm, root.RawTBSCertificate, root.Signature); err != nil {
			return nil, fmt.Errorf("alipay/auth: alipay root certificate is not self-signed: %w", err)
		}
		roots.AddCert(root)
	}
	if err := utils.CheckCertValidity(config.AppCertificate, now); err != nil {
		return nil, fmt.Errorf("alipay/auth: check app certificate validity: %w", err)
	}
	var leaves []*x509.Certificate
	intermediates := x509.NewCertPool()
	for _, cert := range config.AlipayCertificates {
		if cert == nil {
			return nil, nilCredentialsErr("alipay certificate")
		}
		if err := utils.CheckCertValidity(cert, now); err != nil {
			return nil, fmt.Errorf("alipay/auth: check alipay certificate validity: %w", err)
		}
		if cert.IsCA {
			intermediates.AddCert(cert)
			continue
		}
		if _, ok := cert.PublicKey.(*rsa.PublicKey); !ok {
			return nil, fmt.Errorf("alipay/auth: alipay certificate public key is not RSA")
		}
		if cert.KeyUsage != 0 && cert.KeyUsage&x509.KeyUsageDigitalSignature == 0 {
			return nil, fmt.Errorf("alipay/auth: alipay certificate does not allow digital signatures")
		}
		leaves = append(leaves, cert)
	}
	for _, leaf := range leaves {
		_, err := leaf.Verify(x509.VerifyOptions{
			Roots: roots, Intermediates: intermediates, CurrentTime: now,
			KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		})
		if err != nil {
			return nil, fmt.Errorf("alipay/auth: verify alipay certificate chain (include any issuing intermediate CA): %w", err)
		}
	}

	privateKey, ok := appSigner.Public().(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("alipay/auth: app signer public key is %T, want RSA", appSigner.Public())
	}
	certPublicKey, ok := config.AppCertificate.PublicKey.(*rsa.PublicKey)
	if !ok || certPublicKey.N.Cmp(privateKey.N) != 0 || certPublicKey.E != privateKey.E {
		return nil, fmt.Errorf("alipay/auth: app signer does not match the app certificate public key")
	}

	rootCertSN, err := protocol.RootCertSN(config.RootCertificates)
	if err != nil {
		return nil, fmt.Errorf("alipay/auth: compute alipay root certificate SN: %w", err)
	}
	if rootCertSN == "" {
		return nil, fmt.Errorf("alipay/auth: no RSA-signed certificate in the root certificate set; alipay_root_cert_sn cannot be computed")
	}
	appCertSN, err := protocol.CertSN(config.AppCertificate)
	if err != nil {
		return nil, fmt.Errorf("alipay/auth: compute app certificate SN: %w", err)
	}
	set, err := newTrustedCertificateSet(leaves)
	if err != nil {
		return nil, err
	}
	signer, err := NewRSA2Signer(appSigner)
	if err != nil {
		return nil, fmt.Errorf("alipay/auth: build app signer: %w", err)
	}

	return &CertificateCredentials{
		signer:   signer,
		verifier: certificateVerifier{set: set},
		identity: CertificateIdentity{
			AppCertificateSN:  appCertSN,
			RootCertificateSN: rootCertSN,
		},
	}, nil
}

// Mode 实现 Credentials，返回 ModeCertificate。
func (c *CertificateCredentials) Mode() Mode { return ModeCertificate }

// Signer 实现 Credentials。
func (c *CertificateCredentials) Signer() Signer { return c.signer }

// Verifier 实现 Credentials：按验签输入中的证书标识选择可信证书。
func (c *CertificateCredentials) Verifier() Verifier { return c.verifier }

// CertificateIdentity 实现 CertificateProvider，返回身份标识副本。
func (c *CertificateCredentials) CertificateIdentity() CertificateIdentity {
	return c.identity
}

// Validate 实现 Credentials：确认组合对象与身份标识仍然完整。
func (c *CertificateCredentials) Validate() error {
	if c.signer == nil || isNilInterface(c.signer) {
		return nilCredentialsErr("app signer")
	}
	if c.verifier == nil || isNilInterface(c.verifier) {
		return nilCredentialsErr("alipay verifier")
	}
	if c.identity.AppCertificateSN == "" || c.identity.RootCertificateSN == "" {
		return fmt.Errorf("alipay/auth: certificate identity is incomplete")
	}
	if c.signer.Algorithm() != SignatureAlgorithmRSA2 {
		return fmt.Errorf("alipay/auth: unsupported signer algorithm %q", c.signer.Algorithm())
	}
	return nil
}

// certificateVerifier 将可信证书索引适配为 Verifier 契约。
type certificateVerifier struct {
	set *trustedCertificateSet
}

// Verify 实现 Verifier：算法必须与 RSA2 一致；KeyID 来自协议提取的
// 可信材料标识（V3 REST 为响应头 alipay-sn）；协议未提供标识时
// 由索引按本地可信材料语义处理。
func (v certificateVerifier) Verify(ctx context.Context, input VerificationInput) error {
	if input.Signature.Algorithm != SignatureAlgorithmRSA2 {
		return fmt.Errorf("alipay/auth: unsupported signature algorithm %q; only %q is supported",
			input.Signature.Algorithm, SignatureAlgorithmRSA2)
	}
	return v.set.verify(input.Payload, input.Signature.Value, input.KeyID)
}
