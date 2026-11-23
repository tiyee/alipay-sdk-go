package auth

// CertificateIdentity 是证书模式的只读身份标识：
// 请求公共参数所需的证书 SN 均按官方规则计算
// （https://opendocs.alipay.com/common/02kdne.md），
// 不包含私钥或可信材料的可变引用。
type CertificateIdentity struct {
	// AppCertificateSN 是应用公钥证书序列号（app_cert_sn）。
	AppCertificateSN string
	// RootCertificateSN 是支付宝根证书序列号（alipay_root_cert_sn）。
	RootCertificateSN string
}

// CertificateProvider 是证书模式凭证的扩展能力接口：
// 在 Credentials 基础上提供证书身份元数据。
// Client 构造时若认证模式为证书模式，必须同时具备该能力；
// PublicKeyCredentials 只实现 Credentials。
type CertificateProvider interface {
	Credentials
	// CertificateIdentity 返回身份标识副本；修改副本不影响凭证。
	CertificateIdentity() CertificateIdentity
}
