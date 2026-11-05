// Command oauth 演示 V3 授权令牌换取：用户授权令牌与应用授权令牌。
//
// 链路：凭证装载 → Client（V3 REST）→ system.oauth.token（授权码换
// 用户访问令牌，或刷新）→ open.auth.token.app（应用授权码换 ISV 代调用
// 令牌）。令牌属于敏感凭证：示例只打印失效时间与用户标识，
// 不打印令牌内容；实际业务须安全保存，不得写入日志。
//
// 配置全部来自环境变量，示例不包含任何真实凭证：
//
//	ALIPAY_APP_ID           本应用 APPID（必填）
//	ALIPAY_APP_PRIVATE_KEY  应用私钥 PKCS#8 PEM 文件路径（必填）
//	ALIPAY_PUBLIC_KEY       支付宝公钥 PEM 文件路径（必填）
//	ALIPAY_ENV              production 或 sandbox（默认 production）
//	ALIPAY_AUTH_CODE        用户授权码（与 ALIPAY_REFRESH_TOKEN 二选一）
//	ALIPAY_REFRESH_TOKEN    用户授权刷新令牌
//	ALIPAY_APP_AUTH_CODE    应用授权码（可选，演示 ISV 代调用令牌换取）
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/services/open"
	"github.com/tiyee/alipay-sdk-go/services/system"
	"github.com/tiyee/alipay-sdk-go/utils"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	appID := os.Getenv("ALIPAY_APP_ID")
	authCode := os.Getenv("ALIPAY_AUTH_CODE")
	refreshToken := os.Getenv("ALIPAY_REFRESH_TOKEN")
	if appID == "" || (authCode == "" && refreshToken == "") {
		logger.Error("ALIPAY_APP_ID and one of ALIPAY_AUTH_CODE / ALIPAY_REFRESH_TOKEN are required")
		os.Exit(2)
	}

	client := newClient(logger)
	if client == nil {
		os.Exit(2)
	}
	ctx := context.Background()

	systemSvc, err := system.NewService(client)
	if err != nil {
		logger.Error("build system service", slog.String("err", err.Error()))
		os.Exit(2)
	}
	req := system.OAuthTokenRequest{GrantType: system.GrantTypeAuthorizationCode, Code: &authCode}
	if authCode == "" {
		req = system.OAuthTokenRequest{GrantType: system.GrantTypeRefreshToken, RefreshToken: &refreshToken}
	}
	resp, err := systemSvc.OAuthToken(ctx, req)
	if err != nil {
		// 网关业务错误（如 invalid-auth-token）以 *core.APIError 返回，
		// 错误信息只包含错误码；详情经字段访问，不落入日志。
		logger.Error("oauth token failed", slog.String("err", err.Error()))
		os.Exit(1)
	}
	// 令牌内容不打印；只输出元信息。
	logger.Info("user token acquired",
		slog.String("expires_in", deref(resp.Data.ExpiresIn)),
		slog.String("user_id", deref(resp.Data.UserID)),
		slog.String("open_id", deref(resp.Data.OpenID)))

	if appAuthCode := os.Getenv("ALIPAY_APP_AUTH_CODE"); appAuthCode != "" {
		openSvc, err := open.NewService(client)
		if err != nil {
			logger.Error("build open service", slog.String("err", err.Error()))
			os.Exit(2)
		}
		appResp, err := openSvc.AuthTokenApp(ctx, open.AuthTokenAppRequest{
			GrantType: open.GrantTypeAuthorizationCode,
			Code:      &appAuthCode,
		})
		if err != nil {
			logger.Error("app auth token failed", slog.String("err", err.Error()))
			os.Exit(1)
		}
		logger.Info("app token acquired",
			slog.String("auth_app_id", deref(appResp.Data.AuthAppID)),
			slog.String("user_id", deref(appResp.Data.UserID)))
		// 代调用示例：trade.Service 的方法可传入
		// core.WithAppAuthToken(*appResp.Data.AppAuthToken)。
	}
	fmt.Println("done")
}

func newClient(logger *slog.Logger) *core.Client {
	appID := os.Getenv("ALIPAY_APP_ID")
	privateKeyPEM, err := os.ReadFile(os.Getenv("ALIPAY_APP_PRIVATE_KEY"))
	if err != nil {
		logger.Error("read app private key", slog.String("err", err.Error()))
		return nil
	}
	publicKeyPEM, err := os.ReadFile(os.Getenv("ALIPAY_PUBLIC_KEY"))
	if err != nil {
		logger.Error("read alipay public key", slog.String("err", err.Error()))
		return nil
	}
	signer, err := utils.ParsePrivateKeyPKCS8(utils.TrimKeyText(string(privateKeyPEM)))
	if err != nil {
		logger.Error("parse app private key", slog.String("err", err.Error()))
		return nil
	}
	verifier, err := utils.ParsePublicKeyPKIX(utils.TrimKeyText(string(publicKeyPEM)))
	if err != nil {
		logger.Error("parse alipay public key", slog.String("err", err.Error()))
		return nil
	}
	credentials, err := auth.NewPublicKeyCredentials(auth.PublicKeyConfig{
		AppSigner:       signer,
		AlipayPublicKey: verifier,
	})
	if err != nil {
		logger.Error("build credentials", slog.String("err", err.Error()))
		return nil
	}
	// WithEnvironment 按环境成对设置 V3 REST 基地址与产物网关；
	// 本示例只调 REST 接口，无需单独覆盖网关地址。
	opts := []core.ClientOption{option.WithAppID(appID)}
	switch os.Getenv("ALIPAY_ENV") {
	case "", "production":
		opts = append(opts, option.WithEnvironment(core.EnvironmentProduction))
	case "sandbox":
		opts = append(opts, option.WithEnvironment(core.EnvironmentSandbox))
	default:
		logger.Error("ALIPAY_ENV must be production or sandbox")
		return nil
	}
	client, err := core.NewClient(credentials, opts...)
	if err != nil {
		logger.Error("build client", slog.String("err", err.Error()))
		return nil
	}
	return client
}

func deref(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
