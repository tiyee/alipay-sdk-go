package protocol

import (
	"strings"
	"testing"
)

// TestBuildV3AuthString 核对 Authorization 头 authString 的拼装规则：
// app_id 固定在首位，证书 SN 仅在证书模式出现，nonce 与 timestamp 原样拼接。
func TestBuildV3AuthString(t *testing.T) {
	tests := []struct {
		name      string
		appID     string
		certSN    string
		nonce     string
		timestamp string
		want      string
	}{
		{
			name:      "public key mode",
			appID:     "2021000000000001",
			nonce:     "n-1",
			timestamp: "1764000000000",
			want:      "app_id=2021000000000001,nonce=n-1,timestamp=1764000000000",
		},
		{
			name:      "certificate mode",
			appID:     "2021000000000001",
			certSN:    "3e1fcbd2",
			nonce:     "n-2",
			timestamp: "1764000000001",
			want:      "app_id=2021000000000001,app_cert_sn=3e1fcbd2,nonce=n-2,timestamp=1764000000001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildV3AuthString(tt.appID, tt.certSN, tt.nonce, tt.timestamp)
			if got != tt.want {
				t.Fatalf("BuildV3AuthString() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestBuildV3SignContent 核对 V3 请求签名原文：每行以 \n 结尾，
// 请求体为空时保留空行，应用授权令牌非空才追加整行。
func TestBuildV3SignContent(t *testing.T) {
	const authString = "app_id=2021000000000001,nonce=n,timestamp=1"
	tests := []struct {
		name         string
		method       string
		uri          string
		body         string
		appAuthToken string
		want         string
	}{
		{
			name:   "post with json body",
			method: "POST",
			uri:    "/v3/alipay/trade/query",
			body:   `{"out_trade_no":"1"}`,
			want:   authString + "\nPOST\n/v3/alipay/trade/query\n" + `{"out_trade_no":"1"}` + "\n",
		},
		{
			name:   "get without body keeps empty line",
			method: "GET",
			uri:    "/v3/alipay/data/dataservice/bill/downloadurl/query?bill_type=trade",
			want:   authString + "\nGET\n/v3/alipay/data/dataservice/bill/downloadurl/query?bill_type=trade\n\n",
		},
		{
			name:         "app auth token appends trailing line",
			method:       "POST",
			uri:          "/v3/alipay/trade/pay",
			body:         `{"a":1}`,
			appAuthToken: "token-1",
			want:         authString + "\nPOST\n/v3/alipay/trade/pay\n" + `{"a":1}` + "\ntoken-1\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(BuildV3SignContent(authString, tt.method, tt.uri, tt.body, tt.appAuthToken))
			if got != tt.want {
				t.Fatalf("sign content mismatch:\n got  %q\n want %q", got, tt.want)
			}
		})
	}
}

// TestBuildV3VerifyContent 核对响应验签原文：timestamp、nonce、body 三行。
func TestBuildV3VerifyContent(t *testing.T) {
	got := string(BuildV3VerifyContent("1764000000000", "nonce-9", []byte(`{"ok":true}`)))
	want := "1764000000000\nnonce-9\n" + `{"ok":true}` + "\n"
	if got != want {
		t.Fatalf("verify content = %q, want %q", got, want)
	}
}

// TestParseV3Error 覆盖官方错误体契约：code 与 message 必填，
// links 可选；未知错误码按原值保留；缺失字段与非 JSON 体均报错。
func TestParseV3Error(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		want    V3Error
		wantErr bool
	}{
		{
			name: "gateway error",
			body: `{"code":"invalid-signature","message":"签名验证失败"}`,
			want: V3Error{Code: "invalid-signature", Message: "签名验证失败"},
		},
		{
			name: "business error with links",
			body: `{"code":"ACQ.TRADE_NOT_EXIST","message":"交易不存在","links":"https://opendocs.alipay.com"}`,
			want: V3Error{Code: "ACQ.TRADE_NOT_EXIST", Message: "交易不存在", Links: "https://opendocs.alipay.com"},
		},
		{
			name:    "missing code",
			body:    `{"message":"boom"}`,
			wantErr: true,
		},
		{
			name:    "missing message",
			body:    `{"code":"x"}`,
			wantErr: true,
		},
		{
			name:    "not json",
			body:    `oops`,
			wantErr: true,
		},
		{
			name:    "array body",
			body:    `[]`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseV3Error([]byte(tt.body))
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseV3Error(%s) expected error", tt.body)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseV3Error(%s): %v", tt.body, err)
			}
			if got != tt.want {
				t.Fatalf("ParseV3Error(%s) = %+v, want %+v", tt.body, got, tt.want)
			}
		})
	}
}

// TestV3ConstantsFix 锁定协议常量：方案名与响应头名称来自官方契约，
// 变更必须显式更新测试与文档。
func TestV3ConstantsFix(t *testing.T) {
	if V3AuthScheme != "ALIPAY-SHA256withRSA" {
		t.Fatalf("auth scheme = %q", V3AuthScheme)
	}
	for name, want := range map[string]string{
		V3HeaderAuthorization: "Authorization",
		V3HeaderRootCertSN:    "alipay-root-cert-sn",
		V3HeaderRequestID:     "alipay-request-id",
		V3HeaderAppAuthToken:  "alipay-app-auth-token",
		V3HeaderSignature:     "alipay-signature",
		V3HeaderCertSN:        "alipay-sn",
		V3HeaderTimestamp:     "alipay-timestamp",
		V3HeaderNonce:         "alipay-nonce",
	} {
		if name != want {
			t.Fatalf("header constant = %q, want %q", name, want)
		}
	}
	if !strings.HasPrefix(V3PathPrefix, "/v3/") {
		t.Fatalf("v3 path prefix = %q", V3PathPrefix)
	}
}
