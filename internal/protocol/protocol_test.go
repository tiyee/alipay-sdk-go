package protocol

import (
	"encoding/json/jsontext"
	"net/url"
	"strings"
	"testing"
)

func urlValues(m map[string][]string) url.Values {
	values := url.Values{}
	for key, vs := range m {
		values[key] = vs
	}
	return values
}

func TestBuildSignContent(t *testing.T) {
	tests := []struct {
		name   string
		values map[string][]string
		want   string
	}{
		{
			name:   "按 ASCII 升序拼接，值不 URL 编码",
			values: map[string][]string{"method": {"alipay.trade.query"}, "app_id": {"2021000000000001"}},
			want:   "app_id=2021000000000001&method=alipay.trade.query",
		},
		{
			name:   "剔除 sign 参数",
			values: map[string][]string{"app_id": {"1"}, "sign": {"should-exclude"}},
			want:   "app_id=1",
		},
		{
			name:   "剔除空值参数",
			values: map[string][]string{"app_id": {"1"}, "notify_url": {""}},
			want:   "app_id=1",
		},
		{
			name:   "中文与特殊字符保持原文",
			values: map[string][]string{"subject": {"手机-商品/详情"}, "app_id": {"1"}},
			want:   "app_id=1&subject=手机-商品/详情",
		},
		{
			name:   "biz_content 整体参与签名",
			values: map[string][]string{"biz_content": {`{"a":"1","b":"2"}`}, "app_id": {"1"}},
			want:   `app_id=1&biz_content={"a":"1","b":"2"}`,
		},
		{
			name:   "下划线键排序在连字符之后",
			values: map[string][]string{"app_id": {"1"}, "a-b": {"2"}, "a_b": {"3"}},
			want:   "a-b=2&a_b=3&app_id=1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := urlValues(tt.values)
			if got := BuildSignContent(values); got != tt.want {
				t.Fatalf("BuildSignContent() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestBuildNotifySignContent 验证通知待验签串与请求待签名串的规则差异：
// 仅剔除 exclude 列出的参数，空值参数保留。
func TestBuildNotifySignContent(t *testing.T) {
	tests := []struct {
		name    string
		values  map[string][]string
		exclude []string
		want    string
	}{
		{
			name:    "剔除 sign 与 sign_type，其余按升序拼接",
			values:  map[string][]string{"sign": {"sig"}, "sign_type": {"RSA2"}, "trade_status": {"TRADE_SUCCESS"}, "app_id": {"1"}},
			exclude: []string{"sign", "sign_type"},
			want:    "app_id=1&trade_status=TRADE_SUCCESS",
		},
		{
			name:    "空值参数保留，与请求签名规则不同",
			values:  map[string][]string{"app_id": {"1"}, "body": {""}},
			exclude: []string{"sign", "sign_type"},
			want:    "app_id=1&body=",
		},
		{
			name:    "解码后的值保持原文，不 URL 编码",
			values:  map[string][]string{"subject": {"测试 a+b&c"}, "app_id": {"1"}},
			exclude: []string{"sign", "sign_type"},
			want:    "app_id=1&subject=测试 a+b&c",
		},
		{
			name:    "exclude 缺省时不剔除任何参数",
			values:  map[string][]string{"sign_type": {"RSA2"}, "app_id": {"1"}},
			exclude: nil,
			want:    "app_id=1&sign_type=RSA2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := urlValues(tt.values)
			if got := BuildNotifySignContent(values, tt.exclude...); got != tt.want {
				t.Fatalf("BuildNotifySignContent() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestIsProtectedParam(t *testing.T) {
	for _, key := range []string{"app_id", "method", "sign", "sign_type", "biz_content", "timestamp", "notify_url", "app_auth_token"} {
		if !IsProtectedParam(key) {
			t.Errorf("IsProtectedParam(%q) = false, want true", key)
		}
	}
	for _, key := range []string{"out_trade_no", "total_amount", "subject", "custom_field"} {
		if IsProtectedParam(key) {
			t.Errorf("IsProtectedParam(%q) = true, want false", key)
		}
	}
}

func TestExtractRawResponse(t *testing.T) {
	node := `{"code":"10000","msg":"Success","trade_no":"2026010122001400001"}`
	body := `{"alipay_trade_query_response":` + node + `,"sign":"c2lnbmF0dXJlLXZhbHVl"}`
	raw, err := ExtractRawResponse([]byte(body), "alipay_trade_query_response")
	if err != nil {
		t.Fatalf("ExtractRawResponse() error = %v", err)
	}
	if string(raw.Node) != node {
		t.Fatalf("Node = %q, want original bytes %q", raw.Node, node)
	}
	if raw.Sign != "c2lnbmF0dXJlLXZhbHVl" {
		t.Fatalf("Sign = %q", raw.Sign)
	}
	if raw.ErrorResponse != nil {
		t.Fatalf("ErrorResponse = %v, want nil", raw.ErrorResponse)
	}
}

func TestExtractRawResponsePreservesOriginalEncoding(t *testing.T) {
	// \/ 转义与内部空格必须原样保留，重新序列化会导致验签失败。
	node := `{"code":"10000" , "url":"https:\/\/example.com\/pay"}`
	body := `{"alipay_trade_query_response":` + node + `,"sign":"x"}`
	raw, err := ExtractRawResponse([]byte(body), "alipay_trade_query_response")
	if err != nil {
		t.Fatalf("ExtractRawResponse() error = %v", err)
	}
	if string(raw.Node) != node {
		t.Fatalf("Node = %q, want %q", raw.Node, node)
	}
}

func TestExtractRawResponseErrorResponse(t *testing.T) {
	body := `{"error_response":{"code":"40002","msg":"Invalid Arguments","sub_code":"isv.invalid-signature"},"sign":null}`
	raw, err := ExtractRawResponse([]byte(body), "alipay_trade_query_response")
	if err != nil {
		t.Fatalf("ExtractRawResponse() error = %v", err)
	}
	if raw.ErrorResponse == nil {
		t.Fatal("ErrorResponse is nil, want parsed")
	}
	if !strings.Contains(string(raw.ErrorResponse), "isv.invalid-signature") {
		t.Fatalf("ErrorResponse = %q", raw.ErrorResponse)
	}
	if raw.Node != nil {
		t.Fatalf("Node = %q, want nil", raw.Node)
	}
}

func TestExtractRawResponseErrors(t *testing.T) {
	tests := []struct {
		name string
		body string
		key  string
	}{
		{name: "非对象顶层", body: `[1,2]`, key: "alipay_trade_query_response"},
		{name: "节点缺失", body: `{"other":{"code":"10000"}}`, key: "alipay_trade_query_response"},
		{name: "重复业务节点", body: `{"alipay_trade_query_response":{"code":"10000"},"alipay_trade_query_response":{"code":"10000"}}`, key: "alipay_trade_query_response"},
		{name: "重复 sign", body: `{"alipay_trade_query_response":{"code":"10000"},"sign":"a","sign":"b"}`, key: "alipay_trade_query_response"},
		{name: "尾部内容", body: `{"alipay_trade_query_response":{"code":"10000"}} garbage`, key: "alipay_trade_query_response"},
		{name: "非法 JSON", body: `not-json`, key: "alipay_trade_query_response"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ExtractRawResponse([]byte(tt.body), tt.key); err == nil {
				t.Fatalf("ExtractRawResponse() expect error, got nil")
			}
		})
	}
}

func TestParseGatewayErrorAndCode(t *testing.T) {
	node := jsontext.Value(`{"code":"40004","msg":"Business Failed","sub_code":"ACQ.TRADE_NOT_EXIST","sub_msg":"交易不存在"}`)
	gw, err := ParseGatewayError(node)
	if err != nil {
		t.Fatalf("ParseGatewayError() error = %v", err)
	}
	if gw.Code != "40004" || gw.SubCode != "ACQ.TRADE_NOT_EXIST" || gw.SubMsg != "交易不存在" {
		t.Fatalf("gateway error = %+v", gw)
	}
	code, err := ParseResponseCode(node)
	if err != nil || code != "40004" {
		t.Fatalf("ParseResponseCode() = %q, %v", code, err)
	}
}

func TestParseGatewayErrorMissingCode(t *testing.T) {
	if _, err := ParseGatewayError(jsontext.Value(`{"msg":"no code"}`)); err == nil {
		t.Fatal("expect error for missing code")
	}
	if _, err := ParseResponseCode(jsontext.Value(`{"msg":"no code"}`)); err == nil {
		t.Fatal("expect error for missing code")
	}
}
