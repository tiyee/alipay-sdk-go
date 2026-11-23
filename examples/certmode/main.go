// Command certmode 是本地证书模式的最小接入示例（V3 REST 接口）。
//
// 链路：装载证书材料（校验有效期、私钥匹配、根证书链）→
// 请求按 V3 契约携带 alipay-root-cert-sn 头并参与签名 →
// 响应按 alipay-sn 头选择本地可信证书验签 → 类型化业务结果。
//
// 配置全部来自环境变量，示例不包含任何真实凭证：
//
//	ALIPAY_APP_ID          本应用 APPID（必填）
//	ALIPAY_APP_CERT        应用公钥证书 PEM 文件路径（必填）
//	ALIPAY_APP_PRIVATE_KEY 应用私钥 PKCS#8 PEM 文件路径（必填）
//	ALIPAY_CERT            支付宝公钥证书 PEM 文件路径（必填；轮换过渡期可用逗号分隔多个路径）
//	ALIPAY_ROOT_CERT       支付宝根证书 PEM 文件路径（必填）
//	ALIPAY_ENV             production 或 sandbox（默认 production）
//	ALIPAY_OUT_TRADE_NO    要查询的商户订单号（必填）
//
// 证书轮换：更新本地证书文件后重建 Client 实例（配置不可热更新）；
// SDK 不自动下载证书，响应 alipay-sn 与本地证书不匹配时调用失败，
// 需从开放平台下载对应 SN 的证书并更新本地文件。
package main

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/services/trade"
	"github.com/tiyee/alipay-sdk-go/utils"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	appID := os.Getenv("ALIPAY_APP_ID")
	appCert := readRequiredFile(logger, "ALIPAY_APP_CERT")
	appPrivateKey := readRequiredFile(logger, "ALIPAY_APP_PRIVATE_KEY")
	rootCert := readRequiredFile(logger, "ALIPAY_ROOT_CERT")
	outTradeNo := os.Getenv("ALIPAY_OUT_TRADE_NO")
	if appID == "" || outTradeNo == "" {
		logger.Error("ALIPAY_APP_ID and ALIPAY_OUT_TRADE_NO are required")
		os.Exit(2)
	}
	alipayCerts, err := readAlipayCertificateFiles(os.Getenv("ALIPAY_CERT"))
	if err != nil {
		logger.Error("cannot load ALIPAY_CERT files")
		os.Exit(2)
	}

	// 凭证构造完成全部证书校验（有效期、私钥匹配、根证书链、SN 计算）；
	// ALIPAY_CERT 可传多张证书（轮换过渡期），按响应 alipay-sn 头选择。
	credentials, err := buildCertificateCredentials(appCert, appPrivateKey, rootCert, alipayCerts)
	if err != nil {
		// 证书缺失、过期、私钥不匹配或根证书链不通过都会在此显式失败，
		// SDK 不会降级为其他认证模式。
		logger.Error("invalid certificate configuration")
		os.Exit(2)
	}

	opts := []core.ClientOption{
		option.WithAppID(appID),
	}
	switch os.Getenv("ALIPAY_ENV") {
	case "", "production":
		opts = append(opts, option.WithEnvironment(core.EnvironmentProduction))
	case "sandbox":
		opts = append(opts, option.WithEnvironment(core.EnvironmentSandbox))
	default:
		logger.Error("ALIPAY_ENV must be production or sandbox")
		os.Exit(2)
	}

	client, err := core.NewClient(credentials, opts...)
	if err != nil {
		logger.Error("invalid client configuration")
		os.Exit(2)
	}

	service, err := trade.NewService(client)
	if err != nil {
		logger.Error("build trade service")
		os.Exit(2)
	}

	resp, err := service.Query(context.Background(), trade.QueryRequest{
		OutTradeNo: outTradeNo,
	})
	if err != nil {
		// 运行期间证书过期或响应 alipay-sn 缺失、未知均明确报错。
		// 有效期错误可沿错误链定位证书；更新材料后重建实例。
		var validityErr *auth.CertValidityError
		var certErr *core.CertificateError
		var apiErr *core.APIError
		switch {
		case errors.As(err, &validityErr):
			logger.LogAttrs(context.Background(), slog.LevelError,
				"certificate is outside its validity period; update certificates and rebuild client",
				slog.String("field", validityErr.Field),
				slog.Time("not_before", validityErr.NotBefore),
				slog.Time("not_after", validityErr.NotAfter))
		case errors.As(err, &certErr):
			logger.Error("certificate check failed; check response SN and local certificates")
		case errors.As(err, &apiErr):
			// V3 错误响应只携带 code 与 message；未验签错误（Verified=false）
			// 仅作诊断信息，不能作为可信业务结果。
			logger.Error("query rejected", "code", apiErr.Code.String(), "verified", apiErr.Verified)
		default:
			logger.Error("query failed")
		}
		os.Exit(1)
	}

	// code=10000 仅表示查询成功；交易状态需单独判断，
	// TRADE_SUCCESS / TRADE_FINISHED 才表示买家付款成功。
	logger.Info("query succeeded",
		"trade_status", string(resp.Data.TradeStatus),
		"total_amount", resp.Data.TotalAmount)
}

// buildCertificateCredentials 从 PEM 文本装配证书模式凭证。
func buildCertificateCredentials(appCertPEM, appPrivateKeyPEM, rootCertPEM string, alipayCertPEMs []string) (auth.Credentials, error) {
	appCert, err := utils.ParseCertificatePEM(appCertPEM)
	if err != nil {
		return nil, fmt.Errorf("parse app certificate: %w", err)
	}
	rootCerts, err := utils.ParseRootCertificatesPEM(rootCertPEM)
	if err != nil {
		return nil, fmt.Errorf("parse root certificate: %w", err)
	}
	var alipayCerts []*x509.Certificate
	for _, pemText := range alipayCertPEMs {
		certs, err := utils.ParseCertificatesPEM(pemText)
		if err != nil {
			return nil, fmt.Errorf("parse alipay certificate: %w", err)
		}
		alipayCerts = append(alipayCerts, certs...)
	}
	appSigner, err := utils.ParsePrivateKeyPKCS8(utils.TrimKeyText(appPrivateKeyPEM))
	if err != nil {
		return nil, fmt.Errorf("parse app private key: %w", err)
	}
	return auth.NewCertificateCredentials(auth.CertificateConfig{
		AppSigner:          appSigner,
		AppCertificate:     appCert,
		AlipayCertificates: alipayCerts,
		RootCertificates:   rootCerts,
	})
}

func readRequiredFile(logger *slog.Logger, envName string) string {
	path := os.Getenv(envName)
	if path == "" {
		logger.Error("environment variable is required", "name", envName)
		os.Exit(2)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Error("cannot read certificate file", "name", envName)
		os.Exit(2)
	}
	return string(data)
}

func readAlipayCertificateFiles(paths string) ([]string, error) {
	var certificates []string
	for path := range strings.SplitSeq(paths, ",") {
		data, err := os.ReadFile(strings.TrimSpace(path))
		if err != nil {
			return nil, err
		}
		certificates = append(certificates, string(data))
	}
	return certificates, nil
}
