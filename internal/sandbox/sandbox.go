// Package sandbox 为 API 示例测试提供显式启用的沙箱证书模式运行入口。
// 证书配置依据 https://opendocs.alipay.com/common/097jyh.md，
// 请求与响应复用 SDK 的 V3 Service、Client 和凭证实现。
package sandbox

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/tiyee/alipay-sdk-go/core"
	"github.com/tiyee/alipay-sdk-go/core/auth"
	"github.com/tiyee/alipay-sdk-go/core/option"
	"github.com/tiyee/alipay-sdk-go/internal/compat"
	"github.com/tiyee/alipay-sdk-go/internal/jsoncodec"
	"github.com/tiyee/alipay-sdk-go/utils"
)

// Run 使用真实沙箱 Client 调用与离线示例相同的 Service 方法。
// ALIPAY_SANDBOX=1 且 ALIPAY_SANDBOX_API 与 method 完全匹配时才执行；
// 其他接口跳过。ALIPAY_SANDBOX_REQUEST 指向按请求类型编码的 JSON 文件，
// 不复用离线示例的虚构业务参数，也不输出原始报文、令牌或错误内容。
func Run[Service, Req, Resp any](t *testing.T, method string,
	newService func(core.Executor) (*Service, error),
	call func(*Service, context.Context, Req, ...core.CallOption) (*core.Response[Resp], error),
) {
	t.Helper()
	enabled, err := selected(os.Getenv, method)
	if err != nil {
		t.Fatal(err)
	}
	if !enabled {
		t.Skip("sandbox disabled or a different ALIPAY_SANDBOX_API selected")
	}

	body, err := readFile(os.Getenv, "ALIPAY_SANDBOX_REQUEST")
	if err != nil {
		t.Fatal(err)
	}
	req, err := decodeRequest[Req](body)
	if err != nil {
		// JSON 解码错误可能包含授权码等请求字段值，不直接输出。
		t.Fatal("ALIPAY_SANDBOX_REQUEST must be a JSON object matching the Service request type; unknown fields are rejected")
	}
	client, err := newClient(os.Getenv, nil)
	if err != nil {
		t.Fatal(err)
	}
	service, err := newService(client)
	if err != nil {
		t.Fatalf("create Service failed (%T)", err)
	}
	var opts []core.CallOption
	if token := os.Getenv("ALIPAY_APP_AUTH_TOKEN"); token != "" {
		opts = append(opts, core.WithAppAuthToken(token))
	}
	resp, err := call(service, t.Context(), req, opts...)
	if err != nil {
		if apiErr, ok := compat.AsType[*core.APIError](err); ok {
			t.Fatalf("%s rejected: code=%s verified=%t", method, apiErr.Code.String(), apiErr.Verified)
		}
		t.Fatalf("%s failed (%T)", method, err)
	}
	if resp == nil || resp.Data == nil || resp.Result == nil || !resp.Result.SignatureVerified {
		t.Fatal("missing typed data or verified response")
	}
	t.Logf("%s: HTTP %d, signature verified; check the API-specific business status separately", method, resp.Result.HTTPStatus)
}

func selected(getenv func(string) string, method string) (bool, error) {
	switch getenv("ALIPAY_SANDBOX") {
	case "", "0":
		return false, nil
	case "1":
	default:
		return false, fmt.Errorf("ALIPAY_SANDBOX must be 0 or 1")
	}
	api := getenv("ALIPAY_SANDBOX_API")
	if api == "" {
		return false, fmt.Errorf("ALIPAY_SANDBOX_API must name one official API")
	}
	if api != method {
		return false, nil
	}
	if env := getenv("ALIPAY_ENV"); env != "" && env != "sandbox" {
		return false, fmt.Errorf("ALIPAY_ENV must be sandbox when running sandbox examples")
	}
	return true, nil
}

func decodeRequest[T any](body []byte) (T, error) {
	var req T
	trimmed := bytes.TrimSpace(body)
	if len(trimmed) == 0 || trimmed[0] != '{' {
		return req, fmt.Errorf("request must be a JSON object")
	}
	// 混合请求的 query/path 字段标记为 json:"-"，避免它们被写进网关请求体。
	// 沙箱输入文件可以按协议名提供这些字段，其余内容继续严格解码。
	value := reflect.ValueOf(&req).Elem()
	var fields map[string]json.RawMessage
	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			if !field.IsExported() || field.Tag.Get("json") != "-" {
				continue
			}
			name := field.Tag.Get("path")
			if name == "" {
				name, _, _ = strings.Cut(field.Tag.Get("query"), ",")
			}
			if name == "" || name == "-" {
				continue
			}
			if fields == nil {
				if err := jsoncodec.Unmarshal(trimmed, &fields); err != nil {
					return req, err
				}
			}
			if raw, ok := fields[name]; ok {
				if err := jsoncodec.Unmarshal(raw, value.Field(i).Addr().Interface()); err != nil {
					return req, err
				}
				delete(fields, name)
			}
		}
	}
	if fields != nil {
		rest, err := json.Marshal(fields)
		if err != nil {
			return req, err
		}
		trimmed = rest
	}
	err := jsoncodec.UnmarshalStrict(trimmed, &req)
	return req, err
}

func newClient(getenv func(string) string, httpClient *http.Client) (*core.Client, error) {
	appID := getenv("ALIPAY_APP_ID")
	if appID == "" {
		return nil, fmt.Errorf("ALIPAY_APP_ID is required")
	}
	privateKeyPEM, err := readFile(getenv, "ALIPAY_APP_PRIVATE_KEY")
	if err != nil {
		return nil, err
	}
	privateKey, err := utils.ParsePrivateKeyPKCS8(string(privateKeyPEM))
	if err != nil {
		return nil, fmt.Errorf("ALIPAY_APP_PRIVATE_KEY must contain a PKCS#8 RSA private key")
	}
	appCertPEM, err := readFile(getenv, "ALIPAY_APP_CERT")
	if err != nil {
		return nil, err
	}
	appCert, err := utils.ParseCertificatePEM(string(appCertPEM))
	if err != nil {
		return nil, fmt.Errorf("ALIPAY_APP_CERT must contain one application certificate")
	}
	rootPEM, err := readFile(getenv, "ALIPAY_ROOT_CERT")
	if err != nil {
		return nil, err
	}
	roots, err := utils.ParseRootCertificatesPEM(string(rootPEM))
	if err != nil {
		return nil, fmt.Errorf("ALIPAY_ROOT_CERT must contain root certificates")
	}
	var alipayCerts []*x509.Certificate
	for _, path := range strings.Split(getenv("ALIPAY_CERT"), ",") {
		path = strings.TrimSpace(path)
		if path == "" {
			return nil, fmt.Errorf("ALIPAY_CERT requires nonempty certificate file paths")
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("cannot read ALIPAY_CERT file")
		}
		certs, err := utils.ParseCertificatesPEM(string(data))
		if err != nil {
			return nil, fmt.Errorf("ALIPAY_CERT must contain Alipay public certificates")
		}
		alipayCerts = append(alipayCerts, certs...)
	}
	credentials, err := auth.NewCertificateCredentials(auth.CertificateConfig{
		AppSigner: privateKey, AppCertificate: appCert,
		AlipayCertificates: alipayCerts, RootCertificates: roots,
	})
	if err != nil {
		return nil, fmt.Errorf("invalid sandbox certificates: check validity, private key match and root trust")
	}
	opts := []core.ClientOption{
		option.WithAppID(appID),
		option.WithEnvironment(core.EnvironmentSandbox),
	}
	if httpClient != nil {
		opts = append(opts, option.WithHTTPClient(httpClient))
	}
	client, err := core.NewClient(credentials, opts...)
	if err != nil {
		return nil, fmt.Errorf("invalid sandbox Client configuration")
	}
	return client, nil
}

func readFile(getenv func(string) string, name string) ([]byte, error) {
	path := getenv(name)
	if path == "" {
		return nil, fmt.Errorf("%s must contain a file path", name)
	}
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read %s file", name)
	}
	return body, nil
}
