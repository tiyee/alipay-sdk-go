package protocol

import (
	"net/url"
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
